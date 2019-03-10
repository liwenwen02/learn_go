package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"
)
type Liuyan struct {
	Title string
	Content string
	CreateTime string
}

func main() {

	// 连接数据库
	db, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 准备模板
	tpl, err := template.New("liuyanbook").Parse(html)
	if err != nil {
		panic(err)
	}

	// 显示留言页面 /
	requestList := func(w http.ResponseWriter, req *http.Request) {
		// 查询数据
		num, err := db.Do("llen", "mess")
		if err != nil {
			log.Fatal(err)
		}

		total, ok := num.(int64)
		if !ok {
			fmt.Println("convert failed")
		}
		lys := []Liuyan{}

		for i := 0; i < int(total); i++ {
			ly := Liuyan{}
			val, err := redis.Bytes(db.Do("lindex", "mess", i))
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(val, &ly)
			if err != nil {
				log.Fatal(err)
			}
			lys = append(lys,ly)
		}

		// 显示数据
		err = tpl.ExecuteTemplate(w, "list", lys)
		if err != nil {
			log.Fatal(err)
		}
	}



	// 留言页面 /liuyan
	requestLiuyan := func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil{
			log.Fatal(err)
		}

		if "POST" == req.Method {
			if len(req.Form["title"]) < 1 {
				io.WriteString(w, "参数错误!\n")
				return
			}
			if len(req.Form["content"]) < 1 {
				io.WriteString(w, "参数错误!\n")
				return
			}

			title := template.HTMLEscapeString(req.Form.Get("title"))
			content := template.HTMLEscapeString(req.Form.Get("content"))
            liuyan_var := Liuyan{title,content,time.Now().Format("2006-01-02 15:04:05")}
			// sql语句
			b, err :=json.Marshal(&liuyan_var)
			if err != nil {
				log.Fatal(err)
			}
			_, err = db.Do("LPUSH", "mess", string(b))
			if err != nil {
				log.Fatal(err)
			}

			// 跳转
			w.Header().Add("Location", "/")
			w.WriteHeader(302)

			// 提示信息
			io.WriteString(w, "提交成功!\n")

			return
		}

		err = tpl.ExecuteTemplate(w, "liuyan", nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", requestList)
	http.HandleFunc("/liuyan", requestLiuyan)
	err = http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 网页模板
var html string = `{{define "list"}}{{/* 留言列表页面 */}}<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
    <p><a href="/liuyan">给我留言</a></p>
    <table>
{{range .}}
    <tr>
        <td>{{.Title}}</td><td>{{.Content}}</td><td>{{.CreateTime}}</td>
    </tr>
{{end}}
    </table>
</body>
</html>{{end}}
{{define "liuyan"}}{{/* 发布留言页面 */}}<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
    <form method="post">
        标题:<input type="text" name="title" /><br>
        内容:<input type="text" name="content" /><br>
        <input type="submit" value="提交" />
    </form>
</body>
</html>{{end}}`