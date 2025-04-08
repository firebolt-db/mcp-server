# [](#array_any_match)ARRAY\_ANY\_MATCH

- Returns `TRUE` if any element in the array is `TRUE`.
- Returns `FALSE` if all elements in the array are `FALSE` or if the array is empty.
- Returns `NULL` if any element is `NULL` and no element is `TRUE`.

When an optional lambda function is provided, `ARRAY_ANY_MATCH` applies the function to each element and then evaluates the resulting arrayt.

**Alias:** `ANY_MATCH`

## [](#syntax)Syntax

```
{ ANY_MATCH | ARRAY_ANY_MATCH }([<expression> -> <condition>], <array> [, <array2>, ...])
```

## [](#parameters)Parameters

Parameter Description Supported input types     `<expression>` A lambda function applied to each element of the input arrays, returning a `BOOLEAN`. If no lambda function is provided, the function can only evaluate a single `BOOLEAN` array. For more information, see [Manipulating arrays with Lambda functions](/Guides/loading-data/working-with-semi-structured-data/working-with-arrays.html#manipulating-arrays-with-lambda-functions). Same as the element data types of the input arrays.   `<condition>` A `BOOLEAN` expression that evaluates each array value using a comparison operator. See [Comparison operators](/sql_reference/operators.html#comparison).   `<array>` The array to evaluate. `ARRAY`

## [](#return-types)Return Types

The `ARRAY_ANY_MATCH` function returns a result of type `BOOLEAN`.

## [](#examples)Examples

**Example** The following code example checks if an array contains the value `esimpson` as the result `is_he_playing`:

```
SELECT ARRAY_ANY_MATCH(x -> x = 'esimpson', [ 'kennethpark', 'sabrina21', 'steven70']) AS is_he_playing;
```

**Returns** The previous code returns `FALSE` because the array does not contain the specified value:

is\_he\_playing (BOOLEAN) f

**Example** The following code example checks if each element in the first array is divisible by the corresponding element in the second array in a result labeled `divisible`:

```
SELECT ARRAY_ANY_MATCH(x, y -> (x % y) = 0, [ 10, 20, 30, 45 ], [ 12, 3, 42, 15]) AS divisible;
```

**Returns** The previous code returns `TRUE` because each element in the first array is divisible by its corresponding element in the second array:

divisible (BOOLEAN) t

**Example** The following code example evaluates multiple arrays using `ARRAY_ANY_MATCH`:

```
SELECT ARRAY_ANY_MATCH([])          as empty,
    ARRAY_ANY_MATCH([true])         as single_true,
    ARRAY_ANY_MATCH([false])        as single_false,
    ARRAY_ANY_MATCH([NULL])         as single_null ,
    ARRAY_ANY_MATCH([false, NULL])  as false_and_null;
```

**Returns** The previous code returns `FALSE` for the empty array, `TRUE` for the `[TRUE]` array, `FALSE` for the `[FALSE]` array, `NULL` for the `[NULL]` array, and `FALSE` for the `[FALSE, NULL]` array:

empty (BOOLEAN) single\_true (BOOLEAN) single\_false (BOOLEAN) single\_null (BOOLEAN) false\_and\_null (BOOLEAN) f t f NULL NULL