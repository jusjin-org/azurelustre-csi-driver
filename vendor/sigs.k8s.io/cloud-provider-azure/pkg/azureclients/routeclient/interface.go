/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package routeclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"

	"sigs.k8s.io/cloud-provider-azure/pkg/retry"
)

const (
	// APIVersion is the API version for network.
	APIVersion = "2020-08-01"
	// AzureStackCloudAPIVersion is the API version for Azure Stack
	AzureStackCloudAPIVersion = "2018-11-01"
	// AzureStackCloudName is the cloud name of Azure Stack
	AzureStackCloudName = "AZURESTACKCLOUD"
)

// Interface is the client interface for Route.
// Don't forget to run the following command to generate the mock client:
// mockgen -source=$GOPATH/src/sigs.k8s.io/cloud-provider-azure/pkg/azureclients/routeclient/interface.go -package=mockrouteclient Interface > $GOPATH/src/sigs.k8s.io/cloud-provider-azure/pkg/azureclients/routeclient/mockrouteclient/interface.go
type Interface interface {
	// CreateOrUpdate creates or updates a Route.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters network.Route, etag string) *retry.Error

	// Delete deletes a Route by name.
	Delete(ctx context.Context, resourceGroupName string, routeTableName string, routeName string) *retry.Error
}
