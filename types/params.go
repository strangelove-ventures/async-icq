package types

import (
	"fmt"
	"strings"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultHostEnabled is the default value for the host param (set to true)
	DefaultHostEnabled = true
)

var (
	// KeyHostEnabled is the store key for HostEnabled Params
	KeyHostEnabled = []byte("HostEnabled")
	// KeyAllowQueries is the store key for the AllowQueries Params
	KeyAllowQueries = []byte("AllowQueries")
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new parameter configuration
func NewParams(enableHost bool, allowQueries []string) Params {
	return Params{
		HostEnabled:  enableHost,
		AllowQueries: allowQueries,
	}
}

// DefaultParams is the default parameter configuration
func DefaultParams() Params {
	return NewParams(DefaultHostEnabled, nil)
}

// Validate validates all parameters
func (p Params) Validate() error {
	if err := validateEnabled(p.HostEnabled); err != nil {
		return err
	}
	return validateAllowlist(p.AllowQueries)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyHostEnabled, p.HostEnabled, validateEnabled),
		paramtypes.NewParamSetPair(KeyAllowQueries, p.AllowQueries, validateAllowlist),
	}
}

func validateEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateAllowlist(i interface{}) error {
	allowQueries, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, path := range allowQueries {
		if strings.TrimSpace(path) == "" {
			return fmt.Errorf("parameter must not contain empty strings: %s", allowQueries)
		}
	}

	return nil
}
