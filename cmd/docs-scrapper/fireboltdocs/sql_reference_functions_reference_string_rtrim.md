# [](#rtrim)RTRIM

Removes the longest string containing only characters in `<trim_characters>` from the right side of the source string `<expression>`. If no `<trim_characters>` parameter is specified, the longest string containing only whitespace characters (ASCII Decimal 32) is removed from the right side of the specified source string `<expression>`.

## [](#syntax)Syntax

```
RTRIM(<expression>[, <trim_characters>])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` An expression that returns the string to be trimmed. `TEXT` `<trim_characters>` Optional. An expression that returns characters to trim from the right side of the `<expression>` string. If omitted, whitespace (ASCII Decimal 32) is trimmed. `TEXT`

## [](#return-types)Return Types

`TEXT`

## [](#examples)Examples

The following example trims the character `x` from the right side of a string:

```
SELECT
  RTRIM('xxThe Acceleration Cupxxx', 'x') 
```

**Returns**:

```
'xxThe Acceleration Cup'
```

The following example trims the characters `x` and `y` from the right side of a string. Note that the ordering of characters in `<trim_characters>` is irrelevant:

```
SELECT
  RTRIM('xyxyThe Acceleration Cupyyxx', 'xy');
```

**Returns**:

```
'xyxyThe Acceleration Cup'
```

The following example omits the `<trim_characters>` parameter, and thus trims whitespace from the right side of a string:

```
SELECT
  RTRIM('   The Acceleration Cup     ');
```

**Returns**:

```
'   The Acceleration Cup'
```