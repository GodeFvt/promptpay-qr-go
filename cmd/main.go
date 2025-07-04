/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"image"
	"image/png"
	"os"

	promptpayqr "github.com/GodeFvt/promptpay-qr-go"
)

func main() {
	qr, err := promptpayqr.QRForTarget("09999999999")
	if err != nil {
		panic(err)
	}

	base64Img := base64.StdEncoding.EncodeToString(*qr)
	fmt.Println()
	fmt.Println("Base64 Encoded Image:", base64Img)

	byteToImage(*qr)

}

func byteToImage(imgByte []byte) {
	img, _, _ := image.Decode(bytes.NewReader(imgByte))

	//save the imgByte to file
	out, err := os.Create("./QRImg.png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, img)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
