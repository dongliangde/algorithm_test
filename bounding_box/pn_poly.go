package math_util

//矩形坐标点
type Bounding struct {
	BottomLeft Point //左下角
	TopRight   Point //右上方
}

//点坐标
type Point struct {
	X float64 `json:"lng"` //经度
	Y float64 `json:"lat"` //纬度
}

/**
 * @Function: PointInBoundingBox
 * @Description: 判断当前坐标是否存在矩形内
 */
func PointInBoundingBox(pt Point, bb Bounding) bool {
	return pt.X < bb.TopRight.X && pt.X > bb.BottomLeft.X &&
		pt.Y < bb.TopRight.Y && pt.Y > bb.BottomLeft.Y
}

/**
 * @Function: GetBoundingBox
 * @Description: 获取矩形坐标点
 */
func GetBoundingBox(points []Point) Bounding {
	var maxX, maxY, minX, minY float64
	for i := 0; i < len(points); i++ {
		side := points[i]
		if side.X > maxX || maxX == 0.0 {
			maxX = side.X
		}
		if side.Y > maxY || maxY == 0.0 {
			maxY = side.Y
		}
		if side.X < minX || minX == 0.0 {
			minX = side.X
		}
		if side.Y < minY || minY == 0.0 {
			minY = side.Y
		}
	}
	return Bounding{
		BottomLeft: Point{X: minX, Y: minY},
		TopRight:   Point{X: maxX, Y: maxY},
	}
}

/**
 * @Function: UpdateBoundingBox
 * @Description: 更新尾迹最新矩形坐标
 */
func SetBoundingBox(x float64, y float64, box *Bounding) {
	if x > box.TopRight.X || box.TopRight.X == 0.0 {
		box.TopRight.X = x
	}
	if y > box.TopRight.Y || box.TopRight.Y == 0.0 {
		box.TopRight.Y = y
	}
	if x < box.BottomLeft.X || box.BottomLeft.X == 0.0 {
		box.BottomLeft.X = x
	}
	if y < box.BottomLeft.Y || box.BottomLeft.Y == 0.0 {
		box.BottomLeft.Y = y
	}
}

/**
 * @Function: UpdateBoundingBox
 * @Description: 判断矩形是否相交
 */
func BoundingBox(borderBoundingBox Bounding, pointBoundingBox Bounding) bool {
	if borderBoundingBox.BottomLeft.X > pointBoundingBox.TopRight.X ||
		borderBoundingBox.BottomLeft.Y > pointBoundingBox.TopRight.Y ||
		borderBoundingBox.TopRight.X < pointBoundingBox.BottomLeft.X ||
		borderBoundingBox.TopRight.Y < pointBoundingBox.BottomLeft.Y {
		return false
	} else {
		return true
	}
}
