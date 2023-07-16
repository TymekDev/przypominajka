package wizard

import (
	"errors"

	"git.sr.ht/~tymek/przypominajka/storage"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const CallbackSep = ":"

var ErrDone = errors.New("wizard: done")

type Interface interface {
	Name() string
	Start() tg.Chattable
	Next(s storage.Interface, update tg.Update) (tg.Chattable, error)
	Reset()
}
