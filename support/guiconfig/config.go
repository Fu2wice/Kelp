package guiconfig

import (
	"github.com/stellar/kelp/support/utils"
)

type Auth0Config struct {
	Auth0Enabled bool   `valid:"-" toml:"AUTH0_ENABLED" json:"auth0_enabled"`
	Domain       string `valid:"-" toml:"DOMAIN"json:"domain"`
	ClientId     string `valid:"-" toml:"CLIENT_ID"json:"client_id"`
	Audience     string `valid:"-" toml:"AUDIENCE"json:"audience"`
}

type GUIConfig struct {
	Auth0Config 		*Auth0Config `valid:"-" toml:"AUTH0" json:"auth0"`
}

// String impl.
func (g GUIConfig) String() string {
	return utils.StructString(g, 0, map[string]func(interface{}) interface{}{
		"CLIENT_ID":        utils.Hide,
		"DOMAIN":        	utils.Hide,
	})
}

// EmptyGuiConfig returns an empty GUIConfig for when it is not specified on the command-line
func EmptyGuiConfig() GUIConfig {
	return GUIConfig{
		Auth0Config: &Auth0Config{},
	}
}