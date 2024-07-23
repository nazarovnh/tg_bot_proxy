package bot

import (
	"net/http"
	"tg_bot_proxy/internal/config"
	"tg_bot_proxy/internal/repository"
	"tg_bot_proxy/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	API *tgbotapi.BotAPI
	Logger     *zap.Logger
	Config     *config.Config
	Service    *service.Service
	Repository *repository.Repository
}


func Init(config *config.Config, logger *zap.Logger, service *service.Service, repository *repository.Repository) *Bot {
	return &Bot{
		Config:     config,
		Logger:     logger,
		Service:    service,
		Repository: repository,
	}
}

func (b *Bot) Run() {
	botAPI, err := b.NewBotAPI()
	if err != nil {
		b.Logger.Fatal("failed create new bot api instance", zap.String("error", err.Error()))
	}
	b.API = botAPI

	if err := b.SetWebhook(); err != nil {
		b.Logger.Fatal("failed set webhook", zap.String("error", err.Error()))
	}

	go b.StartServer()
	b.Logger.Info("http webhook server started")

	updates := b.API.ListenForWebhook("/")
	for update := range updates {
		go b.HandleUpdate(update)
	}
}

func (b *Bot) NewBotAPI() (*tgbotapi.BotAPI, error) {
	botAPI, err := tgbotapi.NewBotAPI(b.Config.Bot.Token)
	if err != nil {
		return nil, err
	}
	b.Logger.Info("Bot api instance are succesful created", zap.String("account", botAPI.Self.UserName))
	return botAPI, nil
}

func (b *Bot) SetWebhook() error {
	webhook, err := tgbotapi.NewWebhook(b.Config.Bot.WebhookLink + b.API.Token)
	if err != nil {
		return err
	}
	_, err = b.API.Request(webhook)
	if err != nil {
		return err
	}
	info, err := b.API.GetWebhookInfo()
	if err != nil {
		return err
	}
	b.Logger.Info("webhook info", zap.Any("webhook", info))
	if info.LastErrorDate != 0 {
		return err
	}
	return nil
}

func (b *Bot) StartServer(){
	addr := b.Config.Server.Host+b.Config.Server.Port
	if err := http.ListenAndServe(addr, nil); err != nil {
		b.Logger.Fatal("failed start http server", zap.String("error", err.Error()))
	}
}