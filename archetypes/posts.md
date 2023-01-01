+++
title = "{{ replace .Name "-" " " | title }}"
date = {{ .Date }}
tags = []
categories = []

[[resources]]
  name = "featured-image"
  src = "{{ .Name }}.jpg"
  title = "Photo of {{ replace .Name "-" " " | title }}"
+++
