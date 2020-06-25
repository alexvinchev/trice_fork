// Copyright 2020 basti@blackoutcloud.de
//                Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

// Package trice provides subcommands ScLog and ScReceive
// It provides optional decryption and activates receiver package,
// which in turn calls emit.Trice on every received trice.
package trice

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/rokath/trice/internal/disp"
	"github.com/rokath/trice/internal/id"
	"github.com/rokath/trice/internal/receiver"
	"golang.org/x/crypto/xtea"
)

// local config values
var (
	// Password is the key one needs to derypt trice logs if enncrypted
	Password string

	// ShowPassword, if set, allows to see the encryption passphrase
	ShowPassword bool
)

/*
// ScLog is the subcommand log and connects to COM port and displays traces
func ScLog() error {
	cage.Enable()
	defer cage.Disable()

	return DoReceive()
}*/

// Connect starts a display server sy if sv is not empty, otherwise it assumes a running display server.
//
// It connects then to the running display server.
func Connect(sv string) error {
	if "" != sv {
		disp.StartServer(sv)
	}

	err := disp.Connect()
	disp.Out = disp.RemoteOut // re-direct output
	if nil != err {
		return err
	}

	disp.PtrRPC.Call("Server.Out", []string{""}, nil)
	disp.PtrRPC.Call("Server.Out", []string{""}, nil)
	disp.PtrRPC.Call("Server.Out", []string{""}, nil)
	disp.PtrRPC.Call("Server.Out", []string{"att:new connection from ", "read:" + receiver.Source, "..."}, nil)
	disp.PtrRPC.Call("Server.Out", []string{""}, nil)
	disp.PtrRPC.Call("Server.Out", []string{""}, nil)
	return nil
}

/*
// ScReceive is the subcommand remoteDisplay and acts as client connecting to the displayServer
// sv is the executable name to be started as remote display server (typically arg[0] == trice)
func ScReceive(sv string) error {
	err := NewConnection(sv)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cmd.KeyboardInput()
	DoReceive() // does not return
	return nil
}
*/

// DoReceive connects to COM port and displays traces
func SetUp() error {
	if "none" != id.FnJSON {
		// setup ip list
		err := id.List.Read(id.FnJSON)
		if nil != err {
			//fmt.Println("ID list " + path.Base(id.FnJSON) + " not found, exit")
			id.ListNotFoundMsg(id.FnJSON)
			return errors.New("file not found")
		}
		go id.List.FileWatcher()
	}

	var err error
	receiver.Cipher, receiver.Crypto, err = createCipher()
	if nil != err {
		return err
	}
	if true == id.Verbose {
		fmt.Println("id list file", id.FnJSON, "with", len(id.List), "items", "on device", receiver.Source)
	}
	return nil
}

/* DoReceive connects to COM port and displays traces
func DoReceive() error {

	SetUp()
	do()

	return nil
}

*/
// createCipher prepares decryption, with password "none" the encryption flag is set false, otherwise true
func createCipher() (*xtea.Cipher, bool, error) {
	h := sha1.New() // https://gobyexample.com/sha1-hashes
	h.Write([]byte(Password))
	key := h.Sum(nil)
	key = key[:16] // only first 16 bytes needed as key

	c, err := xtea.NewCipher(key)
	if err != nil {
		return nil, false, errors.New("NewCipher returned error")
	}
	var e bool
	if "none" != Password {
		e = true
		if true == ShowPassword {
			fmt.Printf("% 20x is XTEA encryption key\n", key)
		}
	} else if true == ShowPassword {
		fmt.Printf("no encryption\n")
	}
	return c, e, nil
}
