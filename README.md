# cm-model

This repository is dedicated to managing source and target models for cloud migration. **It has been intended to improve and standardize these models across multiple projects.**

* Purpose: A common repository for improving and managing cloud migration models.
* Status This repository is currently experimental.
* Scope: This repository contains only models.
* Separation: This repository is intentionally separated from Damselfly.
  - Reason for separation: To avoid the impact of unrelated packages when importing model packages.


### Usage Instructions
#### Import and use models

To use the on-premise models from the cm-model repository:

Note - `onprem` alias could be applied according to your preference.
```go
import onprem "github.com/cloud-barista/cm-model/infra/onprem"
```

#### Develop models locally and contribute it to the upstream

To develop and test models locally, point to your local cm-model folder in go.mod:

```go
replace github.com/cloud-barista/cm-model => ../cm-model
```

Once you've tested and improved the models locally, we encourage you to contribute your changes back to the cm-model upstream repository. 
This helps improve the models for everyone and supports the ongoing development of cloud migration solutions.
