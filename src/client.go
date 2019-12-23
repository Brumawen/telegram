package telegram

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/kardianos/service"
)

// Client is used to send telegram messages
type Client struct {
	Options        *Options       // Telegram options
	Logger         service.Logger // logger
	VerboseLogging bool           // Switch on verbose logging
}

// Initialize initializes the telegram client
func (c *Client) Initialize() error {
	if c.Options == nil {
		c.Options = &Options{}
	}
	err := c.Options.ReadFromFile("telegramoptions.json")
	if err != nil {
		c.logError("Error reading telegram options file.", err.Error())
	}
	return err
}

// SendMessage sends  the specified message to the telegram chat specified in Options
func (c *Client) SendMessage(m string) error {
	if c.Options == nil {
		err := c.Initialize()
		if err != nil {
			return err
		}
	}

	if c.Options.BotID == "" {
		return errors.New("BotID has not been specified")
	}
	if c.Options.ChatID == "" {
		return errors.New("ChatID has not been specified")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?text=%s&chat_id=%s", c.Options.BotID, m, c.Options.ChatID)
	//c.logDebug("Sending URL ", url)
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return err
	}

	result := Response{}
	result.ReadFrom(response.Body)

	if result.OK {
		return nil
	}
	c.logError("Error sending telegram message.", result.Description, "[", result.ErrorCode, "]")
	return fmt.Errorf("%s [Error Code: %d]", result.Description, result.ErrorCode)
}

func (c *Client) logDebug(v ...interface{}) {
	if c.VerboseLogging {
		a := fmt.Sprint(v...)
		if c.Logger == nil {
			log.Println(a)
		} else {
			c.Logger.Info("Telegram: ", a)
		}
	}
}

func (c *Client) logError(v ...interface{}) {
	a := fmt.Sprint(v...)
	if c.Logger == nil {
		log.Println(a)
	} else {
		c.Logger.Error("Telegram: ", a)
	}
}
