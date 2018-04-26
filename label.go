// label.go - put a label on a packet.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of label, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package label

import (
	"errors"
	"io"
)

/* | len 1 byte | label | payload | */

// Read reads label from reader r.
func Read(r io.Reader) ([]byte, error) {
	lenSlice := make([]byte, 1)
	_, err := io.ReadFull(r, lenSlice)
	if err != nil {
		return nil, err
	}
	len := int(lenSlice[0])
	label := make([]byte, len)
	_, err = io.ReadFull(r, label)
	if err != nil {
		return nil, err
	}
	return label, nil
}

// Write writes label to writer w.
func Write(w io.Writer, label []byte) error {
	len := len(label)
	if len > 255 {
		return errors.New("label is too large")
	}
	_, err := w.Write([]byte{byte(uint8(len))})
	if err != nil {
		return err
	}
	// XXX: one write is probably not enough for large labels
	_, err = w.Write(label)
	if err != nil {
		return err
	}
	return nil
}
