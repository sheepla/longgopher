package main

import (
	"bufio"
	"flag"
	"fmt"
	"strings"

	color "github.com/fatih/color"
)

var (
	appVersion   = "unknown"
	appRevision  = "unknown"
	appBuildDate = "unknown"
)

type asciiArt string

type gopherAsciiArt struct {
	Head asciiArt
	Body asciiArt
	Leg  asciiArt
}

var gopher = gopherAsciiArt{
	// height: 11
	Head: `
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
`,
	Body: `  CCCCCCCCCCCCCCCCCCCCCCCCCCCCCC   `,

	// height: 2
	Leg: ` YYYYYYCCCCCCCCCCCCCCCCCCCYYYYYYY   
  YYYYYCCCCCCCCCCCCCCCCCCCYYYYYY   `,
}

var (
	cyan   = color.New(color.Bold, color.FgHiCyan)
	black  = color.New(color.Bold, color.FgHiBlack)
	white  = color.New(color.Bold, color.FgHiWhite)
	yellow = color.New(color.Bold, color.FgHiYellow)
)

var stdout = bufio.NewWriter(color.Output)

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
	head := gopher.Head.Colorize('C', cyan).Colorize('W', white).Colorize('B', black).Colorize('Y', yellow)
	body := gopher.Body.Colorize('C', cyan)
	leg := gopher.Leg.Colorize('C', cyan).Colorize('Y', yellow)

	fmt.Fprint(stdout, head)
	for i := 0; i < length; i++ {
		fmt.Fprintln(stdout, body)
	}
	fmt.Fprintln(stdout, leg)
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
