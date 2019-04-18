package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	depth := flag.Int("L", 100, "Max display depth of the directory tree.")
	directoryOnly := flag.Bool("d", false, "List directories only.")
	help := flag.Bool("h", false, "Prints usage informationn")
	output := flag.String("o", "", "Send output to filename.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()
	}

	flag.Parse()

	var out *os.File

	if *output == "" {
		out = os.Stdout
	} else {
		var errCreate error
		out, errCreate = os.Create(*output)
		if errCreate != nil {
			log.Fatalln(errCreate)
			os.Exit(-1)
		}
		defer out.Close()
	}

	if *help {
		flag.Usage()
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
	tree(*directoryOnly, currentDir, 0, *depth, out)
	if err != nil {
		log.Fatalln(err)
	}
}

func tree(d bool, dir string, dept int, maxDepth int, out io.Writer) (err error) {
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
		for i := 0; i < dept; i++ {
			tab = tab + "\t"
		}

		// fmt.Printf(tab+"%s\n", file.Name())
		fmt.Fprintf(out, tab+"%s\n", file.Name())

		if file.IsDir() {
			if dept >= maxDepth {
				continue
			}
			child := path.Join(dir, "/", file.Name())
			err = tree(d, child, dept+1, maxDepth, out)
			if err != nil {
				return
			}
		}
	}
	return
}
