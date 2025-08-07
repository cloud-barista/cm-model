---
mode: agent
model: Claude Sonnet 4
tools:
  [
    "read_file",
    "replace_string_in_file",
    "run_in_terminal",
    "get_terminal_output",
    "semantic_search",
    "grep_search",
    "file_search",
    "list_dir",
    "create_directory",
    "get_errors",
    "github_repo",
  ]
description: "Synchronize CB-Tumblebug models in copied-tb-model.go with specified version"
---

# SyncTB - CB-Tumblebug Model Synchronization

Synchronize TB models in `copied-tb-model.go` with the specified CB-Tumblebug version.

## Target Version

${input:version:CB-Tumblebug version (e.g., v0.11.2, v0.12.0, latest)}

## Process Overview

This prompt will help synchronize CB-Tumblebug models by:

1. **Current Version Detection**: Extract current TB version from copied-tb-model.go
2. **Repository Setup**: Clone CB-Tumblebug repository temporarily
3. **Direct Git Diff Execution**: Execute git diff commands directly to compare versions
4. **Direct Model Synchronization**: Apply changes to copied-tb-model.go directly based on diff results
5. **Cleanup**: Remove temporary repository and return to original directory
6. **Validation**: Ensure backward compatibility and proper serialization

## CM-Model Dependencies Priority

**CRITICAL**: Only synchronize structs that are actually used by cm-model. The core dependencies are:

### Primary Dependencies (MUST sync)

- **TbMciReq** - Used by RecommendedVmInfra.TargetVmInfra
- **TbVNetReq** - Used by RecommendedVNet.TargetVNet
- **TbSshKeyReq** - Used by RecommendedSshKey.TargetSshKey
- **TbSpecInfo** - Used by RecommendedVmSpec.TargetVmSpecList ([]TbSpecInfo)
- **TbImageInfo** - Used by RecommendedVmOsImage.TargetVmOsImageList ([]TbImageInfo)
- **TbSecurityGroupReq** - Used by RecommendedSecurityGroup.TargetSecurityGroupList ([]TbSecurityGroupReq)

### Supporting Dependencies (sync if changed)

- **Common types**: KeyValue, Location, RegionInfo, ConnConfig, etc.
- **Enums and constants**: OSArchitecture, OSPlatform, ImageStatus, etc.

### Non-Critical Dependencies (skip unless referenced)

- Internal CB-Tumblebug structs that are not used by cm-model's public API
- Review/validation structs that don't impact cm-model functionality
- Dynamic request extensions that don't affect core model structures

**IMPORTANT**: Before adding new structs, verify they are referenced by cm-model's existing types.

## Tool Usage Guide

### Primary File Operations

- **`read_file`**: Read current TB version from copied-tb-model.go header and examine model structs
- **`replace_string_in_file`**: Apply struct field changes, update version headers, and modify source path comments
- **`get_errors`**: Validate Go compilation after synchronization changes

### Repository and Git Operations

- **`run_in_terminal`**: Execute git commands for cloning, checkout, and diff operations
- **`get_terminal_output`**: Retrieve git diff output and command results for analysis
- **`create_directory`**: Create temporary directories for CB-Tumblebug repository cloning

### Code Analysis and Search

- **`semantic_search`**: Find existing TB model definitions and related code patterns
- **`grep_search`**: Search for specific struct names, validation tags, and field patterns
- **`file_search`**: Locate model files and identify synchronization targets
- **`list_dir`**: Navigate repository structure and verify cleanup operations

### External Repository Access

- **`github_repo`**: Access CB-Tumblebug source code for cross-referencing and validation (if needed)

## Detailed Workflow

### Step 1: Current State Assessment

- **Use `read_file`** to parse current TB version from [copied-tb-model.go](../../infra/cloud-model/copied-tb-model.go) header comment
- **Use `run_in_terminal`** to save current working directory (`pwd`)
- **Use `create_directory`** to create temporary directory for CB-Tumblebug repository

### Step 2: Repository Operations

