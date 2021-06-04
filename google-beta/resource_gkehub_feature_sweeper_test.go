package google

import (
	"context"
	"log"
	"testing"

	gkehub "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("GkeHubFeature", &resource.Sweeper{
		Name: "GkeHubFeature",
		F:    testSweepGkeHubFeature,
	})
}

func testSweepGkeHubFeature(region string) error {
	resourceName := "GkeHubFeature"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

	config, err := sharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
		return err
	}

	t := &testing.T{}
	billingId := getTestBillingAccountFromEnv(t)

	// Setup variables to be used for Delete arguments.
	d := map[string]string{
		"project":         config.Project,
		"region":          region,
		"location":        region,
		"zone":            "-",
		"billing_account": billingId,
	}

	client := CreateGkeHubClient(config, config.userAgent, "")
	err = client.DeleteAllFeature(context.Background(), d["project"], d["location"], isDeletableGkeHubFeature)
	if err != nil {
		return err
	}
	return nil
}

func isDeletableGkeHubFeature(r *gkehub.Feature) bool {
	return isSweepableTestResource(*r.Name)
}
