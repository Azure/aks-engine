//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package metrics

import (
	"log"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
	"github.com/kelseyhightower/envconfig"
)

// Config holds information on how to talk to influxdb
type Config struct {
	Address    string `envconfig:"INFLUX_ADDRESS" required:"true"`
	Username   string `envconfig:"INFLUX_USERNAME" required:"true"`
	Password   string `envconfig:"INFLUX_PASSWORD" required:"true"`
	Database   string `envconfig:"INFLUX_DATABASE" required:"true"`
	IsJenkins  bool   `envconfig:"IS_JENKINS"`
	JenkinsEnv *JenkinsCIEnvironment
}

// JenkinsCIEnvironment holds information about a test run within jenkins
type JenkinsCIEnvironment struct {
	Branch      string `envconfig:"GIT_BRANCH"`
	BuildNumber string `envconfig:"BUILD_NUMBER"`
	CommitSha   string `envconfig:"GIT_COMMIT"`
	Job         string `envconfig:"JOB_NAME"`
}

// Point holds data that will be written to influx
type Point struct {
	ProvisionStart      time.Time
	NodeWaitStart       time.Time
	TestStart           time.Time
	OverallStart        time.Time
	ProvisionDuration   time.Duration
	NodeWaitDuration    time.Duration
	TestDuration        time.Duration
	OverallDuration     time.Duration
	TestErrorCount      float64
	ProvisionErrorCount float64
	NodeWaitErrorCount  float64
	Tags                map[string]string
}

// ParseConfig will parse needed environment variables for running the tests
func ParseConfig() (*Config, error) {
	c := new(Config)
	if err := envconfig.Process("config", c); err != nil {
		return nil, err
	}

	if c.IsJenkins {
		jenkinsci := new(JenkinsCIEnvironment)
		if err := envconfig.Process("jenkinsci-config", jenkinsci); err != nil {
			return nil, err
		}
		c.JenkinsEnv = jenkinsci
	}
	return c, nil
}

// BuildPoint scaffolds a point object that stores information before being written to influx
func BuildPoint(orchestrator, location, clusterDefinition, subscriptionID string) *Point {
	p := Point{
		OverallStart:        time.Now(),
		ProvisionDuration:   0 * time.Second,
		NodeWaitDuration:    0 * time.Second,
		TestDuration:        0 * time.Second,
		OverallDuration:     0 * time.Second,
		ProvisionErrorCount: 0,
		TestErrorCount:      0,
		NodeWaitErrorCount:  0,
		Tags: map[string]string{
			"orchestrator":    orchestrator,
			"location":        location,
			"definition":      clusterDefinition,
			"subscription_id": subscriptionID,
		},
	}
	return &p
}

// SetTestStart will set TestStart value to time.Now()
func (p *Point) SetTestStart() {
	p.TestStart = time.Now()
}

// SetProvisionStart will set ProvisionStart value to time.Now()
func (p *Point) SetProvisionStart() {
	p.ProvisionStart = time.Now()
}

// RecordProvisionError sets appropriate values for when a test error occurs
func (p *Point) RecordProvisionError() {
	p.ProvisionDuration = time.Since(p.ProvisionStart)
	p.ProvisionErrorCount++
}

// RecordProvisionSuccess sets TestErrorCount to 0 to mark tests succeeded
func (p *Point) RecordProvisionSuccess() {
	p.ProvisionDuration = time.Since(p.ProvisionStart)
}

// SetNodeWaitStart will set NodeWaitStart value to time.Now()
func (p *Point) SetNodeWaitStart() {
	p.NodeWaitStart = time.Now()
}

// RecordNodeWait will set NodeWaitDuration to time.Since(p.NodeWaitStart)
func (p *Point) RecordNodeWait(err error) {
	if err != nil {
		p.NodeWaitErrorCount++
	}
	p.NodeWaitDuration = time.Since(p.NodeWaitStart)
}

// RecordTestError sets appropriate values for when a test error occurs
func (p *Point) RecordTestError() {
	p.TestDuration = time.Since(p.TestStart)
	p.TestErrorCount++
}

// RecordTestSuccess sets TestErrorCount to 0 to mark tests succeeded
func (p *Point) RecordTestSuccess() {
	p.TestDuration = time.Since(p.TestStart)
}

// RecordTotalTime captures total runtime of tests
func (p *Point) RecordTotalTime() {
	p.OverallDuration = time.Since(p.OverallStart)
}

// SetProvisionMetrics will parse the csv data retrieved from /opt/m and set appropriate fields
func (p *Point) SetProvisionMetrics(data []byte) {

}

func (p *Point) Write() {
	cfg, err := ParseConfig()
	if err == nil {
		log.Printf("Jenkins?:%v\n", cfg.IsJenkins)

		if cfg.IsJenkins {
			p.Tags["branch"] = cfg.JenkinsEnv.Branch
			p.Tags["commit-sha"] = cfg.JenkinsEnv.CommitSha
			p.Tags["build_number"] = cfg.JenkinsEnv.BuildNumber
			p.Tags["job"] = cfg.JenkinsEnv.Job
			p.Tags["ci"] = "jenkins"
		}

		fields := map[string]interface{}{
			"provision-secs":        p.ProvisionDuration.Seconds(),
			"node-wait-secs":        p.NodeWaitDuration.Seconds(),
			"test-secs":             p.TestDuration.Seconds(),
			"total-secs":            p.OverallDuration.Seconds(),
			"test-error-count":      p.TestErrorCount,
			"provision-error-count": p.ProvisionErrorCount,
			"node-wait-error-count": p.NodeWaitErrorCount,
		}

		c, err := client.NewHTTPClient(client.HTTPConfig{
			Addr:     cfg.Address,
			Username: cfg.Username,
			Password: cfg.Password,
		})
		if err != nil {
			log.Printf("Error trying to create influx http client:%s\n", err)
		}

		// Create a new point batch
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  cfg.Database,
			Precision: "s",
		})
		if err != nil {
			log.Printf("Error trying to create batch points:%s\n", err)
		}

		pt, err := client.NewPoint("test", p.Tags, fields, time.Now())
		log.Printf("Point:%+v\n", pt)
		if err != nil {
			log.Printf("Error trying to create metric point:%s\n", err)
		}
		bp.AddPoint(pt)
		c.Write(bp)
	}
}
