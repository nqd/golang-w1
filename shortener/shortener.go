package shortener

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Shortener struct {
	Short string
	Full  string
	Hit   int
}

var shorten []Shortener

func init() {
	f, err := ioutil.ReadFile("./record.yaml")
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
	if err := yaml.Unmarshal(f, &shorten); err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
	log.Printf("%+v", shorten)
}
