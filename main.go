package main

import (
	"fmt"

	modelo "github.com/S-Kiev/GORM_GO_Modulo/modelos"
	"github.com/S-Kiev/GORM_GO_Modulo/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	//BORRANDO REGISTROS
	//en gorm hay borrado SUAVE y PERMANENTE
	//en el SUAVE se le asigna una fecha de eliminacion ignorando esos registros como si no existiesen, pero no elimina el registro de la tabla
	//en el PERMANENTE se borran de forma permanente

	//BORRADO SUAVE

	miProducto := modelo.Producto{}
	miProducto.ID = 3

	storage.DB().Delete(&miProducto)

	//Ignorara al producto recien eliminado
	productos := make([]modelo.Producto, 0)
	storage.DB().Find(&productos)

	for _, producto := range productos {
		fmt.Printf("%d - %s \n", producto.ID, producto.Nombre)
	}

	//BORRADO PERMANENTE

	miProducto2 := modelo.Producto{}
	miProducto2.ID = 2

	storage.DB().Unscoped().Delete(&miProducto)

	/*
		//MODIFICANDO REGISTROS

		miProducto := modelo.Producto{}
		miProducto.ID = 3

		storage.DB().Model(&miProducto).Updates(
			modelo.Producto{Nombre: "Guayaba",
				Precio: 100,
			},
		)


			//OBTENIENDO UN REGISTRO

			miProducto := modelo.Producto{}

			storage.DB().First(&miProducto, 1)
			fmt.Println(miProducto)



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
