package main

import (
	modelo "github.com/S-Kiev/GORM_GO_Modulo/modelos"
	"github.com/S-Kiev/GORM_GO_Modulo/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	storage.DB().AutoMigrate(
		&modelo.Producto{},
		&modelo.EncabezadoFactura{},
		&modelo.ItemFactura{},
	)
}
