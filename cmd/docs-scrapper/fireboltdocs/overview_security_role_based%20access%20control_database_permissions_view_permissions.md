# [](#view-permissions)View permissions

In Firebolt, **views** are objects that allow users to query data from one or more underlying tables or views. Permissions on these views determine who can interact with the view and what actions they can perform.

To interact with a view, roles must also have **USAGE** permissions on the parent schema and the parent database.

## [](#view-level-privileges)View-level privileges

Privilege Description GRANT Syntax REVOKE Syntax   SELECT Allows selecting data from a view. `GRANT SELECT ON VIEW <view_name> TO <role_name>;` `REVOKE SELECT ON VIEW <view_name> FROM <role_name>;`   MODIFY Allows modifying and dropping a view. `GRANT MODIFY ON VIEW <view_name> TO <role_name>;` `REVOKE MODIFY ON VIEW <view_name> FROM <role_name>;`   ALL \[PRIVLEGES] Grants all privileges over the view to a role. `GRANT ALL ON VIEW <view_name> TO <role_name>;` `REVOKE ALL ON VIEW <view_name> FROM <role_name>;`  

Views are created at the schema level. To grant privileges to create views, refer to the [schema-level privileges documentation](/Overview/Security/Role-Based%20Access%20Control/database-permissions/schema-permissions.html).

## [](#examples-of-granting-view-permissions)Examples of granting view permissions

### [](#select-permission)SELECT permission

To allow querying data from a view, the role must have **SELECT** privileges on the view. Additionally, the **view owner** must have **SELECT** privileges on all underlying tables or views referenced within the view.

The following examples [grant](/sql_reference/commands/access-control/grant.html) the role `read_role` permission to query data from the `viewtest` view and ensure the `view_owner` has the necessary permission to read data from the `referenced_table` table, allowing the view to function correctly.

```
-- Grant SELECT on the view to a user:
GRANT SELECT ON VIEW "viewtest" TO read_role;

-- Grant SELECT on the referenced table to the view owner:
GRANT SELECT ON TABLE "referenced_table" TO view_owner;
```

If the **view owner** loses access to any of these referenced objects, users with **SELECT** on the view will no longer be able to query it, even if their **SELECT** privilege remains.

### [](#modify-permission)MODIFY permission

The following code example grants the role `developer_role` permission to alter or drop the `my_view` view:

```
GRANT MODIFY ON VIEW my_view TO read_role;
```

### [](#all-permissions)ALL permissions

The following code example grants the role `developer_role` with all permissions over the `my_view` view:

```
GRANT ALL ON VIEW my_view TO read_role;
```