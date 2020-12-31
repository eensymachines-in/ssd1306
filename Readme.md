### Sunding 1.9" OLED 
-----------

Its a handy display which suffices for most of the home automation applications. We here have the basic functions to draw some 2D text / geometry on it. 
`Render` and `ResetImage` functions are a bit tricky if used in a loop. So here is some basic guideline.

OLED functions can be chained, and that is a convinience that we can use to our advantage.. 

```go
import (
    "time"
	"gobot.io/x/gobot/platforms/raspi"
)

func RenderText(){
    r := raspi.NewAdaptor()
	r.Connect()
	oled := NewSundingOLED("oled", r)
	oled.ResetImage().Message(10, 10, "Hello world").Render()
	<-time.After(10 * time.Second)
	oled.Clean()
}

```
Above is we trying to display some random text at point 10,10 on the scren for about 10 seconds.
Please see, `Render` is necessay, else you'd end up only changing the image with nothing on the screen. Render paints the image in the memory to the actual screen.
A reason behind this is, the image is updated only in pieces and rendered once the loop is about to repeat. Cases as these would need the Render function to be distinct.

```go 
import (
    "time"
	"gobot.io/x/gobot/platforms/raspi"
)

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
```
> Rectangle points are top-left (x1,y1) and bottom right (x2,y2)

Now we see how we can draw rectangles, given 2 corner co-ordinates. `fill` as int is to blob fill the rectangle else there is only the outline of the rectangle

```go
func (disp *Sunding19OLED) HLine(x1, x2, y int) *Sunding19OLED
func (disp *Sunding19OLED) VLine(x, y1, y2 int) *Sunding19OLED 
```
This can draw a simple horizontal / vertical line, 1px thick for the given co-ordintes 
We havent got many examples for this since this is more likely used internally for drawing rectangle rather than lines