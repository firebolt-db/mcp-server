# [](#custom-roles)Custom roles

In Firebolt, custom roles can be created at the account level to manage access control and permissions. Only users with the [account\_admin](/Overview/organizations-accounts.html#account-administrative-role) system role or those granted the [CREATE ROLE](/sql_reference/commands/access-control/create-role.html) privilege can create custom roles. Once created, custom roles can be assigned to any user or existing role by the `account_admin` or the resource owner.

Privileges can be granted to custom roles by `account_admin` or the [resource owner](/Overview/Security/Role-Based%20Access%20Control/ownership.html). For example, the owner of a table can grant `SELECT` privileges on that table to a custom role.

For more information about creating roles, see [creating a custom role via SQL or the Firebolt UI](/Guides/managing-your-organization/managing-users.html#create-a-role).