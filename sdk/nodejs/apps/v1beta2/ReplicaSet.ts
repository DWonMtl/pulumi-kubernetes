// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { core } from "../..";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import { getVersion } from "../../version";

    /**
     * ReplicaSet ensures that a specified number of pod replicas are running at any given time.
     * 
     * @deprecated apps/v1beta2/ReplicaSet is deprecated by apps/v1/ReplicaSet and not supported by
     * Kubernetes v1.16+ clusters.
     */
    export class ReplicaSet extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"apps/v1beta2">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"ReplicaSet">;

      /**
       * If the Labels of a ReplicaSet are empty, they are defaulted to be the same as the Pod(s)
       * that the ReplicaSet manages. Standard object's metadata. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputs.meta.v1.ObjectMeta>;

      /**
       * Spec defines the specification of the desired behavior of the ReplicaSet. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
       */
      public readonly spec: pulumi.Output<outputs.apps.v1beta2.ReplicaSetSpec>;

      /**
       * Status is the most recently observed status of the ReplicaSet. This data may be out of date
       * by some window of time. Populated by the system. Read-only. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
       */
      public readonly status: pulumi.Output<outputs.apps.v1beta2.ReplicaSetStatus>;

      /**
       * Get the state of an existing `ReplicaSet` resource, as identified by `id`.
       * The ID is of the form `[namespace]/<name>`; if `namespace` is omitted, then (per
       * Kubernetes convention) the ID becomes `default/<name>`.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form `[namespace]/<name>`.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReplicaSet {
          return new ReplicaSet(name, undefined, { ...opts, id: id });
      }

      /** @internal */
      private static readonly __pulumiType = "kubernetes:apps/v1beta2:ReplicaSet";

      /**
       * Returns true if the given object is an instance of ReplicaSet.  This is designed to work even
       * when multiple copies of the Pulumi SDK have been loaded into the same process.
       */
      public static isInstance(obj: any): obj is ReplicaSet {
          if (obj === undefined || obj === null) {
              return false;
          }

          return obj["__pulumiType"] === ReplicaSet.__pulumiType;
      }

      /**
       * Create a apps.v1beta2.ReplicaSet resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputs.apps.v1beta2.ReplicaSet, opts?: pulumi.CustomResourceOptions) {
          const props: pulumi.Inputs = {};

          props["apiVersion"] = "apps/v1beta2";
          props["kind"] = "ReplicaSet";
          props["metadata"] = args?.metadata;
          props["spec"] = args?.spec;

          props["status"] = undefined;

          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }

          opts = pulumi.mergeOptions(opts, {
              aliases: [
                  { type: "kubernetes:apps/v1:ReplicaSet" },
                  { type: "kubernetes:extensions/v1beta1:ReplicaSet" },
              ],
          });
          super(ReplicaSet.__pulumiType, name, props, opts);
      }
    }
