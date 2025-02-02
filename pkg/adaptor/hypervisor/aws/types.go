package aws

import "github.com/confidential-containers/cloud-api-adaptor/pkg/util"

type Config struct {
	AccessKeyId        string
	SecretKey          string
	Region             string
	LoginProfile       string
	LaunchTemplateName string
	ImageId            string
	InstanceType       string
	SecurityGroupId    string
	KeyName            string
	SubnetId           string
	UseLaunchTemplate  bool
}

func (c Config) Redact() Config {
	return *util.RedactStruct(&c, "AccessKeyId", "SecretKey").(*Config)
}
