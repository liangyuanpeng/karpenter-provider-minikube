package cloudprovider

import (
	"context"
	_ "embed"
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/karpenter-provider-cluster-api/pkg/providers/machine"
	"sigs.k8s.io/karpenter-provider-cluster-api/pkg/providers/machinedeployment"
	v1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/karpenter/pkg/cloudprovider"
)

const (
	// TODO (elmiko) if we make these exposed constants from the CAS we can import them instead of redifining and risking drift
	cpuKey          = "capacity.cluster-autoscaler.kubernetes.io/cpu"
	memoryKey       = "capacity.cluster-autoscaler.kubernetes.io/memory"
	gpuCountKey     = "capacity.cluster-autoscaler.kubernetes.io/gpu-count"
	gpuTypeKey      = "capacity.cluster-autoscaler.kubernetes.io/gpu-type"
	diskCapacityKey = "capacity.cluster-autoscaler.kubernetes.io/ephemeral-disk"
	labelsKey       = "capacity.cluster-autoscaler.kubernetes.io/labels"
	taintsKey       = "capacity.cluster-autoscaler.kubernetes.io/taints"
	maxPodsKey      = "capacity.cluster-autoscaler.kubernetes.io/maxPods"
)

func NewCloudProvider(ctx context.Context, kubeClient client.Client) *CloudProvider {
	return &CloudProvider{
		kubeClient: kubeClient,
	}
}

type ClusterAPIInstanceType struct {
	cloudprovider.InstanceType

	MachineDeploymentName      string
	MachineDeploymentNamespace string
}

type CloudProvider struct {
	kubeClient                client.Client
	accessLock                sync.Mutex
	machineProvider           machine.Provider
	machineDeploymentProvider machinedeployment.Provider
}

func (c *CloudProvider) Create(ctx context.Context, nodeClaim *v1.NodeClaim) (*v1.NodeClaim, error) {
	return nil, nil
}

func (c *CloudProvider) Delete(ctx context.Context, nodeClaim *v1.NodeClaim) error {
	return nil
}

// Get returns a NodeClaim for the Machine object with the supplied provider ID, or nil if not found.
func (c *CloudProvider) Get(ctx context.Context, providerID string) (*v1.NodeClaim, error) {
	return nil, nil
}

// GetInstanceTypes enumerates the known Cluster API scalable resources to generate the list
// of possible instance types.
func (c *CloudProvider) GetInstanceTypes(ctx context.Context, nodePool *v1.NodePool) ([]*cloudprovider.InstanceType, error) {
	return nil, nil
}

func (c *CloudProvider) GetSupportedNodeClasses() []schema.GroupVersionKind {
	return []schema.GroupVersionKind{
		{
			Group:   "group",
			Version: "version",
			Kind:    "ClusterAPINodeClass",
		},
	}
}

// Return nothing since there's no cloud provider drift.
func (c *CloudProvider) IsDrifted(ctx context.Context, nodeClaim *v1.NodeClaim) (cloudprovider.DriftReason, error) {
	return "", nil
}

func (c *CloudProvider) List(ctx context.Context) ([]*v1.NodeClaim, error) {
	return nil, nil
}

func (c *CloudProvider) Name() string {
	return "clusterapi"
}
