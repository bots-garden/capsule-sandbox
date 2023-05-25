// Package main → a simple http server
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bots-garden/capsule-host-sdk"
	"github.com/bots-garden/capsule-host-sdk/helpers"

	"github.com/tetratelabs/wazero"
)

var wasmFile []byte
var runtime wazero.Runtime
var ctx context.Context

func main() {

	ctx = context.Background()

	// Create a new WebAssembly Runtime.
	runtime = capsule.GetRuntime(ctx)

	// START: host functions
	// Get the builder and load the default host functions
	builder := capsule.GetBuilder(runtime)

	// Add your host functions here
	// 🏠
	// End of of you hostfunction

	// Instantiate builder and default host functions
	_, err := builder.Instantiate(ctx)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// END: host functions

	// This closes everything this Runtime created.
	defer runtime.Close(ctx)

	// Load the WebAssembly module
	args := os.Args[1:]
	wasmFilePath := args[0]
	httpPort := args[1]

	wasmFile, err = helpers.LoadWasmFile(wasmFilePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Registering the hello handler
	http.HandleFunc("/", callWASMFunction)

	fmt.Println("Cracker is listening on", httpPort)

	// Listening on port 8080
	http.ListenAndServe(":"+httpPort, nil)
}

// A handler for "/" route
func callWASMFunction(w http.ResponseWriter, req *http.Request) {

	mod, err := runtime.Instantiate(ctx, wasmFile)
	if err != nil {
		fmt.Fprintf(w, err.Error()+"\n")
	}

	// Get the reference to the WebAssembly function: "callHandle"
	// callHandle is exported by the Capsule plugin
	handleFunction := capsule.GetHandle(mod)
	/* TODO: explain

	 */

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error()+"\n")
	}

	result, err := capsule.CallHandleFunction(ctx, mod, handleFunction, body)

	if err != nil {
		fmt.Fprintf(w, err.Error()+"\n")
	} else {
		fmt.Fprintf(w, string(result)+"\n")
	}

}
