# [](#show-catalogs)SHOW CATALOGS

Returns a table with a row for each catalog (i.e., database) defined in the current Firebolt account, with columns containing information as listed below.

## [](#syntax)Syntax

```
SHOW CATALOGS;
```

## [](#returns)Returns

The returned table has the following columns.

Column name Data Type Description catalog\_name TEXT The name of the catalog. catalog\_owner TEXT Current owner of the catalog. Default owner is the user who created the catalog created TIMESTAMPTZ Time the catalog was created (UTC) last\_altered TIMESTAMPTZ The date and time that the database was last modified (UTC) last\_altered\_by TEXT User/principal who edited the catalog ddl TEXT Complete DDL of the database. description TEXT User provided description of the database

## [](#example)Example

The following example shows information about CATALOGS in the account:

```
SHOW CATALOGS;
```

catalog\_name catalog\_owner created last\_altered last\_altered\_by ddl description AdTechDB\_v4 firebolt-demo 2024-09-03 11:48:31.683328+00 2024-09-03 11:48:31.683328+00 firebolt-demo CREATE DATABAE AdTechDB\_v4 () WITH DESCRIPTION='DB for AdTech' DB for AdTech