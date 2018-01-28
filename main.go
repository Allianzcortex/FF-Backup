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
	"strings"
	"strconv"
	"os"
	"io"
	"bufio"
	"github.com/BurntSushi/toml"
	"regexp"
)

type Configuration struct {
	Pagestart      string
	Pageend        string
	Loginemail     string
	Loginpasswd    string
	User           string
	BackupFilename string
}

func ReadConfig() Configuration {

	//_, err := os.Open(configfile)
	//if err != nil {
	//	log.Fatal("Config file is missing: ", configfile)
	//}
	var configfile = "config.toml"

	config := Configuration{}

	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

//func readConfig() (string, string, string, string, string) {
//	file, _ := os.Open("config.toml")
//	decoder := json.NewDecoder(file)
//	config.toml := Configuration{}
//	_ := decoder.Decode(&config.toml)
//	return config.toml.pagestart, config.toml.pageend, config.toml.loginemail, config.toml.loginpasswd, config.toml.loginpasswd
//}

// TODO 关于验证码这里也从简。其它尽快完成需要的部分发不过去就好了
// TODO 多 go 来加快执行速度

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
	//file, err := os.Open("config.toml")
	//if err!=nil {
	//	fmt.Println("error:",err)
	//}
	//decoder := json.NewDecoder(file)
	//var config.toml  Configuration
	//err = decoder.Decode(&config.toml)
	//if err!=nil {
	//	fmt.Println("error:",err)
	//}
	//
	//fmt.Println(config.toml.user)

	var configfile = "config.toml"

	var config Configuration

	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}

	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}

	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{Jar: jar}
	// 判断验证码
	resp, err := client.Get("http://fanfou.com/login")
	res, _ := ioutil.ReadAll(resp.Body)
	r := regexp.MustCompile(`<img src="\/\/(.*?)" width`)
	// <img src="\/\/(.*?)" width
	body := string(res)
	captchAddress := r.FindStringSubmatch(body)[1]
	if (1>0) {
		_, err := client.PostForm("https://fanfou.com/login", url.Values{
			"loginname": {"" + config.Loginemail},
			"loginpass": {"" + config.Loginpasswd},
			"action":    {"login"},
			"token":     {"42070970"}})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// 下载图片并打开
		response, err := http.Get("http://" + captchAddress)
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

		// TODO 读取用户输入验证码
		reader := bufio.NewReader(os.Stdin)
		captcha, _ := reader.ReadString('\n')
		fmt.Println(captcha)
		_, err = client.PostForm("http://fanfou.com/login", url.Values{
			"loginname": {"iamwanghz@gmail.com"},
			"loginpass": {"password"},
			"captcha":   {captcha},
			"action":    {"login"},
			"token":     {"127396ab"}})
		if err != nil {
			log.Fatal(err)
		}
	}

	//start, _ := strconv.Atoi(config.toml.pagestart)
	//end, _ := strconv.Atoi(config.toml.pageend)
	start, _ := strconv.Atoi(config.Pagestart)
	end, _ := strconv.Atoi(config.Pageend)

	count := 0

	// 创建文件
	f, err := os.Create(config.BackupFilename)
	w := bufio.NewWriter(f)

	var tweets []string
	var timestamps []string

	for i := start; i <= end; i++ {
		tweets = tweets[:0]
		timestamps = timestamps[:0]
		resp, err := client.Get(config.User + "/p." + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		target := string(data)
		fmt.Println(target)

		// 进行处理
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(target))
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("div#content .content").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			tweet := strconv.Itoa(count) + ". " + s.Text()
			tweets = append(tweets, tweet)

			count += 1
		})

		doc.Find("div#content .stamp").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			timestamp := s.Text()
			timestamps = append(timestamps, timestamp)

		})

		for index, _ := range tweets {

			w.WriteString(tweets[index] + " " + timestamps[index] + "\n---\n")
		}
		w.Flush()
		fmt.Printf("正在处理 %d 页面信息\n", i)
	}
	fmt.Printf("总共读取 %d 条信息", count)
}

func main() {

	parse()

}
