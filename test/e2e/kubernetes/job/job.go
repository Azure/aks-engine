//+build test
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
func CreateJobFromFile(filename, name, namespace string, sleep, timeout time.Duration) (*Job, error) {
	cmd := exec.Command("k", "create", "-f", filename)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to create Job %s:%s\n", name, string(out))
		return nil, err
	}
	job, err := GetWithRetry(name, namespace, sleep, timeout)
	if err != nil {
		log.Printf("Error while trying to fetch Job %s:%s\n", name, err)
		return nil, err
	}
	return job, nil
}

// CreateWindowsJobFromTemplate will create a Job from file with a name
func CreateWindowsJobFromTemplate(filename, name, namespace string, windowsTestImages *engine.WindowsTestImages, sleep, timeout time.Duration) (*Job, error) {
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

	return CreateJobFromFile(tempfile.Name(), name, namespace, sleep, timeout)
}

// CreateWindowsJobFromTemplateDeleteIfExists will create a Job from file, deleting any pre-existing job with the same name
func CreateWindowsJobFromTemplateDeleteIfExists(filename, name, namespace string, windowsTestImages *engine.WindowsTestImages, sleep, timeout time.Duration) (*Job, error) {
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
	return CreateWindowsJobFromTemplate(filename, name, namespace, windowsTestImages, sleep, timeout)
}

// CreateJobFromFileDeleteIfExists will create a Job from file, deleting any pre-existing job with the same name
func CreateJobFromFileDeleteIfExists(filename, name, namespace string, sleep, timeout time.Duration) (*Job, error) {
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
	return CreateJobFromFile(filename, name, namespace, sleep, timeout)
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

// GetAllByPrefixResult is a return struct for GetAllByPrefixAsync
type GetAllByPrefixResult struct {
	jobs []Job
	err  error
}

// GetAllByPrefixAsync wraps GetAllByPrefix with a struct response for goroutine + channel usage
func GetAllByPrefixAsync(prefix, namespace string) GetAllByPrefixResult {
	jobs, err := GetAllByPrefix(prefix, namespace)
	return GetAllByPrefixResult{
		jobs: jobs,
		err:  err,
	}
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

// GetResult is a return struct for GetAsync
type GetResult struct {
	Job *Job
	Err error
}

// GetAsync wraps Get with a struct response for goroutine + channel usage
func GetAsync(jobName, namespace string) GetResult {
	job, err := Get(jobName, namespace)
	return GetResult{
		Job: job,
		Err: err,
	}
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

// GetWithRetry gets a job, allowing for retries
func GetWithRetry(jobName, namespace string, sleep, timeout time.Duration) (*Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetResult)
	var mostRecentGetWithRetryError error
	var job *Job
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetAsync(jobName, namespace)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetWithRetryError = result.Err
			job = result.Job
			if mostRecentGetWithRetryError == nil {
				if job != nil {
					return job, nil
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("GetWithRetry timed out: %s\n", mostRecentGetWithRetryError)
		}
	}
}

// AreAllJobsSucceededResult is a return struct for AreAllJobsSucceededAsync
type AreAllJobsSucceededResult struct {
	succeeded bool
	err       error
}

// AreAllJobsSucceededAsync wraps AreAllJobsSucceeded with a struct response for goroutine + channel usage
func AreAllJobsSucceededAsync(jobPrefix, namespace string) AreAllJobsSucceededResult {
	succeeded, err := AreAllJobsSucceeded(jobPrefix, namespace)
	return AreAllJobsSucceededResult{
		succeeded: succeeded,
		err:       err,
	}
}

// AreAllJobsSucceeded will return true if all jobs with a common prefix in a given namespace are in a Completed State
func AreAllJobsSucceeded(jobPrefix, namespace string) (bool, error) {
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

// WaitOnSucceeded returns true if all jobs matching a prefix substring are in a succeeded state within a period of time
func WaitOnSucceeded(jobPrefix, namespace string, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan AreAllJobsSucceededResult)
	var mostRecentAreAllJobsSucceededError error
	var succeeded bool
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- AreAllJobsSucceededAsync(jobPrefix, namespace)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentAreAllJobsSucceededError = result.err
			succeeded = result.succeeded
			if mostRecentAreAllJobsSucceededError == nil {
				if succeeded {
					return true, nil
				}
			}
		case <-ctx.Done():
			DescribeJobs(jobPrefix, namespace)
			return false, errors.Errorf("WaitOnSucceeded timed out: %s\n", mostRecentAreAllJobsSucceededError)
		}
	}
}

// WaitOnSucceeded will call the static method WaitOnSucceeded passing in p.Metadata.Name and p.Metadata.Namespace
func (j *Job) WaitOnSucceeded(sleep, timeout time.Duration) (bool, error) {
	return WaitOnSucceeded(j.Metadata.Name, j.Metadata.Namespace, sleep, timeout)
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

// DescribeJobs describes all jobs whose name matches a substring
func DescribeJobs(jobPrefix, namespace string) {
	jobs, err := GetAllByPrefix(jobPrefix, namespace)
	if err != nil {
		log.Printf("Unable to get jobs matching prefix %s in namespace %s: %s", jobPrefix, namespace, err)
	}
	for _, j := range jobs {
		err := j.Describe()
		if err != nil {
			log.Printf("Unable to describe job %s: %s", j.Metadata.Name, err)
		}
	}
}

// Describe will describe a Job resource
func (j *Job) Describe() error {
	var commandTimeout time.Duration
	cmd := exec.Command("k", "describe", "jobs/", j.Metadata.Name, "-n", j.Metadata.Namespace)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	log.Printf("\n%s\n", string(out))
	return err
}

// WaitOnDeleted returns when all jobs matching a prefix substring are successfully deleted
func WaitOnDeleted(jobPrefix, namespace string, sleep, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetAllByPrefixResult)
	var mostRecentWaitOnDeletedError error
	var jobs []Job
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetAllByPrefixAsync(jobPrefix, namespace)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentWaitOnDeletedError = result.err
			jobs = result.jobs
			if mostRecentWaitOnDeletedError == nil {
				if len(jobs) == 0 {
					return true, nil
				}
			}
		case <-ctx.Done():
			for _, j := range jobs {
				j.Describe()
			}
			return false, errors.Errorf("WaitOnDeleted timed out: %s\n", mostRecentWaitOnDeletedError)
		}
	}
}
