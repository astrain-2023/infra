# AstrAIn IaC

> "Astreine Infrastruktur" -unknown

## Prerequisites

Requires: Golang, Pulumi

If you use proto to set up golang run `proto setup`.

## Setup

Login with local storage (for now)
```
pulumi login --local
```

Create pulumi stack (once)
```
pulumi stack init
```

Run inside a cmd subfolder to create the infra
```
pulumi up
```

To destroy the infra
```
pulumi down --exclude-protected
```


Create an API key with global access and export the key and secret to your env

```
export CONFLUENT_CLOUD_API_KEY=<your API key>
export CONFLUENT_CLOUD_API_SECRET=<your API secret>
```

## Resources

Terraform examples:
- https://github.com/confluentinc/terraform-provider-confluent/tree/master/examples/configurations
