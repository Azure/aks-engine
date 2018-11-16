// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package v20180331

import "github.com/pkg/errors"

// ErrorInvalidNetworkProfile error
var ErrorInvalidNetworkProfile = errors.New("ServiceCidr, DNSServiceIP, DockerBridgeCidr should all be empty or neither should be empty")

// ErrorPodCidrNotSetableInAzureCNI error
var ErrorPodCidrNotSetableInAzureCNI = errors.New("PodCidr should not be set when network plugin is set to Azure")

// ErrorInvalidNetworkPlugin error
var ErrorInvalidNetworkPlugin = errors.New("Network plugin should be either Azure or Kubenet")

// ErrorInvalidServiceCidr error
var ErrorInvalidServiceCidr = errors.New("ServiceCidr is not a valid CIDR")

// ErrorServiceCidrTooLarge error
var ErrorServiceCidrTooLarge = errors.New("ServiceCidr is too large")

// ErrorInvalidDNSServiceIP error
var ErrorInvalidDNSServiceIP = errors.New("DNSServiceIP is not a valid IP address")

// ErrorInvalidDockerBridgeCidr error
var ErrorInvalidDockerBridgeCidr = errors.New("DockerBridgeCidr is not a valid IP address")

// ErrorDNSServiceIPNotInServiceCidr error
var ErrorDNSServiceIPNotInServiceCidr = errors.New("DNSServiceIP is not within ServiceCidr")

// ErrorDNSServiceIPAlreadyUsed error
var ErrorDNSServiceIPAlreadyUsed = errors.New("DNSServiceIP can not be the first IP address in ServiceCidr")

// ErrorAtLeastAgentPoolNoSubnet error
var ErrorAtLeastAgentPoolNoSubnet = errors.New("At least one agent pool does not have subnet defined")

// ErrorInvalidMaxPods error
var ErrorInvalidMaxPods = errors.New("Max pods per node needs to be at least 5")

// ErrorParsingSubnetID error
var ErrorParsingSubnetID = errors.New("Failed to parse VnetSubnetID")

// ErrorSubscriptionNotMatch error
var ErrorSubscriptionNotMatch = errors.New("Subscription for subnet does not match with other subnet")

// ErrorResourceGroupNotMatch error
var ErrorResourceGroupNotMatch = errors.New("ResourceGroup for subnet does not match with other subnet")

// ErrorVnetNotMatch error
var ErrorVnetNotMatch = errors.New("Vnet for subnet does not match with other subnet")

// ErrorRBACNotEnabledForAAD error
var ErrorRBACNotEnabledForAAD = errors.New("RBAC must be enabled for AAD to be enabled")

// ErrorAADServerAppIDNotSet error
var ErrorAADServerAppIDNotSet = errors.New("ServerAppID in AADProfile cannot be empty string")

// ErrorAADServerAppSecretNotSet error
var ErrorAADServerAppSecretNotSet = errors.New("ServerAppSecret in AADProfile cannot be empty string")

// ErrorAADClientAppIDNotSet error
var ErrorAADClientAppIDNotSet = errors.New("ClientAppID in AADProfile cannot be empty string")

// ErrorAADTenantIDNotSet error
var ErrorAADTenantIDNotSet = errors.New("TenantID in AADProfile cannot be empty string")
