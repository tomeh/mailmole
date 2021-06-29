// Mailmole is an email testing tool.

package main

import (
	"flag"
	"fmt"

	"github.com/tomedharris/mailmole/browser"
	"github.com/tomedharris/mailmole/smtp"
	"github.com/tomedharris/mailmole/web"
	"os/signal"

	//"io/ioutil"
	"log"
	//"net/mail"
	"os"
)

func init() {
	flags()
	log.SetOutput(os.Stdout)
}

func flags() {
	flag.BoolVar(&http, "http", true, "Launch http server (default: true)")
	flag.IntVar(&httpPort, "http_port", 8080, "The port at which to serve http.")
	flag.StringVar(&httpHost, "http_host", "127.0.0.1", "The host at which to serve http.")
	flag.BoolVar(&autoLaunchBrowser, "launchBrowser", true, "Open a browser session (default: true)")

	flag.IntVar(&smtpPort, "smtp_port", 2525, "SMTP Port.")
	flag.StringVar(&smtpHost, "smtp_host", "0.0.0.0", "Bound IP for SMTP connections.")

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

	if http {
		// Launch the http server.
		httpServer := web.NewServer(httpHost, httpPort)

		if autoLaunchBrowser {
			go browser.LaunchBrowser(httpServer.GetBaseUrl())
		}

		go httpServer.Start()
	}

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	smtpServer := smtp.NewServer(smtp.ServerConfig{
		Addr:     smtpHost,
		Port:     smtpPort,
		HostName: hostName,
	})

	// Move these to the server struct
	started, stop := make(chan bool), make(chan bool)
	go smtpServer.ListenAndServe(started, stop)

	<-started

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func(){
		for range c {
			// sig is a ^C, handle it
			log.Println("^C detected - Shutting down")
			stop <- true
		}
	}()

	<-stop

	log.Println("Done. Have a nice day!")
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

var (
	http              bool
	httpPort          int
	httpHost          string
	autoLaunchBrowser bool

	smtpPort int
	smtpHost string
)
