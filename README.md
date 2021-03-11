TerraForm Provider for Kafka

Here is Kafka.  

## Kafka Operation Support
We will take from the toolsets found in the Kafka Distribution Bin 
directory for what will be exposed via the TerraForm Provider:

    Support for the following toolsets (from Kafka CLI/Bin):
        - connect-distributed.sh
        - connect-mirror-maker.sh
        - connect-standalone.sh
        - kafka-acls.sh
        - kafka-broker-api-versions.sh
        - kafka-configs.sh
        - kafka-console-consumer.sh
        - kafka-console-producer.sh
        - kafka-consumer-groups.sh
        - kafka-consumer-perf-test.sh
        - kafka-delegation-tokens.sh
        - kafka-delete-records.sh
        - kafka-dump-log.sh
        - kafka-features.sh
        - kafka-leader-election.sh
        - kafka-log-dirs.sh
        - kafka-mirror-maker.sh
        - kafka-preferred-replica-election.sh
        - kafka-producer-perf-test.sh
        - kafka-reassign-partitions.sh
        - kafka-replica-verification.sh
        - kafka-run-class.sh
        - kafka-server-start.sh
        - kafka-server-stop.sh
        - kafka-streams-application-reset.sh
        - kafka-topics.sh
        - kafka-verifiable-consumer.sh
        - kafka-verifiable-producer.sh
        - trogdor.sh
        - zookeeper-security-migration.sh
        - zookeeper-server-start.sh
        - zookeeper-server-stop.sh
        - zookeeper-shell.sh

## Kafka-CLI Tools

### ./bin/kafka-acls.sh

This tool helps to manage acls on kafka.
Option                                   Description                            
------                                   -----------                            
--add                                               
    Indicates you are trying to add ACLs.  

--allow-host <String: allow-host>                   
    Host from which principals listed in -- allow-principal will have access. 
    If you have specified --allow-principal then the default for this option 
    will be set to * which allows access from all hosts.                      

--allow-principal <String: allow-principal>           
    principal is in principalType:name format. Note that principalType must be supported by 
    the Authorizer being used. For example, User:* is the wild card indicating all users.

--authorizer <String: authorizer>        
    Fully qualified class name of the authorizer, defaults to kafka security.authorizer.AclAuthorizer.

--authorizer-properties <String: authorizer-properties>           
    REQUIRED: properties required to configure an instance of Authorizer. These are key=val pairs. 
    For the default authorizer the example values are: zookeeper. connect=localhost:2181

--bootstrap-server <String: server to connect to>    
    A list of host/port pairs to use for establishing the connection to the Kafka cluster. 
    This list should be in the form host1:port1,host2: port2,... This config is required 
    for acl management using admin client API.                          

--cluster                                
    Add/Remove cluster ACLs.               

--command-config [String: command-config]       
    A property file containing configs to be passed to Admin Client.           

--consumer 
    Convenience option to add/remove ACLs for consumer role. This will generate 
    ACLs that allows READ, DESCRIBE on topic and READ on group. 

--delegation-token <String: delegation-token>  
    Delegation token to which ACLs should be added or removed. A value of * indicates 
    ACL should apply to all tokens.                              

--deny-host <String: deny-host>          
    Host from which principals listed in -- deny-principal will be denied access. 
    If you have specified --deny-principal then the default for this option will 
    be set to * which denies access from all hosts.               

--deny-principal <String: deny-principal>         
    principal is in principalType:name format. By default anyone not added through 
    --allow-principal is denied access. You only need to use this option as negation 
    to already allowed set. Note that principalType must be supported by the Authorizer 
    being used. For example if you wanted to allow access to all users in the system but 
    not test-user you can define an ACL that allows access to User:* and specify 
    --deny-principal=User:test@EXAMPLE.COM. 
    AND PLEASE REMEMBER DENY RULES TAKES PRECEDENCE OVER ALLOW RULES.

--force 
    Assume Yes to all queries and do not prompt.

--group <String: group> 
    Consumer Group to which the ACLs should be added or removed. A value of * indicates 
    the ACLs should apply to all groups.

--help
    Print usage information.               

--idempotent
    Enable idempotence for the producer. This should be used in combination 
    with the --producer option. 
    Note that idempotence is enabled automatically if the producer is authorized to a particular transactional-id.                    

--list 
    List ACLs for the specified resource, use --topic <topic> or --group <group> 
    or --cluster to specify a resource.                            

--operation <String>
    Operation that is being allowed or denied. Valid operation names are: 
        Describe, DescribeConfigs, Alter, IdempotentWrite, Read, Delete, 
        Create, ClusterAction, All, Write, AlterConfigs (default: All)

--principal [String: principal]          
    List ACLs for the specified principal. principal is in principalType:name format. Note that 
    principalType must be supported by the Authorizer being used. Multiple
    
--principal option    
    can be passed.                       

--producer                               
    Convenience option to add/remove ACLs for producer role. This will generate ACLs that allows 
    WRITE, DESCRIBE and CREATE on topic.

--remove
    Indicates you are trying to remove ACLs.                                

--resource-pattern-type <ANY|MATCH|LITERAL|PREFIXED>                  
    The type of the resource pattern or pattern filter. When adding acls, this should be a specific 
    pattern type, e.g. 'literal' or 'prefixed'. When listing or removing acls, a specific pattern 
    type can be used to list or remove acls from specific resource patterns, or use the filter values 
    of 'any' or 'match', where 'any' will match any pattern type, but will match the resource name 
    exactly, where as 'match' will perform pattern matching to list or remove all acls that affect 
    the supplied resource(s). WARNING: 'match', when used in combination with the '--remove' switch, 
    should be used with care. (default: LITERAL)

--topic <String: topic>                  
    topic to which ACLs should be added or removed. A value of * indicates ACL should apply to all topics.
              
--transactional-id <String:transactional-id>              
    The transactionalId to which ACLs should be added or removed. A value of * indicates the ACLs 
    should apply to all transactionalIds.             

--version                                
    Display Kafka version.

--zk-tls-config-file <String:configuration>            
    Identifies the file where ZooKeeper Authorizer ZooKeeper TLS client TLS connectivity 
    properties for the authorizer are defined.  Any properties other than the 
    following (with or without an "authorizer." prefix) are ignored: 
        zookeeper.clientCnxnSocket, zookeeper.ssl.cipher.suites, zookeeper.ssl.client.enable, zookeeper.ssl.crl.enable,    
        zookeeper.ssl.enabled.protocols, zookeeper.ssl.endpoint.identification.algorithm, 
        zookeeper.ssl.keystore.location, zookeeper.ssl.keystore.password, zookeeper.ssl.keystore.type, 
        zookeeper.ssl.ocsp.enable, zookeeper.ssl.protocol,      
        zookeeper.ssl.truststore.location, zookeeper.ssl.truststore.password, zookeeper.ssl.truststore.type. 
                                           
    Note that if SASL is not configured and zookeeper.set.acl is supposed to be true due to mutual certificate
     authentication being used then it is necessary to explicitly specify --  authorizer-properties zookeeper.set.acl=true                             



### ./bin/kafka-cluster.sh
    usage: kafka-cluster [-h] {cluster-id,unregister} ...

    The Kafka cluster tool.

    positional arguments:
    {cluster-id,unregister}
        cluster-id           Get information about the ID of a cluster.
        unregister           Unregister a broker..

    optional arguments:
    -h, --help             show this help message and exit



