package eventhandlers

import (
	"fmt"
	"io"
	"pipelinemkv/makemkv/streamReader"
	"strings"
)

const registerMkvKeyBadKeyPrefix string = "Key not found or invalid"
const registerMkvKeySavedPrefix string = "Registration key saved"

func HandleRegisterMakeMkvEvents(reader io.Reader) int {
	c := streamReader.ReadStream(reader)
	for s := range c {
		s = strings.TrimSpace(s)
		switch {
		case strings.HasPrefix(s, registerMkvKeyBadKeyPrefix):
			fmt.Println(s)
			return 400
		case strings.HasPrefix(s, registerMkvKeySavedPrefix):
			fmt.Println(s)
			return 200
		default:
			fmt.Println(s)
		}
	}
	return 500
}
