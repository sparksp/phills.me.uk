+++
date = "2008-04-13T21:30:00Z"
slug = "how_to_customise_your_ssh_config"
title = "How-To: Customise your SSH Config"
tags = ["howto", "SSH", "config", "unix"]
categories = ["blog", "programming"]
aliases = [
  "/blog/2008/04/13/how_to_customise_your_ssh_config",
  "/post/how_to_customise_your_ssh_config"
]
toc.enable = false
+++

I am someone who feels more comfortable using the command-line than many GUI set-ups and who is often SSH-ing between servers; many of these servers have lengthy host names and require various different usernames; luckily the SSH developers thought ahead and came up with a way of pre-setting most options per-host.

Here's an example that will configure "server1.example.com" with the user "example" and the port "222". I'll set this all up as "ex1". Just drop the following in to `~/.ssh/config`:

```text
Host ex1
HostName server1.example.com
User example
Port 222
```

Now if I call `ssh ex1` it's just like calling `ssh server1.example.com -l example -p 222`. You can set-up defaults by using the Host "\*", I often default the User where I can't have my preferred name...

```text
Host *
User example
```

There are a lot of things you can configure in this way, for a full list check out the man page for ssh_config (5).
