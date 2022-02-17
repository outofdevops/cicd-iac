variable "input" {
    type = object({
        project_id = string
        prefix     = string

        names              = list(string)
        bucket_policy_only = map(string)
        folders            = map(list(string))
        force_destroy      = bool
        lifecycle_rules    = list(object(
            {
                action = object({
                    type = string
                    storage_class = string
                })
                condition = object({
                    age = string
                    matches_storage_class = string
                })
            })    
        )
    })
}
