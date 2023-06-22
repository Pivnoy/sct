package internal

import (
	"gocv.io/x/gocv"
	"log"
	"testing"
)

func BenchmarkVideoService_CleanVideo(b *testing.B) {
	vs := &VideoService{}
	vc, err := gocv.OpenVideoCapture("../outcpp.avi")
	defer vc.Close()
	if err != nil {
		log.Fatal(err)
	}
	videoWriter, err := gocv.VideoWriterFile("./new.avi", "X264", 25, 120, 120, true)
	if err != nil {
		log.Fatal(err)
	}
	defer videoWriter.Close()
	vs.CleanVideo(vc, videoWriter)
}
