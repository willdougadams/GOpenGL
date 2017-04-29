package main

import (
  "Application"
  "runtime"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
  app := new(Application.Application).Init()
  app.Run()
}
