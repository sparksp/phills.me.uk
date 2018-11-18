+++
date = "2008-05-01T19:47:00Z"
modified = "2008-05-02T19:47:00Z"
title = "Related Content"
description = "SQL to hunt down the 5 most relevant pieces of content by tag."
aliases = ["/snips/8-related-content", "/snips/8"]
tags = ["sql", "code"]
categories = ["programming"]
+++
{{< highlight sql >}}
/**
 * SQL to hunt down the 5 most relevant pieces of content by tag.  This schema
 * uses `uri` as the primary key for content (a foreign key in the tags table).
 *
 * @author Phill Sparks <me@phills.me.uk>
 * @license http://creativecommons.org/licenses/by-sa/2.0/uk/ Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales
 */
SELECT uri, COUNT(tag) AS related_tag_count
FROM milk_blog_tags prop
WHERE uri <> 'posts/2008/01/17/finding_the_character_map_in_mac'
AND tag IN ('software','Character Map','Character Palette','for:lsucs','osx')
GROUP BY uri
HAVING COUNT(tag) > 0
ORDER BY related_tag_count DESC, uri DESC
LIMIT 0, 5
{{< /highlight >}}
