// Copyright (c) 2014-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"io"
)

const (
	// MaxFilterAddDataSize is the maximum byte size of a data
	// element to add to the Bloom filter. It is equal to the
	// maximum element size of a script.
	MaxFilterAddDataSize = 520
)

// MsgFilterAdd implements the Message interface and represents a kaspa
// filteradd message. It is used to add a data element to an existing Bloom
// filter.
//
// This message was not added until protocol version BIP0037Version.
type MsgFilterAdd struct {
	Data []byte
}

// KaspaDecode decodes r using the kaspa protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgFilterAdd) KaspaDecode(r io.Reader, pver uint32) error {
	var err error
	msg.Data, err = ReadVarBytes(r, pver, MaxFilterAddDataSize,
		"filteradd data")
	return err
}

// KaspaEncode encodes the receiver to w using the kaspa protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgFilterAdd) KaspaEncode(w io.Writer, pver uint32) error {
	size := len(msg.Data)
	if size > MaxFilterAddDataSize {
		str := fmt.Sprintf("filteradd size too large for message "+
			"[size %d, max %d]", size, MaxFilterAddDataSize)
		return messageError("MsgFilterAdd.KaspaEncode", str)
	}

	return WriteVarBytes(w, pver, msg.Data)
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgFilterAdd) Command() string {
	return CmdFilterAdd
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver. This is part of the Message interface implementation.
func (msg *MsgFilterAdd) MaxPayloadLength(pver uint32) uint32 {
	return uint32(VarIntSerializeSize(MaxFilterAddDataSize)) +
		MaxFilterAddDataSize
}

// NewMsgFilterAdd returns a new kaspa filteradd message that conforms to the
// Message interface. See MsgFilterAdd for details.
func NewMsgFilterAdd(data []byte) *MsgFilterAdd {
	return &MsgFilterAdd{
		Data: data,
	}
}
