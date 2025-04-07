# [](#truncate-table)TRUNCATE TABLE

Removes all rows from a table.

## [](#syntax)Syntax

```
TRUNCATE TABLE <table_name> 
```

## [](#parameters)Parameters

Parameter Description `<table_name>` The name of the table to be truncated.

### [](#example)Example

```
TRUNCATE TABLE product;
```

Table before

```
product
+------------+--------+
| name       | price  |
+---------------------+
| wand       |    125 |
| broomstick |    270 |
| bludger    |      0 |
| robe       |     80 |
| cauldron   |     25 |
| quaffle    |      0 |
+------------+--------+
```

Table after

```
product
+------------+--------+
| name       | price  |
+---------------------+
```