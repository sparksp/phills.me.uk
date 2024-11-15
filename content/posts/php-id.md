+++
date = "2007-12-19T15:29:00Z"
title = "PHP Id"
description = "A handy script to insert $Id$ comments in PHP files."
aliases = [
  "/snip/6*",
  "/snip/06*",
  "/snip/php-id"
]
tags = ["bash", "code"]
categories = ["programming", "snip"]
license = "<a href=\"http://creativecommons.org/licenses/by-sa/2.0/uk/\">Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales</a>"
toc.enable = false
+++

A handy script to insert `$Id$` comments in PHP files.

```bash
#!/bin/bash

# A handy script to insert $Id$ comments in PHP files.
#
# Author: Phill Sparks <me@phills.me.uk>
# License: Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales
# License URL: http://creativecommons.org/licenses/by-sa/2.0/uk/

case $# in
    1);;
    *) echo "Usage: $0 file";exit;;
esac;

# File exists?
if [[ -e $1 ]]; then

    # File already has $Id$ in it?
    if grep -q '\$Id.*\$$' $1 ; then
        echo "$1 [skipped]"
    else

        # After the first <?php (that does not have a ?> on the same line) add the $Id$ comment
        sed '
        /<?php/ {
            /?>/ !{
                x
                /^$/ {
                    g
                    a \
/*\
 * $Id$\
 */
                    h
                }
                g
            }
        }
        ' "$1" >/tmp/sk
        mv /tmp/sk "$1"
        echo "$1"

    fi

else
    echo "$1 [missing]"
fi
```
