# [](#create-service-account)CREATE SERVICE ACCOUNT

Creates a new service account.

For more information, see [Service accounts](/Guides/managing-your-organization/service-accounts.html).

## [](#syntax)Syntax

```
CREATE SERVICE ACCOUNT [ IF NOT EXISTS ] <service_account_name> 
    [ WITH
    [ DESCRIPTION = <description> ] 
    [ NETWORK_POLICY = <network_policy_name> ]
    [ IS_ORGANIZATION_ADMIN = { TRUE|FALSE } ]
    [ CONNECTION_PREFERENCE = { PUBLIC_ONLY | PRIVATE_ONLY | PREFER_PUBLIC | PREFER_PRIVATE | DEFAULT } ]
    ]
```

## [](#parameters)Parameters

Parameter Description `<service_account_name>` The name of the service account. Must start with a letter and may contain only alphanumeric, digit, or underscore (\_) characters. `<description>` An optional description for the service account. `<network_policy_name>` An optional parameter to define the network policy to link to the created service account. `IS_ORGANIZATION_ADMIN` A `BOOLEAN` value specifying if the login is an organization admin. By default, this value is `FALSE`. `CONNECTION_PREFERENCE` Defines how the service account connects to Firebolt. The default value is `PREFER_PUBLIC` if not specified. See **Connection preferences** for details.

### [](#connection-preferences)Connection preferences

The `CONNECTION_PREFERENCE` parameter determines how a [service account](/Guides/managing-your-organization/service-accounts.html) accesses Firebolt:

- **`PUBLIC_ONLY`** : Allows access only through public APIs.
- **`PRIVATE_ONLY`** : Allows access only through AWS PrivateLink.
- **`PREFER_PUBLIC`** (Default): Prefers public APIs but can use AWS PrivateLink if needed.
- **`PREFER_PRIVATE`** : Prefers AWS PrivateLink but can use public APIs if needed.

## [](#example)Example

The following code example creates a service account `sa1` linked to the `my_network_policy` network policy:

```
CREATE SERVICE ACCOUNT "sa1" WITH NETWORK_POLICY = my_network_policy
```