+++
date = "2007-12-07T23:35:00Z"
title = "MP3 Search"
description = "A quick JS snippet to search Google for MP3 files with the given name."
aliases = [
  "/snip/5*",
  "/snip/05*",
  "/snip/mp3-search"
]
tags = ["javascript", "code"]
categories = ["programming", "snip"]
+++

A quick JS snippet to search Google for MP3 files with the given name.

```javascript
/**
 * A quick JS snippet to search Google for MP3 files with the given name.
 *
 * @author Phill Sparks <me@phills.me.uk>
 * @license http://creativecommons.org/licenses/by-sa/2.0/uk/ Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales
 */
if ((t = prompt("Song Title")))
  window.location.href =
    'http://google.co.uk/search?q="' +
    t +
    '" last modified mp3 intitle:"index of" -html -htm -php -asp';
```
