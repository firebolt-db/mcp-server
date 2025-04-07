# [](#create-role)CREATE ROLE

Creates a new role.

For more information, see [Role-based access control](/Guides/security/rbac.html).

## [](#syntax)Syntax

```
CREATE ROLE [ IF NOT EXISTS ] <role_name>
```

## [](#parameters)Parameters

Parameter Description `<role_name>` The name of the role.

## [](#example)Example

The following command will create a role “user\_role”

```
CREATE ROLE user_role;
```

## [](#example-2)Example 2

The following command will create a role “user\_role\_2”

```
CREATE ROLE IF NOT EXISTS my_role_2
```

If “my\_role\_2” exists, no error message is thrown.