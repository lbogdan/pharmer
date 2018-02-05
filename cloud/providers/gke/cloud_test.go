package gke

import (
	"testing"
	//. "github.com/pharmer/pharmer/cloud"
	"golang.org/x/oauth2/google"
	container "google.golang.org/api/container/v1"
	"context"

	"fmt"
)
func TestCL(t *testing.T) {
	data := ``
project := "k8s-qa"
	conf, err := google.JWTConfigFromJSON([]byte(data),
		container.CloudPlatformScope)
	fmt.Println(err)
	client := conf.Client(context.Background())
	containerService, err := container.New(client)
	resp, err := containerService.Projects.Zones.Clusters.Get(project, "us-central1-f", "gk2").Context(context.Background()).Do()
	fmt.Println(resp, err)

}

