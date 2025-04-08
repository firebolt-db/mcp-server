# [](#show-tables)SHOW TABLES

Returns a table with a row for each table in the current database, with columns containing information for each table as listed below.

## [](#syntax)Syntax

```
SHOW TABLES;
```

## [](#returns)Returns

The returned table has the following columns.

Column name Data Type Description table\_name TEXT The name of the table. state TEXT The current table state. table\_type TEXT One of `FACT`, `DIMENSION`, or `EXTERNAL`. column\_count INTEGER The number of columns in the table. primary\_index TEXT An ordered array of the column names comprising the primary index definition, if applicable. schema TEXT The text of the SQL statement that created the table. number\_of\_rows INTEGER The number of rows in the table. size DOUBLE PRECISION The compressed size of the table. size\_uncompressed DOUBLE PRECISION The uncompressed size of the table. compression\_ratio DOUBLE PRECISION The compression ratio (`<size_uncompressed>`/`<size>`). number\_of\_tablets INTEGER The number of tablets comprising the table.

## [](#example)Example

The following example returns information about tables in the database queried:

```
SHOW TABLES;
```

table\_name state table\_type column\_count primary\_index schema number\_of\_rows size size\_uncompressed compression\_ratio number\_of\_tablets ex\_games Valid EXTERNAL 1 “CREATE EXTERNAL TABLE IF NOT EXISTS ““ex\_games”” (““src”” text NOT NULL) "”OBJECT\_PATTERN”” = ‘help\_center\_assets/firebolt\_sample\_dataset/games.json’ ““TYPE”” = (““JSON”” ““PARSE\_AS\_TEXT”” = ‘TRUE’) ““URL”” = ‘s3://firebolt-publishing-public/help\_center\_assets/firebolt\_sample\_dataset/’” 0 0.00 B 0.00 B 0 0