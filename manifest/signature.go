package manifest

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
)

type ModuleHash []byte

func HashModule(modules *pbsubstreams.Modules, module *pbsubstreams.Module, graph *ModuleGraph) ModuleHash {
	buf := bytes.NewBuffer(nil)

	startBlockBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(startBlockBytes, module.StartBlock) //at this
	// point start block should have been resolved
	buf.WriteString("start_block")
	buf.Write(startBlockBytes)

	buf.WriteString("kind")
	switch module.Kind.(type) {
	case *pbsubstreams.Module_KindMap_:
		buf.WriteString("map")
	case *pbsubstreams.Module_KindStore_:
		buf.WriteString("store")
	default:
		panic(fmt.Sprintf("invalid module file %T", module.Kind))
	}

	buf.WriteString("code")
	switch m := module.Code.(type) {
	case *pbsubstreams.Module_WasmCode_:
		code := modules.ModulesCode[m.WasmCode.Index]
		buf.Write(code)
		buf.WriteString(m.WasmCode.Entrypoint)
	case *pbsubstreams.Module_NativeCode_:
		// TODO: get some version of the native code from the registry
		// so it can break compatibility when the native code is updated.
		buf.WriteString(m.NativeCode.Entrypoint)
	}

	buf.WriteString("inputs")
	for _, input := range module.Inputs {
		buf.WriteString(inputName(input))
	}

	buf.WriteString("ancestors")
	ancestors, _ := graph.AncestorsOf(module.Name)
	for _, ancestor := range ancestors {
		sig := HashModule(modules, ancestor, graph)
		buf.Write(sig)
	}

	h := sha1.New()
	h.Write(buf.Bytes())

	return h.Sum(nil)
}
func HashModuleAsString(modules *pbsubstreams.Modules, graph *ModuleGraph, module *pbsubstreams.Module) string {
	return hex.EncodeToString(HashModule(modules, module, graph))
}
func inputName(input *pbsubstreams.Module_Input) string {
	switch input.Input.(type) {
	case *pbsubstreams.Module_Input_Store_:
		return "store"
	case *pbsubstreams.Module_Input_Source_:
		return "source"
	case *pbsubstreams.Module_Input_Map_:
		return "map"
	default:
		panic(fmt.Sprintf("invalid input %T", input.Input))
	}
}