- **Use `run_in_terminal`**: Clone CB-Tumblebug repository: `git clone https://github.com/cloud-barista/cb-tumblebug.git`
- **Use `run_in_terminal`**: Navigate to cloned repository (`cd cb-tumblebug`)
- **Use `read_file`**: Identify current version (from copied-tb-model.go header)
- **Use `run_in_terminal`**: Checkout target version: `git checkout ${input:version}`

### Step 3: Direct Git Diff Execution

Execute git diff commands directly:

- **Use `run_in_terminal`**: Run: `git diff [current_version]..${input:version} -- src/core/model/` in the CB-Tumblebug repository
- **Use `get_terminal_output`**: Capture and analyze diff output line by line
- **Use `grep_search`**: Parse struct modifications from diff hunks
- Focus on files containing models used in copied-tb-model.go

### Step 4: Direct Model Synchronization

Directly apply identified changes to copied-tb-model.go:

- **Use `replace_string_in_file`** to update struct definitions
- Apply field additions, removals, and type changes from git diff
- Update validation tags and JSON serialization tags
- Preserve cm-model specific documentation enhancements
- Update version header and source path comments
- Update version header with target version
- Maintain source path comments with accurate line numbers
- Preserve cm-model specific documentation enhancements

### Step 5: Cleanup and Validation

- **Use `run_in_terminal`**: Remove cloned CB-Tumblebug repository: `rm -rf cb-tumblebug/`
- **Use `run_in_terminal`**: Return to cm-model directory
- **Use `get_errors`**: Compile and validate synchronized models

## Analysis Steps

### 1. Version Extraction and Preparation

- **Use `read_file`** to extract current TB version from [copied-tb-model.go](../../infra/cloud-model/copied-tb-model.go) header comment
- **Use `create_directory`** to set up temporary workspace for CB-Tumblebug repository cloning
- **Use `semantic_search`** to identify all existing TB models and their current source paths

### 2. CB-Tumblebug Repository Preparation and Git Operations

**Repository Setup:**

- **Use `create_directory`**: Clone CB-Tumblebug repository in a temporary directory: `/tmp/sync-tb-${input:version}/`
- **Use `run_in_terminal`**: Navigate to cloned repository (`cd /tmp/sync-tb-${input:version}/cb-tumblebug`)
- **Use `read_file`**: Identify current version from [copied-tb-model.go](../../infra/cloud-model/copied-tb-model.go) header
- **Use `run_in_terminal`**: Checkout target version: `git checkout ${input:version}`

**Git Diff Execution:**

- Execute git diff command: `git diff [current_version]..${input:version} -- src/core/model/`
- **Use `run_in_terminal`** to execute git diff commands directly in the CB-Tumblebug repository
- **Use `get_terminal_output`** to capture complete diff output for detailed analysis
- Focus on key model files that contain structs used by cm-model's core dependencies:
  - `src/core/model/mci.go` (TbMciReq - PRIMARY DEPENDENCY)
  - `src/core/model/vnet.go` (TbVNetReq - PRIMARY DEPENDENCY)
  - `src/core/model/sshkey.go` (TbSshKeyReq - PRIMARY DEPENDENCY)
  - `src/core/model/spec.go` (TbSpecInfo - PRIMARY DEPENDENCY)
  - `src/core/model/image.go` (TbImageInfo - PRIMARY DEPENDENCY)
  - `src/core/model/securitygroup.go` (TbSecurityGroupReq - PRIMARY DEPENDENCY)
  - `src/core/model/common.go` (Supporting types: KeyValue, ConnConfig, etc.)
  - `src/core/model/config.go` (Supporting types: Location, RegionDetail, etc.)
  - `src/core/model/subnet.go` (TbSubnetReq - used by TbVNetReq)

### 3. Git Diff Execution and Analysis

Execute git diff commands in the CB-Tumblebug repository:

**Primary Git Diff Command:**

```bash
# In the CB-Tumblebug repository - Use run_in_terminal tool
git diff [current_version]..${input:version} -- src/core/model/
```

