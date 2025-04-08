# [](#length)LENGTH

Calculates the length of the input string.

When used with an array argument, `LENGTH` is a synonym for [ARRAY\_LENGTH](/sql_reference/functions-reference/array/array-length.html)

## [](#syntax)Syntax

```
LENGTH(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The string or binary data for which to return the length. `TEXT`, `BYTEA`. For `ARRAY`, see [ARRAY\_LENGTH](/sql_reference/functions-reference/array/array-length.html)

## [](#return-type)Return Type

`INTEGER`

## [](#example)Example

Use the `LENGTH` to find the length of any string, such as:

```
SELECT LENGTH('The Accelerator Cup')
```

Spaces are included in the calculation of the total length of the string.

**Returns**: `19`