package math_util

import (
	"math"
)

//点坐标
type Point struct {
	X float64 `json:"lng"` //经度
	Y float64 `json:"lat"` //纬度
}

// Line 代表一个线段
type Line struct {
	Start Point
	End   Point
}

// DistanceToPoint 返回点到线的垂直距离
func (l Line) DistanceToPoint(pt *Point) float64 {
	a, b, c := l.Coefficients()
	return math.Abs(a*pt.X+b*pt.Y+c) / math.Sqrt(a*a+b*b)
}

// Coefficients 返回定义一条线的三个系数
//// 一条线可以用下面的等式表示
func (l Line) Coefficients() (a, b, c float64) {
	a = l.Start.Y - l.End.Y
	b = l.End.X - l.Start.X
	c = l.Start.X*l.End.Y - l.End.X*l.Start.Y
	return a, b, c
}

// SimplifyPath接受一个点列表和epsilon作为阈值，通过删除来简化路径
// 没有通过阈值的点
func SimplifyPath(points []*Point, ep float64) []*Point {
	if len(points) <= 2 {
		return points
	}
	l := Line{Start: *points[0], End: *points[len(points)-1]}
	idx, maxDist := seekMostDistantPoint(l, points)
	if maxDist >= ep {
		left := SimplifyPath(points[:idx+1], ep)
		right := SimplifyPath(points[idx:], ep)
		return append(left[:len(left)-1], right...)
	}
	// 如果最远的点没有通过阈值测试，那么就返回两个点
	return []*Point{points[0], points[len(points)-1]}
}

//寻找最遥远的点
func seekMostDistantPoint(l Line, points []*Point) (idx int, maxDist float64) {
	for i := 0; i < len(points); i++ {
		d := l.DistanceToPoint(points[i])
		if d > maxDist {
			maxDist = d
			idx = i
		}
	}
	return idx, maxDist
}
