package characters

import (
	"context"
	"nutcracker/domain/data"
	"sync"
	"time"
)

type Person interface {
	StartHearing(ctx context.Context)
	GetName() PersonName
	Speak(message data.Message)
}

type PersonName string

func NewMary(ctx context.Context, incoming chan data.Message, wg *sync.WaitGroup) *Mary {
	mary := Mary{
		incoming: incoming,
		wg: wg,
	}

	mary.StartHearing(ctx)
	mary.startCheckingYourself(ctx)

	return &mary
}

type Mary struct {
	happinessIndex HappinessIndex
	faceColor HumanFaceColorVariation
	incoming chan data.Message
	wg *sync.WaitGroup
	name PersonName
}

func (m Mary) GetName() PersonName {
	return m.name
}

func (m Mary) Speak(message data.Message) {
	m.incoming <- message
}


func (m Mary) GetState() HappinessState {
	switch i := m.happinessIndex; {
	case i >= 10 && i <100:
		return IsFine
	case i >-10 && i < 10:
		return IsNeutral
	case i >-100 && i <= -10:
		return IsSad
	case i <= -100:
		return IsCrying
	default:
		return IsHappy
	}
}

func (m *Mary) StartHearing(ctx context.Context) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		for {
			select {
			case msg := <- m.incoming:
				m.messageInterpreter(msg)
			case <- ctx.Done():
				return
			}

		}
	}()

}

func (m Mary) GetFaceColor() HumanFaceColorVariation {
	return m.faceColor
}

func (m *Mary) messageInterpreter(msg data.Message) {
	switch msg {
	case data.MessageMouseStillUnCatch:
		m.happinessIndex = m.happinessIndex - 99
	case data.MessageCalmDownMary:
		m.happinessIndex = m.happinessIndex + 2
	}
}

func (m *Mary) startCheckingYourself(ctx context.Context) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <- ticker.C:
				m.happinessStateInterpreter()
			case <- ctx.Done():
				return
			}
		}
	}()

}

func (m *Mary) happinessStateInterpreter() {
	m.changeFaceColor()
}

func (m *Mary) changeFaceColor() {
	switch m.GetState() {
	case IsCrying: m.faceColor = RedHumanFaceColor
	case IsSad:    m.faceColor = PaleHumanFaceColor
	case IsNeutral: m.faceColor = FleshHumanFaceColor
	default:
		m.faceColor = RosyHumanFaceColor
	}
}

type HappinessIndex int

type HappinessState int
const (
	IsHappy HappinessState = iota
	IsFine
	IsNeutral
	IsSad
	IsCrying

)

type HumanFaceColorVariation int
const (
	PaleHumanFaceColor HumanFaceColorVariation = iota
	GreenHumanFaceColor
	FleshHumanFaceColor
	RedHumanFaceColor
	RosyHumanFaceColor
)