# [](#default-system-roles)Default system roles

In Firebolt, **system-defined** roles are automatically created for each organization and account. These roles provide predefined privileges and serve specific purposes. While system-defined roles **cannot** be modified or dropped, you can grant them additional privileges as needed.

## [](#organization-system-roles)Organization system roles

Role Name Description organization\_admin Enables all the permissions and the ability to manage the organization.

The [organization\_admin](/Overview/organizations-accounts.html#organizational-administrative-role) role cannot be granted using SQL. It can only be granted using the [Firebolt Workspace](https://go.firebolt.io/signup) user interface (UI). To manage resources at the organization level, you must assign the `organization_admin` role to your login using the UI.

## [](#account-system-roles)Account system roles

Role Name Description public Includes `USAGE` on all databases and both `USAGE` and `CREATE` on every public schema. system\_admin Enables managing databases, engines, schemas, tables, and views. This includes setting database and engine properties as well as access to the observability functionality on all engines. account\_admin Grants full permissions to manage the organization.

By default, every newly created user is granted the [public](/Overview/organizations-accounts.html#public-role) role. You can also revoke this role from a user.