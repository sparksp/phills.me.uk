+++
date = "2007-08-18T12:48:00Z"
modified = "2011-11-17T12:48:00Z"
title = "tvrenamer.rc"
aliases = ["/snips/1-tvrenamerrc", "/pb/01"]
tags = ["tvrenamer.rc"]
[menu.main]
  parent = "snip"
+++
{{< highlight text >}}
--include_series
--cache
--pad=2
--detailed
--gap= -
--nogroup
--postproc=s/\((\d+)\)/(Part $1)/g;s/, Part (\d+)/ (Part $1)/g;
{{< /highlight >}}
