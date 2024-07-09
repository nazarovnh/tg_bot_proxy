package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleUpdate(upd tgbotapi.Update) {
	// updLocal := model.DecodeToLocal(upd)
	// if msg := upd.Message; msg != nil {
	// 	if msg.IsCommand() {
	// 		b.SendMessage(b.CommandsHandler(upd.Message.Command(), updLocal))
	// 	} else {
	// 		b.SendMessage(b.MessageHandler(upd))
	// 	}
	// }
	// if cq := upd.CallbackQuery; cq != nil {
	// 	b.SendMessage(b.CallbacksHandler(updLocal))
	// }
}


