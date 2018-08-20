# Terraform Provider for Rancher 2

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">


## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.11.x
-	[Go](https://golang.org/doc/install) 1.10 (to build the provider plugin)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/hashlab/terraform-provider-rancher`

```sh
$ mkdir -p $GOPATH/src/github.com/hashlab
$ cd $GOPATH/src/github.com/hashlab
$ git clone git@github.com:hashlab/terraform-provider-rancher
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/hashlab/terraform-provider-rancher
$ make build
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.10+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
$ $GOPATH/bin/terraform-provider-rancher
```