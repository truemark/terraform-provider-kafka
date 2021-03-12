terraform {
  required_providers {
    truemark-kafka = {
      source = "truemark.io/terraform/truemark-kafka"
      version = "1.0.0"
    }
  }
}

################################################################
################################################################
#
#     --authorizer-properties zookeeper.connect=zookeeper:2181 \
#     --add \
#     --cluster \
#     --operation Alter \
#     --deny-principal User:ANONYMOUS
#
# Adding ACLs for resource 
#     `Cluster:LITERAL:kafka-cluster: User:ANONYMOUS has Deny permission 
#      for operations: Alter from hosts: *
#
# Current ACLs for resource `Cluster:LITERAL:kafka-cluster: 
#     User:ANONYMOUS has Deny permission for operations: Alter from hosts: *
#


provider truemark-kafka {
  bootstrap_servers = ["localhost:9092"]
}

resource "truemark-kafka_acl" "global" {
  principal       = "User:*"
  host            = "*"
  operation       = "all"
  permission      = "allow"
  resource {
    resource_type = "topic"
    resource_name = "*"    
    resource_pattern_type = "any"
  }
}

# resource "truemark-kafka_acl" "acl-config" {
#   principal = "User:ANONYMOUS"
# 	host = "*"
#   operation = "describe"
#   permission = "any"
#   resource {
#     resource_type = "topic"
#     resource_name = "quickstart-events"
#     resource_pattern_type = "literal"
#   }
# }