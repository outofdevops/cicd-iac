locals {
  input_file         = "./input.yaml"
  input_file_content = fileexists(local.input_file) ? file(local.input_file) : "NoInputFileFound: true"
  input  = yamldecode(local.input_file_content)
}

module "cloud_storage" {
  source     = "./modules/wrapper"
  input = local.input
}