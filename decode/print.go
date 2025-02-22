package decode

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/streamingfast/bstream"
	"github.com/streamingfast/substreams"
	"github.com/streamingfast/substreams/manifest"
	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewPrintReturnHandler(manif *manifest.Manifest, fileDescs []*desc.FileDescriptor, outputStreamNames []string) substreams.ReturnFunc {
	decodeMsgTypes := map[string]func(in []byte) string{}
	msgTypes := map[string]string{}

	for _, mod := range manif.Modules {
		for _, outputStreamName := range outputStreamNames {
			if mod.Name == outputStreamName {
				var msgType string
				var isStore bool
				if mod.Kind == "store" {
					isStore = true
					msgType = mod.ValueType
				} else {
					msgType = mod.Output.Type
				}
				msgType = strings.TrimPrefix(msgType, "proto:")

				msgTypes[mod.Name] = msgType

				var msgDesc *desc.MessageDescriptor
				for _, file := range fileDescs {
					msgDesc = file.FindMessage(msgType) //todo: make sure it works relatively-wise
					if msgDesc != nil {
						break
					}
				}

				decodeMsgType := func(in []byte) string {
					if msgDesc == nil {
						return "(unknown proto schema) " + decodeAsString(in)
					}
					msg := dynamic.NewMessageFactoryWithDefaults().NewDynamicMessage(msgDesc)
					if err := msg.Unmarshal(in); err != nil {
						fmt.Printf("error unmarshalling protobuf %s to map: %s\n", msgType, err)
						return decodeAsString(in)
					}

					cnt, err := msg.MarshalJSONIndent()
					if err != nil {
						fmt.Printf("error encoding protobuf %s into json: %s\n", msgType, err)
						return decodeAsString(in)
					}

					return string(cnt)
				}

				if isStore {
					if msgDesc != nil {
						decodeMsgTypeWithIndent := func(in []byte) string {
							out := decodeMsgType(in)
							return strings.Replace(out, "\n", "\n    ", -1)
						}
						decodeMsgTypes[mod.Name] = decodeMsgTypeWithIndent
					} else {
						if msgType == "bytes" {
							decodeMsgTypes[mod.Name] = decodeAsHex
						} else {
							// bigint, bigfloat, int64, float64, string
							decodeMsgTypes[mod.Name] = decodeAsString
						}
					}

				} else {
					decodeMsgTypes[mod.Name] = decodeMsgType
				}

			}
		}
	}

	return func(output *pbsubstreams.BlockScopedData) error {
		printClock(output)
		if output == nil {
			return nil
		}
		if len(output.Outputs) == 0 {
			return nil
		}

		for _, out := range output.Outputs {
			for _, log := range out.Logs {
				fmt.Printf("%s: log: %s\n", out.Name, log)
			}

			switch data := out.Data.(type) {
			case *pbsubstreams.ModuleOutput_MapOutput:
				if len(data.MapOutput.Value) != 0 {
					decodeValue := decodeMsgTypes[out.Name]
					msgType := msgTypes[out.Name]
					if decodeValue != nil {
						cnt := decodeValue(data.MapOutput.GetValue())

						fmt.Printf("%s: message %q: %s\n", out.Name, msgType, cnt)
					} else {
						fmt.Printf("%s: message %q: ", out.Name, msgType)

						marshalledBytes, err := protojson.Marshal(data.MapOutput)
						if err != nil {
							return fmt.Errorf("return handler: marshalling: %w", err)
						}

						fmt.Println(marshalledBytes)
					}
				}

			case *pbsubstreams.ModuleOutput_StoreDeltas:
				if len(data.StoreDeltas.Deltas) != 0 {
					fmt.Printf("%s: store deltas:\n", out.Name)
					decodeValue := decodeMsgTypes[out.Name]
					for _, delta := range data.StoreDeltas.Deltas {
						fmt.Printf("  %s (%d) KEY: %q\n", delta.Operation.String(), delta.Ordinal, delta.Key)

						fmt.Printf("    OLD: %s\n", decodeValue(delta.OldValue))
						fmt.Printf("    NEW: %s\n", decodeValue(delta.NewValue))
					}
				}

			default:
				panic("unsupported module output data type")
			}
		}
		return nil
	}
}

func decodeAsString(in []byte) string { return fmt.Sprintf("%q", string(in)) }
func decodeAsHex(in []byte) string    { return "(hex) " + hex.EncodeToString(in) }

func printClock(block *pbsubstreams.BlockScopedData) {
	fmt.Printf("----------- BLOCK: %d (%s) ---------------\n", block.Clock.Number, stepFromProto(block.Step))
}

func stepFromProto(step pbsubstreams.ForkStep) bstream.StepType {
	switch step {
	case pbsubstreams.ForkStep_STEP_NEW:
		return bstream.StepNew
	case pbsubstreams.ForkStep_STEP_UNDO:
		return bstream.StepUndo
	case pbsubstreams.ForkStep_STEP_IRREVERSIBLE:
		return bstream.StepIrreversible
	}
	return bstream.StepType(0)
}
