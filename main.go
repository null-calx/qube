package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jackpal/bencode-go"
	"gopkg.in/yaml.v3"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Fatal("Usage: program <torrent-file>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var torrentFile MetainfoFile
	err = bencode.Unmarshal(file, &torrentFile)
	if err != nil {
		log.Fatalf("Failed to decode: %v", err)
	}
	torrentYaml, err := yaml.Marshal(torrentFile)
	if err != nil {
		log.Fatalf("Failed to encode yaml: %v", err)
	}

	fmt.Printf("%s", torrentYaml)
}
