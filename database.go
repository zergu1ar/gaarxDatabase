package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	gaarx "github.com/zergu1ar/Gaarx"
)

func GetConnString(user, pass, host, port, dbName string) string {
	return user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4,utf8&parseTime=true&sql_mode=ansi"
}

// Include Database
func WithDatabase(conn string, dbType string, entities ...interface{}) gaarx.Option {
	return func(app *gaarx.App) error {
		db, err := gorm.Open(
			dbType,
			conn,
		)
		if err != nil {
			app.GetLog().Fatal(err)
			panic(err)
		}
		db.SetLogger(app.GetLog())
		db.Set("gorm:table_options", "CHARSET=utf8")
		for _, e := range entities {
			db.AutoMigrate(e)
		}
		app.SetDatabase(db)
		return nil
	}
}
