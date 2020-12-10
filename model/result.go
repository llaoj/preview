package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Result struct {
	Model

	HashKey string `json:"hash_key"`
	Ext     string `json:"ext"`
	OssPath string `json:"oss_path"`
}

func (r *Result) TableName() string {
	return "oconv_result"
}

func (r *Result) Get(hash, ext string) {
	log.Println("[INFO] getting from db: ", hash)
	errors := db.Where("hash_key = ?", hash).Where("ext = ?", ext).First(r).GetErrors()
	for _, err := range errors {
		if !gorm.IsRecordNotFoundError(err) {
			log.Println("[ERROR] ", err)
		}
	}
}

func (r *Result) Create() {
	log.Println("[INFO] inserting into db: ", r.HashKey)
	errors := db.Create(r).GetErrors()
	for _, err := range errors {
		log.Println("[ERROR] ", err)
	}
}
