package kafka_topic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	// bootstrap_servers

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"bootstrap_servers": {
				Type:     schema.TypeList,
				Optional: false,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
	bootstrapServers := d.Get("bootstrap_servers").([]interface{})

	// err := resource.RetryContext(ctx, 30*time.Minute, func() *resource.RetryError {
	// })
	return bootstrapServers, nil // diag.FromErr(err)
}

// config := &Config{
// 	BootstrapServers:        bootstrapServers,
// 	CACert:                  nil,
// 	ClientCert:              nil,
// 	ClientCertKey:           nil,
// 	ClientCertKeyPassphrase: nil,
// 	SkipTLSVerify:           nil,
// 	SASLUsername:            nil,
// 	SASLPassword:            nil,
// 	SASLMechanism:           nil,
// 	TLSEnabled:              nil,
// 	Timeout:                 nil,
// }
