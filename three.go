package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Pornstar struct {
	gorm.Model
	//ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	//	City      string `json:"city"`
	Age         string `json:"age"`
	DustDRating int    `json: "dustdrating"`
	Mask        bool   `json: "mask"`
	Somatype    string `json: "somatype"`
	SomaInfo    bool   `json: "somainfo"`
	Shape       string `json: "shape"`
	ShapeInfo   bool   `json: "shapeinfo"`
	PName       string `json: "pname"`
}
type Camgirl struct {
	gorm.Model
	//ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	//	City      string `json:"city"`
	Age         string `json:"age"`
	DustDRating int    `json: "dustdrating"`
	Mask        bool   `json: "mask"`
	Somatype    string `json: "somatype"`
	SomaInfo    bool   `json: "somainfo"`
	Shape       string `json: "shape"`
	ShapeInfo   bool   `json: "shapeinfo"`
	PName       string `json: "pname"`
}

func main() {

	// NOTE: See we're using = to assign the global var
	//         	instead of := which would assign it only in this function
	//db, err = gorm.Open("sqlite3", "./gorm.db")
	db, _ = gorm.Open("mysql", "pron:Echo--@tcp(127.0.0.1:33061)/porn?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Pornstar{})

	r := gin.Default()
	r.GET("/people/", GetPornstars)
	r.GET("/people/:id", GetPornstar)
	r.POST("/people", CreatePornstar)
	r.PUT("/people/:id", UpdatePornstar)
	r.DELETE("/people/:id", DeletePornstar)

	r.Run(":8080")
}

func DeletePornstar(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Pornstar
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdatePornstar(c *gin.Context) {

	var person Pornstar
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)

	db.Save(&person)
	c.JSON(200, person)

}

func CreatePornstar(c *gin.Context) {

	var person Pornstar
	c.BindJSON(&person)

	db.Create(&person)
	c.JSON(200, person)
}

func GetPornstar(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Pornstar
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}
func GetPornstars(c *gin.Context) {
	var people []Pornstar
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}

}
