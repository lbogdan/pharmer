package aks

import (
	"context"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/azure-sdk-for-go/arm/containerservice"
	"github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/appscode/go/errors"
	. "github.com/appscode/go/types"
	api "github.com/pharmer/pharmer/apis/v1alpha1"
	. "github.com/pharmer/pharmer/cloud"
	"github.com/pharmer/pharmer/credential"
)

const (
	machineIDTemplate = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/%s"
	CloudProviderName = "azure"
)

var providerIDRE = regexp.MustCompile(`^` + CloudProviderName + `://(?:.*)/Microsoft.Compute/virtualMachines/(.+)$`)

type cloudConnector struct {
	ctx     context.Context
	cluster *api.Cluster
	namer   namer

	availabilitySetsClient  compute.AvailabilitySetsClient
	groupsClient            resources.GroupsClient
	containerSvcClient 		containerservice.ContainerServicesClient
}

func NewConnector(ctx context.Context, cluster *api.Cluster) (*cloudConnector, error) {
	cred, err := Store(ctx).Credentials().Get(cluster.Spec.CredentialName)
	if err != nil {
		return nil, err
	}
	typed := credential.Azure{CommonSpec: credential.CommonSpec(cred.Spec)}
	if ok, err := typed.IsValid(); !ok {
		return nil, errors.New().WithMessagef("Credential %s is invalid. Reason: %v", cluster.Spec.CredentialName, err)
	}

	baseURI := azure.PublicCloud.ResourceManagerEndpoint
	config, err := adal.NewOAuthConfig(azure.PublicCloud.ActiveDirectoryEndpoint, typed.TenantID())
	if err != nil {
		return nil, errors.FromErr(err).WithContext(ctx).Err()
	}

	spt, err := adal.NewServicePrincipalToken(*config, typed.ClientID(), typed.ClientSecret(), baseURI)
	if err != nil {
		return nil, errors.FromErr(err).WithContext(ctx).Err()
	}

	client := autorest.NewClientWithUserAgent(fmt.Sprintf("Azure-SDK-for-Go/%s", compute.Version()))
	client.Authorizer = autorest.NewBearerAuthorizer(spt)

	availabilitySetsClient := compute.AvailabilitySetsClient{
		ManagementClient: compute.ManagementClient{
			Client:         client,
			BaseURI:        baseURI,
			SubscriptionID: typed.SubscriptionID(),
		},
	}

	containerSvcClient := containerservice.ContainerServicesClient{
		ManagementClient: containerservice.ManagementClient{
			Client:         client,
			BaseURI:        baseURI,
			SubscriptionID: typed.SubscriptionID(),
		},
	}

	groupsClient := resources.GroupsClient{
		ManagementClient: resources.ManagementClient{
			Client:         client,
			BaseURI:        baseURI,
			SubscriptionID: typed.SubscriptionID(),
		},
	}




	return &cloudConnector{
		cluster: cluster,
		ctx:     ctx,
		availabilitySetsClient:  availabilitySetsClient,
		containerSvcClient: containerSvcClient,
		groupsClient:            groupsClient,
	}, nil
}

func (conn *cloudConnector) detectUbuntuImage() error {
	conn.cluster.Spec.Cloud.OS = "UbuntuServer"
	conn.cluster.Spec.Cloud.InstanceImageProject = "Canonical"
	conn.cluster.Spec.Cloud.InstanceImage = "16.04-LTS"
	conn.cluster.Spec.Cloud.Azure.InstanceImageVersion = "latest"
	return nil
}

func (conn *cloudConnector) getResourceGroup() (bool, error) {
	_, err := conn.groupsClient.Get(conn.namer.ResourceGroupName())
	return err == nil, err
}

func (conn *cloudConnector) ensureResourceGroup() (resources.Group, error) {
	req := resources.Group{
		Name:     StringP(conn.namer.ResourceGroupName()),
		Location: StringP(conn.cluster.Spec.Cloud.Zone),
		Tags: &map[string]*string{
			"KubernetesCluster": StringP(conn.cluster.Name),
		},
	}
	return conn.groupsClient.CreateOrUpdate(conn.namer.ResourceGroupName(), req)
}

func (conn *cloudConnector) getAvailabilitySet() (compute.AvailabilitySet, error) {
	return conn.availabilitySetsClient.Get(conn.namer.ResourceGroupName(), conn.namer.AvailabilitySetName())
}

func (conn *cloudConnector) ensureAvailabilitySet() (compute.AvailabilitySet, error) {
	name := conn.namer.AvailabilitySetName()
	req := compute.AvailabilitySet{
		Name:     StringP(name),
		Location: StringP(conn.cluster.Spec.Cloud.Zone),
		Tags: &map[string]*string{
			"KubernetesCluster": StringP(conn.cluster.Name),
		},
	}
	return conn.availabilitySetsClient.CreateOrUpdate(conn.namer.ResourceGroupName(), name, req)
}
