package tests

// func StartDockerContainer(t *testing.T) (*client.Client, string) {

// }

// func StopAndRemoveContainer(cli *client.Client, id string, t *testing.T) {

// }

// func Test_RegisterUser(t *testing.T) {
// 	cli, id := StartDockerContainer(t)
// 	defer StopAndRemoveContainer(cli, id, t)

// 	tests := []struct {
// 		Name        string
// 		RequestBody string
// 	}{
// 		{
// 			RequestBody: `{
// 				ALSJDNWERJLSDANZXCVK
// 			}`,
// 		},
// 		{},
// 		{},
// 		{},
// 	}

// 	var httpCli = http.DefaultClient

// 	for i := 0; i < len(tests); i++ {
// 		t.Run()

// 		var (
// 			resp responses.RegisterUserResponses
// 		)

// 		buf := bytes.NewBufferString(tests[i].RequestBody)

// 		httpReq, reqErr := http.NewRequest("POST", "http://localhost:8080/register", buf)

// 		if reqErr != nil {
// 			t.Fatalf("Request create error: %v", reqErr)
// 		}

// 		httpResp, doErr := httpCli.Do(httpReq)

// 		if doErr != nil {
// 			t.Fatalf("Do request error: %v", doErr)
// 		}

// 		respBuf, readErr := io.ReadAll(httpResp.Body)

// 		if readErr != nil {
// 			t.Fatalf("Read response error: %v", readErr)
// 		}

// 		unmarshErr := json.Unmarshal(respBuf, &resp)

// 		if unmarshErr != nil {
// 			t.Fatalf("Unmarshal response error: %v", unmarshErr)
// 		}

// 	}
// }
