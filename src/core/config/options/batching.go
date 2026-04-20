package configOptions

import (
	"fmt"
	"os"
	"strconv"
	"trains/src/core/config/types"
)

func BatchingFromEnv(config *types.Config) {
	if v := os.Getenv("TRAINS_BATCHING_TOKEN_LIMIT"); v != "" {
		tokenLimit, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		config.Batching.TokenLimit = tokenLimit
	}
}
