package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/kbinani/screenshot"
)

func capture() {

	// get current path
	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	exPath := filepath.Dir(ex)

	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		now := time.Now()
		dateString := fmt.Sprintf("%d%02d%02d-%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		fileName := fmt.Sprintf("%s-%d.png", dateString, i)

		writeFilePath := filepath.Join(exPath, fileName)

		file, _ := os.Create(writeFilePath)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("write file into path: %s", writeFilePath)
		// fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
}

func main() {
	capture()
}
