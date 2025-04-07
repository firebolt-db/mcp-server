# [](#ownership)Ownership

When you create an object in Firebolt, you become its owner. As the owner, you have full control and can perform all operations on the object without needing additional privileges. This allows you to use objects immediately after creating them.

As the owner of an object, you can do the following:

- Grant privileges on the object to any role.
- Grant roles you own to other users or roles without needing administrator permissions.

## [](#supported-object-types)Supported object types

The following object types support ownership:

- **Role**
- **User**
- **Engine**
- **Database**
- **Schema**
- **Table**
- **View**

The current owner of an object can be viewed in the corresponding `information_schema` view:

Role Name View Role N/A User `information_schema.users` Database `information_schema.catalogs` Engine `information_schema.engines` Schema `information_schema.schemata` Table `information_schema.tables` View `information_schema.views` or `information_schema.tables`

Indexes inherit ownership from their parent table. In `information_schema.indexes`, the table owner is displayed as the index owner.

## [](#changing-an-objects-owner)Changing an object’s owner

The following code example shows how to transfer the ownership of an object:

```
ALTER <object type>  <object name> OWNER TO <user>
```

## [](#examples-of-updating-ownership-permissions)Examples of updating ownership permissions

The following are examples of how to change the ownership of different object types.

### [](#changing-object-ownership)Changing object ownership

The following code example uses `ALTER` to transfer ownership of a database, engine, role, user, schema, table, and view to a new owner:

```
ALTER DATABASE db OWNER TO new_owner
ALTER ENGINE eng OWNER TO new_owner
ALTER ROLE r OWNER TO new_owner
ALTER USER u OWNER TO new_owner
ALTER SCHEMA public OWNER TO new_owner
ALTER TABLE t OWNER TO new_owner
ALTER VIEW v OWNER TO new_owner
```

### [](#dropping-users-that-own-objects)Dropping users that own objects

Before dropping a user who owns objects, you must either drop the objects owned by the owner or transfer ownership of them to another user.

The following code example shows how to drop a table that has dependent views not owned by the table’s owner using the `CASCADE` parameter to enforce the drop:

```
DROP TABLE <table_name> CASCADE;
```