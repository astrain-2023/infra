package main

import (
	"fmt"

	"github.com/pulumi/pulumi-confluentcloud/sdk/go/confluentcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		prefix := "astra-gcp"

		env, err := confluentcloud.NewEnvironment(ctx, prefix, &confluentcloud.EnvironmentArgs{
			DisplayName: pulumi.String(prefix),
		})
		if err != nil {
			return err
		}
		gcpFrankfurt, err := confluentcloud.GetSchemaRegistryRegion(ctx, &confluentcloud.GetSchemaRegistryRegionArgs{
			Cloud:   "GCP",
			Region:  "europe-west3",
			Package: "ESSENTIALS",
		}, nil)
		if err != nil {
			return err
		}

		_, err = confluentcloud.NewSchemaRegistryCluster(ctx, pref(prefix, "sreg-essentials"), &confluentcloud.SchemaRegistryClusterArgs{
			Package: pulumi.String(gcpFrankfurt.Package),
			Environment: &confluentcloud.SchemaRegistryClusterEnvironmentArgs{
				Id: env.ID(),
			},
			Region: &confluentcloud.SchemaRegistryClusterRegionArgs{
				Id: pulumi.String(gcpFrankfurt.Id),
			},
		})
		if err != nil {
			return err
		}

		// https: //www.pulumi.com/registry/packages/confluentcloud/api-docs/kafkacluster/#inputs
		_, err = confluentcloud.NewKafkaCluster(ctx, pref(prefix, "west4"), &confluentcloud.KafkaClusterArgs{
			Availability: pulumi.String("SINGLE_ZONE"),
			Cloud:        pulumi.String("GCP"),
			Environment: &confluentcloud.KafkaClusterEnvironmentArgs{
				Id: env.ID(),
			},
			// Netherlands $0.1155/hour
			Region: pulumi.String("europe-west4"),
			Basic:  &confluentcloud.KafkaClusterBasicArgs{},
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func pref(prefix string, name string) string {
	return fmt.Sprintf("%s-%s", prefix, name)
}
