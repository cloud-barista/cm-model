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
3. **Git Diff Analysis**: Execute git diff to identify all changed structs between versions
4. **Struct Dependency Mapping**: Map all existing structs in copied-tb-model.go and their dependencies
5. **Comprehensive Synchronization**: Update ALL structs that exist in copied-tb-model.go and their dependencies
6. **Cleanup**: Remove temporary repository and return to original directory
7. **Validation**: Ensure compilation and proper serialization

## Synchronization Principles

**CRITICAL GUIDELINES**:

### 1. Dependency-Based Synchronization Rule

- **ALWAYS** synchronize ALL structs currently present in copied-tb-model.go
- **ONLY** add new structs that are **direct or indirect dependencies** of existing structs
- **NEVER** add standalone new structs that have no dependency chain to existing structs
- **FOLLOW dependency chains**: If existing struct A uses new struct B, and B uses new struct C, include both B and C

### 2. Struct Dependency Chain Analysis

- Map ALL existing structs in copied-tb-model.go before synchronization
- For each existing struct, identify ALL field types that reference other structs
- Trace dependency chains: `ExistingStruct → NewDependency → SubDependency → ...`
- **REJECT** new structs that cannot be traced back to any existing struct through dependency chains

### 3. Operations Scope

- **UPDATE**: Modify existing structs to match target version (always required)
- **CREATE**: Add new structs ONLY if they are dependencies of existing/updated structs
- **DELETE**: Remove structs that no longer exist in target version (with impact analysis)

### 4. Dependency Chain Filtering

- **INCLUDE**: New structs referenced in fields of existing structs
- **INCLUDE**: New structs referenced in fields of already-included dependency structs
- **EXCLUDE**: New structs that exist in CB-Tumblebug but have no dependency path to existing cm-model structs
- **EXCLUDE**: Standalone new functionality that doesn't integrate with existing structs

## Tool Usage Guide

### Primary File Operations

- **`read_file`**: Read current TB version from copied-tb-model.go header and examine existing structs
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
- Update version header with target version
- Maintain source path comments with accurate line numbers
- Preserve cm-model specific documentation enhancements

### Step 5: Cleanup and Validation

- **Use `run_in_terminal`**: Remove cloned CB-Tumblebug repository: `rm -rf cb-tumblebug/`
- **Use `run_in_terminal`**: Return to cm-model directory
- **Use `get_errors`**: Compile and validate synchronized models

## Analysis Steps

### 1. Current State Analysis

- **Use `read_file`** to extract current TB version from [copied-tb-model.go](../../infra/cloud-model/copied-tb-model.go) header comment
- **Use `grep_search`** to inventory ALL existing struct definitions in copied-tb-model.go
- **Use `semantic_search`** to map struct dependencies and relationships within copied-tb-model.go
- **Use `create_directory`** to set up temporary workspace for CB-Tumblebug repository cloning

### 2. Repository Setup and Git Diff Analysis

**Repository Setup:**

- **Use `create_directory`**: Clone CB-Tumblebug repository in a temporary directory: `/tmp/sync-tb-${input:version}/`
- **Use `run_in_terminal`**: Navigate to cloned repository (`cd /tmp/sync-tb-${input:version}/cb-tumblebug`)
- **Use `run_in_terminal`**: Checkout target version: `git checkout ${input:version}`

**Comprehensive Git Diff Execution:**

- **Use `run_in_terminal`**: Execute comprehensive diff: `git diff [current_version]..${input:version} -- src/core/model/`
- **Use `get_terminal_output`**: Capture complete diff output for analysis
- **Focus**: ALL model files without exclusion:
  - `src/core/model/mci.go` (MCI-related structs)
  - `src/core/model/vnet.go` (VNet-related structs)
  - `src/core/model/sshkey.go` (SSH key structs)
  - `src/core/model/spec.go` (Specification structs)
  - `src/core/model/image.go` (Image-related structs)
  - `src/core/model/securitygroup.go` (Security group structs)
  - `src/core/model/subnet.go` (Subnet structs)
  - `src/core/model/common.go` (Common types and constants)
  - `src/core/model/config.go` (Configuration structs)
  - All other model files that contain struct changes

