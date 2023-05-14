// Package main => a simple Capsule Module
package main

import (
    capsule "github.com/bots-garden/capsule-module-sdk"
)

func main() {
    capsule.SetHandle(Handle)
}

// Handle function
func Handle(params []byte) ([]byte, error) {
	
    // Display the content of `params`
    capsule.Print("📝 module parameter(s): " + string(params))

    return []byte("👋 Hello " + string(params)), nil
}
