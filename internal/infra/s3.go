package infra

import "boilerplate/pkg/aws"

func NewS3(region, accesskey, secretkey string) (module aws.S3, err error) {
	module = aws.NewS3(region, accesskey, secretkey)
	if err = module.SetConfig(); err != nil {
		return
	}
	return
}
