# Google Chrome Debugger (GCD) Protocol Generator
You should hopefully not have to use this, but this package is for generating commands and types (soon events maybe) for the Google Chrome Debugger Service for remotely debugging a Chrome Process.

Basically it reads in the protocol.json file, does some gnarly processing and spits out the code & types using templates. You'll notice the type spitting out code is a horrible hack because i didn't know how to use templates very well.

Code templates are a bit cleaner but a lot is left to be desired, mainly there is a lot of repeat structs i could replace and I could make the template a bit cleaner but guess what? I did this in like 2 days so gimmie a break.

## Usage
No. Stahp.

## Licensing
The MIT License (MIT)

Copyright (c) 2014 isaac dawson

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
