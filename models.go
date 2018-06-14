package manssh

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/xwjdsh/ssh_config"
)

// HostConfig struct include alias, connect string and other config
type HostConfig struct {
	Alias string
	// Path found in which file
	Path    string
	PathMap map[string][]*ssh_config.Host

	OwnConfig      map[string]string
	ImplicitConfig map[string]string
}

func NewHostConfig(alias, path string, host *ssh_config.Host) *HostConfig {
	return &HostConfig{
		Alias:          alias,
		Path:           path,
		PathMap:        map[string][]*ssh_config.Host{path: {host}},
		OwnConfig:      map[string]string{},
		ImplicitConfig: map[string]string{},
	}
}

func (hc *HostConfig) ConnectStr() string {
	if !hc.Display() {
		return ""
	}

	var (
		user, hostname, port string
		ok                   bool
	)

	if user, ok = hc.OwnConfig["user"]; !ok {
		user = hc.ImplicitConfig["user"]
		delete(hc.ImplicitConfig, "user")
	} else {
		user = color.GreenString(user)
		delete(hc.OwnConfig, "user")
	}

	if hostname, ok = hc.OwnConfig["hostname"]; !ok {
		delete(hc.ImplicitConfig, "hostname")
		hostname = hc.ImplicitConfig["hostname"]
	} else {
		hostname = color.GreenString(hostname)
		delete(hc.OwnConfig, "hostname")
	}

	if port, ok = hc.OwnConfig["port"]; !ok {
		port = hc.ImplicitConfig["port"]
		delete(hc.ImplicitConfig, "port")
	} else {
		port = color.GreenString(port)
		delete(hc.OwnConfig, "port")
	}

	return fmt.Sprintf("%s%s%s%s%s", user, color.GreenString("@"), hostname, color.GreenString(":"), port)
}

func (hc *HostConfig) Display() bool {
	hostname := hc.OwnConfig["hostname"]
	if hostname == "" {
		hostname = hc.ImplicitConfig["hostname"]
	}

	return hostname != ""
}
