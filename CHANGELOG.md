# Changelog (2018)
- October 4th: Updated to latest gcd / protocol.json filee for 69.0.3497.100 for stable branch. Fixed bug in gcdapigen when resolving references to references to base types.
- June 1st: Updated to latest gcd / protocol.json file for 67.0.3396.62 for *stable* branch.
- April 24th: Updated to latest gcd / protocol.json file for 66.0.3359.117 for *stable* branch. Please note there was a change to protocol files again, more details: https://github.com/wirepair/gcd/issues/21.
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