### ./bin/kafka-configs.sh
    This tool helps to manipulate and describe entity config for a topic, client, user, broker or ip
    Option                                 Description                            
    ------                                 -----------                            
    --add-config <String>                 Key Value pairs of configs to add. Square brackets can be used to group values which contain commas: 
                                            'k1=v1,k2=[v1,v2,v2],k3=v3'. 
                                          
                                          The following is a list of valid configurations:   
                                          For entity-type 'topics':            
                                       	    cleanup.policy                        
                                       	    compression.type                      
                                       	    delete.retention.ms                   
                                       	    file.delete.delay.ms                  
                                       	    flush.messages                        
                                       	    flush.ms                              
                                       	    follower.replication.throttled.replicas                             
                                       	    index.interval.bytes                  
                                       	    leader.replication.throttled.replicas 
                                       	    max.compaction.lag.ms                 
                                       	    max.message.bytes                     
                                       	    message.downconversion.enable         
                                       	    message.format.version                
                                       	    message.timestamp.difference.max.ms   
                                       	    message.timestamp.type                
                                            min.cleanable.dirty.ratio             
                                            min.compaction.lag.ms                 
                                            min.insync.replicas                   
                                            preallocate                           
                                            retention.bytes                       
                                            retention.ms                          
                                            segment.bytes                         
                                            segment.index.bytes                   
                                            segment.jitter.ms                     
                                            segment.ms                            
                                            unclean.leader.election.enable        
                                       For entity-type 'brokers':             
                                       	advertised.listeners                  
                                       	background.threads                    
                                       	compression.type                      
                                       	follower.replication.throttled.rate   
                                       	leader.replication.throttled.rate     
                                       	listener.security.protocol.map        
                                       	listeners                             
                                       	log.cleaner.backoff.ms                
                                       	log.cleaner.dedupe.buffer.size        
                                       	log.cleaner.delete.retention.ms       
                                       	log.cleaner.io.buffer.load.factor     
                                       	log.cleaner.io.buffer.size            
                                       	log.cleaner.io.max.bytes.per.second   
                                       	log.cleaner.max.compaction.lag.ms     
                                       	log.cleaner.min.cleanable.ratio       
                                       	log.cleaner.min.compaction.lag.ms     
                                       	log.cleaner.threads                   
                                       	log.cleanup.policy                    
                                       	log.flush.interval.messages           
                                       	log.flush.interval.ms                 
                                       	log.index.interval.bytes              
                                       	log.index.size.max.bytes              
                                       	log.message.downconversion.enable     
                                       	log.message.timestamp.difference.max.ms                                   
                                       	log.message.timestamp.type            
                                       	log.preallocate                       
                                       	log.retention.bytes                   
                                       	log.retention.ms                      
                                       	log.roll.jitter.ms                    
                                       	log.roll.ms                           
                                       	log.segment.bytes                     
                                       	log.segment.delete.delay.ms           
                                       	max.connection.creation.rate          
                                       	max.connections                       
                                       	max.connections.per.ip                
                                       	max.connections.per.ip.overrides      
                                       	message.max.bytes                     
                                       	metric.reporters                      
                                       	min.insync.replicas                   
                                       	num.io.threads                        
                                       	num.network.threads                   
                                       	num.recovery.threads.per.data.dir     
                                       	num.replica.fetchers                  
                                       	principal.builder.class               
                                       	replica.alter.log.dirs.io.max.bytes.per.second                           
                                       	sasl.enabled.mechanisms               
                                       	sasl.jaas.config                      
                                       	sasl.kerberos.kinit.cmd               
                                       	sasl.kerberos.min.time.before.relogin 
                                       	sasl.kerberos.principal.to.local.rules
                                       	sasl.kerberos.service.name            
                                       	sasl.kerberos.ticket.renew.jitter     
                                       	sasl.kerberos.ticket.renew.window.factor                               
                                       	sasl.login.refresh.buffer.seconds     
                                       	sasl.login.refresh.min.period.seconds 
                                       	sasl.login.refresh.window.factor      
                                       	sasl.login.refresh.window.jitter      
                                       	sasl.mechanism.inter.broker.protocol  
                                       	ssl.cipher.suites                     
                                       	ssl.client.auth                       
                                       	ssl.enabled.protocols                 
                                       	ssl.endpoint.identification.algorithm 
                                       	ssl.engine.factory.class              
                                       	ssl.key.password                      
                                       	ssl.keymanager.algorithm              
                                       	ssl.keystore.certificate.chain        
                                       	ssl.keystore.key                      
                                       	ssl.keystore.location                 
                                       	ssl.keystore.password                 
                                       	ssl.keystore.type                     
                                       	ssl.protocol                          
                                       	ssl.provider                          
                                       	ssl.secure.random.implementation      
                                       	ssl.trustmanager.algorithm            
                                       	ssl.truststore.certificates           
                                       	ssl.truststore.location               
                                       	ssl.truststore.password               
                                       	ssl.truststore.type                   
                                       	unclean.leader.election.enable        
                                       For entity-type 'users':               
                                       	SCRAM-SHA-256                         
                                       	SCRAM-SHA-512                         
                                       	consumer_byte_rate                    
                                       	controller_mutation_rate              
                                       	producer_byte_rate                    
                                       	request_percentage                    
                                       For entity-type 'clients':             
                                       	consumer_byte_rate                    
                                       	controller_mutation_rate              
                                       	producer_byte_rate                    
                                       	request_percentage                    
                                       For entity-type 'ips':                 
                                       	connection_creation_rate              
                                       Entity types 'users' and 'clients' may be specified together to update config for clients of a specific user.

--add-config-file <String>             
  Path to a properties file with configs to add. See add-config for a list of valid configurations.                

--all                                  
  List all configs for the given topic, broker, or broker-logger entity (includes static configuration when the entity type is brokers)          

--alter 
  Alter the configuration for the entity.

--bootstrap-server <String: server to connect to>  
  The Kafka server to connect to. This is required for describing and altering broker configs.             

--broker <String>                      
  The broker's ID.                       

--broker-defaults                      
  The config defaults for all brokers.   

--broker-logger <String>               
  The broker's ID for its logger config. 

--client <String>                      
  The client's ID.                       

--client-defaults                      
  The config defaults for all clients.   

--command-config <String: command config property file>      
    Property file containing configs to be passed to Admin Client. This is 
    used only with --bootstrap-server option for describing and altering broker configs.

--delete-config <String>
    config keys to remove 'k1,k2'          

--describe                             
  List configs for the given entity.     

--entity-default                       
  Default entity name for clients/users/brokers/ips (applies to corresponding entity type in command line)                        

--entity-name <String> 
  Name of entity (topic name/client id/user principal name/broker id/ip) 

--entity-type <String>                 
  Type of entity (topics/clients/users/brokers/broker-loggers/ips)                         

--force                                
  Suppress console prompts               

--help                                 
  Print usage information.               

--ip <String>                          
  The IP address.                        

--ip-defaults                          
  The config defaults for all IPs.       

--topic <String>                       
  The topic's name.                      

--user <String>                        
  The user's principal name.             

--user-defaults                        
  The config defaults for all users.     

--version                              
  Display Kafka version.                 

--zk-tls-config-file <String: ZooKeeper TLS configuration>
    Identifies the file where ZooKeeper client TLS connectivity properties are defined.  
    Any properties other than 
        zookeeper.clientCnxnSocket,     
        zookeeper.ssl.cipher.suites,         
        zookeeper.ssl.client.enable,         
        zookeeper.ssl.crl.enable, 
        zookeeper.ssl.enabled.protocols, 
        zookeeper.ssl.endpoint.identification.algorithm,   
        zookeeper.ssl.keystore.location,     
        zookeeper.ssl.keystore.password,     
        zookeeper.ssl.keystore.type,         
        zookeeper.ssl.ocsp.enable, 
        zookeeper.ssl.protocol, 
        zookeeper.ssl.truststore.location, 
        zookeeper.ssl.truststore.password, 
        zookeeper.ssl.truststore.type are ignored.

--zookeeper <String: urls>
    DEPRECATED. The connection string for the zookeeper connection in the form host:port. 
    Multiple URLS can be given to allow fail-over. 
    Replaced by --bootstrap-server, 
    REQUIRED unless --bootstrap-server is given.  



### ./bin/kafka-console-consumer.sh
This tool helps to read data from Kafka topics and outputs it to standard output.
Option                                   Description                            
------                                   -----------                            
--bootstrap-server <String: server to connect to>    
  REQUIRED: The server(s) to connect to. 
                                                                     
--consumer-property <String: consumer_prop>
  A mechanism to pass user-defined properties in the form key=value to the consumer.                        

--consumer.config <String: config file>  
  Consumer config properties file. Note that [consumer-property] takes precedence over this config.         

--enable-systest-events                  
  Log lifecycle events of the consumer in addition to logging consumed messages. 
  (This is specific for system tests.)                       

--formatter <String: class>              
  The name of a class to use for formatting kafka messages for display. 
  (default: kafka.tools. DefaultMessageFormatter)             

