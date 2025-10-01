package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/charmbracelet/log"
	"github.com/jackpal/bencode-go"
	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: program <torrent-file/torrent-url>")
	}
	openFile := os.Args[1]

	if strings.Contains(os.Args[1], "http") {
		dpath, err := DownloadTorrentFileToTmp(os.Args[1])
		if err != nil {
			log.Errorf("unable to download file : %s", err.Error())
		}
		openFile = dpath
	}

	file, err := os.Open(openFile)
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
