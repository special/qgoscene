## qgoscene — start QML applications from Go

**qgoscene** provides a simple API for displaying a QML scene from a Go application, functionally similar to embedding Qt's `qmlscene` tool.

Combined with some form of RPC (not part of this package), this can form the basis of a QML interface for a Go application.

qgoscene is _not_ a Qt/QML binding for Go. There is explicitly no control or communication with the QML scene, and API is very limited. [Past attempts](https://github.com/go-qml/qml) at wrapping the full QML API had serious limitations and weren't useful for non-trivial applications.

### Usage

```
go get -u github.com/special/qgoscene
```
Build with `-tags qt6` to link against Qt 6.

```go
package main

import (
	"os"
	"github.com/special/qgoscene"
)

const (
	qml = `
import QtQuick 2.9
import QtQuick.Window 2.2

Window {
	width: 500
	height: 500
	visible: true
	color: "green"
}`
)

func main() {
	qgoscene.NewSceneData(qml, os.Args).Exec()
}

```

### Todo

* Support loading QML from compiled-in resources, including relative files
* Nicer error handling
