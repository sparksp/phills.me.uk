+++
date = "2008-10-02T19:56:00Z"
modified = "2008-10-20T19:56:00Z"
title = "iSync Reactor"
description = "This script will launch iSync and ask it to syncronize and terminate. It will also refresh calendars in iCal. Protection is added to prevent frequent syncing and leave iCal/iSync running if they were already running."
aliases = ["/snips/12-isync-reactor", "/snips/12"]
tags = ["applescript", "osx", "code"]
categories = ["programming"]
+++
{{< highlight applescript >}}
(*
iSync Reactor

This script will launch iSync and ask it to syncronize and terminate.  It will also refresh calendars in iCal.  
Protection is added to prevent frequent syncing and leave iCal/iSync running if they were already running

Author: Phill Sparks <me@phills.me.uk>
License: Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales
License URL: https://creativecommons.org/licenses/by-sa/2.0/uk/
*)


property sync_frequency : 900 -- seconds

-- Find out if iCal and iSync are already running
tell application "System Events"
    set iCalRunning to (count of (every process whose name is "iCal")) is greater than 0
    set iSyncRunning to (count of (every process whose name is "iSync")) is greater than 0
end tell

-- Syncronize using iSync, at most every 15 minutes
tell application "iSync"
    if last sync is less than ((current date) - sync_frequency) then
        -- Refresh all calendars
        tell application "iCal"
            reload calendars

            if not iCalRunning then quit
        end tell

        -- Do the sync
        synchronize
    end if

    -- If this script started iSync then hide it, wait for it to finish syncing and quit
    if not iSyncRunning then
        tell application "System Events" to set visible of process "iSync" to false

        repeat while (syncing is true)
            -- Wait
        end repeat

        quit
    end if
end tell
{{< /highlight >}}
