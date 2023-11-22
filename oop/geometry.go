package geometry

import "math"

type Point struct{ X, Y float64 }


// 上面的代码里那个附加的参数p，叫做方法的接收器（receiver），
// 早期的面向对象语言留下的遗产将调用一个方法称为“向一个对象发送消息”。
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
