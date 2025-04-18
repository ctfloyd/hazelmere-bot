package user_update

import "github.com/ctfloyd/hazelmere-bot/src/internal/gain"

type UserUpdateOptions struct {
	Username string
	Time     int
	Unit     gain.TimeUnit
}
