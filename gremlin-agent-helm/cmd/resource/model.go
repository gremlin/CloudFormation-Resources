// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource


// Model is autogenerated from the json schema
type Model struct {
    ClusterID *string `json:",omitempty"`
    KubeConfig *string `json:",omitempty"`
    RoleArn *string `json:",omitempty"`
    Namespace *string `json:",omitempty"`
    Name *string `json:",omitempty"`
    Values map[string]string `json:",omitempty"`
    ValueYaml *string `json:",omitempty"`
    Version *string `json:",omitempty"`
    ValueOverrideURL *string `json:",omitempty"`
    ID *string `json:",omitempty"`
    TimeOut *int `json:",omitempty"`
    VPCConfiguration *VPCConfiguration `json:",omitempty"`
}


// VPCConfiguration is autogenerated from the json schema
type VPCConfiguration struct {
    SecurityGroupIds []string `json:",omitempty"`
    SubnetIds []string `json:",omitempty"`
}

