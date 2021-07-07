module github.com/gremlin/CloudFormation-Resources/gremlin-agent-helm

go 1.16

require (
	github.com/aws-cloudformation/cloudformation-cli-go-plugin v1.0.3
	github.com/aws-quickstart/quickstart-helm-resource-provider v1.1.1
	github.com/aws/aws-lambda-go v1.24.0
	github.com/aws/aws-sdk-go v1.38.63
	github.com/stretchr/testify v1.7.0
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
