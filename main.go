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
		Id		int 	`json:"id"`
		Year 	string 	`json:"year"`
		Quarter string 	`json:"quarter"`
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

	//Menambahkan List Of Daily
	//Disini method yang saya gunakan adalah POST
	router.POST("/", func(c *gin.Context){
		var buffer bytes.Buffer
		id := c.PostForm("id")
		year := c.PostForm("year")
		quarter := c.PostForm("quarter")
		//kita akan melakukan query untuk memasukan data kedalam database
		stmt, err := db.Prepare("insert into listofdaily (id, year, quarter) values(?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id, year, quarter)

		if err != nil {
			fmt.Print(err.Error())
		}
		// disini saya membuat cara cepat untuk melakukan append string
		buffer.WriteString(year)
		buffer.WriteString( " ")
		buffer.WriteString(quarter)
		defer stmt.Close()
		data := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"Message": fmt.Sprintf("List of Daiy berhasil ditambahkan %s", data),
		})
	})
// disini kita akan mengubah/mengupdate data
router.PUT("/", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.PostForm("id")
		year := c.PostForm("year")
		quarter := c.PostForm("quarter")
		//saya melakukan query untuk mengupdate data
		stmt, err := db.Prepare("update listofdaily set year= ?, quarter= ? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(year, quarter, id)
		if err != nil {
			fmt.Print(err.Error())
		}

		buffer.WriteString(year)
		buffer.WriteString(" ")
		buffer.WriteString(quarter)
		defer stmt.Close()
		data := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"Message": fmt.Sprintf("Data berhasil diubah %s", data),
		})
	})

}