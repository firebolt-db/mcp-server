# [](#information-schema-for-object_privileges)Information schema for object\_privileges

The `information_schema.object_privileges` view provides information about permissions granted to each role.

To be able to access this information, you must have [role privileges](/Overview/Security/Role-Based%20Access%20Control/role-permissions.html#role-permissions), ownership of the role, or ownership of the object to which the role is granted.

### [](#view-account-role-user-engine-and-database-permissions)View account, role, user, engine, and database permissions

To view account, role, user, engine and database permissions, make sure that current database is **not** selected. Then, query the `information_schema.object_privileges` view as shown in the following examples:

**View privileges directly under an account**

To view all privileges directly under an account, ensure that no database is selected, and query the `information_schema` as follows:

```
SELECT
  *
FROM
  information_schema.object_privileges;
```

You can also deselect the current database in the **Firebolt Develop Space** user interface (UI), by choosing `None` in [the current database selector](/assets/images/current_database_dropdown_none_option.png).

**View privileges in a specific database**

To view all privileges under a user defined database `db`, specify the database in the query as follows:

```
SELECT
  *
FROM
  db.information_schema.object_privileges;
```

### [](#view-object-permissions-in-the-current-database)View object permissions in the current database

When the current database is selected,`information_schema.object_privileges` only shows permissions for objects within that database. It does not show permissions for accounts, roles, users, engines, databases, and objects in other databases.

To view permissions for schemas, tables and views in the current database, set the current database with [USE DATABASE](/sql_reference/commands/data-definition/use-database.html), then select and view privileges in a query as follows:

```
USE DATABASE db;

SELECT
  *
FROM
  information_schema.object_privileges;
```

You can also use the [database selector](/assets/images/current_database_dropdown.png) in the UI.

## [](#columns-in-information_schemaobject_privileges)Columns in `information_schema.object_privileges`

Each row in `information_schema.object_privileges` contains the following columns:

Column Name Data Type Description grantor TEXT The name of the user that granted the privilege. grantee TEXT The name of the role that the privilege was granted to. object\_catalog TEXT The database containing the object on which the privilege is granted. object\_schema TEXT The schema containing the object on which the privilege is granted. object\_name TEXT The name of the object on which the privilege is granted. object\_type TEXT The type of the object on which the privilege is granted. privilege\_type TEXT The type of the privilege granted on the object. is\_grantable TEXT Specify `YES` if the privilege is grantable, and `NO` otherwise. created TIMESTAMPTZ The creation time of the privilege.