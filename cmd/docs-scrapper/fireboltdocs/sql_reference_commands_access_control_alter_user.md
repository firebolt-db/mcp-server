# [](#alter-user)ALTER USER

Updates the configuration of the specified user.

For more information, see [Managing users](/Guides/managing-your-organization/managing-users.html).

Users can modify most of their own account settings without requiring [RBAC](/Overview/Security/Role-Based%20Access%20Control/#role-based-access-control-rbac) permissions, except when altering [LOGIN](/Guides/managing-your-organization/managing-logins.html) configurations or a [SERVICE ACCOUNT](/Guides/managing-your-organization/service-accounts.html).

## [](#alter-user-set)ALTER USER SET

### [](#syntax)Syntax

```
ALTER USER <user_name> SET
	[ LOGIN = <login_name> | DEFAULT ]
        [ SERVICE_ACCOUNT = <service_account_name> | DEFAULT ]
        [ DEFAULT_DATABASE = <database_name> | DEFAULT ]
	[ DEFAULT_ENGINE = <engine_name> | DEFAULT ];
```

or

```
ALTER USER <user_name> RENAME TO <new_user_name>;
```

### [](#parameters)Parameters

Parameter Description `<user_name>` The name of the user, may contain non-alpha-numeric characters such as exclamation points (!), percent signs (%), dots (.), underscores (\_), dashes (-), and asterisks (\*). Strings containing non-alphanumeric characters must be enclosed in single or double quotes. For more information about the full set of naming rules, see the [object identifiers guide](/Reference/object-identifiers.html#user-names). `<new_user_name>` The new name of the user, used with the `RENAME TO` option. The new user name can be any string, and can also contain spaces and non-alpha-numeric characters such as exclamation points (!), percent signs (%), at signs(@), dot signs (.), underscore signs (\_), minus signs (-), and asterisks (\*). If the string contains spaces or non-alphanumeric characters, it must be enclosed in single or double quotes. `<login_name>` An optional, case-insensitive parameter to specify the name of the login to link the user with, used with the `SET` option. This cannot be used in conjunction with the `SERVICE_ACCOUNT` parameter - a user can be linked to a login OR a service account but not both. `DEFAULT` disassociates the user from its login. The user will become unusable before it’s re-associated with some login or service account. `<service_account_name>` An optional parameter to specify the name of the service account to link the user with, used with the `SET` option. This cannot be used in conjunction with the `LOGIN` parameter - a user can be linked to a login OR a service account but not both. `DEFAULT` disassociates the user from its service account. The user will become unusable before it’s re-associated with some login or service account. `<database_name>` An optional parameter to define a default database for the user (for future purposes). Used with the `SET` option. `<engine_name>` An optional parameter to define a default engine for the user (for future purposes).Used with the `SET` option.

### [](#example)Example

The following command will rename the “alex” account to “alexs”.

```
ALTER USER "alex" RENAME TO "alexs";
```

This command will link the user “alex” to the “alexs@acme.com” login.

```
ALTER USER alex SET LOGIN="alexs@acme.com";
```

## [](#alter-user-owner-to)ALTER USER OWNER TO

Change the owner of a user. The current owner of a user can be viewed in the `information_schema.users` view on `user_owner` column.

Check [ownership](/Guides/security/ownership.html) page for more info.

### [](#syntax-1)Syntax

```
ALTER USER <user_name> OWNER TO <user>
```

### [](#parameters-1)Parameters

Parameter Description `<user_name>` The name of the user to change the owner of. `<user>` The new owner of the user.