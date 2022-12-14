+++
title = "{{ replace .Name "-" " " | title }}"
date = {{ .Date }}

[[resources]]
    src = "{{ .Name }}.jpg"
    title = "Description of this photo."
+++
