package main

import (
	"flag"
	"fmt"
	f "github.com/Aszzo/tinker/file"
	"log"
)
var (
	i string
	q int
)

func init()  {
	flag.StringVar(&i, "i", "", "目标文件所在的目录|目标文件路径")
	flag.IntVar(&q, "q", 80, "压缩图片的质量（0-100）")
}
func main()  {
	flag.Parse()

	var fileList []string
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
}
func resize(path string)  {
	mime, err := f.GetFileContentType(path)
	if err != nil {
		log.Fatal(err)
	}

	switch mime {
	case "image/jpeg":
		fmt.Println("正在压缩文件：", path)
		f.ResizeJpg(path, q)
	case "image/png":
		fmt.Println("正在压缩文件：", path)
		f.ResizePng(path, q)
	default:
		fmt.Println(path + "不是jpeg或者png格式的文件")
	}
}
