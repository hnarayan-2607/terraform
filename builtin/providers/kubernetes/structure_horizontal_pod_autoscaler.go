package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
	api "k8s.io/kubernetes/pkg/apis/autoscaling/v1"
)

func expandHorizontalPodAutoscalerSpec(in []interface{}) api.HorizontalPodAutoscalerSpec {
	if len(in) == 0 || in[0] == nil {
		return api.HorizontalPodAutoscalerSpec{}
	}
	spec := api.HorizontalPodAutoscalerSpec{}
	m := in[0].(map[string]interface{})
	if v, ok := m["max_replicas"]; ok {
		spec.MaxReplicas = int32(v.(int))
	}
	if v, ok := m["min_replicas"]; ok {
		spec.MinReplicas = ptrToInt32(v.(int32))
	}
	if v, ok := m["scale_target_ref"]; ok {
		spec.ScaleTargetRef = expandCrossVersionObjectReference(v.([]interface{}))
	}
	if v, ok := m["target_cpu_utilization_percentage"]; ok {
		spec.TargetCPUUtilizationPercentage = ptrToInt32(v.(int32))
	}

	return spec
}

func expandCrossVersionObjectReference(in []interface{}) api.CrossVersionObjectReference {
	if len(in) == 0 || in[0] == nil {
		return api.CrossVersionObjectReference{}
	}
	ref := api.CrossVersionObjectReference{}
	m := in[0].(map[string]interface{})

	if v, ok := m["api_version"]; ok {
		ref.APIVersion = v.(string)
	}
	if v, ok := m["kind"]; ok {
		ref.Kind = v.(string)
	}
	if v, ok := m["name"]; ok {
		ref.Name = v.(string)
	}
	return ref
}

func flattenHorizontalPodAutoscalerSpec(spec api.HorizontalPodAutoscalerSpec) []interface{} {
	m := make(map[string]interface{}, 0)
	m["max_replicas"] = spec.MaxReplicas
	if spec.MinReplicas != nil {
		m["min_replicas"] = *spec.MinReplicas
	}
	m["scale_target_ref"] = flattenCrossVersionObjectReference(spec.ScaleTargetRef)
	if spec.TargetCPUUtilizationPercentage != nil {
		m["target_cpu_utilization_percentage"] = *spec.TargetCPUUtilizationPercentage
	}
	return []interface{}{m}
}

func flattenCrossVersionObjectReference(ref api.CrossVersionObjectReference) []interface{} {
	m := make(map[string]interface{}, 0)
	if ref.APIVersion != "" {
		m["api_version"] = ref.APIVersion
	}
	if ref.Kind != "" {
		m["kind"] = ref.Kind
	}
	if ref.Name != "" {
		m["name"] = ref.Name
	}
	return []interface{}{m}
}

func patchHorizontalPodAutoscalerSpec(keyPrefix string, pathPrefix string, d *schema.ResourceData) PatchOperations {
	return nil
}
