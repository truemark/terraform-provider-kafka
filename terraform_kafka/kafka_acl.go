package kafka_topic

import (

	// clientapi "github.com/cgroschupp/go-client-confluent-cloud/confluentcloud"

	"context"
	"errors"

	sarama "github.com/Shopify/sarama"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//
// This tool helps to manage acls on kafka.
// Option                                   Description
// ------                                   -----------
// --add
// Indicates you are trying to add ACLs.
//
// --allow-host <String: allow-host>
// Host from which principals listed in --
//                                            allow-principal will have access. If
//                                            you have specified --allow-principal
//                                            then the default for this option
//                                            will be set to * which allows access
//                                            from all hosts.
// --allow-principal <String: allow-principal>
// principal is in principalType:name format. Note that principalType must be supported by the Authorizer being used. For example, User:* is the wild card indicating all users.
//
// --authorizer <String: authorizer>
// Fully qualified class name of the authorizer, defaults to kafka.security.authorizer.AclAuthorizer.
// --authorizer-properties <String:authorizer-properties>
// REQUIRED: properties required to configure an instance of Authorizer. These are key=val pairs. For the default authorizer the example values are: zookeeper.connect=localhost:2181

// --bootstrap-server <String: server to connect to>
// A list of host/port pairs to use for establishing the connection to the Kafka cluster.
// This list should be in the form host1:port1,host2:port2,... This config is required
// for acl management using admin client API.
//
// --cluster                                Add/Remove cluster ACLs.
//
// --command-config [String: command-       A property file containing configs to
//   config]                                  be passed to Admin Client.
//
// --consumer                               Convenience option to add/remove ACLs
//                                            for consumer role. This will
//                                            generate ACLs that allows READ,
//                                            DESCRIBE on topic and READ on group.
//
// --delegation-token <String: delegation-  Delegation token to which ACLs should
//   token>                                   be added or removed. A value of *
//                                            indicates ACL should apply to all
//                                            tokens.
//
// --deny-host <String: deny-host>          Host from which principals listed in --
//                                            deny-principal will be denied
//                                            access. If you have specified --deny-
//                                            principal then the default for this
//                                            option will be set to * which denies
//                                            access from all hosts.
//
// --deny-principal <String: deny-          principal is in principalType:name
//   principal>                               format. By default anyone not added
//                                            through --allow-principal is denied
//                                            access. You only need to use this
//                                            option as negation to already
//                                            allowed set. Note that principalType
//                                            must be supported by the Authorizer
//                                            being used. For example if you
//                                            wanted to allow access to all users
//                                            in the system but not test-user you
//                                            can define an ACL that allows access
//                                            to User:* and specify --deny-
//                                            principal=User:test@EXAMPLE.COM. AND
//                                            PLEASE REMEMBER DENY RULES TAKES
//                                            PRECEDENCE OVER ALLOW RULES.
//
// --force                                  Assume Yes to all queries and do not
//                                            prompt.
//
// --group <String: group>                  Consumer Group to which the ACLs
//                                            should be added or removed. A value
//                                            of * indicates the ACLs should apply
//                                            to all groups.
//
// --help                                   Print usage information.
//
// --idempotent                             Enable idempotence for the producer.
//                                            This should be used in combination
//                                            with the --producer option. Note
//                                            that idempotence is enabled
//                                            automatically if the producer is
//                                            authorized to a particular
//                                            transactional-id.
//
// --list                                   List ACLs for the specified resource,
//                                            use --topic <topic> or --group
//                                            <group> or --cluster to specify a
//                                            resource.
//
// --operation <String>                     Operation that is being allowed or
//                                            denied. Valid operation names are:
//                                          	Describe
//                                          	DescribeConfigs
//                                          	Alter
//                                          	IdempotentWrite
//                                          	Read
//                                          	Delete
//                                          	Create
//                                          	ClusterAction
//                                          	All
//                                          	Write
//                                          	AlterConfigs
//                                           (default: All)
//
// --principal [String: principal]          List ACLs for the specified principal.
//                                            principal is in principalType:name
//                                            format. Note that principalType must
//                                            be supported by the Authorizer being
//                                            used. Multiple --principal option
//                                            can be passed.
//
// --producer                               Convenience option to add/remove ACLs
//                                            for producer role. This will
//                                            generate ACLs that allows WRITE,
//                                            DESCRIBE and CREATE on topic.
//
// --remove                                 Indicates you are trying to remove
//                                            ACLs.
//
// --resource-pattern-type                  The type of the resource pattern or
//   <ANY|MATCH|LITERAL|PREFIXED>             pattern filter. When adding acls,
//                                            this should be a specific pattern
//                                            type, e.g. 'literal' or 'prefixed'.
//                                            When listing or removing acls, a
//                                            specific pattern type can be used to
//                                            list or remove acls from specific
//                                            resource patterns, or use the filter
//                                            values of 'any' or 'match', where
//                                            'any' will match any pattern type,
//                                            but will match the resource name
//                                            exactly, where as 'match' will
//                                            perform pattern matching to list or
//                                            remove all acls that affect the
//                                            supplied resource(s). WARNING:
//                                            'match', when used in combination
//                                            with the '--remove' switch, should
//                                            be used with care. (default: LITERAL)
//
// --topic <String: topic>                  topic to which ACLs should be added or
//                                            removed. A value of * indicates ACL
//                                            should apply to all topics.
//
// --transactional-id <String:transactional-id>
// The transactionalId to which ACLs
//                           should be added or removed. A value
//                                            of * indicates the ACLs should apply
//                                            to all transactionalIds.

// --version                                Display Kafka version.
//
// --zk-tls-config-file <String: Authorizer ZooKeeper TLS configuration>
// 		Identifies the file where ZooKeeper client TLS connectivity properties for
// 		the authorizer are defined. Any properties other than the following (with or
//		without an "authorizer." prefix) are ignored:
// 			zookeeper.clientCnxnSocket,
// 			zookeeper.ssl.cipher.suites,
// 			zookeeper.ssl.client.enable,
// zookeeper.ssl.crl.enable,
//                                            zookeeper.ssl.enabled.protocols,
//                                            zookeeper.ssl.endpoint.identification.algorithm,
// zookeeper.ssl.keystore.location,
// zookeeper.ssl.keystore.password,
// zookeeper.ssl.keystore.type,
// zookeeper.ssl.ocsp.enable,
// zookeeper.ssl.protocol,
// zookeeper.ssl.truststore.location,
// zookeeper.ssl.truststore.password,
// zookeeper.ssl.truststore.type.

// Note that if SASL is not configured and zookeeper.set.acl is supposed to be
// true due to mutual certificate authentication being used then it is necessary
// to explicitly specify -- authorizer-properties zookeeper.set.acl=true

func ACLCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	brokers := d.Get("bootstrap_servers").([]string)

	// principle := d.Get("principle")
	host := d.Get("host").(string)
	operationStr := d.Get("operation").(string)
	permissionTypeStr := d.Get("permission").(string)
	resource := d.Get("resource").(map[string]interface{})
	resourceTypeStr := resource["resource_type"].(string)
	resourceName := resource["resource_name"].(string)
	resourcePatternTypeStr := resource["resource_pattern_type"].(string)

	operation, err := operationToOpCode(operationStr)
	permission, err := permissionToOpCode(permissionTypeStr)

	resourceOp, err := resourceToOpCode(resourceTypeStr)
	patternOp, err := resourcePatternToOpCode(resourcePatternTypeStr)
	// r := Resource{ResourceType: AclResourceTopic, ResourceName: "my_topic"}
	// a := Acl{Host: "localhost", Operation: AclOperationAlter, PermissionType: AclPermissionAny}

	resourceStruct := sarama.Resource{
		ResourceType:        sarama.AclResourceType(resourceOp),
		ResourceName:        resourceName,
		ResourcePatternType: sarama.AclResourcePatternType(patternOp),
	}
	aclStruct := sarama.Acl{
		Host:           host,
		Operation:      sarama.AclOperation(operation),
		PermissionType: sarama.AclPermissionType(permission),
	}

	config := sarama.NewConfig()
	client, err := sarama.NewClient(brokers, config)
	clusterAdmin, err := sarama.NewClusterAdminFromClient(client)
	err = clusterAdmin.CreateACL(resourceStruct, aclStruct)
	if err != nil {
		// t.Fatal(err)
	}
	err = clusterAdmin.Close()
	if err != nil {
		// t.Fatal(err)
	}

	return nil
}

func ACLRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func ACLDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func ResourceKafkaACL() *schema.Resource {
	// Acl holds information about acl type
	return &schema.Resource{
		CreateContext: ACLCreate,
		ReadContext:   ACLRead,
		DeleteContext: ACLDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"principal": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"host": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"operation": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: ValidateOperationType,
			},
			"permission": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: ValidatePermissionType,
			},
			"resource": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				// ValidateFunc: ValidateResource,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: ValidateResourceType,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"resource_pattern_type": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: ValidateResourcePatternType,
						},
					},
				},
			},
		},
	}
}

func ValidateOperationType(val interface{}, key string) (warns []string, errs []error) {
	validOperations := []string{
		"unknown",
		"any",
		"all",
		"read",
		"write",
		"create",
		"delete",
		"alter",
		"describe",
		"cluster_action",
		"describe_configs",
		"alter_configs",
		"idempotent_write"}
	for _, operation := range validOperations {
		if val == operation {
			return nil, nil
		}
	}
	return []string{"Error: operation.type value is incorrect. Must be one of [unknown, any, \"all\", \"read\", \"write\", \"create\", \"delete\", \"alter\", \"describe\", \"cluster_action\", \"describe_configs\", \"alter_configs\", \"idempotent_write\"]"}, []error{nil}
}

