+++
date = "2007-12-07T23:35:00Z"
title = "MP3 Search"
description = "A quick JS snippet to search Google for MP3 files with the given name."
aliases = ["/snip/5*", "/snip/05*"]
tags = ["javascript", "code"]
categories = ["programming"]
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
