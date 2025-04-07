# [](#create-database)CREATE DATABASE

Creates a new database.

Each account supports up to 100 databases. If you need more, contact Firebolt’s support team at [support@firebolt.io](mailto:support@firebolt.io).

## [](#syntax)Syntax

```
CREATE DATABASE [IF NOT EXISTS] <database_name>
[ WITH 
[ DESCRIPTION = <description> ]
]
```

## [](#parameters)Parameters

Parameter Description `<database_name>` The name of the database. `DESCRIPTION = 'description'` (Optional) The database’s description, which can contain up to 64 characters.

## [](#example)Example

The following code example creates a database named `my_db` with an optional description, `Testing database`:

```
CREATE DATABASE IF NOT EXISTS my_db
WITH DESCRIPTION = 'Testing database'
```