package promote

import (
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
)

const (
	// AKSRPTests are the rp tests
	AKSRPTests string = "AKSRPTests"
	// AKSEngineTests are the AKS Engine tests
	AKSEngineTests string = "AKSEngineTests"
	// SimDemTests are the SimDem tests
	SimDemTests string = "SimDemTests"
)

// StorageAccount is how we connect to storage
type StorageAccount struct {
	Name string
	Key  string
}

const (
	// Kubernetes is the string constant for the Kubernetes orchestrator type
	Kubernetes string = "Kubernetes"
	// RecordTestRunTableName for storing RecordTestRun
	RecordTestRunTableName string = "RecordTestRun"
	// PromoteToFailureTableName for storing PromoteToFailure
	PromoteToFailureTableName string = "PromoteToFailure"
)

// TestRunQos structure
type TestRunQos struct {
	TimeStampUTC   time.Time // partition key
	TestName       string    // row key
	TestType       string
	SubscriptionID string
	ResourceGroup  string
	Region         string
	Orchestrator   string
	Success        bool
	FailureStr     string
	// DeploymentDurationInSeconds int
}

//DigitalSignalFilter structure
type DigitalSignalFilter struct {
	TestName     string // partition key
	TestType     string // row key
	FailureStr   string
	FailureCount float64
}

// RecordTestRun procedure pushes all test result data to RecordTestRun Table
func RecordTestRun(sa StorageAccount, testRunQos TestRunQos) {
	// fmt.Printf("record test run Qos to '%s': %v\n", sa.Name, testRunQos)

	// get Azure Storage Client
	var err error
	var azureStoreClient storage.Client
	if azureStoreClient, err = storage.NewBasicClient(sa.Name, sa.Key); err != nil {
		fmt.Printf("FAIL to create azure storage basic client, Error: %s\n", err.Error())
		return
	}

	// From storageClient, get Table Service Client
	tsc := azureStoreClient.GetTableService()
	table1 := tsc.GetTableReference(RecordTestRunTableName)

	// Create Table if it does not exist
	if err := table1.Create(30, storage.FullMetadata, nil); err != nil && !strings.Contains(err.Error(), "The table specified already exists") {
		fmt.Printf("Failed to create table: %s, Error: %s\n", RecordTestRunTableName, err.Error())
		return
	}
	// fmt.Printf("Table : %s is created\n", RecordTestRunTableName)

	t := testRunQos.TimeStampUTC.Format("2006-01-02 15:04:05")

	// Insert Entity Entry into Table
	entity := table1.GetEntityReference(t, testRunQos.TestName)

	props := map[string]interface{}{
		"TestType":       testRunQos.TestType,
		"SubscriptionID": testRunQos.SubscriptionID,
		"ResourceGroup":  testRunQos.ResourceGroup,
		"Region":         testRunQos.Region,
		"Orchestrator":   testRunQos.Orchestrator,
		"Success":        testRunQos.Success,
		"FailureStr":     testRunQos.FailureStr,
	}

	entity.Properties = props
	if err = entity.Insert(storage.FullMetadata, nil); err != nil {
		fmt.Printf("Could not insert entity into table, Error: %v\n", err)
		return
	}
}

// RunPromoteToFailure procedure
// Returns True when Error is Promoted, Else False
func RunPromoteToFailure(sa StorageAccount, testRunPromToFail DigitalSignalFilter) (bool, error) {

	// get Azure Storage Client
	var err error
	var azureStoreClient storage.Client
	if azureStoreClient, err = storage.NewBasicClient(sa.Name, sa.Key); err != nil {
		fmt.Printf("FAIL to create azure storage basic client, Error: %s\n", err.Error())
		return false, err
	}

	// From azureStoreClient, get Table Service Client
	tsc := azureStoreClient.GetTableService()
	table1 := tsc.GetTableReference(PromoteToFailureTableName)

	// Create Table if it does not exist
	if err := table1.Create(30, storage.FullMetadata, nil); err != nil && !strings.Contains(err.Error(), "The table specified already exists") {
		fmt.Printf("Failed to create table: %s, Error: %s\n", PromoteToFailureTableName, err.Error())
		return false, err
	}
	// 1. Get the Entity using partition key and row key
	// 2. If doesnt exist, then create the new entity and exit as success
	// 3. If it exists, then increment the FailureCount
	// 4. If FailureCount == 3, push out faillure

	entity := table1.GetEntityReference(testRunPromToFail.TestName, testRunPromToFail.TestType)

	err = entity.Get(30, storage.FullMetadata, &storage.GetEntityOptions{
		Select: []string{"FailureStr", "FailureCount"},
	})

	if err != nil {
		if strings.Contains(err.Error(), "The specified resource does not exist") {
			// Entity does not exist in Table
			// Insert Entity into Table

			err = insertEntity(table1, testRunPromToFail)
			if err != nil {
				fmt.Printf("Error inserting entity :  %v\n", err)
				return false, err
			}
		}
		return false, err
	}

	existingFailureStr := entity.Properties["FailureStr"]
	existingFailureCount := entity.Properties["FailureCount"]

	if existingFailureStr != testRunPromToFail.FailureStr {
		// Perform Update of this entity with testRunPromToFail.FailureStr and testRunPromToFail.FailureCount
		err = updateEntity(entity, testRunPromToFail.FailureCount, testRunPromToFail.FailureStr)
		return false, err
	}

	if testRunPromToFail.FailureCount == 0 {
		// Update the Entity with FailureCount 0
		// Return False
		if err = updateEntity(entity, testRunPromToFail.FailureCount, testRunPromToFail.FailureStr); err != nil {
			fmt.Printf("Failed to reset Failure Count for %s to : %v!\n\n", testRunPromToFail.TestName, testRunPromToFail.FailureCount)
			return false, err
		}
		fmt.Printf("Reset Failure Count for %s to : %v\n\n", testRunPromToFail.TestName, testRunPromToFail.FailureCount)
		return false, nil
	}

	fmt.Printf("Existing Failure Count for %s : %v\n\n", testRunPromToFail.TestName, existingFailureCount)
	newFailureCount := existingFailureCount.(float64) + testRunPromToFail.FailureCount
	fmt.Printf("Incremented Failure Count for %s to : %v\n\n", testRunPromToFail.TestName, newFailureCount)
	if err = updateEntity(entity, newFailureCount, testRunPromToFail.FailureStr); err != nil {
		return false, err
	}

	if newFailureCount >= 3 {
		return true, nil
	}

	return false, nil
}

func insertEntity(table *storage.Table, testRunPromToFail DigitalSignalFilter) error {
	// Insert Entity Entry into Table
	entity := table.GetEntityReference(testRunPromToFail.TestName, testRunPromToFail.TestType)

	props := map[string]interface{}{
		"FailureStr":   testRunPromToFail.FailureStr,
		"FailureCount": testRunPromToFail.FailureCount,
	}

	entity.Properties = props
	if err := entity.Insert(storage.FullMetadata, nil); err != nil {
		fmt.Printf("Could not insert entity into table, Error: %v\n", err)
		return err
	}
	return nil
}

func updateEntity(entity *storage.Entity, failureCount float64, failureStr string) error {

	props := map[string]interface{}{
		"FailureStr":   failureStr,
		"FailureCount": failureCount,
	}
	entity.Properties = props
	if err := entity.Update(false, nil); err != nil {
		fmt.Printf("Error in Updating Entity : %v\n\n", err)
	}
	return nil
}
