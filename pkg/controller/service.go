package controller

import (
	"context"

	"github.com/golang/glog"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ServiceController struct {
	client.Client
	// must?
	//Schema *runtime.Scheme
}

func (r *ServiceController) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	glog.Error("test1111111111111")
	return ctrl.Result{}, nil
}

func (r *ServiceController) RegisterController(mgr ctrl.Manager) error {
	return nil
}
