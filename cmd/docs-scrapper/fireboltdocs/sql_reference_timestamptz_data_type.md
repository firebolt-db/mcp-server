# [](#timestamptz-data-type)TIMESTAMPTZ data type

This topic describes the Firebolt implementation of the `TIMESTAMPTZ` data type.

- [Overview](#overview)
- [Literal string interpretation](#literal-string-interpretation)
- [Daylight savings time transitions](#daylight-savings-time-transitions)
- [Functions and operators](#functions-and-operators)
  
  - [Type conversions](#type-conversions)
    
    - [AT TIME ZONE](#at-time-zone)
  - [Comparison operators](#comparison-operators)
  - [Arithmetic operators](#arithmetic-operators)
- [Serialization and deserialization](#serialization-and-deserialization)

## [](#overview)Overview

Name Size Min Max Resolution `TIMESTAMPTZ` 8 bytes `0001-01-01 00:00:00.000000 UTC` `9999-12-31 23:59:59.999999 UTC` 1 microsecond

The `TIMESTAMPTZ` data type represents an absolute point in time as a date and time with microsecond resolution. Firebolt stores `TIMESTAMPTZ` values as [Unix time](https://en.wikipedia.org/wiki/Unix_time), which is UTC without leap seconds.

### [](#difference-between-timestamptz-and-timestamp)Difference between TIMESTAMPTZ and TIMESTAMP

With `TIMESTAMP`, the time zone is deliberately unspecified. For example, the start of the third millennium was celebrated on New Year’s Day at `TIMESTAMP '2001-01-01 00:00:00'`, independent of the geographical location. However, that doesn’t mean that everybody in the world celebrated at the same absolute point in time. In Munich, Germany, the new year was celebrated at `TIMESTAMPTZ '2001-01-01 00:00:00 Europe/Berlin'`, which was `TIMESTAMPTZ '2000-12-31 23:00:00 UTC'`. Seattle in the United States celebrated nine hours later at `TIMESTAMPTZ '2001-01-01 00:00:00 US/Pacific'`, which was `TIMESTAMPTZ '2001-01-01 08:00:00 UTC'`.

## [](#literal-string-interpretation)Literal string interpretation

`TIMESTAMPTZ` literals can be specified in one of three formats:

1. `local_timestamp [time_zone]`
2. `local_timestamp[time_zone_offset]`
3. `local_timestamp[utc_time_zone]`

<!--THE END-->

- `local_timestamp` follows the ISO 8601 and RFC 3339 format: `YYYY-[M]M-[D]D[( |T)[h]h:[m]m:[s]s[.f]]`.
  
  - `YYYY`: Four-digit year (`0001` - `9999`)
  - `[M]M`: One or two digit month (`01` - `12`)
  - `[D]D`: One or two digit day (`01` - `31`)
  - `( |T)`: A space or `T` separator
  - `[h]h`: One or two digit hour (`00` - `23`)
  - `[m]m`: One or two digit minutes (`00` - `59`)
  - `[s]s`: One or two digit seconds (`00` - `59`)
  - `[.f]`: Up to six digits after the decimal separator (`000000` - `999999`)
- `time_zone` is a string containing the name of the time zone.
- `time_zone_offset` is a string representing the offset from the UTC time zone.
  
  - Format: `(+|-)H[H][:m[m]]`.
- `utc_time_zone`: The letter `Z` representing the UTC time zone.

Time zone names are from the [tz database](http://www.iana.org/time-zones) (see the [list of tz database time zones](http://en.wikipedia.org/wiki/List_of_tz_database_time_zones)). For times in the future, the latest known rule for the given time zone is applied. Firebolt does not support time zone abbreviations, as they cannot account for daylight savings time transitions, and some time zone abbreviations implied different UTC offsets at different times.

If a `TIMESTAMPTZ` literal has an explicit time zone specified, it is converted to Unix time using the appropriate offset. If not, Firebolt uses the session’s `timezone` setting and assumes the `TIMESTAMPTZ` literal is in that time zone. The default value of the `timezone` setting is UTC. To set it to, e.g., `Europe/Berlin`, you can issue: `SET timezone = 'Europe/Berlin';`. For more information, see [system settings](/Reference/system-settings.html#setting-the-time-zone).

If only the date is specified, the time is assumed to be `00:00:00.000000`.

**Example**

```
SET timezone = 'UTC';
SELECT TIMESTAMPTZ '1996-09-03 11:19:33.123456 Europe/Berlin';  --> 1996-09-03 09:19:33.123456+00
SELECT TIMESTAMPTZ '2023-1-29 6:3:42.7-3:30';  --> 2023-01-29 09:33:42.7+00
```

## [](#daylight-savings-time-transitions)Daylight savings time transitions

During a daylight savings time transition, a seemingly valid timestamp can represent a nonexistent or ambiguous timestamp. Firebolt resolves the problem by returning the later time point.

### [](#spring-forward-transitions)“Spring forward” transitions

```
SET timezone = 'UTC';
SELECT TIMESTAMPTZ '2022-03-27 01:59:59 Europe/Berlin';  --> 2022-03-27 00:59:59+00
SELECT TIMESTAMPTZ '2022-03-27 02:00:00 Europe/Berlin';  --> 2022-03-27 01:00:00+00
SELECT TIMESTAMPTZ '2022-03-27 03:00:00 Europe/Berlin';  --> 2022-03-27 01:00:00+00
```

### [](#fall-back-transitions)“Fall back” transitions

```
SET timezone = 'UTC';
SELECT TIMESTAMPTZ '2022-10-30 01:59:59 Europe/Berlin';  --> 2022-10-29 23:59:59+00
SELECT TIMESTAMPTZ '2022-10-30 02:00:00 Europe/Berlin';  --> 2022-10-30 01:00:00+00
SELECT TIMESTAMPTZ '2022-10-30 03:00:00 Europe/Berlin';  --> 2022-10-30 02:00:00+00
```

## [](#functions-and-operators)Functions and operators

### [](#type-conversions)Type conversions

The `TIMESTAMPTZ` data type can be cast to and from types as follows:

#### [](#to-timestamptz)To TIMESTAMPTZ

From type Example Note `TEXT` `SELECT CAST(TEXT '2023-02-13 11:19:42 Europe/Berlin' as TIMESTAMPTZ); --> 2023-02-13 00:00:00+00` Interprets the text according to the literal string format. `DATE` `SELECT CAST(DATE '2023-02-13' as TIMESTAMPTZ); --> 2023-02-13 00:00:00+00` Interprets the timestamp to be midnight in the time zone specified by the session’s `timezone` setting. `TIMESTAMP` `SELECT CAST(TIMESTAMP '2023-02-13 11:19:42' as TIMESTAMPTZ); --> 2023-02-13 11:19:42+00` Interprets the timestamp to be local time in the time zone specified by the session’s `timezone` setting.

#### [](#from-timestamptz)FROM TIMESTAMPTZ

To type Example Note `TEXT` `SELECT CAST(TIMESTAMPTZ '2023-02-13 11:19:42 Europe/Berlin' as TEXT); --> 2023-02-13 10:19:42+00` Converts from Unix time to local time in the time zone specified by the session’s `timezone` setting and appends the UTC offset. `DATE` `SELECT CAST(TIMESTAMPTZ '2023-02-13 11:19:42 Europe/Berlin' as DATE); --> 2023-02-13` Converts from Unix time to local time in the time zone specified by the session’s `timezone` setting and then truncates the timestamp to the date. `TIMESTAMP` `SELECT CAST(TIMESTAMPTZ '2023-02-13 11:19:42 Europe/Berlin' as TIMESTAMP); --> 2023-02-13 10:19:42` Convert from Unix time to local time in the time zone specified by the session’s `timezone` setting.

#### [](#at-time-zone)AT TIME ZONE

The dependence on the session’s `timezone` setting for type conversions can be problematic for implicit conversions. Therefore, we recommend using the `AT TIME ZONE` construct to be explicit about which time zone to use.

- `TIMESTAMP AT TIME ZONE time_zone_str -> TIMESTAMPTZ`  
  Converts the given `TIMESTAMP` to `TIMESTAMPTZ` by interpreting it as local time in the time zone `time_zone_str`.

**Example:** `SELECT TIMESTAMP '1996-09-03' at time zone 'Europe/Berlin' = TIMESTAMPTZ '1996-09-03 Europe/Berlin'; --> 1`

- `TIMESTAMPTZ AT TIME ZONE time_zone_str -> TIMESTAMP`:  
  Converts the given `TIMESTAMPTZ` to `TIMESTAMP` by transforming it from Unix time to local time in the time zone `time_zone_str`.

**Example:** `SELECT TIMESTAMPTZ '1996-09-03 Europe/Berlin' AT TIME ZONE 'US/Pacific'; --> 1996-09-02 15:00:00`

The `AT TIME ZONE` construct cannot be used with values of type `DATE`. However, you can explicitly cast a `DATE` value to `TIMESTAMP` and use the converted value instead.

### [](#comparison-operators)Comparison operators

Operator Description `TIMESTAMPTZ < TIMESTAMPTZ` Less than `TIMESTAMPTZ > TIMESTAMPTZ` Greater than `TIMESTAMPTZ <= TIMESTAMPTZ` Less than or equal to `TIMESTAMPTZ >= TIMESTAMPTZ` Greater than or equal to `TIMESTAMPTZ = TIMESTAMPTZ` Equal to `TIMESTAMPTZ <> TIMESTAMPTZ` Not equal to

A `TIMESTAMPTZ` value is also comparable with a `DATE` or `TIMESTAMP` value:

- The `DATE` value is converted to the `TIMESTAMPTZ` type for comparison with a `TIMESTAMPTZ` value.
- The `TIMESTAMP` value is converted to the `TIMESTAMPTZ` type for comparison with a `TIMESTAMPTZ` value.

### [](#arithmetic-operators)Arithmetic operators

`TIMESTAMPTZ` values can be used for arithmetic with intervals:

```
SET timezone = 'Europe/Berlin';
SELECT TIMESTAMPTZ '2022-10-30 Europe/Berlin' + interval '1 day';  --> 2022-10-31 00:00:00+01
SELECT TIMESTAMPTZ '2022-10-30 Europe/Berlin' + interval '24' hour;  --> 2022-10-30 23:00:00+01

SET timezone = 'US/Pacific';
SELECT TIMESTAMPTZ '2022-10-30 Europe/Berlin' + interval '1 day';  --> 2022-10-30 15:00:00-07
SELECT TIMESTAMPTZ '2022-10-30 Europe/Berlin' + interval '24' hour;  --> 2022-10-30 15:00:00-07
```

For more information, see [Arithmetic with intervals](/Reference/interval-arithmetic.html).

## [](#serialization-and-deserialization)Serialization and deserialization

### [](#text-csv-json)Text, CSV, JSON

In the text, CSV, and JSON format, a `TIMESTAMPTZ` value is shown as local time after conversion from Unix time using the time zone specified in the session’s `timezone` setting. Time zone information using the session’s `timezone` setting is shown as a signed numeric offset from UTC (`hh` if it is an integral number of hours, `hh:mm` if it is an integral number of minutes, else `hh:mm:ss`), with a positive sign for zones east of Greenwich. The date and time components are output as a `YYYY-MM-DD hh:mm:ss[.f]` string. Firebolt outputs as few digits after the decimal separator as possible (at most six). Input is accepted in one of the literal formats described above.

### [](#parquet)Parquet

`TIMESTAMPTZ` maps to Parquet’s 64-bit signed integer `TIMESTAMP` type with the parameter `isAdjustedToUTC` set to `true` and `unit` set to `MICROS`, representing the number of microseconds before or after `1970-01-01 00:00:00.000000 UTC`. It’s also possible to import into a `TIMESTAMPTZ` column from Parquet’s 64-bit signed integer `TIMESTAMP` type with the parameter `isAdjustedToUTC` set to `true` and the `unit` set to `MILLIS` or `NANOS`. In this case, Firebolt implicitly extends or truncates to resolve in microseconds.

### [](#avro)Avro

`TIMESTAMPTZ` maps to Avro’s 64-bit signed integer `timestamp-micros` type, representing the number of microseconds before or after `1970-01-01 00:00:00.000000 UTC`. It’s also possible to import into a `TIMESTAMPTZ` column from Avro’s `timestamp-millis` type.

### [](#orc)ORC

It’s not possible to import directly from ORC into a `TIMESTAMPTZ` column. ORC’s logical `TIMESTAMP` type is independent of a time zone and this would require a possibly unintended time zone conversion on import. Instead, first import using a `TIMESTAMP` column and then use the `AT TIME ZONE` expression to convert to `TIMESTAMPTZ`.