# Changelog (2024)
- 2.3.5 (Oct 16)
  - Fix apigen to use ToT from github for downloading json files. Update to latest protocol 
- 2.3.4 (Mar 18)
  - Fix return types for calls that return two-dimensional arrays 
- 2.3.3 (Jan 31)
  - Improve JSON performance

# Changelog (2023)
- 2.3.2 (Dec 6)
  - Reorder and initialize missing DevTools domains 
- 2.3.1 (May 30) 
  - Applied patch from @camswords to expose dev tools version
  - Updates to chrome 112.0.5615.138
- 2.3.0 (Mar 28) Updated JSON endpoint calls to use PUT instead of GET as GET returns an error since 111.
  - Updates to chrome 111.0.5563.147 protocol
  - gcdapigen updated to handle different revisions thanks to @eatdrinksleepcode

# Changelog (2022)
- 2.2.6 (Sept 5th) Fixes a bug where new types that were just arrays to another ref type were not being output correctly
  - Updates to chrome 105.0.5195.102
- 2.2.5 (March 31st) Upgrade protocol.json to 100.0.4896.60

# Changelog (2021)
- 2.2.4 (November 27th) Upgrade protocol.json to 96.0.4664.45
- 2.2.3 (August 23rd) Applied patch from @camswords to allow using custom output writer for chrome process to aid in debugging.
- 2.2.2 (August 19th) Applied patch from @camswords to fix the cannot connect to debugger message to print seconds out properly.
- 2.2.1 (July 28th) SendDefaultRequests were not returning ChromeErrorResponses
- 2.2.0 (July 21st) Dispatching DevTool events via a newly spawned go routine was causing messages to be delivered out of order. This change synchronizes them using an internal channel. Note: If you previously had Subscriptions try to signal each other (via another channel) it may be blocked since all subscriptions are executed under a single go routine now. Upgrade with caution.
- 2.1.6 (June 9th) Fix go routine leak and add test
- 2.1.5 (June 8th) Fix race condition on error and endpoint
- 2.1.4 (June 8th) Added a chrome exit handler from @camswords.
- 2.1.3 (June 4th): Fix a potential blocked channel if chrome fails to start and debug port probe fails
- 2.1.2 (May 23rd): Update to chrome version 90.0.4430.212
  - Implementation of an observer package to allow inspection of messages between gcd/browser by @camswords

