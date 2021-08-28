terraform {
  required_providers {
    tmc = {
      source = "vmware/tanzu/tmc"
    }
  }
}

// Read TMC cluster : fetch cluster details
data "tmc_cluster" "ready_only_attach_cluster_view" {
  management_cluster_name = "attached"       // Default: attached
  provisioner_name        = "attached"       // Default: attached
  name                    = "terraform-test" // Required
}

// Create TMC attach cluster entry
resource "tmc_cluster" "attach_cluster_without_apply" {
  management_cluster_name = "attached"         // Default: attached
  provisioner_name        = "attached"         // Default: attached
  name                    = "terraform-attach" // Required

  meta {
    description = "create attach cluster from terraform"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" // Default: default
  }

  wait_until_ready = false
}

output "attach_output" {
  value = tmc_cluster.attach_cluster_without_apply
}

