+++
title = "{{ replace .Name "-" " " | title }}"
date = {{ .Date }}
toc.enable = false
meta = false

[[resources]]
    name = "featured-image"
    src = "{{ .Name }}.jpg"
    title = "Photo of {{ replace .Name "-" " " | title }}"
+++
