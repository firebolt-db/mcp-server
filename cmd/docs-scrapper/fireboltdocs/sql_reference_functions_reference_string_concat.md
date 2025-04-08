# [](#concat)CONCAT

Concatenates, i.e. combines, the text representations of all the input parameters without a separator, in the order they are provided.

## [](#syntax)Syntax

```
CONCAT( <expression>[, <expression>[, ...n]] );
```

**—OR—**

```
<expression> || <expression>
```

## [](#parameters)Parameters

### [](#concat-function)`CONCAT` function

Parameter Description Supported input types `<expression>[, ...n]` The expressions to be concatenated. Any type

The parameters to the `CONCAT` function can be of any data type, and will be converted to their text representation before concatenation. `NULL` parameters to the `CONCAT` function are treated as empty strings and ignored. If all parameters are `NULL`, the result will be an empty string.

### [](#-operator)`||` operator

Parameter Description Supported input types `<expression>` The expressions to be concatenated. Any non-array type, but at least one `TEXT`

To enable string concatenation, one parameter to the `||` operator must be of type `TEXT`, while the other parameter may be of any non-array data type. If one parameter to the `||` operator is `NULL`, the result will also be the non-null parameter; if both parameters are `NULL`, the result will be `NULL`.

The concatenation operator `||` can also be used for [array concatenation](/sql_reference/functions-reference/array/array-concat.html).

## [](#return-type)Return Type

`TEXT`

## [](#example)Example

The following example concatenates users’ `nicknames` and `emails` from the players table:

```
SELECT
	CONCAT(nickname, ': ', email) as user_info
FROM players
LIMIT 5;
```

**Returns**:

user\_info steven70: daniellegraham@example.net burchdenise: keith84@example.org stephanie86: zjenkins@example.org sabrina21: brianna65@example.org kennethpark: williamsdonna@example.com