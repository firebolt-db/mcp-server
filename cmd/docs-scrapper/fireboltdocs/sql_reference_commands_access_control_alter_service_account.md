# [](#alter-service-account)ALTER SERVICE ACCOUNT

Updates the configuration of the specified service account.

For more information, see [Service accounts](/Guides/managing-your-organization/service-accounts.html).

## [](#syntax)Syntax

```
ALTER SERVICE ACCOUNT <service_account_name> SET 
      [ NETWORK_POLICY = <network_policy_name> | DEFAULT ] 
      [ DESCRIPTION = <description> | DEFAULT ]
      [ IS_ORGANIZATION_ADMIN = { TRUE | FALSE } ]
      [ CONNECTION_PREFERENCE = { PUBLIC_ONLY | PRIVATE_ONLY | PREFER_PUBLIC | PREFER_PRIVATE | DEFAULT } ]
      [ IS_ENABLED = { TRUE | FALSE } ]
```

or

```
ALTER SERVICE ACCOUNT <service_account_name> RENAME TO <new_name>;
```

## [](#parameters)Parameters

Parameter Description `<service_account_name>` The name of the service account to edit. `<description>` An optional description for the service account. `<network_policy_name>` An optional parameter to define the network policy to link to the edited service account. Specifying `DEFAULT` will detach any linked network policy. `IS_ORGANIZATION_ADMIN` A `BOOLEAN` value specifying if the service account is an organization admin. By default, this value is `FALSE`. `IS_ENABLED` A `BOOLEAN` value specifying whether authentication with this service account should be possible. Disable the service account to prevent access without dropping it. `CONNECTION_PREFERENCE` Defines how the service account connects to Firebolt. The default value is `PREFER_PUBLIC` if not specified. See **Connection preferences** for details. `<new_name>` The new name of the service account. Must start with a letter and may contain only alphanumeric, digit, or underscore (\_) characters.

### [](#connection-preferences)Connection Preferences

The `CONNECTION_PREFERENCE` parameter determines how a [service account](/Guides/managing-your-organization/service-accounts.html) accesses Firebolt:

- **`PUBLIC_ONLY`** : Allows access only through public APIs.
- **`PRIVATE_ONLY`** : Allows access only through AWS PrivateLink.
- **`PREFER_PUBLIC`** (Default): Prefers public APIs but can use AWS PrivateLink if needed.
- **`PREFER_PRIVATE`** : Prefers AWS PrivateLink but can use public APIs if needed.

## [](#example)Example

The following code example renames the service account `sa1` to `serviceaccount1`:

```
ALTER SERVICE ACCOUNT "sa1" RENAME TO "serviceaccount1";
```

The following code example assigns the network policy `my_network_policy` to the `serviceaccount1` service account:

```
ALTER SERVICE ACCOUNT "serviceaccount1" SET NETWORK_POLICY = "my_network_policy";
```