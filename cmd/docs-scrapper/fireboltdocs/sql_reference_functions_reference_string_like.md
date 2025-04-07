# [](#like)LIKE

The LIKE is used for pattern matching to find similar strings in data. It’s often employed in the WHERE clause to filter results based on specific patterns. `LIKE` is case-sensitive; use [ILIKE](/sql_reference/functions-reference/string/ilike.html) for case-insensitive pattern matching.

## [](#syntax)Syntax

```
<expression> LIKE '<pattern>'
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` Any expression that evaluates to `TEXT`. `TEXT` `<pattern>` Specifies the pattern to match (case-sensitive). `TEXT` constant. SQL wildcards are supported:

\* Use an underscore (`_`) to match any single character  
\* Use a percent sign (`%`) to match any number of any characters, including no characters.

## [](#return-type)Return Type

`BOOLEAN`

## [](#example)Example

Find nicknames from the `players` table that partially match the string “joe” and any following characters as follows:

```
SELECT
	playerid, nickname, email
FROM
	players
WHERE
	nickname LIKE 'joe%';
```

**Returns**:

```
+----------+----------+-------------------------+
| playerid | nickname | email                   |
+----------+----------+-------------------------+
| 160      | joedavis | cgarcia@example.org     |
| 519 	   | joe79    | jennifer10@example.net  |
| 3692 	   | joeli    | cperez@example.net      |
| 3891	   | joel11   | joanncain@example.net   |
| 4233 	   | joellong | millerholly@example.net |
| 4627 	   | joebowen | amandalewis@example.net |
+----------+----------+-------------------------+
```

Note that the following would return no results, as `LIKE` matches case-sensitively, unlike `ILIKE`:

```
SELECT
	playerid, nickname, email
FROM
	players
WHERE
	nickname LIKE 'Joe%';
```

Find nicknames from the `players` table that have a specific pattern:

```
SELECT
	playerid, nickname, email
FROM
	players
WHERE
	nickname LIKE '%oe_e%';
```

The pattern %oe\_e% signifies any string that contains “oe” followed by any character (_ acts as a wildcard for a single character) and then “e”.

**Returns**

```
+----------+----------+-------------------------+
| playerid | nickname | email                   |
+----------+----------+-------------------------+
| 3891     | joel11    | joanncain@example.net  |
| 4233     | joellong  | millerholly@example.net|
+----------+----------+-------------------------+
```