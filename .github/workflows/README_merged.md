# 合并后的 Post-processing 工作流说明

## 概述

这个合并后的 `postprocessing_merged.yml` 工作流结合了原有的 C++ 包处理逻辑和新的 Python 包处理逻辑，通过智能检测 `llpkg.cfg` 文件中的 `type` 字段来决定执行哪种处理流程。

## 工作流程

### 1. 包类型检测

工作流会读取每个包的 `llpkg.cfg` 文件，检测 `type` 字段：

```json
{
  "type": "python",  // 如果存在此字段且值为 "python"，则按 Python 包处理
  "upstream": {
    "package": {
      "name": "package_name",
      "version": "1.0.0"
    }
  }
}
```

- **Python 包**：`type` 字段存在且值为 `"python"`
- **C++ 包**：`type` 字段不存在或值不是 `"python"`（默认）

### 2. 处理逻辑

#### Python 包处理流程
1. **环境准备**：
   - 安装 LLGo 和 LLPyg
   - 设置 Python 环境
   - 安装 Python 相关依赖

2. **包处理**：
   - 执行 `llpkgstore postprocessing`
   - 创建 `Python_package/` 目录结构
   - 将包文件复制到 `Python_package/package_name/` 目录
   - 复制 `llpkgstore.json` 到 `Python_package/` 目录

3. **文件结构**：
   ```
   Python_package/
   ├── llpkgstore.json
   └── package_name/
       ├── llpkg.cfg
       ├── package.go
       └── ... (其他包文件)
   ```

#### C++ 包处理流程
1. **环境准备**：
   - 安装 Conan 和 C++ 相关依赖
   - 设置 C++ 构建环境

2. **包处理**：
   - 执行 `llpkgstore postprocessing`
   - 文件直接上传到 `llpkg` 目录下

3. **网站更新**：
   - 更新 `website` 分支
   - 部署到 GitHub Pages

### 3. 智能包发现

工作流使用智能包发现逻辑，只处理当前提交中实际修改的包：

```bash
# 检测修改的包目录
MODIFIED_PACKAGES=$(git diff --name-only HEAD~1 HEAD | grep -E '^[^/]+/' | cut -d'/' -f1 | sort -u)

# 或者从 commit 消息中提取
if [[ $COMMIT_MSG =~ Release-as:\ ([^/]+)/ ]]; then
  PACKAGE_NAME="${BASH_REMATCH[1]}"
  MODIFIED_PACKAGES="$PACKAGE_NAME"
fi
```

### 4. 条件执行

工作流会根据包类型执行不同的步骤：

- **Python 包**：执行 Python 相关的环境设置和处理步骤
- **C++ 包**：执行 C++ 相关的环境设置和网站更新步骤
- **混合包**：同时支持两种类型的包处理

## 使用示例

### Python 包示例

```json
// llpkg.cfg
{
  "type": "python",
  "upstream": {
    "installer": {
      "name": "pip"
    },
    "package": {
      "name": "numpy",
      "version": "1.26.4"
    }
  },
  "llpyg": {
    "output_dir": "./bindings",
    "mod_name": "github.com/goplus/lib/py/numpy",
    "mod_depth": 2
  }
}
```

**处理结果**：
- 包文件上传到 `Python_package/numpy/` 目录
- 版本记录保存在 `Python_package/llpkgstore.json`

### C++ 包示例

```json
// llpkg.cfg
{
  "upstream": {
    "package": {
      "name": "bzip3",
      "version": "1.5.1"
    }
  }
}
```

**处理结果**：
- 包文件直接上传到 `llpkg` 目录
- 更新 `website` 分支和 GitHub Pages

## 优势

1. **向后兼容**：保持原有 C++ 包处理逻辑不变
2. **智能检测**：自动识别包类型，无需手动配置
3. **统一管理**：一个工作流处理所有类型的包
4. **灵活扩展**：易于添加新的包类型支持
5. **清晰分离**：Python 包和 C++ 包有独立的处理路径

## 注意事项

1. **依赖管理**：确保所有必要的依赖都已安装
2. **权限设置**：确保 GitHub Actions 有足够的权限访问仓库
3. **分支管理**：Python 包和 C++ 包使用不同的分支策略
4. **错误处理**：工作流包含适当的错误处理和日志记录

## 故障排除

如果遇到问题，请检查：

1. `llpkg.cfg` 文件格式是否正确
2. 包类型检测是否准确
3. 环境依赖是否完整
4. 权限设置是否正确
5. 分支状态是否正常
