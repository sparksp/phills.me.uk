+++
date = "2008-06-16T22:54:00Z"
title = "SSH Avahi Service"
aliases = ["/snips/11-ssh-avahi-service", "/pb/11"]
tags = ["xml"]
[menu.main]
  parent = "snip"
+++
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
