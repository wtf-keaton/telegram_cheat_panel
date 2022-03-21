package telegram

import (
	"log"
	"strconv"
	"telegram_webpanel/internal/dbApi"
	"telegram_webpanel/pgk/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func hasAccessToBot(e int64) bool {
	for _, a := range config.ApplicationConfig.AccessRights {
		if a == e {
			return true
		}
	}
	return false
}

func handleCommands(message tgbotapi.Message, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	switch message.Command() {

	case "new_cheat":
		if len(message.CommandArguments()) == 0 {
			msg.Text = "Для создания чита введи /new_cheat [НАЗВАНИЕ ЧИТА]"
		} else {
			dbApi.CreateCheat(message.CommandArguments())
			msg.Text = "Чит \"" + message.CommandArguments() + "\" Создан"
		}
		bot.Send(msg)
	/*SCUM*/
	case "generate_key_GAME30":
		if len(message.CommandArguments()) == 0 {
			msg.Text = "Для создания ключа введи /generate_key_[GAME NAME]30 [КОЛИЧЕСТВО]"
			bot.Send(msg)

		} else {
			amount, _ := strconv.Atoi(message.CommandArguments())

			for i := 0; i < amount; i++ {
				res := dbApi.CreateKey(30, 1)
				msg.Text = res
				bot.Send(msg)
			}
		}
	case "generate_key_GAME7":
		if len(message.CommandArguments()) == 0 {
			msg.Text = "Для создания ключа введи /generate_key_[GAME NAME]14 [КОЛИЧЕСТВО]"
			_, _ = bot.Send(msg)

		} else {
			amount, _ := strconv.Atoi(message.CommandArguments())

			for i := 0; i < amount; i++ {
				res := dbApi.CreateKey(7, 1)
				msg.Text = res
				_, _ = bot.Send(msg)
			}
		}
	case "generate_key_GAME1":
		if len(message.CommandArguments()) == 0 {

		} else {
			amount, _ := strconv.Atoi(message.CommandArguments())

			for i := 0; i < amount; i++ {
				res := dbApi.CreateKey(1, 1)
				msg.Text = res
				_, _ = bot.Send(msg)
			}
		}

	case "generate_key":
		msg.Text = "Для создания ключа введи /generate_key_[GAME NAME][ДНИ(1, 7, 30)] [КОЛИЧЕСТВО]\n" +
			"For create key enter: /generate_key_[GAME NAME][DAYS(1, 7, 30)] [COUNT]"
		_, _ = bot.Send(msg)
	case "reset_hwid":
		if len(message.CommandArguments()) == 0 {
			msg.Text = "Для сброса HWID'a введи /reset_hwid [КЛЮЧ]\n" + "For HWID reset enter: /reset_hwid [KEY]"
		} else {
			dbApi.ResetHWID(message.CommandArguments())
			msg.Text = "HWID ключа: \"" + message.CommandArguments() + "\" успешно сброшен"
		}
		_, _ = bot.Send(msg)

	case "ban_key":
		if len(message.CommandArguments()) == 0 {
			msg.Text = "Для бана ключа введи /ban_key [КЛЮЧ]\n" + "For ban key enter /ban_key [KEY]"
		} else {
			dbApi.BanKey(message.CommandArguments())
			msg.Text = "Ключ: \"" + message.CommandArguments() + "\" успешно забанен"
		}
		_, _ = bot.Send(msg)

	case "add_days_all":
		msg.Text = "Дни добавлены всем! Наебал, ещё не работает"

		_, _ = bot.Send(msg)
	default:
		msg.Text = "Я не знаю такую команду, шизоид"
		_, _ = bot.Send(msg)
	}

}

func HandleTelegram() {
	bot, err := tgbotapi.NewBotAPI("BOT TOKEN")

	if err != nil {
		log.Panic("Error: ", err)
	}

	log.Printf("BOT Logined on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil && hasAccessToBot(update.Message.Chat.ID) {
			if !update.Message.IsCommand() { // ignore any non-command Messages
				continue
			}
			handleCommands(*update.Message, bot)
		} else if !hasAccessToBot(update.Message.Chat.ID) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ты сука чё забыл тут шизоид ебаный")

			bot.Send(msg)
		}
	}
}
