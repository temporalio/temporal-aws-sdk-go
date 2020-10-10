# Temporal AWS SDK

| :warning: **This code is an experiment**: Absolutely no guarantee of backwards compatibility |
| --- |

This repository is a prototype of generated [Temporal](https://github.com/temporalio/) activities and workflow stubs for all AWS APIs
exposed by [AWS Go SDK](https://github.com/aws/aws-sdk-go).

Generated activities and stubs are in Go at this point. 
Java SDK has to call them by string name until Java activity interface generation is added.  

Contact [Maxim](https://github.com/mfateev) if you want to collaborate on this project. 

## Generated Code

`awsactivities` package contains generated code for activities that call AWS APIs through AWS Go SDK.

`awsclients` package contains generated code that workflows can use to call these activities in a
strongly typed manner.

## Regenerating Code

Initialize submodule that contains aws-sdk-go
```
git submodule update --init --recursive
```
Regenerate code for all services:
```
go run cmd/temporal-aws-sdk-gen/main.go
```
Regenerate code for a single service:
```
go run cmd/temporal-aws-sdk-gen/main.go
```

## Templates

Templates that are used for code generation are located in 
[cmd/temporal-aws-sdk-gen/templates](cmd/temporal-aws-sdk-gen/templates) directory.
They use Go [text/template](https://golang.org/pkg/text/template/) package. Each template is invoked with
[InterfaceDefinition](cmd/temporal-aws-sdk-gen/internal/parser.go#L31) structure as a parameter.

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
```bash
go run cmd/awsactivities/main.go
```

## Samples

### S3 List Bucket

`cmd/s3list` directory contains very simple example of a workflow that lists S3 buckets under an account. The sample
assumes that AWS SDK activities are already running.

```bash
go run cmd/s3list/worker/main.go
```

### EC2 Instance

`cmd/ec2demo` directory contains an example of a workflow that launches an EC2 instance and keeps
restarting it when terminated. The sample assumes that AWS SDK activities are already running.

To start workflow worker:
```bash
go run cmd/ec2demo/worker/main.go
```

To initiate workflow
```bash
go run cmd/ec2demo/starter/main.go
```