package eventhandlers_test

import (
	"pipelinemkv/makemkv/commands/eventhandlers"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRegisterMakeMkvEventHandlerInvalidKey(t *testing.T) {
	input := "Key not found or invalid"
	responseCode := eventhandlers.HandleRegisterMakeMkvEvents(strings.NewReader(input))

	assert.Equal(t, 400, responseCode)
}

func TestRegisterMakeMkvEventHandlerValidKey(t *testing.T) {
	input := "Registration key saved -- some extra dummy data"
	responseCode := eventhandlers.HandleRegisterMakeMkvEvents(strings.NewReader(input))

	assert.Equal(t, 200, responseCode)
}

func TestRegisterMakeMkvEventHandlerValidKeyKeyAlreadyExists(t *testing.T) {
	input := `Current Key already exists: asdljasdlkjalkdjlk
Registration key saved -- some extra dummy data`
	responseCode := eventhandlers.HandleRegisterMakeMkvEvents(strings.NewReader(input))

	assert.Equal(t, 200, responseCode)
}
