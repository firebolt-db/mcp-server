# [](#date_diff)DATE\_DIFF

Calculates the difference between `start_timestamp` and `end_timestamp` by the indicated unit.

## [](#syntax)Syntax

```
DATE_DIFF('<unit>', <start_timestamp>, <end_timestamp>)
```

## [](#parameters)Parameters

Parameter Description `<unit>` A TEXT literal specifying the time unit. Must be one of `microsecond`, `millisecond`, `second`, `minute`, `hour`, `day`, `week`, `month`, `quarter`, `year`, `decade`, `century`, or `millennium`. `<start_timestamp>` A value expression evaluating to a `TIMESTAMP` or `TIMESTAMPTZ` value. `<end_timestamp>` A value expression evaluating to a `TIMESTAMP` or `TIMESTAMPTZ` value.

## [](#return-type)Return Type

`BIGINT`

## [](#example)Example

```
SELECT DATE_DIFF('day', '2024-01-01'::TIMESTAMP, '2024-04-15'::TIMESTAMP);
```

**Returns**

105