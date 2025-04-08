# [](#bytea-data-type)BYTEA data type

This topic describes the Firebolt implementation of the `BYTEA` data type.

Not all functions support the `BYTEA` data type currently. For more information, see [BYTEA functions](/sql_reference/functions-reference/bytea/)

- [Overview](#overview)
  
  - [Type conversions](#type-conversions)
  - [Comparison operator](#comparison-operator)
  - [Literal string interpretation](#literal-string-interpretation)
  - [Output format](#output-format)
  - [Importing `BYTEA` from external source](#importing-bytea-from-external-source)

## [](#overview)Overview

The `BYTEA` data type is a variable length binary string data type, commonly used to store binary data, like images, other multimedia files, or raw bytes of information. A binary string is a sequence of bytes - unlike `TEXT`, there is no character set. The `BYTEA` data type is nullable.

### [](#type-conversions)Type conversions

The `BYTEA` type can be cast to and from the `TEXT` data type. A cast from `BYTEA` to `TEXT` will interpret the binary string to a hexadecimal representation with `\x` as a prefix. For example `SELECT 'a'::BYTEA::TEXT` returns `\x61`.

Cast from `TEXT` to `BYTEA` supports two formats, **hex** and **escaped**:

**Hex**  
Using hex format, the `TEXT` type data must start with `\x`. Characters `\n`, `\t`, `\r` and ‘ ‘ are ignored if they are not in sequence of two characters representing one byte. Each character must be in one of the following ranges: `a-f`, `A-F`, `0-9`.  
Characters must be in pairs. For example, `\x aa` is a valid hex format, but `\xa a` is invalid.

**Escape**  
Using escape format, an escaped backslash becomes just a single backslash: `\\` -&gt; `\`. One backslash must be followed by 3 numbers representing octal value (base 8) in range of `000-377`. For example, `a \375`

In addition to casting, the [ENCODE](/sql_reference/functions-reference/bytea/encode.html) and [DECODE](/sql_reference/functions-reference/bytea/decode.html) functions can be used to represent `TEXT` as `BYTEA` and vice versa, but will behave slightly differently. For example, `SELECT ENCODE('1'::BYTEA, 'HEX');` returns `31`, while `SELECT CAST('1'::BYTEA as TEXT);` returns `\x31`, both of type `TEXT`.

### [](#comparison-operator)Comparison operator

The `BYTEA` comparison operator will work as lexicographical comparison but with bytes. Two empty `BYTEA` type expressions are equal. If two `BYTEA` type expressions have equivalent bytes and are of the same length, then they are equal. In a greater than (&gt;) or less than (&lt;) comparison, two `BYTEA` type expressions are compared byte by byte, and the first mismatching byte defines which is greater or less than the other.

**Examples:**

`SELECT '\xa3'::BYTEA > '\xa2'::BYTEA;` returns `TRUE`.

`SELECT '\xa3'::BYTEA = '\xa300'::BYTEA;` returns `FALSE`.

`SELECT '\xa3'::BYTEA < '\xa300'::BYTEA;` returns `TRUE`.

### [](#literal-string-interpretation)Literal string interpretation

Literal strings will be interpreted according to the setting [`standard_conforming_strings`,](/Reference/system-settings.html#enable-parsing-for-literal-strings) which controls whether strings are parsed with or without escaping. Similar to [CAST](/sql_reference/functions-reference/conditional-and-miscellaneous/cast.html) from `TEXT` to `BYTEA`, the two text formats hex and escape are supported.

**Examples:**

```
SET standard_conforming_strings = false;
SELECT '\x3132'::BYTEA; -> '\x313332'
SELECT '\x31   32  '::BYTEA; -> '\x3120202033322020'
SELECT 'a b\230a'::BYTEA; -> '\x61206232333061'

set standard_conforming_strings = true;
SELECT '\x3132'::BYTEA; -> '\x3132'
SELECT '\x31   32  '::BYTEA; -> '\x3132'
SELECT 'a b\230a'::BYTEA; -> '\x6120629861'
```

### [](#output-format)Output format

The output format for `BYTEA` is the hexadecimal representation of the bytes in lower case prefixed by `\x` (Note: in JSON `\` is escaped).

**Example:**

```
SELECT 'a'::BYTEA;
```

**Returns:**

```
{
    "data":
    [
        ["\\x61"]
    ]
}
```

### [](#importing-bytea-from-external-source)Importing `BYTEA` from external source

The input format for importing binary data from an external source depends on the external file format.

**For ORC or PARQUET files:** For a specific field type without annotation (UTF-8 for example): BYTE\_ARRAY (binary), bytes will be imported exactly as they are in the source. All the other types will be imported to the corresponding datatype (for example, BYTE\_ARRAY with UTF-8 annotation will be imported to `TEXT` data type) and then cast to type `BYTEA`.

**For CSV, TSV, or JSON files:** The input data are read exactly as they are in the source, and then cast to data type `BYTEA`.

JSON files must be UTF-8 encoded; however this is not required for CSV and TSV files. In the case that these files are not UTF-8 encoded, field values must not start with `\x` - data starting with the characters `\x` will throw an error on ingest. Any data not starting with the characters `\x` will be copied as bytes to the column of data type `BYTEA`.

**CSV File Example:**

*file*

```csv
'row1'
'a�a'
'\xaabf'
15
'15'
```

**SQL**

```
CREATE EXTERNAL TABLE ex_table
(
    column1 BYTEA
) URL = 's3://...'
  OBJECT_PATTERN = '...'
  TYPE = (CSV);

SELECT * FROM ex_table;
```

**Returns:**

```table
| column1    |
| ---------- |
| \x726f7731 |
| \x61ff61   |
| \xaabf     |
| \x3135     |
| \x3135     |
```