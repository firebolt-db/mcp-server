# [](#alter-account)ALTER ACCOUNT

Updates the configuration of the specified database `<database_name>`.

For more information, see [Managing accounts](/Guides/managing-your-organization/managing-accounts.html).

## [](#syntax)Syntax

```
ALTER ACCOUNT <account_name> RENAME TO <new_account_name>;
```

## [](#parameters)Parameters

Parameter Description `<account_name>` The name of the account to be altered. `<new_account_name>` The new name for the account. The account name must start and end with an alphabetic character and cannot contain spaces or special characters except for hyphens (-).

## [](#example)Example

The following command will rename the “dev” account to “staging”.

```
ALTER ACCOUNT dev RENAME TO staging;
```