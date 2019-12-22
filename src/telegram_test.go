package telegram

import (
	"fmt"
	"testing"
)

func TestCanSendTelegramMessage(t *testing.T) {
	c := Client{
		VerboseLogging: true,
	}
	fmt.Println("Initializing")
	c.Initialize()
	fmt.Println("Sending message")
	err := c.SendMessage("test message")
	if err != nil {
		t.Error(err)
	}
}
