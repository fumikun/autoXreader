package selenium

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"time"
)

func Init(userName string, userPass string) *agouti.Page {
	driver := agouti.ChromeDriver()
	defer driver.Stop()
	if err := driver.Start(); err != nil {
		log.Fatalf("ERROR SELENIUM: %v", err)
	}
	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("ERROR PAGE: %v", err)
	}
	if err := page.Navigate("https://xreading.com/login/index.php"); err != nil {
		log.Fatalf("ERROR NAVIGATE: %v", err)
	}
	//page.ClearCookies()
	page.FindByID("username").Fill(userName)
	page.FindByID("password").Fill(userPass)
	btn := page.FindByButton("Log in")
	if err = btn.Click(); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	// 処理完了後、10秒間ブラウザを表示しておく
	time.Sleep(10 * time.Second)
	fmt.Println(page.FindByID("cust-BGpage-heading").Text())
	return page
}
