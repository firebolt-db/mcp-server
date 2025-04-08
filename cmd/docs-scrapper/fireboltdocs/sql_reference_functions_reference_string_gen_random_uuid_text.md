# [](#gen_random_uuid_text)GEN\_RANDOM\_UUID\_TEXT

Returns a version 4 universally unique identifier (UUID) as defined by [RFC-4122](https://tools.ietf.org/html/rfc4122#section-4.4) as a `TEXT` value. This function accepts no arguments.

## [](#syntax)Syntax

```
GEN_RANDOM_UUID_TEXT()
```

## [](#return-type)Return Type

`GEN_RANDOM_UUID_TEXT` returns a result of type `TEXT`.

## [](#example)Example

The following code example generates a random UUID as a `TEXT` value:

```
SELECT gen_random_uuid_text() as id_col
```

id\_col (TEXT) ‘acc612ea-1a7d-4a49-b977-2486c7963d84’