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
	"io/ioutil"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
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
			log.WithField("File", waitfile).Info("Finished reading the contents")
			return contents
		}
	}
}

func main() {
	namespaces := os.Getenv("NAMESPACES")
	if namespaces == "" {
		log.Info("No namespaces provided, please set the NAMESPACES environment variable")
		os.Exit(1)
	}

	namespaceList := strings.Split(namespaces, ",")

	resultsDir := os.Getenv("READ_RESULTS_DIR")
	if resultsDir == "" {
		resultsDir = "/tmp/results"
	}
	log.WithField("resultsDir", resultsDir).Info("Waiting for results to appear in this directory")

	config, err := rest.InClusterConfig()
	if err != nil {
		log.WithError(err).Info("Unable to load in-cluster configuration")
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.WithError(err).Info("Unable to create a new ClientSet for the given config")
		os.Exit(1)
	}

	nsClient := clientset.CoreV1().Namespaces()
	// Wait for the previous container to finish running
	waitForFile(resultsDir + "/done")
	for _, namespace := range namespaceList {
		err = nsClient.Delete(namespace, nil)
		if err != nil {
			log.WithField("namespace", namespace).WithError(err).Info("Could not delete namespace")
		}
	}
}
