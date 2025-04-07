# [](#vacuum)VACUUM

Optimizes tablets for query performance.

`VACUUM` improves query efficiency by restructuring tablets for optimal performance. DML operations such as [DELETE](/sql_reference/commands/data-management/delete.html), [UPDATE](/sql_reference/commands/data-management/update.html), [INSERT](/sql_reference/commands/data-management/insert.html), and [COPY FROM](/sql_reference/commands/data-management/copy-from.html) might create tablets that are not optimally sized. Suboptimal tablets occur because DML efficiently utilizes resources in proportion to the cardinality of the data being inserted. In addition to standard SQL operations, tuples that are deleted by an update are not always physically removed from their table; they remain present until a `VACUUM` is finished operating. In other words, tablets are not necessarily optimal for running queries; therefore, it’s necessary to run `VACUUM` periodically, especially on frequently updated tables.

Any engine that processes a DML operation automatically assesses the health of tables’ data layout and runs the `VACUUM` command when necessary to maintain the underlying table health. The fragmentation ratio—the ratio of rows marked for deletion to the total number of rows—can be monitored using the `information_schema.tables view`. You can also run `VACUUM` manually using the syntax and options described below.

## [](#syntax)Syntax

```
VACUUM [ (option_name = option_value) ] <table|aggregating index>
```

Where `<table|aggregating index>` is the name of the table or aggregating index to be optimized.

## [](#options)Options

Option name Option value and description `INDEXES` `FULL` — (Default) Specifies whether to apply optimizations to both the table and all its aggregating indexes.  
`INCREMENTAL` — Similar to `FULL`, but will apply incremental optimizations to aggregating indexes, instead of complete reevaluation.  
`NONE` — Optimizes only the table. `MAX_CONCURRENCY` `<Number>` — The maximum number of concurrent jobs to use during optimization.

## [](#examples)Examples

**Optimize a table and its aggregating indexes**

Optimizing a table along with its aggregating indexes ensures that both the data and aggregating indexes remain efficient, reducing query latency and improving overall performance.  
The following code example optimizes the `games` table and all its aggregating indexes:

```
VACUUM games;
```

The example above performs a complete rebuild of related aggregating indexes. Alternatively, you can specify to apply incremental optimization, which still improves index layout while using fewer resources, though it may not achieve optimal layout:

```
VACUUM (INDEXES = INCREMENTAL) games;
```

**Optimize a table without its indexes**

If you need to optimize a table without including its aggregating indexes to reduce resource usage, or retain efficient indexes, you can optimize only the table to prevent unnecessary computations.  
The following code example optimizes the `players` table without updating its aggregating indexes:

```
VACUUM (INDEXES = NONE) players;
```

Optimize table named `players`, using a single concurrent stream.

```
VACUUM (MAX_CONCURRENCY = 1) players;
```

### [](#usage-notes)Usage Notes

The following are considerations for running the `VACUUM` command:

- **What happens during VACUUM**  
  `VACUUM` analyzes the tablets, selects the ones that are too small or have too many deleted rows, and produces new versions that are optimized for query execution for both tablets and Aggregate Indexes.  
  `VACUUM` runs as a non-blocking process, alongside other user-initiated operations. Consequently, some changes performed by `VACUUM` may conflict with mutations run by the user. If `VACUUM` and a user mutation modify the same data, the first committed operation takes precedence; see [Transactions and concurrency](/Overview/data-management.html#transactions-and-concurrency) for more details. This means that applications that run mutations in parallel with `VACUUM` should gracefully handle transaction conflicts. It also means that benefits of the `VACUUM` may be diminished by a mutation that committed data first.
- **Space and performance considerations**  
  Users must be aware that `VACUUM` consumes both compute and storage resources.  
  `VACUUM` can consume a considerable amount of compute resources depending on the table size, number of tablets, and number of mutations in the table.  
  `VACUUM` parallelizes its work into multiple concurrent streams, based on the number of CPU cores. While this can be beneficial for the speed of the operation, each stream consumes memory and CPU resources. Use the `MAX_CONCURRENCY` option to limit the number of concurrent streams.  
  `VACUUM` produces optimized versions of the data, while leaving behind older versions subject to the garbage collection (GC) process. These older tablets will continue to consume storage space until the GC process completes the clean-up.  
  If users would like to have precise control over `VACUUM`, it may be preferable to run on a dedicated engine that could be sized and run just for `VACUUM` operations. With `VACUUM` running on a dedicated engine, it would not conflict with other queries’ execution and cache resources, and would provide operational separation from other scenarios.  
  `VACUUM` may introduce a performance penalty as the newly created optimized tablets need to be synchronized with other engines operating on the same table(s).
- **Automatic scheduling**  
  You can enable automatic scheduling of processes such as `VACUUM` by integrating with external tools. Please see section [Integrate with Firebolt](/Guides/integrations/integrations.html) for more detail on our current support for these tools.

### [](#example-with-measuring-the-performance-impact-of-vacuum)Example with measuring the performance impact of VACUUM

Over time, operations such as `INSERT`, `DELETE`, and `UPDATE` can create suboptimal tablets that decrease query performance. The `VACUUM` command restructures these tablets by removing deleted rows and optimizing storage, leading to faster queries.

This example demonstrates the impact of `VACUUM` by:

1. Creating a large table with 10 million rows.
2. Deleting 90% of the rows, leaving behind fragmented data.
3. Running a query before and after `VACUUM` to compare run times.

The following code example loads data from a CSV file in an Amazon S3 bucket into the `tutorial_vacuum` table with headers:

```
COPY tutorial_vacuum
FROM 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv'
WITH HEADER=TRUE;

INSERT INTO tutorial_vacuum
SELECT a.* FROM tutorial_vacuum a, GENERATE_SERIES(1, 1000000); -- This may run a couple of minutes
```

Script above loads 10 rows from S3 csv file; after that it inserts into the same table the cross product of the 10 inserted rows with 1 million integers. Next, the following code example deletes all rows from the `tutorial_vacuum` table where the `LevelID` value is greater than 1, resulting in about 900,000 deleted rows:

```
DELETE FROM tutorial_vacuum WHERE "LevelID" > 1;
```

The following code example runs a query computing checksum of the entire `tutorial_vacuum` table before and after performing `VACUUM`, allowing a comparison of query performance and efficiency improvements after optimization.

```
SELECT hash_agg(*) FROM tutorial_vacuum;
VACUUM tutorial_vacuum;
SELECT hash_agg(*) FROM tutorial_vacuum;
```

In the previous code example, the first `SELECT` is run on data with many deleted rows, while the second runs after `VACUUM`, benefiting from it. The following query history shows the performance benefit of the `VACUUM` operation:

NO STATEMENT STATUS DURATION 1 `SELECT hash_agg(*) FROM tutorial_vacuum;` Success 4.43 s 2 `VACUUM tutorial_vacuum;` Success 17.53 s 3 `SELECT hash_agg(*) FROM tutorial_vacuum;` Success 0.82 s

Note that the initial `SELECT` query ran for over 4 seconds, while the identical `SELECT` query ran for under a second, after running `VACUUM`.