package utils

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func FileHash(fileName string) (hash string, err error) {
	f, _ := os.Open(fileName)
	defer f.Close()
	br := bufio.NewReader(f)
	h := sha256.New()
	_, err = io.Copy(h, br)
	if err != nil {
		log.Println("[ERROR] ", err.Error())
		return
	}

	hash = fmt.Sprintf("%x", h.Sum(nil))
	return
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func FileCopy(r io.Reader, target string) error {
	tmp, err := os.Create(target)
	defer tmp.Close()
	log.Println("[INFO] creating file: " + target)
	if err != nil {
		log.Println("[ERROR] ", err.Error())
		return err
	}
	if _, err = io.Copy(tmp, r); err != nil {
		log.Println("[ERROR] ", err.Error())
		return err
	}

	return nil
}

func FileRemove(files []string) error {
	log.Printf("[INFO] deleting files: %v\n", files)
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			log.Println("[ERROR] ", err.Error())
			return err
		}
	}

	return nil
}

//在指定目录下查找文件名包含str的文件
func FileSearch(dir, str string) (s []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("[ERROR] ", err.Error())
		return
	}
	for _, f := range files {
		if !f.IsDir() && strings.Contains(f.Name(), str) {
			s = append(s, StringBuilder(dir, f.Name()))
		}
	}
	log.Printf("[INFO] found files: %v\n", s)

	return
}

func FileRead(file string, buf []byte) (str string, err error) {
	f, err := os.Open(file)
	if err != nil {
		log.Println("[ERROR] open file error, ", err.Error())
        return
    }
	defer f.Close()

	// 只读一次
    l, err := f.Read(buf)
    if err != nil && err != io.EOF {
		log.Println("[ERROR] read file error, ", err.Error())
        return
    }

    str = string(buf)
	log.Println("[INFO] first ", l, " bytes content: ", str)
    return
}
