package runner

// var fileMap = make(map[string][]byte)
// type TestConfig struct {
// 	ImageDir string
// 	MachineDir string
// }
// func getTestConfig(t testing.TB) *TestConfig {
// 	return nil
// }
// func TestFindMachine(t *testing.T) {
// }
// func populateFakeFS(t testing.TB)  {
// 	t.Helper()
// 	imagedir , machinedir := constructFakeFS(t)
// }
// func constructFakeFS(t testing.TB) (string , string) {
// 	t.Helper()
// 	imagedir , err := ioutil.TempDir("" , "images")
// 	if err != nil {
// 		t.Fatalf("Can't test temp dir %v" , err)
// 	}
// 	defer os.RemoveAll(imagedir)
// 	machinedir , err := ioutil.TempDir("" , "machines")
// 	defer os.RemoveAll(machinedir)
// 	if err != nil {
// 		t.Fatalf("Can't test temp dir %v" , err)
// 	}

// 	return imagedir , machinedir
// }
