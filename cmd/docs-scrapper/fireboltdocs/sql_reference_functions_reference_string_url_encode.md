# [](#url_encode)URL\_ENCODE

Encodes all characters that are not unreserved using [percent-encoding](https://en.wikipedia.org/wiki/Percent-encoding).

Unreserved characters are defined according to [W3C RFC 3986](https://www.rfc-editor.org/rfc/rfc3986.html).

```
unreserved  = ALPHA / DIGIT / "-" / "." / "_" / "~"
```

## [](#syntax)Syntax

```
URL_ENCODE(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The string to be encoded. `TEXT`

## [](#return-type)Return Type

`TEXT`

## [](#example)Example

The example below converts characters that are not unreserved that appear in the parameter section of the URL:

```
SELECT CONCAT('https://www.firebolt.io/?', URL_ENCODE('example_id=1&hl=en'));
```

**Returns**: https://www.firebolt.io/?example\_id%3D1%26hl%3Den

## [](#related)Related

- [URL\_DECODE](/sql_reference/functions-reference/string/url_decode.html)