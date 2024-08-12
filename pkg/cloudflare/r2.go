package cloudflare

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

// Global R2Client instance
var R2 *R2Client

// R2Client wraps the AWS S3 client for Cloudflare R2 operations
type R2Client struct {
	svc        *s3.Client
	bucketName string
}

// NewR2Client creates a new R2Client with the specified configuration
func NewR2Client(endpoint, accessKeyID, secretKey, bucketName string) (*R2Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("auto"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{URL: endpoint}, nil
		})),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	svc := s3.NewFromConfig(cfg)

	R2 = &R2Client{
		svc:        svc,
		bucketName: bucketName,
	}

	return R2, nil
}

// ListObjects lists objects in the R2 bucket
func (c *R2Client) ListObjects(ctx context.Context) (*s3.ListObjectsV2Output, error) {

	result, err := c.svc.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &c.bucketName,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UploadFile uploads a file to the R2 bucket
// UploadFile uploads a file to the R2 bucket
func (c *R2Client) UploadFile(ctx context.Context, file io.Reader, key string) error {

	_, err := c.svc.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &c.bucketName,
		Key:    &key,
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}

// DeleteObject deletes an object from the R2 bucket
func (c *R2Client) DeleteObject(ctx context.Context, key string) error {

	_, err := c.svc.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &c.bucketName,
		Key:    &key,
	})
	if err != nil {
		return err
	}

	return nil
}

// GetObjectMetadata retrieves the metadata of an object from the R2 bucket
func (c *R2Client) GetObjectMetadata(ctx context.Context, key string) (*s3.HeadObjectOutput, error) {

	result, err := c.svc.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &c.bucketName,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
