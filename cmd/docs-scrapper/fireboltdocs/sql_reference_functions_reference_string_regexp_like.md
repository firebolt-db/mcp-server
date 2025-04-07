# [](#regexp_like)REGEXP\_LIKE

Checks whether a text pattern matches a regular expression string. Returns a `BOOLEAN` value, specifically `false` if the text doesn’t match and `true` if it does match. This is a RE2 regular expression.

## [](#syntax)Syntax

```
REGEXP_LIKE(<expression>, '<pattern>'[,'<flag>[...]'])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The text searched for a match using the RE2 pattern. `TEXT` `<pattern>` An RE2 regular expression pattern used to search for a match in the `<expression>`. [RE2 regular expression](https://github.com/google/re2/wiki/Syntax) `<flag>` Optional. Flags allow additional controls over the regular’s expression matching behavior. If using multiple flags, you can include them in the same single-quote block without any separator character. Firebolt supports the following RE2 flags to override default matching behavior. With `-` in front you can disable the flag.  
\* `i` - Specifies case-insensitive matching.  
\* `m` - Specifies multi-line mode. In this mode, `^` and `$` characters in the regex match the beginning and end of line.  
\* `s` - (Enabled per default) Specifies that the `.` metacharacter in regex matches the newline character in addition to any character in `.`  
\* `U` - Specifies non-greedy mode. In this mode, the meaning of the metacharacters `*` and `+` in regex `<pattern>` are swapped with `*?` and `+?`, respectively. See the examples using flags below for the difference in how results are returned.

## [](#return-type)Return Type

`BOOLEAN`

## [](#example)Example

```
SELECT
    REGEXP_LIKE('123','[a-z]');
```

**Returns**: `false`

```
SELECT
    REGEXP_LIKE('123','\\\\d+');
```

**Returns**: `true`

## [](#example-using-flags)Example using flags

The `i` flag causes the regular expression to be case-insensitive. Without this flag, this query would return `false` as no match is found.

```
SELECT
	REGEXP_LIKE('ABC', '[a-z]', 'i');
```

**Returns**: `true`