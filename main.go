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
	//menampilkan data berdasarkan id 
	router.GET("/:id", func(c *gin.Context){
		var (
			list List
			result gin.H
		)
		id := c.Param("id")
		//query untuk mengambul data dalam database listofdaily
		row := db.QueryRow("select id, year, quarter from listofdaily where id = ?;", id)
		err = row.Scan(&list.Id, &list.Year, &list.Quarter)
		if err != nil{
			//jika datanya tidak ada maka akan dikirimkan null
			result = gin.H {
				"Hasil" : "Mohon maaf data tidak ditemukan",
				"Jumlah" : 0,
			}
		}else {
			result = gin.H{
				"Hasil" : "List Of Daily",
			}
		}
		c.JSON(http.StatusOK, result)
	})

	//menampilkan semua data pada database, List Of Daily
	router.GET("/", func(c *gin.Context){
		var (
			list List
			lists []List
		)
		rows, err := db.Query("select id, year, quarter from listofdaily;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next(){
			err = rows.Scan(&list.Id, &list.Year, &list.Quarter)
			lists = append(lists, list)
			if err != nil{
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"Hasil" : lists, 
			"Jumlah" : len(lists),
		})
	})
	
}