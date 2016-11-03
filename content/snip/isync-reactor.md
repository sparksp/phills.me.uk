+++
date = "2008-10-02T19:56:00Z"
modified = "2008-10-20T19:56:00Z"
title = "iSync Reactor"
aliases = ["/snips/12-isync-reactor", "/pb/12"]
tags = ["applescript"]
+++
{{< highlight applescript >}}
(*
iSync Reactor

This script will launch iSync and ask it to syncronize and terminate.  It will also refresh calendars in iCal.  
Protection is added to prevent frequent syncing and leave iCal/iSync running if they were already running

Copyright © Phill Sparks 2008
http://milk-hub.net


Creative Commons Attribution-Share Alike 2.0 UK: England & Wales
(http://creativecommons.org/licenses/by-sa/2.0/uk/legalcode)

You are free:
    * to copy, distribute, display, and perform the work
    * to make derivative works

Under the following conditions:
    * You must attribute this work to Phill Sparks (with link http://milk-hub.net/pb/12)
    * Share Alike. If you alter, transform, or build upon this work, you may distribute the resulting work only under a licence identical to this one.
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
