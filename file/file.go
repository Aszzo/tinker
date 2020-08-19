package file

import (
	"fmt"
	"github.com/h2non/filetype"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)
// 获取文件类型
func GetFileContentType(filename string) (string)  {
	buf, _  := ioutil.ReadFile(filename)
	kind, _ := filetype.Match(buf)

	if kind == filetype.Unknown {
		return "Unknown"
	}
	return kind.MIME.Value
}
// 判断是否是文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return s.IsDir()
}
// 压缩jpg
func ResizeJpg(path string, quality int) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	// decode png into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Printf("decode fail:")
		log.Println(err)
	}
	defer file.Close()

	m := resize.Resize(0, 0, img, resize.NearestNeighbor)

	out, err := os.Create(path)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()

	_ = jpeg.Encode(out, m, &jpeg.Options{Quality: quality})
	defer func() {
		size := GetFileSize(path)
		fmt.Printf("压缩成功，压缩后文件大小%.2f kb \n", size)
	}()
}
// 压缩png
func ResizePng(path string, quality int, filename string)  {
	str, _ := os.Executable()
	dir := filepath.Dir(str)
	qualityString := strconv.Itoa(quality)
	cmd := exec.Command(dir+ "/pngquant", path, "--ext=.png", "--force", "--quality", qualityString)
	//cmd := exec.Command("./pngquant", path, "--ext=.png", "--force", "--quality", qualityString)
	defer func() {
		size := GetFileSize(path)
		fmt.Printf("压缩成功，压缩后文件大小%.2f kb \n", size)
	}()
	if err := cmd.Run(); err != nil {   // 运行命令
		log.Println(err)
	}
}
// 遍历目录下的所有文件，包含子目录
func RangeDir(path string, fileList *[]string) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	for _, fi := range dir{
		if fi.IsDir() {
			RangeDir(path + "/" + fi.Name(), fileList)
		} else {
			filePath := path + "/" + fi.Name()
			*fileList = append(*fileList, filePath)
		}
	}
}
// 读取文件大小
func GetFileSize(path string) float64  {
	fileInfo, _ := os.Stat(path)
	size := float64(fileInfo.Size()) / float64(1000)
	return size
}