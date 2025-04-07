# [](#information-schema-for-service_accounts)Information schema for service\_accounts

Use the `information_schema.service_accounts` view to return information about service accounts. The following code example uses a `SELECT` query to return information about each service account:

```
SELECT
  *
FROM
  information_schema.service_accounts;
```

For more information about service accounts, see [Manage programmatic access to Firebolt](/Guides/managing-your-organization/service-accounts.html).

## [](#columns-in-information_schemaservice_accounts)Columns in information\_schema.service\_accounts

Each row has the following columns with information about a service account:

Column Name Data Type Description service\_account\_id TEXT The ID of the service account. service\_account\_name TEXT The name of the service account. network\_policy\_name TEXT The name of the network policy used by this service account. service\_account\_description TEXT The description of the service account. is\_organization\_admin BOOLEAN Specifies if the user associated with the service account has an [organizational administrative role](/Overview/organizations-accounts.html#organizational-administrative-role). connection\_preference TEXT Defines the connectivity preference for a service account. The default value is `PREFER_PUBLIC` if not specified. A user with an [organizational administrator role](/Overview/organizations-accounts.html#organizational-administrative-role) can configure this setting, and it can be modified after creation. Available options include the following:  
\- **`PUBLIC_ONLY`** : Allows access only through public APIs.  
\- **`PRIVATE_ONLY`** : Allows access only using AWS PrivateLink.  
\- **`PREFER_PUBLIC`** (Default): Prefers public APIs but can use AWS PrivateLink if needed.  
\- **`PREFER_PRIVATE`** : Prefers AWS PrivateLink but can use public APIs if needed. is\_enabled BOOLEAN Specifies if the service account is allowed to authenticate. created TIMESTAMPTZ Time of the service account creation. service\_account\_owner TEXT The name of the login that created the service account. If the service account was created by a service account, the service account name appears instead. last\_altered TIMESTAMPTZ Time the service account was last altered. last\_altered\_by TEXT The name of the login that edited the service account. If the service account was edited by a service account, the service account name appears instead.