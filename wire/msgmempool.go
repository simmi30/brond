// Copyright (c) 2013-2015 The brsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"io"
)

// MsgMemPool implements the Message interface and represents a brocoin mempool
// message.  It is used to request a list of transactions still in the active
// memory pool of a relay.
//
// This message has no payload and was not added until protocol versions
// starting with BIP0035Version.
type MsgMemPool struct{}

//BronDecode decodes r using the brocoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgMemPool)BronDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	if pver < BIP0035Version {
		str := fmt.Sprintf("mempool message invalid for protocol "+
			"version %d", pver)
		return messageError("MsgMemPool.BronDecode", str)
	}

	return nil
}

//BronEncode encodes the receiver to w using the brocoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgMemPool)BronEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	if pver < BIP0035Version {
		str := fmt.Sprintf("mempool message invalid for protocol "+
			"version %d", pver)
		return messageError("MsgMemPool.BronEncode", str)
	}

	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgMemPool) Command() string {
	return CmdMemPool
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgMemPool) MaxPayloadLength(pver uint32) uint32 {
	return 0
}

// NewMsgMemPool returns a new brocoin pong message that conforms to the Message
// interface.  See MsgPong for details.
func NewMsgMemPool() *MsgMemPool {
	return &MsgMemPool{}
}
