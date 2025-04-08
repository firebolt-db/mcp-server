# [](#array_all_match)ARRAY\_ALL\_MATCH

- Returns `TRUE` if all elements in the array are `TRUE` or if the array is empty.
- Returns `FALSE` if any element in the array is `FALSE`.
- Returns `NULL` if any element is `NULL` and no element is `FALSE`.

When an optional lambda function is provided, `ARRAY_ALL_MATCH` applies the function to each element and then evaluates the resulting array.

**Alias:** `ALL_MATCH`

## [](#syntax)Syntax

```
{ ALL_MATCH | ARRAY_ALL_MATCH }([<expression> -> <condition>], <array> [, <array2>, ...])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` A lambda function applied to each element of the input arrays, returning a `BOOLEAN`. If no lambda function is provided, the function can only evaluate a single `BOOLEAN` array. For more information, see [Manipulating arrays with Lambda functions](/Guides/loading-data/working-with-semi-structured-data/working-with-arrays.html#manipulating-arrays-with-lambda-functions). Same as the element data types of the input arrays. `<condition>` A `BOOLEAN` expression that evaluates each array value using a comparison operator. See [Comparison operators](/sql_reference/operators.html#comparison). `<array>` The array to evaluate. `ARRAY`

## [](#return-type)Return Type

The `ARRAY_ALL_MATCH` function returns a result of type `BOOLEAN`.

## [](#examples)Examples

Check if all player nicknames end with `'2024'`:

```
SELECT ARRAY_ALL_MATCH(name -> name like '%2024', [ 'kennethpark2024', 'sabrina2024', 'steven2024']) AS result;
```

result (BOOLEAN) t

Check if all elements in the first array can be divided by the elements in the second array:

```
SELECT ARRAY_ALL_MATCH(x, y -> (x % y) = 0, [ 10, 20, 30, 45 ], [ 5, 10, 2, 15]) AS divisable;
```

divisable (BOOLEAN) t

Check if all elements in an input array are `true`:

```
SELECT ARRAY_ALL_MATCH([])          as empty,
    ARRAY_ALL_MATCH([true])         as single_true,
    ARRAY_ALL_MATCH([false])        as single_false,
    ARRAY_ALL_MATCH([NULL])         as single_null ,
    ARRAY_ALL_MATCH([false, NULL])  as false_and_null;
```

empty (BOOLEAN) single\_true (BOOLEAN) single\_false (BOOLEAN) single\_null (BOOLEAN) false\_and\_null (BOOLEAN) t t f NULL f