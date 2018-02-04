+++
date = "2007-12-07T23:35:00Z"
title = "MP3 Search"
description = "A quick JS snippet to search Google for MP3 files with the given name."
aliases = ["/snips/5-mp3-search", "/snips/5"]
tags = ["javascript", "code"]
categories = ["programming"]
[menu.main]
  parent = "snip"
+++
{{< highlight javascript >}}
/**
 * A quick JS snippet to search Google for MP3 files with the given name.
 *
 * @author Phill Sparks <me@phills.me.uk>
 */
if (t = prompt("Song Title"))
    window.location.href="http://google.co.uk/search?q=\"" + t +"\" last modified mp3 intitle:\"index of\" -html -htm -php -asp"
{{< /highlight >}}
