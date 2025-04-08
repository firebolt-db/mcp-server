# [](#drop-user)DROP USER

Deletes a user.

For more information, see [Managing users](/Guides/managing-your-organization/managing-users.html).

A user cannot be dropped if it owns objects. In this case, an error message will be displayed, and you need to manually drop the objects, or transfer ownership.

for more information, see [Ownership](/Guides/security/ownership.html).

## [](#syntax)Syntax

```
DROP USER [ IF EXISTS ] <user_name> ;
```

## [](#parameters)Parameters

Parameter Description `<user_name>` The name of the user to delete. If the user name contains spaces or non-alphanumeric characters, it must be enclosed in single or double quotes.

## [](#example)Example

The following command will delete the “alex” user.

```
DROP USER alex;
```