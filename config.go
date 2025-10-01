package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

const peerIdPrefix = "-QB0001-"

func (p PeerId) MarshalYAML() (any, error) {
	return string(p[:]), nil
}

func (p PeerId) UnmarshalYAML(value *yaml.Node) error {
	if len(value.Value) != 20 {
		return fmt.Errorf("peerId must be 20 bytes long")
	}
	copy(p[:], value.Value)
	return nil
}

func defaultConfig() (qc QubeConfig, err error) {
	peerId, err := generateRandomPeerId()
	if err != nil {
		return
	}

	qc = QubeConfig{
		PeerId: peerId,
		Port:   6886,
	}

	return
}

func createConfig(qcFilepath string) (qc QubeConfig, err error) {
	qc, err = defaultConfig()
	if err != nil {
		return
	}

	f, err := os.Create(qcFilepath)
	if err != nil {
		return
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	defer encoder.Close()

	err = encoder.Encode(qc)
	return
}

func loadConfig(qcFilepath string) (qc QubeConfig, err error) {
	qcFile, err := os.ReadFile(qcFilepath)
	if err != nil {
		return createConfig(qcFilepath)
	}

	err = yaml.Unmarshal(qcFile, &qc)
	if err != nil {
		return
	}

	return
}

func generateRandomPeerId() (PeerId, error) {
	var peerId PeerId

	copy(peerId[:len(peerIdPrefix)], []byte(peerIdPrefix))

	_, err := io.ReadFull(rand.Reader, peerId[8:])
	if err != nil {
		return peerId, err
	}

	for i := len(peerIdPrefix); i < 20; i++ {
		charset := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
		peerId[i] = charset[int(peerId[i])%len(charset)]
	}

	return peerId, nil
}
