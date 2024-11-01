package config
//aliyun oss cofig
type OSS struct {
	// Endpoint 访问域名
	Endpoint string
	// AccessKeyID AK
	AccessKeyID string
	// AccessKeySecret AKS
	AccessKeySecret string
	// BucketName 桶名称
	BucketName string
}


var OSSConfig = new(OSS)
