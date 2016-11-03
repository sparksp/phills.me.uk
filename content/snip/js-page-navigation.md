+++
date = "2008-05-22T00:28:00Z"
modified = "2011-10-09T00:28:00Z"
title = "JS Page Navigation"
aliases = ["/snips/10-js-page-navigation", "/pb/10"]
tags = ["javascript"]
+++
{{< highlight javascript >}}
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
{{< /highlight }}
