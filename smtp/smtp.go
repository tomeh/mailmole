package smtp

import (
	"fmt"
	"log"
)

type State string
type event string

// An FSM representing the states of an SMTP interaction.
type Smtp struct {
	State State
}

// Constructor for the SMTP FSM.
func NewSmtp() Smtp {
	return Smtp{
		StateInit,
	}
}

type transitionError struct {
	s State
	e event
}

func (e *transitionError) Error() string {
	return fmt.Sprintf("Current state %s does not understand event %s", e.s, e.e)
}

// Transition Smtp using an event.
// An error is returned if the passed event isn't valid for the current state,
// or if the current state is an end state.
func (smtp *Smtp) Transition(e event) error {
	// Check the list of end states. If current state is an end state, return an error.
	if smtp.IsInEndState() {
		return &transitionError{smtp.State, e}
	}
	stateTransitions, ok := transitions[smtp.State]
	if !ok {
		// Current state not registered in the transitions map.
		log.Fatalf("State %s not registered", smtp.State)
	}
	newState, ok := stateTransitions[e]
	if !ok {
		// Cannot apply this event to the current state.
		return &transitionError{smtp.State, e}
	}
	smtp.State = newState
	return nil
}

func (smtp *Smtp) Can(e event) bool {
	if smtp.IsInEndState() {
		return false
	}
	stateTransitions, ok := transitions[smtp.State]
	if !ok {
		return false
	}
	_, ok = stateTransitions[e]
	return ok
}

func (smtp *Smtp) Cannot(e event) bool {
	return !smtp.Can(e)
}

func (smtp *Smtp) IsInEndState() bool {
	return endStatesList.contains(smtp.State)
}

func (smtp *Smtp) Is(s State) bool {
	return smtp.State == s
}

const (
	StateInit State = "init"
	StateHelo State = "helo"
	StateEhlo State = "ehlo"
	StateMail State = "mail"
	StateRcpt State = "rcpt"
	StateData State = "data"
	StateQuit State = "quit"
	StateComp State = "complete"

	eventHelo event = "helo"
	eventEhlo event = "ehlo"
	eventMail event = "mail"
	eventRcpt event = "rcpt"
	eventData event = "data"
	eventRset event = "reset"
	eventPnnd event = "pound"
	eventSend event = "send"
	eventQuit event = "quit"
)

var transitions map[State]map[event]State
var endStatesList endStates
type endStates []State

func (endStates *endStates) contains (s State) bool {
	for _, state := range *endStates {
		if state == s {
			return true
		}
	}
	return false
}

func init() {
	endStatesList = []State{
		StateComp,
		StateQuit,
	}

	transitions = make(map[State]map[event]State)

	transitions[StateInit] = make(map[event]State)
	transitions[StateHelo] = make(map[event]State)
	transitions[StateEhlo] = make(map[event]State)
	transitions[StateMail] = make(map[event]State)
	transitions[StateRcpt] = make(map[event]State)
	transitions[StateData] = make(map[event]State)
	transitions[StateQuit] = make(map[event]State)
	transitions[StateComp] = make(map[event]State)

	transitions[StateInit][eventHelo] = StateHelo
	transitions[StateInit][eventEhlo] = StateEhlo

	transitions[StateHelo][eventMail] = StateMail
	transitions[StateHelo][eventQuit] = StateQuit
	transitions[StateHelo][eventRset] = StateHelo

	transitions[StateEhlo][eventMail] = StateMail
	transitions[StateEhlo][eventQuit] = StateQuit
	transitions[StateEhlo][eventRset] = StateEhlo

	transitions[StateMail][eventRcpt] = StateRcpt
	transitions[StateMail][eventQuit] = StateQuit
	transitions[StateMail][eventRset] = StateHelo // Not sure if this (and others) are suitable? Do we need to reset to Ehlo if that was used?

	transitions[StateRcpt][eventRcpt] = StateRcpt
	transitions[StateRcpt][eventQuit] = StateQuit
	transitions[StateRcpt][eventRset] = StateHelo // And again.

	transitions[StateData][eventSend] = StateComp
}
