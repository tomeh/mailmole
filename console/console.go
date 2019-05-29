// Allows viewing trapped emails through the console output.
package console

import "log"

func NewConsole () *Listener {
	return &Listener{}
}

type Listener struct {

}

func (l *Listener) Publish (s string) {
	log.Println(s)
}

