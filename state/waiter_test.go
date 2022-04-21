package state

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/streamingfast/dstore"
	"github.com/streamingfast/substreams/manifest"
	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/require"
)

func TestKVRegex(t *testing.T) {
	filename := "test-12345.kv"
	res := fullKVRegex.FindAllStringSubmatch(filename, 1)

	assert.Greater(t, len(res[0]), 0)
	assert.Equal(t, res[0][1], "12345")
}

func TestPartialRegex(t *testing.T) {
	filename := "test-01234-12345.partial"
	res := partialKVRegex.FindAllStringSubmatch(filename, 1)
	assert.Greater(t, len(res[0]), 0)
	assert.Equal(t, res[0][1], "01234")
	assert.Equal(t, res[0][2], "12345")
}

var testModules = []*pbsubstreams.Module{
	{
		Name:   "A",
		Kind:   &pbsubstreams.Module_KindMap_{KindMap: &pbsubstreams.Module_KindMap{}},
		Inputs: nil,
	},
	{
		Name: "B",
		Kind: &pbsubstreams.Module_KindStore_{KindStore: &pbsubstreams.Module_KindStore{}},
		Inputs: []*pbsubstreams.Module_Input{
			{
				Input: &pbsubstreams.Module_Input_Store_{Store: &pbsubstreams.Module_Input_Store{
					ModuleName: "A",
				}},
			},
		},
	},
	{
		Name: "C",
		Kind: &pbsubstreams.Module_KindStore_{KindStore: &pbsubstreams.Module_KindStore{}},
		Inputs: []*pbsubstreams.Module_Input{
			{
				Input: &pbsubstreams.Module_Input_Store_{Store: &pbsubstreams.Module_Input_Store{
					ModuleName: "A",
				}},
			},
		},
	},
	{
		Name: "D",
		Kind: &pbsubstreams.Module_KindMap_{KindMap: &pbsubstreams.Module_KindMap{}},
		Inputs: []*pbsubstreams.Module_Input{
			{
				Input: &pbsubstreams.Module_Input_Store_{Store: &pbsubstreams.Module_Input_Store{
					ModuleName: "B",
				}},
			},
		},
	},
	{
		Name: "E",
		Kind: &pbsubstreams.Module_KindStore_{KindStore: &pbsubstreams.Module_KindStore{}},
		Inputs: []*pbsubstreams.Module_Input{
			{
				Input: &pbsubstreams.Module_Input_Store_{Store: &pbsubstreams.Module_Input_Store{
					ModuleName: "C",
				}},
			},
		},

		//Inputs: []*manifest.Input{
		//	{
		//		Map:  "C",
		//		Name: "store:C",
		//	},
		//	{
		//		Map:  "D",
		//		Name: "map:D",
		//	},
		//},
	},
}

