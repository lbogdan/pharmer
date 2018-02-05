package gke

import (
	"context"
	"fmt"

	"github.com/appscode/go/errors"
	api "github.com/pharmer/pharmer/apis/v1alpha1"
	. "github.com/pharmer/pharmer/cloud"
	container "google.golang.org/api/container/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func encodeCluster(ctx context.Context, cluster *api.Cluster) (*container.Cluster, error) {
	nodeGroups, err := Store(ctx).NodeGroups(cluster.Name).List(metav1.ListOptions{})
	if err != nil {
		err = errors.FromErr(err).WithContext(ctx).Err()
		return nil, err
	}

	nodePools := make([]*container.NodePool, 0)
	for _, node := range nodeGroups {
		np := &container.NodePool{
			Config: &container.NodeConfig{
				MachineType: node.Spec.Template.Spec.SKU,
				DiskSizeGb:  node.Spec.Template.Spec.DiskSize,
				ImageType:   cluster.Spec.Cloud.InstanceImage,
			},
			InitialNodeCount: node.Spec.Nodes,
			Name:             node.Name,
		}
		nodePools = append(nodePools, np)
	}

	kluster := &container.Cluster{
		ClusterIpv4Cidr: cluster.Spec.Networking.PodSubnet,
		Name:            cluster.Name,
		//InitialNodeCount: totalNodes,
		InitialClusterVersion: cluster.Spec.KubernetesVersion,

		MasterAuth: &container.MasterAuth{
			Username: cluster.Spec.Cloud.GKE.UserName,
			Password: cluster.Spec.Cloud.GKE.Password,
		},
		Network: cluster.Spec.Cloud.GKE.NetworkName,
		NetworkPolicy: &container.NetworkPolicy{
			Enabled:  true,
			Provider: cluster.Spec.Networking.NetworkProvider,
		},
		NodePools: nodePools,
		/*NodeConfig: &container.NodeConfig{
			MachineType: node.Spec.Template.Spec.SKU,
			DiskSizeGb: node.Spec.Template.Spec.DiskSize,
			ImageType: cluster.Spec.Cloud.InstanceImage,
		},*/
	}

	fmt.Println("........", kluster, ".............")
	return kluster, nil
}

func decodeCluster(cluster *container.Cluster) *api.Cluster {
	return &api.Cluster{}
}
