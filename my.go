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

type Actor struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function
	//db, err = gorm.Open("sqlite3", "./gorm.db")
	db, _ = gorm.Open("mysql", "pron:Echo--@tcp(127.0.0.1:33061)/porn?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&Actor{})
	r := gin.Default()
	r.GET("/Actors/", GetActors)
	r.GET("/Actors/:id", GetActor)
	r.POST("/Actors", CreateActor)
	r.PUT("/Actors/:id", UpdateActor)
	r.DELETE("/Actors/:id", DeleteActor)
	r.Run(":8080")
}
func DeleteActor(c *gin.Context) {
	id := c.Params.ByName("id")
	var actor Actor
	d := db.Where("id = ?", id).Delete(&actor)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
func UpdateActor(c *gin.Context) {
	var actor Actor
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&actor).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&actor)
	db.Save(&actor)
	c.JSON(200, actor)
}
func CreateActor(c *gin.Context) {
	var actor Actor
	c.BindJSON(&actor)
	db.Create(&actor)
	c.JSON(200, actor)
}
func GetActor(c *gin.Context) {
	id := c.Params.ByName("id")
	var actor Actor
	if err := db.Where("id = ?", id).First(&actor).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, actor)
	}
}
func GetActors(c *gin.Context) {
	var Actors []Actor
	if err := db.Find(&Actors).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, Actors)
	}
}
