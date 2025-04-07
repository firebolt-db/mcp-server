# [](#encode)ENCODE

Encode binary data into a SQL expression of type `TEXT`.

## [](#syntax)Syntax

```
ENCODE(<expression>, <format>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` A SQL expression of type `BYTEA` `BYTEA` `<format>` Format to use to encode binary data `HEX`, `ESCAPE`, `BASE64` (case insensitive)

The `HEX` format represents each 4 bits of data as one hexadecimal digit, 0 through f, writing the higher-order digit of each byte first. The `ENCODE` function outputs the a-f hex digits in lower case. Because the smallest unit of data is 8 bits, there are always an even number of characters returned by `ENCODE`.

The `ESCAPE` format converts zero bytes and bytes with the high bit set into octal escape sequences (\\nnn) and doubles backslashes. Other byte values are represented literally.

THE `BASE64` format, per [RFC 2045 Secion 6.8](https://www.rfc-editor.org/rfc/rfc2045#section-6.8), breaks encoded lines at 76 characters using a newline for end of line.

## [](#return-type)Return Type

`TEXT`

## [](#example)Example

The following examples encode binary expressions to the `ESCAPE`, `HEX`, and `BASE64` format:

```
SELECT
	ENCODE('123\000456'::BYTEA, 'ESCAPE');
```

**Returns**: `123\000456`

```
SELECT
	ENCODE('123\000456'::BYTEA, 'HEX');
```

**Returns**: `31323300343536`

```
SELECT
	ENCODE('123\000456'::BYTEA, 'BASE64');
```

**Returns**: `MTIzADQ1Ng==`