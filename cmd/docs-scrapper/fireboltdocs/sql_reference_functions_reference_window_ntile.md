# [](#ntile)NTILE

Divides an ordered data set equally into the number of buckets specified by the argument value. Buckets are sequentially numbered 1 through the argument value.

For more information on usage, please refer to [Window Functions](/sql_reference/functions-reference/window/).

## [](#syntax)Syntax

```
NTILE( <value> ) OVER ( [ PARTITION BY <partition_by> ] ORDER BY <order_by> [ { ASC | DESC } ] )
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` An integer expression used for the `NTILE()` function to specify the number of buckets for division. `INTEGER`/`BIGINT` `<partition_by>` An expression used for the partition by clause. Any `<order_by>` An expression used for the order by clause. Any

## [](#return-type)Return Type

`INTEGER`/`BIGINT`

If there is a remainder after dividing the rows in a partition by the argument value, it will result in buckets of different sizes. For example,

- `NTILE(2)` over 5 rows will result in 2 buckets, the first with 3 rows, the second with 2
- `NTILE(3)` over 5 rows results in 3 buckets, the first 2 with 2 rows each and the last with one.

## [](#example)Example

The example below divides students test results into three buckets depending on their score.

```
SELECT
	student_name,
	score,
	NTILE(3) OVER (ORDER BY score) AS bucket
FROM
	student_test_results
ORDER BY score;
```

**Returns**:

student\_name score bucket James 9 1 Emma 10 1 Harry 11 1 Liam 11 2 Ava 12 2 Charly 12 2 Noah 13 3 Olivia 13 3