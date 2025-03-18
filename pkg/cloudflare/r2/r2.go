package r2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"os"
	"time"
)

var (
	bucketName      = "sdk-example"
	accountId       = "<accountid>"
	accessKeyId     = "<access_key_id>"
	accessKeySecret = "<access_key_secret>"
	baseDir         = "assets/"
	r2Domain        = "abc.com"
)

func init() {
	bucketName = os.Getenv("r2_bucket_name")
	accountId = os.Getenv("r2_account_id")
	accessKeyId = os.Getenv("r2_access_key_id")
	accessKeySecret = os.Getenv("r2_access_key_secret")
	r2Domain = os.Getenv("r2_domain")
}

type S3Handler interface {
	ListObjectsOutput(context.Context) error
	UploadFile(ctx context.Context, fileName string, fileContent []byte, contentType string) (string, error)
}

type R2Server struct {
	S3Client *s3.Client
	S3Manger *manager.Uploader
}

func NewR2Server() *R2Server {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
		// Compatibility https://developers.cloudflare.com/r2/examples/aws/aws-sdk-go/
		config.WithRequestChecksumCalculation(0),
		config.WithResponseChecksumValidation(0),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
	})
	return &R2Server{
		client,
		manager.NewUploader(client),
	}
}

func (r *R2Server) ListObjectsOutput(ctx context.Context) error {
	listObjectsOutput, err := r.S3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}
	for _, object := range listObjectsOutput.Contents {
		obj, _ := json.MarshalIndent(object, "", "\t")
		logx.Info(string(obj))
	}
	return nil
}

func (r *R2Server) UploadFile(ctx context.Context, fileName string, fileContent []byte, contentType string) (string, error) {
	var outKey string
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(fmt.Sprintf("%s%s", baseDir, fileName)),
		Body:        bytes.NewReader(fileContent),
		ContentType: aws.String(contentType),
	}
	output, err := r.S3Manger.Upload(ctx, input)
	if err != nil {
		var noBucket *types.NoSuchBucket
		if errors.As(err, &noBucket) {
			logx.Errorf("Bucket %s does not exist.\n", bucketName)
			err = noBucket
		}
	} else {
		err := s3.NewObjectExistsWaiter(r.S3Client).Wait(ctx, &s3.HeadObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
		}, time.Minute)
		if err != nil {
			logx.Errorf("Failed attempt to wait for object %s to exist in %s.\n", fileName, bucketName)
		} else {
			outKey = *output.Key
		}
	}
	return fmt.Sprintf("https://%s/%s", r2Domain, outKey), err
}
