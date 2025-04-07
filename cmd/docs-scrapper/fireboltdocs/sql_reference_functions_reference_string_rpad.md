# [](#rpad)RPAD

Adds a specified pad string to the end of the string repetitively up until the length of the resulting string is equivalent to an indicated length.

The similar function to pad the start of a string is [LPAD](/sql_reference/functions-reference/string/lpad.html).

## [](#syntax)Syntax

```
RPAD(<expression>, <value>[, <pad>])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The original string. If the length of the original string is larger than the length parameter, this function removes the overflowing characters from the string. `TEXT` `<value>` The integer length that the string will be after it has been left-padded. A negative number returns an empty string. `INTEGER` `<pad>` The string to add to the end of the primary string `<expression`. If left blank, `<pad>` defaults to whitespace characters.  

## [](#example)Example

The following statement adds the string “ABC” to the end of the string “Firebolt” repetitively until the resulting string is equivalent to 20 characters in length.

```
SELECT
	RPAD('Firebolt', 20, 'ABC');
```

**Returns**: `FireboltABCABCABCABC`