package report

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/Azure/aks-engine/pkg/helpers"
)

// ErrorInfo represents the CI error
type ErrorInfo struct {
	TestName string
	Step     string
	ErrName  string
	ErrClass string
	Location string
}

// ErrorStat represents the aggregate error count and region
type ErrorStat struct {
	Count     int            `json:"count"`
	Locations map[string]int `json:"locations"`
}

// Manager represents the details about a build and errors in that build
type Manager struct {
	lock        sync.Mutex
	JobName     string    `json:"job"`
	BuildNum    int       `json:"build"`
	Deployments int       `json:"deployments"`
	Errors      int       `json:"errors"`
	StartTime   time.Time `json:"startTime"`
	Duration    string    `json:"duration"`
	// Failure map: key=error, value=locations
	Failures  map[string]*ErrorStat `json:"failures"`
	LogErrors logErrors             `json:"-"`
}

type logErrors struct {
	LogErrors []logError `json:"Errors"`
}

type logError struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Regex string `json:"regex"`
}

const (
	// ErrClassDeployment represents an error during deployment
	ErrClassDeployment = "Deployment"
	// ErrClassValidation represents an error during validation (tests)
	ErrClassValidation = "Validation"
	// ErrClassAzcli represents an error with Azure CLI
	ErrClassAzcli = "AzCLI"
	// ErrClassNone represents absence of error
	ErrClassNone = "None"

	// ErrSuccess represents a success, for some reason
	ErrSuccess = "Success"
	// ErrUnknown represents an unknown error
	ErrUnknown = "UnspecifiedError"
)

// New creates a new error report
func New(jobName string, buildNum int, nDeploys int, logErrorsFileName string) *Manager {
	h := &Manager{}
	h.JobName = jobName
	h.BuildNum = buildNum
	h.Deployments = nDeploys
	h.Errors = 0
	h.StartTime = time.Now().UTC()
	h.Failures = make(map[string]*ErrorStat)
	h.LogErrors = makeErrorList(logErrorsFileName)
	return h
}

func makeErrorList(fileName string) logErrors {
	dummy := logErrors{}

	if fileName != "" {
		file, e := ioutil.ReadFile(fileName)
		if e != nil {
			// do not exit the tests
			fmt.Printf("ERROR: %v\n", e)
		}
		json.Unmarshal(file, &dummy)
	}
	return dummy
}

// Copy TBD needs definition [ToDo]
func (h *Manager) Copy() *Manager {
	n := New(h.JobName, h.BuildNum, h.Deployments, "")
	n.Errors = h.Errors
	n.StartTime = h.StartTime
	for e, f := range h.Failures {
		locs := make(map[string]int)
		for l, c := range f.Locations {
			locs[l] = c
		}
		n.Failures[e] = &ErrorStat{Count: f.Count, Locations: locs}
	}
	return n
}

// Process TBD needs definition
func (h *Manager) Process(txt, step, testName, location string) *ErrorInfo {
	for _, logErr := range h.LogErrors.LogErrors {
		if match, _ := regexp.MatchString(logErr.Regex, txt); match {
			h.addFailure(logErr.Name, map[string]int{location: 1})
			return NewErrorInfo(testName, step, logErr.Name, logErr.Class, location)
		}
	}
	h.addFailure(ErrUnknown, map[string]int{location: 1})
	return NewErrorInfo(testName, step, ErrUnknown, ErrClassNone, location)
}

func (h *Manager) addFailure(key string, locations map[string]int) {
	h.lock.Lock()
	defer h.lock.Unlock()

	cnt := 0
	if failure, ok := h.Failures[key]; !ok {
		locs := make(map[string]int)
		for l, c := range locations {
			locs[l] = c
			cnt += c
		}
		h.Failures[key] = &ErrorStat{Count: cnt, Locations: locs}
	} else {
		for l, c := range locations {
			cnt += c
			if _, ok := failure.Locations[l]; !ok {
				failure.Locations[l] = c
			} else {
				failure.Locations[l] += c
			}
		}
		failure.Count += cnt
	}
	h.Errors += cnt
}

// CreateTestReport TBD needs definition
func (h *Manager) CreateTestReport(filepath string) error {
	h.Duration = time.Now().UTC().Sub(h.StartTime).String()
	data, err := helpers.JSONMarshalIndent(h, "", "  ", false)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

// CreateCombinedReport TBD needs definition
func (h *Manager) CreateCombinedReport(filepath, testReportFname string) error {
	// "COMBINED_PAST_REPORTS" is the number of recent reports in the combined report
	reports, err := strconv.Atoi(os.Getenv("COMBINED_PAST_REPORTS"))
	if err != nil || reports <= 0 {
		return nil
	}
	combinedReport := h.Copy()
	for i := 1; i <= reports; i++ {
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/%d/%s/%s",
			os.Getenv("JOB_BUILD_ROOTDIR"), h.BuildNum-i, os.Getenv("JOB_BUILD_SUBDIR"), testReportFname))
		if err != nil {
			break
		}
		testReport := &Manager{}
		if err := json.Unmarshal(data, &testReport); err != nil {
			break
		}
		combinedReport.StartTime = testReport.StartTime
		combinedReport.Deployments += testReport.Deployments

		for e, f := range testReport.Failures {
			combinedReport.addFailure(e, f.Locations)
		}
	}
	return combinedReport.CreateTestReport(filepath)
}

// NewErrorInfo TBD needs definition
func NewErrorInfo(testName, step, errName, errClass, location string) *ErrorInfo {
	return &ErrorInfo{TestName: testName, Step: step, ErrName: errName, ErrClass: errClass, Location: location}
}
