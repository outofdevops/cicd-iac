module "wrapper" {
  source     = "../terraform-google-cloud-storage"
  project_id = var.input.project_id
  prefix     = var.input.prefix

  names              = var.input.names
  bucket_policy_only = var.input.bucket_policy_only
  folders            = var.input.folders

  lifecycle_rules = var.input.lifecycle_rules
}