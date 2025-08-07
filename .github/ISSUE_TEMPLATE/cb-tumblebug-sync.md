---
name: CB-Tumblebug Model Synchronization
about: Synchronize CB-Tumblebug model definitions with cm-model dependencies
title: 'Sync CB-Tumblebug models to [TARGET_VERSION] for cm-model dependencies'
labels: ['enhancement', 'cb-tumblebug-sync']
assignees: ''
---

## Overview

Synchronize CB-Tumblebug model definitions in `copied-tb-model.go` with enhanced functionality while maintaining 100% backward compatibility and focusing exclusively on structs used by cm-model dependencies.

## Target Version

**CB-Tumblebug Version**: <!-- e.g., v0.11.2, latest, specific commit hash -->

## Synchronization Scope

### Core cm-model Dependencies (Always Required)
- [ ] `TbMciReq` - Multi-cloud infrastructure requests
- [ ] `TbVNetReq` - Virtual network requests  
- [ ] `TbSshKeyReq` - SSH key requests
- [ ] `TbSpecInfo` - VM specification information
- [ ] `TbImageInfo` - VM image information
- [ ] `TbSecurityGroupReq` - Security group requests

### Related Dependencies (Validate if Required)
- [ ] `TbMciInfo` - MCI information responses
- [ ] `TbVmInfo` - VM information responses
- [ ] `TbVmDynamicReq` - Dynamic VM requests
- [ ] `MciCmdReq` - MCI command requests
- [ ] Other related structs (list if any)

## Synchronization Requirements

### ✅ **DO Include**
- Structs directly used by cm-model dependencies
- Enhanced field descriptions and examples  
- Modern cloud resource specifications (latest AMIs, instance types, etc.)
- Comprehensive validation tags
- Network agent monitoring capabilities
- VM lifecycle management improvements

### ❌ **DO NOT Include**
- Independent scaling operations (`TbScaleOutSubGroupReq`)
- Generic status tracking structs (`ResourceStatusInfo`)  
- VM status monitoring structs (`TbVmStatusInfo`)
- MCI status overview structs (`MciStatusInfo`)
- Specialized monitoring functionality unrelated to cm-model

## Validation Checklist

### Before Changes
- [ ] Document current CB-Tumblebug version in `copied-tb-model.go`
- [ ] Identify target CB-Tumblebug version/commit for synchronization
- [ ] Validate which structs are actually used by cm-model dependencies
- [ ] Review existing struct definitions for backward compatibility

### During Synchronization  
- [ ] Update version header comment with target CB-Tumblebug version
- [ ] Preserve all existing cm-model dependency structs
- [ ] Add enhanced field descriptions and modern examples
- [ ] Include source path comments (file paths and line numbers)
- [ ] Maintain JSON serialization tags and validation rules

### After Changes
- [ ] Verify all cm-model dependencies still function:
  ```go
  var recommended cloudmodel.RecommendedVmInfra
  recommended.TargetVmInfra           // TbMciReq ✓
  recommended.TargetVNet              // TbVNetReq ✓  
  recommended.TargetSshKey            // TbSshKeyReq ✓
  recommended.TargetVmSpecList        // []TbSpecInfo ✓
  recommended.TargetVmOsImageList     // []TbImageInfo ✓
  recommended.TargetSecurityGroupList // []TbSecurityGroupReq ✓
  ```
- [ ] Test JSON serialization/deserialization
- [ ] Confirm no circular dependencies introduced
- [ ] Validate backward compatibility maintained
- [ ] Run go mod tidy and verify module dependencies

## Expected Deliverables

1. **Updated `copied-tb-model.go`**:
   - Synchronized structs with enhanced functionality
   - Modern cloud resource examples  
   - Comprehensive field documentation
   - Updated version header comment

2. **Documentation Updates** (if needed):
   - Update related README sections if new capabilities are added
   - Document any breaking changes (should be avoided)

## Success Criteria

- ✅ **100% Backward Compatibility**: All existing cm-model usage continues to work
- ✅ **Enhanced Functionality**: Improved field descriptions and modern examples
- ✅ **Focused Scope**: Only includes structs directly related to cm-model
- ✅ **No Circular Dependencies**: Maintains independence from CB-Tumblebug
- ✅ **Validation**: All tests pass, JSON serialization works correctly

## Reference

- CB-Tumblebug Repository: https://github.com/cloud-barista/cb-tumblebug
- cm-model Dependencies Documentation: See `copilot-instructions.md` CB-Tumblebug Integration section
- Previous Sync Issues: <!-- Link to related issues if any -->

---

**For Maintainers**: This template ensures efficient and consistent CB-Tumblebug model synchronization while maintaining cm-model's focused scope and avoiding unnecessary complexity.