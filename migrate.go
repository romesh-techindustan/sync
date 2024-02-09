package sync

import (
	"database/sql"

	"github.com/romesh-techindustan/sync/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Database Connection and Auto migration
func DatabaseConn(dsn string)(*gorm.DB, error){
	sqlDB, err := sql.Open("mysql", dsn)
	if err !=nil{
		return nil, err
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
	Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
        return nil, err
    }
	
	gormDB.AutoMigrate(&models.User{})
    return gormDB, nil
}