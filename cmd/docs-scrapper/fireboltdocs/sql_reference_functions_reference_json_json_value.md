# [](#json_value)JSON\_VALUE

Takes a JSON document and extracts a JSON scalar value to SQL `TEXT` value. For JSON strings, removes the outermost quotes and unescapes the values. Other JSON scalars are not changed. Returns a SQL `NULL` if a non-scalar value is given.

This function pairs with the [JSON\_EXTRACT](/sql_reference/functions-reference/JSON/json-extract.html) function, which doesn’t convert the JSON values to SQL values.

## [](#syntax)Syntax

```
JSON_VALUE(<json>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<json>` The JSON document. `TEXT`

## [](#return-type)Return Type

`TEXT`

- If any of the input is `NULL` the output is `NULL` (propagates nulls).

## [](#example)Example

For the JSON document indicated by `<json_common_example>` below, see [JSON common example](/sql_reference/functions-reference/JSON/#json-common-example). The **returned result** is based on this example.

```
SELECT JSON_VALUE(JSON_POINTER_EXTRACT(<json_common_example>, '/value/uid')), JSON_POINTER_EXTRACT(<json_common_example>, '/value/uid')
```

**Returns**

`'987654', '"987654"'`

**Example**

```
SELECT JSON_VALUE(JSON_POINTER_EXTRACT(<json_common_example>, '/key'))::INT
```

**Returns**

`123`

**Example**

```
SELECT JSON_VALUE(JSON_POINTER_EXTRACT(<json_common_example>,'/value/keywords'))
```

**Returns**

`NULL`

**Example**

```
SELECT JSON_VALUE(NULL)
```

**Returns**

`NULL`