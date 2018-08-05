package qgoscene

// #cgo pkg-config: Qt5Gui Qt5Quick
// #include "scene.h"
import "C"
import (
	"runtime"
	"unsafe"
)

// Scene creates and runs a QML scene and Qt application instance.
// Scene explicitly does not wrap or provide any interaction between
// the QML scene and the Go application -- it is purely a way to load
// QML applications within a Go application. With some form of RPC,
// this can be used to build full QML applications with Go backends.
//
// Only one scene may be created in the lifetime of an application,
// and it may be executed only once.
type Scene struct {
	Args []string

	path string
	data string
}

// LoadScene creates a Qt application with the commandline arguments
// from `args`, which will load the QML file at `path`.
func NewScene(path string, args []string) *Scene {
	s := &Scene{
		Args: args,
		path: path,
	}
	s.createApplication()
	C.createEngine()
	return s
}

// LoadSceneData creates a Qt application with the commandline arguments
// from `args` which will load the QML data from `data`.
func NewSceneData(data string, args []string) *Scene {
	s := &Scene{
		Args: args,
		data: data,
	}
	s.createApplication()
	C.createEngine()
	return s
}

// AddImportPath appends a path to the QML engine's import path list,
// equivalent to QQmlEngine::addImportPath.
func (s *Scene) AddImportPath(path string) {
	C.addImportPath(C.CString(path))
}

// SetImportPathList sets the QML engine's import path list, equivalent
// to QQmlEngine::setImportPathList.
func (s *Scene) SetImportPathList(paths []string) {
	n, p := cStringList(paths)
	C.setImportPathList(n, p)
}

// SetContextProperty assigns a string property in the QML root context,
// where it will be exposed to the QML application. Other than commandline
// arguments, this is the only method of passing information into QML.
//
// SetContextProperty cannot be used after Exec() has been called.
func (s *Scene) SetContextProperty(name, value string) {
	C.setContextProperty(C.CString(name), C.CString(value))
}

// Exec executes the Qt application. Exec will block until quit and
// return the exit code from QGuiApplication. After Exec returns, it is
// not possible to load another scene.
func (s *Scene) Exec() int {
	if s.path != "" {
		C.loadScene(C.CString(s.path))
	} else if s.data != "" {
		C.loadSceneData(C.CString(s.data))
	} else {
		return 1
	}

	return int(C.execApplication())
}

// Quit causes the Qt application to quit and Exec() to return. Quit does
// not wait for Exec to return.
func (s *Scene) Quit() {
	C.quitApplication()
}

func (s *Scene) createApplication() {
	runtime.LockOSThread()

	argc, argv := cStringList(s.Args)
	C.createApplication(argc, argv)
}

func cStringList(in []string) (C.int, **C.char) {
	out := C.malloc(C.size_t(len(in)) * C.size_t(unsafe.Sizeof(uintptr(0))))
	a := (*[1<<30 - 1]*C.char)(out)
	for i, v := range in {
		a[i] = C.CString(v)
	}
	return C.int(len(in)), (**C.char)(out)
}
