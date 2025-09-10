#!/bin/bash
echo "=== 测试py文件夹重命名后的工作流逻辑 ==="

# 测试1: 检查包目录检测
echo "1. 测试包目录检测逻辑："
if [ -d "py/tqdm" ] && [ -f "py/tqdm/llpkg.cfg" ]; then
    echo "✅ py/tqdm 目录和配置文件存在"
else
    echo "❌ py/tqdm 目录或配置文件不存在"
fi

# 测试2: 检查包类型检测
echo "2. 测试包类型检测："
if [ -f "py/tqdm/llpkg.cfg" ]; then
    PACKAGE_TYPE=$(grep -o '"type":\s*"[^"]*"' "py/tqdm/llpkg.cfg" | cut -d'"' -f4 || echo "")
    echo "tqdm包类型: ${PACKAGE_TYPE:-'C++ (default)'}"
    if [ "$PACKAGE_TYPE" = "python" ]; then
        echo "✅ Python包类型检测正常"
    else
        echo "❌ Python包类型检测异常"
    fi
else
    echo "❌ 无法读取配置文件"
fi

# 测试3: 模拟CI工作流的包检测逻辑
echo "3. 模拟CI工作流的包检测："
MODIFIED_PACKAGES="tqdm"
FILTERED_PACKAGES=""
for package in $MODIFIED_PACKAGES; do
    if [[ "$package" =~ ^(py|\.github|\.git|public|\.DS_Store)$ ]]; then
        echo "跳过非包目录: $package"
        continue
    fi
    
    if [ -d "$package" ] && [ -f "$package/llpkg.cfg" ]; then
        FILTERED_PACKAGES="$FILTERED_PACKAGES $package"
        echo "✅ 在根目录找到包: $package"
    elif [ -d "py/$package" ] && [ -f "py/$package/llpkg.cfg" ]; then
        FILTERED_PACKAGES="$FILTERED_PACKAGES $package"
        echo "✅ 在py目录找到包: $package"
    else
        echo "❌ 未找到包: $package"
    fi
done

echo "过滤后的包: $FILTERED_PACKAGES"

echo "测试完成"
