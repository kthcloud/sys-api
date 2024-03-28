package service

import (
	"fmt"
	"sys-api/dto/body"
	"sys-api/pkg/timestamp_repository"
)

func GetGpuInfo(n int) ([]body.TimestampedGpuInfo, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch gpu info. details: %s", err)
	}

	if n == 0 {
		n = 1
	}

	result, err := timestamp_repository.NewClient(n).FetchGpuInfo()
	if err != nil {
		return nil, makeError(err)
	}

	return result, nil
}
