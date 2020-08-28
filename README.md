# html2gmi

A command line application to convert HTML to GMI (Gemini text/gemini), written in Go. Written in Go so it should compile on various common platforms (linux/mac/windows etc).

A simple wrapper around the Go library https://github.com/LukeEmmet/html2gemini

## Usage

```
html2gmi <flags>

  -m, --citationMarkers         Use footnote style citation markers
  -c, --citationStart int       Start citations from this index (default 1)
  -i, --input string            Input path. Otherwise uses stdin
  -l, --linkEmitFrequency int   Emit gathered links through the document after this number of paragraphs (default 2)
  -n, --numberedLinks           Number the links
  -o, --output string           Output path. Otherwise uses stdout
  -e, --emitImagesAsLinks       Emit links to included images
  -t, --prettyTables            Pretty tables - works with most simple tables
  -v, --version                 Find out what version of html2gmi you're running

  
```


## Remarks

* linkEmitFrequency - this flag determines the frequency of the output of link lists. For example 1 means any links from the previous paragraph are emitted after each paragraph. 2, would mean they are emitted every two paragraphs and so on. Any gathered links are always emitted before a new heading.
* citationStart - this flag determines the start index of the links. By default this is 1, so the first link is labelled "[1]", but you can set this as required.
* citationMarkers - use a numbered marker in the text to indicate the location of the citation, [1], [2] etc. 
* numberedLinks - number the links with a reference number [1], [2] etc. Certain command line Gemini clients may automatically add these, in which case you can omit them.
* emitImagesAsLinks - add a link for every embedded image in addition to its placemarker
* prettyTables - tables will be displayed as preformatted content. Complex tables may not look perfect. Otherwise each table row is a new line.

You can pipe content in from other applications, for example utilities that download HTML from the web. 

To strip out cruft, use an HTML sanitiser before passing the content to this app.

## Building

Compile using Go.

Retrieve the package using Go get

```
go get github.com/LukeEmmet/html2gmi

```

Build the binary using Go build

```
go build github.com/LukeEmmet/html2gmi
```

# History

## 0.2.5

* new -t flag to emit pretty tables (as preformatted content)
* improve table rendering when prettyTables is off
* don't use a border marker for preformatted tables

## 0.2.4

* option to toggle emitting links for embedded images

## 0.2.2

* option to toggle numbering on links
* option to toggle citation markers

## 0.2.1

* public release