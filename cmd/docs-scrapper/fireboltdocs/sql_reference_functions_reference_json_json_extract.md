# [](#json_extract)JSON\_EXTRACT

Accepts an expression containing a JSON document, a JSON path expression, and an optional path syntax. If the key specified in the JSON path exists, `JSON_EXTRACT` returns the sub-JSON document at the specified path, and otherwise, returns `NULL`.

## [](#syntax)Syntax

```
JSON_EXTRACT
(<json>, <json_path_expression>, path_syntax => <path_syntax>)
```

### [](#aliases)Aliases

```
JSON_POINTER_EXTRACT
(<json>, <json_path_expression>) ->
JSON_EXTRACT(<json>, <json_path_expression>, path_syntax => 'JSONPointer')
```

## [](#parameters)Parameters

Parameter Description Supported input types `<json>` The JSON document. `TEXT` `<json_path_expression>` A JSON path that specifies the location of the desired sub-document within the JSON document. `TEXT` `<path_syntax>` The expected syntax of the `<json_path_expression>`, currently supports only the ‘JSONPointer’ syntax. For more information, see [JSON pointer expression syntax](/sql_reference/functions-reference/JSON/#json-pointer-expression-syntax). `TEXT`

## [](#return-type)Return Type

`TEXT`

- If any input values are `NULL`, the function will return `NULL`.

## [](#examples)Examples

For the JSON document indicated by `<json_common_example>` below, see [JSON common example](/sql_reference/functions-reference/JSON/#json-common-example). The **returned result** is based on the following example.

**Example**

The following code example extracts the value at the path `/value/dyid` from the JSON document represented by `<json_common_example>` using the `JSONPointer` syntax:

```
SELECT JSON_EXTRACT(<json_common_example>, '/value/dyid', 'JSONPointer')
```

**Returns**

The previous example returns “987” because the key `dyid` is associated with the value `987` within the `value` object, and the function retrieves and returns this value as a string.

**Example**

The following code example attempts to extract a value from the path `/value/no_such_key` in the JSON document represented by `<json_common_example>`:

```
SELECT JSON_EXTRACT(<json_common_example>, '/value/no_such_key', 'JSONPointer')
```

**Returns**

The previous code example returns `NULL` because the key `no_such_key` does not exist.

**Example**

The following code example extracts the value at the path `/value/uid` from the JSON document represented by `<json_common_example>`:

```
SELECT JSON_POINTER_EXTRACT(<json_common_example>, '/value/uid')
```

**Returns**

The previous code example returns `'"987654"'` because the value associated with the `uid` key in the JSON document is the string `987654`. The function retrieves this value with double quotes, indicating it’s a JSON string.

**Example**

The following code example extracts the value at the path `/value/keywords` from the JSON document represented by `<json_common_example>`:

```
SELECT JSON_POINTER_EXTRACT(<json_common_example>,'/value/keywords')
```

**Returns**

The previous code example returns the array `["insanely","fast","analytics"]` that is associated with the `keywords` key.

**Example**

The following code example extracts the third element at index `2` from the array located at the path `/value/keywords` in the JSON document represented by `<json_common_example>`:

```
SELECT JSON_POINTER_EXTRACT(<json_common_example>,'/value/keywords/2')
```

**Returns**

The previous code example returns `'"analytics"'` because it accesses the third element in the JSON array, which uses zero-based indexing.