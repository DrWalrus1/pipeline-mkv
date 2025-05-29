package makemkv

import (
	"context"
	"io"
)

type Source = string

type streamTuple struct {
	Reader     *io.Reader
	CancelFunc context.CancelFunc
}

/*
A situation can arrise when we need to reattach onto the output of a command in makemkv.
Hence, we create an instance of
*/
type StreamTracker struct {
	streamSet map[Source]streamTuple
}

func (st *StreamTracker) AddStream(source Source, reader *io.Reader, cancelFunc context.CancelFunc) error {
	st.streamSet[source] = streamTuple{
		Reader:     reader,
		CancelFunc: cancelFunc,
	}
	return nil
}

func (st *StreamTracker) GetStream(source Source) (*io.Reader, bool) {
	reader, ok := st.streamSet[source]
	return reader.Reader, ok
}

func (st *StreamTracker) GetStreamCancelFunc(source string) context.CancelFunc {
	reader := st.streamSet[source]
	return func() {
		delete(st.streamSet, source)
		reader.CancelFunc()
	}
}

func (st *StreamTracker) RemoveStream(source Source) {
	delete(st.streamSet, source)
}

func NewStreamTracker() StreamTracker {
	return StreamTracker{
		streamSet: map[Source]streamTuple{},
	}
}