--from-beginning                         
  If the consumer does not already have an established offset to consume from, start with 
  the earliest message present in the log rather than the latest message.             

--group <String: consumer group id>      
  The consumer group id of the consumer. 

--help                                   
  Print usage information.               

--isolation-level <String>
  Set to read_committed in order to filter out transactional messages which are not 
  committed. Set to read_uncommitted to read all messages. (default: read_uncommitted)

--key-deserializer <String: deserializer for key>                                                         

--max-messages <Integer: num_messages>   
  The maximum number of messages to consume before exiting. If not set, consumption is continual.

--offset <String: consume offset>
  The offset id to consume from (a non-negative number), or 'earliest' which means from beginning, 
  or 'latest' which means from end (default: latest)                    

--partition <Integer: partition>         
  The partition to consume from. Consumption starts from the end of the partition 
  unless '--offset' is specified.                           


--property <String: prop>
  The properties to initialize the message formatter. Default properties include:                  
                                          print.timestamp=true|false            
                                          print.key=true|false                  
                                          print.offset=true|false               
                                          print.partition=true|false            
                                          print.headers=true|false              
                                          print.value=true|false                
                                          key.separator=<key.separator>         
                                          line.separator=<line.separator>       
                                          headers.separator=<line.separator>    
                                          null.literal=<null.literal>           
                                          key.deserializer=<key.deserializer>   
                                          value.deserializer=<value.deserializer>                        
                                          header.deserializer=<header.deserializer>                        
                                         Users can also pass in customized      
                                           properties for their formatter; more 
                                           specifically, users can pass in      
                                           properties keyed with 'key.          
                                           deserializer.', 'value.              
                                           deserializer.' and 'headers.         
                                           deserializer.' prefixes to configure 
                                           their deserializers.

--skip-message-on-error                  
  If there is an error when processing a message, skip it instead of halt.    

--timeout-ms <Integer: timeout_ms>
  If specified, exit if no message is available for consumption for the specified interval.                  

--topic <String: topic> 
  The topic id to consume on.

--value-deserializer <String: deserializer for values>                                                   
                                                        
--version
  Display Kafka version.                 

--whitelist <String: whitelist>
  Regular expression specifying whitelist of topics to include for consumption.                         



### ./bin/kafka-console-producer.sh 
Missing required option(s) [bootstrap-server]
Option                                   Description                            
------                                   -----------                            
--batch-size <Integer: size>
Number of messages to send in a single batch if they are not being sent 
synchronously. (default: 200)

--bootstrap-server <String: server to connect to>
REQUIRED unless --broker-list
deprecated) is specified. The server (s) to connect to. The broker list 
string in the form HOST1:PORT1,HOST2:PORT2.                               

--broker-list <String: broker-list>      
DEPRECATED, use --bootstrap-server     
instead; ignored if --bootstrap-server is specified.  The broker list string 
in the form HOST1:PORT1,HOST2:PORT2.                         

--compression-codec [String:The compression codec: either 'none',compression-codec]
  'gzip', 'snappy', 'lz4', or 'zstd'. If specified without value, then it defaults to 'gzip'                   

--help                                   
  Print usage information.               

--line-reader <String: reader_class>     
  The class name of the class to use for reading lines from standard in. By default 
  each line is read as a separate message. (default: kafka.tools.ConsoleProducer$LineMessageReader)   

--max-block-ms <Long: max block on send>       
  The max time that the producer will block for during a send request (default: 60000)                     

--max-memory-bytes <Long: total memory in bytes>   
  The total memory used by the producer to buffer records waiting to be sent to the server. (default: 33554432)

--max-partition-memory-bytes <Long: memory in bytes per partition>
  The buffer size allocated for a partition. When records are received which are smaller than this size the 
  producer will attempt to optimistically group them together until this size is reached. (default: 16384)

--message-send-max-retries <Integer>     
  Brokers can fail receiving the message for multiple reasons, and being unavailable transiently is just one of 
  them. This property specifies the number of retries before the producer give up and drop this 
  message. (default: 3)

--metadata-expiry-ms <Long: metadata expiration interval>
  The period of time in milliseconds after which we force a refresh of metadata even if we haven't 
  seen any leadership changes. (default: 300000)

--producer-property <String: producer_prop>
  A mechanism to pass user-defined properties in the form key=value to the producer.                        

--producer.config <String: config file>  
  Producer config properties file. Note that [producer-property] takes precedence over this config.

--property <String: prop>                
  A mechanism to pass user-defined properties in the form key=value to the message reader. 
  This allows custom configuration for a user-defined message reader. 
  Default properties include: 
    parse.key=true|false                  
    key.separator=<key.separator>         
    ignore.error=true|false               

--request-required-acks <String: request required acks> 
  The required acks of the producer requests (default: 1)      

--request-timeout-ms <Integer: request timeout ms>
  The ack timeout of the producer requests. Value must be non-negative and non-zero (default: 1500)         

--retry-backoff-ms <Integer>             
  Before each retry, the producer refreshes the metadata of relevant topics. Since leader election takes 
  a bit of time, this property specifies the amount of time that the producer waits before refreshing the 
  metadata. (default: 100)         

--socket-buffer-size <Integer: size>     
  The size of the tcp RECV size. (default: 102400)                    

--sync                                   
  If set message send requests to the brokers are synchronously, one at a time as they arrive.                 

--timeout <Integer: timeout_ms>          
  If set and the producer is running in asynchronous mode, this gives the maximum amount of time a 
  message will queue awaiting sufficient batch size. The value is given in ms. (default: 1000)

--topic <String: topic>
  REQUIRED: The topic id to produce messages to.                         

--version                                
  Display Kafka version.   








### ./bin/kafka-consumer-groups.sh 
Missing required argument "[bootstrap-server]"
Option                                  Description                            
------                                  -----------                            

--all-groups                            
  Apply to all consumer groups.          

--all-topics 
  Consider all topics assigned to a group in the `reset-offsets` process.

--bootstrap-server <String: server to connect to>
  REQUIRED: The server(s) to connect to. 
                                                                 
--by-duration <String: duration>
  Reset offsets to offset by duration from current timestamp. Format: 'PnDTnHnMnS'                  

--command-config <String: command config property file>
  Property file containing configs to be passed to Admin Client and Consumer. 

--delete 
  Pass in groups to delete topic partition offsets and ownership information over the entire consumer 
                                          group. For instance --group g1 --    
                                          group g2                             
--delete-offsets                        Delete offsets of consumer group.      
                                          Supports one consumer group at the   
                                          time, and multiple topics.           
--describe                              Describe consumer group and list       
                                          offset lag (number of messages not   
                                          yet processed) related to given      
                                          group.                               
--dry-run                               Only show results without executing    
                                          changes on Consumer Groups.          
                                          Supported operations: reset-offsets. 
--execute                               Execute operation. Supported           
                                          operations: reset-offsets.           
--export                                Export operation execution to a CSV    
                                          file. Supported operations: reset-   
                                          offsets.                             
--from-file <String: path to CSV file>  Reset offsets to values defined in CSV 
                                          file.                                

--group <String: consumer group>        The consumer group we wish to act on.  
--help                                  Print usage information.               
--list                                  List all consumer groups.              
--members                               Describe members of the group. This    
                                          option may be used with '--describe' 
                                          and '--bootstrap-server' options     
                                          only.                                
                                        Example: --bootstrap-server localhost: 
                                          9092 --describe --group group1 --    
                                          members                              
--offsets                               Describe the group and list all topic  
                                          partitions in the group along with   
                                          their offset lag. This is the        
                                          default sub-action of and may be     
                                          used with '--describe' and '--       
                                          bootstrap-server' options only.      
                                        Example: --bootstrap-server localhost: 
                                          9092 --describe --group group1 --    
                                          offsets                              
--reset-offsets                         Reset offsets of consumer group.       
                                          Supports one consumer group at the   
                                          time, and instances should be        
                                          inactive                             
                                        Has 2 execution options: --dry-run     
                                          (the default) to plan which offsets  
                                          to reset, and --execute to update    
                                          the offsets. Additionally, the --    
                                          export option is used to export the  
                                          results to a CSV format.             
                                        You must choose one of the following   
                                          reset specifications: --to-datetime, 
                                          --by-period, --to-earliest, --to-    
                                          latest, --shift-by, --from-file, --  
                                          to-current.                          
                                        To define the scope use --all-topics   
                                          or --topic. One scope must be        
                                          specified unless you use '--from-    
                                          file'.                               
