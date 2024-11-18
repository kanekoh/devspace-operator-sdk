package v1alpha1

import (
	cachev2alpha1 "example.com/user/memcached/api/v2alpha1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo
func (src *Memcached) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*cachev2alpha1.Memcached)
	dst.Spec.ReplicaSize = src.Spec.Size
	dst.ObjectMeta = src.ObjectMeta
	return nil
}

// ConvertFrom
func (dst *Memcached) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*cachev2alpha1.Memcached)
	dst.Spec.Size = src.Spec.ReplicaSize
	dst.ObjectMeta = src.ObjectMeta
	return nil
}
