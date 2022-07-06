package manage

// import (
//     "testing"
//     "encoding/json"
//     )
//
// func TestRenameMachine(t *testing.T) {
//     t.Run("file exists" , func(t *testing.T) {
// 		runners := []Runner{
// 			{
// 				Name:          "test",
// 				DrivePath:     "test.img",
// 				SystemCommand: "qemu-system-x86_64",
// 				MemSize:       "",
// 				CpuCores:      "",
// 				Iso:           "",
// 			},
// 		}
//         files := make( map[string]([]byte))
// 		for _, runnner := range runners {
// 			content, err := json.Marshal(runnner)
// 			if err != nil {
// 				t.Fatalf("Failed to marshal runner: %v", err)
// 			}
// 			files[runnner.Name] = content
// 		}
//
// 		path, tearDown := setupTestComplex(
// 			t,
// 			files,
// 		)
// 		defer tearDown(t)
//
// 		previousExecProvider := ExecProvider
// 		previousConfigProvider := ConfigProvider
//
// 		ExecProvider = &TestExecutor{
// 			errorExecute: false,
// 		}
// 		ConfigProvider = &TestConfig{
// 			machinepath: path,
// 		}
// 		for _, want := range runners {
//             RenameMachine("test" , "newtest")
//             if *got != want {
//                 t.Errorf("got %v , wanted %v" , got , want)
//             }
//
// 		}
// 		ExecProvider = previousExecProvider
// 		ConfigProvider = previousConfigProvider
//
//
//     })
// }