--shift-by <Long: number-of-offsets>    Reset offsets shifting current offset  
                                          by 'n', where 'n' can be positive or 
                                          negative.                            
--state [String]                        When specified with '--describe',      
                                          includes the state of the group.     
                                        Example: --bootstrap-server localhost: 
                                          9092 --describe --group group1 --    
                                          state                                
                                        When specified with '--list', it       
                                          displays the state of all groups. It 
                                          can also be used to list groups with 
                                          specific states.                     
                                        Example: --bootstrap-server localhost: 
                                          9092 --list --state stable,empty     
                                        This option may be used with '--       
                                          describe', '--list' and '--bootstrap-
                                          server' options only.                
--timeout <Long: timeout (ms)>          The timeout that can be set for some   
                                          use cases. For example, it can be    
                                          used when describing the group to    
                                          specify the maximum amount of time   
                                          in milliseconds to wait before the   
                                          group stabilizes (when the group is  
                                          just created, or is going through    
                                          some changes). (default: 5000)       
--to-current                            Reset offsets to current offset.       
--to-datetime <String: datetime>        Reset offsets to offset from datetime. 
                                          Format: 'YYYY-MM-DDTHH:mm:SS.sss'    
--to-earliest                           Reset offsets to earliest offset.      
--to-latest                             Reset offsets to latest offset.        
--to-offset <Long: offset>              Reset offsets to a specific offset.    
--topic <String: topic>                 The topic whose consumer group         
                                          information should be deleted or     
                                          topic whose should be included in    
                                          the reset offset process. In `reset- 
                                          offsets` case, partitions can be     
                                          specified using this format: `topic1:
                                          0,1,2`, where 0,1,2 are the          
                                          partition to be included in the      
                                          process. Reset-offsets also supports 
                                          multiple topic inputs.               
--verbose                               Provide additional information, if     
                                          any, when describing the group. This 
                                          option may be used with '--          
                                          offsets'/'--members'/'--state' and   
                                          '--bootstrap-server' options only.   
                                        Example: --bootstrap-server localhost: 
                                          9092 --describe --group group1 --    
                                          members --verbose                    
--version                               Display Kafka version.      




./bin/kafka-consumer-perf-test
SLF4J: Actual binding is of type [org.slf4j.impl.Log4jLoggerFactory]
Missing required option(s) [bootstrap-server]
Option                                   Description                            
------                                   -----------                            
--bootstrap-server <String: server to    REQUIRED unless --broker-list          
  connect to>                              (deprecated) is specified. The server
                                           (s) to connect to.                   
--broker-list <String: broker-list>      DEPRECATED, use --bootstrap-server     
                                           instead; ignored if --bootstrap-     
                                           server is specified.  The broker     
                                           list string in the form HOST1:PORT1, 
                                           HOST2:PORT2.                         
--consumer.config <String: config file>  Consumer config properties file.       
--date-format <String: date format>      The date format to use for formatting  
                                           the time field. See java.text.       
                                           SimpleDateFormat for options.        
                                           (default: yyyy-MM-dd HH:mm:ss:SSS)   
--fetch-size <Integer: size>             The amount of data to fetch in a       
                                           single request. (default: 1048576)   
--from-latest                            If the consumer does not already have  
                                           an established offset to consume     
                                           from, start with the latest message  
                                           present in the log rather than the   
                                           earliest message.                    
--group <String: gid>                    The group id to consume on. (default:  
                                           perf-consumer-79410)                 
--help                                   Print usage information.               
--hide-header                            If set, skips printing the header for  
                                           the stats                            
--messages <Long: count>                 REQUIRED: The number of messages to    
                                           send or consume                      
--num-fetch-threads <Integer: count>     DEPRECATED AND IGNORED: Number of      
                                           fetcher threads. (default: 1)        
--print-metrics                          Print out the metrics.                 
--reporting-interval <Integer:           Interval in milliseconds at which to   
  interval_ms>                             print progress info. (default: 5000) 
--show-detailed-stats                    If set, stats are reported for each    
                                           reporting interval as configured by  
                                           reporting-interval                   
--socket-buffer-size <Integer: size>     The size of the tcp RECV size.         
                                           (default: 2097152)                   
--threads <Integer: count>               DEPRECATED AND IGNORED: Number of      
                                           processing threads. (default: 10)    
--timeout [Long: milliseconds]           The maximum allowed time in            
                                           milliseconds between returned        
                                           records. (default: 10000)            
--topic <String: topic>                  REQUIRED: The topic to consume from.   
--version                                Display Kafka version.      




### ./bin/kafka-delegation-tokens.sh
This tool helps to create, renew, expire, or describe delegation tokens.
Option                         Description                                     
------                         -----------                                     
--bootstrap-server <String>    REQUIRED: server(s) to use for bootstrapping.   
--command-config <String>      REQUIRED: A property file containing configs to 
                                 be passed to Admin Client. Token management   
                                 operations are allowed in secure mode only.   
                                 This config file is used to pass security     
                                 related configs.                              
--create                       Create a new delegation token. Use --renewer-   
                                 principal option to pass renewers principals. 
--describe                     Describe delegation tokens for the given        
                                 principals. Use --owner-principal to pass     
                                 owner/renewer principals. If --owner-principal
                                 option is not supplied, all the user owned    
                                 tokens and tokens where user have Describe    
                                 permission will be returned.                  
--expire                       Expire delegation token. Use --expiry-time-     
                                 period option to expire the token.            
--expiry-time-period [Long]    Expiry time period in milliseconds. If the value
                                 is -1, then the token will get invalidated    
                                 immediately.                                  
--help                         Print usage information.                        
--hmac [String]                HMAC of the delegation token                    
--max-life-time-period [Long]  Max life period for the token in milliseconds.  
                                 If the value is -1, then token max life time  
                                 will default to a server side config value    
                                 (delegation.token.max.lifetime.ms).           
--owner-principal [String]     owner is a kafka principal. It is should be in  
                                 principalType:name format.                    
--renew                        Renew delegation token. Use --renew-time-period 
                                 option to set renew time period.              
--renew-time-period [Long]     Renew time period in milliseconds. If the value 
                                 is -1, then the renew time period will default
                                 to a server side config value (delegation.    
                                 token.expiry.time.ms).                        
--renewer-principal [String]   renewer is a kafka principal. It is should be in
                                 principalType:name format.                    
--version                      Display Kafka version.    




### ./bin/kafka-delete-records.sh  
This tool helps to delete records of the given partitions down to the specified offset.
Option                                 Description                           
------                                 -----------                           
--bootstrap-server <String: server(s)  REQUIRED: The server to connect to.   
  to use for bootstrapping>                                                  
--command-config <String: command      A property file containing configs to 
  config property file path>             be passed to Admin Client.          
--help                                 Print usage information.              
--offset-json-file <String: Offset     REQUIRED: The JSON file with offset   
  json file path>                        per partition. The format to use is:
                                       {"partitions":                        
                                         [{"topic": "foo", "partition": 1,   
                                         "offset": 1}],                      
                                        "version":1                          
                                       }                                     
--version                              Display Kafka version.                



### ./bin/kafka-dump-log.sh
This tool helps to parse a log file and dump its contents to the console, useful for debugging a seemingly corrupt log segment.
Option                               Description                                
------                               -----------                                
--deep-iteration                     if set, uses deep instead of shallow       
                                       iteration. Automatically set if print-   
                                       data-log is enabled.                     
--files <String: file1, file2, ...>  REQUIRED: The comma separated list of data 
                                       and index log files to be dumped.        
--help                               Print usage information.                   
--index-sanity-check                 if set, just checks the index sanity       
                                       without printing its content. This is    
                                       the same check that is executed on       
                                       broker startup to determine if an index  
                                       needs rebuilding or not.                 
--key-decoder-class [String]         if set, used to deserialize the keys. This 
                                       class should implement kafka.serializer. 
                                       Decoder trait. Custom jar should be      
                                       available in kafka/libs directory.       
                                       (default: kafka.serializer.StringDecoder)
