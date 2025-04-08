# [](#octet_length)OCTET\_LENGTH

Calculates the length of the input string in bytes.

## [](#syntax)Syntax

```
OCTET_LENGTH(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The string or binary data for which to return the length. `TEXT`, `BYTEA`

## [](#return-type)Return Type

`INTEGER`

## [](#example)Example

Use the `OCTET_LENGTH` to find the length of any string in bytes, such as:

```
SELECT LENGTH('🔥')
```

**Returns**: `4`

Because the UTF8 encoding of ‘🔥’ has the byte sequence `0xF0 0x9F 0x94 0xA5`.