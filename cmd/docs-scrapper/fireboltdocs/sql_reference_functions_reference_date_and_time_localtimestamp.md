# [](#localtimestamp)LOCALTIMESTAMP

Returns the current local timestamp in the time zone specified in the session’s [`timezone` setting](/Reference/system-settings.html#setting-the-time-zone).

## [](#syntax)Syntax

The function can be called with or without parentheses:

```
LOCALTIMESTAMP
LOCALTIMESTAMP()
```

## [](#return-type)Return Type

`TIMESTAMP`

## [](#remarks)Remarks

The function gets the current timestamp from the system, converts it to the time zone specified in the `timezone` setting, and returns it as a `TIMESTAMP` value.

## [](#example)Example

The following example assumes that the current timestamp is `2023-03-03 14:42:31.123456 UTC`. Observe how it returns different `TIMESTAMP` values for different time zone settings:

```
SET timezone = 'Europe/Berlin';
SELECT LOCALTIMESTAMP;  --> 2023-03-03 15:42:31.123456

SET timezone = 'America/New_York';
SELECT LOCALTIMESTAMP;  --> 2023-03-03 09:42:31.123456
```