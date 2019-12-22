package telegram

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Options holds the configuration options for sending Telegram messages
type Options struct {
	EnableTelegram bool   `json:"enableTelegram"` // Enable Telegram messages
	BotID          string `json:"botId"`          // Bot Identifier
	ChatID         string `json:"chatId"`         // Chat Identifier
}

// ReadFromFile will read the configuration settings from the specified file
func (o *Options) ReadFromFile(path string) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		b, err := ioutil.ReadFile(path)
		if err == nil {
			err = json.Unmarshal(b, &o)
		}
	}
	return err
}

// WriteToFile will write the configuration settings to the specified file
func (o *Options) WriteToFile(path string) error {
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0666)
}
