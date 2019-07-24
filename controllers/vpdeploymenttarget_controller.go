/*
Copyright 2019 FinTech Studios, Inc.

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
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
)

// updateResource takes a k8s resource and a VP resource and merges them
func (r *VPDeploymentTargetReconciler) updateResource(req ctrl.Request, resource *ververicaplatformv1beta1.VPDeploymentTarget, depTarget *vpAPI.DeploymentTarget) error {
	ctx := context.Background()

	resource.Name = depTarget.Metadata.Name
	resource.Spec.Metadata = ververicaplatformv1beta1.VPDeploymentTargetMetadata{
		Name:            depTarget.Metadata.Name,
		Namespace:       depTarget.Metadata.Namespace,
		Id:              depTarget.Metadata.Id,
		CreatedAt:       &metav1.Time{Time: depTarget.Metadata.CreatedAt},
		ModifiedAt:      &metav1.Time{Time: depTarget.Metadata.ModifiedAt},
		ResourceVersion: depTarget.Metadata.ResourceVersion,
		Labels:          depTarget.Metadata.Labels,
		Annotations:     depTarget.Metadata.Annotations,
	}

	resource.Spec.Spec = ververicaplatformv1beta1.VPDeploymentTargetSpec{
		Kubernetes: ververicaplatformv1beta1.VPKubernetesTarget{
			Namespace: depTarget.Spec.Kubernetes.Namespace,
		},
	}

	if err := r.Update(ctx, resource); err != nil {
		return err
	}

	return nil
}

// getLogger creates a logger for the controller with the request name
func (r *VPDeploymentTargetReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpdeploymenttarget", req.NamespacedName)
}

// handleCreate creates VP resources
func (r *VPDeploymentTargetReconciler) handleCreate(req ctrl.Request, vpDepTarget ververicaplatformv1beta1.VPDeploymentTarget) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)
	namespace := utils.GetNamespaceOrDefault(&vpDepTarget.Spec.Metadata.Namespace)
	// TODO: load this from
	patchSet := []vpAPI.JsonPatchGeneric{
		{
			Op: "add",
			Path: "/a",
			Value: "",
		},
	}
	depTarget := vpAPI.DeploymentTarget{
		ApiVersion: "v1",
		Metadata: &vpAPI.DeploymentTargetMetadata{
			Name:        req.Name,
			Namespace:   vpDepTarget.Spec.Metadata.Namespace,
			Labels:      vpDepTarget.Spec.Metadata.Labels,
			Annotations: vpDepTarget.Spec.Metadata.Annotations,
		},
		Spec: &vpAPI.DeploymentTargetSpec{
			// Perhaps take this from the req as well?
			Kubernetes: &vpAPI.KubernetesTarget{Namespace: vpDepTarget.Spec.Spec.Kubernetes.Namespace},
			DeploymentPatchSet: patchSet,
		},
	}
	// create it
	_, err := r.VPAPIClient.
		DeploymentTargetsApi.
		CreateDeploymentTarget(ctx, namespace, depTarget)

	if err != nil {
		log.Error(err, "Error creating VP Deployment Target")
		return ctrl.Result{}, err
	}

	// TODO: the depTarget data is already in the res, but for some reason need to un-marshal it
	// 		 most likely a problem with the Swagger
	depTarget, _, err = r.VPAPIClient.DeploymentTargetsApi.GetDeploymentTarget(ctx, namespace, req.Name)
	if err != nil {
		return ctrl.Result{}, err
	}

	log.Info("Created depTarget", "depTarget", depTarget)

	// Now update the k8s resource and status as well
	if err := r.updateResource(req, &vpDepTarget, &depTarget); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// handleUpdate updates the k8s resource when it already exists in the VP
// updates are not supported on Deployment Targets in the VP API, so just need to mirror the latest state
func (r *VPDeploymentTargetReconciler) handleUpdate(req ctrl.Request, vpDepTarget ververicaplatformv1beta1.VPDeploymentTarget, depTarget vpAPI.DeploymentTarget) (ctrl.Result, error)  {
	// Now update the k8s resource
	err := r.updateResource(req, &vpDepTarget, &depTarget)
	return ctrl.Result{}, err
}


// handleDelete will ensure that the Ververica Platform namespace is also cleaned up
func (r *VPDeploymentTargetReconciler) handleDelete(req ctrl.Request, vpDepTarget ververicaplatformv1beta1.VPDeploymentTarget) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	// Let's make sure it's deleted from the ververica platform
	depTarget, _, err := r.VPAPIClient.DeploymentTargetsApi.DeleteDeploymentTarget(ctx, vpDepTarget.Spec.Metadata.Namespace, req.Name)

	if err != nil {
		// If it's already gone, great!
		// TODO: think about adding a wait time if the error
		//		 is about deployments still being attached to the dep target,
		// 		 as perhaps they're still in the deletion process
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	log.Info("Deleting Deployment Target", "name", depTarget.Metadata.Name)
	// Should happen instantaneously
	return ctrl.Result{}, nil
}

// VPDeploymentTargetReconciler reconciles a VPDeploymentTarget object
type VPDeploymentTargetReconciler struct {
	client.Client
	Log         logr.Logger
	VPAPIClient vpAPI.APIClient
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeploymenttargets,verbs=get;list;watch;create;update;patch;delete

// Reconcile tries to make the current state more like the desired state
func (r *VPDeploymentTargetReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	var vpDepTarget ververicaplatformv1beta1.VPDeploymentTarget
	// If it's gone, it's gone!
	if err := r.Get(ctx, req.NamespacedName, &vpDepTarget); err != nil {
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	if vpDepTarget.ObjectMeta.DeletionTimestamp.IsZero() {
		// Not being deleted, add the finalizer
		if utils.AddFinalizerToObjectMeta(&vpDepTarget.ObjectMeta) {
			log.Info("Adding Finalizer")
			if err := r.Update(ctx, &vpDepTarget); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// Being deleted
		log.Info("Deletion event", "name", req.Name)
		res, err := r.handleDelete(req, vpDepTarget)
		if utils.IsRequeueResponse(res, err) {
			return res, err
		}
		// otherwise, we're all good, just remove the finalizer
		if utils.RemoveFinalizerFromObjectMeta(&vpDepTarget.ObjectMeta) {
			if err := r.Update(ctx, &vpDepTarget); err != nil {
				return ctrl.Result{}, err
			}
		}

		return res, nil
	}

	namespace := utils.GetNamespaceOrDefault(&vpDepTarget.Spec.Metadata.Namespace)
	depTarget, _, err := r.VPAPIClient.DeploymentTargetsApi.GetDeploymentTarget(ctx, namespace, req.Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			// Not found, let's create it
			return r.handleCreate(req, vpDepTarget)
		}
		// Other error, not good!
		return ctrl.Result{}, err
	}

	log.Info("Update event")
	return r.handleUpdate(req, vpDepTarget, depTarget)
}

// SetupWithManager is a helper function to initial on manager boot
func (r *VPDeploymentTargetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ververicaplatformv1beta1.VPDeploymentTarget{}).
		Complete(r)
}
