package main

import (
    flag "github.com/spf13/pflag"
    "fmt"
    "io/ioutil"
    "os"
    "github.com/LukeEmmet/html2gemini"
    "bufio"
    "io"
    "errors"
)

var (
    output = flag.StringP("output", "o", "", "Output path. Otherwise uses stdout\n")
    input = flag.StringP("input", "i", "", "Input path. Otherwise uses stdin\n")
)

func check(e error) {
    if e != nil {
        panic(e)
        os.Exit(1)
    }
}

func saveFile(contents []byte, path string) {
    d1 := []byte(contents)
    err := ioutil.WriteFile(path, d1, 0644)
    check(err)
}

func readStdin () (string) {
    // based on https://flaviocopes.com/go-shell-pipes/
    reader := bufio.NewReader(os.Stdin)     //default size is 4096 apparently
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
    
        if info.Mode()&os.ModeNamedPipe != 0 {
            // we have a pipe input
            inputHtml = readStdin()
            
        } else if (*input != "") {
              //get the input file from the command line
              dat, err := ioutil.ReadFile(*input)
              check(err)
              inputHtml = string(dat)

        }  else  {
            //we shouldnt get here 
            return "", errors.New("Invalid option for input - use -i <path> or pipe to stdin")
        } 
        
        return inputHtml, nil
}

func main() {
        var inputHtml string
    
        flag.Parse()

        //get the input from commandline or stdin
        inputHtml, err:= getInput()
        check(err)
        
        //convert html to gmi
        text, err := html2text.FromString(inputHtml, html2text.Options{PrettyTables: true, GeminiCitationStyleLinks: true})
        check(err)
        
        //process the output
        if (*output == "") {
            fmt.Print(text)
        } else {
            //save to the specified output
            gmiBytes := []byte(text)            
            saveFile(gmiBytes, *output)
        }
}
