+++
date = "2007-08-18T12:48:00Z"
modified = "2011-11-17T12:48:00Z"
title = "tvrenamer.rc"
description = "My preferred settings for TVRenamer."
aliases = [
  "/snip/1*",
  "/snip/01*",
  "/snip/tvrenamerrc"
]
tags = ["tvrenamer.rc", "config", "code"]
categories = ["programming", "snip"]
license = "<a href=\"https://unlicense.org\">Unlicense</a>"
toc.enable = false
+++

My preferred settings for [TVRenamer](http://tvrenamer.org/).

```bash
--include_series
--cache
--pad=2
--detailed
--gap= -
--nogroup
--postproc=s/\((\d+)\)/(Part $1)/g;s/, Part (\d+)/ (Part $1)/g;
```
