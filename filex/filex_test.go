package filex

import (
	"testing"
)

func TestFile(t *testing.T) {
	t.Run("GetFileMd5", func(t *testing.T) {
		filePath := "./filex.go"
		md5, ok := GetFileMd5(filePath)
		if ok {
			t.Log(md5)
		} else {
			t.Log("The file or path does not exist")
		}
	})

	t.Run("GetFileSha1", func(t *testing.T) {
		filePath := "./filex.go"
		sha1, ok := GetFileSha1(filePath)
		if ok {
			t.Log(sha1)
		} else {
			t.Log("The file or path does not exist")
		}
	})

	t.Run("GetFileSha256", func(t *testing.T) {
		filePath := "./filex.go"
		sha256, ok := GetFileSha256(filePath)
		if ok {
			t.Log(sha256)
		} else {
			t.Log("The file or path does not exist")
		}
	})
}
