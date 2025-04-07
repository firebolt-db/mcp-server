# [](#timestamp-data-type)TIMESTAMP data type

This topic describes the Firebolt implementation of the `TIMESTAMP` data type.

- [Overview](#overview)
- [Literal string interpretation](#literal-string-interpretation)
- [Functions and operators](#functions-and-operators)
  
  - [Type conversions](#type-conversions)
  - [Comparison operators](#comparison-operators)
  - [Arithmetic operators](#arithmetic-operators)
- [Serialization and deserialization](#serialization-and-deserialization)

## [](#overview)Overview

Name Size Min Max Resolution `TIMESTAMP` 8 bytes `0001-01-01 00:00:00.000000` `9999-12-31 23:59:59.999999` 1 microsecond

The `TIMESTAMP` data type represents a date and time with microsecond resolution independent of a time zone. To represent an absolute point in time, use [TIMESTAMPTZ](/sql_reference/timestamptz-data-type.html).

## [](#literal-string-interpretation)Literal string interpretation

`TIMESTAMP` literals follow the ISO 8601 and RFC 3339 format: `YYYY-[M]M-[D]D[( |T)[h]h:[m]m:[s]s[.f]]`.

- `YYYY`: Four-digit year (`0001` - `9999`)
- `[M]M`: One or two digit month (`01` - `12`)
- `[D]D`: One or two digit day (`01` - `31`)
- `( |T)`: A space or `T` separator
- `[h]h`: One or two digit hour (`00` - `23`)
- `[m]m`: One or two digit minute (`00` - `59`)
- `[s]s`: One or two digit second (`00` - `59`)
- `[.f]`: Up to six digits after the decimal separator (`000000` - `999999`)

If only the date is specified, the time is assumed to be `00:00:00.000000`.

**Examples**

```
SELECT TIMESTAMP '1996-09-03 11:19:33.123456';
SELECT '2023-02-13'::TIMESTAMP;  --> 2023-02-13 00:00:00
SELECT CAST('2019-7-23T16:9:3.1' as TIMESTAMP);  --> 2019-07-23 16:09:03.1
```

## [](#functions-and-operators)Functions and operators

### [](#type-conversions)Type conversions

The `TIMESTAMP` data type can be cast to and from types as follows:

#### [](#to-timestamp)To TIMESTAMP

From type Example Note `TEXT` `SELECT CAST(TEXT '2023-02-13 11:19:42' as TIMESTAMP); --> 2023-02-13 11:19:42` Interprets the text according to the literal string format. `DATE` `SELECT CAST(DATE '2023-02-13' as TIMESTAMP); --> 2023-02-13 00:00:00` Extends the date with `00:00:00`. `TIMESTAMPTZ` `SELECT CAST(TIMESTAMPTZ '2023-02-13 Europe/Berlin' as TIMESTAMP); --> 2023-02-13 22:00:00` Converts the timestamptz value to local time in the time zone specified by the [session’s `time_zone` setting](/Reference/system-settings.html#setting-the-time-zone). This example assumes `SET time_zone = 'UTC';`.

#### [](#from-timestamp)From TIMESTAMP

To type Example Note `TEXT` `SELECT CAST(TIMESTAMP '2023-02-13 11:19:42' as TEXT); --> 2023-02-13 11:19:42` Converts the timestamp to text in the format `YYYY-MM-DD hh:mm:ss[.f]`. `DATE` `SELECT CAST(TIMESTAMP '2023-02-13 11:19:42' as DATE); --> 2023-02-13` Truncates the timestamp to date. `TIMESTAMPTZ` `SELECT CAST(TIMESTAMP '2023-02-13 11:19:42' as TIMESTAMPTZ); --> 2023-02-13 11:19:42+00` Interprets the timestamp to be local time in the time zone specified by the [session’s `time_zone` setting](/Reference/system-settings.html#setting-the-time-zone). This example assumes `SET time_zone = 'UTC';`.

### [](#comparison-operators)Comparison operators

Operator Description `TIMESTAMP < TIMESTAMP` Less than `TIMESTAMP > TIMESTAMP` Greater than `TIMESTAMP <= TIMESTAMP` Less than or equal to `TIMESTAMP >= TIMESTAMP` Greater than or equal to `TIMESTAMP = TIMESTAMP` Equal to `TIMESTAMP <> TIMESTAMP` Not equal to

A `TIMESTAMP` value is also comparable with a `DATE` or `TIMESTAMPTZ` value:

- The `DATE` value is cast to the `TIMESTAMP` type for comparison with a `TIMESTAMP` value.
- The `TIMESTAMP` value is cast to the `TIMESTAMPTZ` type for comparison with a `TIMESTAMPTZ` value.

### [](#arithmetic-operators)Arithmetic operators

Arithmetic with intervals can be used to add or subtract a duration to or from a timestamp. The result is of type `TIMESTAMP`.

```
SELECT TIMESTAMP '1996-09-03' + INTERVAL '42' YEAR;  --> 2038-09-03 00:00:00
SELECT TIMESTAMP '2023-03-18' - INTERVAL '26 years 5 months 44 days 12 hours 41 minutes';  --> 1996-09-03 11:19:00
```

For more information, see [Arithmetic with intervals](/Reference/interval-arithmetic.html).

## [](#serialization-and-deserialization)Serialization and deserialization

### [](#text-csv-json)Text, CSV, JSON

In the text, CSV, and JSON format, a `TIMESTAMP` value is output as a `YYYY-MM-DD hh:mm:ss[.f]` string. Firebolt outputs as few digits after the decimal separator as possible (at most six). Input is accepted in the literal format described above: `YYYY-[M]M-[D]D[( |T)[h]h:[m]m:[s]s[.f]]`.

### [](#parquet)Parquet

`TIMESTAMP` maps to Parquet’s 64-bit signed integer `TIMESTAMP` type with the parameter `isAdjustedToUTC` set to `false` and `unit` set to `MICROS`, also representing the number of microseconds before or after `1970-01-01 00:00:00.000000`.

### [](#avro)Avro

It’s not possible to import directly from Avro into a `TIMESTAMP` column. Instead, first import using a `TIMESTAMPTZ` column and then use the `AT TIME ZONE` expression to convert to `TIMESTAMP`.

### [](#orc)ORC

`TIMESTAMP` maps to ORC’s “timezone-unaware” logical type “timestamp”.