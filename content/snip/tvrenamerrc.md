+++
date = "2007-08-18T12:48:00Z"
modified = "2011-11-17T12:48:00Z"
title = "tvrenamer.rc"
description = "My preferred settings for TVRenamer."
aliases = ["/snips/1-tvrenamerrc", "/snips/1"]
tags = ["tvrenamer.rc", "config", "code"]
categories = ["programming"]
[menu]
  [menu.main]
    parent = "snip"
+++

My preferred settings for [TVRenamer](http://tvrenamer.org/).

{{< highlight bash >}}
--include_series
--cache
--pad=2
--detailed
--gap= -
--nogroup
--postproc=s/\((\d+)\)/(Part $1)/g;s/, Part (\d+)/ (Part $1)/g;
{{< /highlight >}}
