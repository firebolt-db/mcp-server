# [](#drop-table)DROP TABLE

Deletes a table.

## [](#syntax)Syntax

```
DROP TABLE [IF EXISTS] <table_name> [CASCADE]
```

## [](#parameters)Parameters

Parameter Description `<table_name>` The name of the table to be deleted. For external tables, the definition is removed from Firebolt but not from the source. `CASCADE` When specified, causes all dependent database objects such as views and aggregating indexes to be dropped also.