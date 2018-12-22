package gotor_changeip

import (
	"context"
	"github.com/cretz/bine/tor"
	"io/ioutil"
	"log"
	"net/http"
)

func (t *TorClient) Close() {
	t.dialCancel()
	err := t.tor.Close()
	if err != nil {
		log.Fatal(err)
	}
	t.ip = IP_NOT_FOUND
	t.Client = nil
	t.tor = nil
	t.dialCancel = nil
	t.dialCtx = nil
	t.dialer = nil

}

func (t *TorClient) createDialContext() error {
	var err error
	t.dialCtx, t.dialCancel = context.WithCancel(context.Background())
	t.dialer, err = t.tor.Dialer(t.dialCtx, nil)
	if err != nil {
		return err
	}
	return nil
}

func (t *TorClient) createInstance() error {
	instance, err := tor.Start(nil, &TOR_CONFIG)
	if err != nil {
		return err
	}
	t.tor = instance
	return nil
}

func (t *TorClient) GetClient() {
	t.Client = &http.Client{
		Transport: &http.Transport{
			DialContext: t.dialer.DialContext,
		},
	}
}

func (t *TorClient) Renew() error {
	err := t.tor.Control.Signal("NEWNYM")
	if err != nil {
		return err
	}
	t.GetClient()
	t.ip = ""
	return nil
}

func NewClient() (*TorClient, error) {
	t := &TorClient{
		ip: IP_NOT_FOUND,
	}
	err := t.createInstance()
	if err != nil {
		return nil, err
	}
	err = t.createDialContext()
	if err != nil {
		return nil, err
	}
	t.GetClient()
	return t, nil
}

func (t *TorClient) CurrentIP() string {
	if t.ip != "" && t.ip != IP_NOT_FOUND {
		return t.ip
	}
	resp, err := t.Client.Get(IP_CHECK_URL)
	if err != nil {
		return IP_NOT_FOUND
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IP_NOT_FOUND
	}
	return string(body)

}
