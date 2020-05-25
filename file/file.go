package file

import (
	"errors"
	"fmt"
	"github.com/h2non/filetype"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)
// 获取文件类型
func GetFileContentType(filename string) (string, error)  {
	buf, _  := ioutil.ReadFile(filename)
	kind, _ := filetype.Match(buf)

	if kind == filetype.Unknown {
		return "", errors.New("Unknown file type")
	}
	return kind.MIME.Value, nil
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
		log.Fatal(err)
	}
	// decode png into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Printf("decode fail:")
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(0, 0, img, resize.NearestNeighbor)

	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, &jpeg.Options{Quality: quality})
}
// 压缩png
func ResizePng(path string, quality int)  {
	qualityString := strconv.Itoa(quality)
	cmd := exec.Command("./pngquant", path, "--ext=.png", "--force", "--quality", qualityString)
	if err := cmd.Start(); err != nil {   // 运行命令
		log.Fatal(err)
	}
}
// 遍历目录下的所有文件，包含子目录
func RangeDir(path string, fileList *[]string) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
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