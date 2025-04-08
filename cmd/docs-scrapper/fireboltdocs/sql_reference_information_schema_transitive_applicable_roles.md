# [](#information-schema-for-transitive_applicable_roles)Information schema for transitive\_applicable\_roles

The `information_schema.applicable_roles` view shows every role in the account and its grantees, who include other users or roles to whom the role is granted.

Unlike, `information_schema.applicable_roles` that shows only direct grantees, `information_schema.transitive_applicable_roles` also shows indirect grantees. For example, if role `engineer` is granted to role `manager` and role `manager` is granted to user `alice` then user `alice` is a direct grantee of `manager` and an indirect grantee of `engineer`.

You can use a `SELECT` query to return information about each role as shown in the example below.

```
SELECT
  *
FROM
  information_schema.transitive_applicable_roles;
```

See also `information_schema.applicable_roles` [here](/sql_reference/information-schema/applicable-roles.html).

Read more about RBAC roles [here](/Guides/security/rbac.html#check-assigned-privileges-using-sql).

## [](#columns-in-information_schematransitive_applicable_roles)Columns in information\_schema.transitive\_applicable\_roles

Each row has the following columns with information about the role.

Column Name Data Type Description grantee TEXT User or role to whom the role is granted (directly or indirectly). role\_name TEXT Name of the role. is\_grantable TEXT `YES` if the grantee has the admin option on the role, `NO` if not. created TIMESTAMPTZ Creation time of the role.