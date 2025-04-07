# [](#strpos)STRPOS

Returns the position of the substring found in the string, starting from 1. The returned value is for the first matching value, and not for any subsequent valid matches. In case the substring does not exist, functions will return 0.

## [](#syntax)Syntax

```
STRPOS(<string>, <substring>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<string>` The string that will be searched. `TEXT` `<substring>` The substring to search for. `TEXT`

## [](#return-type)Return Type

`INTEGER`

## [](#example)Example

```
SELECT
	STRPOS('hello world','hello') AS res;
```

**Returns**: `1`

```
SELECT
	STRPOS('hello world','world') AS res;
```

**Returns**: `7`

```
SELECT
	STRPOS('hello world','work') AS res;
```

**Returns**: `0`