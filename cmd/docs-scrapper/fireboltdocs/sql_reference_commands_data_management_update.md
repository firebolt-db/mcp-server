# [](#update)UPDATE

Updates rows in the specified table.

## [](#syntax)Syntax

```
UPDATE <table> [ [ AS ] <alias> ] SET <column1> = <expression1> [, <column2> = <expression2> ...] [ FROM from_item [, ...] ] WHERE <condition>
```

## [](#parameters)Parameters

Parameter Description `<table>` The table name to update rows in. `<column>` The name of the column to be updated. `<expression>` An expression which computes a new value to populate the column. The expression can reference any column from the row being updated. `<from_item>` A table expression allowing columns from other tables to appear in the `WHERE` condition and update expressions. This uses the same syntax as the `FROM` clause of a `SELECT` statement; for example, an alias for the table name can be specified. Do not repeat the target table as a `from_item` unless you intend a self-join (in which case it must appear with an alias in the `from_item`). `<condition>` A Boolean expression. Only rows for which this expression returns `true` will be updated. Condition can have subqueries doing semi-join with other table(s).

## [](#remarks)Remarks

Updated rows are marked for deletion, but are not automatically cleaned up. You can monitor fragmentation in `information_schema.tables` to understand how many rows are marked for deletion out of total rows; fragmentation = (rows marked for deletion / total rows). Total row count in `information_schema.tables` excludes the number of rows marked for deletion. Query performance is not materially impacted by delete marks.

To mitigate fragmentation, use the [VACUUM](/sql_reference/commands/data-management/vacuum.html) command to manually clean up deleted rows.

When a `FROM` clause is present, what essentially happens is that the target table is joined to the tables mentioned in the `from_item` list, and each output row of the join represents an update operation for the target table. When using `FROM` the join must produce at most one output row for each row to be modified. In other words, a target row shouldn’t join to more than one row from the other table(s).

### [](#example-with-where)Example with WHERE

The following example restocks `product` for which `quantity` is less than `10`:

```
UPDATE product SET quantity = quantity + 10 WHERE quantity < 10
```

Table before:

name price quantity wand 120 9 robe 80 1 broomstick 270 21 cauldron 20 16 quill 5 100

Table after:

name price quantity wand 120 19 robe 80 11 broomstick 270 21 cauldron 20 16 quill 5 100

### [](#example-updating-multiple-columns)Example updating multiple columns

This example applies a discount and updates the quantity of a specific product.

```
UPDATE product SET price = 15, quantity = 100 WHERE name = 'cauldron'
```

Table before:

name price quantity wand 120 9 broomstick 270 21 robe 80 1 cauldron 20 16 quill 5 100

Table after:

name price quantity wand 120 9 broomstick 270 21 robe 80 1 cauldron 15 100 quill 5 100

### [](#example-updating-with-from)Example updating with FROM

This example updates available stock count.

```
UPDATE product AS p SET quantity = p.quantity - s.amount FROM sales AS s WHERE p.name = s.name
```

Table `product` before:

name price quantity wand 120 9 broomstick 270 21 robe 80 1 cauldron 20 16 quill 5 100

Table `sales` before:

name amount wand 5 broomstick 3 quill 20

Table `products` after:

name price quantity wand 120 4 broomstick 270 18 robe 80 1 cauldron 20 16 quill 5 80