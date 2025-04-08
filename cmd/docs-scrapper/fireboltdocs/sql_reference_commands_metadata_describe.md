# [](#describe)DESCRIBE

Lists all columns and data types for the table. Once the results are displayed, you can also export them to CSV or JSON.

## [](#syntax)Syntax

```
DESCRIBE <table_name>
```

## [](#parameters)Parameters

Parameter Description `<table_name>` The name of the table to be described.

## [](#example)Example

The following lists all columns and data types for the table named `players`:

```
DESCRIBE prices
```

**Returns:**

table\_name column\_name data\_type nullable players agecategory TEXT 0 players email INTEGER 0 players nickname TEXT 0