package kafka_topic

import (

	// clientapi "github.com/cgroschupp/go-client-confluent-cloud/confluentcloud"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ### ./bin/kafka-topics.sh
// Create, delete, describe, or change a topic.
// Option                                   Description
// ------                                   -----------
// --alter                                  Alter the number of partitions,
//                                            replica assignment, and/or
//                                            configuration for the topic.
//
// --at-min-isr-partitions                  if set when describing topics, only
//                                            show partitions whose isr count is
//                                            equal to the configured minimum. Not
//                                            supported with the --zookeeper
//                                            option.
//
// --bootstrap-server <String: server to    REQUIRED: The Kafka server to connect
//   connect to>                              to. In case of providing this, a
//                                            direct Zookeeper connection won't be
//                                            required.
//
// --command-config <String: command        Property file containing configs to be
//   config property file>                    passed to Admin Client. This is used
//                                            only with --bootstrap-server option
//                                            for describing and altering broker
//                                            configs.
//
// --config <String: name=value>
// A topic configuration override for the
//                                            topic being created or altered. The
//                                            following is a list of valid
//                                            configurations:

//                                          	cleanup.policy
//                                          	compression.type
//                                          	delete.retention.ms
//                                          	file.delete.delay.ms
//                                          	flush.messages
//                                          	flush.ms
//                                          	follower.replication.throttled.replicas
//                                          	index.interval.bytes
//                                          	leader.replication.throttled.replicas
//                                          	max.compaction.lag.ms
//                                          	max.message.bytes
//                                          	message.downconversion.enable
//                                          	message.format.version
//                                          	message.timestamp.difference.max.ms
//                                          	message.timestamp.type
//                                          	min.cleanable.dirty.ratio
//                                          	min.compaction.lag.ms
//                                          	min.insync.replicas
//                                          	preallocate
//                                          	retention.bytes
//                                          	retention.ms
//                                          	segment.bytes
//                                          	segment.index.bytes
//                                          	segment.jitter.ms
//                                          	segment.ms
//                                          	unclean.leader.election.enable

//                                          See the Kafka documentation for full
//                                            details on the topic configs. It is
//                                            supported only in combination with --
//                                            create if --bootstrap-server option
//                                            is used (the kafka-configs CLI
//                                            supports altering topic configs with
//                                            a --bootstrap-server option).

// --create                                 Create a new topic.
// --delete                                 Delete a topic
// --delete-config <String: name>           A topic configuration override to be
//                                            removed for an existing topic (see
//                                            the list of configurations under the
//                                            --config option). Not supported with
//                                            the --bootstrap-server option.
// --describe                               List details for the given topics.
// --disable-rack-aware                     Disable rack aware replica assignment
// --exclude-internal                       exclude internal topics when running
//                                            list or describe command. The
//                                            internal topics will be listed by
//                                            default

// --force                                  Suppress console prompts
// --help                                   Print usage information.
// --if-exists                              if set when altering or deleting or
//                                            describing topics, the action will
//                                            only execute if the topic exists.
// --if-not-exists                          if set when creating topics, the
//                                            action will only execute if the
//                                            topic does not already exist.
// --list                                   List all available topics.
// --partitions <Integer: # of partitions>  The number of partitions for the topic
//                                            being created or altered (WARNING:
//                                            If partitions are increased for a
//                                            topic that has a key, the partition
//                                            logic or ordering of the messages
//                                            will be affected). If not supplied
//                                            for create, defaults to the cluster
//                                            default.
// --replica-assignment <String:            A list of manual partition-to-broker
//   broker_id_for_part1_replica1 :           assignments for the topic being
//   broker_id_for_part1_replica2 ,           created or altered.
//   broker_id_for_part2_replica1 :
//   broker_id_for_part2_replica2 , ...>
// --replication-factor <Integer:           The replication factor for each
//   replication factor>                      partition in the topic being
//                                            created. If not supplied, defaults
//                                            to the cluster default.
// --topic <String: topic>                  The topic to create, alter, describe
//                                            or delete. It also accepts a regular
//                                            expression, except for --create
//                                            option. Put topic name in double
//                                            quotes and use the '\' prefix to
//                                            escape regular expression symbols; e.
//                                            g. "test\.topic".
// --topics-with-overrides                  if set when describing topics, only
//                                            show topics that have overridden
//                                            configs
// --unavailable-partitions                 if set when describing topics, only
//                                            show partitions whose leader is not
//                                            available
// --under-min-isr-partitions               if set when describing topics, only
//                                            show partitions whose isr count is
//                                            less than the configured minimum.
//                                            Not supported with the --zookeeper
//                                            option.
// --under-replicated-partitions            if set when describing topics, only
//                                            show under replicated partitions
// --version                                Display Kafka version.
// --zookeeper <String: hosts>              DEPRECATED, The connection string for
//                                            the zookeeper connection in the form
//                                            host:port. Multiple hosts can be
//                                            given to allow fail-over.

func TopicCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// config := sarama.NewConfig()
	// client := sarama.NewClient(config)

	// clusterAdmin := sarama.NewClusterAdminFromClient(client)
	// clusterAdmin.CreateTopic()

	// type CreateTopicsRequest struct {
	// 	Version int16

	// 	TopicDetails map[string]*TopicDetail
	// 	Timeout      time.Duration
	// 	ValidateOnly bool
	// }
	// broker.CreateTopics()

	return nil
}

func TopicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// func ClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
func TopicDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// type TopicDetail struct {
// 	NumPartitions     int32
// 	ReplicationFactor int16
// 	ReplicaAssignment map[int32][]int32
// 	ConfigEntries     map[string]*string
// }

func ResourceKafkaTopic() *schema.Resource {
	return &schema.Resource{
		CreateContext: TopicCreate,
		ReadContext:   TopicRead,
		DeleteContext: TopicDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"num_partitions": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"replication_factor": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"replica_assignment": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
			},
			"config_entries": {
				// Type: schema.TypeMap,
				// Required: true,
				// ForceNew: true,

				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: SchemaKafkaTopicConfigEntries(),
				},
			},
		},
	}
}

func SchemaKafkaTopicConfigEntries() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// cleanup.policy
		// A string that is either "delete" or "compact" or both. This string designates
		// the retention policy to use on old log segments. The default policy ("delete")
		// will discard old segments when their retention time or size limit has been
		// reached. The "compact" setting will enable log compaction on the topic.
		//
		// Type:						list
		// Default:						delete
		// Valid Values:				[compact, delete]
		// Server Default Property:		log.cleanup.policy
		// Importance:					medium
		"cleanup.policy": {
			Type:         schema.TypeList,
			Required:     false,
			ForceNew:     true,
			Description:  "A string that is either \"delete\" or \"compact\" or both. This string designates the retention policy to use on old log segments. The default policy (\"delete\") will discard old segments when their retention time or size limit has been reached. The \"compact\" setting will enable log compaction on the topic.",
			ValidateFunc: ValidateKafkaCleanupPolicy,
		},

		// compression.type
		// Specify the final compression type for a given topic. This configuration accepts
		// the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally
		// accepts 'uncompressed' which is equivalent to no compression; and 'producer' which
		// means retain the original compression codec set by the producer.
		//
		// Type:						string
		// Default:						producer
		// Valid Values:				[uncompressed, zstd, lz4, snappy, gzip, producer]
		// Server Default Property:		compression.type
		// Importance:					medium
		"compression.type": {
			Type:         schema.TypeString,
			Required:     false,
			ForceNew:     true,
			Description:  "Specify the final compression type for a given topic. This configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally accepts 'uncompressed' which is equivalent to no compression; and 'producer' which means retain the original compression codec set by the producer.",
			ValidateFunc: ValidateKafkaCompressionType,
		},

		// delete.retention.ms
		// The amount of time to retain delete tombstone markers for log compacted topics.
		// This setting also gives a bound on the time in which a consumer must complete a
		// read if they begin from offset 0 to ensure that they get a valid snapshot of the
		// final stage (otherwise delete tombstones may be collected before they complete
		// their scan).
		//
		// Type:						long
		// Default:						86400000 (1 day)
		// Valid Values:				[0,...]
		// Server Default Property:		log.cleaner.delete.retention.ms
		// Importance:					medium
		"delete.retention.ms": {
			Type:        schema.TypeInt,
			Required:    false,
			ForceNew:    true,
			Description: "The amount of time to retain delete tombstone markers for log compacted topics. This setting also gives a bound on the time in which a consumer must complete a read if they begin from offset 0 to ensure that they get a valid snapshot of the final stage (otherwise delete tombstones may be collected before they complete their scan).",
		},

		// file.delete.delay.ms
		// The time to wait before deleting a file from the filesystem
		//
		// Type:						long
		// Default:						60000 (1 minute)
		// Valid Values:				[0,...]
		// Server Default Property:		log.segment.delete.delay.ms
		// Importance:					medium
		"file.delete.delay.ms": {
			Type:        schema.TypeInt,
			Required:    false,
			ForceNew:    true,
			Description: "The time to wait before deleting a file from the filesystem",
		},

		// flush.messages
		// This setting allows specifying an interval at which we will force an fsync of data
		// written to the log. For example if this was set to 1 we would fsync after every
		// message; if it were 5 we would fsync after every five messages. In general we
		// recommend you not set this and use replication for durability and allow the
		// operating system's background flush capabilities as it is more efficient. This
		// setting can be overridden on a per-topic basis (see the per-topic configuration
		// section).
		//
		// Type:					long
		// Default:					9223372036854775807
		// Valid Values:			[0,...]
		// Server Default Property:	log.flush.interval.messages
		// Importance:				medium
		"flush.messages": {
			Type:        schema.TypeInt,
			Required:    false,
			ForceNew:    true,
			Description: "This setting allows specifying an interval at which we will force an fsync of data written to the log. For example if this was set to 1 we would fsync after every message; if it were 5 we would fsync after every five messages. In general we recommend you not set this and use replication for durability and allow the operating system's background flush capabilities as it is more efficient. This setting can be overridden on a per-topic basis (see the per-topic configuration  section).",
		},

		// flush.ms
		// This setting allows specifying a time interval at which we will force an fsync of data
		// written to the log. For example if this was set to 1000 we would fsync after 1000 ms
		// had passed. In general we recommend you not set this and use replication for durability
		// and allow the operating system's background flush capabilities as it is more efficient.
		//
		// Type:						long
		// Default:						9223372036854775807
		// Valid Values:				[0,...]
		// Server Default Property:	 	log.flush.interval.ms
		// Importance:					medium
		"flush.ms": {
			Type:        schema.TypeInt,
			Required:    false,
			ForceNew:    true,
			Description: "This setting allows specifying a time interval at which we will force an fsync of data written to the log. For example if this was set to 1000 we would fsync after 1000 ms had passed. In general we recommend you not set this and use replication for durability and allow the operating system's background flush capabilities as it is more efficient.",
		},

		// follower.replication.throttled.replicas
		// A list of replicas for which log replication should be throttled on the follower side.
		// The list should describe a set of replicas in the form
		//    [PartitionId]:[BrokerId],[PartitionId]:[BrokerId]:...
		//
		// or alternatively the wildcard '*' can be used to throttle all replicas for this topic.
		//
		// Type:					list
		// Default:					""
		// Valid Values:			[partitionId]:[brokerId],[partitionId]:[brokerId],...
		// Server Default Property:	follower.replication.throttled.replicas
		// Importance:				medium
		"follower.replication.throttled.replicas": {
			Type:        schema.TypeList,
			Required:    false,
			ForceNew:    true,
			Description: "A list of replicas for which log replication should be throttled on the follower side.",
		},

		// index.interval.bytes
		// This setting controls how frequently Kafka adds an index entry to its offset index.
		// The default setting ensures that we index a message roughly every 4096 bytes. More
		// indexing allows reads to jump closer to the exact position in the log but makes the
		// index larger. You probably don't need to change this.
		//
		// Type:					int
		// Default:					4096 (4 kibibytes)
		// Valid Values:			[0,...]
		// Server Default Property:	log.index.interval.bytes
		// Importance:				medium
		"index.interval.bytes": {
			Type:        schema.TypeInt,
			Required:    false,
			ForceNew:    true,
			Description: "This setting controls how frequently Kafka adds an index entry to its offset index. The default setting ensures that we index a message roughly every 4096 bytes. More indexing allows reads to jump closer to the exact position in the log but makes the index larger. You probably don't need to change this.",
		},

		// leader.replication.throttled.replicas
		// A list of replicas for which log replication should be throttled on the leader
		// side. The list should describe a set of replicas in the form
		//
		//   [PartitionId]:[BrokerId],[PartitionId]:[BrokerId]:...
		//
		// or alternatively the wildcard '*' can be used to throttle all replicas
		// for this topic.
		//
		// Type:					list
		// Default:					""
		// Valid Values:			[partitionId]:[brokerId],[partitionId]:[brokerId],...
		// Server Default Property:	leader.replication.throttled.replicas
		// Importance:				medium
		"leader.replication.throttled.replicas": {
			Type:        schema.TypeList,
			Required:    false,
			ForceNew:    true,
			Description: "A list of replicas for which log replication should be throttled on the leader side. The list should describe a set of replicas in the form",
		},

		// max.compaction.lag.ms
		// The maximum time a message will remain ineligible for compaction in the log.
		// Only applicable for logs that are being compacted.
		//
		// Type:						long
		// Default:						9223372036854775807
		// Valid Values:				[1,...]
		// Server Default Property:		log.cleaner.max.compaction.lag.ms
		// Importance:					medium
		"max.compaction.lag.ms": {
			Type:        schema.TypeInt,
			Required:    false,
			ForceNew:    true,
			Description: "The maximum time a message will remain ineligible for compaction in the log. Only applicable for logs that are being compacted.",
		},

		// max.message.bytes
		// The largest record batch size allowed by Kafka (after compression if compression
		// is enabled). If this is increased and there are consumers older than 0.10.2, the
		// consumers' fetch size must also be increased so that they can fetch record batches
		// this large. In the latest message format version, records are always grouped into
		// batches for efficiency. In previous message format versions, uncompressed records are
		// not grouped into batches and this limit only applies to a single record in that case.
		//
		// Type:						int
		// Default:						1048588
		// Valid Values:				[0,...]
		// Server Default Property:		message.max.bytes
		// Importance:					medium
		"max.message.bytes": {
			Type:        schema.TypeInt,
			Required:    false,
			ForceNew:    true,
			Description: "The largest record batch size allowed by Kafka (after compression if compression is enabled). If this is increased and there are consumers older than 0.10.2, the consumers' fetch size must also be increased so that they can fetch record batches this large. In the latest message format version, records are always grouped into batches for efficiency. In previous message format versions, uncompressed records are not grouped into batches and this limit only applies to a single record in that case.",
		},

		// message.format.version
		// Specify the message format version the broker will use to append messages to the
		// logs. The value should be a valid ApiVersion. Some examples are: 0.8.2, 0.9.0.0,
		// 0.10.0, check ApiVersion for more details. By setting a particular message format
		// version, the user is certifying that all the existing messages on disk are smaller
		// or equal than the specified version. Setting this value incorrectly will cause
		// consumers with older versions to break as they will receive messages with a
		// format that they don't understand.
		//
		// Type:					string
		// Default:					2.7-IV2
		//
		// Valid Values:			[0.8.0, 0.8.1, 0.8.2, 0.9.0, 0.10.0-IV0, 0.10.0-IV1,
		// 							 0.10.1-IV0, 0.10.1-IV1, 0.10.1-IV2, 0.10.2-IV0, 0.11.0-IV0,
		// 							 0.11.0-IV1, 0.11.0-IV2, 1.0-IV0, 1.1-IV0, 2.0-IV0, 2.0-IV1,
		// 							 2.1-IV0, 2.1-IV1, 2.1-IV2, 2.2-IV0, 2.2-IV1, 2.3-IV0, 2.3-IV1,
		//  						 2.4-IV0, 2.4-IV1, 2.5-IV0, 2.6-IV0, 2.7-IV0, 2.7-IV1, 2.7-IV2]
		//
		// Server Default Property:		log.message.format.version
		// Importance:					medium
		"message.format.version": {
			Type:         schema.TypeString,
			Required:     false,
			ForceNew:     true,
			Description:  "Specify the message format version the broker will use to append messages to the logs. The value should be a valid ApiVersion. Some examples are: 0.8.2, 0.9.0.0, 0.10.0, check ApiVersion for more details. By setting a particular message format version, the user is certifying that all the existing messages on disk are smaller or equal than the specified version. Setting this value incorrectly will cause consumers with older versions to break as they will receive messages with a format that they don't understand.",
			ValidateFunc: ValidateMessageFormatVersion,
		},

		// message.timestamp.difference.max.ms
		// The maximum difference allowed between the timestamp when a broker receives a message and
		// the timestamp specified in the message. If message.timestamp.type=CreateTime, a message
		// will be rejected if the difference in timestamp exceeds this threshold. This configuration
		// is ignored if message.timestamp.type=LogAppendTime.
		//
		// Type:						long
		// Default:						9223372036854775807
		// Valid Values:				[0,...]
		// Server Default Property:		log.message.timestamp.difference.max.ms
		// Importance:					medium
		"message.timestamp.difference.max.ms": {
			Type:        schema.TypeInt,
			Required:    false,
			ForceNew:    true,
			Description: "The maximum difference allowed between the timestamp when a broker receives a message and the timestamp specified in the message. If message.timestamp.type=CreateTime, a message will be rejected if the difference in timestamp exceeds this threshold. This configuration is ignored if message.timestamp.type=LogAppendTime.",
		},

		// message.timestamp.type
		// Define whether the timestamp in the message is message create time or log append time.
		// The value should be either `CreateTime` or `LogAppendTime`
		//
		// Type:						string
		// Default:						CreateTime
		// Valid Values:				[CreateTime, LogAppendTime]
		// Server Default Property:		log.message.timestamp.type
		// Importance:					medium
		"message.timestamp.type": {
			Type:         schema.TypeString,
			Required:     false,
			ForceNew:     true,
			Description:  "Define whether the timestamp in the message is message create time or log append time. The value should be either `CreateTime` or `LogAppendTime`",
			ValidateFunc: ValidateMessageTimestampType,
		},

		// min.cleanable.dirty.ratio
		// This configuration controls how frequently the log compactor will attempt to clean the
		// log (assuming log compaction is enabled). By default we will avoid cleaning a log where
		// more than 50% of the log has been compacted. This ratio bounds the maximum space wasted
		// in the log by duplicates (at 50% at most 50% of the log could be duplicates). A higher
		// ratio will mean fewer, more efficient cleanings but will mean more wasted space in the
		// log. If the max.compaction.lag.ms or the min.compaction.lag.ms configurations are also
		// specified, then the log compactor considers the log to be eligible for compaction as
		// soon as either: (i) the dirty ratio threshold has been met and the log has had dirty
		// (uncompacted) records for at least the min.compaction.lag.ms duration, or (ii) if the
		// log has had dirty (uncompacted) records for at most the max.compaction.lag.ms period.
		//
		// Type:						double
		// Default:						0.5
		// Valid Values:				[0,...,1]
		// Server Default Property:		log.cleaner.min.cleanable.ratio
		// Importance:					medium
		"min.cleanable.dirty.ratio": {
			Type:        schema.TypeFloat,
			Required:    false,
			ForceNew:    true,
			Description: "This configuration controls how frequently the log compactor will attempt to clean the log (assuming log compaction is enabled). By default we will avoid cleaning a log where more than 50% of the log has been compacted. This ratio bounds the maximum space wasted in the log by duplicates (at 50% at most 50% of the log could be duplicates). A higher ratio will mean fewer, more efficient cleanings but will mean more wasted space in the log. If the max.compaction.lag.ms or the min.compaction.lag.ms configurations are also specified, then the log compactor considers the log to be eligible for compaction as soon as either: (i) the dirty ratio threshold has been met and the log has had dirty (uncompacted) records for at least the min.compaction.lag.ms duration, or (ii) if the log has had dirty (uncompacted) records for at most the max.compaction.lag.ms period.",
		},

		// min.compaction.lag.ms
		// The minimum time a message will remain uncompacted in the log. Only applicable for
		// logs that are being compacted.
		//
		// Type:						long
		// Default:						0
		// Valid Values:				[0,...]
		// Server Default Property:		log.cleaner.min.compaction.lag.ms
		// Importance:					medium
		"min.compaction.lag.ms": {
			Type:        schema.TypeInt,
			Required:    false,
			ForceNew:    true,
			Description: "The minimum time a message will remain uncompacted in the log. Only applicable for logs that are being compacted.",
		},

		// min.insync.replicas
		// When a producer sets acks to "all" (or "-1"), this configuration specifies the minimum
		// number of replicas that must acknowledge a write for the write to be considered successful.
		// If this minimum cannot be met, then the producer will raise an exception (either NotEnoughReplicas
		// or NotEnoughReplicasAfterAppend).
		//
		// When used together, min.insync.replicas and acks allow you to enforce greater durability
		// guarantees. A typical scenario would be to create a topic with a replication factor of 3,
		// set min.insync.replicas to 2, and produce with acks of "all". This will ensure that the producer
		// raises an exception if a majority of replicas do not receive a write.
		//
		// Type:						int
		// Default:						1
		// Valid Values:				[1,...]
		// Server Default Property:		min.insync.replicas
		// Importance:					medium
		"min.insync.replicas": {
			Type:     schema.TypeInt,
			Required: false,
			ForceNew: true,
			Description: `When a producer sets acks to \"all\" (or \"-1\"), this configuration specifies the minimum 
						   number of replicas that must acknowledge a write for the write to be considered successful. 
						   If this minimum cannot be met, then the producer will raise an exception (either NotEnoughReplicas 
						   or NotEnoughReplicasAfterAppend).
							
						   When used together, min.insync.replicas and acks allow you to enforce greater durability 
						   guarantees. A typical scenario would be to create a topic with a replication factor of 3, 
						   set min.insync.replicas to 2, and produce with acks of \"all\". This will ensure that the producer 
						   raises an exception if a majority of replicas do not receive a write.`,
		},

		// preallocate
		// True if we should preallocate the file on disk when creating a new log segment.
		//
		// Type:						boolean
		// Default:						false
		// Valid Values:
		// Server Default Property:		log.preallocate
		// Importance:					medium
		"preallocate": {
			Type:        schema.TypeBool,
			Required:    false,
			ForceNew:    true,
			Description: "True if we should preallocate the file on disk when creating a new log segment.",
		},

		// retention.bytes
		// This configuration controls the maximum size a partition (which consists of log segments)
		// can grow to before we will discard old log segments to free up space if we are using
		// the "delete" retention policy. By default there is no size limit only a time limit.
		// Since this limit is enforced at the partition level, multiply it by the number of
		// partitions to compute the topic retention in bytes.
		//
		// Type:						long
		// Default:						-1
		// Valid Values:
		// Server Default Property:		log.retention.bytes
		// Importance:					medium
		"retention.bytes": {
			Type:     schema.TypeInt,
			Required: false,
			ForceNew: true,
			Description: `This configuration controls the maximum size a partition (which consists of log segments)
						  can grow to before we will discard old log segments to free up space if we are using
						  the \"delete\" retention policy. By default there is no size limit only a time limit.
						  Since this limit is enforced at the partition level, multiply it by the number of
						  partitions to compute the topic retention in bytes.`,
		},

		// retention.ms
		// This configuration controls the maximum time we will retain a log before we will discard
		// old log segments to free up space if we are using the "delete" retention policy. This
		// represents an SLA on how soon consumers must read their data. If set to -1, no time
		// limit is applied.
		//
		// Type:						long
		// Default:						604800000 (7 days)
		// Valid Values:				[-1,...]
		// Server Default Property:		log.retention.ms
		// Importance:					medium
		"retention.ms": {
			Type:     schema.TypeInt,
			Required: false,
			ForceNew: true,
			Description: `This configuration controls the maximum time we will retain a log before we will discard
						  old log segments to free up space if we are using the "delete" retention policy. This
						  represents an SLA on how soon consumers must read their data. If set to -1, no time
						  limit is applied.`,
		},

		// segment.bytes
		// This configuration controls the segment file size for the log. Retention and cleaning is
		// always done a file at a time so a larger segment size means fewer files but less granular
		// control over retention.
		//
		// Type:						int
		// Default:						1073741824 (1 gibibyte)
		// Valid Values:				[14,...]
		// Server Default Property:		log.segment.bytes
		// Importance:					medium
		"segment.bytes": {
			Type:     schema.TypeInt,
			Required: false,
			ForceNew: true,
			Description: `This configuration controls the segment file size for the log. Retention and cleaning is
						  always done a file at a time so a larger segment size means fewer files but less granular
						  control over retention.`,
		},

		// segment.index.bytes
		// This configuration controls the size of the index that maps offsets to file positions.
		// We preallocate this index file and shrink it only after log rolls. You generally should
		// not need to change this setting.
		//
		// Type:						int
		// Default:						10485760 (10 mebibytes)
		// Valid Values:				[0,...]
		// Server Default Property:		log.index.size.max.bytes
		// Importance:					medium
		"segment.index.bytes": {
			Type:     schema.TypeInt,
			Required: false,
			ForceNew: true,
			Description: `This configuration controls the size of the index that maps offsets to file positions.
						  We preallocate this index file and shrink it only after log rolls. You generally should
						  not need to change this setting.`,
		},

		// segment.jitter.ms
		// The maximum random jitter subtracted from the scheduled segment roll time to avoid
		// thundering herds of segment rolling
		//
		// Type:						long
		// Default:						0
		// Valid Values:				[0,...]
		// Server Default Property:		log.roll.jitter.ms
		// Importance:					medium
		"segment.jitter.ms": {
			Type:     schema.TypeInt,
			Required: false,
			ForceNew: true,
			Description: `The maximum random jitter subtracted from the scheduled segment roll time to avoid
						  thundering herds of segment rolling`,
		},

		// segment.ms
		// This configuration controls the period of time after which Kafka will force the log
		// to roll even if the segment file isn't full to ensure that retention can delete or
		// compact old data.
		//
		// Type:						long
		// Default:						604800000 (7 days)
		// Valid Values:				[1,...]
		// Server Default Property:		log.roll.ms
		// Importance:					medium
		"segment.ms": {
			Type:     schema.TypeInt,
			Required: false,
			ForceNew: true,
			Description: `This configuration controls the period of time after which Kafka will force the log
						  to roll even if the segment file isn't full to ensure that retention can delete or
						  compact old data.`,
		},

		// unclean.leader.election.enable
		// Indicates whether to enable replicas not in the ISR set to be elected as leader as
		// a last resort, even though doing so may result in data loss.
		//
		// Type:						boolean
		// Default:						false
		// Valid Values:
		// Server Default Property:		unclean.leader.election.enable
		// Importance:					medium
		"unclean.leader.election.enable": {
			Type:     schema.TypeBool,
			Required: false,
			ForceNew: true,
			Description: `Indicates whether to enable replicas not in the ISR set to be elected as leader as
						  a last resort, even though doing so may result in data loss.`,
		},

		// message.downconversion.enable
		// This configuration controls whether down-conversion of message formats is enabled to
		// satisfy consume requests. When set to false, broker will not perform down-conversion
		// for consumers expecting an older message format. The broker responds with
		// UNSUPPORTED_VERSION error for consume requests from such older clients. This
		// configurationdoes not apply to any message format conversion that might be
		// required for replication to followers.
		//
		// Type:	boolean
		// Default:	true
		// Valid Values:
		// Server Default Property:	log.message.downconversion.enable
		// Importance:	low
		"message.downconversion.enable": {
			Type:     schema.TypeBool,
			Required: false,
			ForceNew: true,
			Description: `This configuration controls whether down-conversion of message formats is enabled 
						  to satisfy consume requests. When set to false, broker will not perform down-conversion 
						  for consumers expecting an older message format. The broker responds with UNSUPPORTED_VERSION 
						  error for consume requests from such older clients. This configurationdoes not 
						  apply to any message format conversion that might be required for replication to followers.`,
		},
	}
}

