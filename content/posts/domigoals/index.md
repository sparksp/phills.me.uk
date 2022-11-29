+++
date = "2012-06-02T15:41:45Z"
title = "Domigoals 2012"
categories = ["programming", "project"]
aliases = [
  "/project/domigoals"
]
toc.enable = false

[[resources]]
  name = "featured-image"
  src = "domigoals.jpg"
  title = "Three screens from the Domigoals 2012 app"
+++

Domino's UK wanted to give away vouchers for every goal scored during Euro 2012. The first 1000 people to tap the "Goal!" button would win a voucher – &pound;5 for the first goal, &pound;10 for the second by the same player, and &pound;20 for a hat-trick.

My role was to develop the server infrastructure to support the apps, including fetching the game data from a 3rd party, registering new users, serving game data to the apps, and correctly awarding vouchers for each goal.

The biggest challenge for me was in handling the massive spikes of users all wanting to win vouchers as each goal was scored, without dropping or blocking any requests. I was required to think outside of the usual RDBMS box when allocating vouchers and aligning them with the winners.
