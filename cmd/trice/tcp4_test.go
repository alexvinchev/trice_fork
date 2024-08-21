// Use of this source code is governed by a license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"net"
	"testing"

	"github.com/rokath/trice/internal/args"
	"github.com/rokath/trice/pkg/tst"
	"github.com/spf13/afero"
	"github.com/tj/assert"
)

func handleConnection(conn net.Conn) {
	// Close the connection when we're done
	defer conn.Close()

	// No incoming data to read.
	//     buf := make([]byte, 1024)
	//     n, err := conn.Read(buf)

	// That is the used source code line:
	// TRice( iD(16201), "w: Hello! 👋🙂 \a\n" ); // with sound!

	// That is copied from the terminal
	//  ms@DESKTOP-7POEGPB MINGW64 ~/repos/trice/examples/G0B1_inst (master)
	//  $ trice log -p com6 -pw MySecret -pf COBS -debug -showID "deb:%5d"
	//  COBS: 09 92 19 06 45 0b 10 56 3a 00
	//  ->TRICE: 92 19 06 45 0b 10 56 3a
	//  -> DEC:  49 ff 00 00 00 00 c0 00
	//  Aug 21 13:15:54.639047  com6:                main.c    75    0,000_000 16201  Hello! 👋🙂

	// That is the 0-terminated transfer package
	conn.Write([]byte{0x09, 0x92, 0x19, 0x06, 0x45, 0x0b, 0x10, 0x56, 0x3a, 0x00})
}

// TestTCP4Reception sets up a trivial trice log server, which ends on request a single trice package.
// This is a COBS encoded and encrypted package, derived from real data.
// When the Trice tool is started, it read from the trice log server over TCP4 and decodes the received package.
func TestTCP4Reception(t *testing.T) {

	fSys := &afero.Afero{Fs: afero.NewMemMapFs()}
	defer setupTest(t, fSys)()

	til := `{
	"16201": {
		"Type": "TRice",
		"Strg": "w: Hello! 👋🙂\\n"
	}
}`

	// create a minimalistic til.json
	assert.Nil(t, fSys.WriteFile("til.json", []byte(til), 0777))

	// Listen for incoming connections (start server)
	portNumber := "54321"
	ln, err := net.Listen("tcp", ":"+portNumber)
	assert.Nil(t, err)

	// Accept incoming connections and handle them
	srv := func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}

			// Handle the connection in a new goroutine
			go handleConnection(conn)
		}
	}

	go srv() // server active now

	// We use "-port TCP4BUFFER" just for the test, to force the Trice tool to shutdown after receiving a package.
	// In real life, the user will enter "-port TCP4" instead. To keep things simple we switch off all unnecessary information.
	input := []string{"trice", "log", "-port", "TCP4BUFFER", "-args", "localhost:" + portNumber, "-pw", "MySecret", "-pf", "cobs", "-li", "off", "-hs", "off", "-color", "none", "-prefix", "off", "-ts", "off"}
	expect := `Hello! 👋🙂
`

	var out bytes.Buffer
	err = args.Handler(&out, fSys, input)
	if err != nil {
		fmt.Fprint(&out, err)
	}
	assert.Nil(t, err)

	act := out.String()
	tst.EqualLines(t, expect, act)

}
