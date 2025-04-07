# [](#decode)DECODE

Decode binary data from a SQL expression of type `TEXT`.

## [](#syntax)Syntax

```
DECODE(<expression>, <format>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` A SQL expression of type `TEXT` `TEXT` `<format>` Format to use to decode binary data `HEX`, `ESCAPE`, `BASE64` (case insensitive)

The `HEX` format represents each 4 bits of data as one hexadecimal digit, 0 through f, writing the higher-order digit of each byte first. The `DECODE` function accepts the a-f characters in either upper or lower case. An error is raised when `DECODE` is given invalid hex data — including when given an odd number of characters.

The `ESCAPE` format converts zero bytes and bytes with the high bit set into octal escape sequences (\\nnn) and doubles backslashes. Other byte values are represented literally. The `DECODE` function will raise an error if a backslash is not followed by either a second backslash or three octal digits; it accepts other byte values unchanged.

THE `BASE64` format, per [RFC 2045 Secion 6.8](https://www.rfc-editor.org/rfc/rfc2045#section-6.8), breaks encoded lines at 76 characters using a newline for end of line. The `DECODE` function ignores carriage-return, newline, space, and tab characters. Otherwise, an error is raised when `DECODE`is supplied invalid base64 data — including when trailing padding is incorrect.

## [](#return-type)Return type

`BYTEA`

## [](#example)Example

The following examples decode expressions from the `HEX`, `ESCAPE`, and `BASE64` format:

```
SELECT
	DECODE('31323300343536', 'HEX');
```

**Returns**: `\x31323300343536`

```
SELECT
	DECODE('123\000456', 'ESCAPE');
```

**Returns**: `\x31323300343536`

```
SELECT
	DECODE('MTIzADQ1Ng==', 'BASE64');
```

**Returns**: `\x31323300343536`