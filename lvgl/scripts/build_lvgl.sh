#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cp $SCRIPT_DIR/lv_conf_template.h $SCRIPT_DIR/../ext/lvgl/lv_conf.h
cd $SCRIPT_DIR/../

echo "Building lvgl... ${SCRIPT_DIR}"
mkdir -p ext/lib/lvgl
mkdir -p ext/build/lvgl

# Set SDL2 paths if provided or use Homebrew location as fallback
SDL2_PATH=${SDL2_PATH:-"/opt/homebrew/Cellar/sdl2/2.32.2"}

cd ext/build/lvgl 
cmake ../../lvgl \
    -DCMAKE_INSTALL_PREFIX=../../lib/lvgl \
    -DLV_BUILD_EXAMPLES=ON \
    -DLV_USE_SDL=ON \
    -DSDL2_PATH="${SDL2_PATH}" \
    -DCMAKE_PREFIX_PATH="${SDL2_PATH}" \
    -DCMAKE_OSX_ARCHITECTURES=arm64 \
    -DCMAKE_CXX_FLAGS="-arch arm64 -DGL_SILENCE_DEPRECATION -I${SDL2_PATH}/include" \
    -DCMAKE_C_FLAGS="-arch arm64 -DGL_SILENCE_DEPRECATION -I${SDL2_PATH}/include"
make -j4
make install
echo "lvgl built successfully!"

