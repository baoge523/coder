package tools

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// load secure file

type Account struct {
	Address string `json:"address"`
	SK      string `json:"sk"`
	Desc    string `json:"desc"`
}

type SecureFile struct {
	EthSepoliaURL string    `json:"eth-sepolia-url"`
	Accounts      []Account `json:"accounts"`
}

func LoadSecureFile() (*SecureFile, error) {

	path := "/Users/yongzhao/workspace/coder/web3/web3-go/conf/secure.json"
	var secureFile SecureFile
	err := loadFile(path, &secureFile)
	if err != nil {
		return nil, err
	}
	return &secureFile, nil
}

func loadFile[T interface{}](path string, t *T) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	var fileContent []byte
	for scanner.Scan() {
		scannedBytes := scanner.Bytes()
		fileContent = append(fileContent, scannedBytes...)
	}
	err = json.Unmarshal(fileContent, t)
	if err != nil {
		return fmt.Errorf("json unmarshal error: %w", err)
	}
	return nil
}
