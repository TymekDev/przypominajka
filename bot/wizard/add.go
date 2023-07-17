package wizard

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"git.sr.ht/~tymek/przypominajka/format"
	"git.sr.ht/~tymek/przypominajka/models"
	"git.sr.ht/~tymek/przypominajka/storage"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	addStepStart int = iota
	addStepMonth
	addStepDay
	addStepType
	addStepName
	addStepSurname
	addStepDone
)

type Add struct {
	step        int
	e           models.Event
	cancelReset context.CancelFunc
}

var _ Interface = (*Add)(nil)

func (a *Add) Name() string {
	return "add"
}

func (a *Add) Start(update tg.Update) tg.Chattable {
	ctx, cancel := context.WithCancel(context.Background())
	a.cancelReset = cancel
	go func() {
		select {
		case <-ctx.Done():
		case <-time.After(30 * time.Second):
			a.Reset()
		}
	}()
	msg, _ := a.Next(nil, update)
	return msg
}

// TODO: add user error messages
// TODO: add validation
func (a *Add) Next(s storage.Interface, update tg.Update) (tg.Chattable, error) {
	defer func() { // FIXME: this probably shouldn't run on error
		a.step += 1
	}()
	log.Println("DEBUG", "step", a.step)
	log.Println("DEBUG", "callback data", update.CallbackData())

	switch a.step {
	case addStepStart:
		msg := tg.NewMessage(update.FromChat().ID, format.MessageAddStepStart)
		msg.ReplyMarkup = addKeyboardMonths
		return msg, nil
	case addStepMonth:
	case addStepDay:
	case addStepType:
	case addStepName:
	case addStepSurname:
	case addStepDone:
		return nil, ErrDone
	}
	return nil, errors.New("unknown wizard step")
}

func (a *Add) Reset() {
	if cr := a.cancelReset; cr != nil {
		cr()
	}
	a.step = addStepStart
	a.e = models.Event{}
	a.cancelReset = nil
}

var addKeyboardMonths = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("Styczeń", addCallbackMonth(1)),
		tg.NewInlineKeyboardButtonData("Luty", addCallbackMonth(2)),
		tg.NewInlineKeyboardButtonData("Marzec", addCallbackMonth(3)),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("Kwiecień", addCallbackMonth(4)),
		tg.NewInlineKeyboardButtonData("Maj", addCallbackMonth(5)),
		tg.NewInlineKeyboardButtonData("Czerwiec", addCallbackMonth(6)),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("Lipiec", addCallbackMonth(7)),
		tg.NewInlineKeyboardButtonData("Sierpień", addCallbackMonth(8)),
		tg.NewInlineKeyboardButtonData("Wrzesień", addCallbackMonth(9)),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("Październik", addCallbackMonth(10)),
		tg.NewInlineKeyboardButtonData("Listopad", addCallbackMonth(11)),
		tg.NewInlineKeyboardButtonData("Grudzień", addCallbackMonth(12)),
	),
)

func addCallbackMonth(m int) string {
	return newCallbackData(&Add{}, "month", strconv.Itoa(m))
}
