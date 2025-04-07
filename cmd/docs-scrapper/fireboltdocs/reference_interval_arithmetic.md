# [](#arithmetic-with-intervals)Arithmetic with intervals

This topic describes the Firebolt implementation of arithmetic with intervals.

- [Overview](#overview)
- [Literal string interpretation](#literal-string-interpretation)
  
  - [Format examples](#format-examples)
- [Arithmetic between interval and TIMESTAMPTZ](#arithmetic-between-interval-and-timestamptz)
- [Examples](#examples)

## [](#overview)Overview

An `interval` represents a duration. In Firebolt, values of type `interval` can be used to add or subtract a duration to/from a date or timestamp. `Interval` cannot be used as the data type of a column.

The `+` and `*` operators shown below come in commutative pairs (e.g., both `DATE + interval` and `interval + DATE` are accepted).

Operator Description `DATE + interval -> TIMESTAMP` Add an `interval` to a `DATE` `DATE - interval -> TIMESTAMP` Subtract an `interval` from a `DATE` `TIMESTAMP + interval -> TIMESTAMP` Add an `interval` to a `TIMESTAMP` `TIMESTAMP - interval -> TIMESTAMP` Subtract an `interval` from a `TIMESTAMP` `TIMESTAMPTZ + interval -> TIMESTAMPTZ` Add an `interval` to a `TIMESTAMPTZ` `TIMESTAMPTZ - interval -> TIMESTAMPTZ` Subtract an `interval` from a `TIMESTAMPTZ` `interval * DOUBLE PRECISION -> interval` Multiply an `interval` by a scalar

## [](#literal-string-interpretation)Literal string interpretation

`Interval` literals can be specified in two formats.

### [](#format-examples)Format examples

**Cast from text literal**  
The following examples demonstrate different ways to cast a text literal to an interval in SQL. Each method converts a string representing a time interval into the `INTERVAL` data type using various casting syntaxes.

Using interval keyword:

```
interval 'quantity unit [quantity unit...] [direction]'
```

Using double colon:

```
'quantity unit [quantity unit...] [direction]'::interval
```

Using `CAST` function:

```
CAST('quantity unit [quantity unit...] [direction]' AS interval)
```

In these examples, `direction` can be either `ago` or left empty. Using `ago` negates all specified quantities. The `quantity` represents a signed or unsigned integer, and `unit` refers to one of the following time units, matched case-insensitively:

Unit microsecond\[s] / us millisecond\[s] / ms second\[s] / s minute\[s] / m hour\[s] / h day\[s] / d week\[s] / w month\[s] / mon\[s] year\[s] / y decade\[s] / dec\[s] century / centuries / c millennium\[s] / mil\[s]

Each `unit` can appear only once in an interval literal. The value of the interval is determined by adding the quantities of the specified units with the appropriate signs.

**Unit outside of text literal**  
The following example demonstrates how to cast a numeric value to an interval by placing the unit outside of the text literal.

```
interval 'N' unit
```

In this format, `N` represents a signed or unsigned integer, and `unit` specifies the time unit, matched case-insensitively, from the following options:

Unit second minute hour day week month year

## [](#arithmetic-between-interval-and-timestamptz)Arithmetic between interval and TIMESTAMPTZ

Interval arithmetic with `TIMESTAMPTZ` values works as follows:

1. Convert the `TIMESTAMPTZ` value from Unix time to local time according to the rules of the time zone specified by the session’s `time_zone` setting.
2. Add the `millennium`, `century`, `decade`, `year`, `month`, `week` and `day` components of the interval to the local time.
3. Convert the local time back to Unix time according to the rules of the time zone specified by the session’s `time_zone` setting.
4. Add the `hour`, `minute`, `second`, `millisecond`, and `microsecond` components of the interval to the Unix time.

The back and forth between Unix time and local time is necessary to handle the fact that not all days consist of 24 hours due to daylight savings time transitions. For instance, `SELECT TIMESTAMPTZ '2022-10-30 Europe/Berlin' + interval '1 day'` returns `2022-10-31 00:00:00+01` but `SELECT TIMESTAMPTZ '2022-10-30 Europe/Berlin' + interval '24 hours'` returns `2022-10-30 23:00:00+01` (assuming the value of the session’s `time_zone` setting is `'Europe/Berlin'`). Still, the dependence on the session’s `time_zone` setting should be kept in mind when doing arithmetic between interval and `TIMESTAMPTZ`.

### [](#multiplying-an-interval-by-a-scalar)Multiplying an interval by a scalar

You can use the expression `date_time + INTERVAL * d` where `date_time` is a constant or column reference of type `DATE`, `TIMESTAMP`, or `TIMESTAMPTZ`, and `d` is a constant or column reference of type `DOUBLE PRECISION`. The effect is that the INTERVAL is scaled by `d`, and the resulting INTERVAL is added to `date_time`. E.g., `INTERVAL '1 day' * 3` -&gt; `INTERVAL '3 days'`.

## [](#examples)Examples

```
SELECT DATE '1996-09-03' - interval '1 millennium 5 years 42 day 42 ms';  --> 0991-07-22 23:59:59.958
SELECT TIMESTAMP '1996-09-03 11:19:42' + interval '10 years 5 months 42 days 7 seconds';  --> 2007-03-17 11:19:49

SELECT TIMESTAMP '2023-10-20 11:49:52' + interval '1 year 6 months 4 weeks 7 hours' * 7.5;  --> 2035-08-20 16:19:52
SELECT DATE '2023-10-20' - 42 * interval '1 months 1 day 1 hour';  --> 2020-03-07 06:00:00

-- The following example shows a daylight savings time change in the time zone 'Europe/Berlin'
SET time_zone = 'Europe/Berlin';
SELECT TIMESTAMPTZ '2022-10-30 Europe/Berlin' + interval '1 day';  --> 2022-10-31 00:00:00+01
SELECT TIMESTAMPTZ '2022-10-30 Europe/Berlin' + interval '24' hour;  --> 2022-10-30 23:00:00+01

SET time_zone = 'US/Pacific';
SELECT TIMESTAMPTZ '2022-10-30 Europe/Berlin' + interval '1 day';  --> 2022-10-30 15:00:00-07
SELECT TIMESTAMPTZ '2022-10-30 Europe/Berlin' + interval '24' hour;  --> 2022-10-30 15:00:00-07
```