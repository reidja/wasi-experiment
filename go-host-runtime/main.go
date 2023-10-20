package main

import (
	"context"
	_ "embed"
	"log"
	"os"

	wcache "github.com/reidja/wasi_demo/host/cache"
	wconsole "github.com/reidja/wasi_demo/host/console"
	wlog "github.com/reidja/wasi_demo/host/log"
	wrequest "github.com/reidja/wasi_demo/host/request"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {
	log.Println("starting jengine")
	if len(os.Args) != 2 {
		log.Panicf("expects path to wasm file")
	}
	wasmAppName := os.Args[1]
	wasmApp, err := os.ReadFile(wasmAppName)
	if err != nil {
		log.Panicf("unable to read app.wasm: %v", err)
	}

	ctx := context.Background()

	runtime := wazero.NewRuntimeWithConfig(
		ctx,
		wazero.NewRuntimeConfig(),
	)
	defer runtime.Close(ctx)

	// allow stdout
	config := wazero.NewModuleConfig().WithStdout(os.Stdout)

	guest, err := runtime.CompileModule(ctx, wasmApp)
	if err != nil {
		log.Fatalf("error compiling wasm binary: %v", err)
	}

	wasi_snapshot_preview1.MustInstantiate(ctx, runtime)
	wconsole.InstantiateModule(ctx, runtime)
	wlog.InstantiateModule(ctx, runtime)
	wrequest.InstantiateModule(ctx, runtime)
	wcache.InstantiateModule(ctx, runtime)

	_, err = runtime.InstantiateModule(ctx, guest, config)
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}
}