**File-Specific Diffs for Detailed Analysis:**

```bash
# Execute each command with run_in_terminal, capture output with get_terminal_output
# PRIORITY ORDER: Focus on PRIMARY DEPENDENCIES first
git diff [current_version]..${input:version} -- src/core/model/mci.go         # TbMciReq (PRIMARY)
git diff [current_version]..${input:version} -- src/core/model/vnet.go        # TbVNetReq (PRIMARY)
git diff [current_version]..${input:version} -- src/core/model/sshkey.go      # TbSshKeyReq (PRIMARY)
git diff [current_version]..${input:version} -- src/core/model/spec.go        # TbSpecInfo (PRIMARY)
git diff [current_version]..${input:version} -- src/core/model/image.go       # TbImageInfo (PRIMARY)
git diff [current_version]..${input:version} -- src/core/model/securitygroup.go # TbSecurityGroupReq (PRIMARY)
git diff [current_version]..${input:version} -- src/core/model/subnet.go      # TbSubnetReq (used by TbVNetReq)
git diff [current_version]..${input:version} -- src/core/model/common.go      # Supporting types
git diff [current_version]..${input:version} -- src/core/model/config.go      # Supporting types
```

**Diff Analysis Process:**

- **Use `get_terminal_output`** to capture complete diff output for each file
- **Use `grep_search`** to identify specific struct definitions and field patterns
- Parse diff hunks to identify:
  - Added lines (prefixed with `+`)
  - Removed lines (prefixed with `-`)
  - Context lines for struct identification
- Extract struct field changes:
  - New fields added to existing structs
  - Removed fields from structs
  - Modified field types or tags
  - Updated validation tags (`validate:"..."`)
  - Changed JSON tags (`json:"..."`)
  - Updated examples and comments

### 4. Synchronization Process

#### A. Dependency Analysis and Validation

Before synchronizing any struct, validate its usage in cm-model:

1. **Primary Dependency Check**: Verify the struct is used by cm-model's core types:

   ```bash
   # Use grep_search to find struct references in cm-model
   grep -r "TbMciReq\|TbVNetReq\|TbSshKeyReq\|TbSpecInfo\|TbImageInfo\|TbSecurityGroupReq" /home/ubuntu/dev/cloud-barista/cm-model/infra/cloud-model/
   ```

2. **Reference Validation**: Confirm the struct is referenced in model.go or other core files:

   ```bash
   # Check RecommendedVmInfra and related types
   grep -r "TargetVmInfra\|TargetVNet\|TargetSshKey\|TargetVmSpecList\|TargetVmOsImageList\|TargetSecurityGroupList" /home/ubuntu/dev/cloud-barista/cm-model/
   ```

3. **Skip Non-Referenced Structs**: Do not add structs that are not used by cm-model's public API

#### B. Version Header Update

Update the header comment in copied-tb-model.go:

```go
// * To avoid circular dependencies, the following structs are copied from the cb-tumblebug framework.
// TODO: When the cb-tumblebug framework is updated, we should synchronize these structs.
// * Version: CB-Tumblebug ${input:version} (include notable changes or PR references)
```

#### C. Source Path Comments Maintenance

Update path comments for each affected struct:

```go
// * Path: src/core/model/[filename], Line: [start]-[end]
```

#### D. Field Synchronization

For each struct identified in git diff output **AND validated as cm-model dependency**:

1. **Field Additions**: Add new fields exactly as shown in diff `+` lines
2. **Field Removals**: Remove fields shown in diff `-` lines
3. **Field Modifications**: Update field types, tags, and comments based on diff changes
4. **Validation Tag Updates**: Apply validation tag changes (`validate:"required"`, etc.)
5. **JSON Tag Updates**: Update JSON serialization tags (`json:"fieldName"`, `omitempty`)
6. **Example Updates**: Update struct tag examples to match TB source
7. **Comment Preservation**: Maintain cm-model specific documentation while applying TB changes
8. **Tumblebug Comment Protection**: **NEVER DELETE** existing Tumblebug-synchronized field documentation and examples
9. **Path Synchronization**: Update "// \* Path:" line number references to match current CB-Tumblebug source locations

