package aci

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/Azure/azure-sdk-for-go/services/containerinstance/mgmt/2021-09-01/containerinstance"
)

// NewContainerGroupDiagnostics creates a container group diagnostics object
func NewContainerGroupDiagnostics(logAnalyticsID, logAnalyticsKey string) (*containerinstance.ContainerGroupDiagnostics, error) {

	if logAnalyticsID == "" || logAnalyticsKey == "" {
		return nil, errors.New("Log Analytics configuration requires both the workspace ID and Key")
	}

	return &containerinstance.ContainerGroupDiagnostics{
		LogAnalytics: &containerinstance.LogAnalytics{
			WorkspaceID:  &logAnalyticsID,
			WorkspaceKey: &logAnalyticsKey,
		},
	}, nil
}

// NewContainerGroupDiagnosticsFromFile creates a container group diagnostics object from the specified file
func NewContainerGroupDiagnosticsFromFile(filepath string) (*ContainerGroupDiagnostics, error) {

	analyticsdata, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Reading Log Analytics Auth file %q failed: %v", filepath, err)
	}
	// Unmarshal the log analytics file.
	var law LogAnalyticsWorkspace
	if err := json.Unmarshal(analyticsdata, &law); err != nil {
		return nil, err
	}

	return &ContainerGroupDiagnostics{
		LogAnalytics: &law,
	}, nil
}
