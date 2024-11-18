package v1alpha1

import(
	cachev2alpha1 "example.com/user/memcached/api/v2alpha1"
)
//ConvertTo 
func (src *Memached) ConvertTo(dstRaw converesion.Hub) error {
	dst := dstRaw.(*cachev2alpha1.Memcached)
	dst.Spec.ReplicaSize = src.Spec.Size
	dst.ObjectMeta = src.ObjectMeta
	dst.Status.Nodes = src.Status.Nodes
	return nil}

//ConvertFrom
func (dst *Memcached) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*cachev2alpha1.Memcached)
	dst.Spec.Size = src.Spec.ReplicaSize
	dst.ObjectMeta = src.ObjectMeta
	dst.Status.Nodes = src.Status.Nodes
	return nil
}