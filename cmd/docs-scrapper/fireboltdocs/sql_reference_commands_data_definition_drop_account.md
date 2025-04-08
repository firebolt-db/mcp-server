# [](#drop-account)DROP ACCOUNT

Deletes an account.

For more information, see [Managing accounts](/Guides/managing-your-organization/managing-accounts.html).

## [](#syntax)Syntax

```
DROP ACCOUNT [ IF EXISTS ] <account_name> [ RESTRICT | CASCADE ];
```

## [](#parameters)Parameters

Parameter Description `<account_name>` The name of the account to delete. RESTRICT or CASCADE An optional parameter to specify deletion mode.  
RESTRICT mode prevents dropping the account if there is any sub-object contained in the account.  
By default, if the account contains no objects, it will just be dropped.  
CASCADE mode automatically drops the account and all the sub-objects it contains (databases, engines, users, roles, etc.).

All engines in your accounts must be in a stopped state before running the `DROP ACCOUNT … CASCADE` statement.

## [](#example)Example

The following command will delete the “dev” account.

```
DROP ACCOUNT dev
```