+++
date = "2008-05-21T22:13:00Z"
modified = "2008-05-22T22:13:00Z"
title = "New Item Alert"
description = "This Folder Action handler is triggered whenever items are added to the attached folder. The script will display a growl message containing the number of items added."
aliases = [
  "/snip/9*",
  "/snip/09*",
  "/snip/new-item-alert"
]
tags = ["applescript", "osx", "code"]
categories = ["programming", "snip"]
license = "<a href=\"http://creativecommons.org/licenses/by-sa/2.0/uk/\">Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales</a>"
toc.enable = false
+++

This Folder Action handler is triggered whenever items are added to the attached folder. The script will display a [Growl](https://growl.github.io/growl/) message containing the number of items added.

```applescript
(*
New Item Alert

This Folder Action handler is triggered whenever items are added to the attached folder.
The script will display a growl message containing the number of items added.

Author: Phill Sparks <me@phills.me.uk>
License: Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales
License URL: https://creativecommons.org/licenses/by-sa/2.0/uk/
*)

property dialog_timeout : 30 -- set the amount of time before dialogs auto-answer.

on adding folder items to this_folder after receiving added_items
    try
        -- Determine if Growl is running
        tell application "System Events"
            set hasGrowl to (count of (every process whose name is "GrowlHelperApp")) > 0
        end tell

        tell application "Finder"
            -- Get the name of the folder
            set the folder_name to the name of this_folder
        end tell

        -- Find out how many new items have been placed in the folder
        set the item_count to the number of items in the added_items

        -- Create the alert string
        set alert_message to ("") as Unicode text
        if the item_count is greater than 1 then
            set alert_message to alert_message & (the item_count as text) & " new items have "
            set alert_title to "New Items"
        else
            set alert_message to alert_message & "One new item has "
            set alert_title to "New Item"
        end if
        set alert_message to alert_message & "been placed in the " & «data utxt201C» & the folder_name & «data utxt201D» & " folder."

        -- If Growl is running then show the notification, otherwise use a prompt
        if hasGrowl then
            tell application "GrowlHelperApp"
                -- List the notifications for Growl
                set the allNotificationsList to {"New Item"}

                -- Enable the notifications
                set the enabledNotificationsList to {"New Item"}

                -- Let Growl know who we are
                register as application "Directory Monitor" all notifications allNotificationsList default notifications enabledNotificationsList icon of application "Finder"

                -- Do the notify
                notify with name "New Item" title alert_title description the alert_message application name "Directory Monitor"
            end tell
        else
            display dialog the alert_message with title alert_title buttons {"OK"} default button 1 cancel button 1 with icon 1 giving up after dialog_timeout
        end if
    end try
end adding folder items to
```
