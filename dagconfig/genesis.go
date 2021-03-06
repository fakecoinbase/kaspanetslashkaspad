// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"time"

	"github.com/kaspanet/kaspad/util/daghash"
	"github.com/kaspanet/kaspad/util/subnetworkid"
	"github.com/kaspanet/kaspad/wire"
)

var genesisTxOuts = []*wire.TxOut{}

var genesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x17,                                           // Varint
	0xa9, 0x14, 0xda, 0x17, 0x45, 0xe9, 0xb5, 0x49, // OP-TRUE p2sh
	0xbd, 0x0b, 0xfa, 0x1a, 0x56, 0x99, 0x71, 0xc7,
	0x7e, 0xba, 0x30, 0xcd, 0x5a, 0x4b, 0x87,
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = wire.NewSubnetworkMsgTx(1, []*wire.TxIn{}, genesisTxOuts, subnetworkid.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = daghash.Hash{
	0xb3, 0x5d, 0x34, 0x3c, 0xf6, 0xbb, 0xd5, 0xaf,
	0x40, 0x4b, 0xff, 0x3f, 0x83, 0x27, 0x71, 0x1e,
	0xe1, 0x83, 0xf6, 0x41, 0x32, 0x8c, 0xba, 0xe6,
	0xd3, 0xba, 0x13, 0xef, 0x7b, 0x7e, 0x61, 0x65,
}

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = daghash.Hash{
	0xca, 0x85, 0x56, 0x27, 0xc7, 0x6a, 0xb5, 0x7a,
	0x26, 0x1d, 0x63, 0x62, 0x1e, 0x57, 0x21, 0xf0,
	0x5e, 0x60, 0x1f, 0xee, 0x1d, 0x4d, 0xaa, 0x53,
	0x72, 0xe1, 0x16, 0xda, 0x4b, 0xb3, 0xd8, 0x0e,
}

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:              0x10000000,
		ParentHashes:         []*daghash.Hash{},
		HashMerkleRoot:       &genesisMerkleRoot,
		AcceptedIDMerkleRoot: &daghash.Hash{},
		UTXOCommitment:       &daghash.ZeroHash,
		Timestamp:            time.Unix(0x5edf4ce0, 0),
		Bits:                 0x207fffff,
		Nonce:                0,
	},
	Transactions: []*wire.MsgTx{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*wire.TxOut{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x17,                                           // Varint
	0xa9, 0x14, 0xda, 0x17, 0x45, 0xe9, 0xb5, 0x49, // OP-TRUE p2sh
	0xbd, 0x0b, 0xfa, 0x1a, 0x56, 0x99, 0x71, 0xc7,
	0x7e, 0xba, 0x30, 0xcd, 0x5a, 0x4b, 0x87,
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x74, // kaspa-devnet
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = wire.NewSubnetworkMsgTx(1, []*wire.TxIn{}, devnetGenesisTxOuts, subnetworkid.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = daghash.Hash{
	0x50, 0x92, 0xd1, 0x1f, 0xaa, 0xba, 0xd3, 0x58,
	0xa8, 0x22, 0xd7, 0xec, 0x8e, 0xe3, 0xf4, 0x26,
	0x17, 0x18, 0x74, 0xd7, 0x87, 0x05, 0x9d, 0xed,
	0x33, 0xcd, 0xe1, 0x26, 0x1a, 0x69, 0x00, 0x00,
}

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = daghash.Hash{
	0x68, 0x60, 0xe7, 0x77, 0x47, 0x74, 0x7f, 0xd5,
	0x55, 0x58, 0x8a, 0xb5, 0xc2, 0x29, 0x0c, 0xa6,
	0x65, 0x44, 0xb4, 0x4f, 0xfa, 0x31, 0x7a, 0xfa,
	0x55, 0xe0, 0xcf, 0xac, 0x9c, 0x86, 0x30, 0x2a,
}

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:              0x10000000,
		ParentHashes:         []*daghash.Hash{},
		HashMerkleRoot:       &devnetGenesisMerkleRoot,
		AcceptedIDMerkleRoot: &daghash.Hash{},
		UTXOCommitment:       &daghash.ZeroHash,
		Timestamp:            time.Unix(0x5edf4ce0, 0),
		Bits:                 0x1e7fffff,
		Nonce:                0xb3ed,
	},
	Transactions: []*wire.MsgTx{devnetGenesisCoinbaseTx},
}

var regtestGenesisTxOuts = []*wire.TxOut{}

var regtestGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x17,                                           // Varint
	0xa9, 0x14, 0xda, 0x17, 0x45, 0xe9, 0xb5, 0x49, // OP-TRUE p2sh
	0xbd, 0x0b, 0xfa, 0x1a, 0x56, 0x99, 0x71, 0xc7,
	0x7e, 0xba, 0x30, 0xcd, 0x5a, 0x4b, 0x87,
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x72, 0x65, 0x67, 0x74, 0x65, 0x73, 0x74, // kaspa-regtest
}

