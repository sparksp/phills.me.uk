+++
date = "2007-11-11T23:47:00Z"
modified = "2008-11-01T23:47:00Z"
title = "New Terminal"
aliases = ["/snips/3-new-terminal", "/snips/3"]
tags = ["applescript"]
[menu]
  [menu.main]
    parent = "snip"
+++
{{< highlight applescript >}}
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
