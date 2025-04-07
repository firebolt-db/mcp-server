# [](#to_yyyymm)TO\_YYYYMM

Extracts the year and month from a `DATE`, `TIMESTAMP`, or `TIMESTAMPTZ` value and combines them into an integer beginning with the four-digit year followed by the two-digit month. `TO_YYYYMM(<expression>)` is equivalent to `EXTRACT(YEAR FROM <expression>) * 100 + EXTRACT(MONTH FROM <expression>);`

## [](#syntax)Syntax

```
TO_YYYYMM(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The expression from which the time units are extracted. `DATE`, `TIMESTAMP`, `TIMESTAMPTZ`

`TIMESTAMPTZ` values are converted to local time according to the session’s `time_zone` setting before extracting the time units.

## [](#return-types)Return Types

`INT`

## [](#remarks)Remarks

The `TO_YYYYMM` function can be used in the `PARTITION BY` clause of `CREATE TABLE` commands.

```
CREATE TABLE test (
  t TIMESTAMP
)
PARTITION BY TO_YYYYMM(t);
```

## [](#example)Example