package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Comment(message string) revel.Result {
	return c.Render(message)
}

func (c App) SendComment(message string) revel.Result {
    c.Validation.Required(message).Message("Your message is required!")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Comment)
    }

    return c.Render(message)
}