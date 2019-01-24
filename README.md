# terraform-provider-pivotaltracker
a terraform provider to manage pivotal tracker projects and membership

[![CircleCI](https://circleci.com/gh/xchapter7x/terraform-provider-pivotaltracker/tree/master.svg?style=svg)](https://circleci.com/gh/xchapter7x/terraform-provider-pivotaltracker/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/xchapter7x/terraform-provider-pivotaltracker)](https://goreportcard.com/report/github.com/xchapter7x/terraform-provider-pivotaltracker)
[![GoDoc](https://godoc.org/github.com/xchapter7x/terraform-provider-pivotaltracker?status.svg)](https://godoc.org/github.com/xchapter7x/terraform-provider-pivotaltracker)


### Installing the Plugin

```bash
# create the providers dir if needed
$ mkdir -p ~/terraform-providers

# set your version (most recent is here: https://github.com/xchapter7x/terraform-provider-pivotaltracker/releases/latest )
$ export VERSION=v0.0.2
# set your platform (unix|osx|win) supported
$ export PLATFORM=osx

# dload the plugin
$ curl -L https://github.com/xchapter7x/terraform-provider-pivotaltracker/releases/download/${VERSION}/pivotal_tracker_provider_${PLATFORM} -o ~/terraform-providers/pivotal_tracker_provider

# make it executable just in case
$ chmod +x -o ~/terraform-providers/pivotal_tracker_provider

# add the provider to your terraformrc (create one like below if one doesnt exist)
$ cat << EOF > ~/.terraformrc
providers {
    pivotaltracker = "$HOME/terraform-providers/pivotal_tracker_provider"
}
EOF

# create a sample project to manage with terraform
$ cat << EOF > sampleProject.tf 
resource "pivotaltracker_project" "test_project" {
   name  = "some_new_project"
   description = "change description again"
}
EOF

# initialize plugins
$ terraform init

# check your terraform
$ terraform plan

# and have terraform do its thing
$ terraform apply
```


### Available Resources

- Project Resource
  - fields:
    - no_owner
    - new_account_name
    - name
    - status
    - iteration_length
    - week_start_day
    - point_scale
    - bugs_and_chores_are_estimatable
    - automatic_planning
    - enable_tasks
    - start_date
    - time_zone
    - velocity_averaged_over
    - number_of_done_iterations_to_show
    - description
    - profile_content
    - enable_incoming_emails
    - initial_velocity
    - project_type
    - public
    - atom_enabled
    - account_id
    - join_as
