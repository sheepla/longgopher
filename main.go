package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	color "github.com/fatih/color"
)

var (
	appVersion   = "unknown"
	appRevision  = "unknown"
	appBuildDate = "unknown"
)

type asciiArt string

const (
	head asciiArt = `
      CCCCCCCCCCCCCCCCCCCCCC      
 CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC, 
CCCCCCCCWWWWWWWCCCCWWWWWWWCCCCCCCC
CCCCCCWWBBWWWWWWCCWWBBWWWWCCCCCCC  
 CCCCCCWBBWWWWWCCCCWBBWWWWCCCCC   
   CCCCCCCCCCCCBBBBCCCCCCCCCCCCC   
  CCCCCCCCCCCYYYYYYYYCCCCCCCCCCC  
  CCCCCCCCCCCCCWWWWWCCCCCCCCCCCC  
 YYYYYCCCCCCCCCCWWWCCCCCCCCCYYYYY
 YYYYYCCCCCCCCCCCCCCCCCCCCCCYYYYY  
`
	body asciiArt = `  CCCCCCCCCCCCCCCCCCCCCCCCCCCCCC   `
	leg  asciiArt = ` YYYYYYCCCCCCCCCCCCCCCCCCCYYYYYYY   
  YYYYYCCCCCCCCCCCCCCCCCCCYYYYYY   `
)

var (
	cyan   = color.New(color.Bold, color.FgHiCyan)
	black  = color.New(color.Bold, color.FgHiBlack)
	white  = color.New(color.Bold, color.FgHiWhite)
	yellow = color.New(color.Bold, color.FgHiYellow)
)

var stdout = bufio.NewWriter(os.Stdout)

var (
	length  int
	version bool
)

func main() {
	flag.IntVar(&length, "l", 10, "length of gopher's body")
	flag.BoolVar(&version, "V", false, "show version")
	flag.Parse()
	if version {
		fmt.Printf("v%s-%s\nBuild at %s\n", appVersion, appRevision, appBuildDate)
		return
	}
	printGopher(length)
}

func printGopher(length int) {
	h := head.Colorize('C', cyan).Colorize('W', white).Colorize('B', black).Colorize('Y', yellow)
	b := body.Colorize('C', cyan)
	l := leg.Colorize('C', cyan).Colorize('Y', yellow)

	fmt.Fprint(stdout, h)
	for i := 0; i < length; i++ {
		fmt.Fprintln(stdout, b)
	}
	fmt.Fprintln(stdout, l)
	stdout.Flush()
}

func (a asciiArt) Colorize(char rune, color *color.Color) asciiArt {
	str := strings.Replace(
		string(a),
		string(char),
		color.Sprint("#"),
		-1,
	)
	return asciiArt(str)
}
