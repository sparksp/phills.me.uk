+++
date = "2008-05-22T00:28:00Z"
modified = "2011-10-09T00:28:00Z"
title = "JS Page Navigation"
description = "Simple script to find first/prev/next/last links and navigate to them."
aliases = [
  "/snip/10*",
  "/snip/js-page-navigation"
]
tags = ["javascript", "code"]
categories = ["programming", "snip"]
+++

Simple script to find first/prev/next/last links and navigate to them.

```javascript
/**
 * Simple script to find first/prev/next/last links and navigate to them.
 *
 * @author Phill Sparks <me@phills.me.uk>
 * @license http://creativecommons.org/licenses/by-sa/2.0/uk/ Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales
 */
function go(rel) {
    var url;
    if (url = document.evaluate('//*[@rel="' + rel + '"]/href', document, null, XPathResult.STRING_TYPE, null)).stringValue) {
        window.location.href = url;
        return true;
    }
    return false;
}

function goFirst() {
    go('first') || alert('No first found.');
}
function goPrevious() {
    go('prev') || go('previous') || alert('No previous found.');
}
function goNext() {
    go('next') || alert('No next found.');
}
function goLast() {
    go('last') || alert('No last found.');
}
```
