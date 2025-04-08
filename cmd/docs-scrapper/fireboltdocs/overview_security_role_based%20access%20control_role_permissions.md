# [](#role-permissions)Role permissions

In Firebolt, a [role](/Overview/organizations-accounts.html#roles) is a group of permissions that can have privileges assigned to it. You can [grant](/sql_reference/commands/access-control/grant.html#grant-role) a role to another role or to users.

The following table outlines the privileges that can be granted for roles within a particular account:

Privilege Description GRANT Syntax REVOKE Syntax MODIFY Grants the ability to drop the specified role. `GRANT MODIFY ON ROLE <role_name> TO <role>;` `REVOKE MODIFY ON ROLE <role_name> FROM <role>;`

## [](#examples-of-granting-role-permissions)Examples of granting role permissions

### [](#modify-permission)MODIFY permission

The following code example grants the role `developer_role` permission to drop the `my_role` role:

```
GRANT MODIFY ON ROLE my_role TO developer_role;
```