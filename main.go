package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"oconv/config"
	"oconv/model"
	"oconv/utils"
)

func main() {
	time.Sleep(20 * time.Second)
	config.Setup()
	model.Setup()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
	http.HandleFunc("/conv", convHandler)
	log.Print("[INFO] server started on :9097")
	log.Fatal(http.ListenAndServe(":9097", nil))
}

func convHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "http method invalid", http.StatusBadRequest)
		return
	}

	// 限制请求的body最大32M
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out := r.FormValue("out")
	if out != "pdf" && out != "jpg" {
		http.Error(w, "output file format invalid", http.StatusBadRequest)
		return
	}

	// 默认32M的文件
	file, header, err := r.FormFile("file") //io.Reader
	defer file.Close()
	if err != nil {
		log.Println("[ERROR] ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[INFO] request file header: ", header.Header)

	// 检测文件格式
	ctype := header.Header.Get("Content-Type")
	log.Println("[INFO] content type: " + ctype)
	if ctype != "application/msword" && //doc
		ctype != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" && //docx
		ctype != "application/pdf" {
		http.Error(w, "content type invalid", http.StatusBadRequest)
		return
	}

	// 接收文件
	workdir := config.Get().WorkDir
	tmp := utils.StringBuilder(workdir, utils.RandNewStr(8))
	if err = utils.FileCopy(file, tmp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cleanTmp(tmp)
	
	if ctype != "application/pdf" {
		switch pass, err := fileFilter(tmp); {
		case err != nil:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		case !pass:
			http.Error(w, "content type invalid", http.StatusBadRequest)
			return
		}
	}

	// 文件哈希
	hash, err := utils.FileHash(tmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[INFO] file hash: " + hash)

	// 先查存储
	ossUrlPrefix := utils.StringBuilder("https://", config.Get().Oss.Domain, "/")
	var result model.Result
	result.Get(hash, out)
	if result.OssPath != "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]string{utils.StringBuilder(ossUrlPrefix, result.OssPath)})
		return
	}

	// remove files
	defer cleanFiles(workdir, hash)

	// to pdf
	pdf := utils.StringBuilder(workdir, hash, ".pdf")
	if ctype == "application/pdf" {
		if err = os.Rename(tmp, pdf); err != nil {
			log.Println("[ERROR] ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("[INFO] tmp rename to pdf: " + pdf)
	}
	if !utils.FileExist(pdf) {
		if err = utils.OfficeToPdf(tmp, pdf); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// to jpg
	jpg := utils.StringBuilder(workdir, hash, ".jpg")
	if out == "jpg" && !utils.FileExist(jpg) {
		if err = utils.PdfToJpg(pdf, jpg); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// to aliun oss
	filesToUpload := []string{pdf}
	if out == "jpg" {
		filesToUpload = append(filesToUpload, jpg)
	}
	log.Printf("[INFO] files to upload: %v\n", filesToUpload)
	ossPaths, err := utils.UploadToOss(filesToUpload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var rp []string
	for _, v := range ossPaths {
		ext := strings.Replace(path.Ext(v), ".", "", 1)
		// 存储到数据库
		result = model.Result{}
		result.HashKey = hash
		result.Ext = ext
		result.OssPath = v
		result.Create()
		// 创建响应
		if out == ext {
			rp = append(rp, utils.StringBuilder(ossUrlPrefix, v))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rp)
}

func cleanFiles(workdir, hash string) (err error) {
	filesToRm, err := utils.FileSearch(workdir, hash)
	if err != nil {
		log.Println("[ERROR] ", err.Error())
		return
	}
	if err = utils.FileRemove(filesToRm); err != nil {
		log.Println("[ERROR] ", err.Error())
		return
	}

	return
}

func cleanTmp(tmp string) (err error) {
	if utils.FileExist(tmp) {
		if err = utils.FileRemove([]string{tmp}); err != nil {
			return
		}
	}
	return
}

func fileFilter(file string) (pass bool, err error) {
	buf := make([]byte, 200)
	var str string
	str, err = utils.FileRead(file, buf)
	if err != nil {
		return
	}

	pass = true
	s := []string{"multipart/related", "xmlns:pkg"}
	for _, v := range s {
		if strings.Contains(str, v) {
			pass = false
		}
	}

	return
}
