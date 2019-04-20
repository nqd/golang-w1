package shortener

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)
	s, err := NewShortener("../record.yaml")
	log.Println(s)
	assert.Nil(err)

	s.add("hey", "google.com")
	s.remove("hey")
	// t.Log(shorten)
	// t.Error(err)
}
