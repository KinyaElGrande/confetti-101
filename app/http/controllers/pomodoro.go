package controllers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"
)

func Pomodoro(request inter.Request) inter.Response {
	return outcome.Json([]int{1, 2, 3})
}
