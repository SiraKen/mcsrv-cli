package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func main() {

	cSuccess := color.New(color.FgWhite).Add(color.BgGreen)
	cError := color.New(color.FgWhite).Add(color.BgRed)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.minecraft.net/en-us/download/server/", nil)
	if err != nil {
		fmt.Println(cError.Sprintf("Error: %s", err))
		os.Exit(1)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.115 Safari/537.36")
	res, err := client.Do(req)

	if err != nil {
		cError.Println("Error: ", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		cError.Println("Error: ", res.StatusCode)
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		cError.Println("Error: ", err)
		os.Exit(1)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")

		if strings.Contains(href, ".jar") {
			fmt.Printf("Downloading: %s\n", href)

			if err := DownloadFile(href, "./server.jar"); err != nil {
				cError.Println("Error: ", err)
				os.Exit(1)
			} else {
				cSuccess.Println("Downloaded server.jar")

				fmt.Println("Starting server...")
				cmd := exec.Command("java", "-Xms1024M", "-Xmx1024M", "-jar", "./server.jar", "nogui")

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()

			}
		}
	})

}

func DownloadFile(url string, filepath string) error {

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	return err
}
