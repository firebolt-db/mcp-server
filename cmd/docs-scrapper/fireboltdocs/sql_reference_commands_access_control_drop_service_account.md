# [](#drop-service-account)DROP SERVICE ACCOUNT

Deletes a service account.

For more information, see [Service accounts](/Guides/managing-your-organization/service-accounts.html).

## [](#syntax)Syntax

```
DROP SERVICE ACCOUNT <service_account_name>;
```

If the service account is linked to a user, it can not be dropped. In order to drop a service account linked to a user, the link must be reset `alter user foo set service_account=new_service_account|DEFAULT` or the user dropped.

## [](#parameters)Parameters

Parameter Description `<service_account_name>` The name of the service account to delete.

## [](#example)Example

The following command will delete the “sa1” service account.

```
DROP SERVICE ACCOUNT "sa1";
```