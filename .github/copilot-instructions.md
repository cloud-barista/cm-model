# Copilot Instructions for cm-model

This repository is dedicated to managing source and target models for cloud migration. It provides standardized Go structs and types for both on-premise and cloud infrastructure models.

## Project Overview

**cm-model** is a common library for cloud migration models that:

- Provides standardized data models for cloud migration scenarios
- Supports both on-premise infrastructure models and cloud infrastructure models
- Is designed to be imported as a dependency in other cloud migration projects
- Avoids circular dependencies by keeping only model definitions

## Repository Structure

```
cm-model/
├── infra/
│   ├── cloud-model/           # Cloud infrastructure models
│   │   ├── copied-tb-model.go # CB-Tumblebug framework models (copied to avoid dependencies)
│   │   ├── model.go           # Recommended cloud infrastructure models
│   │   └── vm-infra-info.go   # VM infrastructure information models
│   └── on-premise-model/      # On-premise infrastructure models
│       ├── model.go           # Main on-premise infrastructure models
│       ├── network.go         # Network-related models
│       └── server.go          # Server hardware and OS models
├── sw/                        # Software models
│   └── softwaremodel.go       # Software installation and migration models
├── data/                      # Data storage (empty, for future use)
├── go.mod                     # Go module definition
└── README.md                  # Project documentation
```

## Key Model Categories

### 1. On-Premise Models (`infra/on-premise-model/`)

**Purpose**: Model on-premise infrastructure for migration planning.

**Key Types**:

- `OnpremiseInfraModel`: Root model for on-premise infrastructure
- `ServerProperty`: Detailed server specifications (CPU, memory, disk, network, OS)
- `NetworkProperty`: Network configuration including IPv4/IPv6 networks
- `CpuProperty`, `MemoryProperty`, `DiskProperty`: Hardware component specifications
- `NetworkInterfaceProperty`: Network interface details
- `RouteProperty`: Routing table information
- `FirewallRuleProperty`: Firewall configuration
- `OsProperty`: Operating system information

### 2. Cloud Models (`infra/cloud-model/`)

**Purpose**: Model recommended cloud infrastructure configurations and manage VM infrastructure.

**Key Types**:

- `RecommendedVmInfraModel`: Recommended VM infrastructure configuration
- `RecommendedVNet`, `RecommendedSecurityGroup`, `RecommendedVmSpec`: Specific cloud resource recommendations
- `VmInfraInfo`: VM infrastructure information
- TB-prefixed types (e.g., `TbMciReq`, `TbVmReq`): CB-Tumblebug framework models (copied to avoid circular dependencies)

### 3. Software Models (`sw/`)

**Purpose**: Model software components, their dependencies, and migration configurations for cloud migration scenarios.

**Key Types**:

**Software Architecture & Types**:

- `SoftwareArchitecture`: Supported architectures (x86_64, ARM variants, etc.)
- `SoftwareType`: Installation methods (package, container, kubernetes, binary)
- `SoftwarePackageType`: Package formats (deb, rpm)
- `SoftwareContainerRuntimeType`: Container runtimes (docker, podman)

**Software Components**:

- `Binary`: Executable software with dependencies and custom configurations
- `Package`: OS package installations with repository and dependency information
- `Container`: Container-based software with runtime, ports, and environment configurations
- `Kubernetes`: Kubernetes-based software deployments with resource definitions

**Source Software Discovery**:

- `SoftwareList`: Collection of all software types found on source systems
- `SourceConnectionInfoSoftwareProperty`: Software inventory per connection
- `SourceGroupSoftwareProperty`: Grouped software inventory across multiple connections

**Migration Planning**:

- `BinaryMigrationInfo`: Binary migration configuration with installation order
- `PackageMigrationInfo`: Package migration with repository and dependency resolution
- `ContainerMigrationInfo`: Container migration with runtime and configuration details
- `KubernetesMigrationInfo`: Kubernetes migration with Velero backup/restore configuration
- `MigrationList`: Complete migration plan for all software types

## Development Guidelines

### 1. Model Design Principles

