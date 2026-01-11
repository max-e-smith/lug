package common

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetDiskUsageEstimate(bucket string, s3client s3.Client, rootPaths []string) (int64, error) {
	var totalSurveysSize int64 = 0

	for _, surveyRootPath := range rootPaths {
		fmt.Printf("Getting disk usage estimate for s3 files on %s at %s\n", bucket, surveyRootPath)
		// TODO paginate
		result, err := s3client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: aws.String(bucket),
			Prefix: aws.String(surveyRootPath),
		})
		if err != nil {
			return totalSurveysSize, err
		}

		for _, object := range result.Contents {
			//log.Printf("key=%s size=%d", aws.ToString(object.Key), *object.Size)
			totalSurveysSize = totalSurveysSize + *object.Size
		}
	}

	return totalSurveysSize, nil
}
