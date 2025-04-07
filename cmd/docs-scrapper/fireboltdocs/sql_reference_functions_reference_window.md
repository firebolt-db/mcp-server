## [](#window-functions)Window functions

Some functions support an optional `frame_clause`.

The `frame_clause` can be one of the following:

```
    { RANGE | ROWS } <frame_start>
    { RANGE | ROWS } BETWEEN <frame_start> AND <frame_end>
```

where `<frame_start>` and `<frame_end>` is one of the following:

```
  UNBOUNDED PRECEDING
  offset PRECEDING
  CURRENT ROW
  offset FOLLOWING
  UNBOUNDED FOLLOWING
```

The `frame_clause` specifies the set of rows constituting the window frame within the current partition. The frame can be specified in `RANGE` or `ROWS` mode; in each case, the frame runs from the `<frame_start>` to the `<frame_end>`. If `<frame_end>` is omitted, the end defaults to `CURRENT ROW`.

**Usage**

- The default framing option is `RANGE UNBOUNDED PRECEDING`, which is the same as `RANGE BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW`. With an `ORDER BY` clause, this sets the frame to be all rows from the partition start through the current row’s last `ORDER BY` peer. Without an `ORDER BY` clause, all rows of the partition are included in the window frame, since all rows become peers of the current row.
- The number of rows to the end of the frame is limited by the number of rows to the end of the partition; for rows near the partition ends, the frame might contain fewer rows than elsewhere.

**UNBOUNDED PRECEDING, UNBOUNDED FOLLOWING**

- A `<frame_start>` of `UNBOUNDED PRECEDING` means that the frame starts with the first row of the partition. Similarly a `<frame_end>` of `UNBOUNDED FOLLOWING` means that the frame ends with the last row of the partition.

**CURRENT ROW**

- In `RANGE` mode, a `<frame_start>` of `CURRENT ROW` means the frame starts with the current row’s first peer row (a row that the window’s `ORDER BY` clause sorts as equivalent to the current row), while a `<frame_end>` of `CURRENT ROW` means the frame ends with the current row’s last peer row.
- In `ROWS` mode, `CURRENT ROW` simply means the current row.

**offset PRECEDING, offset FOLLOWING**

- For the `offset PRECEDING` and `offset FOLLOWING` frame options, the offset must be an expression not containing any variables, aggregate functions, or window functions. The meaning of the offset depends on the frame mode:
  
  - In `ROWS` mode, the offset must yield a non-null, non-negative integer, and the option means that the frame starts or ends the specified number of rows before or after the current row.
  - In `RANGE` mode, these options require that the `ORDER BY` clause specify exactly one column. The offset specifies the maximum difference between the value of that column in the current row and its value in preceding or following rows of the frame. The data type of the offset expression varies depending on the data type of the ordering column. For numeric ordering columns it is typically of the same type as the ordering column. The offset is still required to be non-null and non-negative, though the meaning of “non-negative” depends on the data type.
  - In `ROWS` mode, `0 PRECEDING` and `0 FOLLOWING` are equivalent to `CURRENT ROW`. This normally holds in `RANGE` mode as well, for an appropriate data-type-specific meaning of “zero”.

**Restrictions**

- `frame_start` cannot be `UNBOUNDED FOLLOWING`
- `frame_end` cannot be `UNBOUNDED PRECEDING`
- `frame_end` cannot appear earlier in the above list of `frame_start` and `frame_end` options than the `frame_start` choice does. For example `RANGE BETWEEN CURRENT ROW AND offset PRECEDING` is not allowed, but `ROWS BETWEEN 7 PRECEDING AND 8 PRECEDING` is allowed, even though it would never select any rows.

**Example**

The example below is querying test scores for students in various grade levels. Unlike a regular `AVG()` aggregation, the window function allows us to see how each student individually compares to the average test score for their grade level, as well as compute the average test score while looking at different slices of the data for different grade levels – narrowing down the set of rows that constitutes the window using framing options such as PRECEDING or FOLLOWING.

