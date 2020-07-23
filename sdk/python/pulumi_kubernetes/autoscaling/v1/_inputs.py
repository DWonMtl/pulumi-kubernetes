# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from ... import meta as _meta

__all__ = [
    'CrossVersionObjectReferenceArgs',
    'HorizontalPodAutoscalerArgs',
    'HorizontalPodAutoscalerSpecArgs',
    'HorizontalPodAutoscalerStatusArgs',
]

@pulumi.input_type
class CrossVersionObjectReferenceArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[str],
                 name: pulumi.Input[str],
                 api_version: Optional[pulumi.Input[str]] = None):
        """
        CrossVersionObjectReference contains enough information to let you identify the referred resource.
        :param pulumi.Input[str] kind: Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
        :param pulumi.Input[str] name: Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
        :param pulumi.Input[str] api_version: API version of the referent
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "apiVersion", api_version)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        """
        Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
        """
        ...

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        API version of the referent
        """
        ...

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class HorizontalPodAutoscalerArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']] = None,
                 status: Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']] = None):
        """
        configuration of a horizontal pod autoscaler.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['HorizontalPodAutoscalerSpecArgs'] spec: behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
        :param pulumi.Input['HorizontalPodAutoscalerStatusArgs'] status: current information about the autoscaler.
        """
        pulumi.set(__self__, "apiVersion", 'autoscaling/v1')
        pulumi.set(__self__, "kind", 'HorizontalPodAutoscaler')
        pulumi.set(__self__, "metadata", metadata)
        pulumi.set(__self__, "spec", spec)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        ...

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        ...

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        """
        Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        ...

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        ...

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']]:
        """
        behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
        """
        ...

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']]):
        ...

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']]:
        """
        current information about the autoscaler.
        """
        ...

    @status.setter
    def status(self, value: Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']]):
        ...


@pulumi.input_type
class HorizontalPodAutoscalerSpecArgs:
    def __init__(__self__, *,
                 max_replicas: pulumi.Input[float],
                 scale_target_ref: pulumi.Input['CrossVersionObjectReferenceArgs'],
                 min_replicas: Optional[pulumi.Input[float]] = None,
                 target_cpu_utilization_percentage: Optional[pulumi.Input[float]] = None):
        """
        specification of a horizontal pod autoscaler.
        :param pulumi.Input[float] max_replicas: upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
        :param pulumi.Input['CrossVersionObjectReferenceArgs'] scale_target_ref: reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
        :param pulumi.Input[float] min_replicas: minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        :param pulumi.Input[float] target_cpu_utilization_percentage: target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
        """
        pulumi.set(__self__, "maxReplicas", max_replicas)
        pulumi.set(__self__, "scaleTargetRef", scale_target_ref)
        pulumi.set(__self__, "minReplicas", min_replicas)
        pulumi.set(__self__, "targetCPUUtilizationPercentage", target_cpu_utilization_percentage)

    @property
    @pulumi.getter(name="maxReplicas")
    def max_replicas(self) -> pulumi.Input[float]:
        """
        upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
        """
        ...

    @max_replicas.setter
    def max_replicas(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="scaleTargetRef")
    def scale_target_ref(self) -> pulumi.Input['CrossVersionObjectReferenceArgs']:
        """
        reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
        """
        ...

    @scale_target_ref.setter
    def scale_target_ref(self, value: pulumi.Input['CrossVersionObjectReferenceArgs']):
        ...

    @property
    @pulumi.getter(name="minReplicas")
    def min_replicas(self) -> Optional[pulumi.Input[float]]:
        """
        minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        """
        ...

    @min_replicas.setter
    def min_replicas(self, value: Optional[pulumi.Input[float]]):
        ...

    @property
    @pulumi.getter(name="targetCPUUtilizationPercentage")
    def target_cpu_utilization_percentage(self) -> Optional[pulumi.Input[float]]:
        """
        target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
        """
        ...

    @target_cpu_utilization_percentage.setter
    def target_cpu_utilization_percentage(self, value: Optional[pulumi.Input[float]]):
        ...


@pulumi.input_type
class HorizontalPodAutoscalerStatusArgs:
    def __init__(__self__, *,
                 current_replicas: pulumi.Input[float],
                 desired_replicas: pulumi.Input[float],
                 current_cpu_utilization_percentage: Optional[pulumi.Input[float]] = None,
                 last_scale_time: Optional[pulumi.Input[str]] = None,
                 observed_generation: Optional[pulumi.Input[float]] = None):
        """
        current status of a horizontal pod autoscaler
        :param pulumi.Input[float] current_replicas: current number of replicas of pods managed by this autoscaler.
        :param pulumi.Input[float] desired_replicas: desired number of replicas of pods managed by this autoscaler.
        :param pulumi.Input[float] current_cpu_utilization_percentage: current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
        :param pulumi.Input[str] last_scale_time: last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
        :param pulumi.Input[float] observed_generation: most recent generation observed by this autoscaler.
        """
        pulumi.set(__self__, "currentReplicas", current_replicas)
        pulumi.set(__self__, "desiredReplicas", desired_replicas)
        pulumi.set(__self__, "currentCPUUtilizationPercentage", current_cpu_utilization_percentage)
        pulumi.set(__self__, "lastScaleTime", last_scale_time)
        pulumi.set(__self__, "observedGeneration", observed_generation)

    @property
    @pulumi.getter(name="currentReplicas")
    def current_replicas(self) -> pulumi.Input[float]:
        """
        current number of replicas of pods managed by this autoscaler.
        """
        ...

    @current_replicas.setter
    def current_replicas(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="desiredReplicas")
    def desired_replicas(self) -> pulumi.Input[float]:
        """
        desired number of replicas of pods managed by this autoscaler.
        """
        ...

    @desired_replicas.setter
    def desired_replicas(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="currentCPUUtilizationPercentage")
    def current_cpu_utilization_percentage(self) -> Optional[pulumi.Input[float]]:
        """
        current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
        """
        ...

    @current_cpu_utilization_percentage.setter
    def current_cpu_utilization_percentage(self, value: Optional[pulumi.Input[float]]):
        ...

    @property
    @pulumi.getter(name="lastScaleTime")
    def last_scale_time(self) -> Optional[pulumi.Input[str]]:
        """
        last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
        """
        ...

    @last_scale_time.setter
    def last_scale_time(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="observedGeneration")
    def observed_generation(self) -> Optional[pulumi.Input[float]]:
        """
        most recent generation observed by this autoscaler.
        """
        ...

    @observed_generation.setter
    def observed_generation(self, value: Optional[pulumi.Input[float]]):
        ...

