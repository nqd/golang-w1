package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	depth := flag.Int64("L", 100, "Max display depth of the directory tree.")
	hasDirectory := flag.Bool("d", false, "List directories only.")
	help := flag.Bool("h", false, "Prints usage informationn")
	output := flag.String("o", "", "Send output to filename.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()
	}

	flag.Parse()

	log.Println(depth, hasDirectory, help, output)

	if *help {
		flag.Usage()
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
	log.Println(currentDir)
}
