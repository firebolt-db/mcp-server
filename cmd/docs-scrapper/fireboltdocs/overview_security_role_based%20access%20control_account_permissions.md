# [](#account-permissions)Account permissions

Accounts represent the physical instance of your data warehouse in Firebolt and are created in Firebolt-supported regions. All objects within an account—such as databases, engines, roles, and users—are scoped to that specific account.

Privilege Description GRANT Syntax REVOKE Syntax CREATE DATABASE Allows creating new databases in an account. `GRANT CREATE DATABASE ON ACCOUNT <account_name> TO <role>;` `REVOKE CREATE DATABASE ON ACCOUNT <account_name> FROM <role>;` USAGE ANY DATABASE Allows using all current and future databases in an account. `GRANT USAGE ANY DATABASE ON ACCOUNT <account_name> TO <role>;` `REVOKE USAGE ANY DATABASE ON ACCOUNT <account_name> FROM <role>;` MODIFY ANY DATABASE Allows editing all current and future databases in an account. `GRANT MODIFY ANY DATABASE ON ACCOUNT <account_name> TO <role>;` `REVOKE MODIFY ANY DATABASE ON ACCOUNT <account_name> FROM <role>;` [CREATE ENGINE](/sql_reference/commands/engines/create-engine.html) Allows creating new engines in an account. `GRANT CREATE ENGINE ON ACCOUNT <account_name> TO <role>;` `REVOKE CREATE ENGINE ON ACCOUNT <account_name> FROM <role>;` USAGE ANY ENGINE Allows using all current and future engines in an account. `GRANT USAGE ANY ENGINE ON ACCOUNT <account_name> TO <role>;` `REVOKE USAGE ANY ENGINE ON ACCOUNT <account_name> FROM <role>;` OPERATE ANY ENGINE Allows starting and stopping all current and future engines in the account. `GRANT OPERATE ANY ENGINE ON ACCOUNT <account_name> TO <role>;` `REVOKE OPERATE ANY ENGINE ON ACCOUNT <account_name> FROM <role>;` MODIFY ANY ENGINE Allows editing all current and future engines in the account. `GRANT MODIFY ANY ENGINE ON ACCOUNT <account_name> TO <role>;` `REVOKE MODIFY ANY ENGINE ON ACCOUNT <account_name> FROM <role>;` [CREATE ROLE](/sql_reference/commands/access-control/create-role.html) Allows creating new roles in the account. `GRANT CREATE ROLE ON ACCOUNT <account_name> TO <role>;` `REVOKE CREATE ROLE ON ACCOUNT <account_name> FROM <role>;` MODIFY ANY ROLE Allows editing all current and future roles in the account. `GRANT MODIFY ANY ROLE ON ACCOUNT <account_name> TO <role>;` `REVOKE MODIFY ANY ROLE ON ACCOUNT <account_name> FROM <role>;` [CREATE USER](/sql_reference/commands/access-control/create-user.html) Allows creating new users in the account. `GRANT CREATE USER ON ACCOUNT <account_name> TO <role>;` `REVOKE CREATE USER ON ACCOUNT <account_name> FROM <role>;` MODIFY ANY USER Allows editing all current and future users in the account. `GRANT MODIFY ANY USER ON ACCOUNT <account_name> TO <role>;` `REVOKE MODIFY ANY USER ON ACCOUNT <account_name> FROM <role>;` MONITOR \[ANY USAGE] Enables the tracking of engine queries through the `engine_running_queries` view for active queries and the `engine_query_history` view for past queries in `information_schema`. `GRANT MONITOR ANY USAGE ON ACCOUNT <account_name> TO <role>;` `REVOKE MONITOR ANY USAGE ON ACCOUNT <account_name> FROM <role>;` ALL \[PRIVILEGES] Grants all direct privileges for a specified account to a specified role. `GRANT ALL ON ACCOUNT <account_name> TO <role>;` `REVOKE ALL ON ACCOUNT <account_name> FROM <role>;`

Revoking a privilege removes it from a role but does not explicitly deny the privilege. If the privilege was not previously granted, revoking it has no effect.

## [](#examples-of-granting-account-level-permissions)Examples of granting account-level permissions

### [](#create-database-permission)CREATE DATABASE permission

The following code example [grants](/sql_reference/commands/access-control/grant.html) the role `developer_role` permission to create new databases within the `account_name`:

```
GRANT CREATE DATABASE ON ACCOUNT account_name TO developer_role;
```

### [](#usage-any-database-permission)USAGE ANY DATABASE permission

The following code example gives permission to the role `developer_role` to access all current and future databases within the `account_name`:

```
GRANT USAGE ANY DATABASE ON ACCOUNT account_name TO developer_role;
```

### [](#modify-any-database-permission)MODIFY ANY DATABASE permission

The following code example grants the role `developer_role` permission to modify or delete all current and future databases within the `account_name`:

```
GRANT MODIFY ANY DATABASE ON ACCOUNT account_name TO developer_role;
```

### [](#create-engine-permission)[CREATE ENGINE](/sql_reference/commands/engines/create-engine.html) permission

The following code example gives the role `developer_role` permission to create new engines within the `account_name`:

```
GRANT CREATE ENGINE ON ACCOUNT account_name TO developer_role;
```

### [](#usage-any-engine-permission)USAGE ANY ENGINE permission

The following code example grants the role `developer_role` permission to use all current and future engines within the `account_name`:

```
GRANT USAGE ANY ENGINE ON ACCOUNT account_name TO developer_role;
```

### [](#operate-any-engine-permission)OPERATE ANY ENGINE permission

The following code example gives the role `developer_role` permission to start and stop all current and future engines within the `account_name`:

```
GRANT USAGE ANY DATABASE ON ACCOUNT account_name TO developer_role;
```

### [](#modify-any-engine-permission)MODIFY ANY ENGINE permission

The following code example grants the role `developer_role` permission to modify or delete all current and future engines within the `account_name`:

```
GRANT MODIFY ANY ENGINE ON ACCOUNT account_name TO developer_role;
```

### [](#create-role-permission)[CREATE ROLE](/sql_reference/commands/access-control/create-role.html) permission

The following code example gives the role `developer_role` permission to create new roles within the `account_name`:

```
GRANT CREATE ROLE ON ACCOUNT account_name TO developer_role;
```

### [](#modify-any-role-permission)MODIFY ANY ROLE permission

The following code example grants the role `developer_role` permission to modify or delete all current and future roles within the `account_name`:

```
GRANT MODIFY ANY ROLE ON ACCOUNT account_name TO developer_role;
```

### [](#create-user-permission)[CREATE USER](/sql_reference/commands/access-control/create-user.html) permission

The following code example gives the role `developer_role` permission to create new users within the `account_name`:

```
GRANT CREATE USER ON ACCOUNT account_name TO developer_role;
```

### [](#modify-any-user-permission)MODIFY ANY USER permission

The following code example grants the role `developer_role` permission to modify or delete all current and future users within the `account_name`:

```
GRANT MODIFY ANY USER ON ACCOUNT account_name TO developer_role;
```

### [](#monitor-any-usage-permission)MONITOR \[ANY USAGE] permission

The following code example grants the role `developer_role` permission to see the query history and currently running queries on all the engines within `account_name`:

```
GRANT MONITOR ANY USAGE ON ACCOUNT "account-1" TO developer_role;
```