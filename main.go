package main

import (
	"sigs.k8s.io/controller-runtime/pkg/log"

	kwok "sigs.k8s.io/karpenter/kwok/cloudprovider"
	"sigs.k8s.io/karpenter/pkg/controllers"
	"sigs.k8s.io/karpenter/pkg/operator"
	"sigs.k8s.io/karpenter/pkg/webhooks"
)

func main() {
	ctx, op := operator.NewOperator()
	instanceTypes, err := kwok.ConstructInstanceTypes()
	if err != nil {
		log.FromContext(ctx).Error(err, "failed constructing instance types")
	}

	cloudProvider := kwok.NewCloudProvider(ctx, op.GetClient(), instanceTypes)
	op.
		WithWebhooks(ctx, webhooks.NewWebhooks()...).
		WithControllers(ctx, controllers.NewControllers(
			op.Manager,
			op.Clock,
			op.GetClient(),
			op.EventRecorder,
			cloudProvider,
		)...).Start(ctx, cloudProvider)
}
