package qubetorrent

import (
	"os"
	"strings"

	log "github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

func qube_main() {
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
