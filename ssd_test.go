package ssd1306

import (
	"testing"
	"time"

	"gobot.io/x/gobot/platforms/raspi"
)

func TestText(t *testing.T) {
	r := raspi.NewAdaptor()
	r.Connect()
	oled := NewSundingOLED("oled", r)
	oled.ResetImage().Message(10, 10, "Pussy").Render()
	<-time.After(10 * time.Second)
	oled.Clean()
}

func TestRectangle(t *testing.T) {
	r := raspi.NewAdaptor()
	r.Connect()
	oled := NewSundingOLED("oled", r)
	oled.ResetImage().Rectangle(0, 0, 10, 10, 0).Render()
	<-time.After(10 * time.Second)
	oled.ResetImage().Render().Rectangle(0, 0, 10, 10, 1).Render()
	<-time.After(10 * time.Second)
	oled.Clean()
}