### 3. Dependency Chain Impact Analysis

**Phase 1: Existing Struct Inventory**

- **Use `grep_search`**: List ALL struct names currently in copied-tb-model.go: `type.*struct`
- Create inventory of current structs: TbMciReq, TbVNetReq, TbMciInfo, etc.
- **Use `semantic_search`**: Map field types and dependencies within each existing struct

**Phase 2: Dependency Chain Mapping**

For each existing struct, identify all struct-type fields:

```bash
# Example dependency mapping for existing structs:
# TbMciInfo → MciSshCmdResult (existing) + CreationErrors (potential new dependency)
# TbMciReq → TbVmReq (existing) + MciCmdReq (existing) + PolicyTypes (potential constants)
# TbVmInfo → Location (existing) + RegionInfo (existing) + ConnConfig (existing) + KeyValue (existing)
```

**Phase 3: Git Diff Analysis with Dependency Focus**

- **Use `run_in_terminal`** + **`get_terminal_output`**: Execute git diff for each model file
- **Priority**: Focus on changes to existing structs first
- **Dependency Tracing**: For each field change in existing structs, identify if it introduces new struct dependencies
- **Chain Following**: If a new dependency is found, recursively analyze its dependencies

**Phase 4: Dependency Chain Validation**

```bash
# For each new struct found in git diff, validate dependency chain:
# 1. Is this struct referenced by any existing struct? → INCLUDE
# 2. Is this struct referenced by any already-included dependency struct? → INCLUDE
# 3. Does this struct have no dependency path to existing structs? → EXCLUDE
```

**Diff Analysis Process:**

- **Use `get_terminal_output`** to capture complete diff output for each file
- **Use `grep_search`** to identify specific struct definitions and field patterns
- Parse diff hunks to identify:

  - Added lines (prefixed with `+`)
  - Removed lines (prefixed with `-`)
  - Context lines for struct identification
  - Context lines for struct identification
    **Git Diff Parsing:**

- **Use `grep_search`** to parse git diff output for:
  - Struct definitions: `type.*struct`
  - Added lines: lines starting with `+`
  - Removed lines: lines starting with `-`
  - Modified struct fields and their types

**Change Classification:**

- Identify changes to structs that exist in copied-tb-model.go
- Identify new struct dependencies introduced by existing struct changes
- Identify removed struct dependencies no longer needed

### 4. Dependency-Based Synchronization Process

#### A. Mandatory Synchronization Rules

**Rule 1: Update ALL Existing Structs**

- **MUST** update every struct that exists in copied-tb-model.go if changed in git diff
- **MUST** include all field additions, modifications, and deletions
- **NO** exceptions for "complexity" or subjective necessity judgments

**Rule 2: Include ONLY Dependency Chain Structs**

- **MUST** add new struct types referenced by existing structs (direct dependencies)
- **MUST** add new struct types referenced by direct dependencies (indirect dependencies)
- **MUST** include nested types, array element types, pointer target types only if they connect to existing structs
- **EXCLUDE** new structs that have no dependency path to any existing struct

**Rule 3: Dependency Chain Operations**

- **CREATE**: Add new dependency structs ONLY if required by existing/updated structs
- **UPDATE**: Modify ALL existing structs according to git diff (mandatory)
- **DELETE**: Remove structs no longer referenced (with careful analysis)

**Rule 4: Dependency Chain Validation Process**

For each new struct found in CB-Tumblebug git diff:

1. **Trace Back**: Can this struct be reached from any existing cm-model struct through field references?
2. **Dependency Path**: Is there a chain: `ExistingStruct → ... → NewStruct`?
3. **Decision**: Include ONLY if dependency path exists, otherwise EXCLUDE

**Example Dependency Chain Analysis:**

