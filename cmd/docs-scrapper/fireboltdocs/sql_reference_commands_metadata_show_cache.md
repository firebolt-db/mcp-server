# [](#show-cache)SHOW CACHE

Returns the current SSD usage for the engine queried. `SHOW CACHE` returns values at the engine level, rather than by each node.

## [](#syntax)Syntax

```
SHOW CACHE;
```

The results of `SHOW CACHE` are formatted as follows:

`<ssd_used>/<ssd_available> GB (<percent_utilization>%)`

## [](#components)Components

Component Description `<ssd_used>` The amount of storage currently used on your engine. This data includes storage that Firebolt reserves for internal usage. `<ssd_available>` The amount of available storage on your engine. `<percent_utilization>` The percent of used storage as compared to available storage.

## [](#example)Example

When the `SHOW CACHE` command is run, the usage displays as a result:

ssd\_usage 3.82/73.28 GB (5.22%)