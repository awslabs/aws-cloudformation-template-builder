# AWS Cloudformation Template Builder

This repository contains `cfn-skeleton`, a command line tool and Go library that consumes the published [CloudFormation specification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-resource-specification.html) and generates skeleton CloudFormation templates with mandatory and optional parameters of chosen resource types pre-filled with placeholder values.

## License

This project is licensed under the Apache 2.0 License. 

## Usage

```console
cfn-skeleton [OPTIONS] [RESOURCE TYPES...]

  cfn-skeleton is a tool that generates skeleton CloudFormation templates
  containing definitions for the resource types that you specify.

  You can use a short name for a resource type so long as it is unambiguous.
  For example 'Bucket', 'S3::Bucket', and 'AWS::S3::Bucket' refer to the same type.
  But 'Instance' would need disambiguation.

Options:

  -b, --bare  Produce a minimal template, omitting all optional resource properties.
  -j, --json  Output the template in JSON format (default: YAML).
  -i, --iam   If any resource includes an IAM policy definition, populate that too.
  --help      Show this message and exit.
```
