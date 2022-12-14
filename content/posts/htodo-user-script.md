+++
date = "2009-05-29T13:05:00Z"
modified = "2010-07-13T13:05:00Z"
title = "hTodo User Script"
description = "Adds hTodo support to the operator toolbar."
aliases = [
  "/snip/13*",
  "/snip/htodo-user-script"
]
tags = ["javascript", "code"]
categories = ["programming", "snip"]
+++

Adds hTodo support to the operator toolbar.

```javascript
/**
 * Adds hTodo support to the operator toolbar.
 *
 * @author Phill Sparks <me@phills.me.uk>
 * @license http://creativecommons.org/licenses/by-sa/2.0/uk/ Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales
 * @version 1.1, 2010-01-11
 * @see http://microformats.org/wiki/htodo
 * @see http://www.kaply.com/weblog/operator/
 */

function hTodo(node) {
  if (node) {
    Microformats.parser.newMicroformat(this, node, "hTodo");
  }
}

hTodo.prototype.toString = function () {
  return this.summary;
};

var hTodo_definition = {
  mfVersion: 0.8,
  description: "Tasks",
  mfObject: hTodo,
  className: "vtodo",
  required: ["summary"],
  properties: {
    attach: {
      plural: true,
      datatype: "anyURI",
    },
    category: {
      plural: true,
      datatype: "microformat",
      microformat: "tag",
      microformat_property: "tag",
    },
    class: {
      values: ["public", "private", "confidential"],
    },
    completed: {
      datatype: "dateTime",
    },
    description: {
      datatype: "HTML",
    },
    dtstart: {
      datatype: "dateTime",
    },
    dtstamp: {
      datatype: "dateTime",
    },
    due: {
      datatype: "dateTime",
    },
    duration: {},
    geo: {
      datatype: "microformat",
      microformat: "geo",
    },
    "last-modified": {
      datatype: "dateTime",
    },
    location: {
      datatype: "microformat",
      microformat: "hCard",
    },
    organizer: {
      datatype: "microformat",
      microformat: "hCard",
    },
    "percent-completed": {
      /* datatype: 'int', min: 0, max: 100 */
    },
    priority: {
      /* datatype: 'int', min: 0, max: 9 */
    },
    summary: {},
    status: {
      values: ["active", "in-process", "completed", "cancelled"],
    },
    uid: {
      datatype: "anyURI",
    },
    url: {
      datatype: "anyURI",
    },

    /* Borrow this from the hCalendar_definition */
    rrule: hCalendar_definition.properties.rrule,
  },
};

Microformats.add("hTodo", hTodo_definition);

/* Firefox Bookmark */
var hTodo_firefox_bookmark = {
  scope: {
    semantic: {
      hTodo: "url",
    },
  },
  doAction: function (semanticObject, semanticObjectType) {
    if (semanticObjectType == "hTodo") {
      name = semanticObject["summary"];
      url = semanticObject["url"];
      SemanticActions.firefox_bookmark.bookmark(name, url);
      return true;
    }
  },
};

SemanticActions.add("firefox_bookmark", hTodo_firefox_bookmark);

/* Delicious Bookmark (v5) */
var hTodo_delicious_bookmark = {
  scope: {
    semantic: {
      hTodo: "url",
    },
  },
  doAction: function (semanticObject, semanticObjectType) {
    if (semanticObjectType == "hTodo") {
      return hTodo_delicious_bookmark.bookmark(
        semanticObject.url,
        semanticObject.summary,
        semanticObject.description
      );
    }
    return false;
  },
  /* Improved delicious bookmarking */
  bookmark: function (link, title, notes) {
    var url = "http://delicious.com/save?v=5&url=" + encodeURIComponent(link);

    if (typeof title != "undefined")
      url += "&title=" + encodeURIComponent(title);
    if (typeof notes != "undefined")
      url += "&notes=" + encodeURIComponent(notes);

    if (
      window.open(
        url + "&noui=1&jump=doclose",
        "deliciousuiv5",
        "location=yes,links=no,scrollbars=no,toolbar=no,width=550,height=550"
      )
    )
      return true;
    else return url + "&jump=yes";
  },
};

SemanticActions.add("delicious_bookmark", hTodo_delicious_bookmark);

/* Google / Yahoo Search */
var hTodo_search = {
  scope: {
    semantic: {
      hTodo: "summary",
    },
  },
};

SemanticActions.add("google_search", hTodo_search);
SemanticActions.add("yahoo_search", hTodo_search);

/* Goto URL (User Script) */
var hTodo_goto = {
  scope: {
    semantic: {
      hTodo: "url",
    },
  },
};

SemanticActions.add("goto_url", hTodo_goto);
```
