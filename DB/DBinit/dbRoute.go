package DBinit

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/osplaza32/Module/DB/DbOPTION"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
type BaseDato struct {
	Direccion string
	driver * gorm.DB
	mongo *mongo.Client
	typebd DbOPTION.Motor
}
func(bd *BaseDato)InitDB(db DbOPTION.Motor){
	bd.typebd= db
	switch os := db; os {
	case DbOPTION.MYSQL:
		db, err := gorm.Open("mysql", bd.Direccion)
		if err != nil {
			fmt.Println(err)
		}
		bd.driver = db
	case DbOPTION.PG:
		db, err := gorm.Open("postgres", bd.Direccion)
		if err != nil {
			fmt.Println(err)
		}
		bd.driver = db
	case DbOPTION.SQLITE:
		db, err := gorm.Open("sqlite", bd.Direccion)
		if err != nil {
			fmt.Println(err)
		}
		bd.driver = db
	case DbOPTION.MSSQL:
		db, err := gorm.Open("mssql", bd.Direccion)
		if err != nil {
			fmt.Println(err)
		}
		bd.driver = db
	case DbOPTION.MONGO:
		bd.initmongo()
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	if bd.typebd != DbOPTION.MONGO {
		fmt.Println(bd.driver.DB().Ping())

	}else{
		fmt.Println(bd.mongo.Ping(context.Background(),readpref.Nearest()))
	}
}
