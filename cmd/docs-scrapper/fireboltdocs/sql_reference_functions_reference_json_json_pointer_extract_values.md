# [](#json_pointer_extract_values)JSON\_POINTER\_EXTRACT\_VALUES

Accepts a JSON document and pointer expression. If the key exists and holds a JSON object (map), `JSON_POINTER_EXTRACT_VALUES` returns all the values in that object as SQL `ARRAY(TEXT)`, while the values remain raw as they appeared in the original JSON document. Otherwise, it returns `NULL`.

## [](#syntax)Syntax

```
JSON_POINTER_EXTRACT_VALUES
(<json>, <json_pointer_expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<json>` The JSON document. `TEXT` `<json_pointer_expression>` A JSON pointer expression to the location of the desired sub-document in the JSON. For more information, see [JSON pointer expression syntax](/sql_reference/functions-reference/JSON/#json-pointer-expression-syntax). `TEXT`

## [](#return-type)Return Type

`ARRAY(TEXT)`

- If any input values are `NULL`, the function will return `NULL`.

## [](#examples)Examples

For the JSON document indicated by `<json_common_example>` below, see [JSON common example](/sql_reference/functions-reference/JSON/#json-common-example). The **returned result** is based on the following example.

**Example**

The following code example extracts all the values at path `/value/events/0/EventProperties` from the JSON document, and returns the raw values as SQL `ARRAY(TEXT)`, labeled as `res`:

```
SELECT JSON_POINTER_EXTRACT_VALUES(<json_common_example>, '/value/events/0/EventProperties') AS res
```

**Returns**

The previous code example returns the following:

res (ARRAY(TEXT)) {‘“John Doe”’, ‘true’}

**Example**

The following code example attempts to extract the keys at the path `/value/no_such_key` from the JSON document:

```
SELECT JSON_POINTER_EXTRACT_VALUES(<json_common_example>, '/value/no_such_key') AS res
```

**Returns**

The previous code example returns the `NULL` with the result labeled as `res`, because the key does not exist:

res (ARRAY(TEXT)) NULL

**Example**

The following code example attempts to extract the value at the path `/value/keywords` from the JSON document:

```
SELECT JSON_POINTER_EXTRACT_VALUES(<json_common_example>,'/value/keywords') AS res
```

**Returns**

The previous code example returns `NULL`, labeled as `res` because the value at the specified path is an array, not an object:

res (ARRAY(TEXT)) NULL