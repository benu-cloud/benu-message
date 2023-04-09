package config

import (
	"fmt"
	"strconv"

	pkgerrors "github.com/benu-cloud/benu-errors"
)

func (p *PortNumber) String() string {
	return fmt.Sprintf("%d", uint(*p))
}

func (p *PortNumber) Set(s string) error {
	portnum, err := strconv.Atoi(s)
	if err != nil {
		goto badFormat
	}
	if portnum < 0 || portnum > 65535 {
		goto badFormat
	}
	return nil
badFormat:
	return pkgerrors.NewBadCommanlineArgument("Port", s, "0-65535")
}
