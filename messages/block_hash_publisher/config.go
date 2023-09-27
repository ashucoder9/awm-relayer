package block_hash_publisher

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/pkg/errors"
)

type destinationInfo struct {
	ChainID  string `json:"chain-id"`
	Interval string `json:"interval"`

	useTimeInterval     bool
	blockInterval       int
	timeIntervalSeconds time.Duration
}

type Config struct {
	DestinationChains []destinationInfo `json:"destination-chains"`
}

func (c *Config) Validate() error {
	for i, destinationInfo := range c.DestinationChains {
		if _, err := ids.FromString(destinationInfo.ChainID); err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid subnetID in block hash publisher configuration. Provided ID: %s", destinationInfo.ChainID))
		}

		// Intervals must be either a positive integer, or a positive integer followed by "s"
		interval, isSeconds, err := parsePositiveIntWithSuffix(destinationInfo.Interval)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid interval in block hash publisher configuration. Provided interval: %s", destinationInfo.Interval))
		}
		if isSeconds {
			c.DestinationChains[i].timeIntervalSeconds = time.Duration(interval) * time.Second
		} else {
			c.DestinationChains[i].blockInterval = interval
		}
		c.DestinationChains[i].useTimeInterval = isSeconds
	}
	return nil
}

func parsePositiveIntWithSuffix(input string) (int, bool, error) {
	// Check if the input string is empty
	if input == "" {
		return 0, false, fmt.Errorf("empty string")
	}

	// Check if the string ends with "s"
	hasSuffix := strings.HasSuffix(input, "s")

	// If it has the "s" suffix, remove it
	if hasSuffix {
		input = input[:len(input)-1]
	}

	// Parse the string as an integer
	intValue, err := strconv.Atoi(input)

	// Check if the parsed value is a positive integer
	if err != nil || intValue < 0 {
		return 0, false, err
	}

	return intValue, hasSuffix, nil
}
