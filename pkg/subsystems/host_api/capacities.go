package host_api

import (
	"fmt"
	"sys-api/utils"
)

func (c *Client) GetCapacities() (*Capacities, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to get host capacities. details: %s", err)
	}

	response, err := utils.DoJsonGetRequest[Capacities](c.URL+"/capacities", nil)
	if err != nil {
		return nil, makeError(err)
	}

	return response, nil
}
