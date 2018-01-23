+++
date = "2008-06-16T22:54:00Z"
title = "SSH Avahi Service"
description = "Simple Avahi definitions for SSH services including Fish and SCP."
aliases = ["/snips/11-ssh-avahi-service", "/snips/11"]
tags = ["xml"]
[menu]
  [menu.main]
    parent = "snip"
+++
Simple [Avahi](https://www.avahi.org/) definitions for SSH services including Fish and SCP.

{{< highlight xml >}}
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
{{< /highlight >}}
