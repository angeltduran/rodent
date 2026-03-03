package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func main() {
	targetURL := flag.String("url", "", "Website URL to convert (required)")
	outFile := flag.String("out", "output.pdf", "Output PDF filename")
	timeout := flag.Duration("timeout", 30*time.Second, "Timeout for page load and rendering")
	flag.Parse()

	if *targetURL == "" {
		exitWithErr("missing required -url flag")
	}

	if _, err := url.ParseRequestURI(*targetURL); err != nil {
		exitWithErr(fmt.Sprintf("invalid URL: %v", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	l := launcher.New().Headless(true)
	browserURL, err := l.Launch()
	if err != nil {
		exitWithErr(fmt.Sprintf("failed to launch browser: %v", err))
	}

	browser := rod.New().ControlURL(browserURL).MustConnect()
	defer browser.MustClose()

	page, err := browser.Page(proto.TargetCreateTarget{URL: "about:blank"})
	if err != nil {
		exitWithErr(fmt.Sprintf("failed to create page: %v", err))
	}
	defer page.MustClose()

	waitCtx, waitCancel := context.WithTimeout(ctx, *timeout)
	defer waitCancel()

	if err := page.Context(waitCtx).Navigate(*targetURL); err != nil {
		exitWithErr(fmt.Sprintf("failed to navigate to URL: %v", err))
	}

	if err := page.Context(waitCtx).WaitLoad(); err != nil {
		exitWithErr(fmt.Sprintf("failed waiting for page load: %v", err))
	}

	pdfStream, err := page.Context(waitCtx).PDF(&proto.PagePrintToPDF{
		PrintBackground: true,
	})
	if err != nil {
		exitWithErr(fmt.Sprintf("failed to create PDF: %v", err))
	}
	defer pdfStream.Close()

	pdfData, err := io.ReadAll(pdfStream)
	if err != nil {
		exitWithErr(fmt.Sprintf("failed to read PDF stream: %v", err))
	}

	if err := os.WriteFile(*outFile, pdfData, 0o644); err != nil {
		exitWithErr(fmt.Sprintf("failed writing file %q: %v", *outFile, err))
	}

	fmt.Printf("PDF saved to %s\n", *outFile)
}

func exitWithErr(msg string) {
	fmt.Fprintln(os.Stderr, "error:", msg)
	os.Exit(1)
}
