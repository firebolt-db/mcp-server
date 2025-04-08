# [](#alter-database)ALTER DATABASE

Updates the configuration of the specified database.

## [](#alter-database-description)ALTER DATABASE DESCRIPTION

### [](#syntax)Syntax

```
ALTER DATABASE <database_name> WITH
    [DESCRIPTION = <description>]
```

### [](#parameters)Parameters

Parameter Description `<database_name>` The name of the database to be altered. `<description>` The description of the database.

### [](#example)Example

The following example alters a description of the database:

```
ALTER DATABASE my_database WITH DESCRIPTION = 'Database for query management';
```

## [](#alter-database-owner-to)ALTER DATABASE OWNER TO

Change the owner of a database. The current owner of a database can be viewed in the [information\_schema.catalogs](/sql_reference/information-schema/catalogs.html) view on `catalog_owner` column.

check [ownership](/Guides/security/ownership.html) page for more info.

### [](#syntax-1)Syntax

```
ALTER DATABASE <database_name> OWNER TO <user>
```

### [](#parameters-1)Parameters

Parameter Description `<database_name>` The name of the database to change the owner of. `<user>` The new owner of the database.