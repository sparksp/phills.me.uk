+++
date = "2007-11-19T16:05:00Z"
title = "Factor Test"
description = "PHP factor test for very large numbers, by lavalamp."
aliases = [
  "/snip/4*",
  "/snip/04*",
  "/snip/factor-test"
]
tags = ["php", "code"]
categories = ["programming", "snip"]
author = "lavalamp"
toc.enable = false
+++

PHP factor test for very large numbers.

```php
<?php

function check_factor($factor, $k, $base, $exponent, $c='-1'){

  $out  = bcpowmod($base, $exponent, $factor);
  $out  = bcmul($out, $k);
  $out  = bcadd($out, $c);
  $out  = bcmod($out, $factor);

  return $out=="0";

}

echo check_factor('28475025393798152885081', '1', '2', '3321931637')? "True!": "False!";
```
