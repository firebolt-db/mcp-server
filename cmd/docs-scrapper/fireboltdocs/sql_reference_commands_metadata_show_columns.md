# [](#show-columns)SHOW COLUMNS

Lists columns and their properties for a specified table. Returns `<table_name>`, `<column_name>`, `<data_type>`, and `nullable` (`TRUE` if nullable, `FALSE` if not) for each column.

## [](#syntax)Syntax

```
SHOW COLUMNS <table>;
```

## [](#parameters)Parameters

Parameter Description `<table>` The name of the table to be analyzed.

## [](#example)Example

The following example highlights all of the columns from the `tournaments` table:

```
SHOW COLUMNS tournaments;
```

table\_name column\_name data\_type nullable tournaments enddatetime TIMESTAMP FALSE tournaments gameid INTEGER FALSE tournaments name TEXT FALSE tournaments rulesdefinition TEXT FALSE tournaments SOURCE\_FILE\_NAME TEXT FALSE tournaments SOURCE\_FILE\_TIMESTAMP TIMESTAMP FALSE tournaments startdatetime TIMESTAMP FALSE tournaments totalprizedollars INTEGER FALSE tournaments tournamentid INTEGER FALSE