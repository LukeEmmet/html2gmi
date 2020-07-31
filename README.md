# html2gmi-cli
A command line application to convert HTML to GMI (Gemini text/gemini), written in Go.

A simple wrapper around the Go library https://github.com/LukeEmmet/html2gemini

## Usage

```
html2gmi-cli -i <inputfile> -o <outputfile>
```

## Flags

-i \<inputfile\> specifies the file to be processed (Otherwise uses stdin)

-o \<outputfile\> specifies the output to be saved (Otherwise uses stdout)

## Remarks

You can pipe content in from other applications, for example utilities that download HTML from the web.

To strip out cruft, you can use an HTML sanitiser before passing the content to this app.