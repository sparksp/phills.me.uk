+++
title = "{{ replace .Name "-" " " | title }}"
date = {{ .Date }}

[menu.main]
  parent = "services"

[[resources]]
  name = "featured-image"
  src = "{{ .Name }}.jpg"
  title = "Photo of {{ replace .Name "-" " " | title }}"
+++
