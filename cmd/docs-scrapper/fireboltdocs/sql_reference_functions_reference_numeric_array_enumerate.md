# [](#array_enumerate)ARRAY\_ENUMERATE

This function takes an array of arbitrary type as input, and produces an integer array of the same length containing increasing numbers. The returned array starts with value one. Every successive element is incremented by one, it holds that `array[i] = array[i - 1] + 1`.

`NULLs` contained in the parameter array are treated like any other value, and result in a non-null element in the returned array.

If the parameter array is `NULL`, then the function also returns `NULL`.

## [](#syntax)Syntax

```
ARRAY_ENUMERATE(<array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The array to be enumerated. The length of the returned array is the same as the length of the parameter array. Any array type.

## [](#return-type)Return Type

`ARRAY(INT)`

## [](#example)Example

The following example returns an array with values one to four:

The array passed to the function can contain arbitrary types:

NULL values are still reflected in the returned result:

If the array passed to the function is NULL, so is the result: