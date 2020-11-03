// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Event is a report of an event somewhere in the cluster.  Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
 */
export class Event extends pulumi.CustomResource {
    /**
     * Get an existing Event resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Event {
        return new Event(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:core/v1:Event';

    /**
     * Returns true if the given object is an instance of Event.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Event {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Event.__pulumiType;
    }

    /**
     * What action was taken/failed regarding to the Regarding object.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"v1">;
    /**
     * The number of times this event has occurred.
     */
    public readonly count!: pulumi.Output<number>;
    /**
     * Time when this Event was first observed.
     */
    public readonly eventTime!: pulumi.Output<string>;
    /**
     * The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
     */
    public readonly firstTimestamp!: pulumi.Output<string>;
    /**
     * The object that this event is about.
     */
    public readonly involvedObject!: pulumi.Output<outputs.core.v1.ObjectReference>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"Event">;
    /**
     * The time at which the most recent occurrence of this event was recorded.
     */
    public readonly lastTimestamp!: pulumi.Output<string>;
    /**
     * A human-readable description of the status of this operation.
     */
    public readonly message!: pulumi.Output<string>;
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
     */
    public readonly reason!: pulumi.Output<string>;
    /**
     * Optional secondary object for more complex actions.
     */
    public readonly related!: pulumi.Output<outputs.core.v1.ObjectReference>;
    /**
     * Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
     */
    public readonly reportingComponent!: pulumi.Output<string>;
    /**
     * ID of the controller instance, e.g. `kubelet-xyzf`.
     */
    public readonly reportingInstance!: pulumi.Output<string>;
    /**
     * Data about the Event series this event represents or nil if it's a singleton Event.
     */
    public readonly series!: pulumi.Output<outputs.core.v1.EventSeries>;
    /**
     * The component reporting this event. Should be a short machine understandable string.
     */
    public readonly source!: pulumi.Output<outputs.core.v1.EventSource>;
    /**
     * Type of this event (Normal, Warning), new types could be added in the future
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Event resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EventArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.involvedObject === undefined) {
                throw new Error("Missing required property 'involvedObject'");
            }
            if (!args || args.metadata === undefined) {
                throw new Error("Missing required property 'metadata'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["apiVersion"] = "v1";
            inputs["count"] = args ? args.count : undefined;
            inputs["eventTime"] = args ? args.eventTime : undefined;
            inputs["firstTimestamp"] = args ? args.firstTimestamp : undefined;
            inputs["involvedObject"] = args ? args.involvedObject : undefined;
            inputs["kind"] = "Event";
            inputs["lastTimestamp"] = args ? args.lastTimestamp : undefined;
            inputs["message"] = args ? args.message : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["reason"] = args ? args.reason : undefined;
            inputs["related"] = args ? args.related : undefined;
            inputs["reportingComponent"] = args ? args.reportingComponent : undefined;
            inputs["reportingInstance"] = args ? args.reportingInstance : undefined;
            inputs["series"] = args ? args.series : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["type"] = args ? args.type : undefined;
        } else {
            inputs["action"] = undefined /*out*/;
            inputs["apiVersion"] = undefined /*out*/;
            inputs["count"] = undefined /*out*/;
            inputs["eventTime"] = undefined /*out*/;
            inputs["firstTimestamp"] = undefined /*out*/;
            inputs["involvedObject"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["lastTimestamp"] = undefined /*out*/;
            inputs["message"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["reason"] = undefined /*out*/;
            inputs["related"] = undefined /*out*/;
            inputs["reportingComponent"] = undefined /*out*/;
            inputs["reportingInstance"] = undefined /*out*/;
            inputs["series"] = undefined /*out*/;
            inputs["source"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:events.k8s.io/v1:Event" }, { type: "kubernetes:events.k8s.io/v1beta1:Event" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Event.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Event resource.
 */
export interface EventArgs {
    /**
     * What action was taken/failed regarding to the Regarding object.
     */
    readonly action?: pulumi.Input<string>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"v1">;
    /**
     * The number of times this event has occurred.
     */
    readonly count?: pulumi.Input<number>;
    /**
     * Time when this Event was first observed.
     */
    readonly eventTime?: pulumi.Input<string>;
    /**
     * The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
     */
    readonly firstTimestamp?: pulumi.Input<string>;
    /**
     * The object that this event is about.
     */
    readonly involvedObject: pulumi.Input<inputs.core.v1.ObjectReference>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"Event">;
    /**
     * The time at which the most recent occurrence of this event was recorded.
     */
    readonly lastTimestamp?: pulumi.Input<string>;
    /**
     * A human-readable description of the status of this operation.
     */
    readonly message?: pulumi.Input<string>;
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    readonly metadata: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
     */
    readonly reason?: pulumi.Input<string>;
    /**
     * Optional secondary object for more complex actions.
     */
    readonly related?: pulumi.Input<inputs.core.v1.ObjectReference>;
    /**
     * Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
     */
    readonly reportingComponent?: pulumi.Input<string>;
    /**
     * ID of the controller instance, e.g. `kubelet-xyzf`.
     */
    readonly reportingInstance?: pulumi.Input<string>;
    /**
     * Data about the Event series this event represents or nil if it's a singleton Event.
     */
    readonly series?: pulumi.Input<inputs.core.v1.EventSeries>;
    /**
     * The component reporting this event. Should be a short machine understandable string.
     */
    readonly source?: pulumi.Input<inputs.core.v1.EventSource>;
    /**
     * Type of this event (Normal, Warning), new types could be added in the future
     */
    readonly type?: pulumi.Input<string>;
}
