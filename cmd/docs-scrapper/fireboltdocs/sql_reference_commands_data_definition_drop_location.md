# [](#drop-location)DROP LOCATION

Removes a location object from your Firebolt account. Be cautious when dropping a location that is referenced by external tables, as it may impact dependent objects.

**Topics:**

- [Syntax](#syntax)
- [Parameters](#parameters)
- [Examples](#examples)
- [Error handling](#error-handling)
- [Best practices](#best-practices)
- [Notes](#notes)

## [](#syntax)Syntax

```
DROP LOCATION [IF EXISTS] <location_name> [WITH FORCE]
```

## [](#parameters)Parameters

Parameter Description `<location_name>` The identifier of the location to remove. `IF EXISTS` A parameter that suppresses the error if the location doesn’t exist. `WITH FORCE` Forces the removal of the location even if it has dependent external tables. **Warning**: This option invalidates all dependent external tables.

## [](#examples)Examples

**Drop a location**

The following code example removes a `LOCATION` object from your Firebolt account:

```
DROP LOCATION my_location
```

**Drop IF EXISTS**

The following code example removes a `LOCATION` object from your Firebolt account if it exists:

```
DROP LOCATION IF EXISTS my_location
```

**Drop with FORCE Option**

If you attempt to drop a location that is referenced by an external table, the drop operation fails by default. You can do either of the following:

1. **Safe Approach**: Drop all dependent external tables before dropping the location.
2. **Force Approach**: Use the `WITH FORCE` option to drop the location immediately regardless of dependencies.

If you choose `WITH FORCE`, the following apply:

- The location is deleted regardless of existing dependencies.
- All external tables that reference the removed location become invalid.
- Queries on any invalid tables will result in an error.
- You must manually clean up the invalid external tables after using `WITH FORCE`.

The following code example deletes a `LOCATION` object from your Firebolt account without considering its dependencies:

```
DROP LOCATION my_location WITH FORCE
```

## [](#error-handling)Error handling

The following table lists the types of errors that may occur during location removal:

Error Type Cause Resolution **Location Not Found or Unauthorized** The specified location does not exist or you lack permission to access it. Use `IF EXISTS` to handle non-existent locations without error. **Dependent Objects Exist** The location is referenced by one or more external tables. Drop the dependent tables first or use `WITH FORCE` to bypass dependency checks. **Invalid External Table** Attempting to query an external table after its location was removed using `WITH FORCE`. Drop the invalid external table.

## [](#best-practices)Best Practices

When dropping a location, consider the following:

1. **Before Dropping**
   
   - Use `IF EXISTS` to handle non-existent locations gracefully.
   - Ensure no active queries are using the location.
   - Check for dependent external tables.

To identify external tables that depend on this location, execute the following query.

```
SELECT * FROM information_schema.tables WHERE location_name = 'my_location_name'
```

1. **Using WITH FORCE**
   
   - Use only when necessary.
   - Clean up invalid external tables immediately.
   - Document the reason for using this option.
2. **After Dropping**
   
   - Verify all dependent objects are handled.
   - Update any scripts or documentation referencing the location.

## [](#notes)Notes

- All identifiers are case-insensitive unless enclosed in double-quotes
- For more information about object identifiers, see [Object identifiers](/Reference/object-identifiers.html)