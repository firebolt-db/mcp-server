# [](#json_pointer_extract_text)JSON\_POINTER\_EXTRACT\_TEXT

Accepts a JSON document and pointer expression. If the key exists and the value is a JSON string, `JSON_POINTER_EXTRACT_TEXT` returns it as SQL `TEXT`, removing outer quotes and decoding characters. Otherwise, it returns `NULL`.

## [](#syntax)Syntax

```
JSON_POINTER_EXTRACT_TEXT
(<json>, <json_pointer_expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<json>` The JSON document. `TEXT` `<json_pointer_expression>` A JSON pointer expression to the location of the desired sub-document in the JSON. For more information, see [JSON pointer expression syntax](/sql_reference/functions-reference/JSON/#json-pointer-expression-syntax). `TEXT`

## [](#return-type)Return Type

`TEXT`

- If any input values are `NULL`, the function will return `NULL`.

## [](#examples)Examples

For the JSON document indicated by `<json_common_example>` below, see [JSON common example](/sql_reference/functions-reference/JSON/#json-common-example). The **returned result** is based on the following example.

**Example**

The following code example extracts the value at path `/value/uid` from the JSON document, removes the outermost quotes, and returns the result as SQL `TEXT`, labeled as `res`:

```
SELECT JSON_POINTER_EXTRACT_TEXT('{"value": {"uid": "987654"} }', '/value/uid') AS res
```

**Returns**

The previous code example returns the following:

res (TEXT) 987654

**Example**

The following code example attempts to extract the value at the path `/value/no_such_key` from the JSON document:

```
SELECT JSON_POINTER_EXTRACT_TEXT('{"value": {"uid": "987654"} }', '/value/no_such_key') AS res
```

**Returns**

The previous code example returns the `NULL` with the result labeled as `res`, because the key does not exist:

res (TEXT) NULL

**Example**

The following code example attempts to extract the value at the path `/value/code` from the JSON document:

```
SELECT JSON_POINTER_EXTRACT_TEXT('{"value": {"code": 12} }', '/value/code') AS res
```

**Returns**

The previous code example returns `NULL`, labeled as `res` because the value at the specified path is an integer, not a string:

res (TEXT) NULL

**Example**

The following code example navigates to the third element at index `2` of the array at `/value/keywords` in the JSON document, removes the outermost quotes, and returns it as SQL `TEXT`, labeled as `res`:

```
SELECT JSON_POINTER_EXTRACT_TEXT(<json_common_example>,'/value/keywords/2') AS res
```

**Returns**

The previous code example returns `"analytics"`, which is the third element in the JSON array, which uses zero-based indexing:

res (TEXT) analytics