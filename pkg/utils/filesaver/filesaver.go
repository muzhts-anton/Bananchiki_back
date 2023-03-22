package filesaver

import (
	"banana/pkg/utils/log"

	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"os"
)

const (
	RootPath = "/home/banana/Bananchiki_back"
)

func UploadFile(reader io.Reader, path, ext string) (string, error) {
	randString, err := generateRandomString(6)
	if err != nil {
		return "", err
	}
	filename := randString + ext
	file, err := createFile(path, filename)
	if err != nil {
		return "", fmt.Errorf("file creating error: %s", err)
	}
	defer file.Close()
	log.Info("Created file with name " + filename)

	filename = path + filename
	_, err = io.Copy(file, reader)
	if err != nil {
		return "", fmt.Errorf("copy error: %s", err)
	}
	return filename, nil
}

func createFile(dir, name string) (*os.File, error) {
	_, err := os.ReadDir(RootPath + dir)
	if err != nil {
		err = os.MkdirAll(RootPath+dir, 0777)
		if err != nil {
			return nil, err
		}
	}
	file, err := os.Create(RootPath + dir + name)
	return file, err
}

func generateRandomString(n uint) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	var i uint
	for i < n {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
		i++
	}
	return string(ret), nil
}
