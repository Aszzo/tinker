package file

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
// 计算文件哈希值
func GetFileHash(file *os.File) string {
	// 计算文件哈希值
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Println(err)
		return ""
	}
	sum := hash.Sum([]byte(""))
	// file.seek(0,0)方法，将文件指针移到开头,
	// https://blog.csdn.net/weixin_40306555/article/details/102779246
	file.Seek(0,0)
	fileHash := new(bytes.Buffer)
	_, hashErr :=fmt.Fprintf(fileHash,"%x", sum)
	if hashErr != nil {
		return ""
	}
	return fileHash.String()
}
// 图片是否被压缩过
func IsResize(file *os.File, quality string)  bool {
	fileHash := GetFileHash(file)
	// 读取.tinker目录下的resize.log信息
	// 里面会存储压缩的历史记录
	isExitFile := isExits(".tinker")
	if !isExitFile {
		os.Mkdir(".tinker", 0777)
	}

	logName := ".tinker/resize.log"
	f, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return false
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return false
	}
	containsHash := strings.Contains(string(content), fileHash + "-" + quality)
	// 如果有压缩日志，直接返回
	if containsHash {
		return  containsHash
	}
	return false
}
// 写入hash
func WriteHash(path string, quality string) error {
	file, err := os.Open(path)
	if err !=nil {
		return err
	}
	fileHash := GetFileHash(file)
	logName := ".tinker/resize.log"
	f, err := os.OpenFile(logName, os.O_WRONLY|os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()
	// 没有压缩日志，返回false,并且创建新的日志
	write := bufio.NewWriter(f)
	write.WriteString(fileHash + "-" + quality + "\n")
	writeErr := write.Flush()
	if writeErr != nil {
		log.Println(writeErr)
		return writeErr
	}
	return nil
}
