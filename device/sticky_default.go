//go:build !linux

package device

import (
	"github.com/kagari-org/wireguard-go/conn"
	"github.com/kagari-org/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
