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

Templates that are used for code generation are located in [cmd/temporal-aws-sdk-gen/templates] directory.
They use Go [text/template](https://golang.org/pkg/text/template/) package. Each template is invoked with
[InterfaceDefinition](cmd/temporal-aws-sdk-gen/internal/parser.go#L31) structure as a parameter.