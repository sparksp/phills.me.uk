+++
date = "2008-08-26T12:55:00Z"
slug = "curl_error_26"
title = "cURL Error 26: Failed to open/read local data from file/application"
tags = ["php", "curl", "error", "web design", "code"]
aliases = ["/blog/2008/08/26/curl_error_26"]
[menu]
  [menu.main]
    parent = "post"
+++

Last night I started working on Lilly's 365 and very quickly ran into a few problems.  Most were simple quirks of [site5](http://www.site5.com/in.php?id=51960) fixed by getting the permissions right.  The thing that kept bugging me was Phlickr, and more specifically its use of [cURL](http://curl.haxx.se/).

[Phlickr](http://phlickr.sourceforge.net) is an open source PHP5 based API kit used to access the Flickr API.  It uses the [REST](http://en.wikipedia.org/wiki/REST) method which requires that some actions are POSTs.  The cURL functions in PHP can happily handle this but for some reason they were failing every time... "Failed to open/read local data from file/application" (Error 26).  As soon as I disabled the `CURLOPT_POST` (making a GET request) cURL made no errors, though the Flickr API didn't like it.

After much searching later I found the solution was very simple...  If you're not submitting any post fields then set `CURLOPT_POSTFIELDS` to blank, like this...

{{< highlight php >}}
<?php
$ch = curl_init($url);
// make sure we submit this as a post
curl_setopt($ch, CURLOPT_POST, true);
if (isset($postParams)) {
    curl_setopt($ch, CURLOPT_POSTFIELDS, $postParams);
}
else {
    curl_setopt($ch, CURLOPT_POSTFIELDS, '');
}
{{< /highlight >}}

I can see that this makes sense; a POST is used for submitting data to a server so you need to submit something, even if it's blank.  It would be nice if the error message was a bit more useful or better documented.
