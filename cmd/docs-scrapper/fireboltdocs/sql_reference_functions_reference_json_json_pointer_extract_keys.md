# [](#json_pointer_extract_keys)JSON\_POINTER\_EXTRACT\_KEYS

Accepts a JSON document and pointer expression. If the key exists and holds a JSON object (map), `JSON_POINTER_EXTRACT_KEYS` returns all the keys in that object as SQL `ARRAY(TEXT)`, removing outer quotes and decoding characters. Otherwise, it returns `NULL`.

## [](#syntax)Syntax

```
JSON_POINTER_EXTRACT_KEYS
(<json>, <json_pointer_expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<json>` The JSON document. `TEXT` `<json_pointer_expression>` A JSON pointer expression to the location of the desired sub-document in the JSON. For more information, see [JSON pointer expression syntax](/sql_reference/functions-reference/JSON/#json-pointer-expression-syntax). `TEXT`

## [](#return-type)Return Type

`ARRAY(TEXT)`

- If any input values are `NULL`, the function will return `NULL`.

## [](#examples)Examples

For the JSON document indicated by `<json_common_example>` below, see [JSON common example](/sql_reference/functions-reference/JSON/#json-common-example).

**Example**

The following code example extracts all the keys at path `/value` from the JSON document, removes the outermost quotes, and returns the result as SQL `ARRAY(TEXT)`, labeled as `res`:

```
SELECT JSON_POINTER_EXTRACT_KEYS(<json_common_example>, '/value') AS res
```

**Returns**

The previous code example returns the following:

res (ARRAY(TEXT)) {‘dyid’, ‘uid’, ‘keywords’, ‘tagIdToHits’, ‘events’}

**Example**

The following code example attempts to extract the keys at the path `/value/no_such_key` from the JSON document:

```
SELECT JSON_POINTER_EXTRACT_KEYS(<json_common_example>, '/value/no_such_key') AS res
```

**Returns**

The previous code example returns the `NULL` with the result labeled as `res`, because the key does not exist:

res (ARRAY(TEXT)) NULL

**Example**

The following code example attempts to extract the value at the path `/value/keywords` from the JSON document:

```
SELECT JSON_POINTER_EXTRACT_KEYS(<json_common_example>,'/value/keywords') AS res
```

**Returns**

The previous code example returns `NULL`, labeled as `res` because the value at the specified path is an array, not an object:

res (ARRAY(TEXT)) NULL