package main
import "fmt"

type Cookie struct {
	Color string
}

func NewCookie(color string) Cookie {
	return Cookie{
		Color:color,
	}
}

func (c *Cookie) GetColor() string {
	return c.Color
}

func (c *Cookie) SetColor(color string) {
	c.Color = color
}

func main() {
	myCookie := NewCookie("green")
	fmt.Println(myCookie.Color)
	myCookieColor := myCookie.GetColor()
	fmt.Println(myCookieColor)
	myCookie.SetColor("yellow")
	fmt.Println(myCookie.Color)
}



