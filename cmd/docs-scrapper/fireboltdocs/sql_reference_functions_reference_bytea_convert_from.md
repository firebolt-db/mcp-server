# [](#convert_from)CONVERT\_FROM

Converts a binary string from a specified encoding to the `TEXT` data type, using the database’s encoding.

## [](#syntax)Syntax

```
CONVERT_FROM
(<bytes>, <src_encoding>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<bytes>` A sequence of bytes representing text in `<src_encoding>`. `BYTEA` `<src_encoding>` The encoding from which the src `<bytes>` are derived, which must be constant. The supported source encoding aligns with those supported in the [ICU library](https://icu.unicode.org/). `TEXT`

## [](#return-type)Return Type

`CONVERT_FROM` returns a value of type `TEXT`.

## [](#errors)Errors

If `<src_encoding>` is invalid, an error is thrown. If `<bytes>` are malformed according to `<src_encoding>`, the behavior is undefined and may result in replacement characters or errors, depending on the encoding. For example, in `UTF-8` encoding, malformed bytes are replaced with the � character.

## [](#examples)Examples

**Example**

The following code example uses `CONVERT_FROM` to convert the binary input `\x1212003100` of type `BYTEA` into a `TEXT` string using the `UTF-16` encoding. All results are displayed in `UTF-8` encoding:

```
SELECT CONVERT_FROM('\x1212003100'::BYTEA, 'utf16');
```

**Returns**

The result, `ሒ1�` consists of:

- `ሒ`, the `UTF-16` representation of `1212`.
- `1`, the `UTF-16` representation of `0031`.
- The replacement character `�`. Because `00` is a single byte, it does not form a valid `UTF-16` sequence.

**Example**

The following code example uses `CONVERT_FROM` to convert the binary input `\x31a031ffffffff` of type `BYTEA` into a `TEXT` string using the `windows-1252` encoding:

```
SELECT CONVERT_FROM('\x31a031ffffffff'::BYTEA, 'windows-1252');
```

**Returns**

The result, `1 1ÿÿÿÿ`, consists of:

- `1`, the windows-1252 representation of `31`.
- ` `, the windows-1252 representation of the [non-breaking space character](https://en.wikipedia.org/wiki/Non-breaking_space) a0.
- `ÿ`, the windows-1252 representation of `ff`.

**Example**

The following code example creates a table, which represents an external table, with a `BYTEA` column that represents text data in a custom encoding that is specified during data loading:

```
CREATE TABLE external_bytes_data (a bytea);
INSERT INTO external_bytes_data VALUES ('\x1212003100');
```

**Example**

The following code example converts the `BYTEA` column `col_bytea` into `TEXT` using the specified encoding `UTF-16`, and inserts it into `text_encoded_data` target table:

```
CREATE TABLE text_encoded_data (a text);
INSERT INTO text_encoded_data SELECT CONVERT_FROM(a, 'UTF-16') FROM external_bytes_data;
SELECT * FROM text_encoded_data;
```

In the previous code example, you can replace the `UTF-16` input encoding with another encoding supported by the [ICU library](https://icu.unicode.org/).