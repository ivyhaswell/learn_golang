/*
练习：图像
还记得之前编写的图片生成器 吗？我们再来编写另外一个，不过这次它将会返回一个 image.Image 的实现而非一个数据切片。

定义你自己的 Image 类型，实现必要的方法并调用 pic.ShowImage。

Bounds 应当返回一个 image.Rectangle ，例如 image.Rect(0, 0, w, h)。

ColorModel 应当返回 color.RGBAModel。

At 应当返回一个颜色。上一个图片生成器的值 v 对应于此次的 color.RGBA{v, v, 255, 255}。
*/

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
	v    uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) ColorModal() color.Model {
	return color.RGBAModel
}

func (i Image) At() color.RGBA {
	return color.RGBA{i.v, i.v, 255, 255}
}

func main() {
	m := Image{100, 100, 100}
	pic.ShowImage(m)
}
