package gke

import (
	"fmt"
	//"github.com/appscode/go/errors"
	api "github.com/pharmer/pharmer/apis/v1alpha1"
	. "github.com/pharmer/pharmer/cloud"
	//	core "k8s.io/api/core/v1"
	//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"k8s.io/client-go/kubernetes"
	container "google.golang.org/api/container/v1"
	//"github.com/appscode/errors"
)

func (cm *ClusterManager) Apply(in *api.Cluster, dryRun bool) ([]api.Action, error) {
	var err error
	var acts []api.Action

	if in.Status.Phase == "" {
		return nil, fmt.Errorf("cluster `%s` is in unknown phase", cm.cluster.Name)
	}
	if in.Status.Phase == api.ClusterDeleted {
		return nil, nil
	}
	cm.cluster = in
	cm.namer = namer{cluster: cm.cluster}
	if cm.ctx, err = LoadCACertificates(cm.ctx, cm.cluster); err != nil {
		return nil, err
	}
	/*if cm.ctx, err = LoadSSHKey(cm.ctx, cm.cluster); err != nil {
		return nil, err
	}*/
	if cm.conn, err = NewConnector(cm.ctx, cm.cluster); err != nil {
		return nil, err
	}
	cm.conn.namer = cm.namer

	if cm.cluster.Status.Phase == api.ClusterPending {
		a, err := cm.applyCreate(dryRun)
		if err != nil {
			return nil, err
		}
		acts = append(acts, a...)
	}

	return acts, nil
}

func (cm *ClusterManager) applyCreate(dryRun bool) (acts []api.Action, err error) {
	found, _ := cm.conn.getNetworks()
	if !found {
		if err = cm.conn.ensureNetworks(); err != nil {
			return
		}
	}
	cluster, err := encodeCluster(cm.ctx, cm.cluster)
	fmt.Println(err)
	clusterRequest := &container.CreateClusterRequest{
		Cluster: cluster,
	}
	resp, err := cm.conn.containerService.Projects.Zones.Clusters.Create(cm.conn.cluster.Spec.Cloud.Project, cm.conn.cluster.Spec.Cloud.Zone, clusterRequest).Do()
	fmt.Println(resp, err, resp.OperationType)
	if err = cm.conn.waitForZoneOperation(resp.Name); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return acts, err
	}

//	cm.conn.containerService.Projects.Zones.Clusters.Get()
	return
}
