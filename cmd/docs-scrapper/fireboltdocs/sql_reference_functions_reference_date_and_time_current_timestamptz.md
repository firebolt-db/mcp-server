# [](#current_timestamp)CURRENT\_TIMESTAMP

Returns the current timestamp as a `TIMESTAMPTZ` value.

**Alias:** `NOW`

## [](#syntax)Syntax

The function can be called with or without parentheses:

```
CURRENT_TIMESTAMP
CURRENT_TIMESTAMP()
```

## [](#return-type)Return Type

`TIMESTAMPTZ`

## [](#remarks)Remarks

The function gets the current timestamp from the system, and returns it as a `TIMESTAMPTZ` value.

## [](#example)Example

The following example assumes that the current Unix timestamp is `2023-03-03 14:42:31.123456 UTC`.

```
SET time_zone = 'Europe/Berlin';
SELECT CURRENT_TIMESTAMP;  --> 2023-03-03 15:42:31.123456+01

SET time_zone = 'America/New_York';
SELECT CURRENT_TIMESTAMP;  --> 2023-03-03 09:42:31.123456-05
```