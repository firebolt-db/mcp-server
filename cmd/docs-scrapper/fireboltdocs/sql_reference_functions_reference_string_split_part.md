# [](#split_part)SPLIT\_PART

Splits `<string>` at occurrences of `<delimiter>` and returns the `<index>`‘th field, with 1 being the first index. If `<index>` is negative, returns the `abs(<index>)`‘th-from-last field. If `<delimiter>` is empty, `<string>` is returned at `<index>` 1. If `abs(<index>)` is larger than the number of fields, returns an empty string. The function raises an error for `<index>` 0.

## [](#syntax)Syntax

```
SPLIT_PART(<string>, <delimiter>, <index>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<string>` The string to be split. `TEXT` `<delimiter>` The character used to split the string. `TEXT` `<index>` The index position of the substring to return within the split parts. `INTEGER`

## [](#return-type)Return Type

`TEXT`

## [](#examples)Examples

**Example**

The following code example uses the delimiter `#` to split the string `'hello#world'`, and returns the first segment:

```
SELECT
	SPLIT_PART('hello#world','#',1) AS res;
```

**Returns**

`'hello'`

**Example**

The following code example uses the delimiter `|` to split the string `'this|is|my|test'`, and returns the second to last segment:

```
SELECT
	SPLIT_PART('this|is|my|test', '|', -2) AS res;
```

**Returns**

`'my'`

**Example**

The following code example uses an empty delimiter `''` to split the string `'hello world'`, and returns the first segment:

```
SELECT
	SPLIT_PART('hello world', '', 1) AS res;
```

**Returns**

`hello world`

In the previous code example, `SPLIT_PART` sees the empty delimiter and interprets the entire string as the first and only part of the string and returns ‘hello world’.

**Example**

The following code example uses an empty delimiter `''` to split the string `'hello world'`, and returns the first segment:

```
SELECT
	SPLIT_PART('hello world', '', 7) AS res;
```

**Returns**

`''`

In the previous code example, `SPLIT_PART` sees the empty delimiter and interprets the entire string as the first and only part of the string. Because there is no seventh part to return, `SPLIT_PART` returns an empty string.