# ffmpeg-dev-go
Go bindings for FFmpeg.

## Supported FFmpeg Version

- [x] [4.4](https://github.com/qrtc/ffmpeg-dev-go/tree/4.4)
- [x] [5.0](https://github.com/qrtc/ffmpeg-dev-go/tree/5.0)
- [x] [5.1](https://github.com/qrtc/ffmpeg-dev-go/tree/5.1)
- [x] [6.0](https://github.com/qrtc/ffmpeg-dev-go/tree/6.0)

Reference: [https://ffmpeg.org/download.html#releases](https://ffmpeg.org/download.html#releases)

## How to use

### Step 1: Prepare

- macOS

	```shell
	brew install pkg-config ffmpeg
	```
- Debian

	```shell
	sudo apt install build-essential pkg-config \
		libavdevice-dev libavformat-dev libavfilter-dev libavcodec-dev \
		libpostproc-dev libswscale-dev libswresample-dev libavutil-dev
	```
- Custom

	```shell
	export PKG_CONFIG_PATH="<CUSTOM_FFMPEG_LIBRARY_PATH>/lib/pkgconfig"
	```

### Step 2: Detecte FFmpeg version

```go
package main

/*
#cgo pkg-config: libavutil
#include <libavutil/ffversion.h>
*/
import "C"
import "fmt"

func main() {
	fmt.Println(string(C.FFMPEG_VERSION)[:3])
}
```

### Step 3: Get ffmpeg-go-dev

```shell
go get github.com/qrtc/ffmpeg-dev-go@<FFMPEG_VERSION>
```

### Step 4: Verify

```go
package main

import (
	"fmt"

	"github.com/qrtc/ffmpeg-dev-go"
)

func main() {
	fmt.Println(ffmpeg.AvVersionInfo())
}
```