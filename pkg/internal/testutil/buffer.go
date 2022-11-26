package testutil

import "bytes"

func ConvertBuffer(b *bytes.Buffer) *Buffer {
	return &Buffer{
		Buffer: *b,
	}
}

func NewBufferString(v string) *Buffer {
	return &Buffer{
		Buffer: *bytes.NewBufferString(v),
	}
}

type Buffer struct {
	bytes.Buffer
}

func (*Buffer) Close() error {
	return nil
}