--max-message-size <Integer: size>   Size of largest message. (default: 5242880)
--offsets-decoder                    if set, log data will be parsed as offset  
                                       data from the __consumer_offsets topic.  
--print-data-log                     if set, printing the messages content when 
                                       dumping data logs. Automatically set if  
                                       any decoder option is specified.         
--transaction-log-decoder            if set, log data will be parsed as         
                                       transaction metadata from the            
                                       __transaction_state topic.               
--value-decoder-class [String]       if set, used to deserialize the messages.  
                                       This class should implement kafka.       
                                       serializer.Decoder trait. Custom jar     
                                       should be available in kafka/libs        
                                       directory. (default: kafka.serializer.   
                                       StringDecoder)                           
--verify-index-only                  if set, just verify the index log without  
                                       printing its content.                    
--version                            Display Kafka version.


### ./bin/kafka-features.sh
This tool describes and updates finalized features.
Option                                 Description                           
------                                 -----------                           
--bootstrap-server <String: server to  REQUIRED: A comma-separated list of   
  connect to>                            host:port pairs to use for          
                                         establishing the connection to the  
                                         Kafka cluster.                      
--command-config [String: command      Property file containing configs to be
  config property file]                  passed to Admin Client. This is used
                                         with --bootstrap-server option when 
                                         required.                           
--describe                             Describe supported and finalized      
                                         features from a random broker.      
--downgrade-all                        Downgrades all finalized features to  
                                         the maximum version levels known to 
                                         the tool. This command deletes      
                                         unknown features from the list of   
                                         finalized features in the cluster,  
                                         but it is guaranteed to not add a   
                                         new feature.                        
--dry-run                              Performs a dry-run of                 
                                         upgrade/downgrade mutations to      
                                         finalized feature without applying  
                                         them.                               
--help                                 Print usage information.              
--upgrade-all                          Upgrades all finalized features to the
                                         maximum version levels known to the 
                                         tool. This command finalizes new    
                                         features known to the tool that were
                                         never finalized previously in the   
                                         cluster, but it is guaranteed to not
                                         delete any existing feature.        
--version                              Display Kafka version.                



### ./bin/kafka-get-offsets.sh
An interactive shell for getting topic-partition offsets.
Option                                   Description                           
------                                   -----------                           
--bootstrap-server <String: HOST1:       REQUIRED. The server(s) to connect to 
  PORT1,...,HOST3:PORT3>                   in the form HOST1:PORT1,HOST2:PORT2.
--broker-list <String: HOST1:PORT1,...,  DEPRECATED, use --bootstrap-server    
  HOST3:PORT3>                             instead; ignored if --bootstrap-    
                                           server is specified. The server(s)  
                                           to connect to in the form HOST1:    
                                           PORT1,HOST2:PORT2.                  
--command-config <String: config file>   Property file containing configs to be
                                           passed to Consumer Client.          
--exclude-internal-topics                By default, internal topics are       
                                           included. If specified, internal    
                                           topics are excluded.                
--partitions <String: partition ids>     Comma separated list of partition ids 
                                           to get the offsets for. If not      
                                           present, all partitions of the      
                                           authorized topics are queried.      
                                           Cannot be used if --topic-partitions
                                           is present.                         
--time <Long: timestamp/-1(latest)/-2    timestamp of the offsets before that. 
  (earliest)>                              [Note: No offset is returned, if the
                                           timestamp greater than recently     
                                           committed record timestamp is       
                                           given.] (default: -1)               
--topic <String: topic>                  The topic to get the offsets for. It  
                                           also accepts a regular expression.  
                                           If not present, all authorized      
                                           topics are queried. Cannot be used  
                                           if --topic-partitions is present.   
--topic-partitions <String: topic1:1,    Comma separated list of topic-        
  topic2:0-3,topic3,topic4:5-,topic5:-3    partition patterns to get the       
  >                                        offsets for, with the format of     
                                           '([^:,]*)(?::(?:([0-9]*)|(?:([0-9]*)
                                           -([0-9]*))))?'. The first group is  
                                           an optional regex for the topic     
                                           name, if omitted, it matches any    
                                           topic name. The section after ':'   
                                           describes a 'partition' pattern,    
                                           which can be: a number, a range in  
                                           the format of 'NUMBER-NUMBER' (lower
                                           inclusive, upper exclusive), an     
                                           inclusive lower bound in the format 
                                           of 'NUMBER-', an exclusive upper    
                                           bound in the format of '-NUMBER' or 
                                           may be omitted to accept all        
                                           partitions.        


#### ./bin/kafka-leader-election.sh
This tool attempts to elect a new leader for a set of topic partitions. The type of elections supported are preferred replicas and unclean replicas.
Option                                  Description                           
------                                  -----------                           
--admin.config <String: config file>    Configuration properties files to pass
                                          to the admin client                 
--all-topic-partitions                  Perform election on all of the        
                                          eligible topic partitions based on  
                                          the type of election (see the --    
                                          election-type flag). Not allowed if 
                                          --topic or --path-to-json-file is   
                                          specified.                          
--bootstrap-server <String: host:port>  A hostname and port for the broker to 
                                          connect to, in the form host:port.  
                                          Multiple comma separated URLs can be
                                          given. REQUIRED.                    
--election-type <[PREFERRED,UNCLEAN]:   Type of election to attempt. Possible 
  election type>                          values are "preferred" for preferred
                                          leader election or "unclean" for    
                                          unclean leader election. If         
                                          preferred election is selection, the
                                          election is only performed if the   
                                          current leader is not the preferred 
                                          leader for the topic partition. If  
                                          unclean election is selected, the   
                                          election is only performed if there 
                                          are no leader for the topic         
                                          partition. REQUIRED.                
--help                                  Print usage information.              
--partition <Integer: partition id>     Partition id for which to perform an  
                                          election. REQUIRED if --topic is    
                                          specified.                          
--path-to-json-file <String: Path to    The JSON file with the list  of       
  JSON file>                              partition for which leader elections
                                          should be performed. This is an     
                                          example format.                     
                                        {"partitions":                        
                                        	[{"topic": "foo", "partition": 1},   
                                        	 {"topic": "foobar", "partition": 2}]
                                        }                                     
                                        Not allowed if --all-topic-partitions 
                                          or --topic flags are specified.     
--topic <String: topic name>            Name of topic for which to perform an 
                                          election. Not allowed if --path-to- 
                                          json-file or --all-topic-partitions 
                                          is specified.                       
--version                               Display Kafka version.                



### ./bin/kafka-log-dirs.sh
This tool helps to query log directory usage on the specified brokers.
Option                                  Description                           
------                                  -----------                           
--bootstrap-server <String: The server  REQUIRED: the server(s) to use for    
  (s) to use for bootstrapping>           bootstrapping                       
--broker-list <String: Broker list>     The list of brokers to be queried in  
                                          the form "0,1,2". All brokers in the
                                          cluster will be queried if no broker
                                          list is specified                   
--command-config <String: Admin client  Property file containing configs to be
  property file>                          passed to Admin Client.             
--describe                              Describe the specified log directories
                                          on the specified brokers.           
--help                                  Print usage information.              
--topic-list <String: Topic list>       The list of topics to be queried in   
                                          the form "topic1,topic2,topic3". All
                                          topics will be queried if no topic  
                                          list is specified (default: )       
--version                               Display Kafka version.                




### ./bin/kafka-mirror-maker.sh 
This tool helps to continuously copy data between two Kafka clusters.
Option                                   Description                           
------                                   -----------                           
--abort.on.send.failure <String: Stop    Configure the mirror maker to exit on 
  the entire mirror maker when a send      a failed send. (default: true)      
  failure occurs>                                                              
--consumer.config <String: config file>  Embedded consumer config for consuming
                                           from the source cluster.            
--consumer.rebalance.listener <String:   The consumer rebalance listener to use
  A custom rebalance listener of type      for mirror maker consumer.          
  ConsumerRebalanceListener>                                                   
--help                                   Print usage information.              
--message.handler <String: A custom      Message handler which will process    
  message handler of type                  every record in-between consumer and
  MirrorMakerMessageHandler>               producer.                           
