+++
date = "2008-06-16T22:54:00Z"
title = "SSH Avahi Service"
description = "Simple Avahi definitions for SSH services including Fish and SCP."
aliases = [
  "/snip/11*",
  "/snip/ssh-avahi-service"
]
tags = ["xml", "unix", "code"]
categories = ["programming", "snip"]
license = "<a href=\"http://creativecommons.org/licenses/by-sa/2.0/uk/\">Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales</a>"
+++

Simple [Avahi](https://www.avahi.org/) definitions for SSH services including Fish and SCP.

### Code

```xml
<?xml version="1.0" standalone='no'?>
<!DOCTYPE service-group SYSTEM "avahi-service.dtd">
<service-group>
    <name>My Server</name>
    <service>
        <type>_ssh._tcp</type>
        <port>22</port>
    </service>
    <service>
        <type>_fish._tcp</type>
        <port>22</port>
    </service>
    <service>
        <type>_scp._tcp</type>
        <port>22</port>
    </service>
</service-group>
```
