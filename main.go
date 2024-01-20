package main

import (
	"autoXreader/src"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/color"
	"github.com/sclevine/agouti"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	envFileErr := godotenv.Load(".env")
	if envFileErr != nil {
		log.Fatal("Error loading .env file")
	}
	userName := os.Getenv("USER_NAME")
	userPass := os.Getenv("USER_PASS")
	loadingTime := os.Getenv("LOADING_WAIT")
	strRandTime := os.Getenv("READ_RANDOM")
	randTime, err := strconv.Atoi(strRandTime)
	waitTime, err := strconv.Atoi(loadingTime)
	waitTimeIncludeRandom := int(math.Abs(float64(rand.Intn((waitTime+randTime)-(waitTime-randTime)) + waitTime - randTime)))
	if os.Getenv("DEBUG") == "1" {
		fmt.Println("DEBUG MODE")
		fmt.Println("USER_NAME:" + userName)
		fmt.Println("USER_PASS:" + userPass)
		fmt.Println("LOADING_TIME:" + loadingTime)
		fmt.Println("READ_RANDOM:" + strRandTime)
		fmt.Println("WAIT_TIME:" + strconv.Itoa(waitTime))
		fmt.Println("WAIT_TIME_INCLUDE_RANDOM:" + strconv.Itoa(waitTimeIncludeRandom))
	}
	if err != nil {
		log.Fatal(err)
	}
	color.Print(color.Cyan("Input UserName(default:" + userName + "):"))
	userNameInput := src.CmdLineInput()
	if userNameInput != "" {
		userName = userNameInput
	}
	color.Print(color.Cyan("Input Password(default:" + userPass + "):"))
	userPassInput := src.CmdLineInput()
	if userPassInput != "" {
		userPass = userPassInput
	}
	color.Print(color.Cyan("Input Target 'FULL' URL:"))
	targetURL := src.CmdLineInput()
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
	time.Sleep(1 * time.Second)
	page.Navigate(targetURL)
	time.Sleep(5 * time.Second)
	pageNumInd, err := page.FindByID("count_signal0").Text()
	if err != nil {
		log.Fatal(err)
	}
	pageNumArray := strings.Split(pageNumInd, "/")
	nowPage, err := strconv.Atoi(pageNumArray[0])
	if err != nil {
		log.Fatal(err)
	}
	maxPage, err := strconv.Atoi(pageNumArray[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nowPage)
	fmt.Println(maxPage)
	for i := nowPage; i < maxPage+1; {
		if os.Getenv("DEBUG") == "1" {
			fmt.Println("NOW_PAGE:" + strconv.Itoa(i))
			fmt.Println("MAX_PAGE:" + strconv.Itoa(maxPage))
		}
		pageNumInd, err = page.FindByID("count_signal0").Text()
		if err != nil {
			log.Fatal(err)
		}
		pageNumArray = strings.Split(pageNumInd, "/")
		i, err = strconv.Atoi(pageNumArray[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i)
		if i == maxPage {
			time.Sleep(time.Duration(waitTimeIncludeRandom) * time.Second)
			err := page.FindByButton("Close").Click()
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Duration(waitTimeIncludeRandom) * time.Second)
			page.Navigate("https://xreading.com/blocks/institution/dashboard.php?tm=dashboard")
		} else {
			time.Sleep(time.Duration(waitTimeIncludeRandom) * time.Second)
			err := page.FindByButton("Next").Click()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	time.Sleep(5 * time.Second)
	err = driver.Stop()
	if err != nil {
		log.Fatal(err)
	}
}
