# [](#ilike)ILIKE

Allows matching of strings based on comparison to a pattern. `ILIKE` is normally used as part of a `WHERE` clause. `ILIKE` is case-insensitive; use [LIKE](/sql_reference/functions-reference/string/like.html) for case-sensitive pattern matching. Note that Firebolt uses the `POSIX` locale, which means that it only classifies the ASCII letters “A” through “Z” and “a” through “z” as letters.

## [](#syntax)Syntax

```
<expression> ILIKE '<pattern>'
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` Any expression that evaluates to `TEXT` `TEXT` `<pattern>` Specifies the pattern to match (case-insensitive). `TEXT` constant. SQL wildcards are supported:

\* Use an underscore (`_`) to match any single character  
\* Use a percent sign (`%`) to match any number of any characters, including no characters.

## [](#return-type)Return Type

`BOOLEAN`

## [](#example)Example

Find nicknames from the `players` table that partially match the string “Joe” and any following characters as follows:

```
SELECT
	playerid, nickname, email
FROM
	players
WHERE
	nickname ILIKE 'Joe%';
```

**Returns**:

playerid nickname email 160 joedavis cgarcia@example.org 519 joe79 jennifer10@example.net 3692 joeli cperez@example.net 3891 joel11 joanncain@example.net 4233 joellong millerholly@example.net 4627 joebowen amandalewis@example.net

## [](#unicode-behavior)Unicode Behavior

Firebolt uses the `POSIX` locale, therefore `ILIKE` case insensitivity is limited to ASCII characters. The uppercase and lowercase versions of non-ASCII characters are not matched:

```
SELECT 'ENCYCLOPÆDIA' ILIKE 'encyclopædia'; -- returns false
SELECT 'ENCYCLOPÆDIA' ILIKE 'encyclopÆdia'; -- returns true
SELECT 'MÜNCHEN' ILIKE 'München'; -- returns false
SELECT 'MÜNCHEN' ILIKE 'mÜnchen'; -- returns true
SELECT 'Πσρ⋈' ILIKE 'πΣΡ%'; -- returns false
```