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
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	etcdbenchv1 "my.domain/guestbook/api/v1"
)

// WriterReconciler reconciles a Writer object
type WriterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=etcd-bench.my.domain,resources=writers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=etcd-bench.my.domain,resources=writers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=etcd-bench.my.domain,resources=writers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Writer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *WriterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	log := log.FromContext(ctx)

	var writer etcdbenchv1.Writer
	if err := r.Get(ctx, req.NamespacedName, &writer); err != nil {
		log.Error(err, "unable to fetch Writer")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if writer.Spec.Len == 0 {
		log.Info("writer length 0, ignoring")
		return ctrl.Result{}, nil
	}
	if len(writer.Spec.Data) != int(writer.Spec.Len) {
		writer.Spec.Data = strings.Repeat("a", int(writer.Spec.Len))
	}
	x := writer.Spec.Data[writer.Spec.Pos]
	x = ((x-97)+1)%26 + 97
	writer.Spec.Data = writer.Spec.Data[:writer.Spec.Pos] + string(x) + writer.Spec.Data[writer.Spec.Pos+1:]
	writer.Spec.Pos = (writer.Spec.Pos + 1) % writer.Spec.Len
	// writer.Spec.Data += string('Z')
	// writer.Spec.Len = int32(len(writer.Spec.Data))

	if err := r.Update(ctx, &writer); err != nil {
		log.Error(err, "unable to update Writer")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WriterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&etcdbenchv1.Writer{}).
		Complete(r)
}
