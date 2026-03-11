package streamtracker

import (
	"context"
	"io"
)

type streamTuple struct {
	Reader     *io.Reader
	CancelFunc context.CancelFunc
}

/*
A situation can arrise when we need to re-attach onto the output of a command in makemkv.
Hence, we create an instance of
*/
type StreamTracker struct {
	streamSet map[string]streamTuple
}

func (st *StreamTracker) AddStream(key string, reader *io.Reader, cancelFunc context.CancelFunc) error {
	st.streamSet[key] = streamTuple{
		Reader:     reader,
		CancelFunc: cancelFunc,
	}
	return nil
}

func (st *StreamTracker) GetStream(key string) (*io.Reader, bool) {
	reader, ok := st.streamSet[key]
	return reader.Reader, ok
}

func (st *StreamTracker) GetStreamCancelFunc(key string) context.CancelFunc {
	reader := st.streamSet[key]
	return func() {
		delete(st.streamSet, key)
		reader.CancelFunc()
	}
}

func (st *StreamTracker) RemoveStream(key string) {
	delete(st.streamSet, key)
}

func NewStreamTracker() StreamTracker {
	return StreamTracker{
		streamSet: map[string]streamTuple{},
	}
}
