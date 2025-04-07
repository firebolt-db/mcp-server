# [](#url_decode)URL\_DECODE

Decodes [percent-encoded](https://en.wikipedia.org/wiki/Percent-encoding) characters and replaces them with their binary value.

## [](#syntax)Syntax

```
URL_DECODE(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The string to be decoded. `TEXT`

## [](#return-type)Return Type

`TEXT`

## [](#example)Example

The example below decodes the percent-encoded parameters of an URL:

```
SELECT URL_DECODE('https://www.firebolt.io/?example_id%3D1%26hl%3Den');
```

**Returns**: https://www.firebolt.io/?example\_id=1&amp;hl=en

## [](#related)Related

- [URL\_ENCODE](/sql_reference/functions-reference/string/url_encode.html)