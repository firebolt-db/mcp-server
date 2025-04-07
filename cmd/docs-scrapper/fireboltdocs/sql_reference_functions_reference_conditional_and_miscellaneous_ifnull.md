# [](#ifnull)IFNULL

Compares two expressions. Returns `<expression1>` if it’s non-NULL, otherwise returns `<expression2>`.

## [](#syntax)Syntax

```
IFNULL(<expression1>, <expression2>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression1>`, `<expression2>` Expressions that evaluate to any data type that Firebolt supports. Any

## [](#return-types)Return Types

Same as input type

## [](#remarks)Remarks

Use `ZEROIFNULL(<expression>)` as a synonym shorthand for `IFNULL(<expression>, 0)`. `IFNULL(a, b)` has the same behaviour as `COALESCE(a, b)`.

## [](#example)Example

The following truth table demonstrates values that `IFNULL` returns based on the values of two columns: `level` and `player_id`:

```
SELECT level, player_id, IFNULL(level, player_id), IFNULL(player_id, level)
FROM players;
```

level player\_id IFNULL(level,player\_id) IFNULL(player\_id,level) 0 32 0 32 1 null 30 30 null 33 33 33 null null null null