```
SELECT
  first_name,
  grade_level,
  test_score,
  ROUND(AVG(test_score) OVER (PARTITION BY grade_level), 2) AS test_score_avg,
  ROUND(AVG(test_score) OVER (PARTITION BY grade_level ORDER BY test_score ASC ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING), 2) AS avg_1p_1f,
  ROUND(AVG(test_score) OVER (PARTITION BY grade_level ORDER BY test_score ASC ROWS BETWEEN 2 PRECEDING AND 2 FOLLOWING), 2) AS avg_2p_2f,
  ROUND(AVG(test_score) OVER (PARTITION BY grade_level ORDER BY test_score ASC ROWS BETWEEN UNBOUNDED PRECEDING AND 1 FOLLOWING), 2) AS avg_up_2f
FROM
  class_test
ORDER BY 
  grade_level, 
  test_score;
```

**Returns**:

```
' | first_name | grade_level | test_score | test_score_avg | avg_1p_1f | avg_2p_2f | avg_up_2f |
  +------------+-------------+------------+----------------+-----------+-----------+-----------+
' | Frank      | 9           | 76         | 81.33          | 77        | 77.67     | 77        |
' | Jojo       | 9           | 78         | 81.33          | 77.67     | 78.25     | 77.67     |
' | Iris       | 9           | 79         | 81.33          | 79        | 79.6      | 78.25     |
' | Peter      | 9           | 80         | 81.33          | 81.33     | 82.4      | 79.6      |
' | Sammy      | 9           | 85         | 81.33          | 85        | 83.5      | 81.33     |
' | Humphrey   | 9           | 90         | 81.33          | 87.5      | 85        | 81.33     |
' | Yolinda    | 10          | 30         | 68.2           | 44.5      | 55.67     | 44.5      |
' | Albert     | 10          | 59         | 68.2           | 55.67     | 63        | 55.67     |
' | Deborah    | 10          | 78         | 68.2           | 74        | 68.2      | 63        |
' | Mary       | 10          | 85         | 68.2           | 84        | 77.75     | 68.2      |
' | Shawn      | 10          | 89         | 68.2           | 87        | 84        | 68.2      |
' | Carol      | 11          | 52         | 73             | 60        | 64.33     | 60        |
' | Larry      | 11          | 68         | 73             | 64.33     | 67        | 64.33     |
' | Wanda      | 11          | 73         | 73             | 72        | 68.8      | 67        |
' | Otis       | 11          | 75         | 73             | 74.67     | 77.2      | 68.8      |
' | Shangxiu   | 11          | 76         | 73             | 81.67     | 79.5      | 73        |
' | Roseanna   | 11          | 94         | 73             | 85        | 81.67     | 73        |
' | Thomas     | 12          | 66         | 89             | 77.5      | 82.33     | 77.5      |
' | Jesse      | 12          | 89         | 89             | 82.33     | 85        | 82.33     |
' | Brunhilda  | 12          | 92         | 89             | 91.33     | 86.8      | 85        |
' | Charles    | 12          | 93         | 89             | 93        | 93.6      | 86.8      |
' | Franco     | 12          | 94         | 89             | 95.67     | 94.75     | 89        |
' | Gary       | 12          | 100        | 89             | 97        | 95.67     | 89        |
  +------------+-------------+------------+----------------+-----------+-----------+-----------+
```

* * *

- [AVG OVER](/sql_reference/functions-reference/window/avg-window.html)
- [COUNT OVER](/sql_reference/functions-reference/window/count-window.html)
- [DENSE\_RANK OVER](/sql_reference/functions-reference/window/dense-rank.html)
- [FIRST\_VALUE](/sql_reference/functions-reference/window/first-value.html)
- [LAG OVER](/sql_reference/functions-reference/window/lag.html)
- [LEAD OVER](/sql_reference/functions-reference/window/lead.html)
- [MAX OVER](/sql_reference/functions-reference/window/max-window.html)
- [MIN OVER](/sql_reference/functions-reference/window/min-window.html)
- [NTILE OVER](/sql_reference/functions-reference/window/ntile.html)
- [PERCENT\_RANK OVER](/sql_reference/functions-reference/window/percent-rank.html)
- [RANK OVER](/sql_reference/functions-reference/window/rank.html)
- [ROW\_NUMBER OVER](/sql_reference/functions-reference/window/row-number.html)
- [SUM OVER](/sql_reference/functions-reference/window/sum-window.html)