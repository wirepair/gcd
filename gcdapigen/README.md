# Google Chrome Debugger (GCD) API Generator
You should hopefully not have to use this, but this package is for generating commands, types and events for the Google Chrome Debugger Service for remotely debugging a Chrome Process.

Basically it reads in the protocol.json file, does some gnarly processing and spits out the code using templates. All references are looked up so we don't have to do nasty type conversions just to get to a base underlying type. This was pretty much a full rewrite, moved all API stuff to it's own package so it doesn't pollute the gcd package.

Updated September 2015 with latest protocol.json.

If you plan on downloading later versions of protocol.json, keep in mind there are some things you'll need to fix by hand. In particular one field which *should* be a bool is actually a string of "true". Go's JSON error handling is absolutely horrible in that it won't tell you the exact line the unmarshalling fails on so I had to fix it by hand. 

## Developer Notes
The protocol.json file while appears simple at first is pretty complex, with lots of crazy edge cases and variables, all around how they use '$ref' fields. You'll notice a *lot* of conditionals when looking up references, i just saw no other way to accomplish generating types without having hundereds of dumb type SomeObject string or type SomeOtherObject []int everywhere. To make using the API easier, i resolved all of these references (at least I think I did) and tried to limit the amount of work necessary for unmarshalling the various data types requires.

If you need to fix references, the two main areas are utils.go:PopulateReferences and domain.go:resolveReferences. We first loop over the entire set of domains and make a map of all reference types, storing various critical metadata about the references (is it array, is it a ref to a underlying type, etc). As types are generated, it looks at this map and replaces any definitions with the proper ones. 

## Usage
Download the latest [protocol.json](https://code.google.com/p/chromium/codesearch#chromium/src/third_party/WebKit/Source/devtools/protocol.json&q=protocol.json&sq=package:chromium&type=cs) and make sure there aren't any string encoded bools (usually hidden or optional fields), fix if necessary, then run:
```
go build
./gcdapigen protocol.json
cp -R output/. ../gcdapi
cd !$
go build && go install
```

## Licensing
The MIT License (MIT)

Copyright (c) 2015 isaac dawson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
