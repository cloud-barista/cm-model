# cm-model

This repository is dedicated to managing source and target models for cloud migration. It provides standardized Go structs and types for cloud migration scenarios.

## Project Overview

**cm-model** is a common library for cloud migration models that:

- Provides standardized data models for cloud migration scenarios
- Supports on-premise infrastructure, cloud infrastructure, and software models
- Is designed to be imported as a dependency in other cloud migration projects
- Avoids circular dependencies by keeping only model definitions

## Repository Structure

```
cm-model/
├── infra/
│   ├── cloud-model/          # Cloud infrastructure models
│   └── on-premise-model/     # On-premise infrastructure models
├── sw/                       # Software models
├── data/                     # Data storage (for future use)
└── go.mod
```

## Model Categories

- **On-Premise Models**: Server specifications, network configuration, OS information
- **Cloud Models**: Recommended cloud infrastructure configurations, CB-Tumblebug integration
- **Software Models**: Software components, dependencies, and migration configurations

## Usage Instructions

### Import and use models

```go
import (
    cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
    onpremisemodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
    softwaremodel "github.com/cloud-barista/cm-model/sw"
)
```

### Local development for other subsystems

To develop and test models locally, add this to your project's go.mod:

```go
replace github.com/cloud-barista/cm-model => ../cm-model
```

Once you've tested your changes, contribute them back to the upstream cm-model repository.
