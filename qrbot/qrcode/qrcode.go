package qrcode

//Генерация QR-кода. Возвращает путь к временному файлу.

import (
	"bytes"
	"image/png"
	"os"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(text string) (string, error) {
	qrCode, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	img, err := png.Decode(bytes.NewReader(qrCode))
	if err != nil {
		return "", err
	}

	filePath := "qrcode.png"
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
