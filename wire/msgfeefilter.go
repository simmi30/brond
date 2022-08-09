// Copyright (c) 2016 The brsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"io"
)

// MsgFeeFilter implements the Message interface and represents a brocoin
// feefilter message.  It is used to request the receiving peer does not
// announce any transactions below the specified minimum fee rate.
//
// This message was not added until protocol versions starting with
// FeeFilterVersion.
type MsgFeeFilter struct {
	MinFee int64
}

//BronDecode decodes r using the brocoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgFeeFilter)BronDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	if pver < FeeFilterVersion {
		str := fmt.Sprintf("feefilter message invalid for protocol "+
			"version %d", pver)
		return messageError("MsgFeeFilter.BronDecode", str)
	}

	return readElement(r, &msg.MinFee)
}

//BronEncode encodes the receiver to w using the brocoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgFeeFilter)BronEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	if pver < FeeFilterVersion {
		str := fmt.Sprintf("feefilter message invalid for protocol "+
			"version %d", pver)
		return messageError("MsgFeeFilter.BronEncode", str)
	}

	return writeElement(w, msg.MinFee)
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgFeeFilter) Command() string {
	return CmdFeeFilter
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgFeeFilter) MaxPayloadLength(pver uint32) uint32 {
	return 8
}

// NewMsgFeeFilter returns a new brocoin feefilter message that conforms to
// the Message interface.  See MsgFeeFilter for details.
func NewMsgFeeFilter(minfee int64) *MsgFeeFilter {
	return &MsgFeeFilter{
		MinFee: minfee,
	}
}
