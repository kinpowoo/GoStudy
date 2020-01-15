package main

import (
	"hello/photo"
)

//可以设置别名，也用用 . 号来将包整体引入，直接调用包中的方法，
//如果没设置别名，也没用 . 号，要用 package 包名来调用包中的方法
//包中被调用的方法名首字母要大写

func main() {

	photo.GenSin()

	//go，模块中要导出的函数，必须首字母大写。
	photo.GenPhoto()
}
