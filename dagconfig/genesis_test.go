// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"bytes"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// TestGenesisBlock tests the genesis block of the main network for validity by
// checking the encoded bytes and hashes.
func TestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := MainnetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), genesisBlockBytes) {
		t.Fatalf("TestGenesisBlock: Genesis block does not appear valid - "+
			"got %v, want %v", spew.Sdump(buf.Bytes()),
			spew.Sdump(genesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := MainnetParams.GenesisBlock.BlockHash()
	if !MainnetParams.GenesisHash.IsEqual(hash) {
		t.Fatalf("TestGenesisBlock: Genesis block hash does not "+
			"appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(MainnetParams.GenesisHash))
	}
}

// TestRegtestGenesisBlock tests the genesis block of the regression test
// network for validity by checking the encoded bytes and hashes.
func TestRegtestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := RegressionNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestRegtestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), regtestGenesisBlockBytes) {
		t.Fatalf("TestRegtestGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(regtestGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := RegressionNetParams.GenesisBlock.BlockHash()
	if !RegressionNetParams.GenesisHash.IsEqual(hash) {
		t.Fatalf("TestRegtestGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(RegressionNetParams.GenesisHash))
	}
}

// TestTestnetGenesisBlock tests the genesis block of the test network for
// validity by checking the encoded bytes and hashes.
func TestTestnetGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := TestnetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestTestnetGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), testnetGenesisBlockBytes) {
		t.Fatalf("TestTestnetGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(testnetGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := TestnetParams.GenesisBlock.BlockHash()
	if !TestnetParams.GenesisHash.IsEqual(hash) {
		t.Fatalf("TestTestnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(TestnetParams.GenesisHash))
	}
}

// TestSimnetGenesisBlock tests the genesis block of the simulation test network
// for validity by checking the encoded bytes and hashes.
func TestSimnetGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := SimnetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestSimnetGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), simnetGenesisBlockBytes) {
		t.Fatalf("TestSimnetGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(simnetGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := SimnetParams.GenesisBlock.BlockHash()
	if !SimnetParams.GenesisHash.IsEqual(hash) {
		t.Fatalf("TestSimnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(SimnetParams.GenesisHash))
	}
}

// TestDevnetGenesisBlock tests the genesis block of the development network
// for validity by checking the encoded bytes and hashes.
func TestDevnetGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := DevnetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestDevnetGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), devnetGenesisBlockBytes) {
		t.Fatalf("TestDevnetGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(simnetGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := DevnetParams.GenesisBlock.BlockHash()
	if !DevnetParams.GenesisHash.IsEqual(hash) {
		t.Fatalf("TestDevnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(DevnetParams.GenesisHash))
	}
}

// genesisBlockBytes are the wire encoded bytes for the genesis block of the
// main network as of protocol version 1.
var genesisBlockBytes = []byte{
	0x00, 0x00, 0x00, 0x10, 0x00, 0xca, 0x85, 0x56, 0x27, 0xc7, 0x6a, 0xb5, 0x7a, 0x26, 0x1d, 0x63,
	0x62, 0x1e, 0x57, 0x21, 0xf0, 0x5e, 0x60, 0x1f, 0xee, 0x1d, 0x4d, 0xaa, 0x53, 0x72, 0xe1, 0x16,
	0xda, 0x4b, 0xb3, 0xd8, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x4c, 0xdf, 0x5e, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xc4, 0x41, 0xe6, 0x78, 0x1d, 0xf7, 0xb3, 0x39, 0x66, 0x4d, 0x1a, 0x03,
	0x97, 0x63, 0xc7, 0x2c, 0xfc, 0x70, 0xd7, 0x75, 0xb6, 0xd9, 0xfc, 0x1a, 0x96, 0xf0, 0xac, 0x07,
	0xef, 0xfa, 0x26, 0x38, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x17, 0xa9, 0x14,
	0xda, 0x17, 0x45, 0xe9, 0xb5, 0x49, 0xbd, 0x0b, 0xfa, 0x1a, 0x56, 0x99, 0x71, 0xc7, 0x7e, 0xba,
	0x30, 0xcd, 0x5a, 0x4b, 0x87,
}

// regtestGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the regression test network as of protocol version 1.
var regtestGenesisBlockBytes = []byte{
	0x00, 0x00, 0x00, 0x10, 0x00, 0x1e, 0x08, 0xae, 0x1f, 0x43, 0xf5, 0xfc, 0x24, 0xe6, 0xec, 0x54,
	0x5b, 0xf7, 0x52, 0x99, 0xe4, 0xcc, 0x4c, 0xa0, 0x79, 0x41, 0xfc, 0xbe, 0x76, 0x72, 0x4c, 0x7e,
	0xd8, 0xa3, 0x43, 0x65, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x4c, 0xdf, 0x5e, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x1e, 0x78, 0x4a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xd4, 0xc4, 0x87, 0x77, 0xf2, 0xe7, 0x5d, 0xf7, 0xff, 0x2d, 0xbb, 0xb6,
	0x2a, 0x73, 0x1f, 0x54, 0x36, 0x33, 0xa7, 0x99, 0xad, 0xb1, 0x09, 0x65, 0xc0, 0xf0, 0xf4, 0x53,
	0xba, 0xfb, 0x88, 0xae, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x17, 0xa9, 0x14,
	0xda, 0x17, 0x45, 0xe9, 0xb5, 0x49, 0xbd, 0x0b, 0xfa, 0x1a, 0x56, 0x99, 0x71, 0xc7, 0x7e, 0xba,
	0x30, 0xcd, 0x5a, 0x4b, 0x87, 0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x72, 0x65, 0x67, 0x74, 0x65,
	0x73, 0x74,
}

// testnetGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the test network as of protocol version 1.
var testnetGenesisBlockBytes = []byte{
	0x00, 0x00, 0x00, 0x10, 0x00, 0xa0, 0xa1, 0x3d, 0xfd, 0x86, 0x41, 0x35, 0xc8, 0xbd, 0xbb, 0xe6,
	0x37, 0x35, 0xbb, 0x4c, 0x51, 0x11, 0x7b, 0x26, 0x90, 0x15, 0x64, 0x0f, 0x42, 0x6d, 0x2b, 0x6f,
	0x37, 0x4d, 0xc1, 0xa9, 0x72, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x28, 0x21, 0xfc, 0x5e, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x1e, 0x24, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xf5, 0x41, 0x4c, 0xf4, 0xa8, 0xa2, 0x8c, 0x47, 0x9d, 0xb5, 0x75, 0x5e,
	0x0f, 0x38, 0xd3, 0x27, 0x82, 0xc6, 0xd1, 0x89, 0xc1, 0x60, 0x49, 0xd9, 0x99, 0xc6, 0x2e, 0xbf,
	0x4b, 0x5a, 0x3a, 0xcf, 0x17, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x6b,
	0x61, 0x73, 0x70, 0x61, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x6e, 0x65, 0x74,
}

// simnetGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the simulation test network as of protocol version 1.
var simnetGenesisBlockBytes = []byte{
	0x00, 0x00, 0x00, 0x10, 0x00, 0x47, 0x52, 0xc7, 0x23, 0x70, 0x4d, 0x89, 0x17, 0xbd, 0x44, 0x26,
	0xfa, 0x82, 0x7e, 0x1b, 0xa9, 0xc6, 0x46, 0x1a, 0x37, 0x5a, 0x73, 0x88, 0x09, 0xe8, 0x17, 0xff,
	0xb1, 0xdb, 0x1a, 0xb3, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x61, 0x52, 0xde, 0x5e, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x20, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xd9, 0x39, 0x5f, 0x40, 0x2a, 0x5e, 0x24, 0x09, 0x1b, 0x9a, 0x4b, 0xdf,
	0x7f, 0x0c, 0x03, 0x7f, 0xf1, 0xd2, 0x48, 0x8c, 0x26, 0xb0, 0xa3, 0x74, 0x60, 0xd9, 0x48, 0x18,
	0x2b, 0x33, 0x22, 0x64, 0x2c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x17, 0xa9, 0x14,
	0xda, 0x17, 0x45, 0xe9, 0xb5, 0x49, 0xbd, 0x0b, 0xfa, 0x1a, 0x56, 0x99, 0x71, 0xc7, 0x7e, 0xba,
	0x30, 0xcd, 0x5a, 0x4b, 0x87, 0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x73, 0x69, 0x6d, 0x6e, 0x65,
	0x74,
}

// devnetGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the development network as of protocol version 1.
var devnetGenesisBlockBytes = []byte{
	0x00, 0x00, 0x00, 0x10, 0x00, 0x68, 0x60, 0xe7, 0x77, 0x47, 0x74, 0x7f, 0xd5, 0x55, 0x58, 0x8a,
	0xb5, 0xc2, 0x29, 0x0c, 0xa6, 0x65, 0x44, 0xb4, 0x4f, 0xfa, 0x31, 0x7a, 0xfa, 0x55, 0xe0, 0xcf,
	0xac, 0x9c, 0x86, 0x30, 0x2a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x4c, 0xdf, 0x5e, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x1e, 0xed, 0xb3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x1d, 0x1c, 0x05, 0x21, 0x10, 0x45, 0x61, 0xed, 0xc6, 0x0b, 0xdc, 0x85,
	0xc0, 0x0a, 0x70, 0x2b, 0x15, 0xd5, 0x3c, 0x07, 0xb0, 0x54, 0x4f, 0x5b, 0x1a, 0x04, 0xcd, 0x49,
	0xf1, 0x7b, 0xd6, 0x27, 0x2c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x17, 0xa9, 0x14,
	0xda, 0x17, 0x45, 0xe9, 0xb5, 0x49, 0xbd, 0x0b, 0xfa, 0x1a, 0x56, 0x99, 0x71, 0xc7, 0x7e, 0xba,
	0x30, 0xcd, 0x5a, 0x4b, 0x87, 0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x64, 0x65, 0x76, 0x6e, 0x65,
	0x74,
}
