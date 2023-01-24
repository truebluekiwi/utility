package template

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"html/template"
)

func GetTemplateByID(downloader *s3manager.Downloader, id string, bucket string) (*template.Template, error) {
	htmlBuf := aws.NewWriteAtBuffer([]byte{})
	if _, err := downloader.Download(htmlBuf,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(id + ".html"),
		},
	); err != nil {
		if err.(awserr.Error).Code() == s3.ErrCodeNoSuchKey {
			return nil, nil
		}
		return nil, err
	}

	tpl, err := template.New(id).Parse(string(htmlBuf.Bytes()))
	if err != nil {
		return nil, err
	}

	return tpl, nil
}
