# [](#count)COUNT

Counts the number of rows or not NULL values.

## [](#syntax)Syntax

```
COUNT([ DISTINCT ] <expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The expression to count Any

Valid values for the input expression include column names or functions that return a column name. When `DISTINCT` is being used, only the unique number of rows with no `NULL` values are counted. `COUNT(*)` returns a total count of all rows in the table, while `COUNT(<column_name>)` returns a count of non-null rows in the specified `<column_name>`.

> By default, `COUNT(DISTINCT)` returns exact results. If you do not require a precise result and want to have faster performance, consider using the APPROX\_COUNT\_DISTINCT function. See below for examples and considerations.

## [](#return-type)Return Type

`NUMERIC`

## [](#example)Example

For this example, see the following table `tournaments`:

name totalprizedollars The Drifting Thunderdome 24768 The Lost Track Showdown 5336 The Acceleration Championship 19274 The Winter Wilderness Rally 21560 The Circuit Championship 9739 The Singapore Grand Prix 19274

Doing a regular `COUNT` returns the total number of rows in the column. As the `tournaments` table contains 6 rows, this will be the returned value.

```
SELECT
	COUNT(name)
FROM
	tournaments;
```

**Returns**

`6`

A `COUNT(DISTINCT)` function on the same column returns the number of unique rows. When applied to the `totalprizedollars` column, the value returned is `5`, as there is a repeated number in the column.

```
SELECT
	COUNT(DISTINCT totalprizedollars)
FROM
	tournaments;
```

**Returns**

`5`

## [](#example-of-countdistinct-vs-approx_count_distinct)Example of COUNT(DISTINCT) vs. APPROX\_COUNT\_DISTINCT

To understand the difference between `COUNT(DISTINCT pk)` with exact precision enabled and using default approximation, consider a table, `count_test` with 8,388,608 unique `pk` values. The `APPROX_COUNT_DISTINCT` function returns the same approximate results as the `COUNT(DISTINCT)` function with exact precision disabled, so we can see the difference between these methods with the following example.

```
SELECT
	COUNT(DISTINCT pk) as count_distinct,
	APPROX_COUNT_DISTINCT(pk) as approx_count
FROM
	count_test;
```

**Returns**

Assuming 8,388,608 unique pk values, the previous code example returns the following:

```
' +----------------+--------------+
' | count_distinct | approx_count |
' +----------------+--------------+
' |      8,388,608 |    8,427,387 |
' +----------------+--------------+
```