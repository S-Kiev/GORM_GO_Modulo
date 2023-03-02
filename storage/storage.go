package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver del storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open("postgres://nombreUsuario:clave@localhost:7530/nombreBD?sslmode=disable"))
		if err != nil {
			log.Fatalf("No se pudo abrir la BD: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open("S-Kiev:sakura1997@tcp(localhost:3306)/bd-cursogo?parseTime=true"))
		if err != nil {
			log.Fatalf("No se pudo conectar con la Base de Datos: %v", err)
		}

		fmt.Println("conectado a MySQL")
	})
}

// Pool retorna una unica instacia de db
func DB() *gorm.DB {
	return db
}
