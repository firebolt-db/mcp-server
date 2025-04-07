# [](#information-schema-for-views)Information schema for views

You can use the `information_schema.views` view to return information about each view in a database. The view is available for each database and contains one row for each view in the database. You can use a `SELECT` query to return information about each view as shown in the example below.

To access information about views, you must have `USAGE` privileges on both the [schema](/Overview/Security/Role-Based%20Access%20Control/database-permissions/schema-permissions.html#schema-level-privileges) and the [database](/Overview/Security/Role-Based%20Access%20Control/database-permissions/schema-permissions.html#schema-level-privileges). You also need ownership of the view or [view-level privileges](/Overview/Security/Role-Based%20Access%20Control/database-permissions/view-permissions.html#view-permissions).

```
SELECT
  *
FROM
  information_schema.views;
```

## [](#columns-in-information_schemaviews)Columns in information\_schema.views

Each row has the following columns with information about each view.

Column Name Data Type Description table\_catalog TEXT The name of the catalog. Firebolt offers a single ‘default’ catalog. table\_schema TEXT The name of the database. table\_name TEXT The name of the view. view\_definition TEXT The query statement that defines the view. check\_option NULL Not applicable for Firebolt. is\_updatable TEXT Always `NO`. insertable\_into TEXT Always `NO`. is\_trigger\_updatable TEXT Always `NO`. is\_trigger\_deletable TEXT Always `NO`. is\_trigger\_insertable\_into TEXT Always `NO`. created TIMESTAMPTZ Time that the view was created. view\_owner TEXT The owner of the view. last\_altered TIMESTAMPTZ Time that the view was last changed. last\_altered\_by TEXT The user who last altered this view.