# [](#user-permissions)User permissions

In Firebolt, a [user](/Overview/organizations-accounts.html#users) is associated with a [login](/Guides/managing-your-organization/managing-logins.html) or [service account](/Guides/managing-your-organization/service-accounts.html), which grants them access to that account. You can assign a [role](/Overview/organizations-accounts.html#roles) to a user, and the role determines the specific actions they are authorized to perform within the account.

The following table outlines the privileges that can be granted for users within a particular account:

Privilege Description GRANT Syntax REVOKE Syntax MODIFY Grants the ability to drop the specified user. `GRANT MODIFY ON USER <user_name> TO <role>;` `REVOKE MODIFY ON USER <user_name> FROM <role>;`

Users can modify most of their own account settings without requiring [RBAC](/Overview/Security/Role-Based%20Access%20Control/#role-based-access-control-rbac) permissions, except when altering [LOGIN](/Guides/managing-your-organization/managing-logins.html) configurations or a [SERVICE ACCOUNT](/Guides/managing-your-organization/service-accounts.html).

## [](#examples-of-granting-user-permissions)Examples of granting user permissions

### [](#modify-permission)MODIFY permission

The following code example grants the role `developer_role` permission to drop the `my_user` user:

```
GRANT MODIFY ON USER my_user TO developer_role;
```