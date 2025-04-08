# [](#information-schema-for-network_policies)Information schema for network\_policies

You can use the `information_schema.network_policies` view to return information about network policies. You can use a `SELECT` query to return information about each network policy as shown in the example below.

```
SELECT
  *
FROM
  information_schema.network_policies;
```

Read more about network policies [here](/Guides/security/network-policies.html).

## [](#columns-in-information_schemanetwork_policies)Columns in information\_schema.network\_policies

Each row has the following columns with information about a network policy.

Column Name Data Type Description network\_policy\_name TEXT The name of the network policy. allowed\_ips ARRAY(TEXT) List of allowed ips blocked\_ips ARRAY(TEXT) List of blocked ips is\_organizational BOOLEAN Specifies if the network policy is active at the organization level . network\_policy\_description TEXT The description of the network policy. created TIMESTAMPTZ Time the service account was created. network\_policy\_owner TEXT The name of the login that created the network policy. If the network policy was created by a service account, the service account name appears instead. last\_altered TIMESTAMPTZ Time the service account was last altered. last\_altered\_by TEXT The name of the login that edited the network policy. If the network policy was edited by a service account, the service account name appears instead.