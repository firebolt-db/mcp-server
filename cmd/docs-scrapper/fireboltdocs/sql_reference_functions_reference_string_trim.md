# [](#trim)TRIM

Removes the longest string containing only characters in `<trim_characters>` from the left, right, or both sides of the source string `<expression>`. If no `<trim_characters>` parameter is specified, the longest string containing only whitespace characters (ASCII Decimal 32) is removed. If neither `LEADING`, `TRAILING`, nor `BOTH` are specified, characters are removed from both sides of the specified source string `<expression>`.

## [](#syntax)Syntax

```
TRIM([LEADING | TRAILING | BOTH] [<trim_characters>] FROM <expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `LEADING | TRAILING | BOTH` Optional. Specifies from which part or parts of the `<expression>` to remove the specified `<trim_characters>`. If omitted, this defaults to `BOTH`.

`LEADING` - trims from the beginning of the specified string

`TRAILING` - trims from the end of the specified string.

`BOTH` - trims from the beginning and the end of the specified string. `<trim_characters>` Optional. An expression that returns characters to trim from the right side of the `<expression>` string. If omitted, whitespace (ASCII Decimal 32) is trimmed. `TEXT` `<expression>` An expression that returns the string to be trimmed. `TEXT`

## [](#return-type)Return Type

`TEXT`

## [](#example)Example

The following example trims the characters `x` and `y` from the right side of a string, since the `TRAILING` parameter is specified. Note that the ordering of characters in `<trim_characters>` is irrelevant:

```
SELECT
  TRIM(TRAILING 'xy' FROM 'xyxyThe Acceleration Cupyyxx');
```

**Returns**:

```
'xyxyThe Acceleration Cup'
```

In the following example, no part of the string is specified for `TRIM`, so it defaults to `BOTH`.

```
SELECT
  TRIM('xy' FROM 'xyxyThe Acceleration Cupyyxx');
```

**Returns**:

```
'The Acceleration Cup'
```

The following example omits the `<trim_characters>` parameter but specifies the `TRAILING` parameter, and thus trims whitespace from the right side of a string:

```
SELECT
  TRIM(TRAILING FROM '   The Acceleration Cup     ');
```

**Returns**:

```
'   The Acceleration Cup'
```

The following example omits the `<trim_characters>` parameter and specifies no part of the string, and thus trims whitespace from both sides of a string:

```
SELECT
  TRIM(FROM '   The Acceleration Cup     ');
```

**Returns**:

```
'The Acceleration Cup'
```