# Gremlin::Agent::Helm

An AWS CloudFormation resource type that deploys the Gremlin agent into EKS clusters using helm.

## Prerequisites

### IAM role
An IAM role is used by CloudFormation to execute this resource type handler code.
A CloudFormation template to create the execution role is available
[here](./execution-role.template.yaml)

### Create an EKS cluster and provide CloudFormation access to the Kubernetes API
EKS clusters use IAM to allow access to the kubernetes API, as the CloudFormation resource types in this project
interact with the kubernetes API, the IAM execution role must be granted access to the kubernetes API. This can be done
in one of two ways:
* Create the cluster using CloudFormation: Currently there is no native way to manage EKS auth using CloudFormation
  (+1 [this GitHub issue](https://github.com/aws/containers-roadmap/issues/554) to help prioritize native support).
  `AWSQS::EKS::Cluster` type fills this gap today. Instructions on activation and usage can be found
  [here](https://github.com/aws-quickstart/quickstart-amazon-eks-cluster-resource-provider/blob/main/README.md).
* Manually: to allow this resource type to access the kubernetes API, follow the
  [instructions in the EKS documentation](https://docs.aws.amazon.com/eks/latest/userguide/add-user-role.html) adding
  the IAM execution role created above to the `system:masters` group. (Note: you can scope this down if you plan to use
  the resource type to only perform specific operations on the kubernetes cluster)

## Activating the resource type
***Coming soon***

## Usage
Properties and return values for the resource type are documented [here](./docs/README.md).

Documentation for the helm chart and it's values are available [here](https://github.com/gremlin/helm/blob/master/gremlin/README.md).

## Examples

### Deploy the Gremlin agent with credentials stored in a managed secret
```yaml
AWSTemplateFormatVersion: "2010-09-09"
Resources:
  KubeStateMetrics:
    Type: "Gremlin::Agent::Helm"
    Properties:
      ClusterID: my-cluster-name
      Name: gremlin-agent
      Namespace: gremlin
      Values:
        gremlin.secret.managed: "true"
        gremlin.secret.type: "secret"
        gremlin.secret.teamID: "YOUR-TEAM-ID"
        gremlin.secret.clusterID: "YOUR-CLUSTER-ID"
        gremlin.secret.teamSecret: "YOUR-TEAM-SECRET"
```
