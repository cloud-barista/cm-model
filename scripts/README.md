# Scripts Directory

This directory contains utility scripts for analyzing and managing the CB-Tumblebug model synchronization in the cm-model project.

## Available Scripts

### 🔍 `analyze_dependencies.py`

Analyzes the struct dependencies across the entire `cloudmodel` package (`copied-tb-model.go`, `model.go`, and `vm-infra-info.go`) and provides comprehensive insights into struct relationships and usage patterns.

#### Features

- **Package-wide Analysis**: Analyzes all Go files in the `infra/cloud-model` directory
- **Cross-file Dependencies**: Identifies dependencies between structs across different files
- **Independence Detection**: Identifies structs with no dependencies on custom types
- **Usage Tracking**: Checks which structs are referenced by other structs within the package
- **File Location Tracking**: Shows which file each struct is defined in
- **Cleanup Candidates**: Identifies completely unused structs across the entire package

#### Usage

```bash
# Basic analysis (all files in cloudmodel package)
python3 scripts/analyze_dependencies.py

# Detailed analysis with dependency chains
python3 scripts/analyze_dependencies.py --verbose

# Show only unused structs (for cleanup)
python3 scripts/analyze_dependencies.py --unused-only

# Show only cross-file dependencies
python3 scripts/analyze_dependencies.py --cross-file-only
```

#### Sample Output

```
🔍 CB-Tumblebug Model Dependency Analysis (CloudModel Package)
=================================================================

📊 Statistics:
   Total files analyzed: 3
   copied-tb-model.go: 25 structs, 3 string types
   model.go: 12 structs, 0 string types
   vm-infra-info.go: 4 structs, 0 string types
   Package total: 41 structs, 3 string types, 44 custom types

✅ INDEPENDENT STRUCTS (no dependencies on custom types) [16]:
   • CreateSubGroupDynamicReq (defined in copied-tb-model.go)
   • CreateSubGroupReq (defined in copied-tb-model.go)
   • FirewallRuleReq (defined in copied-tb-model.go)
   • KeyValue (defined in copied-tb-model.go)
   • Location (defined in copied-tb-model.go)
   • MciCmdReq (defined in copied-tb-model.go)
   • RegionInfo (defined in copied-tb-model.go)
   • SecurityGroupReq (defined in copied-tb-model.go)
   ...

🔗 STRUCTS WITH DEPENDENCIES [25]:
   • ConnConfig (copied-tb-model.go) → RegionZoneInfo, RegionDetail
   • ImageInfo (copied-tb-model.go) → OSArchitecture, OSPlatform, KeyValue, ImageStatus
   • RecommendedVmInfra (model.go) → SecurityGroupReq, SpecInfo, VNetReq, ImageInfo, MciReq, SshKeyReq
   ...

🔄 CROSS-FILE DEPENDENCIES [7]:
   • RecommendedVmInfra (model.go) → MciReq (copied-tb-model.go), SshKeyReq (copied-tb-model.go), ...
   • RecommendedVmOsImage (model.go) → ImageInfo (copied-tb-model.go)
   ...

⚠️  COMPLETELY UNUSED STRUCTS [11]:
   • FirewallRuleReq (defined in copied-tb-model.go)
   • IdList (defined in vm-infra-info.go)
   • MigratedVmInfraModel (defined in vm-infra-info.go)
   ...
```

#### Command Line Options

- `--verbose, -v`: Show detailed dependency information including reference chains and file locations
- `--unused-only, -u`: Show only completely unused structs (useful for cleanup operations)
- `--cross-file-only, -c`: Show only dependencies between structs in different files

#### Use Cases

1. **Before CB-Tumblebug Sync**: Understand current dependency structure across all cloudmodel files
2. **Code Cleanup**: Identify unused structs that might be removed from any file
3. **Refactoring**: Understand impact of struct changes across file boundaries
4. **Architecture Review**: Analyze cross-file dependencies and coupling
5. **Documentation**: Generate comprehensive dependency documentation
6. **Quality Assurance**: Verify struct usage consistency across the package

#### Requirements

- Python 3.6+
- Access to `infra/cloud-model/` directory with Go source files
- Analyzes: `copied-tb-model.go`, `model.go`, `vm-infra-info.go`

#### Technical Details

The script uses regex patterns to:

- Extract struct definitions: `type StructName struct`
- Extract custom types: `type TypeName string`
- Find field references: `FieldName CustomType` or `FieldName []CustomType`
- Analyze dependency chains across multiple Go files in the cloudmodel package
- Track cross-file dependencies to identify architectural coupling

#### Integration with CI/CD

This script can be integrated into automated workflows:

```bash
# Check for unused structs before merging
python3 scripts/analyze_dependencies.py --unused-only

# Check cross-file dependencies for architecture review
python3 scripts/analyze_dependencies.py --cross-file-only

# Generate comprehensive dependency report
python3 scripts/analyze_dependencies.py --verbose > dependency_report.txt
```

## Script Development Guidelines

### Adding New Scripts

1. **Naming Convention**: Use snake_case for script files
2. **Documentation**: Include docstring with usage examples
3. **Error Handling**: Implement proper error handling and user feedback
4. **CLI Interface**: Use `argparse` for command-line options
5. **Project Root Detection**: Auto-detect project root via `go.mod`

### Code Standards

- Follow PEP 8 Python style guidelines
- Include type hints where appropriate
- Add comprehensive error handling
- Provide helpful error messages
- Support both verbose and quiet modes

### Testing Scripts

```bash
# Test basic functionality
cd /path/to/cm-model
python3 scripts/analyze_dependencies.py

# Test with different options
python3 scripts/analyze_dependencies.py --verbose
python3 scripts/analyze_dependencies.py --unused-only

# Test error handling
python3 scripts/analyze_dependencies.py --help
```

## Contributing

When adding new analysis scripts:

1. Follow the existing patterns in `analyze_dependencies.py`
2. Add comprehensive documentation to this README
3. Test scripts from different working directories
4. Ensure scripts work with relative and absolute paths
5. Add appropriate error handling for missing files
6. Update this README with new script information

## Future Enhancements

Potential improvements for the dependency analyzer:

- **Visualization**: Generate dependency graphs (GraphViz, Mermaid)
- **Circular Dependency Detection**: Detect and report circular dependencies
- **Impact Analysis**: Show impact of removing specific structs
- **Version Comparison**: Compare dependencies between different versions
- **Export Formats**: Support JSON, CSV, YAML output formats
- **Web Interface**: Create web-based dependency browser
