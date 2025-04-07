# [](#information-schema-for-schemata)Information schema for schemata

You can use the `information_schema.schemata` view to return information about schemas available in the database. Run a `SELECT` query to return information about each schema as shown in the example below.

To view schema information, you must have [schema privilege](/Overview/Security/Role-Based%20Access%20Control/database-permissions/schema-permissions.html#schema-level-privileges) or ownership of the schema object.

```
SELECT
  *
FROM
  information_schema.schemata;
```

## [](#columns-in-information_schemaschemata)Columns in information\_schema.schemata

Each row has the following columns with information about the schema.

Column Name Data Type Description   catalog\_name TEXT Name of the catalog.   schema\_name TEXT Name of the schema.   schema\_owner TEXT Owner of the schema.   default\_character\_set\_catalog TEXT The catalog that contains the character set. Defaults to NULL.   default\_character\_set\_schema TEXT The schema that contains the character set. Defaults to NULL.   default\_character\_set\_name TEXT Default character set of the schema. Defaults to `UTF-8`.   sql\_path TEXT SQL path of the schema.   description TEXT Description of the schema.