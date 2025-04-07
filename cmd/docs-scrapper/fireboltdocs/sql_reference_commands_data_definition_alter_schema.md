# [](#alter-schema)ALTER SCHEMA

Updates the specified SCHEMA.

## [](#alter-schema-owner-to)ALTER SCHEMA OWNER TO

Change the owner of a schema.  
The current owner of a view can be viewed in the `information_schema.schemata` view on `schema_owner` column.

check [ownership](/Guides/security/ownership.html) page for more info.

### [](#syntax)Syntax

```
ALTER SCHEMA <schema> OWNER TO <user>
```

### [](#parameters)Parameters

Parameter Description `<schema>` Name of the schema to change the owner of. `<user>` The new owner of the view.