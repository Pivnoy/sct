package main

import (
	"gocv.io/x/gocv"
	"log"
	"stc/internal"
)

func SetVecbAt(m *gocv.Mat, row int, col int, v gocv.Vecb) {
	ch := m.Channels()
	for c := 0; c < ch; c++ {
		m.SetUCharAt(row, col*ch+c, v[c])
	}
}

func main() {
	vs := &internal.VideoService{}
	vc, err := gocv.OpenVideoCapture("outcpp.avi")
	defer vc.Close()
	if err != nil {
		log.Fatal(err)
	}
	videoWriter, err := gocv.VideoWriterFile("./captured_video.avi", "X264", 25, 120, 120, true)
	if err != nil {
		log.Fatal(err)
	}
	defer videoWriter.Close()
	vs.CleanVideo(vc, videoWriter)
}
