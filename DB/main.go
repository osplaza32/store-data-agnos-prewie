package main

import (
	"github.com/osplaza32/Module/DB/DBinit"
	"github.com/osplaza32/Module/DB/DbOPTION"
)

func main() {
	baseDato := DBinit.BaseDato{Direccion: "root:1234@/testame?charset=utf8&parseTime=True&loc=Local"}
	baseDato.InitDB(DbOPTION.MYSQL)
	baseDato2 := DBinit.BaseDato{Direccion: "host=localhost port=5432 user=kibernum dbname=gorm sslmode=disable password="}
	baseDato2.InitDB(DbOPTION.PG)
	baseDato3 := DBinit.BaseDato{Direccion: "mongodb://localhost:27017/corp"}
	baseDato3.InitDB(DbOPTION.MONGO)
}
