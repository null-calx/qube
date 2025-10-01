package main

import (
	"log"
	"os"

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

	torrent, err := readTorrentFile(file)
	if err != nil {
		log.Fatalf("Failed to unmarshal torrent file: %v", err)
	}

	qc, err := loadConfig("qubetorrent.yaml")
	if err != nil {
		log.Fatalf("Failed to load qubetorrent config: %v", err)
	}

	trackerResponse, err := torrent.requestPeers(qc.PeerId, uint16(qc.Port))
	if err != nil {
		log.Fatalf("Failed requesting peers: %v", err)
	}

	trackerYaml, err := yaml.Marshal(trackerResponse)
	if err != nil {
		log.Fatalf("Failed to encode yaml: %v", err)
	}

	log.Printf("%s", trackerYaml)
}
