# Terraform Provider for SUSE NeuVector

The Terraform NeuVector provider is a plugin that allows Terraform to manage the installation and maintenance of SUSE NeuVector on any Kubernetes cluster.

## Terraform useful links

- [Plugin Development](https://developer.hashicorp.com/terraform/plugin)
- [Terraform Plugin SDKv2](https://developer.hashicorp.com/terraform/plugin/sdkv2)
- [Writing Custom Terraform Providers](https://www.hashicorp.com/blog/writing-custom-terraform-providers)

## SUSE NeuVector useful links

- [Welcome to the NeuVector Docs](https://open-docs.neuvector.com/)
- [NeuVector Helm Chart](https://github.com/neuvector/neuvector-helm/tree/master/charts/core)

## Requirements

- [Terraform](./terraform.md)
- [Go](./go.md)

**To find out which version of the tools above was used to build the provider in the repository, look at this [file](./versions.md).**

## Develop the provider (macOS example)

#### [Fork the repository](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/fork-a-repo)

#### Download the repository locally

```console
git clone git@github.com:<YOUR-GITHUB-USERNAME>/terraform-provider-neuvector.git
cd terraform-provider-neuvector
```

#### Make your changes

#### Edit the *main.go* file to point to your repository

```console
sed -i '' "s@github.com/glovecchi0/terraform-provider-neuvector/neuvector@github.com/<YOUR-GITHUB-USERNAME>/terraform-provider-neuvector/neuvector@g" main.go
```

#### [Generate a Personal GitHub Access Token](https://github.com/settings/tokens) 

#### Configure the correct environment variables

```console
export GOPRIVATE=github.com/<YOUR-GITHUB-USERNAME>
export GIT_TERMINAL_PROMPT=1
export GITHUB_TOKEN=<YOUR-PERSONAL-ACCESS-TOKEN>
```

#### Build the provider

```console
go mod init github.com/<YOUR-GITHUB-USERNAME>/terraform-provider-neuvector
go mod tidy
go build -o terraform-provider-neuvector_v<VERSION>
```

## Test the provider (locally)

#### Make sure you are inside the repository

```console
$ pwd
  ~/terraform-provider-neuvector
```

#### Copy the provider to the local path *~/.terraform.d/plugins/*

```console
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/local/neuvector/<VERSION>/darwin_arm64
cp terraform-provider-neuvector_v<VERSION> ~/.terraform.d/plugins/registry.terraform.io/local/neuvector/<VERSION>/darwin_arm64/
chmod +x ~/.terraform.d/plugins/registry.terraform.io/local/neuvector/<VERSION>/darwin_arm64/terraform-provider-neuvector_v<VERSION>
```
 
#### Run the sample Terraform files

```console
cd tf-projects/example/

```
