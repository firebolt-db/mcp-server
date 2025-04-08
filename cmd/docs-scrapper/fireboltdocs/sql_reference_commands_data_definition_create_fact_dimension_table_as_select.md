# [](#create-fact-or-dimension-tableas-select)CREATE FACT or DIMENSION TABLE…AS SELECT

Creates a table and loads data into it based on the [SELECT](/sql_reference/commands/queries/select.html) query. The table column names and types are automatically inferred based on the output columns of the [SELECT](/sql_reference/commands/queries/select.html). When specifying explicit column names, those override the column names inferred from `SELECT`.

## [](#syntax)Syntax

Fact table:

```
CREATE FACT TABLE [IF NOT EXISTS] <table_name>
[(<column_name>[, ...n] )]
PRIMARY INDEX <column_name>[, <column_name>[, ...n]]
[PARTITION BY <column_name>[, <column_name>[, ...n]]]
AS <select_query>
```

Dimension table:

```
CREATE DIMENSION TABLE [IF NOT EXISTS] <table_name>
[(<column_name>[, ...n] )]
[PRIMARY INDEX <column_name>[, <column_name>[, ...n]]]
AS <select_query>
```

## [](#parameters)Parameters

Parameter Description `<table_name>` An identifier that specifies the name of the external table. This name should be unique within the database. `<column_name>` An identifier that specifies the name of the column. This name should be unique within the table. `<select_query`&gt; Any valid select query.

## [](#remarks)Remarks

Unlike a traditional `CREATE` statement, the primary index and partition definition must come *before* the `AS` clause.