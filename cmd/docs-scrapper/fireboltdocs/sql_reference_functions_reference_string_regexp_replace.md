# [](#regexp_replace)REGEXP\_REPLACE

Matches a pattern in the input string and replaces the first matched portion (from the left) with the specified replacement.

## [](#syntax)Syntax

```
REGEXP_REPLACE(<input>, <pattern>, <replacement>)
```

# [](#regexp_replace_all)REGEXP\_REPLACE\_ALL

Matches a pattern in the input string and replaces all matched portions with the specified replacement.

## [](#syntax-1)Syntax

```
REGEXP_REPLACE_ALL(<input>, <pattern>, <replacement>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<input>` The string to search for a matching pattern `TEXT` `<pattern>` An [RE2 regular expression](https://github.com/google/re2/wiki/Syntax) for matching with the string input. `TEXT` `<replacement>` The string to replace the matching pattern found in the input. This argument can include the following special sequences:  
\* `\&` - To indicate that the substring matching the entire pattern should be inserted.  
\* `\n` - Where *n* is a digit from 1 to 9, to indicate that the substring matching the n’th capturing group (parenthesized subexpression) of the pattern should be inserted. If pattern doesn’t have *n* capturing groups, the `\n` is ignored.  
\* `\\` - results in a single &lt;br&gt;* `\c` - Specifies for any other character, c results in the same sequence \\c  
Note, that for string literals the above escaping rules apply *after* string literals escaping rules for `\`. See examples below. `TEXT`

If any of the arguments to these functions is `NULL`, the return value is `NULL`.

### [](#return-type)Return Type

`TEXT`

## [](#examples)Examples

Replace first occurence of `!` with `!!!`

```
SELECT REGEXP_REPLACE('Hello, world!', '!', '!!!');
```

**Returns**: `'Hello, world!!!'`

Remove leading and trailing spaces

```
SELECT REGEXP_REPLACE_ALL('     Hello world ! ', '^[ ]+|[ ]+$', '');
```

**Returns**: `'Hello world !'`

Duplicate every character

```
SELECT REGEXP_REPLACE_ALL('Hello, World!', '.', '\&\&')
```

**Returns**: `'HHeelllloo,, WWoorrlldd!!'`

Mask email address by leaving first character only (Note: this is for illustrative purposes only, the email matching pattern is too simplistic)

```
SELECT REGEXP_REPLACE(email, '(\w)[\w\.]+@([\w]+\.)+([\w]+)', '\1***@\2\3')
FROM UNNEST([
  'matt123@hotmail.com',
  'joe.doe@gmail.com',
  '12345@www.atg.wa.gov'
]) email
```

**Returns**:

```
'm***@hotmail.com'
'j***@gmail.com'
'1***@www.atg.wa.gov'
```

Convert dates into US format

```
SELECT REGEXP_REPLACE(event_date::TEXT, '(\d{4})-(\d{2})-(\d{2})', '\2/\3/\1')
FROM UNNEST([
  DATE '1970-08-07',
  DATE '2000-04-22',
  DATE '2002-07-25',
  DATE '2010-11-11'
]) event_date
```

**Returns**

```
08/07/1970
04/22/2000
07/25/2002
11/11/2010
```