func ValidateKafkaCleanupPolicy(val interface{}, key string) (warns []string, errs []error) {
	if val != "compact" && val != "delete" {
		return []string{"Error: cleanup.policy value is incorrect. Must be one of [compact, delete]"}, []error{nil}
	}
	return nil, nil
}

func ValidateKafkaCompressionType(val interface{}, key string) (warns []string, errs []error) {
	if val != "uncompressed" && val != "zstd" && val != "lz4" && val != "snappy" && val != "gzip" && val != "producer" {
		return []string{"Error: compression.type value is incorrect. Must be one of [\"uncompressed\", \"zstd\", \"lz4\", \"snappy\", \"gzip\", \"producer\"]"}, []error{nil}
	}
	return nil, nil
}

func ValidateMessageFormatVersion(val interface{}, key string) (warns []string, errs []error) {
	validVersions := []string{"0.8.0", "0.8.1", "0.8.2", "0.9.0", "0.10.0-IV0", "0.10.0-IV1",
		"0.10.1-IV0", "0.10.1-IV1", "0.10.1-IV2", "0.10.2-IV0", "0.11.0-IV0",
		"0.11.0-IV1", "0.11.0-IV2", "1.0-IV0", "1.1-IV0", "2.0-IV0", "2.0-IV1",
		"2.1-IV0", "2.1-IV1", "2.1-IV2", "2.2-IV0", "2.2-IV1", "2.3-IV0", "2.3-IV1",
		"2.4-IV0", "2.4-IV1", "2.5-IV0", "2.6-IV0", "2.7-IV0", "2.7-IV1", "2.7-IV2"}
	for _, knownVersion := range validVersions {
		if val == knownVersion {
			return []string{"Error: known-version value is incorrect."}, []error{nil}
		}
	}
	return nil, nil
}

func ValidateMessageTimestampType(val interface{}, key string) (warns []string, errs []error) {
	if val != "CreateTime" && val != "LogAppendTime" {
		return []string{"Error: message.timestamp.type value is incorrect. Must be one of [\"CreateTime\", \"LogAppendTime\"]"}, []error{nil}
	}
	return nil, nil
}
