package main

//melakukan import package dan library yang dibutuhkan
import (
	"bytes"
	"fmt"
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	//membuat koneksi ke dalam database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/taptalk-api")
	err = db.Ping()
	if err!= nil {
		panic("Gagal menghubungkan")
	}
	defer db.Close()

	router := gin.Default()
//struct ini berisikan field yang terdapat di database
	type List struct {
		Id	int `json: "id"`
		Year string `json: "year"`
		Quarter string `json: "quarter"`
	}
}