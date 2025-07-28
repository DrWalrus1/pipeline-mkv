package dummymkv

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"time"
)

//go:embed diskInfoBytes.txt
var diskInfoBytes []byte

type DummyMakeMkvHandler struct{}

func (m *DummyMakeMkvHandler) LoadConfig(configPath string) error {
	return nil
}

func (m DummyMakeMkvHandler) TriggerDiskInfo(source string) (io.Reader, context.CancelFunc, error) {
	return bytes.NewReader(diskInfoBytes), func() {}, nil
}

func (h DummyMakeMkvHandler) TriggerInitialInfoLoad(timeout time.Duration) (io.Reader, context.CancelFunc, error) {
	panic("NOT IMPLEMENTED")
}

func (h DummyMakeMkvHandler) TriggerDiskBackup(decrypt bool, source string, destination string) (io.Reader, context.CancelFunc, error) {
	panic("NOT IMPLEMENTED")
}

func (h DummyMakeMkvHandler) TriggerSaveMkv(source string, title string, destination string) (io.Reader, context.CancelFunc, error) {
	return nil, func() {}, nil
}

func (h DummyMakeMkvHandler) RegisterMakeMkv(key string) error {
	return nil
}
