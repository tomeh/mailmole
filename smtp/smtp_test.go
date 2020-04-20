package smtp

import (
	"testing"
)

func TestNewSmtp(t *testing.T) {
	s := NewSmtp()
	s.assertState(StateInit, t)
}

// StateInit tests.

func TestStateInitEventHelo(t *testing.T) {
	s := NewSmtp()
	err := s.Transition(eventHelo)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateHelo, t)
}

func TestStateInitEventEhlo(t *testing.T) {
	s := NewSmtp()
	err := s.Transition(eventEhlo)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateEhlo, t)
}

var stateInitInvalidEvents = []event{
	eventMail,
	eventRcpt,
	eventData,
	eventRset,
	eventPnnd,
	eventQuit,
}
func TestStateInitInvalidEvent(t *testing.T) {
	s := NewSmtp()
	for _, event := range stateInitInvalidEvents {
		err := s.Transition(event)
		if err == nil {
			t.Fatalf("Expected error on state %s for event %s but got nil", StateInit, event)
		}
	}
	s.assertState(StateInit, t)
}

// StateHelo tests.

func TestStateHeloEventQuit(t *testing.T) {
	s := NewSmtp()
	s.State = StateHelo
	err := s.Transition(eventQuit)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateQuit, t)
}

func TestStateHeloEventMail(t *testing.T) {
	s := NewSmtp()
	s.State = StateHelo
	err := s.Transition(eventMail)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateMail, t)
}

var stateHeloInvalidEvents = []event{
	eventHelo,
	eventEhlo,
	eventRcpt,
	eventData,
	eventPnnd,
}
func TestStateHeloInvalidEvent(t *testing.T) {
	s := NewSmtp()
	s.State = StateHelo
	for _, event := range stateHeloInvalidEvents {
		err := s.Transition(event)
		if err == nil {
			t.Fatalf("Expected error on state %s for event %s but got nil", StateHelo, event)
		}
	}
	s.assertState(StateHelo, t)
}

// StateEhlo tests.

func TestStateEhloEventQuit(t *testing.T) {
	s := NewSmtp()
	s.State = StateEhlo
	err := s.Transition(eventQuit)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateQuit, t)
}

func TestStateEhloEventMail(t *testing.T) {
	s := NewSmtp()
	s.State = StateEhlo
	err := s.Transition(eventMail)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateMail, t)
}

var stateEhloInvalidEvents = []event{
	eventHelo,
	eventEhlo,
	eventRcpt,
	eventData,
	eventPnnd,
}
func TestStateEhloInvalidEvent(t *testing.T) {
	s := NewSmtp()
	s.State = StateEhlo
	for _, event := range stateEhloInvalidEvents {
		err := s.Transition(event)
		if err == nil {
			t.Fatalf("Expected error on state %s for event %s but got nil", StateEhlo, event)
		}
	}
	s.assertState(StateEhlo, t)
}

// StateMail tests.

func TestStateMailEventQuit(t *testing.T) {
	s := NewSmtp()
	s.State = StateMail
	err := s.Transition(eventQuit)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateQuit, t)
}

func TestStateMailEventRcpt(t *testing.T) {
	s := NewSmtp()
	s.State = StateMail
	err := s.Transition(eventRcpt)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateRcpt, t)
}

var stateMailInvalidEvents = []event{
	eventHelo,
	eventEhlo,
	eventMail,
	eventData,
	eventPnnd,
}
func TestStateMailInvalidEvents(t *testing.T) {
	s := NewSmtp()
	s.State = StateMail
	for _, event := range stateMailInvalidEvents {
		err := s.Transition(event)
		if err == nil {
			t.Fatalf("Expected error on state %s for event %s but got nil", StateMail, event)
		}
	}
	s.assertState(StateMail, t)
}

// StateRcpt tests.

func TestStateRcptEventQuit(t *testing.T) {
	s := NewSmtp()
	s.State = StateRcpt
	err := s.Transition(eventQuit)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateQuit, t)
}

func TestStateRcptEventRcpt(t *testing.T) {
	s := NewSmtp()
	s.State = StateRcpt
	err := s.Transition(eventRcpt)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateRcpt, t)
}

var stateRcptInvalidEvents = []event{
	eventHelo,
	eventEhlo,
	eventMail,
	eventPnnd,
}
func TestStateRcptInvalidEvents(t *testing.T) {
	s := NewSmtp()
	s.State = StateRcpt
	for _, event := range stateRcptInvalidEvents {
		err := s.Transition(event)
		if err == nil {
			t.Fatalf("Expected error on state %s for event %s but got nil", StateRcpt, event)
		}
	}
	s.assertState(StateRcpt, t)
}

// StateData tests.

func TestStateDataEventSend(t *testing.T) {
	s := NewSmtp()
	s.State = StateData
	err := s.Transition(eventSend)
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateComp, t)
}

var stateDataInvalidEvents = []event{
	eventHelo,
	eventEhlo,
	eventMail,
	eventRcpt,
	eventData,
	eventRset,
	eventPnnd,
	eventQuit,
}
func TestStateDataInvalidEvents(t *testing.T) {
	s := NewSmtp()
	s.State = StateData
	for _, event := range stateDataInvalidEvents {
		err := s.Transition(event)
		if err == nil {
			t.Fatalf("Expected error on state %s for event %s but got nil", StateData, event)
		}
	}
	s.assertState(StateData, t)
}

// StateComplete tests.

var stateCompleteInvalidEvents = []event{
	eventHelo,
	eventEhlo,
	eventMail,
	eventRcpt,
	eventData,
	eventRset,
	eventPnnd,
	eventSend,
	eventQuit,
}
func TestStateCompleteInvalidEvents(t *testing.T) {
	s := NewSmtp()
	s.State = StateComp
	for _, event := range stateCompleteInvalidEvents {
		err := s.Transition(event)
		if err == nil {
			t.Fatalf("Expected error on state %s for event %s but got nil", StateComp, event)
		}
	}
	s.assertState(StateComp, t)
}

// StateQuit tests.

var stateQuitInvalidEvents = []event{
	eventHelo,
	eventEhlo,
	eventMail,
	eventRcpt,
	eventData,
	eventRset,
	eventPnnd,
	eventSend,
	eventQuit,
}
func TestStateQuitInvalidEvents(t *testing.T) {
	s := NewSmtp()
	s.State = StateQuit
	for _, event := range stateQuitInvalidEvents {
		err := s.Transition(event)
		if err == nil {
			t.Fatalf("Expected error on state %s for event %s but got nil", StateQuit, event)
		}
	}
	s.assertState(StateQuit, t)
}

func (smtp *Smtp) assertState(s State, t *testing.T) {
	if !smtp.Is(s) {
		t.Errorf("Expected state to be %s, got %s", s, smtp.State)
	}
}
