package main

import (
	"flag"
	"fmt"
	f "github.com/Aszzo/tinker/file"
	"log"
	"os"
	p "path"
)
var (
	i string
	q int
)

func init()  {
	flag.IntVar(&q, "q", 80, "压缩图片的质量（0-100）")
}
func main()  {
	flag.Parse()
	var fileList []string
	var i string

	if len(os.Args) > 1{
		i = os.Args[1]

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
	mime, err := f.GetFileContentType(path)
	if err != nil {
		log.Fatal(err)
	}
	filenameWithSuffix := p.Base(path) //获取文件名带后缀

	switch mime {
		case "image/jpeg":
			size := getFileSize(path)
			fmt.Printf("压缩文件：%v %.2f kb \n", p.Base(path), size)
			f.ResizeJpg(path, q)

		case "image/png":
			size := getFileSize(path)
			fmt.Printf("压缩文件：%v %.2f kB \n", p.Base(path), size)
			f.ResizePng(path, q, filenameWithSuffix)
		default:
			fmt.Println(path + "不是jpeg或者png格式的文件")
	}
}
// 读取文件大小
func getFileSize(path string) float64  {
	fileInfo, _ := os.Stat(path)
	size := float64(fileInfo.Size()) / float64(1000)
	return size
}
