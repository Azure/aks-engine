// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package job

import (
	"bufio"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"text/template"
	"time"

	"github.com/Azure/aks-engine/test/e2e/engine"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/pod"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

// List is a container that holds all jobs returned from doing a kubectl get jobs
type List struct {
	Jobs []Job `json:"items"`
}

// Job is used to parse data from kubectl get jobs
type Job struct {
	Metadata pod.Metadata `json:"metadata"`
	Spec     Spec         `json:"spec"`
	Status   Status       `json:"status"`
}

// Spec holds job spec metadata
type Spec struct {
	Completions int `json:"completions"`
	Parallelism int `json:"parallelism"`
}

// Status holds job status information
type Status struct {
	Active    int `json:"active"`
	Succeeded int `json:"succeeded"`
}

// CreateJobFromFile will create a Job from file with a name
func CreateJobFromFile(filename, name, namespace string) (*Job, error) {
	cmd := exec.Command("k", "create", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create Job %s:%s\n", name, string(out))
		return nil, err
	}
	job, err := Get(name, namespace)
	if err != nil {
		log.Printf("Error while trying to fetch Job %s:%s\n", name, err)
		return nil, err
	}
	return job, nil
}

// CreateWindowsJobFromTemplate will create a Job from file with a name
func CreateWindowsJobFromTemplate(filename, name, namespace string, windowsTestImages *engine.WindowsTestImages) (*Job, error) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		return nil, err
	}

	tempfile, err := ioutil.TempFile("", "*.yaml")
	if err != nil {
		return nil, err
	}
	defer tempfile.Close()

	w := bufio.NewWriter(tempfile)
	err = t.Execute(w, windowsTestImages)
	if err != nil {
		return nil, err
	}
	w.Flush()

	return CreateJobFromFile(tempfile.Name(), name, namespace)
}

// CreateWindowsJobFromTemplateDeleteIfExists will create a Job from file, deleting any pre-existing job with the same name
func CreateWindowsJobFromTemplateDeleteIfExists(filename, name, namespace string, windowsTestImages *engine.WindowsTestImages) (*Job, error) {
	j, err := Get(name, namespace)
	if err == nil {
		err := j.Delete(util.DefaultDeleteRetries)
		if err != nil {
			return nil, err
		}
		_, err = WaitOnDeleted(j.Metadata.Name, j.Metadata.Namespace, 5*time.Second, 1*time.Minute)
		if err != nil {
			return nil, err
		}
	}
	return CreateWindowsJobFromTemplate(filename, name, namespace, windowsTestImages)
}

// CreateJobFromFileDeleteIfExists will create a Job from file, deleting any pre-existing job with the same name
func CreateJobFromFileDeleteIfExists(filename, name, namespace string) (*Job, error) {
	j, err := Get(name, namespace)
	if err == nil {
		err := j.Delete(util.DefaultDeleteRetries)
		if err != nil {
			return nil, err
		}
		_, err = WaitOnDeleted(j.Metadata.Name, j.Metadata.Namespace, 5*time.Second, 1*time.Minute)
		if err != nil {
			return nil, err
		}
	}
	return CreateJobFromFile(filename, name, namespace)
}

