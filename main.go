/*
Copyright 2018 Heptio Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func waitForFile(waitfile string) []byte {
	// TODO: add maximum amount of waiting time
	for {
		contents, err := ioutil.ReadFile(waitfile) // For read access.
		if err != nil {
			time.Sleep(1 * time.Second)
		} else {
			return contents
		}
	}
}

func main() {
	namespaces := os.Getenv("NAMESPACES")
	if namespaces == "" {
		panic("You must specify the NAMESPACES env variable.")
	}

	namespaceList := strings.Split(namespaces, ",")

	resultsDir := os.Getenv("READ_RESULTS_DIR")
	if resultsDir == "" {
		resultsDir = "/tmp/results"
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	nsClient := clientset.CoreV1().Namespaces()
	// Wait for the previous container to finish running
	waitForFile(resultsDir + "/done")
	for _, namespace := range namespaceList {
		err = nsClient.Delete(namespace, nil)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