func ValidatePermissionType(val interface{}, key string) (warns []string, errs []error) {
	permissionTypes := []string{"unknown", "any", "deny", "allow"}
	for _, permission := range permissionTypes {
		if val == permission {
			return nil, nil
		}
	}
	return []string{"Error: permission.type value is incorrect. Must be one of [\"unknown\", \"any\", \"deny\", \"allow\"]"}, []error{nil}
}

func ValidateResourceType(val interface{}, key string) (warns []string, errs []error) {
	validResourceTypes := []string{"unknown", "any", "topic", "group",
		"cluster", "transaction_id", "delegation_token"}

	for _, resource := range validResourceTypes {
		if val == resource {
			return nil, nil
		}
	}
	return []string{"Error: resource.type value is incorrect. Must be one of [\"unknown\", \"any\", \"topic\", \"group\", \"cluster\", \"transaction_id\", \"delegation_token\"]"}, []error{nil}
}

func ValidateResourcePatternType(val interface{}, key string) (warns []string, errs []error) {
	validResourcePatternTypes := []string{"unknown", "any", "match", "literal", "prefixed"}
	for _, resourcePattern := range validResourcePatternTypes {
		if val == resourcePattern {
			return nil, nil
		}
	}
	return []string{"Error: resource.pattern.type value is incorrect. Must be one of [\"unknown\", \"any\", \"match\", \"literal\", \"prefixed\"]"}, []error{nil}
}

func operationToOpCode(operation string) (int, error) {
	switch operation {
	case "unknown":
		return 0, nil
	case "any":
		return 1, nil
	case "all":
		return 2, nil
	case "read":
		return 3, nil
	case "write":
		return 4, nil
	case "create":
		return 5, nil
	case "delete":
		return 6, nil
	case "alter":
		return 7, nil
	case "describe":
		return 8, nil
	case "cluster_action":
		return 9, nil
	case "describe_configs":
		return 10, nil
	case "alter_configs":
		return 11, nil
	case "idempotent_write":
		return 12, nil
	}

	return -1, errors.New("Unable to locate operation code")
}

func permissionToOpCode(permission string) (int, error) {
	switch permission {
	case "unknown":
		return 0, nil
	case "any":
		return 1, nil
	case "deny":
		return 2, nil
	case "allow":
		return 3, nil
	}
	return -1, errors.New("Unable to locate permission code")
}

func resourceToOpCode(resource string) (int, error) {
	switch resource {
	case "unknown":
		return 0, nil
	case "any":
		return 1, nil
	case "topic":
		return 2, nil
	case "group":
		return 3, nil
	case "cluster":
		return 4, nil
	case "transaction_id":
		return 5, nil
	case "delegation_token":
		return 6, nil
	}
	return -1, errors.New("Unable to locate resource code")
}

func resourcePatternToOpCode(resourcePattern string) (int, error) {
	switch resourcePattern {
	case "unknown":
		return 0, nil
	case "any":
		return 1, nil
	case "match":
		return 2, nil
	case "literal":
		return 3, nil
	case "prefixed":
		return 4, nil
	}
	return -1, errors.New("Unable to locate resource code")
}
