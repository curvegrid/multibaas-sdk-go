package multibaas

import (
	"errors"
)

// IsMultiBaasErr returns a *multibaas.Error and `true` if the if the error provided is a standard MultiBaas API error.
func IsMultiBaasErr(err error) (*Error, bool) {
	var oaErr *GenericOpenAPIError
	if errors.As(err, &oaErr) {
		if mbErr, ok := oaErr.Model().(Error); ok {
			return &mbErr, ok
		}
	}
	return nil, false
}
