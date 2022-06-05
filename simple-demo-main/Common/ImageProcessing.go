package Common

import (
	"bufio"
	"gocv.io/x/gocv"
	"image"
	"image/png"
	"os"
)

func GetVideoMoment(filePath string, time float64) (i image.Image, err error) {
	//load video
	vc, err := gocv.VideoCaptureFile(filePath)
	if err != nil {
		return i, err
	}

	frames := vc.Get(gocv.VideoCaptureFrameCount)
	fps := vc.Get(gocv.VideoCaptureFPS)
	duration := frames / fps

	frames = (time / duration) * frames

	// Set Video frames
	vc.Set(gocv.VideoCapturePosFrames, frames)

	img := gocv.NewMat()

	vc.Read(&img)

	imageObject, err := img.ToImage()
	if err != nil {
		return i, err
	}
	return imageObject, err
}

func SaveCover(videoPath string, fileName string, savePath string) bool {
	img, err := GetVideoMoment(videoPath, 0) //"./public/2_周杰伦-暗号"
	if err != nil {
		panic(err)
		return false
	}

	outFile, err := os.Create(savePath + fileName)
	defer outFile.Close()
	if err != nil {
		panic(err)
		return false
	}

	b := bufio.NewWriter(outFile)
	err = png.Encode(b, img)
	if err != nil {
		panic(err)
		return false
	}

	err = b.Flush()
	if err != nil {
		panic(err)
		return false
	}
	return true
}
