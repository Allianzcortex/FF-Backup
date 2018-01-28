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
	//"go/doc"
	"strings"
)

type Jar struct {
    cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
    jar.cookies = cookies
}
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
    return jar.cookies
}

func parse(){
	options:=cookiejar.Options{
		PublicSuffixList:publicsuffix.List,
	}

	jar, err := cookiejar.New(&options)
    if err != nil {
        log.Fatal(err)
    }
    client := http.Client{Jar: jar}
    resp, err := client.PostForm("http://fanfou.com/login", url.Values{
        "loginname": {"iamwanghz@gmail.com"},
        "loginpass" : {"password"},
        "action":{"login"},
        "token":{"68b423d5"}})
    if err != nil {
        log.Fatal(err)
    }

    resp,err=client.Get("http://fanfou.com/Allianzcortex")
    if err!=nil {
    	log.Fatal(err)
	}
	data,err:=ioutil.ReadAll(resp.Body)
	if err!=nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	target :=string(data)
	fmt.Println(target)


	// 进行处理
	doc,err:=goquery.NewDocumentFromReader(strings.NewReader(target))
	if err!=nil {
		log.Fatal(err)
	}

	doc.Find("div#content .content").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    tweet := s.Text()
    fmt.Println(tweet)
    fmt.Println("---")
  })
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
}
