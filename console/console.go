// Allows viewing trapped emails through the console output.
package console

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
)

func NewConsole () *Listener {
	return &Listener{}
}

type Listener struct {

}

func (l *Listener) Pub(m *mail.Message) {
	b, err := ioutil.ReadAll(m.Body)

	if err != nil {
		log.Println(fmt.Sprintf("Error reading message body: %s", err))
		return
	}

	log.Println(string(b))
}

