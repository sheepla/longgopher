package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	color "github.com/fatih/color"
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
 YYYYCCCCCCCCCCCWWWCCCCCCCCCYYYYY
 YYYYCCCCCCCCCCCCCCCCCCCCCCCYYYYY  
`
	body asciiArt = `  CCCCCCCCCCCCCCCCCCCCCCCCCCCCCC   `
	leg  asciiArt = ` YYYYYYYCCCCCCCCCCCCCCCCYYYYYYYYY  `
)

var (
	cyan   = color.New(color.Bold, color.FgHiCyan)
	black  = color.New(color.Bold, color.FgHiBlack)
	white  = color.New(color.Bold, color.FgHiWhite)
	yellow = color.New(color.Bold, color.FgHiYellow)
)

var stdout = bufio.NewWriter(os.Stdout)

var length int

func main() {
	flag.IntVar(&length, "l", 10, "length")
	flag.Parse()
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
