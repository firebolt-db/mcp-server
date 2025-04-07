# [](#create-user)CREATE USER

Creates a new user in Firebolt.

For more information, see [Managing users and roles](/Guides/managing-your-organization/managing-users.html).

## [](#syntax)Syntax

```
CREATE USER [ IF NOT EXISTS ] <user_name>  
[ WITH 
[ LOGIN = <login_name> | SERVICE_ACCOUNT = <service_account> ]
[ DEFAULT_DATABASE = <database_name> ]
[ DEFAULT_ENGINE = <engine_name> ]
[ ROLE = <role_name>[,...<role_name>] ]
]
```

## [](#parameters)Parameters

Parameter Description `<user_name>` The name of the user, may contain non-alpha-numeric characters such as exclamation points (!), percent signs (%), dots (.), underscores (\_), dashes (-), and asterisks (\*). Strings containing non-alphanumeric characters must be enclosed in single or double quotes. For more information about the full set of naming rules, see the [object identifiers guide](/Reference/object-identifiers.html#user-names). `<login>` (Optional) Specifies the name of the login to link the user with. This cannot be used in conjunction with the `SERVICE_ACCOUNT` parameter because a user can be linked to either a login OR a service account but not both. `<service_account>` (Optional) Specifies the name of the service account to link the user with. The `<service_account>` parameter cannot be used in conjunction with the `LOGIN_NAME` parameter because a user can be linked to a login OR a service account but not both. `<database_name>` (Optional) Defines the default database for the user. `<engine_name>` (Optional) Defines the default engine for the user. `<role_name>[, ...<role_name>]` (Optional) Defines a role for the user. Additional roles can be granted after the user is created. When assigning multiple roles, enclose the list of roles in parentheses. If no role is specified, the user is not granted any roles.

## [](#example)Example

The following code example creates a user named `alex`, links it to the login `alexs@acme.com`, and assigns it the roles of `analyst` and `data_engineer`:

```
CREATE USER alex WITH LOGIN= "alexs@acme.com" ROLE= (analyst, data_engineer);
```