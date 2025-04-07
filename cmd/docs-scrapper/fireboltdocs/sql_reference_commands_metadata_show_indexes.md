# [](#show-indexes)SHOW INDEXES

Returns a table with a row for each Firebolt index defined in the current database, with columns containing information about each index as listed below.

## [](#syntax)Syntax

```
SHOW INDEXES;
```

## [](#returns)Returns

The returned table has the following columns.

Column name Data Type Description index\_name TEXT The name of the index. table\_name TEXT The name of the table associated with the index. type TEXT One of `primary` or `aggregating`. expression ARRAY (TEXT) An ordered array of the expression in SQL that defined the index. size\_compressed DOUBLE PRECISION The size of the index in bytes. size\_uncompressed DOUBLE PRECISION The uncompressed size of the index in bytes. compression\_ratio DOUBLE PRECISION The compression ratio (`<size_uncompressed>`/`<size_compressed>`). number\_of\_segments INTEGER The number of segments comprising the table.

## [](#example)Example

The following example returns information about indexes in the database queried:

```
SHOW INDEXES;
```

index\_name table\_name type expression size\_compressed size\_uncompressed compression\_ratio number\_of\_tablets players\_join\_idx players join \[“playerid”,”nickname”,”email”,”agecategory”] 819.98 KiB 819.98 KiB 1 0