# [](#alter-table)ALTER TABLE

Updates the specified table.

## [](#alter-table-add-column)ALTER TABLE ADD COLUMN

Adds a column to an existing table.

### [](#syntax)Syntax

```
ALTER TABLE <table> ADD COLUMN [IF NOT EXISTS] <column_name> <column_type> [NULL | NOT NULL] [DEFAULT <expression>]
```

### [](#parameters)Parameters

Parameter Description `<table>` Name of the table to which to add the column. `<column_name>` An identifier that specifies the name of the column that will be added to the table. `<column_type>` Specifies the data type for the column. `IF NOT EXISTS` If specified, this clause prevents an error message that would occur if the column already exists in the table.

### [](#column-constraints--default-expression)Column constraints &amp; default expression

`DEFAULT <expression>` Determines the value that will be used for the column when this column is omitted in an `INSERT` statement. It also determines the value that is used for the rows that were inserted before the column was added. `NULL` | `NOT NULL` Determines if the column may or may not contain `NULL`s.

### [](#limitations)Limitations

The query can only be executed under the following conditions:

- Only for managed tables created on Firebolt version 4.10 or higher (external tables are not supported).
- The table must not have any dependent views.
- The default expression must contain only literals or functions `CURRENT_DATE()`, `LOCALTIMESTAMP()`, `CURRENT_TIMESTAMP()`, and `NOW()`.

## [](#alter-table-drop-partition)ALTER TABLE DROP PARTITION

Use to delete a partition from a fact or dimension table.

Dropping a partition deletes the partition and the data stored in that partition.

### [](#syntax-1)Syntax

```
ALTER TABLE <table> DROP PARTITION <value1>[,...<value2]
```

### [](#parameters-1)Parameters

Parameter Description `<table>` Name of the table from which to drop the partition. `<value1>[,...<value2>]` An ordered set of one or more values corresponding to the partition key definition. This specifies the partition to drop. When dropping partitions with composite keys (more than one key value), specify all key values in the same order as they were defined. Only partitions with values that match the entire composite key are dropped.

### [](#examples)Examples

See the examples in [Working with partitions](/Overview/indexes/using-indexes.html#partitions-in-tables).

## [](#alter-table-owner-to)ALTER TABLE OWNER TO

Change the owner of a table. The current owner of a table can be viewed in the `information_schema.tables` view on `table_owner` column.

check [ownership](/Guides/security/ownership.html) page for more info.

### [](#syntax-2)Syntax

```
ALTER TABLE <table> OWNER TO <user>
```

### [](#parameters-2)Parameters

Parameter Description `<table>` Name of the table to change the owner of. `<user>` The new owner of the table.

## [](#alter-table-rename-to)ALTER TABLE RENAME TO

Renames a table.

### [](#syntax-3)Syntax

```
ALTER TABLE [IF EXISTS] <table_name> RENAME TO <new_table_name>
```

### [](#parameters-3)Parameters

Parameter Description `<table_name>` The name of the table to rename. `<new_table_name>` The new name of the table.

### [](#limitations-1)Limitations

The query can only be executed under the following conditions:

- Only for managed tables created on Firebolt version 4.10 or higher (external tables are not supported).
- The table must not have any dependent views.
- Renaming tables across schemas and databases is not supported. Consider using [CREATE TABLE CLONE](/sql_reference/commands/data-definition/create-table-clone.html)