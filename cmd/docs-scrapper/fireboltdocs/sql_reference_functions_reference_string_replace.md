# [](#replace)REPLACE

Replaces all occurrences of the `<pattern>` substring within the `<expression>` with the `<replacement>` substring.

## [](#syntax)Syntax

```
REPLACE (<expression>, <pattern, <replacement>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The original string that will be searched for instances of the `<pattern>`. `TEXT` `<pattern>` The substring to be searched and replaced in the string. `TEXT` `<replacement>` The substring to replace the original substring defined by `<pattern>`. To remove the `<pattern>` substring with no replacement, you can use a empty string `''` as the replacement value. `TEXT`

## [](#example)Example

In the example below, “two” in “Level two” is replaced with “three”.

```
SELECT
	REPLACE('Level two','two','three') AS level; 
```

**Returns**: `Level three`

In this example below, “eight” is replaced by an empty string.

```
SELECT
	REPLACE('Level eight',' eight','') AS level;
```

**Returns**: `Level`

In this following example, the substring “five” is not found in the original string, so the string is returned unchanged.

```
SELECT
	REPLACE('Level four','five','six') AS level;
```

**Returns**: `Level four`