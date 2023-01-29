- can add a tags

### Build [#](https://learn.hashicorp.com/tutorials/terraform/google-cloud-platform-build)

- 1. Init
  - run `terraform init`
  - check there are `.terraform` folder
- 2. format and validate
  - `terraform fmt` : formatting terraform files
  - `terraform validate`: validate the terraform configuration
- 3. Apply terraform config
  - `terraform apply`: apply `main.tf`

### Inspect

When you applied your configuration, Terraform wrote data into a file called `terraform.tfstate`. Terraform stores the IDs and properties of the resources it manages in this file, so that it can update or destroy those resources going forward.

- `terraform show`: inspect terraform state

### Change infrastructure

- modify `main.tf`
- run `terraform apply`

#### destructive changes[#](https://learn.hashicorp.com/tutorials/terraform/google-cloud-platform-change?in=terraform/gcp-get-started#introduce-destructive-changes)

A destructive change is a change that requires the provider to replace the existing resource rather than updating it.

### Destroy infrastructure

- `terraform destroy`: destroy current infrastructure managed by terraform
  - It doesn't affect to existing but not managed by terraform

## Variables

```
Tip: Terraform loads all files ending in `.tf` in the working directory, so you can name your configuration files however you choose. We recommend defining variables in their own file to make your configuration easier to organize and understand.
```

### Assign values to your variables[#](https://learn.hashicorp.com/tutorials/terraform/google-cloud-platform-variables?in=terraform/gcp-get-started#assign-values-to-your-variables)

You can populate variables using values from a file. Terraform automatically loads files called terraform.tfvars or matching \*.auto.tfvars in the working directory when running operations.

## Query data with output variables

- run `terraform output` after apply: Will output the configured output

## Log meaning

- `~`: The prefix `~` means that Terraform will update the resource in-place
- `-/+`: The prefix `-/+` means that Terraform will destroy and recreate the resource, rather than updating it in-place.
- `-`: The `-` prefix indicates that Terraform will destroy the instance and the network.
