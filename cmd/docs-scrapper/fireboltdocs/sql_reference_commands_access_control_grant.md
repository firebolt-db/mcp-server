# [](#grant)GRANT

Grants permission or assignment to a role. `GRANT` can also be used to assign a role to another role or a user.

For more information, see [Role-based access control](/Overview/Security/Role-Based%20Access%20Control/role-management/).

## [](#grant-privilege)GRANT PRIVILEGE

Grants a permission to a role.

Only an account\_admin or a role owner can grant a permission to a role. To grant a permission, you must first have that permission granted to you.

### [](#syntax)Syntax

```
GRANT <permission> ON <object_type> <object_name> [IN <object_type> <object_name>] TO <role_name>
```

### [](#parameters)Parameters

Parameter Description `<permission>` The name of the permission to grant to a role. Available permissions vary depending on the object that they apply to. For a full list, see [Permissions](/Overview/Security/Role-Based%20Access%20Control/role-management/). `<object_type>` The type of object to grant permissions on. `<object_name>` The name of the object to grant permissions on. `<role_name>` The name of the role to grant the permission to.

### [](#examples)Examples

**Grant `USAGE` on a single database**

The following code example grants the `USAGE` privilege on the `db` database to the role `user_role`, allowing it to access the database:

```
GRANT USAGE ON DATABASE db TO user_role;
```

**Grant `USAGE` on all databases within an account**

The following code example grants the `USAGE` privilege on all databases in the `dev` account to the role `user_role`, allowing access to them:

```
GRANT USAGE ANY DATABASE ON ACCOUNT dev TO user_role;
```

**Grant access to a databse, schema, and a specific table**

The following code example grants the role `user_role` access to the `db` database, the `public` schema within the `db` database, and permission to read data from the `my_table` table in the `public` schema:

```
GRANT USAGE ON DATABASE db TO user_role;
GRANT USAGE ON SCHEMA public IN DATABASE db TO user_role;
USE DATABASE db;
GRANT SELECT ON TABLE my_table IN SCHEMA public TO user_role;
```

**Grant access to a database, schema, and all operations on a specific table**

The following code example grants the role `user_role` access to the `db` database, the `public` schema within the `db` database, and all permissions on the `my_table` table in the `public` schema:

```
GRANT USAGE ON DATABASE db TO user_role;
GRANT USAGE ON SCHEMA public IN DATABASE db TO user_role;
USE DATABASE db;
GRANT ALL ON TABLE my_table IN SCHEMA public TO user_role;
```

**Grant access to all existing and future tables or views in a schema**

The following code example grants `user_role` access to the `db` database, the `public` schema within the `db` database, and permission to query all existing and future tables or views in the `public` schema:

```
GRANT USAGE ON DATABASE db TO user_role;
GRANT USAGE ON SCHEMA public IN DATABASE db TO user_role;
GRANT SELECT ANY ON SCHEMA public IN DATABASE db TO user_role;
```

## [](#grant-role)GRANT ROLE

Grants a role to either a user or another role, allowing the recipient to inherit the permissions associated with the granted role.

### [](#syntax-1)Syntax

```
GRANT ROLE <role_name> TO { USER <user_name> | ROLE <role_name_2> }
```

### [](#parameters-1)Parameters

Parameter Description `<role_name>` The name of the role to grant. `<user_name>` The name of the user to grant `<role_name>` to. `<role_name_2>` The name of the role to assign the role to.

### [](#examples-1)Examples

**Grant a role to another role**

The following code example assigns the `role_name` role to `role_name_2`, allowing `role_name_2` to inherit all the permissions granted to `role_name`:

```
GRANT ROLE role_name TO ROLE role_name_2;
```

**Grant a role to a user**

The following code example assigns the `role_name` role to `user_name`, allowing the user to inherit all the permissions granted to `role_name`:

```
GRANT ROLE role_name TO USER user_name;
```