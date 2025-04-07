# [](#coalesce)COALESCE

Checks from left to right for the first non-NULL argument found for each entry parameter pair.

## [](#syntax)Syntax

```
COALESCE(<expression> [,...])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The expression(s) to coalesce. Any

## [](#return-types)Return Types

Same as input type

## [](#example)Example

The following example returns the first non-NULL value provided, which is the username `esimpson`:

```
SELECT COALESCE(NULL, 'esimpson','sabrina21') AS nicknames;
```

**Returns:** `esimpson`