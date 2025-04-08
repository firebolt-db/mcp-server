# [](#database-permissions)Database permissions

In Firebolt, a **database** is a logical container that organizes your data warehouse by holding components such as **tables**, **views**, **indexes**, and other database objects, as shown in the following diagram:

![Firebolt's object model contains schema under databases, and tables, views, and indexes under schema.](../../../../assets/images/database-hierarchy.png)

Database-level permissions define what actions roles can perform within a database and its associated objects.

## [](#database-level-privileges)Database-level privileges

Privilege Description GRANT Syntax REVOKE Syntax USAGE Allows access to the database and enables attaching engines to it. `GRANT USAGE ON DATABASE <database_name> TO <role>;` `REVOKE USAGE ON DATABASE <database_name> FROM <role>;` MODIFY Allows altering database properties and dropping the database. `GRANT MODIFY ON DATABASE <database_name> TO <role>;` `REVOKE MODIFY ON DATABASE <database_name> FROM <role>;` USAGE ANY SCHEMA Allows access to all current and future schemas within the database. `GRANT USAGE ANY SCHEMA ON DATABASE <database_name> TO <role>;` `REVOKE USAGE ANY SCHEMA ON DATABASE <database_name> FROM <role>;` [VACUUM](/sql_reference/commands/data-management/vacuum.html) ANY Allows running the `VACUUM` operation on all current and future tables. `GRANT VACUUM ANY ON DATABASE <database_name> TO <role>;` `REVOKE VACUUM ANY ON DATABASE <database_name> FROM <role>;` ALL \[PRIVILEGES] Grants all direct privileges over the database to a role. `GRANT ALL ON DATABASE <database_name> TO <role>;` `REVOKE ALL ON DATABASE <database_name> FROM <role>;`

## [](#examples-of-granting-database-permissions)Examples of granting database permissions

### [](#usage-permission)USAGE permission

The following code example [grants](/sql_reference/commands/access-control/grant.html) the role `developer_role` access to use the specified database:

```
GRANT USAGE ON DATABASE "database-1" TO developer_role;
```

### [](#modify-permission)MODIFY permission

The following code example gives the role `developer_role` permission to alter properties or drop the specified database:

```
GRANT MODIFY ON DATABASE "database-1" TO developer_role;
```

### [](#usage-any-schema-permission)USAGE ANY SCHEMA permission

The following code example grants the role `developer_role` access to all current and future schemas within the specified database:

```
GRANT USAGE ANY SCHEMA ON DATABASE "database-1" TO developer_role;
```

### [](#vacuum-any-permission)VACUUM ANY permission

The following code example gives the role `developer_role` permission to run [VACUUM](/sql_reference/commands/data-management/vacuum.html) operations on all current and future tables in the specified database:

```
GRANT VACUUM ANY ON DATABASE "database-1" TO developer_role;
```

### [](#all-permissions)ALL permissions

The following code example gives the role `developer_role` all the direct permissions over database `database-1`:

```
GRANT ALL ON DATABASE "database-1" TO developer_role;
```

* * *

- [Schema permissions](/Overview/Security/Role-Based%20Access%20Control/database-permissions/schema-permissions.html)
- [Table permissions](/Overview/Security/Role-Based%20Access%20Control/database-permissions/table-permissions.html)
- [View permissions](/Overview/Security/Role-Based%20Access%20Control/database-permissions/view-permissions.html)