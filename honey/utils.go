package honey

import (
    "os"
    "errors"
    "encoding/json"
)

const GOOGLECLOUD_PLATFORM_CREDS = "GCP_CREDS"

var (
    ErrCredsNotFound = errors.New("could not find JSON creds in environment")
)

/*
Check the environment for 
*/
func checkEnvForCreds() (creds []byte, projectId string) {
    var structuredCreds map[string]string
    val, present := os.LookupEnv(GOOGLECLOUD_PLATFORM_CREDS)
    if !present {
        panic(ErrCredsNotFound)
    }
    creds = []byte(val)
    jErr := json.Unmarshal(creds, &structuredCreds)
    if jErr != nil {
        panic(jErr)
    }
    projectId = structuredCreds["project_id"]  // Handle missing key
    return
}
