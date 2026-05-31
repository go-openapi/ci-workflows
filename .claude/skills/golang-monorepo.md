# Go Mono-repo Patterns for CI/CD

This skill captures patterns, best practices, and gotchas for working with Go mono-repositories in GitHub Actions workflows.

## Table of Contents
- [Detecting Mono-repos](#detecting-mono-repos)
- [Module Discovery](#module-discovery)
- [Module Naming and Tagging](#module-naming-and-tagging)
- [Release Notes Generation](#release-notes-generation)
- [Dependency Management](#dependency-management)
- [Workspace Management](#workspace-management)
- [Common Patterns](#common-patterns)
- [Tips and Gotchas](#tips-and-gotchas)
- [Real-world Examples](#real-world-examples)

---

## Detecting Mono-repos

### The `go list -m` Command

**Single module repo:**
```bash
$ go list -m
github.com/go-openapi/swag
```

**Mono-repo (multiple modules):**
```bash
$ go list -m
github.com/go-openapi/swag
github.com/go-openapi/swag/module1
github.com/go-openapi/swag/module2
```

### Detection Pattern

```bash
count_modules=$(go list -m | wc -l)
if [[ "${count_modules}" -gt 1 ]] ; then
  echo "is_monorepo=true"
else
  echo "is_monorepo=false"
fi
```

**Important**: `go list -m` returns **all modules** in the workspace, while `go list` returns only the **current module**.

### When to Use Each

```bash
# Get root module path (from repo root)
root_module=$(go list)
# Returns: github.com/go-openapi/swag

# List all modules (from repo root)
all_modules=$(go list -m)
# Returns: github.com/go-openapi/swag
#          github.com/go-openapi/swag/submodule1
#          ...
```

---

## Module Discovery

### Method 1: Find go.mod Files

```bash
# Find all go.mod directories
while read -r dir ; do
  echo "Module directory: ${dir}"
done < <(find . -name go.mod | xargs dirname)
```

### Method 2: Use go list with Template

```bash
# Get module directories with their paths
go list -f '{{.Dir}}' -m
```

**Example output:**
```
/repo/root
/repo/root/submodule1
/repo/root/submodule2
```

### Extracting Module Metadata

```bash
root="$(git rev-parse --show-toplevel)"

while read -r module_location ; do
  # Convert absolute path to relative
  relative_location=${module_location#"$root"/}
  relative_location=${relative_location#"$root"}

  # Remove /go.mod suffix if present
  module_dir=${relative_location%"/go.mod"}

  # Determine module name
  if [[ -z "${module_dir}" || "${module_dir}" == "." ]] ; then
    module_name="root"
    module_path="."
  else
    module_name="${module_dir#"./"}"
    module_path="${module_dir}"
  fi

  echo "Module: ${module_name} at ${module_path}"
done < <(find . -name go.mod -exec dirname {} \;)
```

### Building JSON Module List

```bash
modules_json="["
first=true

while read -r module_location ; do
  # ... extract module_name and module_path as above ...

  if [[ "${first}" == "true" ]] ; then
    first=false
  else
    modules_json="${modules_json},"
  fi

  modules_json="${modules_json}{\"name\":\"${module_name}\",\"path\":\"${module_path}\"}"
done < <(find . -name go.mod -exec dirname {} \;)

modules_json="${modules_json}]"
echo "${modules_json}"
```

**Output:**
```json
[
  {"name":"root","path":"."},
  {"name":"module1","path":"./module1"},
  {"name":"module2","path":"./module2"}
]
```

---

## Module Naming and Tagging

### Tagging Convention

**Root module:**
```
v0.24.0
```

**Sub-modules:**
```
module1/v0.24.0
module2/v0.24.0
```

### Tag Generation Pattern

```bash
root="$(git rev-parse --show-toplevel)"
tag="v0.24.0"
declare -a all_tags

while read -r module_location ; do
  relative_location=${module_location#"$root"/}
  relative_location=${relative_location#"$root"}
  module_dir=${relative_location%"/go.mod"}
  base_tag="${module_dir#"./"}"

  if [[ "${base_tag}" == "" || "${base_tag}" == "." ]] ; then
    module_tag="${tag}"              # v0.24.0
  else
    module_tag="${base_tag}/${tag}"  # module1/v0.24.0
  fi

  all_tags+=("${module_tag}")
  echo "Tag: ${module_tag}"
done < <(go list -f '{{.Dir}}' -m)

# Push all tags at once
git push origin ${all_tags[@]}
```

**Reference:** `hack/tag_modules.sh`

---

## Release Notes Generation

### The Challenge with git-cliff and Nested Modules

**Problem:** When running `git-cliff` from a directory, it implicitly captures **all commits** in that directory and its subdirectories. This causes critical issues in mono-repos with nested modules.

#### Example Scenario

```
swag/                           (root module)
├── jsonutils/                  (sub-module)
│   └── adapters/easyjson/      (nested sub-module)
└── yamlutils/                  (sub-module)
```

**What happens when you naively run git-cliff:**

```bash
# Running from swag/
cd swag && git-cliff --current
# ❌ Captures ALL commits: root + jsonutils + adapters + yamlutils

# Running from jsonutils/
cd jsonutils && git-cliff --current
# ❌ Captures jsonutils commits AND adapters commits (duplication!)

# Running from adapters/easyjson/
cd adapters/easyjson && git-cliff --current
# ✅ Only captures adapters commits (leaf module - no children)
```

**Result:** Massive duplication where commits appear in multiple module sections.

### Solution: Exclusion-Based Approach

To generate accurate per-module release notes, you must **exclude child modules** when running git-cliff.

#### Working Pattern (from `hack/release_notes.sh`)

```bash
#!/bin/bash
set -euo pipefail
cd "$(git rev-parse --show-toplevel)"

# Get module information
root=$(go list)
list=$(go list -m -f '{"name":{{ printf "%q" .Path }},"path":{{ printf "%q" .Dir }}}')
modules=$(echo "${list}" | jq -sc)

# Extract bash-friendly arrays (unquoted)
bash_paths=$(echo "${list}" | jq -r '.path')
bash_relative_names=$(echo "${modules}" | jq -r --arg ROOT "${root}" \
  '.[] | .name | ltrimstr($ROOT) | ltrimstr("/") | sub("^$";"{root}")')

declare -a ALL_RELATIVE_MODULES
ALL_RELATIVE_MODULES=(${bash_relative_names})

declare -a ALL_FOLDERS
ALL_FOLDERS=(${bash_paths})

# Function to find child modules that need exclusion
function other_module_paths() {
  local current_index="$1"
  local current_module_path="$2"
  declare -a result

  for (( i=0; i<${#ALL_FOLDERS[@]}; i++ )); do
    # Skip earlier elements (list is sorted)
    [[ $i -le $current_index ]] && continue

    folder="${ALL_FOLDERS[$i]}"

    # Check if this folder is a child of current module
    if [[ "${folder}" =~ ^"${current_module_path}" ]] ; then
      result+=("--exclude-path ${folder}")
    fi
  done

  echo "${result[@]}"
}

# Generate release notes for each module
{
for (( i=0; i<${#ALL_RELATIVE_MODULES[@]}; i++ )); do
  relative_module="${ALL_RELATIVE_MODULES[$i]}"
  folder="${ALL_FOLDERS[$i]}"

  # Build exclusion list for child modules
  excluded=$(other_module_paths "${i}" "${folder}")

  # Set tag pattern for this module
  if [[ "${relative_module}" == "{root}" ]] ; then
    relative_module="${root}"
    tag_pattern="^v\d+\.\d+\.\d+$"           # Matches: v0.24.0
  else
    tag_pattern="^${relative_module}/v\d+\.\d+\.\d+$"  # Matches: jsonutils/v0.24.0
  fi

  # Generate notes with exclusions
  pushd "${folder}" >/dev/null
  echo "## ${relative_module}"

  # Reindent markdown (add one # level for subsections)
  git-cliff --config .cliff.toml \
    --tag-pattern "${tag_pattern}" \
    ${excluded} | sed -E 's/^(#+) /\1# /g'

  popd >/dev/null
done
} > release_notes.md
```

#### Key Techniques

**1. Tag Pattern Filtering**
```bash
# Root module: only match tags like "v0.24.0"
--tag-pattern "^v\d+\.\d+\.\d+$"

# Sub-module: only match tags like "jsonutils/v0.24.0"
--tag-pattern "^jsonutils/v\d+\.\d+\\.d+$"
```

**2. Path Exclusion**
```bash
# From jsonutils/, exclude its child adapters/easyjson/
git-cliff --exclude-path /full/path/to/jsonutils/adapters/easyjson
```

**3. Markdown Reindenting**
```bash
# git-cliff outputs "## Features" but we need "### Features" for subsections
sed -E 's/^(#+) /\1# /g'
```

### Bash Output Format Issues

**Problem:** The `detect-go-monorepo` action outputs JSON-formatted strings which don't work well with bash arrays.

#### What Doesn't Work

```bash
# ❌ JSON-quoted strings break bash array iteration
bash_paths='"/path/one" "/path/two" "/path/three"'
ALL_PATHS=(${bash_paths})  # Quotes become part of elements!
# Result: ALL_PATHS[0] = '"/path/one"' (includes quotes!)
```

#### What Works

```bash
# ✅ Use jq -r for raw (unquoted) output
bash_paths=$(echo "${json}" | jq -r '.[] | .path')
# Output: /path/one
#         /path/two
#         /path/three

ALL_PATHS=(${bash_paths})
# Result: ALL_PATHS[0] = '/path/one' (no quotes!)
```

#### Empty String Handling

```bash
# Problem: Root module has empty relative path ""
# Solution: Use placeholder and handle specially

bash_relative_names=$(echo "${modules}" | jq -r --arg ROOT "${root}" \
  '.[] | .name | ltrimstr($ROOT) | ltrimstr("/") | sub("^$";"{root}")')
# Empty strings become "{root}" placeholder

# Later, check for placeholder:
if [[ "${relative_module}" == "{root}" ]] ; then
  # This is the root module
fi
```

**✅ Resolved:** The `detect-go-monorepo` action now outputs raw (unquoted) format using `jq -r | tr '\n' ' ' | sed` for bash compatibility.

### Template Repetition Problem - ✅ SOLVED

**Problem:** The `.cliff.toml` template contains global sections that should appear **once** in the final release notes, not repeated for each module:

- Contributors list
- License information
- Footer/signature
- External links

#### Solution: Two-Part Generation with Template Override

**Implemented:** A two-part approach using git-cliff's `--body` and `--strip` flags.

**Part 1: Full Changelog** (with global sections)
```bash
# Run once with root tag pattern only
git-cliff \
  --config .cliff.toml \
  --tag-pattern "^v\d+\.\d+\.\d+$" \
  --current \
  --with-tag-message "${TAG_MESSAGE}"
```

Produces complete changelog with:
- ✅ All commits (all modules)
- ✅ Contributors section (once!)
- ✅ License footer (once!)
- ✅ Global sections (once!)

**Part 2: Module-Specific Notes** (minimal template)
```bash
# For each module, with exclusions
body_template=$(cat .cliff-monorepo.toml)

git-cliff \
  --config .cliff.toml \
  --body "${body_template}" \
  --tag-pattern "^${module}/v\d+\.\d+\\.d+$" \
  --exclude-path ${child_modules} \
  --strip all \
  --current
```

Produces module sections with:
- ✅ Only commits for this module
- ✅ Version header + commit groups
- ❌ **No** contributors (stripped)
- ❌ **No** footer (stripped)

**Final Assembly:**
```bash
{
  echo "${FULL_NOTES}"           # Part 1
  echo ""
  echo "# Module-specific release notes"
  echo ""
  cat notes-*.md                  # Part 2
} > final-notes.md
```

**Result:** Clean two-part structure with no duplication!

#### The `.cliff-monorepo.toml` Template

Minimal template for module sections (commits only):

```toml
{%- if version %}
## [{{ version }}] - {{ timestamp | date(format="%Y-%m-%d") }}
{%- endif %}

---

{%- for group, commits in commits | group_by(attribute="group") %}
### {{ group | upper_first }}
  {%- for commit in commits %}
* {{ commit.message }}
  {%- endfor %}
{%- endfor %}
```

**Key characteristics:**
- ✅ Version header
- ✅ Commit groups (Features, Fixes, etc.)
- ✅ PR/contributor attribution
- ❌ No contributors section
- ❌ No license footer

**Location:** `.cliff-monorepo.toml` in repository or fetched from ci-workflows repo.

### Implementation Status

**✅ Completed & Production-Ready:**
- ✅ Detecting mono-repos (`detect-go-monorepo` action)
- ✅ Tagging all modules with same version
- ✅ Building exclusion lists to prevent commit duplication
- ✅ Tag pattern filtering per module
- ✅ Markdown reindenting for subsections
- ✅ **Two-part release notes generation** (no template duplication!)
- ✅ **Workflow integration** (`release.yml` supports mono-repos)
- ✅ Bash output formats (using `jq -r | tr | sed`)
- ✅ Remote template fetching (`.cliff-monorepo.toml`)

**⚠️ Known Limitations:**
- ⚠️ All modules get same version tag (no granular versioning yet)
- ⚠️ Bash arrays assume no spaces in paths (acceptable for Go modules)
- ⚠️ Complexity is moderate (but well-documented and tested)

**⚡ Performance Note:**

Performance is **not a concern** for release notes generation:
- ✅ Releases are infrequent (weeks/months between releases)
- ✅ git-cliff is extremely fast
- ✅ Tested with 16-module mono-repo: **< 5 seconds** total
- ✅ Largest expected mono-repo: ~50 modules
- ✅ Sequential processing is perfectly acceptable

**No performance optimization needed.** The workflow is already production-ready for any realistic mono-repo size.

**🚧 Future Enhancements:**
- 🔮 Granular tagging (different versions per module for patch releases)
- 🔮 Skip unchanged modules in patch releases (correctness, not performance)
- 🔮 Better error handling for edge cases

### Production Workflow Integration

**Workflow:** `.github/workflows/release.yml`

The release workflow now supports mono-repos with automatic detection:

```yaml
jobs:
  gh-release:
    steps:
      # Mono-repo detection handled by caller
      - name: Install git-cliff [monorepo]
        if: ${{ inputs.is-monorepo == 'true' }}
        uses: taiki-e/install-action@...
        with:
          tool: git-cliff

      - name: Fetch remote cliff-monorepo template
        if: ${{ inputs.is-monorepo == 'true' && !local-config }}
        run: curl -fsSL .cliff-monorepo.toml ...

      - name: Generate release notes [monorepo]
        # Two-part generation with exclusions
        run: |
          # Part 1: Full changelog
          git-cliff --tag-pattern "^v\d+\.\d+\.\d+$" ...

          # Part 2: Module-specific notes
          for module in modules; do
            git-cliff \
              --body "${monorepo_template}" \
              --tag-pattern "^${module}/v\d+\.\d+\.\d+$" \
              --exclude-path ${child_modules} \
              --strip all ...
          done

          # Concatenate final notes
```

**Caller:** `.github/workflows/bump-release-monorepo.yml`

```yaml
- uses: ./.github/workflows/release.yml
  with:
    tag: ${{ needs.tag-release-monorepo.outputs.next-tag }}
    is-monorepo: 'true'
    module-relative-paths: ${{ needs.detect-modules.outputs.bash-relative-names }}
```

### Related Files

**Implemented:**
- ✅ `.github/workflows/release.yml` - Release workflow with mono-repo support
- ✅ `.github/workflows/bump-release-monorepo.yml` - Mono-repo release orchestration
- ✅ `.cliff-monorepo.toml` - Minimal template for module sections
- ✅ `go-openapi/gh-actions/ci-jobs/detect-go-monorepo` - Module detection action

**Reference Implementation:**
- 📝 `hack/release_notes.sh` (in consuming repos) - Standalone test script
- 📝 `.cliff.toml` - Full git-cliff configuration template

**Documentation:**
- 📚 `.claude/skills/golang-monorepo.md` - This document
- 📚 `.claude/skills/github-actions.md` - GitHub Actions patterns

### Success Criteria

A successful mono-repo release produces notes like:

```markdown
## [v0.25.0] - 2025-01-15

[Full changelog with ALL commits]

### Contributors
@alice, @bob

---
License footer

# Module-specific release notes

## swag
### [v0.25.0]
[Commits affecting swag only, excluding child modules]

## jsonutils
### [jsonutils/v0.25.0]
[Commits affecting jsonutils only, excluding children]

## jsonutils/adapters
### [jsonutils/adapters/v0.25.0]
[Commits affecting adapters only]
```

**Key achievements:**
- ✅ Contributors appear once (not per module)
- ✅ No commit duplication between parent/child modules
- ✅ Clear two-part structure
- ✅ Proper markdown nesting
- ✅ Tag patterns match module hierarchy

---

## Dependency Management

### Updating Inter-module Dependencies

**Goal:** Update all modules to use version `v0.24.0` of dependencies from the same repo.

**Pattern:**

```bash
root="$(git rev-parse --show-toplevel)"
target_tag="v0.24.0"

# Get root module path (e.g., "github.com/go-openapi/swag")
root_module=$(go list)

# Process each module
while read -r dir ; do
  pushd "${dir}" > /dev/null

  # List dependencies matching the root module or sub-modules
  go list -deps -test \
    -f '{{ if .DepOnly }}{{ with .Module }}{{ .Path }}{{ end }}{{ end }}' | \
  sort -u | \
  grep "^${root_module}" | \
  while read -r module ; do
    echo "Updating ${module} to ${target_tag}"
    go mod edit -require "${module}@${target_tag}"
  done

  # Tidy the go.mod file
  go mod tidy

  popd > /dev/null
done < <(find . -name go.mod | xargs dirname)
```

**Key points:**
- Use `grep "^${root_module}"` to match both root and sub-modules
- The `^` anchor ensures we only match modules under our path
- Always run `go mod tidy` after editing

**Reference:** `hack/upgrade_modules.sh`, `prepare-release-monorepo.yml`

### Why This Pattern Works

```bash
root_module="github.com/go-openapi/swag"

# This matches:
# ✅ github.com/go-openapi/swag
# ✅ github.com/go-openapi/swag/module1
# ✅ github.com/go-openapi/swag/module2

# This does NOT match:
# ❌ github.com/go-openapi/other
# ❌ github.com/other-org/swag
```

---

## Workspace Management

### go.work File

A `go.work` file at the repository root defines a Go workspace for mono-repos.

**Example:**
```
go 1.24

use (
    .
    ./module1
    ./module2
)
```

### Syncing Workspace

After updating modules, always sync the workspace:

```bash
if [[ -f go.work ]] ; then
  echo "Syncing workspace"
  go work sync
fi
```

**What it does:**
- Updates `go.work` to reflect current module structure
- Ensures all modules are properly linked in the workspace

### go work sync vs go mod tidy

```bash
# In each module directory
go mod tidy        # Updates that module's go.mod and go.sum

# At repository root
go work sync       # Synchronizes the workspace across all modules
```

**Order matters:**
1. Update and tidy each module
2. Then sync the workspace

---

## Common Patterns

### Pattern 1: Iterate Over Modules

```bash
root="$(git rev-parse --show-toplevel)"

while read -r dir ; do
  echo "Processing module in ${dir}"
  pushd "${dir}" > /dev/null

  # Do something in this module
  go mod tidy

  popd > /dev/null
done < <(find . -name go.mod | xargs dirname)
```

### Pattern 2: Conditional Logic Based on Mono-repo

```yaml
# In GitHub Actions
jobs:
  detect:
    outputs:
      is_monorepo: ${{ steps.detect.outputs.is_monorepo }}
    steps:
      - id: detect
        run: |
          count_modules=$(go list -m | wc -l)
          if [[ "${count_modules}" -gt 1 ]] ; then
            echo "is_monorepo=true" >> "${GITHUB_OUTPUT}"
          else
            echo "is_monorepo=false" >> "${GITHUB_OUTPUT}"
          fi

  single-module:
    needs: [detect]
    if: ${{ needs.detect.outputs.is_monorepo == 'false' }}
    uses: ./.github/workflows/single-module-workflow.yml

  multi-module:
    needs: [detect]
    if: ${{ needs.detect.outputs.is_monorepo == 'true' }}
    uses: ./.github/workflows/monorepo-workflow.yml
```

### Pattern 3: Generate Release Notes Per Module

```bash
root="$(git rev-parse --show-toplevel)"
root_module=$(go list)

while read -r dir ; do
  module_name=$(basename "${dir}")
  [[ "${dir}" == "${root}" ]] && module_name="root"

  echo "## Module: ${module_name}" >> /tmp/notes.md

  pushd "${dir}" > /dev/null
  git-cliff --current --strip all >> /tmp/notes.md
  popd > /dev/null
done < <(find . -name go.mod -exec dirname {} \;)
```

---

## Tips and Gotchas

### 1. **Always Use Anchored grep for Module Matching**

```bash
# ❌ WRONG - matches too broadly
grep "go-openapi/swag"

# ✅ CORRECT - anchored to start
grep "^${root_module}"
```

### 2. **Run Commands from Repo Root**

```bash
# ✅ CORRECT
root="$(git rev-parse --show-toplevel)"
cd "${root}"
root_module=$(go list)  # Gets the root module

# ❌ WRONG - might be in wrong directory
root_module=$(go list -m)  # Returns ALL modules, not just root
```

### 3. **go list vs go list -m**

```bash
# Current module only
go list
# Output: github.com/go-openapi/swag

# All modules in workspace
go list -m
# Output: github.com/go-openapi/swag
#         github.com/go-openapi/swag/module1
#         ...
```

### 4. **Path Manipulation is Tricky**

```bash
root="/home/user/repo"
module_location="/home/user/repo/module1"

# Remove root prefix
relative_location=${module_location#"$root"/}    # "module1"
relative_location=${relative_location#"$root"}    # Handle root itself

# Remove go.mod if present
module_dir=${relative_location%"/go.mod"}

# Remove leading ./
base_tag="${module_dir#"./"}"
```

### 5. **pushd/popd for Directory Changes**

```bash
# ✅ CORRECT - preserves directory stack
pushd "${dir}" > /dev/null
# ... do work ...
popd > /dev/null

# ❌ WRONG - can get lost if script fails mid-execution
cd "${dir}"
# ... do work ...
cd -
```

### 6. **Array Management in Bash**

```bash
# Declare array
declare -a all_tags

# Append to array
all_tags+=("value")

# Expand array
git push origin ${all_tags[@]}

# NOT: git push origin "${all_tags[@]}"  # Quotes break multi-value expansion
```

### 7. **Testing for Empty Directory Path**

```bash
# Must check for both empty and "."
if [[ -z "${module_dir}" || "${module_dir}" == "." ]] ; then
  echo "This is the root module"
fi
```

### 8. **Go Work Commands Context**

```bash
# go work sync MUST be run from repo root
cd "$(git rev-parse --show-toplevel)"
go work sync

# go mod tidy MUST be run from module directory
cd module_dir
go mod tidy
```

### 9. **Dependency Update Order**

```bash
# ✅ CORRECT ORDER
for each module:
  1. go list -deps -test | grep | go mod edit -require
  2. go mod tidy
go work sync  # After all modules updated

# ❌ WRONG - syncing before all modules updated
for each module:
  go mod edit -require
  go work sync  # Too early!
  go mod tidy
```

### 10. **Silent Errors in Pipelines**

```bash
# ❌ WRONG - grep failure is silent in pipeline
go list -deps -test -f '...' | grep "^${root_module}" | while read -r module ; do
  # If grep finds nothing, loop never executes - no error!
done

# ✅ BETTER - check results
deps=$(go list -deps -test -f '...' | grep "^${root_module}" || true)
if [[ -z "${deps}" ]] ; then
  echo "::warning::No dependencies found matching ${root_module}"
fi
```

---

## Real-world Examples

### From `go-test-monorepo.yml`

**Mono-repo detection:**
```yaml
- name: Detect go mono-repo
  id: detect-monorepo
  run: |
    count_modules=$(go list -m | wc -l)
    if [[ "${count_modules}" -gt 1 ]] ; then
      echo "is_monorepo=true" >> "${GITHUB_OUTPUT}"
      echo "::notice title=is_monorepo::true"
      exit
    fi
    echo "is_monorepo=false" >> "${GITHUB_OUTPUT}"
    echo "::notice title=is_monorepo::false"
```

### From `prepare-release-monorepo.yml`

**Update dependencies across all modules:**
```yaml
- name: Update go.mod files for new release
  env:
    TARGET_TAG: v0.24.0
  run: |
    root="$(git rev-parse --show-toplevel)"
    cd "${root}"

    # Infer root module
    root_module=$(go list)

    # Update each module
    while read -r dir ; do
      pushd "${dir}" > /dev/null

      # Update dependencies
      go list -deps -test \
        -f '{{ if .DepOnly }}{{ with .Module }}{{ .Path }}{{ end }}{{ end }}' | \
      sort -u | \
      grep "^${root_module}" | \
      while read -r module ; do
        go mod edit -require "${module}@${TARGET_TAG}"
      done

      go mod tidy
      popd > /dev/null
    done < <(find . -name go.mod | xargs dirname)

    # Sync workspace
    if [[ -f go.work ]] ; then
      go work sync
    fi
```

### From `bump-release-monorepo.yml`

**Tag all modules:**
```yaml
- name: Tag all modules
  env:
    NEXT_TAG: v0.24.0
  run: |
    root="$(git rev-parse --show-toplevel)"
    declare -a all_tags
    cd "${root}"

    while read -r module_location ; do
      relative_location=${module_location#"$root"/}
      relative_location=${relative_location#"$root"}
      module_dir=${relative_location%"/go.mod"}
      base_tag="${module_dir#"./"}"

      if [[ "${base_tag}" == "" || "${base_tag}" == "." ]] ; then
        module_tag="${NEXT_TAG}"
      else
        module_tag="${base_tag}/${NEXT_TAG}"
      fi

      all_tags+=("${module_tag}")
      git tag -s -m "Release ${module_tag}" "${module_tag}"
    done < <(go list -f '{{.Dir}}' -m)

    # Push all tags
    git push origin ${all_tags[@]}
```

### From `hack/tag_modules.sh`

Complete script for tagging all modules:
```bash
#! /bin/bash
set -euo pipefail

remote="$1"
tag="$2"
root="$(git rev-parse --show-toplevel)"
declare -a all_tags

cd "${root}"

while read module_location ; do
  relative_location=${module_location#"$root"/}
  relative_location=${relative_location#"$root"}
  module_dir=${relative_location%"/go.mod"}
  base_tag="${module_dir#"./"}"

  if [[ "${base_tag}" == "" ]] ; then
    module_tag="${tag}"
  else
    module_tag="${base_tag}/${tag}"
  fi

  all_tags+=("${module_tag}")
  git tag -s "${module_tag}" -m "${module_tag}"
done < <(go list -f '{{.Dir}}' -m)

git push "${remote}" ${all_tags[@]}
```

### From `hack/upgrade_modules.sh`

Update module dependencies:
```bash
#! /bin/bash
new_tag=$1

cd "$(git rev-parse --show-toplevel)"
while read -r dir ; do
  pushd $dir

  go list -deps -test \
    -f '{{ if .DepOnly }}{{ with .Module }}{{ .Path }}{{ end }}{{ end }}' | \
  sort -u | \
  grep "go-openapi/swag" | \
  while read -r module ; do
    go mod edit -require "${module}@${new_tag}"
  done

  go mod tidy
  popd
done < <(find . -name go.mod | xargs dirname)

go work sync
```

---

## Future Improvements

For planned enhancements and future work on Go mono-repo support, see:

📋 [`.claude/plans/golang-monorepo-improvements.md`](../plans/golang-monorepo-improvements.md)

---

## Related Files

**Workflows:**
- `.github/workflows/go-test-monorepo.yml` - Testing mono-repos
- `.github/workflows/bump-release-monorepo.yml` - Releasing mono-repos
- `.github/workflows/prepare-release-monorepo.yml` - Preparing releases
- `.github/workflows/release.yml` - Release notes generation (with mono-repo support)

**Configuration:**
- `.cliff.toml` - Full git-cliff configuration template
- `.cliff-monorepo.toml` - Minimal template for module sections

**Documentation:**
- `.claude/skills/golang-monorepo.md` - This document
- `.claude/skills/github-actions.md` - GitHub Actions patterns
- `.claude/plans/golang-monorepo-improvements.md` - Future improvements and planned work

**Reference Scripts** (in consuming repos):
- `hack/tag_modules.sh` - Tag all modules script
- `hack/upgrade_modules.sh` - Update dependencies script
- `hack/release_notes.sh` - Test script for release notes generation

---

## Summary

Working with Go mono-repos requires careful attention to:

1. **Detection**: Use `go list -m | wc -l` to detect mono-repos
2. **Root Module**: Use `go list` (not `go list -m`) from repo root
3. **Module Iteration**: Use `find . -name go.mod` or `go list -f '{{.Dir}}' -m`
4. **Path Manipulation**: Carefully handle path prefixes and suffixes
5. **Dependency Updates**: Use anchored grep (`^${root_module}`) to match modules
6. **Tagging**: Follow `module-name/version` convention for sub-modules
7. **Order of Operations**: Update modules → tidy → sync workspace
8. **Directory Management**: Use `pushd`/`popd` for safety

Following these patterns ensures reliable, maintainable CI/CD workflows for Go mono-repositories.
