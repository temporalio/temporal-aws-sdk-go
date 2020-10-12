# Temporal AWS SDK

| :warning: **This code is an experiment**: Absolutely no guarantee of backwards compatibility |
| --- |

This repository is a prototype of generated [Temporal](https://github.com/temporalio/) activities and workflow stubs for all AWS APIs
exposed by [AWS Go SDK](https://github.com/aws/aws-sdk-go).

Generated activities and stubs are in Go at this point.
Java SDK has to call them by string name until Java activity interface generation is added.

This repository relies on the [temporal-aws-sdk-generator](https://github.com/temporalio/temporal-aws-sdk-generator).

## Generated Code

`activities` package contains generated code for activities that call AWS APIs through AWS Go SDK.

`clients` package contains generated code that workflows can use to call these activities in a
strongly typed manner.

## Regenerating Code

Update a template in templates directory and run:
```
make bins
```
to regenerate and compile the generated code.

## Templates

Templates that are used for code generation are located in
[templates](templates) directory.
They use Go [text/template](https://golang.org/pkg/text/template/) package. Each template is invoked with
AWSSDKDefinition](internal/definitions.go) structure as a parameter. A template file must call `SetFileName`
function at the beginning to specify the output file name.
It is allowed to call `SetFileName` multiple times to generate multiple files from the same template.

## Running Temporal AWS SDK Activities

To run Temporal AWS SDK activities:

Create `~/.aws/credentials` file with aws access key id and secret or set corresponding environment variables.
The basic credentials file format:
```
[default]
aws_access_key_id = <ID>
aws_secret_access_key = <SECRET>
```
See [AWS Go SDK documentation](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html) for details.
Set AWS_REGION environment variable. Another option is to set AWS_SDK_LOAD_CONFIG variable to true and set Region config value.
```bash
AWS_REGION=us-west-2 bin/aws-sdk-worker
```

## Samples

### S3 List Bucket

`samples/s3list` directory contains very simple example of a workflow that lists S3 buckets under an account. The sample
assumes that AWS SDK activities are already running.

```bash
bin/samples/s3list
```

### EC2 Instance

`samples/ec2demo` directory contains an example of a workflow that launches an EC2 instance and keeps
restarting it when terminated. The sample assumes that AWS SDK activities are already running.

To start workflow worker:
```bash
bin/samples/ec2demo-worker
```

To initiate workflow
```bash
bin/samples/ec2demo-starter
```
