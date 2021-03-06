package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path"
)

// Proverb ...
type Proverb struct {
	ID       int
	Text     string
	reviewed bool
}

// UnmarshalBinary decodes the form genereated by MarshalBinary
func (p *Proverb) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanf(b, "ID=%d, Text=%q, reviewed=%t\n", &p.ID, &p.Text, &p.reviewed)
	return err
}

func main() {
	filename := path.Join("..", "proverbs.gob")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var proverbs []Proverb

	dec := gob.NewDecoder(file)
	if err := dec.Decode(&proverbs); err != nil {
		log.Fatalln(err)
	}

	for _, p := range proverbs {
		log.Printf("%#v\n", p)
	}
}
