# [](#to_yyyymmdd)TO\_YYYYMMDD

Extracts year, month and day from a `DATE`, `TIMESTAMP`, or `TIMESTAMPTZ` value and combines them into an integer beginning with the four-digit year followed by the two-digit month followed by the two-digit day. `TO_YYYYMMDD(<expression>)` is equivalent to `EXTRACT(YEAR FROM <expression>) * 10000 + EXTRACT(MONTH FROM <expression>) * 100 + EXTRACT(DAY FROM <expression>);`

## [](#syntax)Syntax

```
TO_YYYYMMDD(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The expression from which the time units are extracted. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ`

`TIMESTAMPTZ` values are converted to local time according to the session’s `time_zone` setting before extracting the time units.

## [](#return-types)Return Types

`INT`

## [](#remarks)Remarks

The `TO_YYYYMMDD` function can be used in the `PARTITION BY` clause of `CREATE TABLE` commands.

```
CREATE TABLE test (
  t TIMESTAMP
)
PARTITION BY TO_YYYYMMDD(t);
```

## [](#example)Example