```go
// ✅ INCLUDE: TbMciInfo (existing) → MciCreationErrors (new) → VmCreationError (new)
// ✅ INCLUDE: Dependency chain exists from existing struct

// ❌ EXCLUDE: ReviewMciDynamicReqInfo (standalone new struct)
// ❌ EXCLUDE: No existing struct references this new struct

// ✅ INCLUDE: SpiderVpcInfo (new) ← IF referenced by existing struct
// ❌ EXCLUDE: SpiderVpcInfo (new) ← IF NOT referenced by existing struct
```

#### B. Version Header Update

Update the header comment in copied-tb-model.go:

```go
// * To avoid circular dependencies, the following structs are copied from the cb-tumblebug framework.
// TODO: When the cb-tumblebug framework is updated, we should synchronize these structs.
// * Version: CB-Tumblebug ${input:version} (include notable changes or PR references)
```

#### C. Source Path Comments Maintenance

Update path comments for each affected struct using consistent line number rules:

```go
// * Path: src/core/model/[filename], Line: [comment_start_or_struct_start]-[end]
```

**Line Number Rule**: Use comment start line if struct has a comment (e.g., "// TbMciReq is struct..."), otherwise use struct definition start line (e.g., "type TbVmInfo struct").

#### D. Complete Field Synchronization

For EVERY struct that exists in copied-tb-model.go AND appears in git diff:

1. **Field Additions**: Add ALL new fields exactly as shown in git diff `+` lines
2. **Field Removals**: Remove ALL fields shown in git diff `-` lines
3. **Field Modifications**: Update ALL field types, tags, and comments based on diff changes
4. **Validation Tag Updates**: Apply ALL validation tag changes (`validate:"required"`, etc.)
5. **JSON Tag Updates**: Update ALL JSON serialization tags (`json:"fieldName"`, `omitempty`)
6. **Example Updates**: Update ALL struct tag examples to match TB source
7. **Comment Preservation**: Maintain ALL existing Tumblebug field documentation and examples
8. **Path Synchronization**: Update ALL "// \* Path:" line number references to match CB-Tumblebug source

#### E. Dependency Struct Addition

For NEW structs referenced by existing structs that appear in git diff:

1. **Complete Addition**: Add the ENTIRE new struct definition from CB-Tumblebug source
2. **All Fields**: Include ALL fields with complete documentation
3. **Proper Placement**: Add in logical order near related structs
4. **Full Documentation**: Include ALL comments, examples, and validation tags from TB source

#### F. File Operations

Execute file editing operations using VS Code tools:

- **Use `replace_string_in_file`** to apply EVERY struct change from git diff
- **Use `read_file`** to verify changes and ensure proper context
- **Use `get_errors`** to validate Go compilation after each change
- **Use `grep_search`** to verify all structs are properly synchronized
- Maintain proper Go syntax and formatting
- Preserve existing cm-model documentation patterns

### 5. Repository Cleanup

After successful synchronization:

- **Use `run_in_terminal`** to remove the cloned CB-Tumblebug repository
- **Use `list_dir`** to verify cleanup and directory restoration
- **Use `read_file`** to validate final changes in copied-tb-model.go
- Return to original working directory

### 6. Path Line Number Synchronization (CRITICAL)

**MANDATORY Line Number Verification Process:**

- [ ] **`run_in_terminal`**: For each synchronized struct, execute `grep -n "// StructName is struct" /tmp/sync-tb-${input:version}/cb-tumblebug/src/core/model/*.go` to find actual comment start line numbers
- [ ] **`run_in_terminal`**: For structs WITHOUT comment lines, use `grep -n "type StructName struct" /tmp/sync-tb-${input:version}/cb-tumblebug/src/core/model/[filename]` to find struct definition start line
- [ ] **`run_in_terminal`**: For each struct, execute `sed -n 'start_line,estimated_end' /tmp/sync-tb-${input:version}/cb-tumblebug/src/core/model/[filename] | grep -n "^}"` to find exact end line
- [ ] **`replace_string_in_file`**: Update ALL "// \* Path:" comments to match exact CB-Tumblebug source line numbers
- [ ] **CRITICAL**: **Line Number Consistency Rule**: ALWAYS use comment start line (if exists) or struct definition start line (if no comment) as the starting line number
- [ ] **CRITICAL**: **`read_file`**: Verify ALL Tumblebug-synchronized field comments and examples are preserved
- [ ] **CRITICAL**: **`grep_search`**: Confirm Path line numbers match actual CB-Tumblebug source file locations
- [ ] **CRITICAL**: **`semantic_search`**: Ensure no valuable documentation was unintentionally deleted during synchronization

