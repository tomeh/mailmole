package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/tomeh/mailmole/web"
)

func init() {
	flags()
	log.SetOutput(os.Stdout)
}

func flags() {
	flag.IntVar(&port, "port", 8084, "The port at which to serve http.")
	flag.StringVar(&host, "host", "127.0.0.1", "The host at which to serve http.")
	flag.BoolVar(&autoLaunchBrowser, "launchBrowser", true, "Open a browser session (default: true)")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	flag.Parse()

	server := web.NewServer(host, port)

	if autoLaunchBrowser {
		go launchBrowser(server.GetBaseUrl())
	}

	server.Start()
}

func browserCmd() (string, bool) {
	browser := map[string]string{
		"darwin":  "open",
		"linux":   "xdg-open",
		"windows": "start",
	}
	cmd, ok := browser[runtime.GOOS]
	return cmd, ok
}

func launchBrowser(addr string) {
	browser, ok := browserCmd()
	if !ok {
		log.Printf("Cannot launch browser on %s systems", runtime.GOOS)
		return
	}

	url := fmt.Sprintf("http://%s", addr)
	log.Printf("Launching browser at %s", url)
	cmd := exec.Command(browser, url)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(output))
}

var (
	port              int
	host              string
	autoLaunchBrowser bool
)
