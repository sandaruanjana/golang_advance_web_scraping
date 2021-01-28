package main

import (
	"fmt"
	"github.com/mxschmitt/playwright-go"
	"log"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func main() {

	var s []string

	var title string
	var date string
	var price string
	var address string
	var desc string
	var tel string

	fmt.Print("Enter Category Type: ")
	var category string
	fmt.Scanln(&category)

	fmt.Print("Enter District: ")
	var district string
	fmt.Scanln(&district)
	fmt.Println("------------------------------")

	mainDomain := "https://ikman.lk"
	url := "https://ikman.lk/en/ads/" + district + "/" + category

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto(url); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	entries, err := page.QuerySelectorAll(".list--3NxGO")

	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}

	for _, entry := range entries {

		//titleElement, err := entry.QuerySelectorAll(".normal--2QYVk gtm-normal-ad")

		titleElement, err := entry.QuerySelectorAll("li > a")
		if err != nil {
			log.Fatalf("could not get title element: %v", err)
		}

		for _, titleElementarray := range titleElement {
			titlenew, _ := titleElementarray.GetAttribute("href")

			fullPath := (mainDomain + titlenew + "\n")
			s = append(s, fullPath)

		}

		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
	}

	if category == "property" {

		for _, prop := range s {

			Block{
				Try: func() {
					page.Goto(prop)

					getTitle, _ := page.QuerySelector(".title--3s1R8")
					getPrice, _ := page.QuerySelector(".amount--3NTpl")
					getDate, _ := page.QuerySelector(".sub-title--37mkY")
					getAddress, _ := page.QuerySelector(".full-width--XovDn")
					getDesc, _ := page.QuerySelector(".description--1nRbz")

					page.Click("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[2]/div/div/div[1]/div[2]/div[1]/button")
					getPhone, _ := page.QuerySelector(".phone-numbers--2COKR")

					title, _ = getTitle.TextContent()
					price, _ = getPrice.TextContent()
					date, _ = getDate.TextContent()
					address, _ = getAddress.TextContent()
					desc, _ = getDesc.TextContent()
					tel, _ = getPhone.TextContent()

					fmt.Println(title)
					fmt.Println(date)
					fmt.Println(price)
					fmt.Println(address)
					fmt.Println(tel)
					fmt.Println(desc)

					fmt.Println("----------------------------------------")

				},
				Catch: func(e Exception) {
					fmt.Printf("Caught %v\n", e)
				},
				Finally: func() {

				},
			}.Do()

		}

	}

	if category == "electronics" {

		for _, prop := range s {

			Block{
				Try: func() {
					page.Goto(prop)

					getTitle, _ := page.QuerySelector(".title--3s1R8")
					getPrice, _ := page.QuerySelector(".amount--3NTpl")
					getDate, _ := page.QuerySelector(".sub-title--37mkY")
					getAddress, _ := page.QuerySelector(".full-width--XovDn")
					getDesc, _ := page.QuerySelector(".description--1nRbz")

					page.Click("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[2]/div/div/div[1]/div[2]/div[1]/button")
					getPhone, _ := page.QuerySelector(".phone-numbers--2COKR")

					title, _ = getTitle.TextContent()
					price, _ = getPrice.TextContent()
					date, _ = getDate.TextContent()
					address, _ = getAddress.TextContent()
					desc, _ = getDesc.TextContent()
					tel, _ = getPhone.TextContent()

					fmt.Println(title)
					fmt.Println(date)
					fmt.Println(price)
					fmt.Println(address)
					fmt.Println(tel)
					fmt.Println(desc)

					fmt.Println("----------------------------------------")

				},
				Catch: func(e Exception) {
					//fmt.Printf("Caught %v\n", e)
				},
				Finally: func() {

				},
			}.Do()

		}

	}

	if category == "vehicles" {

		for _, prop := range s {

			Block{
				Try: func() {
					page.Goto(prop)

					getTitle, _ := page.QuerySelector(".title--3s1R8")
					getPrice, _ := page.QuerySelector(".amount--3NTpl")
					getDate, _ := page.QuerySelector(".sub-title--37mkY")
					getAddress, _ := page.QuerySelector(".full-width--XovDn")
					getDesc, _ := page.QuerySelector(".description--1nRbz")

					page.Click("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[2]/div/div/div[1]/div[2]/div[1]/button")
					getPhone, _ := page.QuerySelector(".phone-numbers--2COKR")

					title, _ = getTitle.TextContent()
					price, _ = getPrice.TextContent()
					date, _ = getDate.TextContent()
					address, _ = getAddress.TextContent()
					desc, _ = getDesc.TextContent()
					tel, _ = getPhone.TextContent()

					fmt.Println(title)
					fmt.Println(date)
					fmt.Println(price)
					fmt.Println(address)
					fmt.Println(tel)
					fmt.Println(desc)

					fmt.Println("----------------------------------------")

				},
				Catch: func(e Exception) {
					//fmt.Printf("Caught %v\n", e)
				},
				Finally: func() {

				},
			}.Do()

		}

	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