// GetAll will return all jobs in a given namespace
func GetAll(namespace string) (*List, error) {
	cmd := exec.Command("k", "get", "jobs", "-n", namespace, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	jl := List{}
	err = json.Unmarshal(out, &jl)
	if err != nil {
		log.Printf("Error unmarshalling jobs json:%s\n", err)
		return nil, err
	}
	return &jl, nil
}

// GetAllByPrefix will return all jobs in a given namespace that match a prefix
func GetAllByPrefix(prefix, namespace string) ([]Job, error) {
	jl, err := GetAll(namespace)
	if err != nil {
		return nil, err
	}
	jobs := []Job{}
	for _, j := range jl.Jobs {
		matched, err := regexp.MatchString(prefix+"-.*", j.Metadata.Name)
		if err != nil {
			log.Printf("Error trying to match pod name:%s\n", err)
			return nil, err
		}
		if matched {
			jobs = append(jobs, j)
		}
	}
	return jobs, nil
}

// Get will return a job with a given name and namespace
func Get(jobName, namespace string) (*Job, error) {
	cmd := exec.Command("k", "get", "jobs", jobName, "-n", namespace, "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	j := Job{}
	err = json.Unmarshal(out, &j)
	if err != nil {
		log.Printf("Error unmarshalling jobs json:%s\n", err)
		return nil, err
	}
	return &j, nil
}

// AreAllJobsCompleted will return true if all jobs with a common prefix in a given namespace are in a Completed State
func AreAllJobsCompleted(jobPrefix, namespace string) (bool, error) {
	jl, err := GetAll(namespace)
	if err != nil {
		return false, err
	}

	var status []bool
	for _, job := range jl.Jobs {
		matched, err := regexp.MatchString(jobPrefix, job.Metadata.Name)
		if err != nil {
			log.Printf("Error trying to match job name:%s\n", err)
			return false, err
		}
		if matched {
			if job.Status.Active > 0 {
				status = append(status, false)
			} else if job.Status.Succeeded == job.Spec.Completions {
				status = append(status, true)
			}
		}
	}

	if len(status) == 0 {
		return false, nil
	}

	for _, s := range status {
		if !s {
			return false, nil
		}
	}

	return true, nil
}

// WaitOnReady is used when you dont have a handle on a job but want to wait until its in a Succeeded state.
func WaitOnReady(jobPrefix, namespace string, sleep, duration time.Duration) (bool, error) {
	readyCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Jobs (%s) to complete in namespace (%s)", duration.String(), jobPrefix, namespace)
			default:
				ready, _ := AreAllJobsCompleted(jobPrefix, namespace)
				if ready {
					readyCh <- true
				} else {
					time.Sleep(sleep)
				}
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			pods, getPodsErr := pod.GetAllByPrefix(jobPrefix, namespace)
			if getPodsErr != nil {
				log.Printf("Error trying to get job pods: %s\n", getPodsErr)
			}
			for _, p := range pods {
				p.Logs()
			}
			return false, err
		case ready := <-readyCh:
			return ready, nil
		}
	}
}

// WaitOnReady will call the static method WaitOnReady passing in p.Metadata.Name and p.Metadata.Namespace
func (j *Job) WaitOnReady(sleep, duration time.Duration) (bool, error) {
	return WaitOnReady(j.Metadata.Name, j.Metadata.Namespace, sleep, duration)
}

// Delete will delete a Job in a given namespace
func (j *Job) Delete(retries int) error {
	var kubectlOutput []byte
	var kubectlError error
	for i := 0; i < retries; i++ {
		cmd := exec.Command("k", "delete", "job", "-n", j.Metadata.Namespace, j.Metadata.Name)
		util.PrintCommand(cmd)
		kubectlOutput, kubectlError = cmd.CombinedOutput()
		if kubectlError != nil {
			log.Printf("Error while trying to delete Job %s in namespace %s:%s\n", j.Metadata.Namespace, j.Metadata.Name, string(kubectlOutput))
			continue
		}
		break
	}

	return kubectlError
}

// WaitOnDeleted returns when a job is successfully deleted
func WaitOnDeleted(jobPrefix, namespace string, sleep, duration time.Duration) (bool, error) {
	succeededCh := make(chan bool, 1)
	errCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				errCh <- errors.Errorf("Timeout exceeded (%s) while waiting for Jobs (%s) to be deleted in namespace (%s)", duration.String(), jobPrefix, namespace)
			default:
				p, err := GetAllByPrefix(jobPrefix, namespace)
				if err != nil {
					errCh <- errors.Errorf("Got error while getting Jobs with prefix \"%s\" in namespace \"%s\"", jobPrefix, namespace)
				}
				if len(p) == 0 {
					succeededCh <- true
				}
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return false, err
		case deleted := <-succeededCh:
			return deleted, nil
		}
	}
}
