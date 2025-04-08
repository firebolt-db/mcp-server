# [](#ownership)Ownership

Ownership allows users to perform all operations on any object they created without having to grant privileges for these operations manually. This provides a smoother user experience because objects are immediately available to use as they are created. These operations include granting privileges on owned objects.

## [](#supported-object-types)Supported object types

The object types that support ownership are:

- Role
- User
- Engine
- Database
- Schema
- Table
- View

The current owner of an object can be viewed in the corresponding information\_schema view:

Object View Role N/A User [information\_schema.users](/sql_reference/information-schema/users.html) Database [information\_schema.catalogs](/sql_reference/information-schema/catalogs.html) Engine [information\_schema.engines](/sql_reference/information-schema/engines.html) Schema [information\_schema.schemata](/sql_reference/information-schema/schemata.html) Table [information\_schema.tables](/sql_reference/information-schema/tables.html) View [information\_schema.views](/sql_reference/information-schema/views.html) or [information\_schema.tables](/sql_reference/information-schema/tables.html)

Index ownership, shown in [information\_schema.indexes](/sql_reference/information-schema/indexes.html), will always show the table owner as an index’s owner.

## [](#changing-an-objects-owner)Changing an object’s owner

The owner of an object may alter its ownership using the following syntax:

```
ALTER <object type>  <object name> OWNER TO <user>
```

Examples:

```
ALTER DATABASE db OWNER TO new_owner
ALTER ENGINE eng OWNER TO new_owner
ALTER ROLE r OWNER TO new_owner
ALTER USER u OWNER TO new_owner
ALTER SCHEMA public OWNER TO new_owner
ALTER TABLE t OWNER TO new_owner
ALTER VIEW v OWNER TO new_owner
```

## [](#dropping-users-that-own-objects)Dropping users that own objects

Any objects owned by a user must first be dropped or have their owner changed before dropping the user.

A table owner can drop the table even if there are views referencing it that are not owned by the table’s owner, using the `CASCADE` parameter to [DROP TABLE](/sql_reference/commands/data-definition/drop-table.html).

## [](#transfer-ownership-using-the-firebolt-workspace)Transfer ownership using the Firebolt Workspace

You can use the user interface in the **Firebolt Workspace** to transfer ownership of objects as follows:

1. Log in to the [Firebolt Workspace](https://firebolt.go.firebolt.io/signup). If you don’t yet have an account with Firebolt, you can sign up for one.
2. Select the Govern icon (![The icon to open the Govern Space.](../../assets/images/govern-icon.png)) in the left navigation pane to open the **Govern Space**.
3. Select **Ownership** from the left navigation pane.
4. Select the three horizontal dots (…) to the right of the object that you want to transfer ownership of.
5. Select **Transfer ownership** from the drop-down list.
6. In the **Transfer ownership** window that opens, choose a new owner from the drop-down list.
7. Select the **Transfer ownership** button to confirm.

##### [](#viewing-all-objects-owned-by-a-user)Viewing all objects owned by a user

1. From the [Firebolt Workspace](https://firebolt.go.firebolt.io/signup), select the Govern icon (![The icon to open the Govern Space.](../../assets/images/govern-icon.png)) in the left navigation pane to open the **Govern Space**.
2. Select **Users** from the left navigation pane.
3. Select the user from the **User Name** column.
4. Select the **Ownership** tab to view a list of objects owned by the selected user.

##### [](#bulk-transferring-or-deleting-objects-owned-by-a-user)Bulk transferring or deleting objects owned by a user

1. From the [Firebolt Workspace](https://firebolt.go.firebolt.io/signup), select the Govern icon (![The icon to open the Govern Space.](../../assets/images/govern-icon.png)) in the left navigation pane to open the **Govern Space**.
2. Select **Users** from the left navigation pane.
3. Select the three horizontal dots (…) to the right of the user whose objects you want to transfer ownership of.
4. Select **Transfer ownership** from the drop-down list.
5. In the window that opens, select the checkboxes next to objects that you want to delete or transfer ownership of.
6. Select the **Delete object** or **Transfer ownership** button to apply changes.

Ownership transfer using the **Firebolt Workspace** is not available for `Schema`, `Table`, and `View` objects. These must be modified using SQL commands in the **Develop Workspace** or using the [Firebolt API](/API-reference/).