func TestFileWaiter_Wait(t *testing.T) {
	graph, err := manifest.NewModuleGraph(testModules)
	_ = graph
	assert.NoError(t, err)

	tests := []struct {
		name          string
		graph         *manifest.ModuleGraph
		stores        []*Store
		targetBlock   uint64
		expectedError bool
	}{
		{
			name:  "files all present",
			graph: graph,
			stores: []*Store{
				mustGetWaiterTestStore("B", "module.hash.1", func(ctx context.Context, prefix, ignoreSuffix string, max int) ([]string, error) {
					files := map[string][]string{
						"B-1000": {"B-1000.kv"},
						"B-2000": {"B-2000-1000.partial"},
						"B-3000": {"B-3000-2000.partial"},
					}
					return files[prefix], nil

				}),
				mustGetWaiterTestStore("C", "module.hash.1", func(ctx context.Context, prefix, ignoreSuffix string, max int) ([]string, error) {
					files := map[string][]string{
						"C-1000": {"C-1000.kv"},
						"C-2000": {"C-2000-1000.partial"},
						"C-3000": {"C-3000-2000.partial"},
					}
					return files[prefix], nil

				}),
			},
			targetBlock:   3000,
			expectedError: false,
		},
		{
			name:  "file missing on one store",
			graph: graph,
			stores: []*Store{
				mustGetWaiterTestStore("B", "module.hash.1", func(ctx context.Context, prefix, ignoreSuffix string, max int) ([]string, error) {
					files := map[string][]string{
						"B-1000": {"B-1000.kv"},
						"B-2000": {"B-2000-1000.partial"},
						"B-3000": {"B-3000-2000.partial"},
					}

					return files[prefix], nil

				}),
				mustGetWaiterTestStore("C", "module.hash.1", func(ctx context.Context, prefix, ignoreSuffix string, max int) ([]string, error) {
					files := map[string][]string{
						"C-1000": {"C-1000.kv"},
						"C-3000": {"C-3000-2000.partial"},
					}
					return files[prefix], nil
				}),
			},
			targetBlock:   3000,
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			waiter := NewFileWaiter(test.targetBlock, test.stores)

			ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
			defer cancel()

			err = waiter.Wait(ctx, test.targetBlock, 1000)
			if test.expectedError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func Test_pathToState(t *testing.T) {
	tests := []struct {
		name             string
		store            *Store
		storeName        string
		moduleStartBlock uint64
		targetBlock      uint64
		expectedOk       bool
		expectedFiles    []string
		expectedError    bool
	}{
		{
			name:      "happy path",
			storeName: "A",
			store: mustGetWaiterTestStore("A", "module.hash.1", func(ctx context.Context, prefix, ignoreSuffix string, max int) ([]string, error) {
				files := map[string][]string{
					"A-1000": {"A-1000-0.kv"},
					"A-2000": {"A-2000-1000.partial"},
					"A-3000": {"A-3000-2000.partial"},
				}
				return files[prefix], nil

			}),
			moduleStartBlock: 0,
			targetBlock:      3000,
			expectedOk:       true,
			expectedFiles:    []string{"A-1000-0.kv", "A-2000-1000.partial", "A-3000-2000.partial"},
			expectedError:    false,
		},
		{
			name:      "happy path all partial with start block",
			storeName: "A",
			store: mustGetWaiterTestStore("A", "module.hash.1", func(ctx context.Context, prefix, ignoreSuffix string, max int) ([]string, error) {
				files := map[string][]string{
					"A-2000": {"A-2000-1000.partial"},
					"A-3000": {"A-3000-2000.partial"},
				}
				return files[prefix], nil

			}),
			moduleStartBlock: 1000,
			targetBlock:      3000,
			expectedOk:       true,
			expectedFiles:    []string{"A-2000-1000.partial", "A-3000-2000.partial"},
			expectedError:    false,
		},
		{
			name:      "happy path take shortest path",
			storeName: "A",
			store: mustGetWaiterTestStore("A", "module.hash.1", func(ctx context.Context, prefix, ignoreSuffix string, max int) ([]string, error) {
				files := map[string][]string{
					"A-2000": {"A-2000-1000.partial"},
					"A-3000": {"A-3000-0.kv", "module.hash.1-A-3000-2000.partial"},
				}
				return files[prefix], nil
			}),
			moduleStartBlock: 1000,
			targetBlock:      3000,
			expectedOk:       true,
			expectedFiles:    []string{"A-3000-0.kv"},
			expectedError:    false,
		},
		{
			name:      "happy path take shortest path part 2",
			storeName: "A",
			store: mustGetWaiterTestStore("A", "module.hash.1", func(ctx context.Context, prefix, ignoreSuffix string, max int) ([]string, error) {
				files := map[string][]string{
					"A-2000": {"A-2000-1000.partial"},
					"A-3000": {"A-3000-2000.partial", "A-3000-0.kv"},
				}
				return files[prefix], nil
			}),
			moduleStartBlock: 1000,
			targetBlock:      3000,
			expectedOk:       true,
			expectedFiles:    []string{"A-3000-0.kv"},
			expectedError:    false,
		},
		{
			name:      "conflicting partial files",
			storeName: "A",
			store: mustGetWaiterTestStore("A", "module.hash.1", func(ctx context.Context, prefix, ignoreSuffix string, max int) ([]string, error) {
				files := map[string][]string{
					"A-1000": {"A-1000-0.kv"},
					"A-2000": {"A-2000-1000.partial"},
					"A-3000": {"A-3000-1000.partial", "module.hash.1-A-3000-2000.partial"},
				}
				return files[prefix], nil
			}),
			moduleStartBlock: 0,
			targetBlock:      3000,
			expectedOk:       true,
			expectedFiles:    nil,
			expectedError:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			files, err := pathToState(context.TODO(), test.store, test.targetBlock, test.moduleStartBlock)
			assert.Equal(t, test.expectedFiles, files)
			assert.Equal(t, test.expectedError, err != nil)
		})
	}
}

func mustGetWaiterTestStore(moduleName string, moduleHash string, listFilesFunc func(ctx context.Context, prefix, ignoreSuffix string, max int) ([]string, error)) *Store {
	mockDStore := &dstore.MockStore{
		ListFilesFunc: listFilesFunc,
	}
	s, err := NewStore(moduleName, moduleHash, 0, mockDStore)
	if err != nil {
		panic(fmt.Sprintf("faild to create mock store: %s", err))
	}
	return s
}
