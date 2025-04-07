# [](#generate_series)GENERATE\_SERIES

Generates a single rowset of values from `start` to `stop`, with a step size of `step`. `GENERATE_SERIES` is a table-valued function.

## [](#syntax)Syntax

```
GENERATE_SERIES ( <start>, <stop> [, <step> ] )
```

## [](#parameters)Parameters

Parameter Description Supported input types `<start>` The first value in the interval. `BIGINT`, `INTEGER` `<stop>` The last value in the interval.  
The series stops once the last generated step value exceeds the stop value. `BIGINT`, `INTEGER` `<step>` Optional literal integer value to set step. If not included, the default step is 1. `BIGINT`, `INTEGER`

## [](#return-type)Return Type

Setof `INTEGER` if all operands are of type `INTEGER`, otherwise setof `BIGINT`.

## [](#example)Example

```
SELECT n, DATE_ADD('DAY', n, '2023-02-02') result 
FROM GENERATE_SERIES(1, 10, 2) s(n)
```

**Returns**:

n result 1 2023-02-03 00:00:00 3 2023-02-05 00:00:00 5 2023-02-07 00:00:00 7 2023-02-09 00:00:00 9 2023-02-11 00:00:00