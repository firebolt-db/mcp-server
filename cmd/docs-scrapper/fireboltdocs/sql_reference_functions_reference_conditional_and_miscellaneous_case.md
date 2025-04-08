# [](#case)CASE

Conditional expression similar to if-then-else statements. If the result of the condition is true, then the value of the CASE expression is the result that follows the condition, and the remainder of the CASE expression is not processed. If the result is not true, any subsequent WHEN clauses (conditions) are evaluated in the same manner. If no WHEN condition is true, then the value of the case expression is the result specified in the ELSE clause. If the ELSE clause is omitted and no condition matches, the result is NULL.

## [](#syntax)Syntax

```
CASE
    WHEN <condition> THEN <result>
    [ WHEN ...n ]
    [ ELSE <result> ]
END;
```

## [](#parameters)Parameters

Parameter Description Supported input types `<condition>` A condition can be defined for each `WHEN` clause. `BOOLEAN` `<result>` The result of the CASE expression when the preceding condition holds. Every `THEN` clause receives a single result. All results in a single `CASE` expression must share the same data type. Any

## [](#return-type)Return type

Same data type as `<result>`

## [](#example)Example

This example references a table `player_level` with the following columns and values:

player current\_level kennethpark 3 esimpson 8 sabrina21 11 rileyjon 15 burchdenise 4

The following example categorizes each player by level. If the level is higher than zero and less than 5 they are categorized as beginner. When the level is 6-12, they are categorized as intermediate, and when even higher, they are categorized as expert.

```
SELECT
	player,
	current_level,
	CASE
		WHEN current_level > 0 AND current_level <= 5 
			THEN 'Beginner'
		WHEN current_level > 5 AND current_level <= 12 
			THEN 'Intermediate'
		WHEN current_level > 12
			THEN 'Expert'
	END ranking
FROM
	player_level
ORDER BY
	player;
```

**Returns**:

player current\_level ranking burchdenise 4 Beginner esimpson 8 Intermediate kennethpark 3 Beginner rileyjon 15 Expert sabrina21 11 Intermediate