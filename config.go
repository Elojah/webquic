package webquic

import (
	"errors"
)

// Config is web quic server structure config.
type Config struct {
	URL     string `json:"url"`
	Port    string `json:"port"`
	CertKey string `json:"cert-key"`
	CertPem string `json:"cert-pem"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c.URL == rhs.URL &&
		c.Port == rhs.Port
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}
	cURL, ok := fconf["url"]
	if !ok {
		return errors.New("missing key url")
	}
	if c.URL, ok = cURL.(string); !ok {
		return errors.New("key url invalid. must be string")
	}
	cPort, ok := fconf["port"]
	if !ok {
		return errors.New("missing key port")
	}
	if c.Port, ok = cPort.(string); !ok {
		return errors.New("key port invalid. must be string")
	}
	cKey, ok := fconf["cert-key"]
	if !ok {
		return errors.New("missing key cert-key")
	}
	if c.CertKey, ok = cKey.(string); !ok {
		return errors.New("key cert-key invalid. must be string")
	}
	cPerm, ok := fconf["cert-pem"]
	if !ok {
		return errors.New("missing key cert-pem")
	}
	if c.CertPem, ok = cPerm.(string); !ok {
		return errors.New("key cert-pem invalid. must be string")
	}
	return nil
}
