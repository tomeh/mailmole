// Mailmole is an email testing tool.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/tomedharris/mailmole/console"
	"github.com/tomedharris/mailmole/contracts"
	"github.com/tomedharris/mailmole/smtp"
	"github.com/tomedharris/mailmole/web"
)

func init() {
	flags()
	log.SetOutput(os.Stdout)
}

func flags() {
	flag.BoolVar(&quiet, "quiet", false, "Dont log messages to console.")
	flag.BoolVar(&http, "http", true, "Launch http server (default: true)")

	flag.IntVar(&port, "port", 8084, "The port at which to serve http.")
	flag.StringVar(&host, "host", "127.0.0.1", "The host at which to serve http.")
	flag.BoolVar(&autoLaunchBrowser, "launchBrowser", true, "Open a browser session (default: true)")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	flag.Parse()

	// Create smtp server with a channel to subscribe listeners.
	var listeners []contracts.Listener
	sub := make(chan contracts.Listener)
	smtpServer := smtp.NewServer(sub)

	if !quiet {
		// Create the console listener.
		con := console.NewConsole()
		sub <- con
		listeners = append(listeners, con)
	}

	if http {
		// Launch the http server.
		httpServer := web.NewServer(host, port)

		if autoLaunchBrowser {
			go launchBrowser(httpServer.GetBaseUrl())
		}

		go httpServer.Start()
	}

	smtpServer.Start()
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
	cmd := execCommand(browser, url)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(output))
}

var (
	quiet bool
	http  bool

	port              int
	host              string
	autoLaunchBrowser bool

	// We hold a reference to exec.Command here so we can mock the function
	// in testing. See launchBrowser.
	execCommand = exec.Command
)
