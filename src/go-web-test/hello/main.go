package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"html/template"
	"strconv"
	"regexp"
	"time"
	"crypto/md5"
	"io"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprint(w, "hello astaxie")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mehtod:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("D://Projects/go/src/ge-web/login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form.Get("username"))
		fmt.Println("password:", r.Form["password"])

		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
		}

		if getint > 100 {
		}

		// Go实现的正则是RE2，所有的字符都是UTF-8编码的
		// 数字
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		}

		// 匹配汉字
		if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
		}

		// 匹配字母
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
		}

		// 匹配mail
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("no")
		}else{
			fmt.Println("yes")
		}

		// 匹配手机号
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
		}

		t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		fmt.Printf("Go launched at %s\n", t.Local())

		//验证15位身份证，15位的是全部数字
		if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
		}

		//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		}

		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
