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
	kbkb "github.com/omakeno/kbkb/pkg"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	k8sv1beta1 "github.com/omakeno/kbkb-operator/api/v1beta1"
)

// KbkbReconciler reconciles a Kbkb object
type KbkbReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=k8s.omakenoyouna.net,resources=kbkbs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=k8s.omakenoyouna.net,resources=kbkbs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;delete;watch
// +kubebuilder:rbac:groups="",resources=nodes,verbs=get;list;watch

func (r *KbkbReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	reqLogger := r.Log.WithValues("kbkb", req.NamespacedName)

	reqLogger.Info("Reconciling")

	kbkbObj := &k8sv1beta1.Kbkb{}
	if err := r.Client.Get(ctx, req.NamespacedName, kbkbObj); err != nil {
		if errors.IsNotFound(err) {
			reqLogger.Info("kbkb not found. Ignore not found")
			return ctrl.Result{}, nil
		}
		reqLogger.Error(err, "failed to get kbkb")
		return ctrl.Result{}, err
	}
	kokeshi := *(kbkbObj.Spec.Kokeshi)

	podList := &corev1.PodList{}
	if err := r.Client.List(ctx, podList); err != nil {
		reqLogger.Error(err, "failed to get list of pods")
		return ctrl.Result{}, err
	}

	nodeList := &corev1.NodeList{}
	if err := r.Client.List(ctx, nodeList); err != nil {
		reqLogger.Error(err, "failed to get list of nodes")
		return ctrl.Result{}, err
	}

	kf := kbkb.BuildKbkbFieldFromList(podList, nodeList)
	erasablePods := kf.ErasableKbkbPodList(kokeshi)

	for _, kp := range erasablePods {
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: kp.ObjectMeta.Namespace,
				Name:      kp.ObjectMeta.Name,
			},
		}

		if err := r.Client.Delete(ctx, pod); err != nil {
			reqLogger.Error(err, "failed to delete pod", pod.ObjectMeta.Name)
		} else {
			reqLogger.Info("delete pod: ", pod.ObjectMeta.Name)
		}
	}

	return ctrl.Result{}, nil
}

func (r *KbkbReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&k8sv1beta1.Kbkb{}).
		For(&corev1.Pod{}).
		Complete(r)
}
