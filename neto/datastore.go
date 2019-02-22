// Copyright 2019 The LootBag Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neto

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

// DatastoreEmulator spawn datastore emulator
//
//   Input:
//     projectID -- Google Cloud project ID
//     port -- datastore emulator port
//
//   Output:
//     stop -- function to stop emulator and clean up data
//
//   Example:
//     stop := DatastoreEmulator("epop-web", "8081")
//     defer stop()
//
//   Check for background processes with:
//     ps -eF | grep datastore-emulator
//
func DatastoreEmulator(projectID, port string) (stop func()) {

	// create command
	buf := new(bytes.Buffer)
	cmd := exec.Command("gcloud", "beta", "emulators", "datastore", "start", "--no-legacy", "--no-store-on-disk", "--host-port=localhost:"+port)
	cmd.Stdout = buf
	cmd.Stderr = buf
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} // grandchild process IDS get -IDS of parents

	// stop kills datastore emulator process
	stop = func() {
		err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		if err != nil {
			log.Printf("############ cannot kill emulator processes: %v\n", err)
			return
		}
		log.Printf("############ datastore emulator stopped\n")
	}

	// define environment variables
	envvars := []string{
		"DATASTORE_DATASET=" + projectID,
		"DATASTORE_PROJECT_ID=" + projectID,
		"DATASTORE_EMULATOR_HOST=localhost:" + port,
		"DATASTORE_EMULATOR_HOST_PATH=localhost:" + port + "/datastore",
		"DATASTORE_HOST=http://localhost:" + port,
	}

	// set environment variables for the background process
	cmd.Env = append(os.Environ(), envvars...)

	// set environment variables for this process
	for _, pair := range envvars {
		words := strings.Split(pair, "=")
		os.Setenv(words[0], words[1])
	}

	// spawn emulator processes
	log.Printf("############ starting datastore emulator\n")
	err := cmd.Start()
	if err != nil {
		log.Printf("############ cannot start emulator: %v\n\t%v\n", err, buf.String())
		stop()
		os.Exit(1)
	}

	// reset database
	log.Printf("############ resetting datastore\n")
	retry, numberOfRetries := 0, 10
	for retry = 1; retry <= numberOfRetries; retry++ {
		time.Sleep(1000 * time.Millisecond)
		response, err := http.Post("http://localhost:"+port+"/reset", "application/json", nil)
		if response != nil && err == nil {
			if response.StatusCode == http.StatusOK {
				break
			}
		}
	}
	if retry >= numberOfRetries {
		log.Printf("############ cannot reset datastore emulator\n")
		stop()
		os.Exit(1)
	}
	return
}
