# [](#cast)CAST

Converts data types into other data types based on specified parameters. If the conversion cannot be performed, `CAST` returns an error. To return a `NULL` value instead, use [TRY\_CAST](/sql_reference/functions-reference/conditional-and-miscellaneous/try-cast.html).

## [](#syntax)Syntax

```
CAST(<value> AS <type>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The value to convert or an expression that results in a value to convert. Any `<type>` The target [data type](/sql_reference/data-types.html) (case-insensitive) Any

## [](#return-types)Return Types

Same data type as `<type>`

## [](#example)Example

The following example returns `1` as an integer:

```
SELECT CAST('1' AS INTEGER) as level;
```

**Returns**: `1`

`CAST` can also be done by writing the format before the object, for example - `SELECT DATE '2022-01-01'` , `SELECT TIMESTAMP '2022-01-01 01:02:03'.`

`CAST` can also be done by using the `::` operator. For more information, see [:: operator for CAST](/sql_reference/operators.html#-type-cast).