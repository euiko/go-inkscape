package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/galihrivanto/go-inkscape"
)

var (
	svgInput  string
	pdfOutput string
	verbose   bool
)

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	flag.StringVar(&svgInput, "input", "", "svg input")
	flag.StringVar(&pdfOutput, "output", "result.pdf", "pdf output")
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.Parse()

	if svgInput == "" {
		fmt.Println("svg input is missing")
		os.Exit(1)
	}

	proxy := inkscape.NewProxy(inkscape.Verbose(verbose))

	err := proxy.Run()
	handleErr(err)
	defer proxy.Close()

	err = proxy.Svg2Pdf(svgInput, pdfOutput)
	handleErr(err)

	fmt.Println("done!!")
}
