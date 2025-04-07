# [](#upper)UPPER

Converts the input string to uppercase characters. Note that Firebolt uses the `POSIX` locale, therefore `upper` only converts the ASCII characters “a” through “z” to uppercase. Non-ASCII characters remain unchanged.

## [](#syntax)Syntax

```
UPPER(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The string to be converted to uppercase characters. `TEXT`

## [](#return-type)Return Type

`TEXT`

## [](#example)Example

The following example converts a game player’s username from lowercase to uppercase characters:

```
SELECT
	UPPER('esimpson') as username
```

**Returns**: `ESIMPSON`

Because Firebolt uses the `POSIX` locale, non-ASCII characters are not uppercased:

```
SELECT UPPER('München')
```

**Returns**: `MüNCHEN`