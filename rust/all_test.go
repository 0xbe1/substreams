package rust

import (
	"io/ioutil"
	"math/big"
	"os"
	"testing"

	"github.com/streamingfast/substreams/state"

	"github.com/streamingfast/substreams/wasm"
	"github.com/test-go/testify/require"
)

func TestRustScript(t *testing.T) {
	cases := []struct {
		wasmFile     string
		functionName string
		parameters   []interface{}
		builder      *state.Builder
		assert       func(t *testing.T, module *wasm.Module, instance *wasm.Instance, builder *state.Builder)
	}{
		{
			wasmFile:     "./test/target/wasm32-unknown-unknown/release/testing_substreams.wasm",
			functionName: "test_sum_big_int",
			builder:      state.NewBuilder("builder.name.1", "", "", "", nil),
			assert: func(t *testing.T, module *wasm.Module, instance *wasm.Instance, builder *state.Builder) {
				data, found := builder.GetLast("test.key.1")
				require.True(t, found)
				require.Equal(t, big.NewInt(20).String(), string(data))
			},
		},
	}

	for _, c := range cases {
		t.Run(c.wasmFile, func(t *testing.T) {
			file, err := os.Open(c.wasmFile)
			require.NoError(t, err)
			byteCode, err := ioutil.ReadAll(file)
			require.NoError(t, err)
			module, err := wasm.NewModule(byteCode, c.functionName)
			require.NoError(t, err)

			instance, err := module.NewInstance(c.functionName, nil)
			require.NoError(t, err)
			instance.SetOutputStore(c.builder)
			err = instance.Execute()
			require.NoError(t, err)

			c.assert(t, module, instance, c.builder)
		})
	}
}