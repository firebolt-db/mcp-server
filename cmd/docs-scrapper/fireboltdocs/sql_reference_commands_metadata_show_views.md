# [](#show-views)SHOW VIEWS

Lists the views defined in the current database and the `CREATE VIEW` statement that defines each view.

## [](#syntax)Syntax

```
SHOW VIEWS;
```

## [](#example)Example

The following example displays the view name and view definition of a views in the database queried:

**Returns**:

view\_name schema v14 CREATE VIEW “v14” AS SELECT a.* FROM (SELECT 1 AS “x”) AS “a” INNER JOIN (SELECT 1 AS “x”) AS “b” USING(x) v15 CREATE VIEW IF NOT EXISTS “v15” AS SELECT * FROM “bf\_test\_t” WHERE ( “n” = 0 ) v16 CREATE VIEW “v16” AS WITH x7 AS (SELECT * FROM “oz\_x6” ) SELECT * FROM “x7”