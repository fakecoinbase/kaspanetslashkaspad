package main

import (
	"github.com/kaspanet/kaspad/rpcclient"
	"github.com/kaspanet/kaspad/util"
	"github.com/kaspanet/kaspad/wire"
	"github.com/pkg/errors"
	"io/ioutil"
	"time"
)

type minerClient struct {
	*rpcclient.Client
	onBlockAdded chan struct{}
}

func newMinerClient(connCfg *rpcclient.ConnConfig) (*minerClient, error) {
	client := &minerClient{
		onBlockAdded: make(chan struct{}, 1),
	}
	notificationHandlers := &rpcclient.NotificationHandlers{
		OnFilteredBlockAdded: func(_ uint64, header *wire.BlockHeader,
			txs []*util.Tx) {
			client.onBlockAdded <- struct{}{}
		},
	}
	var err error
	client.Client, err = rpcclient.New(connCfg, notificationHandlers)
	if err != nil {
		return nil, errors.Errorf("Error connecting to address %s: %s", connCfg.Host, err)
	}

	if err = client.NotifyBlocks(); err != nil {
		return nil, errors.Errorf("Error while registering client %s for block notifications: %s", client.Host(), err)
	}
	return client, nil
}

func connectToServer(cfg *configFlags) (*minerClient, error) {
	cert, err := readCert(cfg)
	if err != nil {
		return nil, err
	}

	rpcAddr, err := cfg.NetParams().NormalizeRPCServerAddress(cfg.RPCServer)
	if err != nil {
		return nil, err
	}

	connCfg := &rpcclient.ConnConfig{
		Host:           rpcAddr,
		Endpoint:       "ws",
		User:           cfg.RPCUser,
		Pass:           cfg.RPCPassword,
		DisableTLS:     cfg.DisableTLS,
		RequestTimeout: time.Second * 10,
		Certificates:   cert,
	}

	client, err := newMinerClient(connCfg)
	if err != nil {
		return nil, err
	}

	log.Infof("Connected to server %s", client.Host())

	return client, nil
}

func readCert(cfg *configFlags) ([]byte, error) {
	if cfg.DisableTLS {
		return nil, nil
	}

	cert, err := ioutil.ReadFile(cfg.RPCCert)
	if err != nil {
		return nil, errors.Errorf("Error reading certificates file: %s", err)
	}

	return cert, nil
}
