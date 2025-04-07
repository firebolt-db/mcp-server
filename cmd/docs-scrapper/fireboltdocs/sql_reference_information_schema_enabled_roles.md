# [](#information-schema-for-enabled_roles)Information schema for enabled\_roles

The `information_schema.enabled_roles` view lists roles in a Firebolt account that is either owned by a user or comes with [role privileges](/Overview/Security/Role-Based%20Access%20Control/role-permissions.html#role-permissions) to access.

The following code example uses a `SELECT` query to return information about each role:

```
SELECT
  *
FROM
  information_schema.enabled_roles;
```

For more information about permissions to access and perform operations on specific objects by role, see [Manage role-based access control](/Guides/security/rbac.html).

## [](#columns-in-information_schemaenabled_roles)Columns in information\_schema.enabled\_roles

Each row contains the following columns with information about the role:

Column Name Data Type Description role\_name TEXT The name of the role. created TIMESTAMPTZ The time that the role was created. role\_owner TEXT The name of the user who created the role.