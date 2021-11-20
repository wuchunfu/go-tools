package qrcodex

import "testing"

func TestQrcode(t *testing.T) {
	t.Run("GenerateDefaultQrCode", func(t *testing.T) {
		url := "http://www.baidu.com/"
		iconPath := "./images/icon.jpeg"
		savePath := "./images/show.png"
		GenerateDefaultQrCode(url, iconPath, savePath)
	})
}
