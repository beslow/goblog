package helpers

import (
	"fmt"
	"time"

	"github.com/beslow/goblog/models"
	"github.com/ungerik/go-gravatar"
)

func ToHashID(id int) string {
	h, _ := models.HashID.Encode([]int{id})

	return h
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()

	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func AvatarURL(email string) string {
	if email == "" {
		return "/public/static/images/avatar/2.jpg"
	} else {
		return gravatar.Url(email)
	}
}
