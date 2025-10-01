package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jackpal/bencode-go"
)

func readTorrentFile(r io.Reader) (torr TorrentFile, err error) {
	var metainfofile MetainfoFile
	err = bencode.Unmarshal(r, &metainfofile)
	if err != nil {
		return
	}
	torr, err = metainfofile.toTorrentFile()
	return
}

func (metainfofile MetainfoFile) toTorrentFile() (TorrentFile, error) {
	infoHash, err := metainfofile.infoHash()
	if err != nil {
		return TorrentFile{}, err
	}

	piecesHash, err := metainfofile.piecesHash()
	if err != nil {
		return TorrentFile{}, err
	}

	torr := TorrentFile{
		Announce:   metainfofile.Announce,
		InfoHash:   infoHash,
		PiecesHash: piecesHash,
		Length:     metainfofile.Info.PieceLength * len(piecesHash),
	}
	return torr, nil
}

func (metainfofile MetainfoFile) infoHash() ([20]byte, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, metainfofile.Info)
	if err != nil {
		return [20]byte{}, err
	}
	hash := sha1.Sum(buf.Bytes())
	return hash, nil
}

func (metainfofile MetainfoFile) piecesHash() ([][20]byte, error) {
	pieces := []byte(metainfofile.Info.Pieces)
	if len(pieces)%20 != 0 {
		return [][20]byte{}, fmt.Errorf("malformed pieces of length %d", len(pieces))
	}
	numPieces := len(pieces) / 20
	splitPieces := make([][20]byte, numPieces)

	for i := range numPieces {
		copy(splitPieces[i][:], pieces[20*i:20*(i+1)])
	}

	return splitPieces, nil
}

func (torr TorrentFile) getAnnounceURL(peerId PeerId, port uint16) (string, error) {
	base, err := url.Parse(torr.Announce)
	if err != nil {
		return "", err
	}

	params := url.Values{
		"info_hash":  []string{string(torr.InfoHash[:])},
		"peer_id":    []string{string(peerId[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(torr.Length)},
	}

	base.RawQuery = params.Encode()

	return base.String(), nil
}

func (torr TorrentFile) requestPeers(peerId PeerId, port uint16) (tresp TrackerResponse, err error) {
	announceURL, err := torr.getAnnounceURL(peerId, port)
	if err != nil {
		return
	}

	resp, err := http.Get(announceURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var trackerResponse RawTrackerResponse
	err = bencode.Unmarshal(resp.Body, &trackerResponse)
	if err != nil {
		return
	}

	tresp, err = trackerResponse.parseTrackerResponse()
	return
}

func (tresp RawTrackerResponse) parseTrackerResponse() (TrackerResponse, error) {
	peers := []byte(tresp.Peers)
	numPeers := len(peers) / 6
	splitPeers := make([]Peer, numPeers)

	for i := range numPeers {
		offset := i * 6
		splitPeers[i].IP = net.IPv4(peers[offset], peers[offset+1], peers[offset+2], peers[offset+3])
		splitPeers[i].Port = binary.BigEndian.Uint16(peers[offset+4 : offset+6])
	}

	peers6 := []byte(tresp.Peers6)
	numPeers6 := len(peers6) / 18
	splitPeers6 := make([]Peer, numPeers6)

	for i := range numPeers6 {
		offset := i * 18
		splitPeers6[i].IP = net.IP(peers6[offset : offset+16])
		splitPeers6[i].Port = binary.BigEndian.Uint16(peers[offset+16 : offset+18])
	}

	return TrackerResponse{
		Interval: tresp.Interval,
		Peers:    splitPeers,
		Peers6:   splitPeers6,
	}, nil
}

func (p Peer) String() string {
	return fmt.Sprintf("%s:%d", p.IP.String(), p.Port)
}

// func (p Peer) getConnection() (err error) {
// 	conn, err := net.DialTimeout("tcp", p.String(), 3*time.Second)
// 	if err != nil {
// 		return
// 	}
// }
