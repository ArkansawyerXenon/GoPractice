// main.go
package main

import (
    "github.com/sclevine/agouti"
    "log"
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
}
