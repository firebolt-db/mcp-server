# [](#percent_rank)PERCENT\_RANK

Calculates the relative rank of the current row within an ordered data set, as `( rank - 1 ) / ( rows - 1 )` where rank is the current row’s rank within the partition, and rows is the number of rows in the partition. PERCENT\_RANK always returns values from 0 to 1 inclusive. The first row in any set has a `PERCENT_RANK` of 0.

## [](#syntax)Syntax

```
PERCENT_RANK() OVER ( [ PARTITION BY <partition_by> ] ORDER BY <order_by> [ASC|DESC] )
```

## [](#parameters)Parameters

Parameter Description Supported input types Parameter Description   `<partition_by>` An expression used for the partition by clause. Any `<order_by>` An expression used for the order by clause. Any

## [](#return-type)Return Type

`DOUBLE PRECISION`

This function respects `NULL` values, and results will be ordered with default null ordering `NULLS LAST` unless otherwise specified in the `ORDER BY` clause.

## [](#example)Example

The example below calculates, for each student in grade nine, the percent rank of the student’s test score by their grade level.

```
SELECT
	nickname, current_score,
	PERCENT_RANK() OVER (PARTITION BY level ORDER BY current_score DESC) as percent_rank
FROM
	class_test
WHERE grade_level=9;
```

**Returns**:

nickname current\_score percent\_rank kennethpark 90 0 sabrina21 85 0.2 rileyjon 80 0.4 ymatthews 79 0.6