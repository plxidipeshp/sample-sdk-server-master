package config

import (
	"flag"
	"log"
)

type config struct {
	ProdClientID string

	ProdClientSecret string

	TestClientID string

	TestClientSecret string

	ProdURL string

	TestURL string
}

// Config stores the configs
var Config config

func init() {
	prodClientID := flag.String("production-client-id", "PwfsiztXNvfeiaAPU4zCozaXIn15pQMUW9omVdcG", "Production Client ID")
	prodClientSecret := flag.String("production-client-secret", "wS2spkqBREEJety3kLJ6idxSRIOrg1gTl4jHnM7EHfk92PjGWDULQJeIbl65bq4oLvbsuDOI81oA1wvFvEGzmBu9fFHgiCvtZGkmAVU22kyi4LQM9CHTrh7kiiOJKi5f", "Production Client Secret")
	testClientID := flag.String("test-client-id", "test_YoTUOtL9yPqtD6pKm1yekhgFgmBBKMZ39Mx", "Test Client ID")
	testClientSecret := flag.String("test-client-secret", "test_bSr0i0DkYzZcoMTN4PRRrWHsbtS4V2JLOhTE7RbTM6403SlDbl3az2PGZjw4GpwQ7ahemvhIM3iPni5ZGxFKOlWcKypzelXZFslNwbnmMt2EdB56KVDb1gW56DW", "Test Client Secret")
	flag.Parse()

	if *prodClientID == "" {
		log.Fatal("Production Client ID is missing")
	}

	if *prodClientSecret == "" {
		log.Fatal("Production Client secret is missing")
	}

	if *testClientID == "" {
		log.Fatal("Test Client ID is missing")
	}

	if *testClientSecret == "" {
		log.Fatal("Test Client Secret is missing")
	}

	Config = config{
		ProdClientID:     *prodClientID,
		ProdClientSecret: *prodClientSecret,
		TestClientID:     *testClientID,
		TestClientSecret: *testClientSecret,
		ProdURL:          "https://api.instamojo.com",
		TestURL:          "https://test.instamojo.com",
	}
}
