package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/LukeEmmet/html2gemini"
	flag "github.com/spf13/pflag"
	"io"
	"io/ioutil"
	"os"
)

var version = "0.2.6"

var (
	output            = flag.StringP("output", "o", "", "Output path. Otherwise uses stdout")
	input             = flag.StringP("input", "i", "", "Input path. Otherwise uses stdin")
	citationStart     = flag.IntP("citationStart", "c", 1, "Start citations from this index")
	citationMarkers   = flag.BoolP("citationMarkers", "m", false, "Use footnote style citation markers")
	numberedLinks     = flag.BoolP("numberedLinks", "n", false, "Number the links")
	prettyTables      = flag.BoolP("prettyTables", "t", false, "Pretty tables - works with most simple tables")
	emitImagesAsLinks = flag.BoolP("emitImagesAsLinks", "e", false, "Emit links to included images")
	linkEmitFrequency = flag.IntP("linkEmitFrequency", "l", 2, "Emit gathered links through the document after this number of paragraphs")
	verFlag           = flag.BoolP("version", "v", false, "Find out what version of html2gmi you're running")
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func saveFile(contents []byte, path string) {
	d1 := contents
	err := ioutil.WriteFile(path, d1, 0644)
	check(err)
}

func readStdin() string {
	// based on https://flaviocopes.com/go-shell-pipes/
	reader := bufio.NewReader(os.Stdin) //default size is 4096 apparently
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return string(output)
}

func getInput() (string, error) {
	var inputHtml string

	info, err := os.Stdin.Stat()
	check(err)

	if *input != "" {
		//get the input file from the command line
		dat, err := ioutil.ReadFile(*input)
		check(err)
		inputHtml = string(dat)
	} else if info.Mode()&os.ModeNamedPipe != 0 {
		// we have a pipe input
		inputHtml = readStdin()

	} else {
		//we shouldn't get here
		return "", errors.New("invalid option for input - use -i <path> or pipe to stdin")
	}

	return inputHtml, nil
}

func main() {
	var inputHtml string

	flag.Parse()

	if *verFlag {
		fmt.Println("html2gmi " + version)
		return
	}

	//get the input from commandline or stdin
	inputHtml, err := getInput()
	check(err)

	//convert html to gmi
	options := html2gemini.NewOptions()
	options.PrettyTables = *prettyTables
	options.CitationStart = *citationStart
	options.LinkEmitFrequency = *linkEmitFrequency
	options.CitationMarkers = *citationMarkers
	options.NumberedLinks = *numberedLinks
	options.EmitImagesAsLinks = *emitImagesAsLinks


	//dont use an extra line to separate header from body, but
	//do separate each row visually
	options.PrettyTablesOptions.HeaderLine = false
	options.PrettyTablesOptions.RowLine = true

	//use slightly nicer Unicode borders, otherwise can use +,|,-
	//options.PrettyTablesOptions.CenterSeparator = "┼"
	//options.PrettyTablesOptions.ColumnSeparator = "│"
	//options.PrettyTablesOptions.RowSeparator = "─"

    //pretty tables option is somewhat experimental
    //and the column positions not always correct
    //so use invisible borders of spaces for now
    options.PrettyTablesOptions.CenterSeparator = " "
	options.PrettyTablesOptions.ColumnSeparator = " "
	options.PrettyTablesOptions.RowSeparator = " "

	ctx := html2gemini.NewTraverseContext(*options)

	text, err := html2gemini.FromString(inputHtml, *ctx)

	check(err)

	//process the output
	if *output == "" {
		fmt.Print(text + "\n") //terminate with a new line
	} else {
		//save to the specified output
		gmiBytes := []byte(text + "\n")
		saveFile(gmiBytes, *output)
	}
}
