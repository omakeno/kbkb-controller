/*


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

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	omakenoyounanetv1alpha1 "github.com/omakeno/kbkb-operator/api/v1alpha1"
)

// KbkbReconciler reconciles a Kbkb object
type KbkbReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=omakenoyouna.net.omakenoyouna.net,resources=kbkbs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=omakenoyouna.net.omakenoyouna.net,resources=kbkbs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;delete

func (r *KbkbReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	reqLogger := r.Log.WithValues("kbkb", req.NamespacedName)

	// your logic here

	reqLogger.Info("Reconciling")
	ctx := context.Background()

	kbkb := &omakenoyounanetv1alpha1.Kbkb{}
	if err := r.Client.Get(ctx, req.NamespacedName, kbkb); err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("kbkb not found. Ignore not found")
			return ctrl.Result{}, nil
		}
		reqLogger.Error(err, "failed to get kbkb")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *KbkbReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&omakenoyounanetv1alpha1.Kbkb{}).
		Owns(&corev1.Pod{}).
		Complete(r)
}
