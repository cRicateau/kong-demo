package main

import (
	"github.com/Kong/go-pdk"
)

// Config represents to config parameters into the config.yml
type Config struct {
	Apikey string
}

// New is a func that your plugin must define
// it creates an instance of this type and returns as an interface{}
func New() interface{} {
	return &Config{}
}

// Access is called on the access kong event
// To handle Kong events, define the relevant methods on your configuration structure.
// For example, to handle the “access” event, define a function like this
// The event methods you can define are Certificate, Rewrite, Access, Preread and Log. The signature is the same for all of them
func (conf Config) Access(kong *pdk.PDK) {
	key, err := kong.Request.GetQueryArg("key")
	apiKey := conf.Apikey

	if err != nil {
		kong.Log.Err(err.Error())
	}

	//it adjusts the header parameters in this way.
	x := make(map[string][]string)
	x["Content-Type"] = append(x["Content-Type"], "application/json")

	//If the key of the consumer is not equal to the claimed key, kong doesn't ensure the proxy
	if apiKey != key {
		kong.Response.Exit(403, "You have no correct consumer key.", x)
	}
}
