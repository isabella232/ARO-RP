package subnet

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	mgmtnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-07-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/golang/mock/gomock"

	"github.com/Azure/ARO-RP/pkg/api"
	mock_network "github.com/Azure/ARO-RP/pkg/util/mocks/azureclient/mgmt/network"
)

func TestGet(t *testing.T) {
	ctx := context.Background()

	type test struct {
		name       string
		subnetID   string
		mocks      func(*test, *mock_network.MockSubnetsClient)
		wantSubnet *mgmtnetwork.Subnet
		wantErr    string
	}

	for _, tt := range []*test{
		{
			name:     "valid",
			subnetID: "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet",
			mocks: func(tt *test, subnets *mock_network.MockSubnetsClient) {
				subnets.EXPECT().
					Get(ctx, "vnetResourceGroup", "vnet", "subnet", "").
					Return(*tt.wantSubnet, nil)
			},
			wantSubnet: &mgmtnetwork.Subnet{
				Name: to.StringPtr("subnet"),
			},
		},
		{
			name:     "internal error",
			subnetID: "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet",
			mocks: func(tt *test, subnets *mock_network.MockSubnetsClient) {
				subnets.EXPECT().
					Get(ctx, "vnetResourceGroup", "vnet", "subnet", "").
					Return(mgmtnetwork.Subnet{}, fmt.Errorf("random error"))
			},
			wantErr: "random error",
		},
		{
			name:     "path error",
			subnetID: "/invalid/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet",
			wantErr:  "parsing failed for /invalid/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet. Invalid resource Id format",
		},
		{
			name:     "invalid",
			subnetID: "invalid",
			wantErr:  `subnet ID "invalid" has incorrect length`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			subnets := mock_network.NewMockSubnetsClient(controller)
			if tt.mocks != nil {
				tt.mocks(tt, subnets)
			}

			m := &manager{
				subnets: subnets,
			}

			subnet, err := m.Get(ctx, tt.subnetID)
			if err != nil && err.Error() != tt.wantErr ||
				err == nil && tt.wantErr != "" {
				t.Error(err)
			}

			if !reflect.DeepEqual(subnet, tt.wantSubnet) {
				t.Error(subnet)
			}
		})
	}
}

func TestCreateOrUpdate(t *testing.T) {
	ctx := context.Background()

	type test struct {
		name     string
		subnetID string
		mocks    func(*test, *mock_network.MockSubnetsClient)
		wantErr  string
	}

	for _, tt := range []*test{
		{
			name:     "valid",
			subnetID: "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet",
			mocks: func(tt *test, subnets *mock_network.MockSubnetsClient) {
				subnets.EXPECT().
					CreateOrUpdateAndWait(ctx, "vnetResourceGroup", "vnet", "subnet", mgmtnetwork.Subnet{}).
					Return(nil)
			},
		},
		{
			name:     "internal error",
			subnetID: "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet",
			mocks: func(tt *test, subnets *mock_network.MockSubnetsClient) {
				subnets.EXPECT().
					CreateOrUpdateAndWait(ctx, "vnetResourceGroup", "vnet", "subnet", mgmtnetwork.Subnet{}).
					Return(fmt.Errorf("random error"))
			},
			wantErr: "random error",
		},
		{
			name:     "path error",
			subnetID: "/invalid/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet",
			wantErr:  "parsing failed for /invalid/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet. Invalid resource Id format",
		},
		{
			name:     "invalid",
			subnetID: "invalid",
			wantErr:  `subnet ID "invalid" has incorrect length`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			subnets := mock_network.NewMockSubnetsClient(controller)
			if tt.mocks != nil {
				tt.mocks(tt, subnets)
			}

			m := &manager{
				subnets: subnets,
			}

			err := m.CreateOrUpdate(ctx, tt.subnetID, &mgmtnetwork.Subnet{})
			if err != nil && err.Error() != tt.wantErr ||
				err == nil && tt.wantErr != "" {
				t.Error(err)
			}
		})
	}
}

func TestNetworkSecurityGroupID(t *testing.T) {
	oc := &api.OpenShiftCluster{
		Properties: api.Properties{
			ClusterProfile: api.ClusterProfile{
				ResourceGroupID: "/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup",
			},
			MasterProfile: api.MasterProfile{
				SubnetID: "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master",
			},
			WorkerProfiles: []api.WorkerProfile{
				{
					SubnetID: "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker",
				},
			},
		},
	}

	for _, tt := range []struct {
		name      string
		subnetID  string
		wantNSGID string
		wantErr   string
	}{
		{
			name:      "master",
			subnetID:  "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master",
			wantNSGID: "/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup/providers/Microsoft.Network/networkSecurityGroups/aro-controlplane-nsg",
		},
		{
			name:      "worker",
			subnetID:  "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker",
			wantNSGID: "/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup/providers/Microsoft.Network/networkSecurityGroups/aro-node-nsg",
		},
		{
			name:     "invalid",
			subnetID: "invalid",
			wantErr:  `unknown subnetID "invalid"`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			nsgID, err := NetworkSecurityGroupID(oc, tt.subnetID)
			if err != nil && err.Error() != tt.wantErr ||
				err == nil && tt.wantErr != "" {
				t.Error(err)
			}

			if nsgID != tt.wantNSGID {
				t.Error(nsgID)
			}
		})
	}
}
