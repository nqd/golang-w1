package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	depth := flag.Int64("L", 100, "Max display depth of the directory tree.")
	directoryOnly := flag.Bool("d", false, "List directories only.")
	help := flag.Bool("h", false, "Prints usage informationn")
	output := flag.String("o", "", "Send output to filename.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()
	}

	flag.Parse()

	log.Println(depth, directoryOnly, help, output)

	if *help {
		flag.Usage()
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
	tree(*directoryOnly, currentDir, 0)
	if err != nil {
		log.Fatalln(err)
	}
}

func tree(d bool, dir string, deep int) (err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	for _, file := range files {
		// ignore file when needed
		if d && file.IsDir() == false {
			continue
		}

		tab := ""
		for i := 0; i < deep; i++ {
			tab = tab + "\t"
		}

		fmt.Printf(tab+"%s\n", file.Name())

		if file.IsDir() {
			child := dir + "/" + file.Name()
			err = tree(d, child, deep+1)
			if err != nil {
				return
			}
		}
	}
	return
}
