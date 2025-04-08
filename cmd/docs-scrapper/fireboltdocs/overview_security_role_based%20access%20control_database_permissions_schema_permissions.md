# [](#schema-permissions)Schema permissions

In Firebolt, a **schema** is a logical namespace within a database that organizes **tables**, **views**, and other objects. Schema-level permissions allow roles to perform specific actions, such as accessing, modifying, or managing objects within a schema.

To perform actions on a schema or its objects, the role must also have the **USAGE** privilege on the schema’s parent database.

## [](#schema-level-privileges)Schema-level privileges

Privilege Description GRANT Syntax REVOKE Syntax USAGE Allows access to the schema and its objects `GRANT USAGE ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE USAGE ON SCHEMA public IN <database_name> FROM <role_name>;` MODIFY Allows altering the schema properties, including renaming or dropping the schema them. `GRANT MODIFY ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE MODIFY ON SCHEMA public IN <database_name> FROM <role_name>;` CREATE Allows creating new objects, such as tables and views, within the schema. `GRANT CREATE ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE CREATE ON SCHEMA public IN <database_name> FROM <role_name>;` [DELETE](/sql_reference/commands/data-management/delete.html) ANY Allows deleting rows and partitions from all current and future tables. `GRANT DELETE ANY ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE DELETE ANY ON SCHEMA public IN <database_name> FROM <role_name>;` [INSERT](/sql_reference/commands/data-management/insert.html) ANY Allows inserting rows into all current and future tables within the schema. `GRANT INSERT ANY ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE INSERT ANY ON SCHEMA public IN <database_name> FROM <role_name>;` [UPDATE](/sql_reference/commands/data-management/update.html) ANY Allows updating rows in all current and future tables within the schema. `GRANT UPDATE ANY ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE UPDATE ANY ON SCHEMA public IN <database_name> FROM <role_name>;` [TRUNCATE](/sql_reference/commands/data-management/truncate-table.html) ANY Allows truncating all current and future tables within the schema. `GRANT TRUNCATE ANY ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE TRUNCATE ANY ON SCHEMA public IN <database_name> FROM <role_name>;` [VACUUM](/sql_reference/commands/data-management/vacuum.html) ANY Allows running the `VACUUM` operation on all current and future tables. `GRANT VACUUM ANY ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE VACUUM ANY ON SCHEMA public IN <database_name> FROM <role_name>;` MODIFY ANY Allows modifying or dropping all current and future objects in the schema. `GRANT MODIFY ANY ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE MODIFY ANY ON SCHEMA public IN <database_name> FROM <role_name>;` SELECT ANY Allows reading data from all current and future objects within the schema. `GRANT SELECT ANY ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE SELECT ANY ON SCHEMA public IN <database_name> FROM <role_name>;` ALL \[PRIVILEGES] Grants all direct privileges over the schema to a role. `GRANT ALL ON SCHEMA public IN <database_name> TO <role_name>;` `REVOKE ALL ON SCHEMA public IN <database_name> FROM <role_name>;`

## [](#examples-of-granting-schema-permissions)Examples of granting schema permissions

### [](#usage-permission)USAGE permission

The following code example [grants](/sql_reference/commands/access-control/grant.html) the role `developer_role` permission to use the specified schema.

```
GRANT USAGE ON SCHEMA "public" TO developer_role;
```

### [](#modify-permission)MODIFY permission

The following code example gives the role `developer_role` permission to alter properties or drop the specified schema.

```
GRANT MODIFY ON SCHEMA "public" TO developer_role;
```

### [](#create-permission)CREATE permission

The following code example grants the role `developer_role` the ability to create new objects in the specified schema:

```
GRANT CREATE ON SCHEMA "public" TO developer_role;
```

### [](#delete-any-permission)DELETE ANY permission

The following code example gives the role `developer_role` permission to [delete](/sql_reference/commands/data-management/delete.html) rows and partitions from all current and future tables in the specified schema:

```
GRANT DELETE ANY ON SCHEMA "public" TO developer_role;
```

### [](#insert-any-permission)INSERT ANY permission

The following code example grants the role `developer_role` permission to [insert](/sql_reference/commands/data-management/insert.html) rows into all current and future tables in the specified schema:

```
GRANT INSERT ANY ON SCHEMA "public" TO developer_role;
```

### [](#update-any-permission)UPDATE ANY permission

The following code example gives the role `developer_role` permission to [update](/sql_reference/commands/data-management/update.html) rows in all current and future tables in the specified schema:

```
GRANT UPDATE ANY ON SCHEMA "public" TO developer_role;
```

### [](#truncate-any-permission)TRUNCATE ANY permission

The following code example grants the role `developer_role` the ability to [truncate](/sql_reference/commands/data-management/truncate-table.html) all current and future tables in the specified schema:

```
GRANT TRUNCATE ANY ON SCHEMA "public" TO developer_role;
```

### [](#vacuum-any-permission)VACUUM ANY permission

The following code example gives the role `developer_role` permission to run [`VACUUM`](/sql_reference/commands/data-management/vacuum.html) operations on all current and future tables in the specified schema:

```
GRANT VACUUM ANY ON SCHEMA "public" TO developer_role;
```

### [](#modify-any-permission)MODIFY ANY permission

The following code example grants the role `developer_role` permission to modify or drop all current and future objects in the specified schema:

```
GRANT MODIFY ANY ON SCHEMA "public" TO developer_role;
```

### [](#select-any-permission)SELECT ANY permission

The following code example gives the role `developer_role` permission to select data from all current and future objects in the specified schema:

```
GRANT SELECT ANY ON SCHEMA "public" TO developer_role;
```

### [](#all-permissions)ALL permissions

The following code example gives the role `developer_role` all the direct permissions over schema `public`:

```
GRANT ALL ON SCHEMA "public" TO developer_role;
```