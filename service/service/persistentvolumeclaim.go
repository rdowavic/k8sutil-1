package service

import (
	. "bitbucket.org/welovetravel/xops/service/lint"
	v1 "k8s.io/api/core/v1"
)

func PersistentVolumeClaimRules(resource *YamlDerivedKubernetesResource) []*Rule {
	if _, isPersistentVolumeClaim := resource.Resource.(*v1.PersistentVolumeClaim); !isPersistentVolumeClaim {
		return nil
	}
	return nil
}
