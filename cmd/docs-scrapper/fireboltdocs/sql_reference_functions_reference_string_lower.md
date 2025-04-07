# [](#lower)LOWER

Converts the input string to lowercase characters. Note that Firebolt uses the `POSIX` locale, therefore `lower` only converts the ASCII characters “A” through “Z” to lowercase. Non-ASCII characters remain unchanged.

## [](#syntax)Syntax

```
LOWER(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The string to be converted to lowercase characters. `TEXT`

## [](#return-type)Return Type

`TEXT`

## [](#example)Example

The following example converts a game player’s username from uppercase to lowercase characters:

```
SELECT
	LOWER('ESIMPSON') as username
```

**Returns**: `esimpson`

Because Firebolt uses the `POSIX` locale, non-ASCII characters are not lowercased:

```
SELECT LOWER('MÜNCHEN')
```

**Returns**: `mÜnchen`