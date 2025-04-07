# [](#ltrim)LTRIM

Removes the longest string containing only characters in `<trim_characters>` from the left side of the source string `<expression>`. If no `<trim_characters>` parameter is specified, the longest string containing only whitespace characters (ASCII Decimal 32) is removed from the left side of the specified source string `<expression>`.

## [](#syntax)Syntax

```
LTRIM(<expression>[, <trim_characters>])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` An expression that returns the string to be trimmed. `TEXT` `<trim_characters>` Optional. An expression that returns characters to trim from the left side of the `<expression>` string. If omitted, whitespace (ASCII Decimal 32) is trimmed. `TEXT`

## [](#return-type)Return Type

`TEXT`

## [](#examples)Examples

The following example trims the character `x` from the left side of a string:

```
SELECT
  LTRIM('xxThe Acceleration Cupxxx', 'x') 
```

**Returns**:

```
'The Acceleration Cupxxx'
```

The following example trims the characters `x` and `y` from the left side of a string. Note that the ordering of characters in `<trim_characters>` is irrelevant:

```
SELECT
  LTRIM('xyxyThe Acceleration Cupyyxx', 'xy');
```

**Returns**:

```
'The Acceleration Cupyyxx'
```

The following example omits the `<trim_characters>` parameter, and thus trims whitespace from the left side of a string:

```
SELECT
  LTRIM('   The Acceleration Cup     ');
```

**Returns**:

```
'The Acceleration Cup     '
```