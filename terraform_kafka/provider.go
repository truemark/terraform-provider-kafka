package kafka_topic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/diag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sarama "gopkg.in/Shopify/sarama.v2"
)

func Provider() *schema.Provider {
	// bootstrap_servers

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"bootstrap_servers": {
				Type:     schema.TypeList,
				Optional: false,
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"truemark-kafka_acl":   ResourceKafkaACL(),
			"truemark-kafka_topic": ResourceKafkaTopic(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// TODO: log.Printf("[INFO] Initializing ConfluentCloud client")
	bootstrapServers := d.Get("bootstrap_servers").(string)

	config := &sarama.Config{
		BootstrapServers:        bootstrapServers,
		CACert:                  nil,
		ClientCert:              nil,
		ClientCertKey:           nil,
		ClientCertKeyPassphrase: nil,
		SkipTLSVerify:           nil,
		SASLUsername:            nil,
		SASLPassword:            nil,
		SASLMechanism:           nil,
		TLSEnabled:              nil,
		Timeout:                 nil,
	}

	// err := resource.RetryContext(ctx, 30*time.Minute, func() *resource.RetryError {
	// })
	return config, nil // diag.FromErr(err)
}
