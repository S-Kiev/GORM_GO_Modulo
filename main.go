package main

import (
	"fmt"

	modelo "github.com/S-Kiev/GORM_GO_Modulo/modelos"
	"github.com/S-Kiev/GORM_GO_Modulo/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	//OBTENIENDO UN REGISTRO

	miProducto := modelo.Producto{}

	storage.DB().First(&miProducto, 1)
	fmt.Println(miProducto)

	/*

		OBTENINEDO TODOS LOS REGISTROS

		productos := make([]modelo.Producto, 0)
		storage.DB().Find(&productos)

		for _, producto := range productos {
			fmt.Printf("%d - %s \n", producto.ID, producto.Nombre)
		}


			CREANDO PRODUCTOS

			producto1 := modelo.Producto{
				Nombre: "Papas",
				Precio: 120,
			}

			//Se debe mandar un puntero al string pues la columna detalle acepta nulos
			detalle := "Tipo frances"
			producto2 := modelo.Producto{
				Nombre:  "Pan",
				Precio:  60,
				Detalle: &detalle,
			}

			producto3 := modelo.Producto{
				Nombre: "Espinaca",
				Precio: 150,
			}

			storage.DB().Create(&producto1)
			storage.DB().Create(&producto2)
			storage.DB().Create(&producto3)
	*/

	//	MIGRACIONES DE TABLAS
	//storage.DB().AutoMigrate(
	//	&modelo.Producto{},
	//	&modelo.EncabezadoFactura{},
	//	&modelo.ItemFactura{},
	//)

}
