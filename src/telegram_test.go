package telegram

import (
	"fmt"
	"testing"
)

func TestCanSendTelegramMessageWithInitialize(t *testing.T) {
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

func TestCanSendTelegramMessageWithoutInitialize(t *testing.T) {
	c := Client{
		VerboseLogging: true,
	}
	fmt.Println("Sending message")
	err := c.SendMessage("test message")
	if err != nil {
		t.Error(err)
	}
}
