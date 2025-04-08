# [](#data-types)Data types

This topic lists the data types available in Firebolt.

Category Data Type Aliases Description Numeric [INTEGER](https://docs.firebolt.io/sql_reference/data-types.html#integer) INT, INT4 A four-byte signed integer. [BIGINT](https://docs.firebolt.io/sql_reference/data-types.html#bigint) LONG, INT8 An eight-byte signed integer. [NUMERIC](https://docs.firebolt.io/sql_reference/data-types.html#numeric-1) DECIMAL An exact numeral defined by a fixed precision and scale, with a default of \`38\` for precision and \`9\` for scale. [REAL](https://docs.firebolt.io/sql_reference/data-types.html#real) FLOAT4 A four-bye floating point number with six decimal digits of precision. [DOUBLE PRECISION](https://docs.firebolt.io/sql_reference/data-types.html#double-precision) DOUBLE, FLOAT, FLOAT8, FLOAT(p) An eight-byte floating point number with fifteen decimal digits of precision. Boolean [BOOLEAN](https://docs.firebolt.io/sql_reference/data-types.html#boolean-1) BOOL A logical boolean value of true or false. Composite [ARRAY](#array) An array that holds multiple values of the same data type. Date &amp; Timestamp [DATE](#date) A calendar date including the year, month, and day. [TIMESTAMP](#timestamp) A calendar date and time in Coordinated Universal Time (UTC), including the year, month, day, hour, minute, second, and microsecond. [TIMESTAMPTZ](#timestamptz) A calendar date and time in the local timezone, including year, month, day, hour, minute, second, and microsecond. String [TEXT](#text) A character string of variable length. Binary [BYTEA](#bytea) Binary data of variable length. Spatial [GEOGRAPHY](#geography) Geographic data, including points, LineStrings, Polygons and collections thereof.

*For more information about converting a value with a given data type to another, see [Type Conversion](#type-conversion).*

## [](#numeric)Numeric

### [](#integer)INTEGER

A whole number ranging from -2,147,483,648 to 2,147,483,647. `INTEGER` data types require 4 bytes of storage. Aliases: `INT`, `INT4`.

### [](#bigint)BIGINT

A whole number ranging from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807. `BIGINT` data types require 8 bytes of storage. Aliases: `LONG`, `INT8`.

### [](#numeric-1)NUMERIC

A fixed-point numeric data type defined by its precision (total number of digits) and scale (number of digits to the right of the decimal point). For more information, see [NUMERIC data type](/sql_reference/numeric-data-type). Aliases: `DECIMAL`.

### [](#real)REAL

A floating-point number that has six decimal-digit precision. `REAL` data types require 4 bytes of storage. Aliases: `FLOAT4`.

### [](#double-precision)DOUBLE PRECISION

A floating-point number that has 15 decimal-digit precision. `DOUBLE` data types require 8 bytes. Aliases: `DOUBLE`, `FLOAT`, `FLOAT8`, `FLOAT(p)` where 25 &lt;= p &lt;= 53.

## [](#boolean)Boolean

### [](#boolean-1)BOOLEAN

Represents boolean value of `TRUE` or `FALSE`. Aliases: `BOOL`

## [](#composite)Composite

### [](#array)ARRAY

Represents an array of values. All elements of the array must have the same data type. Elements of the array can be of any supported data type including nested arrays (array with arrays).

Array columns must be defined with the data type of the array elements, and optionally whether or not those elements are nullable. The following syntax options are supported:

- `ARRAY(<data-type> [NULL | NOT NULL])`
- `<data-type> ARRAY`
- `<data-type>[]`

For example, the following three queries will create tables with the same nullable `demo_array` column of `TEXT` elements:

```
  CREATE TABLE demo1 (
  demo_array ARRAY(TEXT NULL) 
  );
  
  CREATE TABLE demo2 (
  demo_array TEXT[]
  );

  CREATE TABLE demo3 (
  demo_array TEXT ARRAY 
  );
```

You can also specify that an array be NOT NULL, but you must then use the `ARRAY(<data-type> NOT NULL)` syntax.

You can access a specific array element with an array subscript expression: `array_value[index]`. The supplied index must be of type `INT` or `BIGINT`. An array of n elements starts with `array_value[1]` and ends with `array_value[n]`. Array subscript expressions:

- raise an error if the subscript expression evaluates to a negative number or 0,
- return NULL if the array is NULL, or if the subscript expression evaluates to NULL or an index larger than the size of the array,
- return the specific element of the array for subscript expressions evaluating to a number in the range \[1, array\_size].

#### [](#example)Example

The following `CREATE TABLE` statement shows arrays of different element types and different nullabilities.

```
CREATE TABLE demo (
  a_t ARRAY(TEXT NULL) NULL,
  a_i ARRAY(INTEGER NULL) NOT NULL,
  a_d ARRAY(DATE NOT NULL) NULL,
  a_f ARRAY(REAL NOT NULL) NOT NULL,
  a_a ARRAY(ARRAY(INTEGER NULL) NULL) NULL
);
```

And the following `INSERT INTO` statement demonstrates examples of values for these arrays:

```
INSERT INTO demo VALUES
  (
    ['Hello', NULL, 'world'],
    [1, 42, NULL],
    [DATE '2000-01-01'],
    [3.14, 2.71, 9.8],
    [ [1, 2], [NULL], NULL]
  ),
  (
    NULL,
    [],
    NULL,
    [],
    NULL
  )
```

## [](#date-and-timestamp)Date and timestamp

Type Size Min Max Resolution `DATE` 4 bytes `0001-01-01` `9999-12-31` 1 day `TIMESTAMP` 8 bytes `0001-01-01 00:00:00.000000` `9999-12-31 23:59:59.999999` 1 microsecond `TIMESTAMPTZ` 8 bytes `0001-01-01 00:00:00.000000 UTC` `9999-12-31 23:59:59.999999 UTC` 1 microsecond

Dates are counted according to the [proleptic Gregorian calendar](https://en.wikipedia.org/wiki/Proleptic_Gregorian_calendar). Each year consists of 365 days, with leap days added to February in leap years.

### [](#date)DATE

A year, month, and day calendar date independent of a time zone. For more information, see [DATE data type](/sql_reference/date-data-type.html).

### [](#timestamp)TIMESTAMP

A year, month, day, hour, minute, second, and microsecond timestamp independent of a time zone. For more information, see [TIMESTAMP data type](/sql_reference/timestampntz-data-type.html).

### [](#timestamptz)TIMESTAMPTZ

A year, month, day, hour, minute, second, and microsecond timestamp associated with a time zone. For more information, see [TIMESTAMPTZ data type](/sql_reference/timestamptz-data-type.html).

## [](#string)String

### [](#text)TEXT

The `TEXT` type can be used to store character strings of any length using the UTF-8 encoding standard. Only the ASCII letters “A” through “Z” and “a” through “z” are classified as letters (e.g., `UPPER('aäuü')` returns `'AäUü'`). The sort order of two strings is determined using the collation `ucs_basic`, which sorts strings by Unicode code point (e.g., `'Ab' < 'ab'` is true as `A` (`U+0041`) is less than `a` (`U+0061`)). The character with code zero cannot be in a string.

Regular string literals are enclosed in single quotes and don’t recognize escape sequences. Write two adjacent single quotes to include a single-quote character within a string literal (e.g., `'Leonard''s bicycle'`).

Escape string literals are specified by writing the letter `E` (upper or lower case) before the opening single quote, e.g., `E'Firebolt 🔥\U0001F680'`. Use backslash escape sequences within an escape string literal to represent special byte values:

Backslash escape sequence Interpretation `\b` backspace `\f` form feed `\n` newline `\r` carriage return `\t` tab `\o`, `\oo`, `\ooo` (o = 0–7) octal byte value (decimal value must be between 1 and 255) `\xh`, `\xhh` (h = 0–9, A–F) hexadecimal byte value `\uxxxx`, `\Uxxxxxxxx` (x = 0–9, A–F) 16 or 32-bit hexadecimal Unicode character value

Any other character following a backslash is taken literally (e.g., write two backslashes `\\` to include one backslash character). The byte sequences you create must be valid UTF-8. For historic reasons, if you set the setting `standard_conforming_strings` to `false`, regular string literals will also recognize backslash escape sequences.

## [](#binary)Binary

### [](#bytea)BYTEA

Represents variable size binary data. A binary string is a sequence of bytes - unlike TEXT, there is no character set. The BYTEA data type is nullable. For more information, see [BYTEA data type](/sql_reference/bytea-data-type.html).

## [](#spatial)Spatial

## [](#geography)GEOGRAPHY

Represents geographic objects like points, LineStrings, Polygons, or collections thereof. For more information, see [GEOGRAPHY data type](geography-data-type).

## [](#type-conversion)Type Conversion

Values with a given data type can be converted to another data type. There are three contexts in which this happens:

- *Explicit*: With an explicit invocation of the [CAST](/sql_reference/functions-reference/conditional-and-miscellaneous/cast.html) function.
- *Assignment*: Assigning values to a column of the target data type, as happens in the insert statement.
- *Implicit*: Using a SQL function where none of the available signatures match the argument types. The planner inserts implicit casts.

The following table lists which type conversions are supported and in which context.

“Explicit”, “Assignment”, and “Implicit” indicate in which type conversion context the conversion operation can be invoked.

- “Explicit” means only as an explicit cast (using [CAST](/sql_reference/functions-reference/conditional-and-miscellaneous/cast.html) or :: syntax).
- “Assignment” means implicitly in assignment to a target column, as well as explicitly.
- “Implicit” means implicitly in expressions, as well as the other cases.

From \\ To UNKNOWN INT BIGINT REAL DOUBLE TEXT BYTEA BOOLEAN NUMERIC ARRAY DATE TIMESTAMP TIMESTAMPTZ GEOGRAPHY **UNKNOWN**   Implicit Implicit Implicit Implicit Implicit Implicit Implicit Implicit Implicit Implicit Implicit Implicit Implicit **INT**     Implicit Implicit Implicit Assignm.   Explicit Implicit           **BIGINT**   Assignm.   Assignm. Implicit Assignm.   Explicit Implicit           **REAL**   Assignm. Assignm.   Implicit Assignm.   Explicit Assignm.           **DOUBLE**   Assignm. Assignm. Assignm.   Assignm.   Explicit Assignm.           **TEXT**   Assignm. Assignm. Assignm. Assignm.   Explicit Explicit Explicit Assignm. Assignm. Assignm. Assignm. Assignm. **BYTEA**           Explicit               Implicit **BOOLEAN**   Explicit       Assignm.                 **NUMERIC**   Assignm. Assignm. Explicit Implicit Assignm.     Explicit           **DATE**           Assignm.           Implicit Implicit   **TIMESTAMP**           Assignm.         Implicit   Implicit   **TIMESTAMPTZ**           Assignm.         Implicit Implicit     **GEOGRAPHY**           Assignm. Implicit