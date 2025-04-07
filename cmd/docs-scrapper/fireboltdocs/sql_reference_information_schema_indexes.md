# [](#information-schema-for-indexes)Information schema for indexes

You can use the `information_schema.indexes` view to return information about each index in a database. The view is available for each database and contains one row for each index in the database. You can use a `SELECT` query to return information about each index.

In order to view index information, you need the USAGE privilege on both the [schema](/Overview/Security/Role-Based%20Access%20Control/database-permissions/schema-permissions.html#schema-level-privileges) and the [database](/Overview/Security/Role-Based%20Access%20Control/database-permissions/#database-level-privileges). You also need ownership of the table or the necessary [table-level privileges](/Overview/Security/Role-Based%20Access%20Control/database-permissions/table-permissions.html#table-level-privileges) required for the intended action.

The following query returns all aggregating indexes defined within the current database.

```
SELECT
  *
FROM
  information_schema.indexes
WHERE
  index_type='aggregating`;
```

## [](#columns-in-information_schemaindexes)Columns in information\_schema.indexes

Each row has the following columns with information about the database.

Column Name Data Type Description table\_catalog TEXT Name of the catalog. Firebolt provides a single ‘default’ catalog. table\_schema TEXT Name of the database. table\_name TEXT The name of the table for which the index is defined. index\_name TEXT The name defined for the index. index\_type TEXT One of either `primary` or `aggregating`. index\_owner TEXT The owner of the table, which is the owner of the index. index\_definition TEXT The part of the index statement that specifies the columns and any aggregations included in the index. compressed\_bytes BIGINT The compressed size of the index, in bytes. uncompressed\_bytes BIGINT The uncompressed size of the index, in bytes. number\_of\_tablets BIGINT The number of tablets in the index. created TIMESTAMPTZ Time that the index was created.