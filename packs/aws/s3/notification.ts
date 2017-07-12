// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

export class Notification extends lumi.NamedResource implements NotificationArgs {
    public readonly bucket: string;
    public readonly lambdaFunction?: { events: string[], filterPrefix?: string, filterSuffix?: string, _id?: string, lambdaFunctionArn?: string }[];
    public readonly queue?: { events: string[], filterPrefix?: string, filterSuffix?: string, _id?: string, queueArn: string }[];
    public readonly topic?: { events: string[], filterPrefix?: string, filterSuffix?: string, _id?: string, topicArn: string }[];

    constructor(name: string, args: NotificationArgs) {
        super(name);
        this.bucket = args.bucket;
        this.lambdaFunction = args.lambdaFunction;
        this.queue = args.queue;
        this.topic = args.topic;
    }
}

export interface NotificationArgs {
    readonly bucket: string;
    readonly lambdaFunction?: { events: string[], filterPrefix?: string, filterSuffix?: string, _id?: string, lambdaFunctionArn?: string }[];
    readonly queue?: { events: string[], filterPrefix?: string, filterSuffix?: string, _id?: string, queueArn: string }[];
    readonly topic?: { events: string[], filterPrefix?: string, filterSuffix?: string, _id?: string, topicArn: string }[];
}
