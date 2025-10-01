package qubetorrent

import "net"

type PeerId [20]byte

type QubeConfig struct {
	PeerId PeerId `yaml:"peerId"`
	Port   int    `yaml:"port"`
}

type MetainfoFile_File struct {
	Path   string `bencode:"path" yaml:"path"`
	Length int    `bencode:"length" yaml:"length"`
	MD5Sum string `bencode:"md5sum,omitempty" yaml:"md5sum,omitempty"`
}

type MetainfoFile_Info struct {
	PieceLength int    `bencode:"piece length" yaml:"pieceLength"`
	Pieces      string `bencode:"pieces" yaml:"pieces"`

	Name   string              `bencode:"name" yaml:"name"`
	Files  []MetainfoFile_File `bencode:"files,omitempty" yaml:"file,omitempty"`
	Length int                 `bencode:"length,omitempty" yaml:"length,omitempty"`
	MD5Sum string              `bencode:"md5sum,omitempty" yaml:"md5sum,omitempty"`

	Private int `bencode:"private,omitempty" yaml:"private,omitempty"`
}

type MetainfoFile struct {
	Announce string            `bencode:"announce" yaml:"announce"`
	Info     MetainfoFile_Info `bencode:"info" yaml:"info"`

	AnnounceList [][]string `bencode:"announce-list,omitempty" yaml:"announceList,omitempty"`
	CreationDate int        `bencode:"creation date,omitempty" yaml:"creationDate,omitempty"`
	Comment      string     `bencode:"comment,omitempty" yaml:"comment,omitempty"`
	CreatedBy    string     `bencode:"created by,omitempty" yaml:"createdBy,omitempty"`
	Encoding     string     `bencode:"encoding,omitempty" yaml:"encoding,omitempty"`
}

type TorrentFile struct {
	Announce   string     `yaml:"announce"`
	InfoHash   [20]byte   `yaml:"infoHash"`
	PiecesHash [][20]byte `yaml:"piecesHash"`
	Length     int        `yaml:"length"`
}

type RawTrackerResponse struct {
	Interval int    `bencode:"interval" yaml:"interval"`
	Peers    string `bencode:"peers" yaml:"peers"`
	Peers6   string `bencode:"peers6" yaml:"peers6"`
}

type TrackerResponse struct {
	Interval int    `yaml:"interval"`
	Peers    []Peer `yaml:"peers"`
	Peers6   []Peer `yaml:"peers6"`
}

type Peer struct {
	IP   net.IP `yaml:"ip"`
	Port uint16 `yaml:"port"`
}
