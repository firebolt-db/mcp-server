# [](#operators)Operators

Operators are represented by special characters or keywords, they do not use function call syntax.

## [](#operator-precedence)Operator Precedence

The following table lists all SQL operators from highest to lowest precedence, alongside with their associativity. The precedence and associativity of the operators is hard-coded into the parser. Add parentheses if you want an expression with multiple operators to be interpreted in some other way than what the precedence rules imply.

Operator Description Associativity `.` Table/column name separator left `::` Type cast left `[]` Array element left `+` `-` Unary plus and minus right `AT` `AT TIME ZONE` left `^` exponentiation left `*` `/` `%` multiplication, division, modulo left `+` `/` addition, subtraction left `\|\|` concatenation left `BETWEEN` range containment   `LIKE` `ILIKE` string matching   `<` `>` `=` `<=` `>=` `<>` comparison   `IS` `IS NULL`, `IS DISTINCT FROM`, `IS TRUE`   `NOT` logical negation right `AND` logical conjunction left `OR` logical disjunction left

- [Operator Precedence](#operator-precedence)
- [:: Type cast](#-type-cast)
- [Arithmetic (numbers)](#arithmetic-numbers)
- [Logical](#logical)
- [Comparison](#comparison)
- [CASE](#case)
  
  - [Return type](#return-type)
- [Date and time arithmetic](#date-and-time-arithmetic)
- [String](#string)
- [Subquery operators](#subquery-operators)

## [](#-type-cast):: Type cast

Use can use the `::` operator instead of the [CAST](/sql_reference/functions-reference/conditional-and-miscellaneous/cast.html) function to convert one [data type](/sql_reference/data-types.html) to another.

### [](#syntax)Syntax

```
 -- CAST function
 CAST(<value> AS <type>)
 -- :: operator
 <value>::<type>
```

Component Description `<value>` The value to convert or an expression that results in a value to convert. Can be a column name, a function applied to a column or another function, or a literal value. `<type>` The target [data type](/sql_reference/data-types.html) (case-insensitive).

### [](#example)Example

```
SELECT '2021-12-31'::DATE;
SELECT 8.5::REAL;
SELECT col_a::BIGINT;
```

## [](#arithmetic-numbers)Arithmetic (numbers)

Operator Operator description Input Data Types Output Data Types Example Result + addition INTEGER, BIGINT, NUMERIC, REAL, DOUBLE PRECISION Corresponding type `SELECT 2 + 3;` 5 - subtraction INTEGER, BIGINT, NUMERIC, REAL, DOUBLE PRECISION Corresponding type `SELECT 2 - 3;` -1 * multiplication INTEGER, BIGINT, NUMERIC, REAL, DOUBLE PRECISION Corresponding type `SELECT 2 * 3;` 6 / division (integer division truncates the result) INTEGER, BIGINT, NUMERIC, REAL, DOUBLE PRECISION Corresponding type `SELECT 4 / 2;` 2 % modulo (remainder) INTEGER, BIGINT Corresponding type `SELECT 5 % 4;` 1 ^ exponentiation INTEGER, BIGINT, NUMERIC, REAL, DOUBLE PRECISION DOUBLE PRECISION `SELECT 2 ^ 3.0;` 8

In arithmetic operations like +, -, \*, and / , the result’s data type aligns with the most encompassing type of the operands indicated as “ Corresponding type” in the table above. For clarity:

- When both operands are of the same data type (e.g., two INTEGERs or two NUMERICs), the result will also be of that same data type.
  
  - `INTEGER <op> INTEGER = INTEGER`
  - `INTEGER <op> BIGINT = BIGINT`
- For operations involving two different numeric data types, the result will typically be of the more precise or larger data type.
  
  - `INTEGER <op> REAL = DOUBLE PRECISION`
- Overflow checks and floating point errors are applied according to the result data type only.

> Floating point note:
> 
> Precision means that the representation of a number is accurate up to a certain number of digits. In Firebolt, `REAL` data types have 6-digit precision and `DOUBLE PRECISION` have 16-digit precision. This means that calculations have a precision of 6 or 16 respectively, and numbers are truncated to that precision. For example, if a number is stored as 1.234567, it is automatically truncated to 1.23456 for `REAL`.
> 
> When performing arithmetic, the number of leading digits in the output is the product of the leading digits in both inputs. This means that if either or both of the input numbers are larger than 6, those numbers are the first truncated, and then the arithmetic is performed.

## [](#logical)Logical

Logical operators return the result of a boolean operation using [three valued logic](https://en.wikipedia.org/wiki/Three-valued_logic)

Operator Example Explanation `AND` `x AND y` True if both x and y are true `NOT` `NOT x` True if x is false `OR` `x OR y` True if either x or y is true

## [](#comparison)Comparison

Operator Syntax Explanation = `a=b` a is equal to b. != `a!=b` a is not equal to b. &lt;&gt; `a<>b` a is not equal to b. &lt;= `a<=b` a is less than or equal to b. &gt; `a>b` a is greater than b. &gt;= `a>=b` a is greater than or equal to b. &lt; `a<b` a is less than b. BETWEEN `a BETWEEN b AND c` equivalent to b &lt;= a &lt;= c IS NULL `a IS NULL` a contains a NULL value IS NOT NULL `a IS NOT NULL` a doesn’t contain a NULL value IS NOT DISTINCT FROM `a IS NOT DISTINCT FROM b` equivalent to a=b where NULL is considered equal to NULL. IS DISTINCT FROM `a IS DISTINCT FROM b` equivalent to a!=b where NULL is considered equal to NULL.

Example of using comparison operator in `WHERE` clause

```
SELECT
  *
FROM
  Table
WHERE
  Price >= 100;
```

## [](#case)CASE

Conditional expression similar to if-then-else statements. If the result of the condition is true, then the value of the CASE expression is the result that follows the condition. If the result is false, any subsequent WHEN clauses (conditions) are searched in the same manner. If no WHEN condition is true, then the value of the case expression is the result specified in the ELSE clause. If the ELSE clause is omitted and no condition matches, the result is NULL.

### [](#syntax-1)Syntax

```
CASE
    WHEN <condition> THEN <result>
    [ WHEN ...n ]
    [ ELSE <result> ]
END;
```

### [](#parameters)Parameters

Parameter Description Supported input types `<condition>` A condition can be defined for each `WHEN`, and `ELSE` clause. `BOOLEAN` `<result>` The result of any condition. Every `THEN` clause receives a single result. All results in a single `CASE` function must share the same data type. Any

### [](#return-type)Return type

Same data type as `<result>`

### [](#example-1)Example

This example references a table `player_level` with the following columns and values:

player currentlevel kennethpark 3 esimpson 8 sabrina21 11 rileyjon 15 burchdenise 4

The following example categorizes each entry by length. If the movie is longer than zero minutes and less than 50 minutes it is categorized as SHORT. When the length is 50-120 minutes, it’s categorized as Medium, and when even longer, it’s categorized as Long.

```
SELECT
	player,
	currentlevel,
	CASE
		WHEN length > 0
		AND length <= 5 THEN 'Beginner'
		WHEN length > 5
		AND length <= 12 THEN 'Intermediate'
		WHEN length > 12 THEN 'Expert'
	END ranking
FROM
	player_level
ORDER BY
	player;
```

**Returns**:

player currentlevel ranking kennethpark 3 Beginner esimpson 8 Intermediate sabrina21 11 Intermediate rileyjon 15 Expert burchdenise 4 Beginner

## [](#date-and-time-arithmetic)Date and time arithmetic

Use the `INTERVAL` operator to add to or subtract from a period of time in `DATE`, `TIME`, or `TIMESTAMP` data types.

### [](#syntax-2)Syntax

```
{ +|- } INTERVAL '<quantity> [ <date_unit> ] [ ...]'
```

Component Description `<quantity>` An integer. Multiple `<quantities>` and `<date_units>` can be used in the same `INTERVAL` command if they are separated by spaces. `<date_unit>` A date measurement including any of the following: `millennium`, `century`, `decade`, `year`, `month`, `week`, `day`, `hour`, `minute`, `second`, `millisecond`, `microsecond`or their plural forms. If unspecified, `<date_unit>` defaults to `second`.

### [](#example-2)Example

```
<date_column> + INTERVAL '1 year 2 months 3 days'
<date_column> - INTERVAL '2 weeks'
<date_column> - INTERVAL '1 year 3 hours 20 minutes'
```

## [](#string)String

To concatenate strings, you can use the `CONCAT` function.

```
SELECT concat('This', ' is', ' a', ' parenthetical', 'concantenation.') AS concatenated_String
```

Alternatively, you can use the double pipe `||` operator.

```
SELECT 'This' || ' is' || ' a' || ' double pipe' || ' concantenation.' AS concatenated_String
```

## [](#subquery-operators)Subquery operators

Subqueries are queries nested within another query. They allow complex data retrieval by enabling a query to filter results based on the outcome of another query. Subquery operators are crucial in constructing these nested queries, especially within the WHERE clause, to filter data based on specific conditions.

Operator Explanation `EXISTS` The `EXISTS` operator is used to check for the existence of any record in a subquery. It returns TRUE if the subquery returns one or more records. The subquery within EXISTS is executed repeatedly, once for each row that might be selected by the outer query. If the subquery returns any row, the EXISTS condition is met, and the outer query processing continues for that row. `NOT EXISTS` It is opposite of EXISTS and is used to find records in one table that have no related records in another table. If the subquery returns no rows, NOT EXISTS returns TRUE. `IN` The `IN` operator checks if a specific value is present in a list of values or the results of a subquery. Commonly utilized in a `WHERE` clause, it compares a column’s value against a predefined set. When the column’s value matches any value in this set, `IN` yields `TRUE`. `NOT IN` Retrieve all entries from the value list that don’t match the required value.

### [](#exampleusing-exists-to-find-all-suppliers-with-products-equal-to-the-price-of-22)Example–using EXISTS to find all suppliers with products equal to the price of 22

```
SELECT supplier_name
FROM suppliers
WHERE EXISTS (
  SELECT
    product_name
  FROM
    products
  WHERE
    products.supplier_id = suppliers.supplier_id
  AND
    price = 22);
```

### [](#exampleusing-the-in-operator-to-return-all-customers-from-mannheim-or-london)Example–using the IN operator to return all customers from Mannheim or London

```
SELECT
  customer_name
FROM
  customers
WHERE
  customer_address IN ('Mannheim','London');
```

### [](#exampleusing-a-correlated-subquery-to-retrieve-all-the-products-that-cost-more-than-the-avgerage-price)Example–using a correlated subquery to retrieve all the products that cost more than the avgerage price

```
SELECT
  product_id,
  product_name,
  list_price
FROM
  products p
WHERE
  list_price > (
    SELECT
      AVG( list_price )
    FROM
      products
    WHERE
      category_id = p.category_id);
```

### [](#exampleusing-a-scalar-boolean-subquery-to-retrieve-rows-based-on-truefalse-condition)Example–using a scalar boolean subquery to retrieve rows based on true/false condition

```
SELECT
  *
FROM
  products
WHERE (
  SELECT CASE WHEN
    MIN(list_price) > 100
  THEN
    true
  ELSE
    false
  END
  FROM
    products);
```