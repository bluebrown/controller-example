/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "example.io/demo/api/v1"
)

// InternalReconciler reconciles a Internal object
type InternalReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.example.io,resources=internals,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.example.io,resources=internals/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.example.io,resources=internals/finalizers,verbs=update

func (r *InternalReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	obj := demov1.Internal{}
	_ = r.Get(ctx, req.NamespacedName, &obj)
	return logic(ctx, r.Client, obj)

}

func (r *InternalReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.Internal{}).
		Complete(r)
}
