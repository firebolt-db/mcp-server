# [](#information-schema-for-catalogs)Information schema for catalogs

You can use the `information_schema.catalogs` view to get information about catalogs (databases in SQL terminology). You can use a `SELECT` query to return information about each database as shown in the example below.

In order to view catalog information, you must have [catalog](/Overview/Security/Role-Based%20Access%20Control/database-permissions/#database-level-privileges) privileges or have ownership of the catalog object.

```
SELECT
  *
FROM
  information_schema.catalogs;
```

## [](#columns-in-information_schemacatalogs)Columns in information\_schema.catalogs

Each row has the following columns with information about the database.

Column Name Data Type Description catalog\_name TEXT Name of the database. default\_collation TEXT Always ‘POSIX’. default\_character\_set TEXT Always ‘UTF-8’. description TEXT The description of the database. created TIMESTAMPTZ The time the database was created. ddl TEXT The text of the SQL statement that created the database. catalog\_owner TEXT The owner of the database, `NULL` if there is none. last\_altered TIMESTAMPTZ Time the database was last altered. last\_altered\_by TEXT Name of the last user to alter the database.