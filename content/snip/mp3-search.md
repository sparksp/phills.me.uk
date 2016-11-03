+++
date = "2007-12-07T23:35:00Z"
title = "MP3 Search"
aliases = ["/snips/5-mp3-search", "/pb/05"]
tags = ["javascript"]
+++
{{< highlight javascript >}}
if (t = prompt("Song Title"))
    window.location.href="http://google.co.uk/search?q=\"" + t +"\" last modified mp3 intitle:\"index of\" -html -htm -php -asp"
{{< /highlight >}}