// regtestGenesisCoinbaseTx is the coinbase transaction for
// the genesis blocks for the regtest network.
var regtestGenesisCoinbaseTx = wire.NewSubnetworkMsgTx(1, []*wire.TxIn{}, regtestGenesisTxOuts, subnetworkid.SubnetworkIDCoinbase, 0, regtestGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var regtestGenesisHash = daghash.Hash{
	0xf8, 0x1d, 0xe9, 0x86, 0xa5, 0x60, 0xe0, 0x34,
	0x0f, 0x02, 0xaa, 0x8d, 0xea, 0x6f, 0x1f, 0xc6,
	0x2a, 0xb4, 0x77, 0xbd, 0xca, 0xed, 0xad, 0x3c,
	0x99, 0xe6, 0x98, 0x7c, 0x7b, 0x5e, 0x00, 0x00,
}

// regtestGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the regtest.
var regtestGenesisMerkleRoot = daghash.Hash{
	0x1e, 0x08, 0xae, 0x1f, 0x43, 0xf5, 0xfc, 0x24,
	0xe6, 0xec, 0x54, 0x5b, 0xf7, 0x52, 0x99, 0xe4,
	0xcc, 0x4c, 0xa0, 0x79, 0x41, 0xfc, 0xbe, 0x76,
	0x72, 0x4c, 0x7e, 0xd8, 0xa3, 0x43, 0x65, 0x94,
}

// regtestGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var regtestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:              0x10000000,
		ParentHashes:         []*daghash.Hash{},
		HashMerkleRoot:       &regtestGenesisMerkleRoot,
		AcceptedIDMerkleRoot: &daghash.Hash{},
		UTXOCommitment:       &daghash.ZeroHash,
		Timestamp:            time.Unix(0x5edf4ce0, 0),
		Bits:                 0x1e7fffff,
		Nonce:                0x4a78,
	},
	Transactions: []*wire.MsgTx{regtestGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*wire.TxOut{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x17,                                           // Varint
	0xa9, 0x14, 0xda, 0x17, 0x45, 0xe9, 0xb5, 0x49, // OP-TRUE p2sh
	0xbd, 0x0b, 0xfa, 0x1a, 0x56, 0x99, 0x71, 0xc7,
	0x7e, 0xba, 0x30, 0xcd, 0x5a, 0x4b, 0x87,
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x73, 0x69, 0x6d, 0x6e, 0x65, 0x74, // kaspa-simnet
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = wire.NewSubnetworkMsgTx(1, []*wire.TxIn{}, simnetGenesisTxOuts, subnetworkid.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = daghash.Hash{
	0x34, 0x43, 0xed, 0xdc, 0xab, 0x0c, 0x39, 0x53,
	0xa2, 0xc5, 0x6d, 0x12, 0x4b, 0xc2, 0x41, 0x1c,
	0x1a, 0x05, 0x24, 0xb4, 0xff, 0xeb, 0xe8, 0xbd,
	0xee, 0x6e, 0x9a, 0x77, 0xc7, 0xbb, 0x70, 0x7d,
}

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var simnetGenesisMerkleRoot = daghash.Hash{
	0x47, 0x52, 0xc7, 0x23, 0x70, 0x4d, 0x89, 0x17,
	0xbd, 0x44, 0x26, 0xfa, 0x82, 0x7e, 0x1b, 0xa9,
	0xc6, 0x46, 0x1a, 0x37, 0x5a, 0x73, 0x88, 0x09,
	0xe8, 0x17, 0xff, 0xb1, 0xdb, 0x1a, 0xb3, 0x3f,
}

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:              0x10000000,
		ParentHashes:         []*daghash.Hash{},
		HashMerkleRoot:       &simnetGenesisMerkleRoot,
		AcceptedIDMerkleRoot: &daghash.Hash{},
		UTXOCommitment:       &daghash.ZeroHash,
		Timestamp:            time.Unix(0x5ede5261, 0),
		Bits:                 0x207fffff,
		Nonce:                0x2,
	},
	Transactions: []*wire.MsgTx{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*wire.TxOut{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x01,                                                                         // Varint
	0x00,                                                                         // OP-FALSE
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x6e, 0x65, 0x74, // kaspa-testnet
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = wire.NewSubnetworkMsgTx(1, []*wire.TxIn{}, testnetGenesisTxOuts, subnetworkid.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = daghash.Hash{
	0xFC, 0x21, 0x64, 0x1A, 0xB5, 0x59, 0x61, 0x8E,
	0xF3, 0x9A, 0x95, 0xF1, 0xDA, 0x07, 0x79, 0xBD,
	0x11, 0x2F, 0x90, 0xFC, 0x8B, 0x33, 0x14, 0x8A,
	0x90, 0x6B, 0x76, 0x08, 0x4B, 0x52, 0x00, 0x00,
}

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = daghash.Hash{
	0xA0, 0xA1, 0x3D, 0xFD, 0x86, 0x41, 0x35, 0xC8,
	0xBD, 0xBB, 0xE6, 0x37, 0x35, 0xBB, 0x4C, 0x51,
	0x11, 0x7B, 0x26, 0x90, 0x15, 0x64, 0x0F, 0x42,
	0x6D, 0x2B, 0x6F, 0x37, 0x4D, 0xC1, 0xA9, 0x72,
}

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:              0x10000000,
		ParentHashes:         []*daghash.Hash{},
		HashMerkleRoot:       &testnetGenesisMerkleRoot,
		AcceptedIDMerkleRoot: &daghash.ZeroHash,
		UTXOCommitment:       &daghash.ZeroHash,
		Timestamp:            time.Unix(0x5efc2128, 0),
		Bits:                 0x1e7fffff,
		Nonce:                0x1124,
	},
	Transactions: []*wire.MsgTx{testnetGenesisCoinbaseTx},
}
