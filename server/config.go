package server

import (
	"github.com/blend/go-sdk/web"
)

// Config is the config for a slack server
type Config struct {
	web.Config           `json:",inline" yaml:",inline"`
	SlackSignatureSecret string `json:"slackSignatureSecret" yaml:"slackSignatureSecret"`
	AcknowledgeOnVerify  bool   `json:"acknowledgeOnVerify" yaml:"acknowledgeOnVerify"`
}

// Status is the status of the server
type Status struct {
	Ready bool `json:"ready"`
}
