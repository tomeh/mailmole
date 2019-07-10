package smtp

import (
	"testing"
)


func TestNewSmtpStateMachine(t *testing.T) {
	NewSmtpStateMachine()
}

func TestEventHelo(t *testing.T) {
	s := NewSmtpStateMachine()
	err := s.S.Event(string(eventHelo))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateHelo, t)
}

func TestEventEhlo(t *testing.T) {
	s := NewSmtpStateMachine()
	err := s.S.Event(string(eventHelo))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateHelo, t)
}

func TestEventMail(t *testing.T) {
	s := NewSmtpStateMachine()
	s.S.SetState(string(StateHelo))
	err := s.S.Event(string(eventMail))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateMail, t)
}

func TestEventRcpt(t *testing.T) {
	s := NewSmtpStateMachine()
	s.S.SetState(string(StateMail))
	err := s.S.Event(string(eventRcpt))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateRcpt, t)
}

func TestEventData(t *testing.T) {
	s := NewSmtpStateMachine()
	s.S.SetState(string(StateRcpt))
	err := s.S.Event(string(eventData))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateData, t)
}

func TestEventMailRset(t *testing.T) {
	s := NewSmtpStateMachine()
	s.S.SetState(string(StateMail))
	err := s.S.Event(string(eventRset))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateHelo, t)
}

func TestEventRcptRset(t *testing.T) {
	s := NewSmtpStateMachine()
	s.S.SetState(string(StateRcpt))
	err := s.S.Event(string(eventRset))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateHelo, t)
}

func TestEventPnnd(t *testing.T) {
	s := NewSmtpStateMachine()
	s.S.SetState(string(StateData))
	err := s.S.Event(string(eventPnnd))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateHelo, t)
}

func TestEventInitQuit(t *testing.T) {
	s := NewSmtpStateMachine()
	s.S.SetState(string(StateInit))
	err := s.S.Event(string(eventQuit))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateFini, t)
}

func TestEventHeloQuit(t *testing.T) {
	s := NewSmtpStateMachine()
	s.S.SetState(string(StateHelo))
	err := s.S.Event(string(eventQuit))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateFini, t)
}

func TestEventMailQuit(t *testing.T) {
	s := NewSmtpStateMachine()
	s.S.SetState(string(StateMail))
	err := s.S.Event(string(eventQuit))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateFini, t)
}

func TestEventRcptQuit(t *testing.T) {
	s := NewSmtpStateMachine()
	s.S.SetState(string(StateRcpt))
	err := s.S.Event(string(eventQuit))
	if err != nil {
		t.Error(err)
	}
	s.assertState(StateFini, t)
}

func (fsm *StateMachine) assertState(state State, t *testing.T) {
	if !fsm.S.Is(string(state)) {
		t.Errorf("Expected state to be %s, got %s", string(state), fsm.S.Current())
	}
}
