package commands_test

import (
	"servermakemkv/commands"
	"testing"
)

func TestGetInfoCommand(t *testing.T) {

}

func TestRegisterMkvKey(t *testing.T) {
	key := "VALIDKEY"

	commands.RegisterMkvKey(key)
}
