# [](#regexp_extract)REGEXP\_EXTRACT

Returns the first match of `<pattern>` within the `<expression>`. If the pattern does not match, returns `NULL`. If you want to extract all matches, use [REGEXP\_EXTRACT\_ALL](/sql_reference/functions-reference/string/regexp-extract-all.html).

## [](#syntax)Syntax

```
REGEXP_EXTRACT(<expression>, <pattern>[,'<flag>[...]',[<index>]])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The string from which to extract a substring, based on a regular expression. `TEXT` `<pattern` A [re2 regular expression](https://github.com/google/re2/wiki/Syntax) for matching with the string. `TEXT` `<flag>` Optional. Flag that allows additional controls over the regular’s expression matching behavior. If using multiple flags, you can include them in the same single-quote block without any separator character. Firebolt supports the following RE2 flags to override default matching behavior. With `-` in front, you can disable the flag.  
\* `i` - Specifies case-insensitive matching.  
\* `m` - Specifies multi-line mode. In this mode, `^` and `$` characters in the regex match the beginning and end of the line.  
\* `s` - (Enabled per default) Specifies that the `.` metacharacter in regex matches the newline character in addition to any character in `.`  
\* `U` - Specifies non-greedy mode. In this mode, the meaning of the metacharacters `*` and `+` in regex `<pattern>` are swapped with `*?` and `+?`, respectively. See the examples using flags below for the difference in how results are returned. `<index>` Optional. Indicates which subgroup of the expression match should be returned. The default value is `0` which means the whole match is returned, independent of any number of given subgroups. An `INTEGER` between `0` and `N` where `N` is the number subgroups in the `<pattern>`.

## [](#return-types)Return Types

`TEXT`

## [](#example)Example

```
SELECT
	REGEXP_EXTRACT('ABC 2024', '^[A-Z]+');
```

**Returns**: `"ABC"`

Despite using subgroups in the regular expression, the full match will be returned as the optional `<index>` argument is not set (the default value `0` is used instead).

```
SELECT
	REGEXP_EXTRACT('Learning about #REGEX in #Firebolt 2024', '#([A-Za-z]+) (\d+)');
```

**Returns**: `"#Firebolt 2024"`

The regular expression contains two subgroups which allows us to set the `<index>` argument to something between `0` and `2`. Every other value will cause an exception to be thrown. Setting `<index>` to `0` would cause the full match `"#Firebolt 2024"` to be returned (same behavior as not setting this value, see the example above), while a `2` would return the second subgroup `"2024"`.

```
SELECT
	REGEXP_EXTRACT('Learning about #REGEX in #Firebolt 2024', '#([a-z]+) (\d+)', 'i', 1);
```

**Returns**: `"Firebolt"`