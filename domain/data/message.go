package data

var (
	MessageMouseStillUnCatch Message = Message{
		Header: "злая мышь еще не поймана",
	}
	MessageCalmDownMary Message = Message{
		Header: "Полно, успокойся, милая Мари",
		Description: "Поверь, что мы прогоним всех мышей. Если не помогут мышеловки, то Фриц достанет нам своего кота",
	}
)

type Message struct {
	Header MessageHeader
	Description string
}

type MessageHeader string