--message.handler.args <String:          Arguments used by custom message      
  Arguments passed to message handler      handler for mirror maker.           
  constructor.>                                                                
--new.consumer                           DEPRECATED Use new consumer in mirror 
                                           maker (this is the default so this  
                                           option will be removed in a future  
                                           version).                           
--num.streams <Integer: Number of        Number of consumption streams.        
  threads>                                 (default: 1)                        
--offset.commit.interval.ms <Integer:    Offset commit interval in ms.         
  offset commit interval in                (default: 60000)                    
  millisecond>                                                                 
--producer.config <String: config file>  Embedded producer config.             
--rebalance.listener.args <String:       Arguments used by custom rebalance    
  Arguments passed to custom rebalance     listener for mirror maker consumer. 
  listener constructor as a string.>                                           
--version                                Display Kafka version.                
--whitelist <String: Java regex          Whitelist of topics to mirror.        
  (String)>                                                                    



### ./bin/kafka-preferred-replica-election.sh 
This tool is deprecated. Please use kafka-leader-election tool. Tracking issue: KAFKA-8405
This tool helps to causes leadership for each partition to be transferred back to the 'preferred replica', it can be used to balance leadership among the servers.
Option                                  Description                           
------                                  -----------                           
--admin.config <String: config file>    Admin client config properties file to
                                          pass to the admin client when --    
                                          bootstrap-server is given.          
--bootstrap-server <String: host:port>  A hostname and port for the broker to 
                                          connect to, in the form host:port.  
                                          Multiple comma-separated URLs can be
                                          given. REQUIRED unless --zookeeper  
                                          is given.                           
--help                                  Print usage information.              
--path-to-json-file <String: list of    The JSON file with the list of        
  partitions for which preferred          partitions for which preferred      
  replica leader election needs to be     replica leader election should be   
  triggered>                              done, in the following format -     
                                        {"partitions":                        
                                        	[{"topic": "foo", "partition": 1},   
                                        	 {"topic": "foobar", "partition": 2}]
                                        }                                     
                                        Defaults to all existing partitions   
--version                               Display Kafka version.                
--zookeeper <String: urls>              DEPRECATED. The connection string for 
                                          the zookeeper connection in the form
                                          host:port. Multiple URLS can be     
                                          given to allow fail-over. Replaced  
                                          by --bootstrap-server, REQUIRED     
                                          unless --bootstrap-server is given. 



### ./bin/kafka-producer-perf-test.sh    
usage: producer-performance [-h] --topic TOPIC --num-records NUM-RECORDS [--payload-delimiter PAYLOAD-DELIMITER] --throughput THROUGHPUT
                            [--producer-props PROP-NAME=PROP-VALUE [PROP-NAME=PROP-VALUE ...]] [--producer.config CONFIG-FILE]
                            [--print-metrics] [--transactional-id TRANSACTIONAL-ID] [--transaction-duration-ms TRANSACTION-DURATION]
                            (--record-size RECORD-SIZE | --payload-file PAYLOAD-FILE)

This tool is used to verify the producer performance.

optional arguments:
  -h, --help             show this help message and exit
  --topic TOPIC          produce messages to this topic
  --num-records NUM-RECORDS
                         number of messages to produce
  --payload-delimiter PAYLOAD-DELIMITER
                         provides delimiter to be used when --payload-file is provided.  Defaults to new line. Note that this parameter will
                         be ignored if --payload-file is not provided. (default: \n)
  --throughput THROUGHPUT
                         throttle maximum message  throughput  to  *approximately*  THROUGHPUT  messages/sec.  Set  this  to  -1  to disable
                         throttling.
  --producer-props PROP-NAME=PROP-VALUE [PROP-NAME=PROP-VALUE ...]
                         kafka  producer  related  configuration  properties  like   bootstrap.servers,client.id  etc.  These  configs  take
                         precedence over those passed via --producer.config.
  --producer.config CONFIG-FILE
                         producer config properties file.
  --print-metrics        print out metrics at the end of the test. (default: false)
  --transactional-id TRANSACTIONAL-ID
                         The transactionalId to use if transaction-duration-ms is  >  0.  Useful  when testing the performance of concurrent
                         transactions. (default: performance-producer-default-transactional-id)
  --transaction-duration-ms TRANSACTION-DURATION
                         The max age of each transaction. The  commitTransaction  will  be  called after this time has elapsed. Transactions
                         are only enabled if this value is positive. (default: 0)

  either --record-size or --payload-file must be specified but not both.

  --record-size RECORD-SIZE
                         message size in bytes. Note that you must provide exactly one of --record-size or --payload-file.
  --payload-file PAYLOAD-FILE
                         file to read the message payloads from. This works  only  for  UTF-8 encoded text files. Payloads will be read from
                         this file and a payload will be randomly selected when  sending messages. Note that you must provide exactly one of
                         --record-size or --payload-file.



### ./bin/kafka-reassign-partitions.sh
This tool helps to move topic partitions between replicas.
Option                                  Description                           
------                                  -----------                           
--additional                            Execute this reassignment in addition 
                                          to any other ongoing ones. This     
                                          option can also be used to change   
                                          the throttle of an ongoing          
                                          reassignment.                       
--bootstrap-server <String: Server(s)   the server(s) to use for              
  to use for bootstrapping>               bootstrapping. REQUIRED if an       
                                          absolute path of the log directory  
                                          is specified for any replica in the 
                                          reassignment json file, or if --    
                                          zookeeper is not given.             
--broker-list <String: brokerlist>      The list of brokers to which the      
                                          partitions need to be reassigned in 
                                          the form "0,1,2". This is required  
                                          if --topics-to-move-json-file is    
                                          used to generate reassignment       
                                          configuration                       
--cancel                                Cancel an active reassignment.        
--command-config <String: Admin client  Property file containing configs to be
  property file>                          passed to Admin Client.             
--disable-rack-aware                    Disable rack aware replica assignment 
--execute                               Kick off the reassignment as specified
                                          by the --reassignment-json-file     
                                          option.                             
--generate                              Generate a candidate partition        
                                          reassignment configuration. Note    
                                          that this only generates a candidate
                                          assignment, it does not execute it. 
--help                                  Print usage information.              
--list                                  List all active partition             
                                          reassignments.                      
--preserve-throttles                    Do not modify broker or topic         
                                          throttles.                          
--reassignment-json-file <String:       The JSON file with the partition      
  manual assignment json file path>       reassignment configurationThe format
                                          to use is -                         
                                        {"partitions":                        
                                        	[{"topic": "foo",                    
                                        	  "partition": 1,                    
                                        	  "replicas": [1,2,3],               
                                        	  "log_dirs": ["dir1","dir2","dir3"] 
                                          }],                                 
                                        "version":1                           
                                        }                                     
                                        Note that "log_dirs" is optional. When
                                          it is specified, its length must    
                                          equal the length of the replicas    
                                          list. The value in this list can be 
                                          either "any" or the absolution path 
                                          of the log directory on the broker. 
                                          If absolute log directory path is   
                                          specified, the replica will be moved
                                          to the specified log directory on   
                                          the broker.                         
--replica-alter-log-dirs-throttle       The movement of replicas between log  
  <Long: replicaAlterLogDirsThrottle>     directories on the same broker will 
                                          be throttled to this value          
                                          (bytes/sec). This option can be     
                                          included with --execute when a      
                                          reassignment is started, and it can 
                                          be altered by resubmitting the      
                                          current reassignment along with the 
                                          --additional flag. The throttle rate
                                          should be at least 1 KB/s. (default:
                                          -1)                                 
--throttle <Long: throttle>             The movement of partitions between    
                                          brokers will be throttled to this   
                                          value (bytes/sec). This option can  
                                          be included with --execute when a   
                                          reassignment is started, and it can 
                                          be altered by resubmitting the      
                                          current reassignment along with the 
                                          --additional flag. The throttle rate
                                          should be at least 1 KB/s. (default:
                                          -1)                                 
--timeout <Long: timeout>               The maximum time in ms to wait for log
                                          directory replica assignment to     
                                          begin. (default: 10000)             