- **Struct-based**: Use Go structs with JSON tags for serialization
- **Validation**: Include validation tags where appropriate (`validate:"required"`)
- **Documentation**: Include comprehensive field comments and examples
- **Extensibility**: Design for future extension with optional fields and TODO comments

### 2. Package Organization

- Use separate packages for different model categories (`cloudmodel`, `onpremisemodel`, `softwaremodel`)
- Keep models focused and cohesive within each package
- Avoid cross-package dependencies where possible

### 3. JSON Serialization

- Use `json:` tags for all exported fields
- Include `omitempty` for optional fields
- Provide meaningful field names in JSON format

### 4. CB-Tumblebug Integration

- **Version Management**: TB-prefixed models in `copied-tb-model.go` are maintained to match specific CB-Tumblebug versions
- **Current Version**: Check the version comment at the top of `copied-tb-model.go` for the current synchronized CB-Tumblebug version
- **Maintainer-Driven Updates**: Only maintainers can initiate updates to specific CB-Tumblebug versions (e.g., v0.10.3, latest)
- **Update Process**: When updating, compare with the target TB version to identify:
  - Changed struct fields and their types
  - New struct definitions
  - Removed or deprecated models
  - Updated validation tags and examples
- **Synchronization Requirements**:
  - Maintain source path comments with file paths and line numbers from CB-Tumblebug
  - Update version header comment to reflect the target TB version
  - Preserve all existing struct documentation and validation tags
  - Ensure backward compatibility where possible
- **Dependency Avoidance**: Models are copied (not imported) to prevent circular dependencies with CB-Tumblebug

### 5. Field Documentation Standards

- Include units for numeric fields (e.g., "Unit GiB" for memory/disk sizes)
- Provide examples in struct tags
- Reference system commands where relevant (e.g., `df -h`, `ifconfig`, `ip route`)
- Use clear, descriptive field names

## Common Patterns

### Hardware Specifications

```go
type CpuProperty struct {
    Architecture string  `json:"architecture" example:"x86_64"`
    Cpus         uint32  `json:"cpus" validate:"required" example:"2"`
    Cores        uint32  `json:"cores" validate:"required" example:"18"`
    Threads      uint32  `json:"threads" validate:"required" example:"36"`
    MaxSpeed     float32 `json:"maxSpeed,omitempty" example:"3.6"` // GHz
    Vendor       string  `json:"vendor,omitempty" example:"GenuineIntel"`
    Model        string  `json:"model,omitempty"`
}
```

### Network Configuration

```go
type NetworkInterfaceProperty struct {
    Name           string   `json:"name,omitempty" validate:"required"`
    MacAddress     string   `json:"macAddress,omitempty"`
    IPv4CidrBlocks []string `json:"ipv4CidrBlocks,omitempty"`
    IPv6CidrBlocks []string `json:"ipv6CidrBlocks,omitempty"`
    Mtu            int      `json:"mtu,omitempty"`
    State          string   `json:"state,omitempty"`
}
```

### Resource Lists and Collections

```go
type RecommendedVmSpecList struct {
    Status                string              `json:"status"`
    Description           string              `json:"description"`
    Count                 int                 `json:"count"`
    RecommendedVmSpecList []RecommendedVmSpec `json:"recommendedVmSpecList"`
}
```

### Software Component Patterns

```go
type Container struct {
    Name              string                       `json:"name,omitempty" validate:"required"`
    Runtime           SoftwareContainerRuntimeType `json:"runtime,omitempty" validate:"required"`
    ContainerId       string                       `json:"container_id" validate:"required"`
    ContainerImage    ContainerImage               `json:"container_image,omitempty" validate:"required"`
    ContainerPorts    []ContainerPort              `json:"container_ports"`
    ContainerStatus   string                       `json:"container_status" validate:"required"`
    MountPaths        []string                     `json:"mount_paths"`
    Envs              []Env                        `json:"envs"`
}
```

### Migration Configuration Patterns

