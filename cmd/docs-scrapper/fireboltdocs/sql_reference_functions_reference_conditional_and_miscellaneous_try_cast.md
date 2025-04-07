# [](#try_cast)TRY\_CAST

Converts data types into other data types based on the specified parameters. If the conversion cannot be performed, returns a NULL. To return an error message instead, use [CAST](/sql_reference/functions-reference/conditional-and-miscellaneous/cast.html).

TRY\_CAST replaces only execution errors with NULLs. However, during planning, impossible casts between two non-castable types still produce an error because the query is invalid.

## [](#syntax)Syntax

```
TRY_CAST(<value> AS <type>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The value to convert or an expression that results in a value to convert Any `<type>` The target [data type](/sql_reference/data-types.html) (case-insensitive) Any

## [](#return-type)Return Type

Returns `NULL` if the conversion cannot be performed. Otherwise, returns the data type of `<type>`.

## [](#example)Example

The following example attempts to cast the level input as an integer:

```
SELECT TRY_CAST('1' AS INTEGER) as level, TRY_CAST('level 2' AS INTEGER) as current_level;
```

**Returns**: `1, null`