package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"regexp"

	"github.com/charmbracelet/log"
)

func DownloadTorrentFileToTmp(url string) (string, error) {
	filename := "unspecified.torrent"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if val, ok := resp.Header["Content-Disposition"]; ok {
		re := regexp.MustCompile(`filename="?([^"]+)"?`)
		matchesWithQuotes := re.FindStringSubmatch(val[0])[1]
		log.Info(matchesWithQuotes)
		filename = matchesWithQuotes
	}
	fpath := path.Join(os.TempDir(), filename)
	floc, err := os.Create(fpath)
	log.Infof("Writing to path : %s", fpath)
	if err != nil {
		return "", err
	}

	downloadChan := make(chan int64, 1)
	go func(dchan chan int64) {
		// Use Copy Buffer
		sizeWritten, err := io.Copy(floc, resp.Body)
		if err != nil {
			return
		}
		downloadChan <- sizeWritten
	}(downloadChan)
	downloadSize := <-downloadChan
	log.Infof("Downloaded %.2f MiB", (float64(downloadSize))/(1024.0*1024.0))
	return fpath, nil
}
