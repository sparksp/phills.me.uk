+++
title = "{{ replace .Name "-" " " | title }}"
date = {{ .Date }}
link = "//www.{{ .Name }}.com/"
association = "{{ .Name }}"

[[resources]]
  name = "logo"
  src = "{{ .Name }}.png"
  title = "{{ replace .Name "-" " " | title }} logo"
+++
