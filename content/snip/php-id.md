+++
date = "2007-12-19T15:29:00Z"
title = "PHP Id"
aliases = ["/snips/6-php-id", "/snips/6"]
tags = ["bash"]
[menu]
  [menu.main]
    parent = "snip"
+++
{{< highlight bash >}}
#!/bin/bash

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
{{< /highlight >}}
