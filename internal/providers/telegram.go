package providers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Habeebamoo/tunnl-backend/internal/configs"
	"github.com/Habeebamoo/tunnl-backend/internal/models"
)

type TelegramProvider struct {
	botToken string
	client   *http.Client
}

func NewTelegramProvider(cfg *configs.Config) *TelegramProvider {
	return &TelegramProvider{
		botToken: cfg.TelegramBotToken,
		client:   &http.Client{},
	}
}

func (t *TelegramProvider) Send(ctx context.Context, n models.Notification) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.botToken)

	resp, err := t.client.PostForm(apiURL, url.Values{
		"chat_id": {n.To},
		"text":    {fmt.Sprintf("*%s*\n\n%s", n.Title, n.Body)},
		"parse_mode": {"Markdown"},
	})
	if err != nil {
		return fmt.Errorf("telegram request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram returned status: %d", resp.StatusCode)
	}

	return nil
}