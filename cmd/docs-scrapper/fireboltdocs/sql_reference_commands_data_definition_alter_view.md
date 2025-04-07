# [](#alter-view)ALTER VIEW

Updates the specified VIEW.

## [](#alter-view-owner-to)ALTER VIEW OWNER TO

Change the owner of a view. The current owner of a view can be viewed in the `information_schema.views` view on `view_owner` column.

check [ownership](/Guides/security/ownership.html) page for more info.

### [](#syntax)Syntax

```
ALTER VIEW <view> OWNER TO <user>
```

### [](#parameters)Parameters

Parameter Description `<view>` Name of the view to change the owner of. `<user>` The new owner of the view.