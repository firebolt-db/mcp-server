# [](#json_value_array)JSON\_VALUE\_ARRAY

Takes a JSON document and extracts a JSON array of scalar values to SQL `ARRAY(TEXT)` value. For JSON strings, removes the outermost quotes and unescapes the values. Other JSON scalars are not changed. Returns a SQL `NULL` if a non-array is given, or non-scalar value is given (inside the array).

This function pairs with the [JSON\_EXTRACT](/sql_reference/functions-reference/JSON/json-extract.html) function, which doesn’t convert the JSON values to SQL values.

## [](#syntax)Syntax

```
JSON_VALUE_ARRAY(<json>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<json>` The JSON document. `TEXT`

## [](#return-type)Return Type

`ARRAY(TEXT)`

- If any of the input is `NULL` the output is `NULL` (propagates nulls).

## [](#example)Example

For the JSON document indicated by `<json_common_example>` below, see [JSON common example](/sql_reference/functions-reference/JSON/#json-common-example). The **returned result** is based on this example.

```
SELECT JSON_VALUE_ARRAY(JSON_POINTER_EXTRACT(<json_common_example>, '/value/uid')), JSON_POINTER_EXTRACT(<json_common_example>, '/value/uid')
```

**Returns**: `NULL, '"987654"'`

```
SELECT JSON_VALUE_ARRAY(JSON_POINTER_EXTRACT(<json_common_example>,'/value/keywords'))
```

**Returns**: `{'insanely','fast','analytics'}`

```
SELECT JSON_VALUE_ARRAY(NULL)
```

**Returns**: `NULL`