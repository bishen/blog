package main

import (
	"log"

	"./controllers"
	"./fun"
	"./models"
	"github.com/go-martini/martini"
	"github.com/go-ozzo/ozzo-config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/martini-contrib/binding"
	"github.com/zdebeer99/gojade"
)

func checkErr(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}

func main() {
	config := config.New()
	config.Load("config.json")
	cfbase := config.GetString("database", "sqlite3")
	cfconn := config.GetString("connection", "./app.db")
	db, err := gorm.Open(cfbase, cfconn)
	checkErr(err, "Database connection failed")
	db.AutoMigrate(&models.Blog{})
	db.AutoMigrate(&models.Classify{})
	defer db.Close()

	admin := &fun.Admins{Account: config.GetString("web_account", "admin"), Password: config.GetString("web_password", "admin")}

	jade := gojade.New()
	jade.ViewPath = "./jade"
	jade.RegisterFunction("date", fun.Date)
	jade.RegisterFunction("dateformat", fun.DateFormat)
	jade.RegisterFunction("compare", fun.Compare)
	jade.RegisterFunction("compare_not", fun.CompareNot)
	jade.RegisterFunction("substr", fun.Substr)
	jade.RegisterFunction("subtext", fun.Subtext)
	jade.RegisterFunction("img", fun.Img)
	jade.RegisterFunction("html2str", fun.HTML2str)
	jade.RegisterFunction("str2html", fun.Str2html)
	jade.RegisterFunction("htmlquote", fun.Htmlquote)
	jade.RegisterFunction("htmlunquote", fun.Htmlunquote)
	jade.RegisterFunction("map_get", fun.MapGet)

	m := martini.Classic()
	m.Map(db)
	m.Map(jade)
	m.Map(admin)
	m.Use(martini.Static("static"))
	m.Get("/", controllers.Home)
	m.Get("/:tag.html", controllers.Tag)
	m.Get("/view/:id.html", controllers.View)
	m.Get("/_login", controllers.Login)
	m.Get("/_logout", controllers.Logout)
	m.Get("/_add", controllers.Check, controllers.Create)
	m.Get("/_class", controllers.Check, controllers.Classify)
	m.Get("/_edit/:id", controllers.Check, controllers.Edit)
	m.Get("/_delete/:id", controllers.Check, controllers.Delete)

	m.Post("/_login", binding.Bind(models.Admin{}), controllers.LoginPost)
	m.Post("/_add", binding.Bind(models.Blog{}), controllers.CreatePost)
	m.Post("/_class", binding.Bind(models.Classify{}), controllers.ClassifyPost)
	m.Post("/_edit/:id", binding.Bind(models.Blog{}), controllers.UpdatePost)

	m.NotFound(controllers.NotFound)
	m.Run()
}
