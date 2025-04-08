# [](#string_to_array)STRING\_TO\_ARRAY

Splits a string into an array of strings based on a specified delimiter, with the following behaviors:

- If the delimiter is an empty string `''`, the result is an array containing the entire original input string as a single element.
- If the delimiter is `NULL`, the string is split into individual characters, with one character per array element.

## [](#syntax)Syntax

```
STRING_TO_ARRAY(<string>, <delimiter>)
```

## [](#parameters)Parameters

Parameter Description `<string>` The string to split. `<delimiter>` The separator to split the string by.

## [](#return-types)Return Types

`ARRAY(TEXT)`

## [](#examples)Examples

**Example**

The following code example splits the string `stephen70|esimpson|ruthgill|` at each `|` character and returns the resulting array as `nicknames`:

```
SELECT STRING_TO_ARRAY('stephen70|esimpson|ruthgill|', '|') AS nicknames;
```

**Returns**

nicknames (ARRAY(TEXT)) {stephen70,esimpson,ruthgill,””}

**Example**

The following code example calls `STRING_TO_ARRAY` with an empty delimiter, producing an array containing a single element which contains the input text:

```
SELECT STRING_TO_ARRAY('firebolt', '') as size_one_array;
```

**Returns**

size\_one\_array (ARRAY(TEXT)) {firebolt}

**Example**

The following example calls `STRING_TO_ARRAY` with `NULL` as the delimiter, splitting the text into individual characters:

```
SELECT STRING_TO_ARRAY('firebolt', NULL) AS single_characters;
```

**Returns**

single\_characters (ARRAY(TEXT)) {f,i,r,e,b,o,l,t}