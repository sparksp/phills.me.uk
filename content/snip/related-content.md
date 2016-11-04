+++
date = "2008-05-01T19:47:00Z"
modified = "2008-05-02T19:47:00Z"
title = "Related Content"
aliases = ["/snips/8-related-content", "/pb/08"]
tags = ["sql"]
[menu]
  [menu.main]
    parent = "snip"
+++
{{< highlight sql >}}
SELECT uri, COUNT(tag) AS related_tag_count
FROM milk_blog_tags prop
WHERE uri <> 'posts/2008/01/17/finding_the_character_map_in_mac'
AND tag IN ('software','Character Map','Character Palette','for:lsucs','osx')
GROUP BY uri
HAVING COUNT(tag) > 0
ORDER BY related_tag_count DESC, uri DESC
LIMIT 0, 5
{{< /highlight >}}
