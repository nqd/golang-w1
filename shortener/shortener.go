package shortener

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type short struct {
	Short string
	Full  string
	Hit   int
}

type Shortener struct {
	shorts []short
	file   string
}

func NewShortener(file string) (shorten Shortener, err error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	if err = yaml.Unmarshal(f, &shorten.shorts); err != nil {
		return
	}
	shorten.file = file
	log.Printf("%+v", shorten)
	return
}

func (st Shortener) Add(s string, f string) (err error) {
	// not check duplicated keys
	st.shorts = append(st.shorts, short{
		Short: s,
		Full:  f,
		Hit:   0,
	})

	// write back
	if err = st.update(); err != nil {
		return
	}

	return
}
func (st Shortener) Remove(s string) (err error) {
	// crazy searching
	find := false
	for i, short := range st.shorts {
		if short.Short == s {
			// crazy removing
			if i < len(st.shorts)-1 {
				st.shorts = append(st.shorts[:i], st.shorts[i+1:]...)
			} else {
				st.shorts = st.shorts[:i]
			}
			find = true
		}
	}

	if find == true {
		// write back
		if err = st.update(); err != nil {
			return
		}
	}

	return
}
func (st Shortener) List() {
	for _, short := range st.shorts {
		fmt.Printf("Alias: %s, Origin: %s\n", short.Short, short.Full)
	}
}

func (st Shortener) query(s string) (f string, ok bool) {
	for _, short := range st.shorts {
		if short.Short == s {
			f = short.Full
			ok = true
			return
		}
	}
	return
}

func (st Shortener) update() (err error) {
	out, err := yaml.Marshal(st.shorts)
	if err != nil {
		return
	}
	log.Println(string(out))

	f, err := os.Create(st.file)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.Write(out)
	if err != nil {
		return
	}

	return
}
