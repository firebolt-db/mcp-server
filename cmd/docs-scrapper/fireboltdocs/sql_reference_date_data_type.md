# [](#date-data-type)DATE data type

This topic describes the Firebolt implementation of the `DATE` data type.

- [Overview](#overview)
- [Literal string interpretation](#literal-string-interpretation)
- [Functions and operators](#functions-and-operators)
  
  - [Type conversions](#type-conversions)
  - [Comparison operators](#comparison-operators)
  - [Arithmetic operators](#arithmetic-operators)
- [Serialization and deserialization](#serialization-and-deserialization)

## [](#overview)Overview

Name Size Min Max Resolution `DATE` 4 bytes `0001-01-01` `9999-12-31` 1 day

The `DATE` type represents a calendar date, independent of a time zone.

## [](#literal-string-interpretation)Literal string interpretation

`DATE` literals follow the ISO 8601 format: `YYYY-[M]M-[D]D`.

- `YYYY`: Four-digit year (`0001` - `9999`)
- `[M]M`: One or two digit month (`01` - `12`)
- `[D]D`: One or two digit day (`01` - `31`)

**Examples**

```
SELECT DATE '2023-02-13';
SELECT '2023-02-13'::DATE;
SELECT CAST('2023-6-03' AS DATE);
SELECT DATE('2023-02-13');
```

## [](#functions-and-operators)Functions and operators

### [](#type-conversions)Type conversions

The `DATE` data type can be cast to and from types as follows:

#### [](#to-date)To DATE

From type Example Note `TEXT` `SELECT CAST(TEXT '2023-02-13' as DATE); --> 2023-02-13` Interprets the text according to the literal string format. `TIMESTAMP` `SELECT CAST(TIMESTAMP '2023-02-13 11:19:42' as DATE); --> 2023-02-13` Truncates the timestamp to the date. `TIMESTAMPTZ` `SELECT CAST(TIMESTAMPTZ '2023-02-13 Europe/Berlin' as DATE); --> 2023-02-13` Converts the timestamptz to local time in the time zone specified by the [session’s `timezone` setting](/Reference/system-settings.html#setting-the-time-zone), then truncates the timestamp to the date. This example assumes `SET timezone = 'UTC';`.

#### [](#from-date)From DATE

To type Example Note `TEXT` `SELECT CAST(DATE '2023-02-13' as TEXT); --> 2023-02-13` Converts the date to text in the format `YYYY-MM-DD`. `TIMESTAMP` `SELECT CAST(DATE '2023-02-13' as TIMESTAMP); --> 2023-02-1 00:00:00` Extends the date with `00:00:00`. `TIMESTAMPTZ` `SELECT CAST(DATE '2023-02-13' as TIMESTAMPTZ); --> 2023-02-13` Interprets the date as midnight in the time zone specified by the [session’s `timezone` setting](/Reference/system-settings.html#setting-the-time-zone). This example assumes `SET timezone = 'UTC';`.

### [](#comparison-operators)Comparison operators

Operator Description `DATE < DATE` Less than `DATE > DATE` Greater than `DATE <= DATE` Less than or equal to `DATE >= DATE` Greater than or equal to `DATE = DATE` Equal to `DATE <> DATE` Not equal to

A `DATE` value is also comparable with a `TIMESTAMP` or `TIMESTAMPTZ` value:

- The `DATE` value is cast to the `TIMESTAMP` type for comparison with a `TIMESTAMP` value.
- The `DATE` value is cast to the `TIMESTAMPTZ` type for comparison with a `TIMESTAMPTZ` value.

### [](#arithmetic-operators)Arithmetic operators

The `+` operators described below come in commutative pairs:

Operator Description Example `DATE + INTEGER -> DATE` Add a number of days to a date `SELECT DATE '2023-03-03' + 42; --> 2023-04-14` `DATE - INTEGER -> DATE` Subtract a number of days from a date `SELECT DATE '2023-03-03' - 42; --> 2023-01-20` `DATE - DATE -> INTEGER` Subtract dates, producing the number of elapsed days `SELECT DATE '2023-03-03' - DATE '1996-09-03'; --> 9677` `DATE + INTERVAL -> TIMESTAMP` Add an interval to a date `SELECT DATE '1996-09-03' + INTERVAL '42' YEAR; --> 2038-09-03 00:00:00` `DATE - INTERVAL -> TIMESTAMP` Subtract an interval from a date `SELECT DATE '2023-03-18' - INTERVAL '26 years 5 months 44 days 12 hours 41 minutes'; --> 1996-09-03 11:19:00`

#### [](#interval-arithmetic)Interval arithmetic

Arithmetic with intervals can be used to add or subtract a duration to or from a date. The result is of type `TIMESTAMP`.

**Example**

```
SELECT DATE '1996-09-03' + INTERVAL '42' YEAR;  --> 2038-09-03 00:00:00
SELECT DATE '2023-03-18' - INTERVAL '26 years 5 months 44 days 12 hours 41 minutes';  --> 1996-09-03 11:19:00
```

For more information, see [Arithmetic with intervals](/Reference/interval-arithmetic.html).

## [](#serialization-and-deserialization)Serialization and deserialization

### [](#text-csv-json)Text, CSV, JSON

In the text, CSV, and JSON format, a `DATE` value is output as a `YYYY-MM-DD` string. Input is accepted in the literal format described above: `YYYY-[M]M-[D]D`.

### [](#parquet)Parquet

`DATE` maps to Parquet’s 32-bit signed integer `DATE` type, also representing the number of days before or after `1970-01-01`.

### [](#avro)Avro

`DATE` maps to Avro’s 32-bit signed integer `DATE` type, also representing the number of days before or after `1970-01-01`.

### [](#orc)ORC

`DATE` maps to ORC’s signed integer `DATE` type, also representing the number of days before or after `1970-01-01`.