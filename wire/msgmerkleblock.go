// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"io"

	"github.com/kaspanet/kaspad/util/daghash"
)

// maxFlagsPerMerkleBlock is the maximum number of flag bytes that could
// possibly fit into a merkle block. Since each transaction is represented by
// a single bit, this is the max number of transactions per block divided by
// 8 bits per byte. Then an extra one to cover partials.
const maxFlagsPerMerkleBlock = maxTxPerBlock / 8

// MsgMerkleBlock implements the Message interface and represents a kaspa
// merkleblock message which is used to reset a Bloom filter.
//
// This message was not added until protocol version BIP0037Version.
type MsgMerkleBlock struct {
	Header       BlockHeader
	Transactions uint32
	Hashes       []*daghash.Hash
	Flags        []byte
}

// AddTxHash adds a new transaction hash to the message.
func (msg *MsgMerkleBlock) AddTxHash(hash *daghash.Hash) error {
	if len(msg.Hashes)+1 > maxTxPerBlock {
		str := fmt.Sprintf("too many tx hashes for message [max %d]",
			maxTxPerBlock)
		return messageError("MsgMerkleBlock.AddTxHash", str)
	}

	msg.Hashes = append(msg.Hashes, hash)
	return nil
}

// KaspaDecode decodes r using the kaspa protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgMerkleBlock) KaspaDecode(r io.Reader, pver uint32) error {
	err := readBlockHeader(r, pver, &msg.Header)
	if err != nil {
		return err
	}

	err = ReadElement(r, &msg.Transactions)
	if err != nil {
		return err
	}

	// Read num block locator hashes and limit to max.
	count, err := ReadVarInt(r)
	if err != nil {
		return err
	}
	if count > maxTxPerBlock {
		str := fmt.Sprintf("too many transaction hashes for message "+
			"[count %d, max %d]", count, maxTxPerBlock)
		return messageError("MsgMerkleBlock.KaspaDecode", str)
	}

	// Create a contiguous slice of hashes to deserialize into in order to
	// reduce the number of allocations.
	hashes := make([]*daghash.Hash, count)
	msg.Hashes = make([]*daghash.Hash, 0, count)
	for i := uint64(0); i < count; i++ {
		hash := &daghash.Hash{}
		err := ReadElement(r, hash)
		if err != nil {
			return err
		}
		msg.AddTxHash(hash)
		hashes[i] = hash
	}

	msg.Flags, err = ReadVarBytes(r, pver, maxFlagsPerMerkleBlock,
		"merkle block flags size")
	return err
}

// KaspaEncode encodes the receiver to w using the kaspa protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgMerkleBlock) KaspaEncode(w io.Writer, pver uint32) error {
	// Read num transaction hashes and limit to max.
	numHashes := len(msg.Hashes)
	if numHashes > maxTxPerBlock {
		str := fmt.Sprintf("too many transaction hashes for message "+
			"[count %d, max %d]", numHashes, maxTxPerBlock)
		return messageError("MsgMerkleBlock.KaspaDecode", str)
	}
	numFlagBytes := len(msg.Flags)
	if numFlagBytes > maxFlagsPerMerkleBlock {
		str := fmt.Sprintf("too many flag bytes for message [count %d, "+
			"max %d]", numFlagBytes, maxFlagsPerMerkleBlock)
		return messageError("MsgMerkleBlock.KaspaDecode", str)
	}

	err := writeBlockHeader(w, pver, &msg.Header)
	if err != nil {
		return err
	}

	err = WriteElement(w, msg.Transactions)
	if err != nil {
		return err
	}

	err = WriteVarInt(w, uint64(numHashes))
	if err != nil {
		return err
	}
	for _, hash := range msg.Hashes {
		err = WriteElement(w, hash)
		if err != nil {
			return err
		}
	}

	return WriteVarBytes(w, pver, msg.Flags)
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgMerkleBlock) Command() string {
	return CmdMerkleBlock
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver. This is part of the Message interface implementation.
func (msg *MsgMerkleBlock) MaxPayloadLength(pver uint32) uint32 {
	return MaxMessagePayload
}

// NewMsgMerkleBlock returns a new kaspa merkleblock message that conforms to
// the Message interface. See MsgMerkleBlock for details.
func NewMsgMerkleBlock(bh *BlockHeader) *MsgMerkleBlock {
	return &MsgMerkleBlock{
		Header:       *bh,
		Transactions: 0,
		Hashes:       make([]*daghash.Hash, 0),
		Flags:        make([]byte, 0),
	}
}
