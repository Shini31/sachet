package main

import (
	"io/ioutil"

	"github.com/Shini31/sachet/provider/aspsms"
	"github.com/Shini31/sachet/provider/cm"
	"github.com/Shini31/sachet/provider/exotel"
	"github.com/Shini31/sachet/provider/freemobile"
	"github.com/Shini31/sachet/provider/infobip"
	"github.com/Shini31/sachet/provider/mediaburst"
	"github.com/Shini31/sachet/provider/messagebird"
	"github.com/Shini31/sachet/provider/nexmo"
	"github.com/Shini31/sachet/provider/otc"
	"github.com/Shini31/sachet/provider/sipgate"
	"github.com/Shini31/sachet/provider/telegram"
	"github.com/Shini31/sachet/provider/turbosms"
	"github.com/Shini31/sachet/provider/twilio"

	"github.com/prometheus/alertmanager/template"
	"gopkg.in/yaml.v2"
)

type ReceiverConf struct {
	Name     string
	Provider string
	To       []string
	From     string
	Text     string
}

var config struct {
	Providers struct {
		MessageBird messagebird.MessageBirdConfig
		Nexmo       nexmo.NexmoConfig
		Twilio      twilio.TwilioConfig
		Infobip     infobip.InfobipConfig
		Exotel      exotel.ExotelConfig
		CM          cm.CMConfig
		Telegram    telegram.TelegramConfig
		Turbosms    turbosms.TurbosmsConfig
		OTC         otc.OTCConfig
		MediaBurst  mediaburst.MediaBurstConfig
		FreeMobile  freemobile.Config
		AspSms      aspsms.Config
		Sipgate     sipgate.Config
	}

	Receivers []ReceiverConf
	Templates []string
}
var tmpl *template.Template

// LoadConfig loads the specified YAML configuration file.
func LoadConfig(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}

	tmpl, err = template.FromGlobs(config.Templates...)
	return err
}
