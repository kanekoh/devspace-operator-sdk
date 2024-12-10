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

package v2alpha1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	// appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Memcached Webhook", func() {

	Context("When creating Memcached under Defaulting Webhook", func() {
		It("Should fill in the default value if a required field is empty", func() {

			// TODO(user): Add your logic here

		})
	})

	Context("When creating Memcached under Validating Webhook", func() {
		It("Should deny if a required field is empty", func() {

			// TODO(user): Add your logic here

		})

		It("Should admit if all required fields are provided", func() {

			// TODO(user): Add your logic here

		})
	})

	Context("When creating Memcached under Conversion Webhook", func() {
		const resourceName = "test-resource"

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default",
		}
		memcached := &Memcached{}

		BeforeEach(func() {
			By("creating the custom resource for the Kind Memcached")
			err := k8sClient.Get(ctx, typeNamespacedName, memcached)
			if err != nil && errors.IsNotFound(err) {
				resource := &Memcached{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					Spec: MemcachedSpec{
						ReplicaSize: 1,
					},
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		It("Should get the converted version of Memcached", func() {
			resource := &Memcached{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			// TODO(user): Add your logic here
			Expect(resource.Spec.ReplicaSize).Should(Equal(int32(1)))

		})
	})

})
