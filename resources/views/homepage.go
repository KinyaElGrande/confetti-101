package views

import (
	"confetti-framework/config"
	"github.com/confetti-framework/contract/inter"
)

type HomepageView struct {
	Title       string
	Description string
	Locale      string
}

func Homepage(app inter.App, title string, description string) *HomepageView {
	return &HomepageView{
		Title:       title,
		Description: description,
		Locale:      Locale(app),
	}
}

func (h HomepageView) Template() string {
	return config.Path.Views + "/homepage.gohtml"
}
