# LVGL Go Bindings

This package provides Go bindings for the LVGL graphics library, allowing you to create embedded GUI applications using Go.

## Prerequisites

Before using these bindings, you need to install the required dependencies:

1. Install SDL2 (for simulator support):
   ```bash
   # macOS
   brew install sdl2
   
   # Ubuntu/Debian
   sudo apt-get install libsdl2-dev
   
   # Fedora
   sudo dnf install SDL2-devel
   ```

## Installation

1. Clone the LVGL C library:
   ```bash
   mkdir ext
   cd ext/
   git clone git@github.com:lvgl/lvgl.git
   cd ../
   ```

2. Build the LVGL library:
   ```bash
   sh scripts/build_lvgl.sh
   ```

## Usage

Here's a simple example of how to use the LVGL Go bindings:

```go
package main

import (
	"github.com/yourusername/llgo/c/lvgl/display"
	"github.com/yourusername/llgo/c/lvgl/widgets"
	"time"
)

func main() {
	// Initialize display
	disp := display.LvDisplayCreate(800, 480)
	display.LvDisplaySetDefault(disp)
	
	// Create a label
	label := widgets.LvLabelCreate(display.LvScreenActive())
	widgets.LvLabelSetText(label, "Hello, LVGL from Go!")
	
	// Main loop
	for {
		// Handle LVGL tasks
		// ...
		time.Sleep(10 * time.Millisecond)
	}
}
```

## API Documentation

The Go bindings follow the same API structure as the original LVGL C library, with the following naming conventions:

- C function `lv_display_create` becomes Go function `display.LvDisplayCreate`
- C function `lv_label_set_text` becomes Go function `widgets.LvLabelSetText`

For detailed API documentation, refer to the [LVGL documentation](https://docs.lvgl.io/).

## Building for Different Platforms

### Desktop Simulator

The default build targets the desktop simulator using SDL2, which is useful for development and testing.

### Embedded Targets

To build for embedded targets, you'll need to:

1. Set up the appropriate cross-compiler for your target platform
2. Modify the build scripts to use the cross-compiler
3. Implement the appropriate display and input drivers for your hardware

## Contributing

Contributions to improve these Go bindings are welcome! Please feel free to submit issues and pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.




