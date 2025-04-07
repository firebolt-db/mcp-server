# [](#random)RANDOM

Returns a pseudo-random unsigned value greater than 0 and less than 1 of type `DOUBLE PRECISION`.

## [](#syntax)Syntax

```
RANDOM()
```

## [](#return-types)Return Types

`DOUBLE PRECISION`

## [](#examples)Examples

The following code example demonstrates using `RANDOM` without any other numeric functions. This generates a `DOUBLE PRECISION` value less than 1:

To create a random integer number between two values, you can use `RANDOM` with the `FLOOR` function. If `a` is the lesser value and `b` is the greater value, compute `FLOOR(RANDOM() * (b - a + 1)) + a`. The following code example generates a random integer between 50 and 100: