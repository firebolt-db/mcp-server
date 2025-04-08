# [](#drop-view)DROP VIEW

Deletes a view.

## [](#syntax)Syntax

```
DROP VIEW [IF EXISTS] <view_name> [CASCADE]
```

## [](#parameters)Parameters

Parameter Description `<view_name>` The name of the view to be deleted. `CASCADE` When specified, causes all dependent database objects such as views and aggregating indexes to be dropped also.