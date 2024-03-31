package service

import (
	"fmt"
	"sys-api/dto/body"
	"sys-api/pkg/timestamp_repository"
)

func GetStatus(n int) ([]body.TimestampedStatus, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch status. details: %s", err)
	}

	if n == 0 {
		n = 1
	}

	result, err := timestamp_repository.NewClient(n).FetchStatus()
	if err != nil {
		return nil, makeError(err)
	}

	return result, nil
}
