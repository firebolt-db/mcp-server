# [](#manage-role-based-access-control)Manage role-based access control

Role-Based Access Control (RBAC) allows you to manage user permissions by controlling who can access or perform operations on specific objects in Firebolt. This guide provides a step-by-step process for setting RBAC in Firebolt.

## [](#prerequisites)Prerequisites

The following material can help you understand key concepts related to organizations and RBAC in Firebolt:

- [Organizations and accounts](/Overview/organizations-accounts.html) – How Firebolt provides a structure for managing users, resources, and permissions.
- [Role-Based Access Control](/Overview/Security/Role-Based%20Access%20Control/) – How administrators manage user permissions and control access to resources based on predefined roles.

## [](#view-all-roles)View all roles

To view all roles using the **Firebolt Workspace**, do the following:

1. Login to the [Firebolt Workspace](https://firebolt.go.firebolt.io/signup).
2. Select the **Govern** icon (![The Firebolt Govern Space icon.](../../assets/images/govern-icon.png))from the left navigation bar to open the **Govern Space**.
3. Choose **Roles** from the left panel under **Govern**.

To view all roles using SQL, query the [information\_schema.applicable\_roles](/sql_reference/information-schema/applicable-roles.html) view as shown in the following code example:

```
SELECT
  *
FROM
  information_schema.applicable_roles;
```

## [](#create-a-role)Create a role

You can create a role using the **Firebolt Workspace** user interface (UI) or using SQL.

### [](#create-a-role-using-sql)Create a role using SQL

The following code example uses [CREATE ROLE](/sql_reference/commands/access-control/create-role.html) to create the role `user_role`:

```
CREATE ROLE user_role;
```

### [](#create-a-role-using-the-ui)Create a role using the UI

To create a custom role using the UI:

1. Select the **Govern** icon (![The Firebolt Govern Space icon.](../../assets/images/govern-icon.png)) from the left navigation bar to open the **Govern Space**.
2. Choose **Roles** from the left panel under **Govern**.
3. Choose the **+ New Role** button in the upper-right corner of the page.
4. Under **Create role**, enter a role name.
5. Under **Role privileges:**, select the object type that you want to grant permissions for. You can choose either **Databases** or **Engines**.
6. Configure permissions for the role:
   
   - Toggle the buttons under **Databases privileges** or **Engine privileges** to grant permissions to **create** or **modify** an objects across all databases or engines. If you want to apply permissions to a specific engine or database, select it from the table under the toggle buttons.
   - Define permissions more granularly using table views.

## [](#delete-a-role)Delete a role

You can delete a role using either the UI in the **Govern Workspace** or using SQL.

### [](#delete-a-role-using-sql)Delete a role using SQL

To delete a role using SQL, use [DROP ROLE](/sql_reference/commands/access-control/drop-role.html) as shown in the following code example:

```
DROP ROLE user_role;
```

### [](#delete-a-role-using-the-ui)Delete a role using the UI

To delete a role via the UI:

1. Select the **Govern** icon (![The Firebolt Govern Space icon.](../../assets/images/govern-icon.png)) from the left navigation bar to open the **Govern Space**.
2. Choose **Roles** from the left panel under **Govern**.
3. Search for the relevant role using the top search filters or by scrolling through the list. Hover over the right-most column to make the role menu appear, then choose **Delete role**.
4. Choose **Confirm**.

## [](#grant-permissions-to-a-role)Grant permissions to a role

### [](#grant-permissions-using-sql)Grant permissions using SQL

To grant a permission to a role using SQL, use [GRANT](/sql_reference/commands/access-control/grant.html) as shown in the following code example:

```
GRANT USAGE ON DATABASE my_db TO user_role;
```

### [](#grant-permissions-using-the-ui)Grant permissions using the UI

To grant a permission to a role via the UI:

1. Select **Govern** to open the govern space, then choose **Roles** from the menu:
2. Search for the relevant role either by using the search filters at the top of the page, or by scrolling through the list of logins. Hover over the right-most column to make the role menu appear, then choose **Edit role**.
3. Navigate to the permissions tab and select the desired permissions. To grant permissions over all objects of that type, choose the topmost line.
4. Select **Update**.

## [](#grant-a-role-to-users)Grant a role to users

### [](#grant-a-role-to-users-using-sql)Grant a role to users using SQL

To grant a role to a user or another role using SQL, use [GRANT ROLE](/sql_reference/commands/access-control/grant.html) as shown in the following code example:

```
GRANT ROLE user_role TO ROLE user2_role;
```

### [](#grant-a-role-using-the-ui)Grant a role using the UI

To grant a role to a user via the UI:

1. Select **Govern**, then choose **Users** from the menu:
2. In the user’s row, select the three horizontal dots to the right.
3. Select **Edit user details**.
4. Select the drop-down list next to **Role**.
5. Select the checkbox next to the roles that you want to grant.
6. Select **Edit user**.

## [](#revoke-permissions)Revoke permissions

You can revoke permissions using the UI in the **Govern Space** or using SQL.

### [](#revoke-permissions-using-sql)Revoke permissions using SQL

To revoke a permission from a role using SQL, use [REVOKE](/sql_reference/commands/access-control/revoke.html) as shown in the following example:

```
REVOKE USAGE ON DATABASE my_db FROM user_role;
```

### [](#revoke-permissions-using-the-ui)Revoke permissions using the UI

To revoke permissions, follow the same steps described in [Grant permissions to a role](#grant-permissions-to-a-role).

## [](#revoke-role)Revoke role

You can revoke a role from either a user or another role using either the UI in the **Govern Space** or SQL.

### [](#revoke-a-role-using-sql)Revoke a role using SQL

To revoke a role from a user or another role using SQL, use the [REVOKE ROLE](/sql_reference/commands/access-control/revoke.html) statement. For example:

```
REVOKE ROLE user_role FROM USER alex;
```

### [](#revoke-a-role-using-the-ui)Revoke a role using the UI

To revoke a role, follow the steps in [Grant a role to users](#grant-a-role-to-users).

### [](#check-assigned-privileges-using-sql)Check assigned privileges using SQL

To check the effective privileges for the current user, run the following example query:

```
SELECT
  AR.grantee,
  AR.role_name,
  OP.privilege_type,
  OP.object_type,
  OP.object_name
FROM information_schema.transitive_applicable_roles AS AR
JOIN information_schema.object_privileges AS OP
ON (AR.role_name = OP.grantee)
WHERE
  AR.grantee = session_user();
```

**Returns**:

grantee role\_name privilege\_type object\_type object\_name test\_user account\_admin USAGE engine engine1 test\_user account\_admin USAGE database db1

#### [](#owner-rights)Owner rights

When a query is run on a view, the database checks and uses the permissions of the view’s owner to access the underlying objects that view references, rather than the permissions of the user that ran the query on the view. The view’s owner is the user that created the view.

The following code example shows how granting and revoking privileges affects access to a base table and its view, ultimately causing an authorization failure when the view’s owner loses schema usage privileges:

```
CREATE USER user1 WITH ROLE=role1;
CREATE USER user2 WITH ROLE=role2;

CREATE TABLE base_table (a int); -- executed by user1
CREATE VIEW view_over_base_table AS SELECT * FROM base_table; -- executed by user1

GRANT SELECT ON VIEW view_over_base_table TO role2;
REVOKE SELECT ON TABLE base_table FROM role2;

SELECT * FROM base_table; -- executed by user2, fails with an authorization error
SELECT * FROM view_over_base_table; -- executed by user2, successfully

REVOKE USAGE ON SCHEMA public FROM role1;
-- role1 no longer has no access to the table due to missing schema usage privileges
SELECT * FROM view_over_base_table; -- executed by user2 and fails because the view owner's role1 cannot access table t
```

If the view owner’s privileges are revoked, the query will fail even if the user has access to the view.