package service

import (
	"fmt"
	"sys-api/dto/body"
	"sys-api/pkg/timestamp_repository"
)

func GetCapacities(n int) ([]body.TimestampedCapacities, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch capacities. details: %s", err)
	}

	if n == 0 {
		n = 1
	}

	result, err := timestamp_repository.NewClient(n).FetchCapacities()
	if err != nil {
		return nil, makeError(err)
	}

	return result, nil
}
