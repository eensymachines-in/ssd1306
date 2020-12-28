package ssd1306

import (
	"image"

	"gobot.io/x/gobot/drivers/i2c"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// This provides functions to draw geometry and call initilizations on the Sunding 0.96 OLED
// This though is not the TCP server that you were expecting, that you'd have to hand roll it as per application needs

const (
	// I2Cbus : bus on which the display is expected to be connected
	I2Cbus = 1
	// I2Caddr : oled appears on this address - from experience
	I2Caddr = 0x3c
	// DisplayWidth : width in pixels that the oled is at max
	DisplayWidth = 128
	// DisplayHeight : height in pixels that the oled is at max
	DisplayHeight = 64
)

// Sunding19OLED : represents the oled as hardware
type Sunding19OLED struct {
	id     string
	i      *image.RGBA        // the image that gets rewritten and displayed
	screen *i2c.SSD1306Driver // link to the hardware
}

// NewSundingOLED : creates a new OLED device, initiates theI2C connection
func NewSundingOLED(id string, r i2c.Connector) *Sunding19OLED {
	scr := i2c.NewSSD1306Driver(r, i2c.WithBus(I2Cbus), i2c.WithAddress(I2Caddr), i2c.WithSSD1306DisplayWidth(DisplayWidth), i2c.WithSSD1306DisplayHeight(DisplayHeight))
	scr.Start()
	scr.Clear()
	return &Sunding19OLED{
		id:     id,
		i:      image.NewRGBA(image.Rect(0, 0, 128, 64)),
		screen: scr,
	}
}

/*Notice how all the functions can be chained*/

// Render : paints the image on the screen
func (disp *Sunding19OLED) Render() *Sunding19OLED {
	disp.screen.ShowImage(disp.i)
	return disp
}

// ResetImage : Changes the image, does not change anything on the screen, unless Render is called
func (disp *Sunding19OLED) ResetImage() *Sunding19OLED {
	disp.i = image.NewRGBA(image.Rect(0, 0, 128, 64))
	return disp
}

// Clean : clears the screen blank
func (disp *Sunding19OLED) Clean() *Sunding19OLED {
	disp.ResetImage()
	disp.Render()
	disp.screen.Halt()
	disp.screen.Clear()
	return disp
}

// Message : given the x y position in pixels on the display this only puts out a text
func (disp *Sunding19OLED) Message(posx, posy int, msg string) *Sunding19OLED {
	drawer := &font.Drawer{
		Dst:  disp.i,
		Src:  image.NewUniform(image.White.C),
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(posx * 116), Y: fixed.Int26_6(posy * 64)},
	}
	drawer.DrawString(msg)
	return disp
}

/*Geometric functions, this is used for drawing simple 2D geomtery on the screen*/

// HLine : draws a horizontal line, needs 2 x values and a single y value
func (disp *Sunding19OLED) HLine(x1, x2, y int) *Sunding19OLED {
	for ; x1 <= x2; x1++ {
		disp.i.Set(x1, y, image.White.C)
	}
	return disp
}

// VLine : Draws a vertical line on the image sent in uses 2 y while only one x will suffice.
func (disp *Sunding19OLED) VLine(x, y1, y2 int) *Sunding19OLED {
	for ; y1 <= y2; y1++ {
		disp.i.Set(x, y1, image.White.C)
	}
	return disp
}

// Rectangle : Can draw a thin or thick Rectangle, needs the top left and bottom right extremities as points it appends to the image argument
func (disp *Sunding19OLED) Rectangle(x1, y1, x2, y2 int, fill int) *Sunding19OLED {
	if fill == 1 {
		// just a bunch of closely drawn vertical lines
		// this is in case you need a thick Rectangle
		for ; x1 <= x2; x1++ {
			disp.VLine(x1, y1, y2)
		}
		return disp
	}
	// incase if we want to draw thin Rectangle
	disp.VLine(x1, y1, y2).VLine(x2, y1, y2).HLine(x1, x2, y1).HLine(x1, x2, y2)
	return disp
}
