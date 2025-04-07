# [](#drop-network-policy)DROP NETWORK POLICY

Deletes a network policy.

For more information, see [Network policies](/Guides/security/network-policies.html).

## [](#syntax)Syntax

```
DROP NETWORK POLICY <network_policy_name> [ RESTRICT | CASCADE ]
```

## [](#parameters)Parameters

Parameter Description `<network_policy_name>` The name of the network policy to delete. `RESTRICT` or `CASCADE` An optional parameter to specify deletion mode.  
RESTRICT mode prevents dropping the network policy if there is any login, service account or organization linked. RESTRICT mode is used by default.  
CASCADE mode automatically drops the network policy and all its links to other objects.

## [](#example)Example

The following command will delete “my\_network\_policy”.

```
DROP NETWORK POLICY my_network_policy [ RESTRICT | CASCADE ]
```