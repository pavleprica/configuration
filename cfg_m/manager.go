package cfg_m

import (
	"errors"
	"fmt"
)

var (
	configurationValueNotFoundMessage = "configuration value not found"
)

// ConfigStringGetter is responsible for allowing getting of loaded configuration from the implemented types.
type ConfigStringGetter interface {
	// Get returns the configuration value that has been previously loaded/
	// In event that the value isn't present it returns an error.
	Get(configurationName string) (string, error)
}

// ConfigStringGetterWithDefault is responsible for allowing getting of loaded configuration from the implemented types.
// With the addition of not returning an error, rather the provided default value
type ConfigStringGetterWithDefault interface {
	// GetOrDefault returns the configuration value that has been previously loaded/
	// In event that the value isn't present it returns the provided default value.
	GetOrDefault(configurationName string, defaultValue string) string
}

// Manager is used as a loaded configuration store.
// It provides a set of methods
type Manager struct {
	loadedConfiguration map[string]string
}

func NewManager(configuration map[string]string) *Manager {
	return &Manager{
		loadedConfiguration: configuration,
	}
}

func (m *Manager) Get(configurationName string) (string, error) {
	cfgValue, present := m.loadedConfiguration[configurationName]
	if present {
		return cfgValue, nil
	} else {
		return "", errors.New(fmt.Sprintf("%s - %s", configurationName, configurationValueNotFoundMessage))
	}
}

func (m *Manager) GetOrDefault(configurationName string, defaultValue string) string {
	cfgValue, present := m.loadedConfiguration[configurationName]
	if present {
		return cfgValue
	} else {
		return defaultValue
	}
}
