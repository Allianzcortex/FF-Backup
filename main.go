package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"net/http/cookiejar"
	"golang.org/x/net/publicsuffix"
	"log"
	"github.com/PuerkitoBio/goquery"
	//"os"
	//"encoding/json"
	"strings"

	"strconv"
	// "regexp"
	"os"
	"io"
	"bufio"
)

type Configuration struct {
	pagestart   string
	pageend     string
	loginemail  string
	loginpasswd string
	user        string
}

//func readConfig() (string, string, string, string, string) {
//	file, _ := os.Open("config.json")
//	decoder := json.NewDecoder(file)
//	config := Configuration{}
//	_ := decoder.Decode(&config)
//	return config.pagestart, config.pageend, config.loginemail, config.loginpasswd, config.loginpasswd
//}

type Jar struct {
	cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

func parse() {

	//配置文件
	//fmt.Println(os.Getwd())
	//file, err := os.Open("config.json")
	//if err!=nil {
	//	fmt.Println("error:",err)
	//}
	//decoder := json.NewDecoder(file)
	//var config  Configuration
	//err = decoder.Decode(&config)
	//if err!=nil {
	//	fmt.Println("error:",err)
	//}
	//
	//fmt.Println(config.user)
	fmt.Println("----")

	config := Configuration{}
	config.user = "http://fanfou.com/Allianzcortex"
	config.loginpasswd = "password"
	config.loginemail = "iamwanghz@gmail.com"
	config.pagestart = "1"
	config.pageend = "2"

	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}

	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{Jar: jar}
	// 判断验证码
	//resp, err := client.Get("http://fanfou.com/login")
	//res, _ := ioutil.ReadAll(resp.Body)
	//r := regexp.MustCompile(`<img src="\/\/(.*?)" width`)
	//// <img src="\/\/(.*?)" width
	//body := string(res)
	//captchAddress := r.FindStringSubmatch(body)[0]
	if (1 > 0) {
		_, err := client.PostForm("http://fanfou.com/login", url.Values{
			"loginname": {"%s", config.loginemail},
			"loginpass": {"%s", config.loginpasswd},
			"action":    {"login"},
			"token":     {"68b423d5"}})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// 下载图片并打开
		response, err := http.Get("http://")
		if err != nil {
			fmt.Println(err)
		}
		picture, _ := os.Create("captcha.png")
		_, err = io.Copy(picture, response.Body)
		if err != nil {
			fmt.Println(err)
		}
		if err != nil {
			fmt.Println(err)
		}
		picture.Close()

		// TODO 很多文件流打开后就没有关闭
		reader := bufio.NewReader(os.Stdin)
		captcha, _ := reader.ReadString('\n')
		fmt.Println(captcha)
		_, err = client.PostForm("http://fanfou.com/login", url.Values{
			"loginname": {"iamwanghz@gmail.com"},
			"loginpass": {"password"},
			"captcha":   {"\"" + captcha + "\""},
			"action":    {"login"},
			"token":     {"127396ab"}})
		if err != nil {
			log.Fatal(err)
		}
	}

	//start, _ := strconv.Atoi(config.pagestart)
	//end, _ := strconv.Atoi(config.pageend)
	start := 1
	end := 4
	count:=0
	for i := start; i <= end; i++ {

		resp, err := client.Get(config.user + "/p." + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		target := string(data)

		// fmt.Println(target)

		// 进行处理
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(target))
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("div#content .content").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			tweet := s.Text()
			fmt.Println(tweet)
			fmt.Println("---")
			count += 1
		})
	}
	fmt.Printf("总共读取 %d 条信息",count)
}

func main() {
	//  response, err := http.PostForm("http://127.0.0.1:8080/admin/user/login", map[string][]string{"username": []string{"admin"}, "pass": []string{"admin"}})

	//jar := new(Jar)
	//
	//client := &http.Client{nil, nil, jar, 99999999999992}
	//
	////req, _ := http.NewRequest("POST", "http://127.0.0.1:8080/admin/user/login", nil)
	//
	//resp, err := client.PostForm("http://fanfou.com/login", url.Values{"loginname": {"iamwanghz@gmail.com"}, "loginpass": {"402840evened"}, "action":{"login"}, "token":{"68b423d5"}})
	//
	//cookie    :=    http.Cookie{
	//	"__utmz=208515845.1515266115.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmc=208515845; __utma=208515845.1741402146.1515266115.1517038704.1517054440.50; __utmt=1; __utmv=208515845.visitor; u=Allianzcortex; uuid=f1675fd41a51ab2e29ba.1515266127.7; PHPSESSID=dfufps6rut8aqkl4vvl4s4q913; m=iamwanghz%40gmail.com; __utmb=208515845.48.10.1517054440"
	//}
	//if err != nil {
	//    panic(err.Error())
	//}
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//
	//resp, err = client.Get("http://fanfou.com/Allianzcortex")
	//
	//if err != nil {
	//    panic(err.Error())
	//}
	//
	//body, err = ioutil.ReadAll(resp.Body)
	//
	//fmt.Println(string(body))
	//
	//cookies := resp.Cookies()
	//
	//fmt.Println("xxxx")
	//for _, cookie := range cookies {
	//    fmt.Println("cookie:", cookie.Name)
	//}
	parse()

	// 所以 golag
}
