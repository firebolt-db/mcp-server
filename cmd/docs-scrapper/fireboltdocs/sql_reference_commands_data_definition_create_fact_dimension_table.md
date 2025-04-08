# [](#create-table)CREATE TABLE

Creates a new table in the current database.

- [CREATE TABLE](#create-table)
  
  - [Column constraints and the default expression](#column-constraints-and-the-default-expression)
    
    - [Example: Creating a table with `NULL` and `NOT NULL` values](#example-creating-a-table-with-null-and-not-null-values)
    - [PRIMARY INDEX](#primary-index)
    - [PARTITION BY](#partition-by)
    - [Table type](#table-type)
  - [Related functions](#related-functions)

## [](#syntax)Syntax

```
CREATE [FACT|DIMENSION] TABLE [IF NOT EXISTS] <table_name>
(
    <column_name> <column_type> [constraints]
    [, <column_name> <column_type> [constraints]
    [, ...n]]
)
[PRIMARY INDEX <column_name>[, <column_name>[, ...k]]]
[PARTITION BY <column_name>[, <column_name>[, ...m]]]
[WITH ( <storage_parameter> = <storage_parameter_value>[, <storage_parameter> = <storage_parameter_value>[, ...p]] ) ]
```

## [](#parameters)Parameters

Parameter Description `<table_name>` An identifier that specifies the name of the table. This name should be unique within the database. `<column_name>` An identifier that specifies the name of the column. This name should be unique within the table. `<column_type>` Specifies the data type for the column. `<storage_parameter>` The name of a [storage parameter](#storage-parameters) for controlling behaviors related to storage and indexes. `<storage_parameter_value>` The value assigned to a `<storage_parameter>`.

All identifiers are case-insensitive unless enclosed in double-quotes. For more information, see [Object identifiers](/Reference/object-identifiers.html).

## [](#column-constraints-and-the-default-expression)Column constraints and the default expression

Firebolt supports the following column constraints:

```
<column_name> <column_type> [NULL | NOT NULL] [DEFAULT <expression>]
```

Constraint Description Default value `DEFAULT <expression>` Determines the default value used when no value is provided, instead of inserting a `NULL` value.   `NULL` | `NOT NULL` Determines if the column may or may not contain `NULL` values. `NOT NULL`

Nullable columns cannot be used in Firebolt primary or aggregating indexes. Additionally, only literals and the following functions are supported in default expressions: [CURRENT\_DATE](/sql_reference/functions-reference/date-and-time/current-date.html), [LOCALTIMESTAMP](/sql_reference/functions-reference/date-and-time/localtimestamp.html), [CURRENT\_TIMESTAMP](/sql_reference/functions-reference/date-and-time/current-timestamptz.html), and NOW, the alias for CURRENT\_TIMESTAMP.

### [](#example-creating-a-table-with-null-and-not-null-values)Example: Creating a table with `NULL` and `NOT NULL` values

The following example illustrates different use cases for column definitions and `INSERT` statements:

- An **Explicit** `NULL` insert – a direct insertion of a `NULL` value into a particular column.
- An **Implicit** `NULL` insert – an `INSERT` statement with missing values for a particular column.

The following example creates a fact table `t1` with five columns, specifying if each column can contain `NULL` values, their default values, and a primary index on `col2`:

```
CREATE FACT TABLE t1
(
    col1 INTEGER  NULL,
    col2 INTEGER  NOT NULL,
    col3 INTEGER  NULL DEFAULT 1,
    col4 INTEGER  NOT NULL DEFAULT 1,
    col5 TEXT
)
PRIMARY INDEX col2;
```

After creating a table, you can manipulate the values using different `INSERT` statements, as shown in the following examples:

INSERT statement Results and explanation `INSERT INTO t1 VALUES (1,1,1,1,1)` This code example inserts `1` into each column. `INSERT INTO t1 VALUES (NULL,1,1,1,1)` This code example explicitly inserts a `NULL` value into `col1`. Because `col1` can contain `NULL` values, this operation is successful. `INSERT INTO t1 (col2,col3,col4,col5) VALUES (1,1,1,1)` This code example shows both explicit and implicit `INSERT` statements. Because `col1` has no value specified, and lacks a default expression, it is implicitly set to `NULL`. `INSERT INTO t1 VALUES (1,NULL,1,1,1)`

`INSERT INTO t1 (col1,col3,col4,col5) VALUES (1,1,1,1)` This code example shows how a **null mismatch** error is generated. Because `col2` is defined as `NOT NULL` with no default expression, both `INSERT` statements implicitly try to insert `NULL` values into `col2`, and generate “null mismatch” events. `INSERT INTO t1 VALUES (1,1,NULL,1,1)` This code example explicitly inserts a `NULL` value into `col3`. Because `col3` is defined as `NULL DEFAULT 1`, the operation is successful. `INSERT INTO t1 (col1,col2,col4,col5) VALUES (1,1,1,1)` By not specifying a value for `col3`, this code example implicitly inserts the default value `1` into `col3`, which is defined as `NULL DEFAULT 1`. `INSERT INTO t1 VALUES (1,1,1,NULL,1)` This code example shows how another **null mismatch** error is generated. Because `col4` is defined as `NOT NULL DEFAULT 1`, the explicit insertion of a `NULL` value violates the `NOT NULL` constraint, and results in a null mismatch event. `INSERT INTO t1 (col1,col2,col3,col5) VALUES (1,1,1,1)` This code example shows how omitting a value in an implicit insert invokes the default value `1` for `col4`. `INSERT INTO t1 VALUES (1,1,1,1,NULL)`

`INSERT INTO t1 (col1,col2,col3,col4) VALUES (1,1,1,1)` This code example shows how explicit and implicit inserts cause a **null mismatch** error. Because `col5` was neither defined with a default expression nor allowed to contain `NULL` values, Firebolt treats `col5` as `NOT NULL DEFAULT NULL`. Both `INSERT` statements attempt to insert a `NULL` value into a `NOT NULL TEXT` column, invoking in a null mismatch event.

### [](#primary-index)PRIMARY INDEX

The `PRIMARY INDEX` is an optional sparse index that sorts and organizes data based on the indexed field as it is ingested, without affecting data scan performance. For more information, see [Primary index](/Overview/indexes/primary-index.html).

#### [](#syntax-1)Syntax

```
PRIMARY INDEX <column_name>[, <column_name>[, ...n]]
```

The following table describes the primary index parameters:

Parameter. Description Mandatory? `<column_name>` Specifies the name of the column in the Firebolt table which composes the index. At least one column is required. Y

### [](#partition-by)PARTITION BY

The `PARTITION BY` clause defines one or more columns that determine how the table is divided into physical parts. These columns serve as the partition key and cannot allow `NULL` values. When multiple columns are used as the partition key, the combination of all of these columns define the partition boundaries.

```
PARTITION BY <column_name>[, <column_name>[, ...n]]
```

The following subset of SQL functions can be used in `PARTITION BY` expressions:

- [TO\_YYYYMM](/sql_reference/functions-reference/date-and-time/to-yyyymm.html)
- [TO\_YYYYMMDD](/sql_reference/functions-reference/date-and-time/to-yyyymmdd.html)
- [EXTRACT](/sql_reference/functions-reference/date-and-time/extract.html)`(year|month|day|hour from <column_name>)`
- [DATE\_TRUNC](/sql_reference/functions-reference/date-and-time/date-trunc.html)

For more information, see [Working with partitions](/Overview/indexes/using-indexes.html#partitions-in-tables).

### [](#table-type)Table type

Firebolt supports two types of [tables](/Overview/indexes/using-indexes.html#tables):

- `FACT` table - the data is distributed across all nodes of the engine.
- `DIMENSION` table - the entire table is replicated in every node of the engine.

The [CREATE TABLE](/sql_reference/commands/data-definition/create-fact-dimension-table.html) command defaults to a `FACT` table. `DIMENSION` tables are ideal for relatively small tables, up to tens of gigabytes, that are used in joins with `FACT` tables.

## [](#storage-parameters)Storage Parameters

Storage parameters are specified in the optional `WITH (...)` clause as comma separated `<storage_parameter> = <storage_parameter_value>` assignments.

Storage Parameter Description `<index_granularity>` The maximum number of rows in each granule. `<storage_parameter_value>` must be a power of 2 between 128 and 8192. The default value is 8192. For more information, see [Index granularity](/Overview/indexes/primary-index.html#advanced-option-index-granularity).

All identifiers are case-insensitive unless enclosed in double-quotes. For more information, see [Object identifiers](/Reference/object-identifiers.html).

## [](#related-functions)Related functions

Firebolt also supports the following related functions:

- [CREATE TABLE AS SELECT (CTAS)](/sql_reference/commands/data-definition/create-fact-dimension-table-as-select.html) – - Creates a table and loads data into it based on a `SELECT` query.
- [CREATE TABLE CLONE](/sql_reference/commands/data-definition/) – Creates a table that is a copy of an existing table in the database.