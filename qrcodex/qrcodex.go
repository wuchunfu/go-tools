package qrcodex

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// GenerateDefaultQrCode 生成一个默认大小的中间带 logo 的二维码
func GenerateDefaultQrCode(content string, logoPath string, savePath string) {
	// 二维码大小
	qrCodeSize := 256
	// 百分比
	percent := 25
	// logo 宽高大小
	size := uint(float64(qrCodeSize) * float64(percent) / 100)
	GenerateQrCode(content, false, logoPath, savePath, qrCodeSize, size, size)
}

// GenerateQrCode 生成一个中间带icon的二维码
// content 二维码资源
// logoPath logo 图片地址
// savePath 图片的保存地址
// qrCodeSize 二维码图片大小
// logoWidth 重置 logo 图片宽度
// logoHeight 重置 logo 图片高度
func GenerateQrCode(content string, borderDisable bool, logoPath string, savePath string, qrCodeSize int, logoWidth uint, logoHeight uint) {
	if len(content) == 0 {
		fmt.Printf("%s :it is empty!\n", content)
		return
	}
	logoExists := FilePathExists(logoPath)
	if !logoExists {
		fmt.Printf("Logo path not exists: %s\n", logoPath)
		return
	}
	outAbs, err := filepath.Abs(savePath)
	if err != nil {
		fmt.Printf("Get save file abs path failed：%s", err.Error())
		return
	}
	fileDir := filepath.Dir(outAbs)
	saveExists := FilePathExists(logoPath)
	if !saveExists {
		fmt.Printf("%s :Save path not exists!\n", savePath)
		MkdirAll(fileDir)
		fmt.Printf("%s :create successfully!\n", fileDir)
	}

	qrCode, err := qrcode.New(content, qrcode.Highest)
	if err != nil {
		fmt.Printf("Create qrcode failed: %v\n", err)
	}
	qrCode.DisableBorder = borderDisable
	qrCodeImg := qrCode.Image(qrCodeSize)

	logoImg, resizeErr := AddLogo(logoPath, qrCodeImg, logoWidth, logoHeight)
	if resizeErr != nil {
		fmt.Printf("Resize failed: %v\n", resizeErr)
	}
	saveErr := SaveImage(savePath, logoImg)
	if saveErr != nil {
		fmt.Printf("Save image failed: %v\n", saveErr)
	}
}

// AddLogo 重置图片大小
// logoPath logo 图片地址
// qrCodeImg 二维码图片
// width 重置 logo 图片宽度
// height 重置 logo 图片高度
func AddLogo(logoPath string, qrCodeImg image.Image, width uint, height uint) (*image.RGBA, error) {
	avatarImg, err := ResizeLogo(logoPath, width, height)
	if err != nil {
		return nil, err
	}
	// 得到 icon 图片的大小
	logoSize := avatarImg.Bounds()
	// 得到背景图的大小
	qrCodeSize := qrCodeImg.Bounds()
	// 居中设置icon到二维码图片
	offset := image.Pt((qrCodeSize.Dx()-logoSize.Dx())/2, (qrCodeSize.Dy()-logoSize.Dy())/2)
	imgRgba := image.NewRGBA(qrCodeSize)
	draw.Draw(imgRgba, qrCodeSize, qrCodeImg, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(imgRgba, avatarImg.Bounds().Add(offset), avatarImg, image.Point{X: 0, Y: 0}, draw.Over)
	return imgRgba, nil
}

// ResizeLogo 修改图片的大小
// logoPath logo 图片地址
// width 重置 logo 图片宽度
// height 重置 logo 图片高度
func ResizeLogo(logoPath string, width uint, height uint) (image.Image, error) {
	file, err := os.Open(logoPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	avatarImg, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	avatarImg = resize.Resize(width, height, avatarImg, resize.Lanczos3)
	return avatarImg, nil
}

// ImageType 判断图片类型
func ImageType(iconPath string) (image.Image, error) {
	avatarFile, err := os.Open(iconPath)
	if err != nil {
		return nil, err
	}
	ext := filepath.Ext(iconPath)
	var avatarImg image.Image
	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
		avatarImg, err = jpeg.Decode(avatarFile)
	} else if strings.EqualFold(ext, ".png") {
		avatarImg, err = png.Decode(avatarFile)
	} else if strings.EqualFold(ext, ".gif") {
		avatarImg, err = gif.Decode(avatarFile)
	}
	return avatarImg, err
}

// SaveImage 保存图片
// savePath 图片的保存地址
// logoImg logo 图片
func SaveImage(savePath string, logoImg image.Image) error {
	saveFile, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer saveFile.Close()
	ext := filepath.Ext(savePath)
	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
		err = jpeg.Encode(saveFile, logoImg, &jpeg.Options{Quality: 100})
	} else if strings.EqualFold(ext, ".png") {
		err = png.Encode(saveFile, logoImg)
	} else if strings.EqualFold(ext, ".gif") {
		err = gif.Encode(saveFile, logoImg, &gif.Options{NumColors: 256})
	}
	return err
}

// FilePathExists Judge whether the given path file / folder exists.
func FilePathExists(filePath string) bool {
	_, err := os.Lstat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// MkdirAll 递归创建目录
func MkdirAll(filePath string) bool {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}
