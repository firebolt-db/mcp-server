# [](#work-with-arrays)Work with arrays

This section covers querying and manipulating arrays in Firebolt.

- [Declaring ARRAY data types in table definitions](#declaring-array-data-types-in-table-definitions)
- [Simple array functions](#simple-array-functions)
- [Manipulating arrays with Lambda functions](#manipulating-arrays-with-lambda-functions)
  
  - [Lambda function general syntax](#lambda-function-general-syntax)
  - [Lambda function example–single array](#lambda-function-examplesingle-array)
  - [Lambda function example–multiple arrays](#lambda-function-examplemultiple-arrays)
- [UNNEST](#unnest)
  
  - [Example–single UNNEST with single ARRAY-typed column](#examplesingle-unnest-with-single-array-typed-column)
  - [Example–single UNNEST using multiple ARRAY-typed columns](#examplesingle-unnest-using-multiple-array-typed-columns)
  - [Example–multiple UNNEST clauses resulting in a Cartesian product](#examplemultiple-unnest-clauses-resulting-in-a-cartesian-product)
  - [Example–error on UNNEST of multiple arrays with different lengths](#exampleerror-on-unnest-of-multiple-arrays-with-different-lengths)
- [ARRAY input and output syntax](#array-input-and-output-syntax)
  
  - [Converting ARRAY to TEXT](#converting-array-to-text)
  - [Converting TEXT to ARRAY](#converting-text-to-array)
  - [Nested ARRAYs](#nested-arrays)

## [](#declaring-array-data-types-in-table-definitions)Declaring ARRAY data types in table definitions

Array types are declared using `ARRAY(<type>)` where `<type>` can be any data type that Firebolt supports. This includes the `ARRAY` data type, so arrays can be arbitrarily nested.

If you load an array from a CSV file, the arrays in the CSV file must be enclosed in double quotes (`""`).

For example, if a CSV file contains a row containing `value1 , value2 , "[array_value3 , array_value4]"` , you can create a table using the following code to read `array_value3` and `array_value4` into `array_values`.

```
CREATE TABLE IF NOT EXISTS array_example
    (  
      value1 STRING,  
      value2 STRING,  
      array_values ARRAY(TEXT)
    )
    URL = 's3://path_to_your_data/'
    TYPE = (csv);
```

Array literals are also supported. For example, the `SELECT` statement shown below is valid.

```
SELECT [1,2,3,4]
```

### [](#basis-for-examples)Basis for examples

All examples in this topic are based on the table below, named `visits`. The column `id` is of type `INTEGER`. All other columns are of type `ARRAY(TEXT)`.

![](/assets/images/array_example_table.png)

## [](#simple-array-functions)Simple array functions

There are several fundamental functions that you can use to work with arrays, including [ARRAY\_LENGTH](/sql_reference/functions-reference/array/array-length.html), [ARRAY\_CONCAT](/sql_reference/functions-reference/array/array-concat.html), and [ARRAY\_FLATTEN](/sql_reference/functions-reference/array/flatten.html). See the respective reference for a full description. Brief examples are shown below.

### [](#length-example)LENGTH example

`LENGTH` returns the number of elements in an array.

```
SELECT 
  id,
  LENGTH(agent_props_keys) as key_array_length
FROM visits;
```

**Returns**:

```
+-----------------------+
| id | key_array_length |
+-----------------------+
| 1  | 3                |
| 2  | 2                |
| 3  | 3                |
+-----------------------+
```

### [](#array_concat-example)ARRAY\_CONCAT example

`ARRAY_CONCAT` combines multiple arrays into a single array.

```
SELECT 
  id, 
  ARRAY_CONCAT(agent_props_keys, agent_props_vals) as concat_keys_and_vals
FROM visits;
```

**Returns**:

```
+----+------------------------------------------------------------------------------+
| id | concat_keys_and_vals                                                         |
+----+------------------------------------------------------------------------------+
| 1  | ["agent","platform","resolution","Mozilla/5.0","Windows NT 6.1","1024X4069"] |
| 2  | ["agent","platform","Safari","iOS 14"]                                       |
| 3  | ["agent","platform","platform","Safari","iOS 14","Windows 11"]               |
+----+------------------------------------------------------------------------------+
```

### [](#array_flatten-example)ARRAY\_FLATTEN example

`ARRAY_FLATTEN` converts an ARRAY of ARRAYs into a single flattened ARRAY. Note that this operation flattens only one level of the array hierarchy.

```
SELECT ARRAY_FLATTEN([ [[1,2,3],[4,5]], [[2]] ]) as flattened_array;
```

**Returns**:

```
+---------------------+
| flattened_array     |
+---------------------+
| [[1,2,3],[4,5],[2]] |
+---------------------+
```

## [](#manipulating-arrays-with-lambda-functions)Manipulating arrays with Lambda functions

Firebolt *Lambda functions* are a powerful tool that you can use on arrays to extract results. Lambda functions iteratively perform an operation on each element of one or more arrays. Arrays and the operation to perform are specified as arguments to the Lambda function.

### [](#lambda-function-general-syntax)Lambda function general syntax

The general syntax pattern of a Lambda function is shown below. For detailed syntax and examples see the reference topics for [Lambda functions](/sql_reference/functions-reference/Lambda/).

```
<LAMBDA_FUNC>(<arr1_var>[, <arr2_var>][, ...<arrN_var>]) -> <operation>, <array1>[, <array2>][, ...<arrayN>])
```

Parameter Description `<LAMBDA_FUNC>` Any array function that accepts a Lambda expression as an argument. For a list, see [Lambda functions](/sql_reference/functions-reference/Lambda/). `<arr1_var>[, <arr2_var>][, ...<arrN_var>]` A list of one or more variables that you specify. The list is specified in the same order and must be the same length as the list of array expressions (`<array1>[, <array2>][, ...<arrayN>]`). At runtime, each variable contains an element of the corresponding array. The specified `<operation>` is performed for each variable. The operation that is performed for each element of the array. This is typically a function or Boolean expression. \[, ]\[, ...] A comma-separated list of expressions, each of which evaluates to an `ARRAY` data type.

### [](#lambda-function-examplesingle-array)Lambda function example–single array

Consider the following [TRANSFORM](/sql_reference/functions-reference/Lambda/transform.html) array function that uses a single array variable and reference in the Lambda expression. This example applies the `UPPER` function to each element `t` in the `ARRAY`-typed column `tags`. This converts each element in each `tags` array to upper-case.

```
SELECT 
  id, 
  TRANSFORM(t -> UPPER(t), tags) AS up_tags
FROM visits;
```

**Returns:**

```
+----+--------------------------+
| id | up_tags                  |
+----+--------------------------+
| 1  | ["SUMMER-SALE","SPORTS"] |
| 2  | ["GADGETS","AUDIO"]      |
| 3  | ["SUMMER-SALE","AUDIO"]  |
+----+--------------------------+
```

### [](#lambda-function-examplemultiple-arrays)Lambda function example–multiple arrays

[ARRAY\_FIRST](/sql_reference/functions-reference/Lambda/array-first.html) is an example of a function that takes multiple arrays as arguments in a map of key-value pairs. One array represents the keys and the other represents the values.

`ARRAY_FIRST` uses a Boolean expression that you specify to find the key in the key array. If the Boolean expression resolves to true, the function returns the first value in the value array that corresponds to the key’s element position. If there are duplicate keys, only the first corresponding value is returned.

The example below returns the first value in the `agent_props_vals` array where the corresponding position in the `agent_props_keys` array contains the key `platform`.

```
SELECT 
  id, 
  ARRAY_FIRST(v, k -> k = 'platform', agent_props_vals, agent_props_keys) AS platform
FROM visits;
```

**Returns**:

```
+----+----------------+
| id | platform       |
+----+----------------+
| 1  | Windows NT 6.1 |
| 2  | iOS 14         |
| 3  | iOS 14         |
+----+----------------+
```

[ARRAY\_SORT](/sql_reference/functions-reference/array/array-sort.html) sorts one array by another. One array represents the values and the other represents the sort order.

The example below sorts the first array by the positions defined in the second array

```
SELECT 
  ARRAY_SORT(x,y -> y, [ 'A','B','C'],[3,2,1]) AS res;
```

**Returns**:

```
+-----------------+
| res             |
+-----------------+
| ["C", "B", "A"] |
+-----------------+
```

## [](#unnest)UNNEST

You might want to transform a nested array structure to a standard tabular format. `UNNEST` serves this purpose.

[UNNEST](/sql_reference/commands/queries/select.html#unnest) is a table-valued function (TVF) that transforms an input row containing an array into a set of rows. `UNNEST` unfolds the elements of the array and duplicates all other columns found in the `SELECT` clause for each array element. If the input array is empty, the corresponding row is eliminated.

You can use a single `UNNEST` command to unnest several arrays if the arrays are the same length.

Multiple `UNNEST` statements in a single `FROM` clause result in a Cartesian product. Each element in the first array has a record in the result set corresponding to each element in the second array.

### [](#examplesingle-unnest-with-single-array-typed-column)Example–single UNNEST with single ARRAY-typed column

The following example unnests the `tags` column from the `visits` table.

```
SELECT 
    id, 
    tag
FROM 
    visits,
    UNNEST(tags) as r(tag);
```

**Returns**:

```
+----+---------------+
| id |     tag       |
+----+---------------+
|  1 | "summer-sale" |
|  1 | "sports"      |
|  2 | "gadgets"     |
|  2 | "audio"       |
+----+---------------+
```

### [](#examplesingle-unnest-using-multiple-array-typed-columns)Example–single UNNEST using multiple ARRAY-typed columns

The following query specifies both the `agent_props_keys` and `agent_props_vals` columns to unnest.

```
SELECT
    id,
    a_key,
    a_val
FROM
    visits,
    UNNEST(agent_props_keys, agent_props_vals) as r(a_key, a_val);
```

**Returns**:

```
+----+------------+------------------+
| id | a_key      | a_val            |
+----+------------+------------------+
| 1  | agent      | “Mozilla/5.0”    |
| 1  | platform   | “Windows NT 6.1” |
| 1  | resolution | “1024x4069”      |
| 2  | agent      | “Safari”         |
| 2  | platform   | “iOS 14”         |
+----+------------+------------------+
```

### [](#examplemultiple-unnest-clauses-resulting-in-a-cartesian-product)Example–multiple UNNEST clauses resulting in a Cartesian product

The following query, while valid, creates a Cartesian product.

```
SELECT
    id,
    a_key,
    a_val
FROM
    visits,
    UNNEST(agent_props_keys as a_keys) as r1(a_key),
    UNNEST(agent_props_vals as a_vals) as r2(a_val);
```

**Returns**:

```
+-----+------------+------------------+
| id  | a_key      |   a_val          |
+-----+------------+------------------+
|   1 | agent      | "Mozilla/5.0"    |
|   1 | agent      | "Windows NT 6.1" |
|   1 | agent      | "1024x4069"      |
|   1 | platform   | "Mozilla/5.0"    |
|   1 | platform   | "Windows NT 6.1" |
|   1 | platform   | "1024x4069"      |
|   1 | resolution | "Mozilla/5.0"    |
|   1 | resolution | "Windows NT 6.1" |
|   1 | resolution | "1024x4069"      |
|   2 | agent      | "Safari"         |
|   2 | agent      | "iOS 14"         |
|   2 | platform   | "Safari"         |
|   2 | platform   | "iOS 14"         |
+-----+------------+------------------+
```

### [](#exampleerror-on-unnest-of-multiple-arrays-with-different-lengths)Example–error on UNNEST of multiple arrays with different lengths

The following query is **invalid** and will result in an error as the `tags` and `agent_props_keys` arrays have different lengths for row 1.

```
SELECT
    id,
    tag,
    a_key
FROM
    visits,
    UNNEST(tags, agent_props_keys) as r(tag, a_key);
```

## [](#array-input-and-output-syntax)ARRAY input and output syntax

`ARRAY` values can be converted from and to `TEXT`. This happens, for example, when an explicit `CAST` is added to a query, or when `ARRAY` values are (de-)serialized in a `COPY` statement.

### [](#converting-array-to-text)Converting ARRAY to TEXT

Broadly, the `TEXT` representation of an `ARRAY` value starts with an opening curly brace (`{`). This is followed by the `TEXT` representations of the individual array elements separated by commas (`,`). It ends with a closing curly brace (`}`). `NULL` array elements are represented by the literal string `NULL`. For example, the query

```
SELECT
    CAST([1,2,3,4,NULL] AS TEXT)
```

returns the `TEXT` value `'{1,2,3,4,NULL}'`.

When converting `ARRAY` values containing `TEXT` elements to `TEXT`, some additional rules apply. Specifically, array elements are enclosed by double quotes (`"`) in the following cases:

- The array element is an empty string.
- The array element contains curly or square braces (`{`,`[`,`]`,`}`), commas (`,`), double quotes (`"`), backslashes (`\`), or white space.
- The array element matches the word `NULL` (case-insensitively).

Additionally, double quotes and backslashes embedded in array elements will be backslash-escaped. For example, the query

```
SELECT
    CAST(['1','2','3','4',NULL,'','{impostor,array}','["impostor","array","back\slash"]',' padded and spaced ', 'only spaced', 'null'] AS TEXT)
```

returns the `TEXT` value `'{1,2,3,4,NULL,"","{impostor,array}","[\"impostor\",\"array\",\"back\\slash\"]"," padded and spaced ","only spaced","null"}'`.

### [](#converting-text-to-array)Converting TEXT to ARRAY

When converting the `TEXT` representation of an array back to `ARRAY`, the same quoting and escaping rules as above apply. Unquoted whitespace surrounding array elements is trimmed, but whitespace contained within array elements is preserved. The array elements themselves are converted according to the conversion rules for the requested array element type. For example, the query

```
SELECT
    CAST('{1, 2, 3, 4, null, "", "{impostor,array}", "[\"impostor\",\"array\",\"back\\slash\"]", " padded and spaced ", "null",   unquoted padded and spaced   }' AS ARRAY(TEXT))
```

returns the `ARRAY(TEXT)` value `[1,2,3,4,NULL,'','{impostor,array}','["impostor","array","back\slash"]',' padded and spaced ','null','unquoted padded and spaced']`.

It is also possible to enclose arrays with square braces (`[` and `]`) instead of curly braces (`{` and `}`) when converting from `TEXT` to `ARRAY`. For example, the query

```
SELECT
    CAST('[1, 2, 3, 4, NULL]' AS ARRAY(INTEGER))
```

returns the `ARRAY(INTEGER)` value `[1,2,3,4,NULL]`.

### [](#nested-arrays)Nested ARRAYs

Finally, the same prodedure applies when converting nested `ARRAY` values from and to `TEXT`. For example, the query

```
SELECT
    CAST([NULL,[],[NULL],[1,2],[3,4]] AS TEXT)
```

returns the `TEXT` value `{NULL,{},{1,2},{3,4}}`, and the query

```
SELECT
    CAST('{NULL,{},{1,2},{3,4}}' AS ARRAY(ARRAY(INTEGER)))
```

returns the `ARRAY(ARRAY(INTEGER))` value `[NULL,[],[NULL],[1,2],[3,4]]`.