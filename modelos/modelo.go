package modelo

import (
	//"database/sql/driver"
	//"fmt"

	"gorm.io/gorm"
)

type Detalle string

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
	EncabezadoFacturaID uint
	ProductoID          uint
}

/*
// Value implementa la interfaz driver.Valuer
func (d Detalle) Value() (driver.Value, error) {
	return string(d), nil
}

// Scan implementa la interfaz sql.Scanner
func (d *Detalle) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	v, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Detalle: %v", value)
	}
	*d = Detalle(v)
	return nil
}
*/