--topics-to-move-json-file <String:     Generate a reassignment configuration 
  topics to reassign json file path>      to move the partitions of the       
                                          specified topics to the list of     
                                          brokers specified by the --broker-  
                                          list option. The format to use is - 
                                        {"topics":                            
                                        	[{"topic": "foo"},{"topic": "foo1"}],
                                        "version":1                           
                                        }                                     
--verify                                Verify if the reassignment completed  
                                          as specified by the --reassignment- 
                                          json-file option. If there is a     
                                          throttle engaged for the replicas   
                                          specified, and the rebalance has    
                                          completed, the throttle will be     
                                          removed                             
--version                               Display Kafka version.                
--zookeeper <String: urls>              DEPRECATED: The connection string for 
                                          the zookeeper connection in the form
                                          host:port. Multiple URLS can be     
                                          given to allow fail-over.  Please   
                                          use --bootstrap-server instead.     


### ./bin/kafka-replica-verification.sh 
Validate that all replicas for a set of topics have the same data.
Option                                  Description                           
------                                  -----------                           

--broker-list <String: hostname: port,...,hostname:port>        
  REQUIRED: The list of hostname and port of the server to connect to.   

--fetch-size <Integer: bytes>           
  The fetch size of each request. (default: 1048576)

--help                                  
  Print usage information.              

--max-wait-ms <Integer: ms>             
  The max amount of time each fetch request waits. (default: 1000)

--report-interval-ms <Long: ms>         
  The reporting interval. (default: 30000)

--time <Long: timestamp/-1(latest)/-2 (earliest)>
   Timestamp for getting the initial offsets. (default: -1)              

--topic-white-list <String: Java regex  (default: .*) (String)>                       
  White list of topics to verify replica consistency. Defaults to all topics.

--version                               
  Print version information and exit.   



### ./bin/kafka-run-class.sh
USAGE: ./bin/kafka-run-class.sh [-daemon] [-name servicename] [-loggc] classname [opts]



### ./bin/kafka-server-start.sh 
USAGE: ./bin/kafka-server-start.sh [-daemon] server.properties [--override property=value]*



### ./bin/kafka-storage.sh
usage: kafka-storage [-h] {info,format,random-uuid} ...
The Kafka storage tool.

positional arguments:
  {info,format,random-uuid}
    info                 Get information about the Kafka log directories on this node.
    format               Format the Kafka log directories on this node.
    random-uuid          Print a random UUID.

optional arguments:
  -h, --help             show this help message and exit



### ./bin/kafka-streams-application-reset.sh
Missing required option(s) [application-id]
Option (* = required)                 Description                              
---------------------                 -----------                              
* --application-id <String: id>       
  The Kafka Streams application ID (application.id).

--bootstrap-servers <String: urls>    
  Comma-separated list of broker urls with format: HOST1:PORT1,HOST2:PORT2 (default: localhost:9092)              

--by-duration <String>
Reset offsets to offset by duration from current timestamp. Format: 'PnDTnHnMnS'

--config-file <String: file name>     
Property file containing configs to be passed to admin clients and embedded consumer.                              

--dry-run                             
Display the actions that would be performed without executing the reset commands.                              

--execute
  Execute the command.                     

--force                               
  Force the removal of members of the consumer group (intended to remove stopped members 
  if a long session timeout was used). Make sure to shut down all stream applications when 
  this option is specified to avoid unexpected rebalances.                            

--from-file <String>                  
  Reset offsets to values defined in CSV file.                                  

--help                                
  Print usage information.                 

--input-topics <String: list>         
  Comma-separated list of user input topics. For these topics, the tool will reset
  the offset to the earliest available offset.                      

--intermediate-topics <String: list>  
  Comma-separated list of intermediate user topics (topics that are input and output 
  topics, e.g., used in the deprecated through() method). For these topics, the tool 
  will skip to the end. 

--shift-by <Long: number-of-offsets>  
  Reset offsets shifting current offset by 'n', where 'n' can be positive or negative                               

--to-datetime <String>                
  Reset offsets to offset from datetime. Format: 'YYYY-MM-DDTHH:mm:SS.sss'      

--to-earliest                         
  Reset offsets to earliest offset.        

--to-latest                           
  Reset offsets to latest offset.          

--to-offset <Long>                    
  Reset offsets to a specific offset.      

--version                             
  Print version information and exit.      

--zookeeper                           
  Zookeeper option is deprecated by bootstrap.servers, as the reset tool 
  would no longer access Zookeeper directly.               



### ./bin/kafka-topics.sh
Create, delete, describe, or change a topic.
Option                                   Description                            
------                                   -----------                            
--alter                                  Alter the number of partitions,        
                                           replica assignment, and/or           
                                           configuration for the topic.         
--at-min-isr-partitions                  if set when describing topics, only    
                                           show partitions whose isr count is   
                                           equal to the configured minimum. Not 
                                           supported with the --zookeeper       
                                           option.                              
--bootstrap-server <String: server to    REQUIRED: The Kafka server to connect  
  connect to>                              to. In case of providing this, a     
                                           direct Zookeeper connection won't be 
                                           required.                            
--command-config <String: command        Property file containing configs to be 
  config property file>                    passed to Admin Client. This is used 
                                           only with --bootstrap-server option  
                                           for describing and altering broker   
                                           configs.                             
--config <String: name=value>            
A topic configuration override for the 
                                           topic being created or altered. The  
                                           following is a list of valid         
                                           configurations:                      

                                         	cleanup.policy                        
                                         	compression.type                      
                                         	delete.retention.ms                   
                                         	file.delete.delay.ms                  
                                         	flush.messages                        
                                         	flush.ms                              
                                         	follower.replication.throttled.replicas                             
                                         	index.interval.bytes                  
                                         	leader.replication.throttled.replicas 
                                         	max.compaction.lag.ms                 
                                         	max.message.bytes                     
                                         	message.downconversion.enable         
                                         	message.format.version                
                                         	message.timestamp.difference.max.ms   
                                         	message.timestamp.type                
                                         	min.cleanable.dirty.ratio             
                                         	min.compaction.lag.ms                 
                                         	min.insync.replicas                   
                                         	preallocate                           
                                         	retention.bytes                       
                                         	retention.ms                          
                                         	segment.bytes                         
                                         	segment.index.bytes                   
                                         	segment.jitter.ms                     
                                         	segment.ms                            
                                         	unclean.leader.election.enable  

                                         See the Kafka documentation for full   
                                           details on the topic configs. It is  
                                           supported only in combination with --
                                           create if --bootstrap-server option  
                                           is used (the kafka-configs CLI       
                                           supports altering topic configs with 
                                           a --bootstrap-server option).     

--create                                 Create a new topic.                    
--delete                                 Delete a topic                         
--delete-config <String: name>           A topic configuration override to be   
                                           removed for an existing topic (see   
                                           the list of configurations under the 
                                           --config option). Not supported with 
                                           the --bootstrap-server option.       
--describe                               List details for the given topics.     
--disable-rack-aware                     Disable rack aware replica assignment  
--exclude-internal                       exclude internal topics when running   
                                           list or describe command. The        
                                           internal topics will be listed by    
                                           default         
                                                                
--force                                  Suppress console prompts               
--help                                   Print usage information.               
--if-exists                              if set when altering or deleting or    
                                           describing topics, the action will   
                                           only execute if the topic exists.    
--if-not-exists                          if set when creating topics, the       
                                           action will only execute if the      
                                           topic does not already exist.        
--list                                   List all available topics.             
--partitions <Integer: # of partitions>  The number of partitions for the topic 
                                           being created or altered (WARNING:   
                                           If partitions are increased for a    
                                           topic that has a key, the partition  
                                           logic or ordering of the messages    
                                           will be affected). If not supplied   
                                           for create, defaults to the cluster  
                                           default.                             
--replica-assignment <String:            A list of manual partition-to-broker   
  broker_id_for_part1_replica1 :           assignments for the topic being      
  broker_id_for_part1_replica2 ,           created or altered.                  
  broker_id_for_part2_replica1 :                                                
  broker_id_for_part2_replica2 , ...>                                           
--replication-factor <Integer:           The replication factor for each        
  replication factor>                      partition in the topic being         
                                           created. If not supplied, defaults   
                                           to the cluster default.              
