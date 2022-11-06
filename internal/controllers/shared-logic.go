package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// this interface defines something
// that has methods to get the fields
// required by the logic function
type fielder interface {
	GetFoo() string
}

func logic(ctx context.Context, cli client.Client, obj fielder) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("demo", "foo", obj.GetFoo())
	// vPass , err := vps.Get(ctx, client, obj.GetFoo())
	return ctrl.Result{}, nil
}
