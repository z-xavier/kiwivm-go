package kiwivm

import (
	"io"
	"net"
	"time"

	"github.com/bytedance/sonic"
	"resty.dev/v3"
)

type ClientOptFunc func(client *resty.Client)

func WithDebug(debug bool) ClientOptFunc {
	return func(client *resty.Client) {
		client.SetDebug(debug)
	}
}

func WithLogger(logger resty.Logger) ClientOptFunc {
	return func(client *resty.Client) {
		client.SetLogger(logger)
	}
}

func WithTimeout(timeout time.Duration) ClientOptFunc {
	return func(client *resty.Client) {
		client.SetTimeout(timeout)
	}
}

func DecodeJson(r io.Reader, v any) error {
	dec := sonic.ConfigStd.NewDecoder(r)
	for {
		if err := dec.Decode(v); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}

func MarshalString(v any) string {
	s, _ := sonic.MarshalString(v)
	return s
}

func IntToNetIP4(ipv4Int int64) net.IP {
	return net.IPv4(byte(ipv4Int>>24), byte(ipv4Int>>16), byte(ipv4Int>>8), byte(ipv4Int))
}

func IntToNetIP4Str(ipv4Int int64) string {
	return net.IPv4(byte(ipv4Int>>24), byte(ipv4Int>>16), byte(ipv4Int>>8), byte(ipv4Int)).To4().String()
}
