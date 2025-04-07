# [](#drop-login)DROP LOGIN

Deletes an account.

For more information, see [Managing logins](/Guides/managing-your-organization/managing-logins.html).

## [](#syntax)Syntax

```
DROP LOGIN [ IF EXISTS ] <login_name>;
```

If the login is linked to a user, it can not be dropped. In order to drop a login linked to a user, the link must be reset `alter user foo set login="new-login@acme.com"|DEFAULT` or the user dropped.

## [](#parameters)Parameters

Parameter Description `<login_name>` The name of the login to delete.

## [](#example)Example

The following command will delete the “alexs@acme.com” login.

```
DROP LOGIN "alexs@acme.com";
```