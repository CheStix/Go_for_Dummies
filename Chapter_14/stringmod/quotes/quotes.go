package quotes

import (
	"github.com/hackebrot/turtle"
)

func GetEmoji(s string) string {
	emoji, ok := turtle.Emojis[s]
	if !ok {
		return ""
	}
	return emoji.Char
}
