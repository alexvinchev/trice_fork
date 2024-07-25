// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package id_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/rokath/trice/internal/args"
	"github.com/rokath/trice/internal/id"
	"github.com/spf13/afero"
	"github.com/tj/assert"
)

func TestInsertIDsAndJSONUpward(t *testing.T) {

	fSys := &afero.Afero{Fs: afero.NewMemMapFs()}
	defer id.SetupTest(t, fSys)()

	// create src files
	src0 := `
	TRice( "Hi!" );
	TRice( "Hi!" );
	`
	fn0 := t.Name() + "file0.c"
	assert.Nil(t, fSys.WriteFile(fn0, []byte(src0), 0777))

	// action
	var b bytes.Buffer
	assert.Nil(t, args.Handler(io.Writer(&b), fSys, []string{"TRICE", "insert", "-src", "./", "-til", id.FnJSON, "-li", id.LIFnJSON, "-IDMin", "10", "-IDMax", "20", "-IDMethod", "upward"}))

	// check source files
	expSrc0 := `
	TRice( iD(10), "Hi!" );
	TRice( iD(11), "Hi!" );
	`
	actSrc0, e := fSys.ReadFile(fn0)
	assert.Nil(t, e)
	assert.Equal(t, expSrc0, string(actSrc0))

	// check til.json
	expJSON := `{
	"10": {
		"Type": "TRice",
		"Strg": "Hi!"
	},
	"11": {
		"Type": "TRice",
		"Strg": "Hi!"
	}
}`
	actJSON, e := fSys.ReadFile(id.FnJSON)
	assert.Nil(t, e)
	assert.Equal(t, expJSON, string(actJSON))

	// check location information
	expLI := `{
	"10": {
		"File": "` + fn0 + `",
		"Line": 2
	},
	"11": {
		"File": "` + fn0 + `",
		"Line": 3
	}
}`
	actLI, e := fSys.ReadFile(id.LIFnJSON)
	assert.Nil(t, e)
	assert.Equal(t, expLI, string(actLI))

}
