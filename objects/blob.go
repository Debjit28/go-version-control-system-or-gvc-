package objects

import (
	"bytes"
	"compress/zlib"
	"crypto/sha3"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateBlob(filePath string) (string, error) {
	//here i will try to implement sha 3
	//sha 3 is cryptographic algorithm which essentially uses a 64 charectser for hashing
	// why as significantly high resistance to collision

	//This will function will read the source file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file : %w", err)
	}

	//Making custom headers for the file

	header := fmt.Sprintf("blob %d\x00", len(content))
	fullData := append([]byte(header), content...)

	//applying hash algorithm
	hashBytes := sha3.Sum256(fullData)
	hash := hex.EncodeToString(hashBytes[:])

	dir := filepath.Join(".git", "objects", hash[:2])
	path := filepath.Join(dir, hash[:2])

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory : %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	zw := zlib.NewWriter(file)

	if _, err := zw.Write(fullData); err != nil {
		return "", err
	}
	zw.Close()

	return hash, nil

}

func ReadBlob(hash string) ([]byte, error) {
	dir := hash[:2]
	file := hash[2:]
	path := filepath.Join(".git", "objects", dir, file)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r, err := zlib.NewReader(f)

	if err != nil {
		return nil, err
	}

	defer r.Close()
	data, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	nullIndex := bytes.IndexByte(data, 0)

	if nullIndex < 0 {
		return nil, fmt.Errorf("Invalid blob")
	}

	return data[nullIndex+1:], nil
}
