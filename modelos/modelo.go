package modelo

import "gorm.io/gorm"

// Producto es el modelo de la tabla productos
type Producto struct {
	gorm.Model
	Nombre       string  `gorm:"type:varchar(100); not null"`
	Detalle      *string `gorm:"type:varchar(100)"`
	Precio       int     `gorm:"not null"`
	ItemsFactura []ItemFactura
}

// EncabezadoFactura es el modelo de la tabla encabezado_factura
type EncabezadoFactura struct {
	gorm.Model
	Cliente      string `gorm:"type:varchar(100); not null"`
	ItemsFactura []ItemFactura
}

// ItemFactura es el modelo de la tabla item_factura
type ItemFactura struct {
	gorm.Model
	InvoiceHeaderID uint
	ProductID       uint
}
