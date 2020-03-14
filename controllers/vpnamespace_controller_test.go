package controllers

import (
	"context"
	"time"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1/converters"
	platformApiClient "github.com/fintechstudios/ververica-platform-k8s-operator/internal/vvp/platform-api-client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func timeMustParse(layout, value string) *time.Time {
	date, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}

	return &date
}

var _ = Describe("VpNamespace Controller", func() {
	var reconciler VpNamespaceReconciler

	BeforeEach(func() {
		platformClient := &platformApiClient.APIClient{}

		reconciler = VpNamespaceReconciler{
			Client:            k8sClient,
			Log:               logger,
			PlatformAPIClient: platformClient,
		}
	})

	Describe("updateResource", func() {
		var (
			key              types.NamespacedName
			created, fetched *v1beta1.VpNamespace
		)

		BeforeEach(func() {
			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &v1beta1.VpNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1beta1.VpNamespaceSpec{
					RoleBindings: []v1beta1.NamespaceRoleBinding{
						{
							Members: []string{"system:authenticated"},
							Role:    "owner",
						},
						{
							Members: []string{"austin@fintechstudios.com"},
							Role:    "viewer",
						},
					},
				},
			}
			Expect(k8sClient.Create(context.TODO(), created)).To(Succeed())
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(context.TODO(), created)).To(Succeed())
		})

		It("should update a k8s vp namespace with a Platform namespace", func() {
			phase := "LIFECYCLE_PHASE_ACTIVE"
			namespace := &platformApiClient.Namespace{
				CreateTime:     timeMustParse(time.RFC3339, "2019-10-18T14:27:58.328Z"),
				LifecyclePhase: &phase,
				Name:           "foo",
				RoleBindings: []platformApiClient.RoleBinding{
					{
						Members: []string{"system:authenticated"},
						Role:    "owner",
					},
					{
						Members: []string{"austin@fintechstudios.com"},
						Role:    "viewer",
					},
				},
			}

			Expect(reconciler.updateResource(created, namespace)).To(Succeed())

			fetched = &v1beta1.VpNamespace{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			updatedPhase, err := converters.NamespaceLifecyclePhaseFromNative(fetched.Status.LifecyclePhase)
			Expect(err).To(BeNil())
			Expect(updatedPhase).To(Equal(phase))
		})
	})
})
