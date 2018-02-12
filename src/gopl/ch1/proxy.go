package main

import (
	"net/http"
	"net/url"
	"fmt"
)

func main() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1080")
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}
	resp, err := client.Get("http://www.google.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Body)

}

//func main() {
//	reg,_ := regexp.Compile("47.88.223.49");
//	proxy := func(_ *http.Request) (*url.URL, error) {
//		return url.Parse("http://127.0.0.1:1080")//根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
//	}
//	transport := &http.Transport{Proxy: proxy}
//	client := &http.Client{Transport: transport}
//	resp, err := client.Get("http://google.com/") //请求并获取到对象,使用代理
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	res, err := http.Get("http://google.com/") //请求并获取到对象
//	dataproxy, err := ioutil.ReadAll(resp.Body) //取出主体的内容
//	if err != nil {
//		log.Fatal(err)
//	}
//	data, err := ioutil.ReadAll(res.Body) //取出主体的内容
//	if err != nil {
//		log.Fatal(err)
//	}
//	//fmt.Printf("%s",data) //打印
//	sproxy:=reg.FindString(string(dataproxy));
//	s:=reg.FindString(string(data));
//	res.Body.Close()
//	resp.Body.Close()
//	fmt.Printf("不使用代理:%s",s) //打印
//	fmt.Printf("使用代理:%s",sproxy) //打印
//}
