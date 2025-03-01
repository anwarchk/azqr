// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package fdfp

import "github.com/Azure/azqr/internal/scanners"

func init() {
	scanners.ScannerList["fdfp"] = []scanners.IAzureScanner{&FrontDoorWAFPolicyScanner{}}
}

// FrontDoorWAFPolicyScanner - Scanner for Front Door Web Application Policy
type FrontDoorWAFPolicyScanner struct {
	config *scanners.ScannerConfig
}

// Init - Initializes the Front Door Web Application Policy Scanner
func (a *FrontDoorWAFPolicyScanner) Init(config *scanners.ScannerConfig) error {
	a.config = config
	return nil
}

// Scan - Scans all Front Door Web Application Policy in a Resource Group
func (a *FrontDoorWAFPolicyScanner) Scan(scanContext *scanners.ScanContext) ([]scanners.AzqrServiceResult, error) {
	scanners.LogSubscriptionScan(a.config.SubscriptionID, a.ResourceTypes()[0])
	return []scanners.AzqrServiceResult{}, nil
}

func (a *FrontDoorWAFPolicyScanner) ResourceTypes() []string {
	return []string{"Microsoft.Network/frontdoorWebApplicationFirewallPolicies"}
}

func (a *FrontDoorWAFPolicyScanner) GetRecommendations() map[string]scanners.AzqrRecommendation {
	return map[string]scanners.AzqrRecommendation{}
}
