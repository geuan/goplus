package AppInit
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {


	var err error
	db, err = gorm.Open("mysql",
		"root:xc456789110@tcp(192.168.19.138:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	err = db.DB().Ping()
	if err != nil {
		log.Println(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)



	// db.LogMode(true)
}
func  GetDB() *gorm.DB {
	return db
}
