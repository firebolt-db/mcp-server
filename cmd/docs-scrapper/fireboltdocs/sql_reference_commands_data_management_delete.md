# [](#delete)DELETE

Deletes rows from the specified table.

## [](#syntax)Syntax

```
DELETE FROM <table> [[AS] <alias>] [USING <from_item>] WHERE <condition>
```

## [](#parameters)Parameters

Parameter Description `<table>` The table to delete rows from. `<from_item>` A table expression allowing columns from other tables to appear in the `WHERE` condition. This uses the same syntax as the `FROM` clause of a `SELECT` statement; for example, an alias for the table name can be specified. Do not repeat the target table as a `from_item` unless you wish to set up a self-join (in which case it must appear with an alias in the `from_item`). `<condition>` A Boolean expression. Only rows for which this expression returns `true` will be deleted. Condition can have subqueries doing semi-join with other table(s).

The `DELETE FROM <table>` without `<expression>` will delete *all* rows from the table. It is equivalent to a [TRUNCATE TABLE](/sql_reference/commands/data-management/truncate-table.html) statement.

## [](#remarks)Remarks

Deleted rows are marked for deletion, but are not automatically cleaned up. You can monitor fragmentation in [information\_schema.tables](/sql_reference/information-schema/tables.html) to understand how many rows are marked for deletion out of total rows; fragmentation = (rows marked for deletion / total rows). Total row count in `information_schema.tables` excludes the number of rows marked for deletion. Query performance is not materially impacted by delete marks.

To mitigate fragmentation, use the [VACUUM](/sql_reference/commands/data-management/vacuum.html) command to manually clean up deleted rows.

## [](#example)Example

The following example deletes entries from the `products` table where the `quantity` is less than `10`:

```
DELETE FROM products WHERE quantity < 10
```

Table before:

product quantity wand 9 broomstick 21 robe 1 quidditch gloves 10 cauldron 16 quill 100

Table after:

product quantity broomstick 21 quidditch gloves 10 cauldron 16 quill 100

### [](#example-specifying-other-tables-in-the-using-clause)Example specifying other tables in the USING clause

This example deletes all the products from stores that went out of business.

```
DELETE FROM products USING suppliers WHERE products.product = suppliers.product AND suppliers.store = 'Ollivanders'
```

Table `products` before:

product quantity wand 9 broomstick 21 robe 1 quidditch gloves 10 cauldron 16 quill 100

Table `suppliers` before:

product store wand Ollivanders broomstick Quality Quidditch Supplies robe Madam Malkin’s quidditch gloves Quality Quidditch Supplies cauldron Apothecary quill Amanuensis Quills

Table `products` after:

product quantity broomstick 21 robe 1 quidditch gloves 10 cauldron 16 quill 100