**CRITICAL**: Only apply changes to structs that are part of cm-model's core dependencies.

**CRITICAL SAFEGUARDS**:

- **Comment Preservation**: Preserve ALL existing field comments, examples, and documentation from CB-Tumblebug source
- **Path Accuracy**: Verify and update Path line numbers to match actual CB-Tumblebug source file locations
- **Documentation Enhancement**: Add new documentation from TB source while preserving existing content

#### E. File Operations

Execute file editing operations using VS Code tools:

- **Use `replace_string_in_file`** to apply each struct change systematically
- **Use `read_file`** to verify changes and ensure proper context
- **Use `get_errors`** to validate Go compilation after each major change
- Update multiple structs sequentially based on git diff results
- Maintain proper Go syntax and formatting
- Preserve existing cm-model documentation patterns

### 5. Repository Cleanup

After successful synchronization:

- **Use `run_in_terminal`** to remove the cloned CB-Tumblebug repository
- **Use `list_dir`** to verify cleanup and directory restoration
- **Use `read_file`** to validate final changes in copied-tb-model.go
- Return to original working directory

### 6. Validation Checklist

After synchronization (use appropriate tools for each validation):

- [ ] **`list_dir`**: Temporary CB-Tumblebug repository removed
- [ ] **`run_in_terminal`**: Working directory restored to cm-model
- [ ] **`get_errors`**: No compilation errors detected
- [ ] **`grep_search`**: All structs have proper JSON serialization tags
- [ ] **`semantic_search`**: Validation tags match TB source patterns
- [ ] **`read_file`**: Documentation is preserved and enhanced
- [ ] **Manual Review**: Backward compatibility maintained where possible
- [ ] **`grep_search`**: Source path comments are accurate and reflect target version
- [ ] **`read_file`**: Version header reflects target version with change summary
- [ ] **`semantic_search`**: Verify only cm-model dependencies are synchronized (TbMciReq, TbVNetReq, TbSshKeyReq, TbSpecInfo, TbImageInfo, TbSecurityGroupReq)
- [ ] **`grep_search`**: Confirm no unused structs were added to copied-tb-model.go
- [ ] **CRITICAL**: **`read_file`**: Verify ALL Tumblebug-synchronized field comments and examples are preserved
- [ ] **CRITICAL**: **`grep_search`**: Confirm Path line numbers match actual CB-Tumblebug source file locations
- [ ] **CRITICAL**: **`semantic_search`**: Ensure no valuable documentation was unintentionally deleted during synchronization

## Files to Update

- [copied-tb-model.go](../../infra/cloud-model/copied-tb-model.go)

## Reference Guidelines

Follow the patterns and guidelines defined in:

