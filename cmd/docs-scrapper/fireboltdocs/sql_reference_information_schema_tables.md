# [](#information-schema-for-tables)Information schema for tables

You can use the `information_schema.tables` view to return information about each table in a database. The view is available for each database and contains one row for each table in the database. You can use a `SELECT` query to return information about each table as shown in the example below.

To view table information, you must have `USAGE` privileges on both the [schema](/Overview/Security/Role-Based%20Access%20Control/database-permissions/schema-permissions.html#schema-level-privileges) and the [database](/Overview/Security/Role-Based%20Access%20Control/database-permissions/#database-level-privileges). You also need ownership of the table or the necessary [table-level privileges](/Overview/Security/Role-Based%20Access%20Control/database-permissions/table-permissions.html#table-level-privileges) required for the intended action.

```
SELECT
  *
FROM
  information_schema.tables;
```

## [](#columns-in-information_schematables)Columns in information\_schema.tables

Each row has the following columns with information about each table.

Column Name Data Type Description table\_catalog TEXT The name of the database. table\_schema TEXT The name of the schema. table\_name TEXT The name of the table. table\_type TEXT The table’s type, such as `BASE TABLE`, `EXTERNAL` `VIEW`. table\_owner TEXT The owner of the table, or `NULL` if there is no owner. created TIMESTAMPTZ The time that the table or view was created. last\_altered TIMESTAMPTZ Not applicable for Firebolt. last\_altered\_by TEXT Not applicable for Firebolt. primary\_index TEXT An ordered array of the column names that comprise the primary index definition, if applicable. number\_of\_rows BIGINT The number of rows in the table. compressed\_bytes BIGINT The compressed size of the table in bytes. uncompressed\_bytes BIGINT The uncompressed size of the table in bytes. compression\_ratio NUMERIC The compression ratio (`<uncompressed_bytes>`/`<compressed_bytes>`). number\_of\_tablets INTEGER The number of tablets that comprise the table. fragmentation DECIMAL The table fragmentation percentage (between 0-100). type TEXT The table’s type. location\_name TEXT Not applicable for Firebolt. ddl TEXT The text of the SQL statement that created the table. self\_referencing\_column\_name NULL Not applicable for Firebolt. reference\_generation NULL Not applicable for Firebolt. user\_defined\_type\_catalog NULL Not applicable for Firebolt. user\_defined\_type\_schema NULL Not applicable for Firebolt. user\_defined\_type\_name NULL Not applicable for Firebolt. is\_insertable\_into TEXT `YES` if the table is insertable, `NO` otherwise. is\_typed TEXT Always `NO`. commit\_action NULL Not applicable for Firebolt.