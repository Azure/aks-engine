// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package v20170701

import "testing"

func TestMerge(t *testing.T) {
	newCS := &ContainerService{
		Properties: &Properties{
			ServicePrincipalProfile: &ServicePrincipalProfile{
				ClientID: "fakeID",
				Secret:   "",
				KeyvaultSecretRef: &KeyvaultSecretRef{
					VaultID:    "keyVaultRefNew",
					SecretName: "secret-name",
				},
			},
			WindowsProfile: &WindowsProfile{
				AdminUsername: "azureuser",
				AdminPassword: "",
			},
		},
	}

	existingCS := &ContainerService{
		Properties: &Properties{
			ServicePrincipalProfile: &ServicePrincipalProfile{
				ClientID: "existingFakeID",
				Secret:   "existingSecret",
				KeyvaultSecretRef: &KeyvaultSecretRef{
					VaultID:    "keyVaultRefExisting",
					SecretName: "secret-name",
				},
			},
			WindowsProfile: &WindowsProfile{
				AdminUsername: "azureuser",
				AdminPassword: "existingPassword",
			},
		},
	}
	if err := newCS.Merge(existingCS); err != nil {
		t.Fatalf("unexpectedly detected merge failure, %+v", err)
	}
	if newCS.Properties.ServicePrincipalProfile.ClientID != "fakeID" {
		t.Fatalf("unexpected Properties.ServicePrincipalProfile.ClientID changed")
	}
	if newCS.Properties.ServicePrincipalProfile.Secret != "existingSecret" {
		t.Fatalf("unexpected Properties.ServicePrincipalProfile.Secret not updated")
	}
	if newCS.Properties.ServicePrincipalProfile.KeyvaultSecretRef.VaultID != "keyVaultRefNew" {
		t.Fatalf("unexpected Properties.ServicePrincipalProfile.KeyvaultSecretRef changed")
	}
	if newCS.Properties.WindowsProfile.AdminPassword != "existingPassword" {
		t.Fatalf("unexpected Properties.WindowsProfile.AdminPassword not updated")
	}

}

func TestMergeWithNil(t *testing.T) {
	newCS := &ContainerService{
		Properties: &Properties{
			ServicePrincipalProfile: &ServicePrincipalProfile{
				ClientID: "fakeID",
				Secret:   "",
				KeyvaultSecretRef: &KeyvaultSecretRef{
					VaultID:    "keyVaultRefNew",
					SecretName: "secret-name",
				},
			},
		},
	}

	existingCS := &ContainerService{
		Properties: &Properties{
			ServicePrincipalProfile: &ServicePrincipalProfile{
				ClientID: "existingFakeID",
				Secret:   "existingSecret",
				KeyvaultSecretRef: &KeyvaultSecretRef{
					VaultID:    "keyVaultRefExisting",
					SecretName: "secret-name",
				},
			},
			WindowsProfile: &WindowsProfile{
				AdminUsername: "azureuser",
				AdminPassword: "existingPassword",
			},
		},
	}
	if err := newCS.Merge(existingCS); err != nil {
		t.Fatalf("unexpectedly detected merge failure, %+v", err)
	}
	if newCS.Properties.ServicePrincipalProfile.ClientID != "fakeID" {
		t.Fatalf("unexpected Properties.ServicePrincipalProfile.ClientID changed")
	}
	if newCS.Properties.ServicePrincipalProfile.Secret != "existingSecret" {
		t.Fatalf("unexpected Properties.ServicePrincipalProfile.Secret not updated")
	}
	if newCS.Properties.ServicePrincipalProfile.KeyvaultSecretRef.VaultID != "keyVaultRefNew" {
		t.Fatalf("unexpected Properties.ServicePrincipalProfile.KeyvaultSecretRef changed")
	}
	if newCS.Properties.WindowsProfile == nil {
		t.Fatalf("unexpected Properties.WindowsProfile not updated")
	}
	if newCS.Properties.WindowsProfile.AdminUsername != "azureuser" {
		t.Fatalf("unexpected Properties.WindowsProfile.AdminUsername not updated")
	}
	if newCS.Properties.WindowsProfile.AdminPassword != "existingPassword" {
		t.Fatalf("unexpected Properties.WindowsProfile.AdminPassword not updated")
	}
}
