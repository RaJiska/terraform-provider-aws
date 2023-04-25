package ivs_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccIVS_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"PlaybackKeyPair": {
			"basic":      testAccPlaybackKeyPair_basic,
			"update":     testAccPlaybackKeyPair_update,
			"tags":       testAccPlaybackKeyPair_tags,
			"disappears": testAccPlaybackKeyPair_disappears,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