**Line Number Update Commands (REQUIRED):**

```bash
# For structs WITH comment lines, find comment start:
grep -n "// StructName is struct" /tmp/sync-tb-${input:version}/cb-tumblebug/src/core/model/[filename]

# For structs WITHOUT comment lines, find struct definition start:
grep -n "type StructName struct" /tmp/sync-tb-${input:version}/cb-tumblebug/src/core/model/[filename]

# Find end line (for both cases):
sed -n 'start,estimated_end' /tmp/sync-tb-${input:version}/cb-tumblebug/src/core/model/[filename] | grep -n "^}"
# Calculate exact end line: start_line + matched_line_number - 1
```

**Path Comment Format (EXACT):**

```go
// * Path: src/core/model/[filename], Line: [comment_start_or_struct_start]-[actual_end]
```

**Line Number Consistency Examples:**

```go
// ✅ CORRECT (Comment exists - use comment start line):
// * Path: src/core/model/mci.go, Line: 261-314
// TbMciDynamicReq is struct for requirements to create MCI dynamically (with default resource option)
type TbMciDynamicReq struct {

// ✅ CORRECT (No comment - use struct definition start line):
// * Path: src/core/model/mci.go, Line: 564-634
type TbVmInfo struct {
```

### 7. Final Validation Checklist

After synchronization (use appropriate tools for each validation):

- [ ] **`list_dir`**: Temporary CB-Tumblebug repository removed
- [ ] **`run_in_terminal`**: Working directory restored to cm-model
- [ ] **`get_errors`**: No compilation errors detected
- [ ] **`grep_search`**: All existing structs synchronized with git diff changes
- [ ] **`semantic_search`**: All new dependency structs added ONLY if connected to existing structs
- [ ] **Dependency Chain Verification**: No standalone new structs included without dependency path
- [ ] **`read_file`**: Documentation is preserved and enhanced
- [ ] **Manual Review**: Backward compatibility maintained where possible
- [ ] **`grep_search`**: Source path comments are accurate and reflect target version
- [ ] **`read_file`**: Version header reflects target version with change summary
- [ ] **CRITICAL**: **`semantic_search`**: Verify NO orphaned structs exist (all new structs must trace back to existing structs)
- [ ] **CRITICAL**: **Dependency Path Validation**: Each new struct has clear dependency chain to existing cm-model structs
- [ ] **`grep_search`**: Confirm ALL dependency structs are present
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
- **Complete Synchronization**: ALL existing structs MUST be synchronized according to git diff
- **No Arbitrary Filtering**: NEVER skip structs based on subjective complexity judgments
- **Dependency Inclusion**: MUST include ALL dependency structs required by existing structs
- **Documentation Critical**: Maintain comprehensive change documentation
- **🚨 CRITICAL SAFEGUARD**: **NEVER DELETE Tumblebug-synchronized field comments** - These contain valuable examples and documentation from CB-Tumblebug source that must be preserved
- **🚨 CRITICAL REQUIREMENT**: **Path line numbers must stay synchronized** - Always verify and update "// \* Path:" comments to match actual CB-Tumblebug source file locations

### Complete Synchronization Guidelines

