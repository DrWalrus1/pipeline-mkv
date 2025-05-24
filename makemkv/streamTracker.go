package makemkv

import "io"

type Source = string

/*
A situation can arrise when we need to reattach onto the output of a command in makemkv.
Hence, we create an instance of
*/
type StreamTracker struct {
	streamSet map[Source]*io.Reader
}

func (st *StreamTracker) AddStream(source Source, reader *io.Reader) error {
	st.streamSet[source] = reader
	return nil
}

func (st *StreamTracker) GetStream(source Source) (*io.Reader, bool) {
	reader, ok := st.streamSet[source]
	return reader, ok
}

func (st *StreamTracker) RemoveStream(source Source) {
	delete(st.streamSet, source)
}

func NewStreamTracker() StreamTracker {
	return StreamTracker{
		streamSet: map[Source]*io.Reader{},
	}
}
