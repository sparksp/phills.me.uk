+++
date = "2007-11-19T16:05:00Z"
title = "Factor Test (by lavalamp)"
aliases = ["/snips/4-factor-test", "/pb/04"]
tags = ["php"]
[menu.main]
  parent = "snip"
+++
{{< highlight php >}}
<?php

function check_factor($factor, $k, $base, $exponent, $c='-1'){

  $out  = bcpowmod($base, $exponent, $factor);
  $out  = bcmul($out, $k);
  $out  = bcadd($out, $c);
  $out  = bcmod($out, $factor);

  return $out=="0";

}

echo check_factor('28475025393798152885081', '1', '2', '3321931637')? "True!": "False!";

?>
{{< /highlight >}}
