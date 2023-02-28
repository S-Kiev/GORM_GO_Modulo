package main

import (
	"github.com/S-Kiev/GORM_GO_Modulo/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)
}
