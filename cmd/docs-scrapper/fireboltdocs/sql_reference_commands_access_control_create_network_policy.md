# [](#create-network-policy)CREATE NETWORK POLICY

Creates a new network policy.

For more information, see [Network policies](/Guides/security/network-policies.html).

## [](#syntax)Syntax

```
CREATE NETWORK POLICY [IF NOT EXISTS] <network_policy_name>
WITH ALLOWED_IP_LIST = ( '<allowed_ip1>', '<allowed_ip2>', ... )
[ BLOCKED_IP_LIST = ( '<blocked_ip1>', '<blocked_ip2>', ... ) ]
[ DESCRIPTION = '<description>' ]
```

## [](#parameters)Parameters

Parameter Description `<network_policy_name>` The name of the network policy. Must start with a letter, and may contain only alphanumeric and underscore(\_) characters. `<allowed_ip>` The comma-separated and quoted list of IP addresses to allow in the created network policy. `<blocked_ip>` An optional comma-separated and quoted list of IP addresses to block in the created network policy. `<description>` An optional description for the created network policy.

## [](#example)Example

The following command will create a network policy that allows IPs ‘4.5.6.1’ and ‘2.4.5.1’ and blocks the IP address ‘6.7.8.1’, with a description:

```
CREATE NETWORK POLICY my_network_policy WITH ALLOWED_IP_LIST = ('4.5.6.1', '2.4.5.1') DESCRIPTION = 'my new network policy'
```