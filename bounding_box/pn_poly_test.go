package math_util

import (
	"log"
	"math"
	"testing"
)

func TestUpdateBoundingBox(t *testing.T) {
	//视觉矩形
	initPoints1 := make([]Point, 0)
	BottomLeft1 := Point{X: 109.874267578125, Y: 25.393660521998022}
	initPoints1 = append(initPoints1, BottomLeft1)
	TopRight1 := Point{X: 112.467041015625, Y: 26.686729520004036}
	initPoints1 = append(initPoints1, TopRight1)
	boundingBox1 := GetBoundingBox(initPoints1)
	//相交矩形
	initPoints2 := make([]Point, 0)
	BottomLeft2 := Point{X: 112.0880126953125, Y: 24.72188526321623}
	initPoints2 = append(initPoints2, BottomLeft2)
	TopRight2 := Point{X: 113.1976318359375, Y: 25.512700007620513}
	initPoints2 = append(initPoints2, TopRight2)
	boundingBox2 := GetBoundingBox(initPoints2)
	//不相交矩形
	//initPoints2 := make([]Point, 0)
	//BottomLeft2 := Point{X: 110.27526855468749, Y: 26.88777988202911}
	//initPoints2 = append(initPoints2, BottomLeft2)
	//TopRight2 := Point{X: 111.676025390625, Y: 27.425414052729582}
	//initPoints2 = append(initPoints2, TopRight2)
	//boundingBox2 := GetBoundingBox(initPoints2)
	if boundingBox1.BottomLeft.X > boundingBox2.TopRight.X ||
		boundingBox1.BottomLeft.Y > boundingBox2.TopRight.Y ||
		boundingBox1.TopRight.X < boundingBox2.BottomLeft.X ||
		boundingBox1.TopRight.Y < boundingBox2.BottomLeft.Y {
		log.Println("-----")
	} else {
		log.Println("++++++++++++++")
	}
}

func TestUpdateBoundingBox2(t *testing.T) {
	//视觉矩形
	initPoints1 := make([]Point, 0)
	BottomLeft1 := Point{X: 109.874267578125, Y: 25.393660521998022}
	initPoints1 = append(initPoints1, BottomLeft1)
	TopRight1 := Point{X: 112.467041015625, Y: 26.686729520004036}
	initPoints1 = append(initPoints1, TopRight1)
	boundingBox1 := GetBoundingBox(initPoints1)
	//相交矩形
	initPoints2 := make([]Point, 0)
	BottomLeft2 := Point{X: 112.0880126953125, Y: 24.72188526321623}
	initPoints2 = append(initPoints2, BottomLeft2)
	TopRight2 := Point{X: 113.1976318359375, Y: 25.512700007620513}
	initPoints2 = append(initPoints2, TopRight2)
	boundingBox2 := GetBoundingBox(initPoints2)
	//不相交矩形
	//initPoints2 := make([]Point, 0)
	//BottomLeft2 := Point{X: 110.27526855468749, Y: 26.88777988202911}
	//initPoints2 = append(initPoints2, BottomLeft2)
	//TopRight2 := Point{X: 111.676025390625, Y: 27.425414052729582}
	//initPoints2 = append(initPoints2, TopRight2)
	//boundingBox2 := GetBoundingBox(initPoints2)
	//力扣算法判断相交矩形
	w1 := boundingBox1.TopRight.X - boundingBox1.BottomLeft.X
	h1 := boundingBox1.TopRight.Y - boundingBox1.BottomLeft.Y
	w2 := boundingBox2.TopRight.X - boundingBox2.BottomLeft.X
	h2 := boundingBox2.TopRight.Y - boundingBox2.BottomLeft.Y
	w := math.Abs((boundingBox1.BottomLeft.X+boundingBox1.TopRight.X)/2 - (boundingBox2.BottomLeft.X+boundingBox2.TopRight.X)/2)
	h := math.Abs((boundingBox1.BottomLeft.Y+boundingBox1.TopRight.Y)/2 - (boundingBox2.BottomLeft.Y+boundingBox2.TopRight.Y)/2)
	if w < (w1+w2)/2 && h < (h1+h2)/2 {
		log.Println("-------------")
	} else {
		log.Println("++++++++++++++")
	}
}
