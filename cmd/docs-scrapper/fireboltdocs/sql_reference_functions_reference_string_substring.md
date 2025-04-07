## [](#substring-substr)SUBSTRING, SUBSTR

Returns a substring starting at the character indicated by the `<start>` index and including the number of characters defined by `<length>`. Character indexing starts at index 1.

## [](#syntax)Syntax

```
SUBSTRING(<expression> [FROM <start>] [FOR <length>])
SUBSTRING(<expression>, <start> [, <length>])
SUBSTR(<expression>, <start> [, <length>])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The input string. `TEXT`,`BYTEA` `<start>` Optional. The starting position for the substring. 1 is the first character. If value &lt; 1 then it start from negative position (all negative positions treated as empty). `BIGINT` `<length>` Optional. The number of characters to be returned by the `SUBSTRING` function. Must be positive or 0. If not specified, `count` by default returns all of the string not specified by the `start` parameter. `BIGINT`

At least one of the optional arguments `<start>` and `<length>` must be given.

Negative `<start>` values begin the substring before the input `<expression>`. Positions before the input `<expression>` are considered empty.

## [](#return-type)Return Type

`TEXT` or `BYTEA`. The type of `<expression>`

##### [](#example)Example

In the example below, the string is offset by 1 and so the `SUBSTRING` command begins at the first letter, “h”. The `<length>` of 5 indicates the resulting string should be only five characters long.

```
SELECT
	SUBSTRING('hello world' FROM 1 FOR 5);
```

**Returns**: `hello`

In this next example `<length>` is not provided. This means all characters are included after the `<start>` index, which is 7.

```
SELECT
	SUBSTRING('hello world' FROM 7);
```

**Returns**: `world`

For negative `<start>` values, empty positions before the input string are selected. Thus, without `<length>` negative values will select the entire input `<string>`.

```
SELECT
	SUBSTRING('hello world' FROM -2);
```

**Returns**: `hello world`

With a negative `<start>` value and `<length>`, the empty positions before the string are counted towards the substring length. In this example, beginning the substring at `-2` for `8` characters includes 3 empty positions and 5 positions with characters from the input `<expression>`.

```
SELECT
	SUBSTRING('hello world' FROM -2 FOR 8);
```

**Returns**: `hello`