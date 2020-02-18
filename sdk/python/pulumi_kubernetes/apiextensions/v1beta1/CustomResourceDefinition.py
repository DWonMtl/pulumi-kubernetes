# *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
from typing import Optional

import pulumi
import pulumi.runtime
from pulumi import Input, ResourceOptions

from ... import tables, version


class CustomResourceDefinition(pulumi.CustomResource):
    """
    CustomResourceDefinition represents a resource that should be exposed on the API server.  Its
    name MUST be in the format <.spec.name>.<.spec.group>. Deprecated in v1.16, planned for removal
    in v1.19. Use apiextensions.k8s.io/v1 CustomResourceDefinition instead.
    """

    apiVersion: pulumi.Output[str]
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should
    convert recognized schemas to the latest internal value, and may reject unrecognized values.
    More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
    """

    kind: pulumi.Output[str]
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer
    this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
    info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
    """

    metadata: pulumi.Output[dict]
    

    spec: pulumi.Output[dict]
    """
    spec describes how the user wants the resources to appear
    """

    status: pulumi.Output[dict]
    """
    status indicates the actual state of the CustomResourceDefinition
    """

    def __init__(self, resource_name, opts=None, spec=None, metadata=None, __name__=None, __opts__=None):
        """
        Create a CustomResourceDefinition resource with the given unique name, arguments, and options.

        :param str resource_name: The _unique_ name of the resource.
        :param pulumi.ResourceOptions opts: A bag of options that control this resource's behavior.
        :param pulumi.Input[dict] spec: spec describes how the user wants the resources to appear
        :param pulumi.Input[dict] metadata: 
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apiVersion'] = 'apiextensions.k8s.io/v1beta1'
        __props__['kind'] = 'CustomResourceDefinition'
        if spec is None:
            raise TypeError('Missing required property spec')
        __props__['spec'] = spec
        __props__['metadata'] = metadata

        __props__['status'] = None

        parent = opts.parent if opts and opts.parent else None
        aliases = [
            pulumi.Alias(type_="kubernetes:apiextensions.k8s.io/v1:CustomResourceDefinition"),
        ]
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(
            version=version.get_version(),
            aliases=aliases,
        ))

        super(CustomResourceDefinition, self).__init__(
            "kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinition",
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get the state of an existing `CustomResourceDefinition` resource, as identified by `id`.
        The ID is of the form `[namespace]/[name]`; if `[namespace]` is omitted,
        then (per Kubernetes convention) the ID becomes `default/[name]`.

        Pulumi will keep track of this resource using `resource_name` as the Pulumi ID.

        :param str resource_name: _Unique_ name used to register this resource with Pulumi.
        :param pulumi.Input[str] id: An ID for the Kubernetes resource to retrieve.
               Takes the form `[namespace]/[name]` or `[name]`.
        :param Optional[pulumi.ResourceOptions] opts: A bag of options that control this
               resource's behavior.
        """
        opts = ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))
        return CustomResourceDefinition(resource_name, opts)

    def translate_output_property(self, prop: str) -> str:
        return tables._CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return tables._CASING_BACKWARD_TABLE.get(prop) or prop
