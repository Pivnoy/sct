package internal

import (
	"gocv.io/x/gocv"
	"log"
)

type VideoService struct{}

func (vs *VideoService) SetVecbAt(m *gocv.Mat, row int, col int, v gocv.Vecb) {
	ch := m.Channels()
	for c := 0; c < ch; c++ {
		m.SetUCharAt(row, col*ch+c, v[c])
	}
}

func (vs *VideoService) CleanVideo(vc *gocv.VideoCapture, vw *gocv.VideoWriter) {
	img := gocv.NewMat()
	dst := gocv.NewMat()
	ok := vc.Read(&img)
	for ok {
		if img.Empty() {
			log.Fatal("Empty frame")
		}
		img.CopyTo(&dst)
		for i := 0; i < img.Rows(); i++ {
			for j := 0; j < img.Rows(); j++ {
				if img.GetVecbAt(i, j)[0] > 50 {
					vs.SetVecbAt(&dst, i, j, gocv.Vecb{255, 255, 255})
				}
			}
		}
		err := vw.Write(dst)
		if err != nil {
			log.Fatal("Can not write frame")
		}
		ok = vc.Read(&img)
	}
}
