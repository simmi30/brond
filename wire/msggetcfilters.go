// Copyright (c) 2017 The brsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"io"

	"github.com/brsuite/brond/chaincfg/chainhash"
)

// MaxGetCFiltersReqRange the maximum number of filters that may be requested in
// a getcfheaders message.
const MaxGetCFiltersReqRange = 1000

// MsgGetCFilters implements the Message interface and represents a brocoin
// getcfilters message. It is used to request committed filters for a range of
// blocks.
type MsgGetCFilters struct {
	FilterType  FilterType
	StartHeight uint32
	StopHash    chainhash.Hash
}

//BronDecode decodes r using the brocoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgGetCFilters)BronDecode(r io.Reader, pver uint32, _ MessageEncoding) error {
	err := readElement(r, &msg.FilterType)
	if err != nil {
		return err
	}

	err = readElement(r, &msg.StartHeight)
	if err != nil {
		return err
	}

	return readElement(r, &msg.StopHash)
}

//BronEncode encodes the receiver to w using the brocoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgGetCFilters)BronEncode(w io.Writer, pver uint32, _ MessageEncoding) error {
	err := writeElement(w, msg.FilterType)
	if err != nil {
		return err
	}

	err = writeElement(w, &msg.StartHeight)
	if err != nil {
		return err
	}

	return writeElement(w, &msg.StopHash)
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgGetCFilters) Command() string {
	return CmdGetCFilters
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgGetCFilters) MaxPayloadLength(pver uint32) uint32 {
	// Filter type + uint32 + block hash
	return 1 + 4 + chainhash.HashSize
}

// NewMsgGetCFilters returns a new brocoin getcfilters message that conforms to
// the Message interface using the passed parameters and defaults for the
// remaining fields.
func NewMsgGetCFilters(filterType FilterType, startHeight uint32,
	stopHash *chainhash.Hash) *MsgGetCFilters {
	return &MsgGetCFilters{
		FilterType:  filterType,
		StartHeight: startHeight,
		StopHash:    *stopHash,
	}
}
