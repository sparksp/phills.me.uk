+++
title = "{{ replace .Name "-" " " | title }}"
date = {{ .Date }}

[menu.main]
  parent = "course"

[[resources]]
  src = "{{ .Name }}.jpg"
  title = "Photo of {{ replace .Name "-" " " | title }}"
+++
