# [](#drop-role)DROP ROLE

Deletes a role. Note that role cannot be dropped if there are permissions granted to the role, in this case error message will be displayed and you need manually to drop permissions granted to the role and retry.

For more information, see [Role-based access control](/Guides/security/rbac.html).

A role cannot be dropped if there are permissions granted to the role. In this case, an error message will be displayed, and you need to manually drop the permissions granted to the role and retry.

## [](#syntax)Syntax

```
DROP ROLE [ IF EXISTS ] <role_name>
```

## [](#parameters)Parameters

Parameter Description `<role_name>` The name of the role.

## [](#example)Example

The following command will delete the role “user\_role”

```
DROP ROLE user_role;
```

### [](#example-2)Example 2

The following command will delete the role “my\_role\_2”

```
DROP ROLE IF EXISTS my_role_2
```

If “my\_role\_2” does not exist, no error message is thrown.