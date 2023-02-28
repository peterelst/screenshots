package screenshot

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/pkg/browser"
)

func main() {
	url := flag.String("url", "https://www.golang.org", "URL of the webpage to capture")
	file := flag.String("file", "screenshot.png", "File name of the capture")
	open := flag.Bool("open", false, "Open the image when captured")
	flag.Parse()

	var buf []byte
	buf, err := screenshot(url)
	if err != nil {
		log.Fatal(err)
	}

	writeFile(file, buf)
	if err != nil {
		log.Fatal(err)
	}

	if *open {
		browser.OpenFile(*file)
	}
}

func screenshot(url *string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(*url),
		chromedp.CaptureScreenshot(&buf),
	})
	return buf, err
}

func writeFile(file *string, buf []byte) error {
	return os.WriteFile(*file, buf, 0o644)
}
