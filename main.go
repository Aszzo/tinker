package main

import (
	"flag"
	"fmt"
	f "github.com/Aszzo/tinker/file"
	p "path"
)
var (
	q int
	i string
)


func main()  {
	flag.StringVar(&i, "i", "", "目标文件夹|目标文件")
	flag.IntVar(&q, "q", 80, "压缩图片的质量（0-100）")

	flag.Parse()

	var fileList []string
	//var i string
	//
	if i != ""{
	//	i = os.Args[1]

		// 首先判断是文件还是文件夹
		isDir := f.IsDir(i)

		if isDir {
			// 如果是文件夹，遍历文件夹,获取文件夹内的所有文件
			f.RangeDir(i, &fileList)

			for _, path := range fileList {
				resize(path)
			}
		} else {
			// 如果是图像文件直接压缩
			resize(i)
		}
	} else {
		flag.Usage()
	}
}
func resize(path string)  {
	mime := f.GetFileContentType(path)

	filenameWithSuffix := p.Base(path) //获取文件名带后缀
	switch mime {
		case "image/jpeg":
			size := f.GetFileSize(path)
			fmt.Printf("正在压缩图片：%v，图片大小： %.2f kb >>>>>>>>>>>>>>> ", p.Base(path), size)
			f.ResizeJpg(path, q)

		case "image/png":
			size := f.GetFileSize(path)
			fmt.Printf("正在压缩图片：%v，图片大小： %.2f kb >>>>>>>>>>>>>>> ", p.Base(path), size)
			f.ResizePng(path, q, filenameWithSuffix)
		default:
			fmt.Println(path + "不是jpeg或者png格式的文件")
	}
}
