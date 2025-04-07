# [](#position)POSITION

Returns the position of the substring found in the string, starting from 1. The returned value is for the first matching value, and not for any subsequent valid matches. In case the substring does not exist, functions will return 0.

## [](#syntax)Syntax

```
POSITION(<substring> IN <string>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<substring>` The substring to search for. `TEXT` `<string>` The string that will be searched. `TEXT`

## [](#return-type)Return Type

`INTEGER`

## [](#example)Example

```
SELECT
	POSITION('hello' IN 'hello world') AS res;
```

**Returns**: `1`

```
SELECT
	POSITION('world' IN 'hello world') AS res;
```

**Returns**: `7`

```
SELECT
	POSITION('work' IN 'hello world') AS res;
```

**Returns**: `0`