```go
type PackageMigrationInfo struct {
    Order                    int      `json:"order"`
    Name                     string   `json:"name" validate:"required"`
    Version                  string   `json:"version" validate:"required"`
    NeededPackages           []string `json:"needed_packages" validate:"required"`
    RepoURL                  string   `json:"repo_url"`
    PackageMigrationConfigID string   `json:"package_migration_config_id"`
}
```

## Usage Examples

### Importing Models

```go
import (
    cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
    onpremisemodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
    softwaremodel "github.com/cloud-barista/cm-model/sw"
)
```

### Local Development for Other Subsystems

When developing other cloud migration subsystems that depend on cm-model, you can use a local version of cm-model for testing and development purposes.

To use your local cm-model repository in another subsystem project:

1. **Add replace directive in go.mod** of the subsystem project:

   ```go
   replace github.com/cloud-barista/cm-model => ../cm-model
   ```

2. **Import and use the models** as usual:

   ```go
   import onprem "github.com/cloud-barista/cm-model/infra/on-premise-model"
   import cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
   ```

3. **Test your changes** locally before contributing back to the upstream repository.

**Important Notes:**

- The replace directive should point to the relative or absolute path of your local cm-model directory
- This configuration allows you to develop and test model improvements locally
- Once you've tested your changes, contribute them back to the upstream cm-model repository
- Remove the replace directive when using the official released version

## Contributing

When adding or modifying models:

1. **Follow naming conventions**: Use descriptive struct and field names
2. **Add comprehensive documentation**: Include field comments, examples, and units
3. **Consider validation**: Add validation tags for required fields
4. **Test JSON serialization**: Ensure models serialize/deserialize correctly
5. **Update documentation**: Keep README and comments in sync
6. **Maintain backward compatibility**: Avoid breaking changes to existing models

### CB-Tumblebug Model Updates

**For Maintainers Only**: CB-Tumblebug model synchronization requires maintainer privileges and careful coordination.

**Update Process**:

1. **Version Selection**: Maintainer specifies target CB-Tumblebug version (e.g., `v0.10.3`, `latest`, `v0.12.0`)
2. **Comparison Analysis**: Compare current `copied-tb-model.go` with target TB version:
   - Identify struct field changes (added, removed, type changes)
   - Check for new model definitions
   - Verify validation tag updates
   - Note breaking changes that affect compatibility
3. **Synchronization Steps**:
   - Update version header comment in `copied-tb-model.go`
   - Apply struct changes while preserving existing documentation
   - Update source path comments with correct file paths and line numbers
   - Test model serialization and validation
4. **Validation Requirements**:
   - Ensure no circular dependencies are introduced
   - Verify backward compatibility with existing cm-model usage
   - Update related documentation if needed
   - Create detailed PR with change summary

**Example Version Header**:

```go
// * To avoid circular dependencies, the following structs are copied from the cb-tumblebug framework.
// TODO: When the cb-tumblebug framework is updated, we should synchronize these structs.
// * Version: CB-Tumblebug v0.11.2 (includes Security Group firewall rule model refactor from PR #2063)
```

## GitHub Workflows

The repository includes automated workflows:

- **Prow commands**: Support for `/approve`, `/lgtm`, `/assign` etc.
- **Auto-merge**: Automatic merging based on maintainer approval
- Team-based access control through `cloud-barista` organization

### Using GitHub Copilot in Issues

For complex model design or major refactoring tasks, you can leverage GitHub Copilot:

**When to use Copilot in Issues**:

- Designing new model categories with multiple related structs
- Large-scale TB model synchronization with breaking changes
- Cross-package model integration requirements
- Performance optimization for large model structures

**Example Copilot prompt in Issues**:

```markdown
@github-copilot please help design complete model structures for [specific use case]
following cm-model patterns:

- JSON serialization with proper tags
- Validation tags for required fields
- Comprehensive field documentation
- Backward compatibility considerations
```

## Related Projects

This model library is designed to be used by:

- Cloud migration tools and services
- Infrastructure assessment applications
- Cloud resource planning systems
- Multi-cloud management platforms

The models support migration scenarios from on-premise infrastructure to various cloud providers through the CB-Tumblebug framework integration.
