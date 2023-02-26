package main

import (
	"context"
	"os"
	"log"
	"flag"

	"github.com/chromedp/chromedp"
	"github.com/pkg/browser"
)

func main() {
    url := flag.String("url", "https://www.golang.org", "URL of the webpage to capture")
	file := flag.String("file", "screenshot.png", "File name of the capture")
	open := flag.Bool("open", false, "Open the image when captured")

	flag.Parse()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(*url),
		chromedp.CaptureScreenshot(&buf),
	}); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(*file, buf, 0o644); err != nil {
		log.Fatal(err)
	}

	if *open {
		browser.OpenFile(*file)
	}
}