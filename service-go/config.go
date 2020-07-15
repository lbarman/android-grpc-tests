package servicego

import (
	"fmt"
	"bytes"
	"go.dedis.ch/onet/log"
	"golang.org/x/mobile/asset"
)

const CONFIG_FILENAME = "config.toml"

func LoadConfig() error {
	fmt.Println("Reading config file", CONFIG_FILENAME)
	_, _ = readTomlFromAssets(CONFIG_FILENAME)

	//...
	return nil
}

func readTomlFromAssets(filename string) (string, error) {
	file, err := asset.Open(filename)
	defer file.Close()

	if err != nil {
		log.Error("Could not open file ", filename)
		return "", err
	}

	tomlRawDataBuffer := new(bytes.Buffer)
	_, err = tomlRawDataBuffer.ReadFrom(file)

	if err != nil {
		log.Error("Could not read file ", filename)
		return "", err
	}

	return tomlRawDataBuffer.String(), err
}

func GenerateNewKeyPairAndAssign() error {
	return nil
}

func GetPublicKey() (string, error) {
	return "", nil
}

func SetPublicKey(pubKey string) error {
	return nil
}

func GetPrivateKey() (string, error) {
	return "", nil
}

func SetPrivateKey(priKey string) error {
	return nil
}