// main.go
package main

import (
    "github.com/sclevine/agouti"
    "github.com/PuerkitoBio/goquery"
//    "github.com/microcosm-cc/bluemonday"
//    "golang.org/x/text/width"
//    "golang.org/x/text/collate"
//    "golang.org/x/text/language"
//    "golang.org/x/text/unicode/norm"
//    "unicode/utf8"
    "strings"
    "log"
    "io/ioutil"
    "fmt"
//    "os"
)

func main() {
    driver := agouti.PhantomJS()
    if err := driver.Start(); err != nil {
        log.Fatalf("Failed to start driver:%v", err)
    }
    defer driver.Stop()

    page, err := driver.NewPage(agouti.Browser("phantomjs"))
    if err != nil {
        log.Fatalf("Failed to open page:%v", err)
    }

    if err := page.Navigate("https://www.youtube.com/"); err != nil {
        log.Fatalf("Failed to navigate:%v", err)
    }
    page.Screenshot("./phantomjs_youtube.jpg")
    html, _ := page.HTML()
    fmt.Println(html)

//    fmt.Println(page.Find("span").Text())


    if false {
        fileInfos, _ := ioutil.ReadFile("index.html")
        stringReader := strings.NewReader(string(fileInfos))
        doc, err := goquery.NewDocumentFromReader(stringReader)
        if err != nil {
            fmt.Print("url scarapping failed")
        }
        doc.Find("a").Each(func(_ int, s *goquery.Selection) {
            url, _ := s.Attr("href")
            fmt.Println(url)
        })
    }
}
