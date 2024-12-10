/*
Copyright 2024 redhat.

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

package controller

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	cachev1alpha1 "example.com/user/memcached/api/v1alpha1"
)

var _ = Describe("Memcached Controller", func() {
	Context("When reconciling a resource", func() {
		const resourceName = "test-resource"

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default",
		}
		memcached := &cachev1alpha1.Memcached{}

		BeforeEach(func() {
			By("creating the custom resource for the Kind Memcached")
			err := k8sClient.Get(ctx, typeNamespacedName, memcached)
			if err != nil && errors.IsNotFound(err) {
				resource := &cachev1alpha1.Memcached{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					Spec: cachev1alpha1.MemcachedSpec{
						Size: 1,
					},
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		AfterEach(func() {
			// TODO(user): Cleanup logic after each test, like removing the resource instance.
			resource := &cachev1alpha1.Memcached{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource instance Memcached")
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		})
		It("should successfully reconcile the resource", func() {
			By("Reconciling the created resource")
			controllerReconciler := &MemcachedReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())

			// Deployment が作成されていることを確認
			By("Checking if Deployment was successfully created in the reconciliation")
			Eventually(func() error {
				found := &appsv1.Deployment{}
				return k8sClient.Get(ctx, typeNamespacedName, found)
			}, time.Minute, time.Second).Should(Succeed())
		})
	})

	Context("When defining a resource", func() {
		const resourceName = "test-resource"

		ctx := context.Background()

		It("should fail when size exceeds the maximum value", func() {
			resource := &cachev1alpha1.Memcached{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resourceName,
					Namespace: "default",
				},
				Spec: cachev1alpha1.MemcachedSpec{
					Size: 6,
				},
			}
			err := k8sClient.Create(ctx, resource)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("spec.size in body should be less than or equal to 5"))
		})

		It("should fail when size is less than the minimum value", func() {
			resource := &cachev1alpha1.Memcached{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resourceName,
					Namespace: "default",
				},
				Spec: cachev1alpha1.MemcachedSpec{
					Size: 0,
				},
			}
			err := k8sClient.Create(ctx, resource)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("spec.size in body should be greater than or equal to 1"))
		})

	})
})
