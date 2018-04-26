// label_test.go
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of label, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package label

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/matryer/is"
)

func TestFull(t *testing.T) {
	is := is.New(t)
	payload := []byte("hello")
	l := []byte("localhost")
	buf := &bytes.Buffer{}
	err := Write(buf, l)
	is.NoErr(err)
	_, err = buf.Write(payload)
	is.NoErr(err)

	l2, err := Read(buf)
	is.NoErr(err)
	payload2, err := ioutil.ReadAll(buf)
	is.NoErr(err)
	is.Equal(l2, l)
	is.Equal(payload2, payload)
}
