package main

import (
	"gocv.io/x/gocv"
	"log"
	"os"
	"stc/internal"
)

func SetVecbAt(m *gocv.Mat, row int, col int, v gocv.Vecb) {
	ch := m.Channels()
	for c := 0; c < ch; c++ {
		m.SetUCharAt(row, col*ch+c, v[c])
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Videofile argument is required")
	}
	vs := &internal.VideoService{}
	vc, err := gocv.OpenVideoCapture(args[1])
	defer vc.Close()
	if err != nil {
		log.Fatal(err)
	}
	videoWriter, err := gocv.VideoWriterFile("edited_" + args[1], "X264", 25, 120, 120, true)
	if err != nil {
		log.Fatal(err)
	}
	defer videoWriter.Close()
	vs.CleanVideo(vc, videoWriter)
}