--topic <String: topic>                  The topic to create, alter, describe   
                                           or delete. It also accepts a regular 
                                           expression, except for --create      
                                           option. Put topic name in double     
                                           quotes and use the '\' prefix to     
                                           escape regular expression symbols; e.
                                           g. "test\.topic".                    
--topics-with-overrides                  if set when describing topics, only    
                                           show topics that have overridden     
                                           configs                              
--unavailable-partitions                 if set when describing topics, only    
                                           show partitions whose leader is not  
                                           available                            
--under-min-isr-partitions               if set when describing topics, only    
                                           show partitions whose isr count is   
                                           less than the configured minimum.    
                                           Not supported with the --zookeeper   
                                           option.                              
--under-replicated-partitions            if set when describing topics, only    
                                           show under replicated partitions     
--version                                Display Kafka version.                 
--zookeeper <String: hosts>              DEPRECATED, The connection string for  
                                           the zookeeper connection in the form 
                                           host:port. Multiple hosts can be     
                                           given to allow fail-over.            



### ./bin/kafka-verifiable-consumer.sh
usage: verifiable-consumer [-h] --topic TOPIC --group-id GROUP_ID [--group-instance-id GROUP_INSTANCE_ID] [--max-messages MAX-MESSAGES]
                           [--session-timeout TIMEOUT_MS] [--verbose] [--enable-autocommit] [--reset-policy RESETPOLICY]
                           [--assignment-strategy ASSIGNMENTSTRATEGY] [--consumer.config CONFIG_FILE]
                           (--bootstrap-server HOST1:PORT1[,HOST2:PORT2[...]] | --broker-list HOST1:PORT1[,HOST2:PORT2[...]])

This tool consumes messages from  a  specific  topic  and  emits  consumer  events  (e.g.  group  rebalances, received messages, and offsets
committed) as JSON objects to STDOUT.

optional arguments:
  -h, --help             show this help message and exit
  --topic TOPIC          Consumes messages from this topic.
  --group-id GROUP_ID    The groupId shared among members of the consumer group
  --group-instance-id GROUP_INSTANCE_ID
                         A unique identifier of the consumer instance
  --max-messages MAX-MESSAGES
                         Consume this many messages. If -1 (the default), the  consumer  will consume until the process is killed externally
                         (default: -1)
  --session-timeout TIMEOUT_MS
                         Set the consumer's session timeout (default: 30000)
  --verbose              Enable to log individual consumed records (default: false)
  --enable-autocommit    Enable offset auto-commit on consumer (default: false)
  --reset-policy RESETPOLICY
                         Set reset policy (must be either 'earliest', 'latest', or 'none' (default: earliest)
  --assignment-strategy ASSIGNMENTSTRATEGY
                         Set assignment  strategy  (e.g.  org.apache.kafka.clients.consumer.RoundRobinAssignor)  (default: org.apache.kafka.
                         clients.consumer.RangeAssignor)
  --consumer.config CONFIG_FILE
                         Consumer config properties file (config options shared with command line parameters will be overridden).

Connection Group:
  Group of arguments for connection to brokers

  --bootstrap-server HOST1:PORT1[,HOST2:PORT2[...]]
                         REQUIRED unless --broker-list(deprecated) is specified. The server(s)  to connect to. Comma-separated list of Kafka
                         brokers in the form HOST1:PORT1,HOST2:PORT2,...
  --broker-list HOST1:PORT1[,HOST2:PORT2[...]]
                         DEPRECATED, use --bootstrap-server instead; ignored  if  --bootstrap-server  is specified.  Comma-separated list of
                         Kafka brokers in the form HOST1:PORT1,HOST2:PORT2,...



### ./bin/kafka-verifiable-producer.sh 
usage: verifiable-producer [-h] --topic TOPIC [--max-messages MAX-MESSAGES] [--throughput THROUGHPUT] [--acks ACKS]
                           [--producer.config CONFIG_FILE] [--message-create-time CREATETIME] [--value-prefix VALUE-PREFIX]
                           [--repeating-keys REPEATING-KEYS] (--bootstrap-server HOST1:PORT1[,HOST2:PORT2[...]] |
                           --broker-list HOST1:PORT1[,HOST2:PORT2[...]])

This tool produces increasing integers to the specified topic and  prints  JSON metadata to stdout on each "send" request, making externally
visible which messages have been acked and which have not.

optional arguments:
  -h, --help             show this help message and exit
  --topic TOPIC          Produce messages to this topic.
  --max-messages MAX-MESSAGES
                         Produce this many messages. If -1, produce messages until the process is killed externally. (default: -1)
  --throughput THROUGHPUT
                         If set >= 0, throttle maximum message throughput to *approximately* THROUGHPUT messages/sec. (default: -1)
  --acks ACKS            Acks required on each produced message. See Kafka docs on acks for details. (default: -1)
  --producer.config CONFIG_FILE
                         Producer config properties file.
  --message-create-time CREATETIME
                         Send messages with creation time starting at the arguments value, in milliseconds since epoch (default: -1)
  --value-prefix VALUE-PREFIX
                         If specified, each produced value will have this prefix with a dot separator
  --repeating-keys REPEATING-KEYS
                         If specified, each produced record will  have  a  key  starting  at  0  increment  by  1 up to the number specified
                         (exclusive), then the key is set to 0 again

Connection Group:
  Group of arguments for connection to brokers

  --bootstrap-server HOST1:PORT1[,HOST2:PORT2[...]]
                         REQUIRED: The server(s) to  connect  to.  Comma-separated  list  of  Kafka  brokers  in the form HOST1:PORT1,HOST2:
                         PORT2,...
  --broker-list HOST1:PORT1[,HOST2:PORT2[...]]
                         DEPRECATED, use --bootstrap-server instead; ignored  if  --bootstrap-server  is specified.  Comma-separated list of
                         Kafka brokers in the form HOST1:PORT1,HOST2:PORT2,...




### ./bin/trogdor.sh                  
The Trogdor fault injector.

Usage:
  ./bin/trogdor.sh [action] [options]

Actions:
  agent: Run the trogdor agent.
  coordinator: Run the trogdor coordinator.
  client: Run the client which communicates with the trogdor coordinator.
  agent-client: Run the client which communicates with the trogdor agent.
  help: This help message.





### ./bin/zookeeper-security-migration.sh
ZooKeeper Migration Tool Help. This tool updates the ACLs of znodes as part of the process of setting up ZooKeeper authentication.
Option                                 Description                            
------                                 -----------                            
--enable.path.check                    Checks if all the root paths exist in ZooKeeper before migration. If not, exit the command.

--help                                 Print usage information.               

--version                              Display Kafka version.                 

--zk-tls-config-file <String: ZooKeeper TLS configuration>        
Identifies the file where ZooKeeper client TLS connectivity properties are defined. Any properties other than: 
  zookeeper.ssl.keystore.password, 
  zookeeper.ssl.truststore.location, 
  zookeeper.ssl.truststore.type, 
  zookeeper.ssl.ocsp.enable,     
  zookeeper.ssl.crl.enable, 
  zookeeper.ssl.truststore.password, 
  zookeeper.ssl.protocol, 
  zookeeper.ssl.keystore.type, 
  zookeeper.ssl.client.enable,   
  zookeeper.ssl.cipher.suites,         
  zookeeper.clientCnxnSocket,          
  zookeeper.ssl.endpoint.identification.algorithm, 
  zookeeper.ssl.enabled.protocols, 
  zookeeper.ssl.keystore.location   
  are ignored.

--zookeeper.acl <String>               
  Indicates whether to make the Kafka znodes in ZooKeeper secure or unsecure. The options are 'secure' and 'unsecure'                       

--zookeeper.connect <String>           
Sets the ZooKeeper connect string (ensemble). This parameter takes a comma-separated list of host:port pairs. (default: localhost:2181)     

--zookeeper.connection.timeout         Sets the ZooKeeper connection timeout. 
  <Integer>                              (default: 30000)                     

--zookeeper.session.timeout <Integer>  Sets the ZooKeeper session timeout.    
                                         (default: 30000)                     

### ./bin/zookeeper-shell.sh
USAGE: ./bin/zookeeper-shell.sh zookeeper_host:port[/path] [-zk-tls-config-file file] [args...]


## Kafka Domain Model
   - ACLs 
   - Consumer Group
      - Consumer