# [](#stop-engine)STOP ENGINE

Stops a running engine.

## [](#syntax)Syntax

```
STOP ENGINE [IF EXISTS] <engine_name>
[WITH 
    [TERMINATE = <true/false>]]
```

## [](#parameters)Parameters

Parameter Description `<engine_name>` The name of the engine to be stopped. `<TERMINATE>` When `false`, the engine will wait for running queries to finish before stopping.  
When `true`, the engine will stop without waiting for running queries to complete.

If not specified, `false` is used as default.

**Limitations:**

- When `TERMINATE=false` (default), the engine will wait for up to 24 hours for running queries to finish before stopping. After 24 hours, the engine will finish shutting down. Any queries still running may not run to completion.

If you would like to remove any of these limitations, reach out to Firebolt Support.

## [](#example-1)Example 1

The following example waits for queries on my\_engine to finish before stopping:

```
STOP ENGINE my_engine
```

## [](#example-2)Example 2

The following example stops my\_engine without waiting for running queries to finish:

```
STOP ENGINE my_engine WITH TERMINATE=true
```