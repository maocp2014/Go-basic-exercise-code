/*
问题：
还记得之前编写的图片生成器吗？现在来另外编写一个，不过这次将会返回image.Image来代替slice的数据。
自定义的Image类型，要实现必要的方法，并且调用pic.ShowImage。

Bounds应当返回一个image.Rectangle，例如 `image.Rect(0, 0, w, h)`。
ColorModel应当返回color.RGBAModel。

At 应当返回一个颜色；在这个例子里，在最后一个图片生成器的值 v 匹配 `color.RGBA{v, v, 255, 255}`。

提示：

官方 image的结构
type Image interface {
	// 颜色模式
	 ColorModel() color.Model
	// 图片边界
	Bounds() Rectangle
	// 某个点的颜色
	At(x, y int) color.Color
}
重点是实现 Image的方法，这样就可以自己使用了
*/ 

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

// 定义空结构体
type Image struct{}

// 实现系统Image接口的3个方法
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	// 宽和高写死了(简单展示)，应该从i中获取  
	return image.Rect(0, 0, 200, 200)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}

func main() {
// 可以自己设置宽高,传递进去
	m := Image{}
	// 3 调用
	pic.ShowImage(m)
}