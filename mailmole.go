// Mailmole is an email testing tool.

package main

import (
	"flag"
	"fmt"
	smtp "github.com/tomeh/mailmole/smtp"
	"os/signal"

	//"io/ioutil"
	"log"
	//"net/mail"
	"os"
	"os/exec"
	"runtime"
	//"github.com/tomeh/mailmole/web"
)

func init() {
	flags()
	log.SetOutput(os.Stdout)
}

func flags() {
	flag.BoolVar(&http, "http", true, "Launch http server (default: true)")

	flag.IntVar(&port, "port", 8084, "The port at which to serve http.")
	flag.StringVar(&host, "host", "127.0.0.1", "The host at which to serve http.")
	flag.BoolVar(&autoLaunchBrowser, "launchBrowser", true, "Open a browser session (default: true)")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	flag.Parse()

	//var listeners []contracts.Sub
	// Register listeners globally

	//if !quiet {
	//	// Create the console listener.
	//	con := console.NewConsole()
	//	listeners = append(listeners, con)
	//}

	//if http {
	//	// Launch the http server.
	//	httpServer := web.NewServer(host, port)
	//
	//	if autoLaunchBrowser {
	//		go launchBrowser(httpServer.GetBaseUrl())
	//	}
	//
	//	go httpServer.Start()
	//}

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	smtpServer := smtp.NewServer(smtp.ServerConfig{
		Addr:     "0.0.0.0",
		Port:     2525,
		HostName: hostName,
	})

	// Move these to the server struct
	start := make(chan bool)
	stop := make(chan bool)
	go smtpServer.ListenAndServe(start, stop)

	start <- true

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func(){
		for _ = range c {
			// sig is a ^C, handle it
			fmt.Println("Shutting down ...")
			stop <- true
		}
	}()

	<-stop

	fmt.Println("Done. Have a nice day!")
}

//func mailHandler(msg *mail.Message) {
//	b, err := ioutil.ReadAll(msg.Body)
//
//	if err != nil {
//		log.Println(fmt.Sprintf("Error reading message body: %s", err))
//		return
//	}
//
//	log.Println(string(b))
//}

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
	http bool

	port              int
	host              string
	autoLaunchBrowser bool

	// We hold a reference to exec.Command here so we can mock the function
	// in testing. See launchBrowser.
	execCommand = exec.Command
)
