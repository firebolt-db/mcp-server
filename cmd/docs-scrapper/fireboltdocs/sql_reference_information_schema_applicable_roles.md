# [](#information-schema-for-applicable_roles)Information schema for applicable\_roles

The `information_schema.applicable_roles` view shows every role in the account and its grantees, who include other users or roles to whom the role is granted.

You can use a `SELECT` query to return information about each role as shown in the example below.

```
SELECT
  *
FROM
  information_schema.applicable_roles;
```

See also `information_schema.transitive_applicable_roles` [here](/sql_reference/information-schema/transitive-applicable-roles.html).

Read more about RBAC roles [here](/Guides/security/rbac.html).

## [](#columns-in-information_schemaapplicable_roles)Columns in information\_schema.applicable\_roles

Each row has the following columns with information about the role.

Column Name Data Type Description grantee TEXT User or role to whom the role is granted. role\_name TEXT Name of the role. is\_grantable TEXT `YES` if the grantee has the admin option on the role, `NO` if not. created TIMESTAMPTZ Creation time of the role.