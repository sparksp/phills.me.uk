+++
date = "2007-11-11T23:47:00Z"
modified = "2008-11-01T23:47:00Z"
title = "New Terminal"
description = "Launch a new terminal window, regardless of Terminal already running."
aliases = ["/snips/3-new-terminal", "/snips/3"]
tags = ["applescript"]
[menu]
  [menu.main]
    parent = "snip"
+++
{{< highlight applescript >}}
(*
New Terminal

Launch a new Terminal window, regardless of Terminal already running.

Author: Phill Sparks <me@phills.me.uk>
License: Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales
License URL: https://creativecommons.org/licenses/by-sa/2.0/uk/
*)
tell application "System Events"
    if (count (processes whose name is "Terminal")) is 0 then
        tell application "Terminal"
            activate
        end tell
    else
        tell application "Terminal"
            activate
            do script ""
        end tell
    end if
end tell
{{< /highlight >}}
