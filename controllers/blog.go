package controllers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"../fun"
	"../models"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/zdebeer99/gojade"
)

func IsLogin(us string, req *http.Request) bool {
	cookie, err := req.Cookie("_ac")
	if err != nil {
		return false
	}
	if cookie.Value != us {
		return false
	}
	return true
}

func Home(ad *fun.Admins, req *http.Request, res http.ResponseWriter, jade *gojade.Engine, db *gorm.DB) {
	blogs := []models.Blog{}
	db.Find(&blogs)
	datas := &struct {
		Title string
		Is    bool
		Data  []models.Blog
	}{
		Title: "首页",
		Is:    IsLogin(ad.Account, req),
		Data:  blogs,
	}
	jade.RenderFileW(res, "home", datas)
}

func Login(res http.ResponseWriter, jade *gojade.Engine) {
	jade.RenderFileW(res, "login", nil)
}

func Logout(res http.ResponseWriter, req *http.Request) string {
	cookie := http.Cookie{Name: "_ac", Path: "/", MaxAge: -1}
	http.SetCookie(res, &cookie)
	return "<script>alert('成功登出！');window.location.assign('/')</script>"
}

func Tag(ad *fun.Admins, req *http.Request, params martini.Params, res http.ResponseWriter, jade *gojade.Engine, db *gorm.DB) {
	cs := models.Classify{}
	tag := params["tag"]
	db.Where("url=?", tag).First(&cs)
	if cs.Title != "" {
		blogs := []models.Blog{}
		db.Where("cid=?", cs.ID).Find(&blogs)
		data := &struct {
			Cls  models.Classify
			Is   bool
			Data []models.Blog
		}{
			Cls:  cs,
			Is:   IsLogin(ad.Account, req),
			Data: blogs,
		}
		jade.RenderFileW(res, "tag", data)
	} else {
		jade.RenderFileW(res, "404", nil)
	}
}

func View(ad *fun.Admins, params martini.Params, res http.ResponseWriter, req *http.Request, jade *gojade.Engine, db *gorm.DB) {
	blog := &models.Blog{}
	db.First(blog, params["id"])
	if blog.Title != "" {
		data := &struct {
			B  models.Blog
			Is bool
		}{
			B:  *blog,
			Is: IsLogin(ad.Account, req),
		}
		jade.RenderFileW(res, "view", data)
	} else {
		jade.RenderFileW(res, "404", nil)
	}
}

func Edit(params martini.Params, res http.ResponseWriter, jade *gojade.Engine, db *gorm.DB) {
	blog := models.Blog{}
	db.First(&blog, params["id"])
	jade.RenderFileW(res, "edit", blog)
}

func Delete(params martini.Params, db *gorm.DB) string {
	db.Delete(models.Blog{}, params["id"])
	return "<script>alert('删除完成');window.history.back()</script>"
}

func Create(rw http.ResponseWriter, jade *gojade.Engine, db *gorm.DB) {
	cls := []models.Classify{}
	db.Find(&cls)
	datas := &struct {
		Cls []models.Classify
	}{
		Cls: cls,
	}
	jade.RenderFileW(rw, "create", datas)
}

func Classify(rw http.ResponseWriter, jade *gojade.Engine, db *gorm.DB) {
	cls := []models.Classify{}
	db.Find(&cls)
	data := &struct {
		Data []models.Classify
	}{
		Data: cls,
	}
	jade.RenderFileW(rw, "class", data)
}

func NotFound(rw http.ResponseWriter, jade *gojade.Engine) {
	jade.RenderFileW(rw, "404", nil)
}

func LoginPost(ad *fun.Admins, res http.ResponseWriter, req *http.Request, admin models.Admin) string {
	if admin.User == ad.Account && admin.Pass == ad.Password {
		fmt.Println(admin.User + "   -  " + ad.Account)
		cookie := http.Cookie{Name: "_ac", Value: admin.User, Path: "/", MaxAge: 86400}
		http.SetCookie(res, &cookie)
		return "<script>alert('登录成功！');window.location.assign('/')</script>"
	} else {
		return "<script>alert('登录失败！');window.history.back()</script>"
	}
}

func CreatePost(params martini.Params, res http.ResponseWriter, blog models.Blog, db *gorm.DB) string {
	db.Save(&blog)
	return "创建成功"
}

func ClassifyPost(params martini.Params, res http.ResponseWriter, cls models.Classify, db *gorm.DB) string {
	db.Save(&cls)
	fi, err := os.Open("./jade/layout-tmpl.jade")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err2 := ioutil.ReadAll(fi)

	f, err2 := os.OpenFile("./jade/layout.jade", os.O_RDWR, 0755)
	if err2 != nil {
		panic(err)
	}

	css := []models.Classify{}
	db.Find(&css)
	var buf bytes.Buffer
	for _, v := range css {
		buf.WriteString("        a(href=\"/" + v.Url + ".html\") " + v.Title + "\n")
	}
	w := strings.Replace(string(fd), "####", buf.String(), -1)
	n, err1 := io.WriteString(f, w) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("写入 %d 个字节n", n)
	return "1"
}

func UpdatePost(params martini.Params, res http.ResponseWriter, req *http.Request, updated models.Blog, db *gorm.DB) {
	blog := models.Blog{}
	db.First(&blog, params["id"])
	db.Model(&blog).Updates(updated)
	http.Redirect(res, req, "../view/"+params["id"]+".html", 302)
}

func Check(ad *fun.Admins, res http.ResponseWriter, req *http.Request) {
	if !IsLogin(ad.Account, req) {
		res.Write([]byte("<script>alert('请先登录！');window.history.back()</script>"))
	}
}
