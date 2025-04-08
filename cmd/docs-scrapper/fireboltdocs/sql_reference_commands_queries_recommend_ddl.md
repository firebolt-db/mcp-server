# [](#recommend-ddl)RECOMMEND DDL

`CALL recommend_ddl` can help you optimize schema configurations to enhance query performance by early data pruning. The statement finds [primary indexes](/Overview/indexes/primary-index.html) (PIs) and [partition](/Overview/indexes/using-indexes.html#partitions-in-tables) recommendations for the specified table tailored to the given workload.

## [](#syntax)Syntax

```
CALL recommend_ddl(<table_name>, (<select_statement>))
```

## [](#parameters)Parameters

Parameter Description `<table_name>` The name of the table for which primary indexes and partition keys should be recommended. `<select_statement>` [SELECT](/sql_reference/commands/queries/select.html) statement that returns the workload that the DDL recommendation is based on. The `<select_statement>` must return exactly one column of type `TEXT`.

## [](#example)Example

The example below demonstrates retrieving schema recommendation using the `CALL recommend_ddl` statement for a table named `lineitem` tailored to the workload in the query history of the past week.

```
CALL recommend_ddl(lineitem, (SELECT query_text FROM information_schema.engine_query_history WHERE start_time > NOW() - INTERVAL '1 week'))
```

The `<select_statement>` returns exactly one column of type `TEXT` containing the SQL statements that the `CALL recommend_ddl` command should analyze.

**Returns:**

recommended\_partition\_key recommended\_primary\_index average\_pruning\_improvement analyzed\_queries DATE\_TRUNC(‘month’, l\_orderdate) l\_shipmode, l\_returnflag, l\_shipinstruct 0.42 393

The `CALL recommend_ddl` results indicate that the amount of bytes scanned can be decreased by up to 42% by configuring `PRIMARY INDEX l_shipmode, l_returnflag, l_shipinstruct` and `PARTITION BY DATE_TRUNC('month', l_orderdate)`. The statement analyzed 393 queries that scanned the `lineitem` table and applied filters to any of the `lineitem` columns.

## [](#quick-setup)Quick Setup

The following steps will guide you to achieve great query performance within the first few minutes after joining Firebolt. First, create a table without any primary index and parition key configurations.

```
CREATE TABLE <table_name>(
    ...
);
```

Next, load a workload that you want to run on this table from S3 into Firebolt utilizing [COPY INTO]().

```
COPY INTO workload_table FROM 's3://bucket/workload/' with ... 
```

Now you can use the `CALL recommend_ddl` command to find primary index and parition key configurations.

```
CALL recommend_ddl(<table_name>, (select * from workload_table));
```

Finally, recreate the table with the recommended primary index and parition key configurations and ingest the data into this table.

```
DROP TABLE <table_name>;
CREATE TABLE <table_name>(
    ...
)
PRIMARY INDEX ...
PARTITION BY ..;
```

## [](#under-the-hood)Under The Hood

Primary index and partition key configurations are chosen to maximize pruning potential and thus reduce query runtime time. Columns with selective filters and low cardinality are suggested as primary index and partition key columns. Recommendations can be run on empty tables as well as on tables with production data. The more queries executed on a populated table, the better the recommendations become. If additionally the workload of a table changes over time, it can be beneficial to run the `CALL recommend_ddl` command periodically to check for better table configurations.