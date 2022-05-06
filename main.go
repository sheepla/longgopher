package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	color "github.com/fatih/color"
)

var cyan = color.New(color.Bold, color.FgHiCyan)

var (
	head = cyan.Sprint(`
         ,_---~~~~~----._         
  _,,_,*^____      _____\\*g*\"*, 
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/  
 \*/   \~____ / __ \_____/    \   
  |           _l__l_           I   
  }          [______]          I  
  ]            | | |           |  
  ]             ~ ~            |  
  |                            |   
`,
	)
	body = cyan.Sprint(`  |                            |   `)
	leg  = cyan.Sprint(`  (_///).,______________(_//.,_)   `)
)

var stdout = bufio.NewWriter(os.Stdout)

var length int

func main() {
	flag.IntVar(&length, "l", 10, "length")
	flag.Parse()
	printGopher(length)
}

func printGopher(length int) {
	fmt.Fprint(stdout, head)
	for i := 0; i < length; i++ {
		fmt.Fprintln(stdout, body)
	}
	fmt.Fprintln(stdout, leg)
	stdout.Flush()
}