- **Existing Struct Rule**: ALL structs in copied-tb-model.go MUST be updated according to git diff
- **Dependency Chain Rule**: Add new structs ONLY if they have dependency chains to existing structs
- **No Orphaned Structs**: Do NOT include standalone new structs without dependency paths to existing structs
- **Full Operations**: Perform CREATE (dependencies only), UPDATE (existing structs), DELETE operations as needed
- **Complete Documentation**: Include ALL field comments, examples, and validation tags from TB source

### Dependency Chain Guidelines

- **Trace Dependencies**: For each new struct in git diff, verify if it's referenced by existing or dependency-connected structs
- **Follow Chains**: Include multi-level dependencies: `ExistingStruct → NewStruct1 → NewStruct2 → ...`
- **Exclude Orphans**: Reject new structs that cannot be traced back to existing structs
- **Examples of Valid Chains**:
  - `TbMciInfo (existing) → MciCreationErrors (new) → VmCreationError (new)` ✅
  - `TbSpecInfo (existing) → NewSpecExtension (new)` ✅
- **Examples of Invalid Chains**:
  - `ReviewMciDynamicReqInfo (standalone new)` ❌
  - `ProvisioningLog (standalone new)` ❌

### Tool Usage Best Practices

- **Terminal Operations**: Use `run_in_terminal` for all git commands and `get_terminal_output` for capturing results
- **File Modifications**: Always use `replace_string_in_file` with sufficient context (3-5 lines before/after)
- **Validation**: Run `get_errors` after each significant change to catch compilation issues early
- **Search Operations**: Combine `grep_search` and `semantic_search` for comprehensive code analysis
- **Safety Checks**: Use `list_dir` and `read_file` to verify operations and cleanup

## Expected Output

1. **Current State Analysis**: Complete inventory of existing structs in copied-tb-model.go
2. **Repository Setup**: Clone CB-Tumblebug repository to temporary directory `/tmp/sync-tb-${input:version}/`
3. **Dependency Chain Analysis**: Identify git diff changes and trace dependency chains from existing structs
4. **Change Classification**: Categorize changes into existing struct updates vs. dependency-connected new structs
5. **Selective Synchronization**: Apply ALL changes to existing structs and add ONLY dependency-connected new structs
6. **Full Documentation Update**: Update ALL path references and version headers
7. **Compilation Verification**: Ensure ALL changes compile without errors
8. **Repository Cleanup**: Removal of temporary CB-Tumblebug repository
9. **Final Validation**: Confirmation that ALL existing structs are synchronized and dependencies complete

## Execution Steps

### Phase 1: Dependency Chain Analysis

1. **`read_file`**: Read current version from copied-tb-model.go header
2. **`grep_search`**: Inventory ALL existing struct definitions in copied-tb-model.go
3. **`create_directory`** + **`run_in_terminal`**: Create temporary directory and clone CB-Tumblebug repository
4. **`run_in_terminal`**: Checkout target version in CB-Tumblebug repository
5. **`run_in_terminal`** + **`get_terminal_output`**: Execute comprehensive git diff commands

### Phase 2: Selective Synchronization

6. **`semantic_search`**: Analyze git diff output to identify ALL changes to existing structs
7. **Dependency Chain Tracing**: For each new struct in git diff, verify dependency path to existing structs
8. **`read_file`**: **BEFORE EDITING** - Read current struct documentation to preserve existing comments
9. **`replace_string_in_file`**: Apply ALL struct changes to existing structs according to git diff
10. **`replace_string_in_file`**: Add ONLY dependency-connected new structs with complete definitions
11. **`replace_string_in_file`**: Update version header and ALL source path comments
12. **`get_errors`**: Validate Go syntax and compilation after each major change

### Phase 3: Dependency Chain Validation

13. **`run_in_terminal`** + **`list_dir`**: Remove temporary CB-Tumblebug repository
14. **`get_errors`** + **`read_file`**: Run final validation in cm-model directory
15. **`semantic_search`**: Verify NO orphaned structs exist - all new structs must connect to existing structs
16. **Dependency Path Review**: Generate summary showing dependency chains for all included new structs
