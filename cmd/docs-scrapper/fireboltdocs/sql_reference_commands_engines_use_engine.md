# [](#use-engine)USE ENGINE

Changes the client to direct queries against the specified engine.

The client can be updated to direct queries against the system engine by using `system` for the `engine_name`.

Changing the engine used by the client will reset the user’s session.

## [](#syntax)Syntax

```
USE ENGINE <engine_name>
```

## [](#parameters)Parameters

Parameter Description `<engine_name>` The name of the engine to be used.

## [](#example-1)Example 1

The following example changes the client to use my\_engine:

```
USE ENGINE my_engine
```

## [](#example-2)Example 2

The following example changes the client to use system:

```
USE ENGINE system
```