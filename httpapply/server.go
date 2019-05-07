package httpapply

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func ExampleGet() {
	response, err := http.Get("http://httpbin.org/get")

	if err == nil {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}

	defer response.Body.Close()

	fmt.Println("request bad!!!")
}

func ExamplePost() {
	url := "http://httpbin.org/post"
	//contentType 用于指定参数放在data中还是form中
	resp, err := http.Post(url,
		"application/x-www-urlencoded",
		strings.NewReader("name=ygh&sex=body"))

	if err != nil {
		panic(err.Error())
	}

	//使用完需要关闭一下
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(content))
}

func ExamplePostForm() {
	href := "http://httpbin.org/post"

	resp, err := http.PostForm(href, url.Values{
		"name": {"smitch"},
		"age":  {"30"},
	})

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(content))
}
/**
	复杂的请求，使用client 发送 以及 设置header参数
 */
func ExamplePostHeader() {
	//第一步，构造client 和 request请求
	client := &http.Client{}

	request , _ := http.NewRequest("post",
		"http://httpbin.org/post",
		strings.NewReader("name=Smith"))

	request.Header.Set("Content-Type","application/x-www-urlencoded")
	request.Header.Set("Cookies","this is cookie")

	//第二步，模拟请求
	response , _ := client.Do(request)

	//第三步，获取响应值

	defer response.Body.Close()

	content , _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
