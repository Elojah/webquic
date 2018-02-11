package webquic

import (
	"errors"
)

// Config is web quic server structure config.
type Config struct {
	Address    string `json:"address"`
	CertKey    string `json:"cert-key"`
	CertPem    string `json:"cert-pem"`
	Dispatcher int    `json:"dispatcher"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c.Address == rhs.Address
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}
	if err := c.dialAdress(fconf); err != nil {
		return err
	}
	if err := c.dialCert(fconf); err != nil {
		return err
	}
	cDispatcher, ok := fconf["dispatcher"]
	if !ok {
		return errors.New("missing key dispatcher")
	}
	cDispatcherFloat, ok := cDispatcher.(float64)
	if !ok {
		return errors.New("key dispatcher invalid. must be number")
	}
	c.Dispatcher = int(cDispatcherFloat)
	return nil
}

func (c *Config) dialAdress(fconf map[string]interface{}) error {
	cAddress, ok := fconf["address"]
	if !ok {
		return errors.New("missing key address")
	}
	if c.Address, ok = cAddress.(string); !ok {
		return errors.New("key address invalid. must be string")
	}
	return nil
}

func (c *Config) dialCert(fconf map[string]interface{}) error {
	cKey, ok := fconf["cert-key"]
	if !ok {
		return errors.New("missing key cert-key")
	}
	if c.CertKey, ok = cKey.(string); !ok {
		return errors.New("key cert-key invalid. must be string")
	}
	cPem, ok := fconf["cert-pem"]
	if !ok {
		return errors.New("missing key cert-pem")
	}
	if c.CertPem, ok = cPem.(string); !ok {
		return errors.New("key cert-pem invalid. must be string")
	}
	return nil
}
