package filex

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// GetFileMd5 MD5 摘要
func GetFileMd5(filePath string) (string, bool) {
	exists := FilePathExists(filePath)
	if !exists {
		return "", false
	}
	openFile, err := os.Open(filePath)
	if err != nil {
		return "", false
	}
	const bufferSize = 65536
	hash := md5.New()
	readFile := bufio.NewReader(openFile)
	for buf, reader := make([]byte, bufferSize), readFile; ; {
		num, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", false
		}
		hash.Write(buf[:num])
	}
	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, true
}

// GetFileSha1 sha1 摘要
func GetFileSha1(filePath string) (string, bool) {
	exists := FilePathExists(filePath)
	if !exists {
		return "", false
	}
	openFile, err := os.Open(filePath)
	if err != nil {
		return "", false
	}
	const bufferSize = 65536
	hash := sha1.New()
	readFile := bufio.NewReader(openFile)
	for buf, reader := make([]byte, bufferSize), readFile; ; {
		num, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", false
		}
		hash.Write(buf[:num])
	}
	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, true
}

// GetFileSha256 sha256 摘要
func GetFileSha256(filePath string) (string, bool) {
	exists := FilePathExists(filePath)
	if !exists {
		return "", false
	}
	openFile, err := os.Open(filePath)
	if err != nil {
		return "", false
	}
	const bufferSize = 65536
	hash := sha256.New()
	readFile := bufio.NewReader(openFile)
	for buf, reader := make([]byte, bufferSize), readFile; ; {
		num, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", false
		}
		hash.Write(buf[:num])
	}
	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, true
}

// FilePathExists Judge whether the given path file / folder exists.
func FilePathExists(filePath string) bool {
	fileAbsPath, _ := filepath.Abs(filePath)
	_, err := os.Lstat(fileAbsPath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
