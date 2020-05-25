package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	f "github.com/Aszzo/tinker/file"
)
var (
	h bool
	i string
	q int
)

func init()  {

	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&i, "i", "", "目标文件所在的目录|目标文件路径")
	flag.IntVar(&q, "q", 80, "压缩图片的质量（0-100）")
}
func usage()  {
	fmt.Fprint(os.Stderr, `version/1.0.0
Usage: tinker [-i input] [-q quality]
Options:
`)
flag.PrintDefaults()
}
func main()  {
	flag.Parse()
	if h {
		usage()
		return
	}
	var fileList []string
	// 首先判断是文件还是文件夹
	isDir := f.IsDir("file")

	if isDir {
		// 如果是文件夹，遍历文件夹,获取文件夹内的所有文件
		f.RangeDir(i, &fileList)
	}
	for _, path := range fileList{
		mime, err := f.GetFileContentType(path)
		if err != nil{
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
}
