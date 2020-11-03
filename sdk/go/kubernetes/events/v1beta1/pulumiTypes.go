// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type EventType struct {
	// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at most 128 characters.
	Action *string `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedCount *int `pulumi:"deprecatedCount"`
	// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedFirstTimestamp *string `pulumi:"deprecatedFirstTimestamp"`
	// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedLastTimestamp *string `pulumi:"deprecatedLastTimestamp"`
	// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedSource *corev1.EventSource `pulumi:"deprecatedSource"`
	// eventTime is the time when this Event was first observed. It is required.
	EventTime string `pulumi:"eventTime"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string           `pulumi:"kind"`
	Metadata metav1.ObjectMeta `pulumi:"metadata"`
	// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note *string `pulumi:"note"`
	// reason is why the action was taken. It is human-readable. This field can have at most 128 characters.
	Reason *string `pulumi:"reason"`
	// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding *corev1.ObjectReference `pulumi:"regarding"`
	// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related *corev1.ObjectReference `pulumi:"related"`
	// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
	ReportingController *string `pulumi:"reportingController"`
	// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
	ReportingInstance *string `pulumi:"reportingInstance"`
	// series is data about the Event series this event represents or nil if it's a singleton Event.
	Series *EventSeries `pulumi:"series"`
	// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.
	Type *string `pulumi:"type"`
}

// EventTypeInput is an input type that accepts EventTypeArgs and EventTypeOutput values.
// You can construct a concrete instance of `EventTypeInput` via:
//
//          EventTypeArgs{...}
type EventTypeInput interface {
	pulumi.Input

	ToEventTypeOutput() EventTypeOutput
	ToEventTypeOutputWithContext(context.Context) EventTypeOutput
}

// Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type EventTypeArgs struct {
	// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at most 128 characters.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedCount pulumi.IntPtrInput `pulumi:"deprecatedCount"`
	// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedFirstTimestamp pulumi.StringPtrInput `pulumi:"deprecatedFirstTimestamp"`
	// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedLastTimestamp pulumi.StringPtrInput `pulumi:"deprecatedLastTimestamp"`
	// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedSource corev1.EventSourcePtrInput `pulumi:"deprecatedSource"`
	// eventTime is the time when this Event was first observed. It is required.
	EventTime pulumi.StringInput `pulumi:"eventTime"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput  `pulumi:"kind"`
	Metadata metav1.ObjectMetaInput `pulumi:"metadata"`
	// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note pulumi.StringPtrInput `pulumi:"note"`
	// reason is why the action was taken. It is human-readable. This field can have at most 128 characters.
	Reason pulumi.StringPtrInput `pulumi:"reason"`
	// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding corev1.ObjectReferencePtrInput `pulumi:"regarding"`
	// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related corev1.ObjectReferencePtrInput `pulumi:"related"`
	// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
	ReportingController pulumi.StringPtrInput `pulumi:"reportingController"`
	// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
	ReportingInstance pulumi.StringPtrInput `pulumi:"reportingInstance"`
	// series is data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesPtrInput `pulumi:"series"`
	// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (EventTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventType)(nil)).Elem()
}

func (i EventTypeArgs) ToEventTypeOutput() EventTypeOutput {
	return i.ToEventTypeOutputWithContext(context.Background())
}

func (i EventTypeArgs) ToEventTypeOutputWithContext(ctx context.Context) EventTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTypeOutput)
}

// EventTypeArrayInput is an input type that accepts EventTypeArray and EventTypeArrayOutput values.
// You can construct a concrete instance of `EventTypeArrayInput` via:
//
//          EventTypeArray{ EventTypeArgs{...} }
type EventTypeArrayInput interface {
	pulumi.Input

	ToEventTypeArrayOutput() EventTypeArrayOutput
	ToEventTypeArrayOutputWithContext(context.Context) EventTypeArrayOutput
}

type EventTypeArray []EventTypeInput

func (EventTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventType)(nil)).Elem()
}

func (i EventTypeArray) ToEventTypeArrayOutput() EventTypeArrayOutput {
	return i.ToEventTypeArrayOutputWithContext(context.Background())
}

func (i EventTypeArray) ToEventTypeArrayOutputWithContext(ctx context.Context) EventTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTypeArrayOutput)
}

// Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type EventTypeOutput struct{ *pulumi.OutputState }

func (EventTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventType)(nil)).Elem()
}

func (o EventTypeOutput) ToEventTypeOutput() EventTypeOutput {
	return o
}

func (o EventTypeOutput) ToEventTypeOutputWithContext(ctx context.Context) EventTypeOutput {
	return o
}

// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at most 128 characters.
func (o EventTypeOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventType) *string { return v.Action }).(pulumi.StringPtrOutput)
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o EventTypeOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventType) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
func (o EventTypeOutput) DeprecatedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EventType) *int { return v.DeprecatedCount }).(pulumi.IntPtrOutput)
}

// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
func (o EventTypeOutput) DeprecatedFirstTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventType) *string { return v.DeprecatedFirstTimestamp }).(pulumi.StringPtrOutput)
}

// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
func (o EventTypeOutput) DeprecatedLastTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventType) *string { return v.DeprecatedLastTimestamp }).(pulumi.StringPtrOutput)
}

// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
func (o EventTypeOutput) DeprecatedSource() corev1.EventSourcePtrOutput {
	return o.ApplyT(func(v EventType) *corev1.EventSource { return v.DeprecatedSource }).(corev1.EventSourcePtrOutput)
}

// eventTime is the time when this Event was first observed. It is required.
func (o EventTypeOutput) EventTime() pulumi.StringOutput {
	return o.ApplyT(func(v EventType) string { return v.EventTime }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o EventTypeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventType) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o EventTypeOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v EventType) metav1.ObjectMeta { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
func (o EventTypeOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventType) *string { return v.Note }).(pulumi.StringPtrOutput)
}

// reason is why the action was taken. It is human-readable. This field can have at most 128 characters.
func (o EventTypeOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventType) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
func (o EventTypeOutput) Regarding() corev1.ObjectReferencePtrOutput {
	return o.ApplyT(func(v EventType) *corev1.ObjectReference { return v.Regarding }).(corev1.ObjectReferencePtrOutput)
}

// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
func (o EventTypeOutput) Related() corev1.ObjectReferencePtrOutput {
	return o.ApplyT(func(v EventType) *corev1.ObjectReference { return v.Related }).(corev1.ObjectReferencePtrOutput)
}

// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
func (o EventTypeOutput) ReportingController() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventType) *string { return v.ReportingController }).(pulumi.StringPtrOutput)
}

// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
func (o EventTypeOutput) ReportingInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventType) *string { return v.ReportingInstance }).(pulumi.StringPtrOutput)
}

// series is data about the Event series this event represents or nil if it's a singleton Event.
func (o EventTypeOutput) Series() EventSeriesPtrOutput {
	return o.ApplyT(func(v EventType) *EventSeries { return v.Series }).(EventSeriesPtrOutput)
}

// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.
func (o EventTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type EventTypeArrayOutput struct{ *pulumi.OutputState }

func (EventTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventType)(nil)).Elem()
}

func (o EventTypeArrayOutput) ToEventTypeArrayOutput() EventTypeArrayOutput {
	return o
}

func (o EventTypeArrayOutput) ToEventTypeArrayOutputWithContext(ctx context.Context) EventTypeArrayOutput {
	return o
}

func (o EventTypeArrayOutput) Index(i pulumi.IntInput) EventTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventType {
		return vs[0].([]EventType)[vs[1].(int)]
	}).(EventTypeOutput)
}

// EventList is a list of Event objects.
type EventListType struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is a list of schema objects.
	Items []EventType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// EventListTypeInput is an input type that accepts EventListTypeArgs and EventListTypeOutput values.
// You can construct a concrete instance of `EventListTypeInput` via:
//
//          EventListTypeArgs{...}
type EventListTypeInput interface {
	pulumi.Input

	ToEventListTypeOutput() EventListTypeOutput
	ToEventListTypeOutputWithContext(context.Context) EventListTypeOutput
}

// EventList is a list of Event objects.
type EventListTypeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// items is a list of schema objects.
	Items EventTypeArrayInput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput `pulumi:"metadata"`
}

func (EventListTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventListType)(nil)).Elem()
}

func (i EventListTypeArgs) ToEventListTypeOutput() EventListTypeOutput {
	return i.ToEventListTypeOutputWithContext(context.Background())
}

func (i EventListTypeArgs) ToEventListTypeOutputWithContext(ctx context.Context) EventListTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventListTypeOutput)
}

// EventList is a list of Event objects.
type EventListTypeOutput struct{ *pulumi.OutputState }

func (EventListTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventListType)(nil)).Elem()
}

func (o EventListTypeOutput) ToEventListTypeOutput() EventListTypeOutput {
	return o
}

func (o EventListTypeOutput) ToEventListTypeOutputWithContext(ctx context.Context) EventListTypeOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o EventListTypeOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventListType) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// items is a list of schema objects.
func (o EventListTypeOutput) Items() EventTypeArrayOutput {
	return o.ApplyT(func(v EventListType) []EventType { return v.Items }).(EventTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o EventListTypeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventListType) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o EventListTypeOutput) Metadata() metav1.ListMetaPtrOutput {
	return o.ApplyT(func(v EventListType) *metav1.ListMeta { return v.Metadata }).(metav1.ListMetaPtrOutput)
}

// EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
type EventSeries struct {
	// count is the number of occurrences in this series up to the last heartbeat time.
	Count int `pulumi:"count"`
	// lastObservedTime is the time when last Event from the series was seen before last heartbeat.
	LastObservedTime string `pulumi:"lastObservedTime"`
	// Information whether this series is ongoing or finished. Deprecated. Planned removal for 1.18
	State *string `pulumi:"state"`
}

// EventSeriesInput is an input type that accepts EventSeriesArgs and EventSeriesOutput values.
// You can construct a concrete instance of `EventSeriesInput` via:
//
//          EventSeriesArgs{...}
type EventSeriesInput interface {
	pulumi.Input

	ToEventSeriesOutput() EventSeriesOutput
	ToEventSeriesOutputWithContext(context.Context) EventSeriesOutput
}

// EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
type EventSeriesArgs struct {
	// count is the number of occurrences in this series up to the last heartbeat time.
	Count pulumi.IntInput `pulumi:"count"`
	// lastObservedTime is the time when last Event from the series was seen before last heartbeat.
	LastObservedTime pulumi.StringInput `pulumi:"lastObservedTime"`
	// Information whether this series is ongoing or finished. Deprecated. Planned removal for 1.18
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (EventSeriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSeries)(nil)).Elem()
}

func (i EventSeriesArgs) ToEventSeriesOutput() EventSeriesOutput {
	return i.ToEventSeriesOutputWithContext(context.Background())
}

func (i EventSeriesArgs) ToEventSeriesOutputWithContext(ctx context.Context) EventSeriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSeriesOutput)
}

func (i EventSeriesArgs) ToEventSeriesPtrOutput() EventSeriesPtrOutput {
	return i.ToEventSeriesPtrOutputWithContext(context.Background())
}

func (i EventSeriesArgs) ToEventSeriesPtrOutputWithContext(ctx context.Context) EventSeriesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSeriesOutput).ToEventSeriesPtrOutputWithContext(ctx)
}

// EventSeriesPtrInput is an input type that accepts EventSeriesArgs, EventSeriesPtr and EventSeriesPtrOutput values.
// You can construct a concrete instance of `EventSeriesPtrInput` via:
//
//          EventSeriesArgs{...}
//
//  or:
//
//          nil
type EventSeriesPtrInput interface {
	pulumi.Input

	ToEventSeriesPtrOutput() EventSeriesPtrOutput
	ToEventSeriesPtrOutputWithContext(context.Context) EventSeriesPtrOutput
}

type eventSeriesPtrType EventSeriesArgs

func EventSeriesPtr(v *EventSeriesArgs) EventSeriesPtrInput {
	return (*eventSeriesPtrType)(v)
}

func (*eventSeriesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSeries)(nil)).Elem()
}

func (i *eventSeriesPtrType) ToEventSeriesPtrOutput() EventSeriesPtrOutput {
	return i.ToEventSeriesPtrOutputWithContext(context.Background())
}

func (i *eventSeriesPtrType) ToEventSeriesPtrOutputWithContext(ctx context.Context) EventSeriesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSeriesPtrOutput)
}

// EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
type EventSeriesOutput struct{ *pulumi.OutputState }

func (EventSeriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSeries)(nil)).Elem()
}

func (o EventSeriesOutput) ToEventSeriesOutput() EventSeriesOutput {
	return o
}

func (o EventSeriesOutput) ToEventSeriesOutputWithContext(ctx context.Context) EventSeriesOutput {
	return o
}

func (o EventSeriesOutput) ToEventSeriesPtrOutput() EventSeriesPtrOutput {
	return o.ToEventSeriesPtrOutputWithContext(context.Background())
}

func (o EventSeriesOutput) ToEventSeriesPtrOutputWithContext(ctx context.Context) EventSeriesPtrOutput {
	return o.ApplyT(func(v EventSeries) *EventSeries {
		return &v
	}).(EventSeriesPtrOutput)
}

// count is the number of occurrences in this series up to the last heartbeat time.
func (o EventSeriesOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v EventSeries) int { return v.Count }).(pulumi.IntOutput)
}

// lastObservedTime is the time when last Event from the series was seen before last heartbeat.
func (o EventSeriesOutput) LastObservedTime() pulumi.StringOutput {
	return o.ApplyT(func(v EventSeries) string { return v.LastObservedTime }).(pulumi.StringOutput)
}

// Information whether this series is ongoing or finished. Deprecated. Planned removal for 1.18
func (o EventSeriesOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSeries) *string { return v.State }).(pulumi.StringPtrOutput)
}

type EventSeriesPtrOutput struct{ *pulumi.OutputState }

func (EventSeriesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSeries)(nil)).Elem()
}

func (o EventSeriesPtrOutput) ToEventSeriesPtrOutput() EventSeriesPtrOutput {
	return o
}

func (o EventSeriesPtrOutput) ToEventSeriesPtrOutputWithContext(ctx context.Context) EventSeriesPtrOutput {
	return o
}

func (o EventSeriesPtrOutput) Elem() EventSeriesOutput {
	return o.ApplyT(func(v *EventSeries) EventSeries { return *v }).(EventSeriesOutput)
}

// count is the number of occurrences in this series up to the last heartbeat time.
func (o EventSeriesPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EventSeries) *int {
		if v == nil {
			return nil
		}
		return &v.Count
	}).(pulumi.IntPtrOutput)
}

// lastObservedTime is the time when last Event from the series was seen before last heartbeat.
func (o EventSeriesPtrOutput) LastObservedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSeries) *string {
		if v == nil {
			return nil
		}
		return &v.LastObservedTime
	}).(pulumi.StringPtrOutput)
}

// Information whether this series is ongoing or finished. Deprecated. Planned removal for 1.18
func (o EventSeriesPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSeries) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EventTypeOutput{})
	pulumi.RegisterOutputType(EventTypeArrayOutput{})
	pulumi.RegisterOutputType(EventListTypeOutput{})
	pulumi.RegisterOutputType(EventSeriesOutput{})
	pulumi.RegisterOutputType(EventSeriesPtrOutput{})
}