- [copilot-instructions.md](../../copilot-instructions.md) - CB-Tumblebug Integration section
- [TB Synchronization Guidelines](../../copilot-instructions.md#cb-tumblebug-model-updates)

## Important Notes

- **Maintainer-Only Process**: Only maintainers should initiate TB model synchronization
- **Git-Based Comparison**: Uses git diff for accurate change detection between versions
- **Temporary Repository**: CB-Tumblebug repository is cloned temporarily and cleaned up after use
- **Working Directory Safety**: Process saves and restores original working directory
- **Careful Review Required**: All changes should be reviewed for backward compatibility
- **Testing Essential**: Verify model serialization after updates
- **Documentation Critical**: Maintain comprehensive change documentation
- **Cleanup Mandatory**: Ensure temporary files are removed after synchronization
- **🚨 CRITICAL SAFEGUARD**: **NEVER DELETE Tumblebug-synchronized field comments** - These contain valuable examples and documentation from CB-Tumblebug source that must be preserved
- **🚨 CRITICAL REQUIREMENT**: **Path line numbers must stay synchronized** - Always verify and update "// \* Path:" comments to match actual CB-Tumblebug source file locations

### Dependency Management Guidelines

- **Primary Dependencies Only**: Only synchronize structs that are actively used by cm-model
- **Reference Validation**: Before adding any struct, verify it's referenced in cm-model's codebase
- **Core Dependencies**: Focus on TbMciReq, TbVNetReq, TbSshKeyReq, TbSpecInfo, TbImageInfo, TbSecurityGroupReq
- **Avoid Bloat**: Do not add internal CB-Tumblebug structs that don't impact cm-model functionality
- **Supporting Types**: Include supporting types (KeyValue, enums, etc.) only if they're used by primary dependencies

### Documentation Preservation Guidelines

- **Comment Safeguarding**: **ABSOLUTE REQUIREMENT** - Never delete existing Tumblebug-synchronized field comments, examples, or documentation
- **Path Synchronization**: Always verify and update "// \* Path:" line numbers to match current CB-Tumblebug source file locations
- **Documentation Enhancement**: Add new TB documentation while preserving all existing content
- **Field Examples**: Maintain all `example:` tags and field documentation from both cm-model and TB sources
- **Validation Consistency**: Preserve all validation constraints and their explanatory comments
- **Pre-Edit Verification**: Before modifying any struct, read and document current field comments to ensure preservation

### Tool Usage Best Practices

- **Terminal Operations**: Use `run_in_terminal` for all git commands and `get_terminal_output` for capturing results
- **File Modifications**: Always use `replace_string_in_file` with sufficient context (3-5 lines before/after)
- **Validation**: Run `get_errors` after each significant change to catch compilation issues early
- **Search Operations**: Combine `grep_search` and `semantic_search` for comprehensive code analysis
- **Safety Checks**: Use `list_dir` and `read_file` to verify operations and cleanup

## Expected Output

1. **Current Version Extraction**: Extract current TB version from copied-tb-model.go header
2. **Repository Setup**: Clone CB-Tumblebug repository to temporary directory `/tmp/sync-tb-${input:version}/`
3. **Git Diff Execution**: Run git diff commands and capture complete diff output
4. **Diff Analysis Report**: Detailed analysis of struct changes from git diff results
5. **Model Updates**: Apply changes to `copied-tb-model.go` using file editing tools
6. **Change Summary**: List of all modifications made based on git diff output
7. **Breaking Change Analysis**: Identification of backward compatibility issues
8. **Testing Validation**: Go compilation and model serialization verification
9. **Repository Cleanup**: Removal of temporary CB-Tumblebug repository
10. **Final Validation**: Confirmation of successful synchronization in cm-model directory

## Execution Steps

### Phase 1: Setup and Analysis

1. **`read_file`**: Read current version from copied-tb-model.go header
2. **`create_directory`** + **`run_in_terminal`**: Create temporary directory and clone CB-Tumblebug repository
3. **`run_in_terminal`**: Checkout target version in CB-Tumblebug repository
4. **`run_in_terminal`** + **`get_terminal_output`**: Execute git diff commands directly

### Phase 2: Synchronization

5. **`grep_search`** + **`semantic_search`**: Parse git diff output for struct changes
6. **`read_file`**: **BEFORE EDITING** - Read current struct documentation to preserve existing comments
7. **`replace_string_in_file`**: Apply changes directly to copied-tb-model.go using file editing tools
8. **`replace_string_in_file`**: Update version header and source path comments
9. **CRITICAL**: **`read_file`** + **`run_in_terminal`**: Verify Path line numbers against CB-Tumblebug source files
10. **CRITICAL**: **`read_file`**: Validate that ALL Tumblebug-synchronized field comments are preserved
11. **`get_errors`**: Validate Go syntax and compilation

### Phase 3: Cleanup and Validation

9. **`run_in_terminal`** + **`list_dir`**: Remove temporary CB-Tumblebug repository
10. **`get_errors`** + **`read_file`**: Run final validation in cm-model directory
11. **`semantic_search`**: Generate change summary and compatibility report
