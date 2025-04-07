# [](#table-permissions)Table permissions

In Firebolt, a **table** is a structured data object within a database, composed of rows and columns. Tables are the foundational units for organizing, querying, and managing data in your Firebolt data warehouse. Table-level permissions allow roles to perform actions such as selecting, modifying, or managing data within specific tables.

To perform actions on a table, roles must also have **USAGE** permissions on both the parent schema and the parent database of the table.

## [](#table-level-privileges)Table-level privileges

Privilege Description GRANT Syntax REVOKE Syntax SELECT Allows selecting rows from the table. `GRANT SELECT ON TABLE <table_name> TO <role_name>;` `REVOKE SELECT ON TABLE <table_name> FROM <role_name>;` [INSERT](/sql_reference/commands/data-management/insert.html) Allows inserting rows into the table. Applies to managed tables only. `GRANT INSERT ON TABLE <table_name> TO <role_name>;` `REVOKE INSERT ON TABLE <table_name> FROM <role_name>;` MODIFY Allows modifying and dropping the table. `GRANT MODIFY ON TABLE <table_name> TO <role_name>;` `REVOKE MODIFY ON TABLE <table_name> FROM <role_name>;` [DELETE](/sql_reference/commands/data-management/delete.html) Allows deleting rows and dropping partitions from the table. Applies to managed tables only. `GRANT DELETE ON TABLE “<table_name>” TO <role_name>;` `REVOKE DELETE ON TABLE “<table_name>” FROM <role_name>;` [UPDATE](/sql_reference/commands/data-management/update.html) Allows updating rows in the table. Applies to managed tables only. `GRANT UPDATE ON TABLE <table_name> TO <role_name>;` `REVOKE UPDATE ON TABLE <table_name> FROM <role_name>;` [TRUNCATE](/sql_reference/commands/data-management/truncate-table.html) Allows truncating a table. Applies to managed tables only. `GRANT TRUNCATE ON TABLE <table_name> TO <role_name>;` `REVOKE TRUNCATE ON TABLE <table_name> FROM <role_name>;` [VACUUM](/sql_reference/commands/data-management/vacuum.html) Allows running the `VACUUM` operation. Applies to managed tables only. `GRANT VACUUM ON TABLE <table_name> TO <role_name>;` `REVOKE VACUUM ON TABLE <table_name> FROM <role_name>;` ALL \[PRIVILEGES] Grants all privileges over the table to a role. `GRANT ALL ON TABLE <table_name> TO <role_name>;` `REVOKE ALL ON TABLE <table_name> FROM <role_name>;`

To grant permissions across all tables in a schema, use [schema-level privileges](/Overview/Security/Role-Based%20Access%20Control/database-permissions/schema-permissions.html). For example, privileges like **SELECT ANY**, **INSERT ANY**, or **DELETE ANY** at the schema level will apply to all current and future tables within that schema.

## [](#aggregating-indexes)Aggregating Indexes

An [aggregating index](/Overview/indexes/aggregating-index.html) in Firebolt accelerates queries involving aggregate functions on large tables. This reduces compute usage and improves query performance.

To **create** or **drop** an aggregating index, a role must have the following permissions:

- `MODIFY` permission on the table.
- `CREATE` permission on the parent schema.
- `USAGE` permission on the parent schema.
- `USAGE` permission on the parent database.

To drop an aggregating index, the role requires:

- `MODIFY` permission on the table.
- `USAGE` permission on the parent schema.
- `USAGE` permission on the parent database.

## [](#examples-of-modifying-table-permissions)Examples of modifying table permissions

The following example use [`GRANT`](/sql_reference/commands/access-control/grant.html) to grant permissions. You can also replace `GRANT` with [REVOKE](/sql_reference/commands/access-control/revoke.html) in any of the examples to remove any granted privileges.

### [](#select-permission)SELECT permission

The following code example [grants](/sql_reference/commands/access-control/grant.html) the role `developer_role` permission to read data from the `games` table:

```
GRANT SELECT ON TABLE games TO developer_role;
```

### [](#insert-permission)INSERT permission

The following code example gives the role `developer_role` permissions to [insert](/sql_reference/commands/data-management/insert.html) rows into the `games` table:

```
GRANT INSERT ON TABLE games TO developer_role;
```

### [](#modify-permission)MODIFY permission

The following code example grants the role `developer_role` permission to alter or drop the `games` table:

```
GRANT MODIFY ON TABLE games TO developer_role;
```

### [](#delete-permission)DELETE permission

The following code example gives the role `developer_role` permission to [delete](/sql_reference/commands/data-management/delete.html) rows or partitions from the `games` table:

```
GRANT DELETE ON TABLE games TO developer_role;
```

### [](#update-permission)UPDATE permission

The following code example grants the role `developer_role` permission to [update](/sql_reference/commands/data-management/update.html) rows in the `games` table:

```
GRANT UPDATE ON TABLE games TO developer_role;
```

### [](#truncate-permission)TRUNCATE permission

The following code example gives the role `developer_role` permission to [truncate](/sql_reference/commands/data-management/truncate-table.html) the `games` table, removing all rows:

```
GRANT TRUNCATE ON TABLE games TO developer_role;
```

### [](#vacuum-permission)VACUUM permission

The following code example grants the role `developer_role` permission to run the [`VACUUM`](/sql_reference/commands/data-management/vacuum.html) operation on the `games` table:

```
GRANT VACUUM ON TABLE games TO developer_role;
```

### [](#all-permissions)ALL permissions

The following code example grants the role `developer_role` with all permissions on the table `games`:

```
GRANT ALL ON TABLE games TO developer_role;
```

## [](#considerations)Considerations

- Use the [REVOKE](/sql_reference/commands/access-control/revoke.html) statement to remove any granted privileges. Replace [`GRANT`](/sql_reference/commands/access-control/grant.html) with [`REVOKE`](/sql_reference/commands/access-control/revoke.html) in the examples above.
- Table-level permissions apply only to the specified table. For broader control, consider granting schema-level privileges.