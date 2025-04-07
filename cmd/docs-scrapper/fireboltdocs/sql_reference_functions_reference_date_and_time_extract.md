# [](#extract)EXTRACT

Retrieves the time unit, such as `year` or `hour`, from a `DATE`, `TIMESTAMP`, or `TIMESTAMPTZ` value.

## [](#syntax)Syntax

```
EXTRACT(<time_unit> FROM <expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<time_unit>` The time unit to extract from the expression. `microseconds`, `milliseconds`, `second`, `minute`, `hour`, `day`, `week`, `month`, `quarter`, `year`, `decade`, `century`, `millennium` `<expression>` The expression from which the time unit is extracted. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ`

`TIMESTAMPTZ` values are converted to local time according to the session’s `time_zone` setting before extracting the `time_unit`. The set of allowed `time_unit` values depends on the data type of `<expression>`. Furthermore, the return type depends on the `time_unit`.

### [](#time-units)Time Units

Unit Description Supported input types Return type Example `century` Extract the century. The first century starts on `0001-01-01` and ends on `0100-12-31`. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(century FROM TIMESTAMP '0100-12-31'); --> 1` `day` Extract the day (of the month) field. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(day FROM DATE '2001-02-16'); --> 16` `decade` Extract the year field divided by 10. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(decade FROM DATE '0009-12-31'); --> 0` `dow` Extract the day of the week as Sunday (0) to Saturday (6). `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(dow FROM DATE '2022-10-13'); --> 4` `doy` Extract the day of the year (1–365/366). `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(doy FROM DATE '1972-02-29'); --> 60` `epoch` For `TIMESTAMPTZ`, extract the number of seconds since `1970-01-01 00:00:00 UTC`. For `TIMESTAMP`, extract the number of seconds since `1970-01-01 00:00:00` independent of a time zone. `DATE` expressions are implicitly converted to `TIMESTAMP`. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `DECIMAL(38, 9)` `SELECT EXTRACT(epoch FROM TIMESTAMP '2001-02-16 20:38:40.12'); --> 982355920.120000000` `hour` Extract the hour field (0–23). `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(hour FROM TIMESTAMP '2001-02-16 20:38:40.12'); --> 20` `isodow` Extract the day of the week as Monday (1) to Sunday (7). `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(isodow FROM DATE '2022-10-13'); --> 4` `isoyear` Extract the ISO 8601 week-numbering year that the date falls in. Each ISO 8601 week-numbering year begins with the Monday of the week containing the 4th of January; so in early January or late December the ISO year may be different from the Gregorian year. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(isoyear FROM DATE '2006-01-01'); --> 2005` `microseconds` Extract the seconds field, including fractional parts, multiplied by 1,000,000. `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(microseconds FROM TIMESTAMP '2001-02-16 20:38:40.12'); --> 40120000` `millennium` Extract the millennium. The third millennium started on 2001-01-01. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(millennium FROM TIMESTAMP '1000-12-31 23:59:59.999999'); --> 1` `milliseconds` Extract the seconds field, including fractional parts, multiplied by 1,000. `TIMESTAMP`, `TIMESTAMPTZ` `DECIMAL(38, 9)` `SELECT EXTRACT(milliseconds FROM TIMESTAMP '2001-02-16 20:38:40.12'); --> 40120.000000000` `minute` Extract the minutes field (0–59). `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(minute FROM TIMESTAMP '1000-12-31 23:42:59'); --> 42` `month` Extract the number of the month within the year (1–12). `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(month FROM DATE '1000-12-31'); --> 12` `quarter` Extract the quarter of the year (1–4) that the date is in:  
`[01, 03] -> 1`  
`[04, 06] -> 2`  
`[07, 09] -> 3`  
`[10, 12] -> 4` `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(quarter FROM DATE '1000-10-31'); --> 4` `second` Extract the second’s field, including fractional parts. `TIMESTAMP`, `TIMESTAMPTZ` `DECIMAL(38, 9)` `SELECT EXTRACT(second FROM TIMESTAMP '2001-02-16 20:38:40.12'); --> 40.120000000` `timezone` Extract the time zone offset from UTC, measured in seconds, with a positive sign for zones east of Greenwich. `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(timezone FROM TIMESTAMPTZ '2022-11-29 13:58:23 Europe/Berlin'); --> -28800` (assumes set time zone is ‘US/Pacific’) `timezone_hour` Extract the hour component of the time zone offset. `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(timezone_hour FROM TIMESTAMPTZ '2022-11-29 13:58:23 Europe/Berlin'); --> -8` (assumes set time zone is ‘US/Pacific’) `timezone_minute` Extract the minute component of the time zone offset. `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(timezone_minute FROM TIMESTAMPTZ '2022-11-29 13:58:23 Europe/Berlin'); --> 0` (assumes set time zone is ‘US/Pacific’) `week` Extract the number of the ISO 8601 week-numbering week of the year. By definition, ISO weeks start on Mondays and the first week of a year contains January 4 of that year. It is possible for early-January dates to be part of the 52nd or 53rd week of the previous year, and for late-December dates to be part of the first week of the next year. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ`. `INTEGER` `SELECT EXTRACT(week FROM DATE '2005-01-01'); --> 53`  
`SELECT EXTRACT(week from DATE '2006-01-01'); --> 52` `year` Extract the year field. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` `INTEGER` `SELECT EXTRACT(year FROM TIMESTAMP '2001-02-16'); --> 2001`

## [](#return-types)Return Types

Depending on the requested time unit, either an integer or a decimal.

## [](#remarks)Remarks

The `EXTRACT` function can be used in the `PARTITION BY` clause of `CREATE TABLE` commands.

```
CREATE DIMENSION TABLE test (
  d DATE,
  t TIMESTAMP
)
PARTITION BY EXTRACT(month FROM d), EXTRACT(hour FROM t);
```