package main

import (
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	schema = runtime.NewScheme()
)

func main() {

	mgr, err := ctrl.NewManager(config.GetConfigOrDie(), ctrl.Options{
		Scheme: schema,
	})
	if err != nil {
		glog.Error("fail to create controller manager: %v", err)
	}

	if err = mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		glog.Error("fail to start manager: %v", err)
	}
}