- 2.1.1 (April 29th): Update to chrome version 90.0.4430.93.
  - Replaced websocket with [rod's](https://go-rod.github.io/#/) implementation (You should use their library!).
  - Added functional opts for creating Gcd objects.
  - Allow custom logger

# Changelog (2020)

- September 1st: Update to chrome version 85.0.4183.83, add a lock around process for -race.
- August 7th: Update to chrome version 84.0.4147.105
- July 24th: Fix only closing wsconn in the event of specific ws errors, not io/timeout which would happen if context was canceled.
- July 24th: Return proper error if context is done.
- July 23rd: Finally move to using context.Context for all messages, this is a breaking change so updated to v2.0.0. Going forward you'll need to go get github.com/wirepair/gcd/v2
- July 5th: Fix more timeout issues and wsConnsize was set to 1mb which is insufficient
- July 2nd: Fix SendCustomReturn to timeout properly if send failed. Upgrade to latest gcd /protocol.json file for 83.0.4103.116
- May 30th: Fix wsconn.Close, add retries for tabs and updated to latest gcd/ protocol.json file for 83.0.4103.61
- May 19th: Replace websocket code with github.com/gobwas/ws and fixed a terrible bug in handling websocket reads. Performance should be significantly improved.
- May 18th: Replace encoding/json with https://github.com/json-iterator/go, heavy DOM usage should see reduction of memory by half. Updated to latest gcd / protocol.json file for 81.0.4044.138
- April 19th: Fix bad merge
- April 19th: Move to go modules, remove vendor directory. Fix gcdapigen to output version properly. Updated to latest gcd / protocol.json file for 81.0.4044.113

# Changelog (2019)

- July 23rd: Updated to latest gcd / protocol.json file for 75.0.3770.142. Added WebAuthN APIs.
- May 4th: Updated to latest gcd / protocol.json file for 74.0.3729.131. Added Audits and WebAudio APIs.
- March 13th: Updated to latest gcd / protocol.json file for 73.0.3683.75. Added ability custom start the process

# Changelog (2018)

- December 8th: Updated to latest gcd / protocol.json file for 71.0.3578.80
- October 18th: Updated to latest gcd / protocol.json file for 70.0.3538.67.
- October 17th: Added new helper features to Gcd, GetFirstTab()) and DeleteProfileOnExit()
- October 4th: Updated to latest gcd / protocol.json file for 69.0.3497.100 for stable branch. Fixed bug in gcdapigen when resolving references to references to base types.
- June 1st: Updated to latest gcd / protocol.json file for 67.0.3396.62 for _stable_ branch.
- April 24th: Updated to latest gcd / protocol.json file for 66.0.3359.117 for _stable_ branch. Please note there was a change to protocol files again, more details: https://github.com/wirepair/gcd/issues/21.
- Feburary 22nd: Updated to latest gcd / protocol.json file for 66.0.3346.8. Added dep init/dep ensure to repository for package management.
- January 17th: Updated to latest protocol. Removed examples from README.md because API changes too often, added a link to examples instead which I ensure work after update. Fixed \n in descriptions of protocol.json files causing template generation errors. Current is 65.0.3322.3. Updated examples to work with latest API changes.

# Changelog (2017)

- November 20th: Updated to latest protocol. Current is 64.0.3269.3. Updated tests and examples to work with latest API changes.
- October 30th: Updated to latest protocol. Current is 64.0.3251.0. Updated examples to work with latest API changes.
- September 9th: Updated to latest protocol files.
  We now download specific channels protocol files. By default we download the 'dev' channel. You can override this if you wish in gcdapigen by passing gcdapigen -update -channel {canary,stable,...}. To see what version gcd is bound to, check the gcdapi/version.go file, or call gcd.GetRevision(). Current is 62.0.3202.9.
- August 15th: Updated to the latest protocol.json
- July 8th: Updated to latest protocol.json which supports network interception. However it does not appear to work in most versions yet. (I could only get it working in the latest https://download-chromium.appspot.com/ download.) Also, going forward GCDVERSION will be Major.Year.Month.Day.Minor format.
- June: Due to the excellent efforts of [Ian L.](https://github.com/MrSaints) gcd can now issue requests ignoring fields. A non-breaking change was implemented allowing callers to use new methods (denoted by WithParams) to pass in structures, and choosing which fields to populate. Users can continue to use the old method passing individual arguments.
  Example:

```Go
networkParams := &gcdapi.NetworkEnableParams{
    MaxTotalBufferSize:    -1,
    MaxResourceBufferSize: -1,
}

if _, err := network.EnableWithParams(networkParams); err != nil {
    log.Fatal("error enabling network")
}
```

- May: Updated to latest protocol.json
- May: Fixed a bug with templating causing certain parameters that were slices to only show up as the base type.
- May: Changed templating where 'any' protocol.json field types from string to interface{} to allow caller to decide how to decode.
- April: Updated to the latest protocol.json (version 1.2). Note this changes quite a few APIs.

# Changelog (2016)

- June: Updated to the latest protocol.json, gcdapigen will download the js_protocol and browser_protocol json files from chromium repositories. It will also fix them up and merge them into a single file and output it.
  Note that several API endpoints have been removed and method calls have changed since the last update.
- February: I created a new library for actual automation purposes. If you want something with more functionality and more usability I suggest checking out [autogcd](https://github.com/wirepair/autogcd).
