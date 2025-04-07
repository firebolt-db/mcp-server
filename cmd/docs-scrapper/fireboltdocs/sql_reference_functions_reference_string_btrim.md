# [](#btrim)BTRIM

Removes the longest string containing only characters in `<trim_characters>` from both sides of the source string `<expression>`. If no `<trim_characters>` parameter is specified, the longest string containing only whitespace characters (ASCII Decimal 32) is removed from both sides of the specified source string `<expression>`.

## [](#syntax)Syntax

```
BTRIM(<expression>[, <trim_characters>])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` An expression that returns the string to be trimmed. `TEXT` `<trim_characters>` Optional. An expression that returns characters to trim from both sides of the `<expression>` string. If omitted, whitespace (ASCII Decimal 32) is trimmed. `TEXT`

## [](#return-type)Return Type

`TEXT`

## [](#examples)Examples

The following example trims the character `x` from both sides of a string:

```
SELECT BTRIM('xxThe Acceleration Cupxxx', 'x') as result
```

result (TEXT) ‘The Acceleration Cup’

The following example trims the characters `x` and `y` from both sides of a string. Note that the ordering of characters in `<trim_characters>` is irrelevant:

```
SELECT BTRIM('xyxyThe Acceleration Cupyyxx', 'xy') as result;
```

result (TEXT) ‘The Acceleration Cup’

The following example omits the parameter, and thus trims whitespace from both sides of a string:

```
SELECT BTRIM('   The Acceleration Cup     ') as result;
```

result (TEXT) ‘The Acceleration Cup’