package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/bigtable"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func main() {
	// write data to bigtable
	btInstance := "localhost:8086"
	WriteSimple(os.Stdout, "test", btInstance, "mobile-time-series")

	// Retrieve credentials from environment variables
	endpoint := os.Getenv("TABLESTORE_ENDPOINT")
	instanceName := os.Getenv("TABLESTORE_INSTANCE")
	accessKeyID := os.Getenv("TABLESTORE_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("TABLESTORE_ACCESS_KEY_SECRET")

	if endpoint == "" || instanceName == "" || accessKeyID == "" || accessKeySecret == "" {
		fmt.Println("Error: Missing environment variables for Tablestore credentials.")
		return
	}
	// Initialize the Tablestore client
	client := tablestore.NewClient(endpoint, instanceName, accessKeyID, accessKeySecret)

	tableExists := checkTable(client, "SampleTable")

	if !tableExists {
		// Create the table
		createTable(client)
	}

	// Write data to the table
	writeData(client, "SampleTable", "Row1", "example value")
}

func checkTable(client *tablestore.TableStoreClient, tableName string) bool {
	exists := false
	tables, err := client.ListTable()
	if err != nil {
		fmt.Println("Failed to list tables:", err)
		return exists
	}

	for _, table := range tables.TableNames {
		if table == tableName {
			fmt.Println("Table already exists. Skipping table creation.")
			exists = true
			break
		}
	}

	return exists
}

// Function to create a table
func createTable(client *tablestore.TableStoreClient) {
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = "SampleTable"
	tableMeta.AddPrimaryKeyColumn("PK1", tablestore.PrimaryKeyType_STRING)

	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1 // Data will not expire
	tableOption.MaxVersion = 1   // Maximum number of versions for each cell

	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0

	createTableRequest := new(tablestore.CreateTableRequest)
	createTableRequest.TableMeta = tableMeta
	createTableRequest.TableOption = tableOption
	createTableRequest.ReservedThroughput = reservedThroughput

	_, err := client.CreateTable(createTableRequest)
	if err != nil {
		fmt.Println("Failed to create table:", err)
	} else {
		fmt.Println("Table created successfully.")
	}
}

// Function to write data to the table
func writeData(client *tablestore.TableStoreClient, tableName, primaryKeyValue, attributeValue string) {
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = tableName

	// Set primary key
	primaryKey := new(tablestore.PrimaryKey)
	fmt.Println("Primary key value:", primaryKeyValue, "Attribute value:", attributeValue)
	primaryKey.AddPrimaryKeyColumn("PK1", primaryKeyValue)
	putRowChange.PrimaryKey = primaryKey

	// Set attribute column
	putRowChange.AddColumn("Attribute1", attributeValue)
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)

	putRowRequest.PutRowChange = putRowChange

	_, err := client.PutRow(putRowRequest)
	if err != nil {
		fmt.Println("Failed to write data:", err)
	} else {
		fmt.Println("Data written successfully.")
	}
}

func WriteSimple(w io.Writer, projectID, instanceID string, tableName string) error {
	// projectID := "my-project-id"
	// instanceID := "my-instance-id"
	// tableName := "mobile-time-series"

	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, projectID, instanceID)
	if err != nil {
		return fmt.Errorf("bigtable.NewClient: %w", err)
	}
	defer client.Close()
	tbl := client.Open(tableName)
	columnFamilyName := "stats_summary"
	timestamp := bigtable.Now()

	mut := bigtable.NewMutation()
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, int64(1))

	mut.Set(columnFamilyName, "connected_cell", timestamp, buf.Bytes())
	mut.Set(columnFamilyName, "connected_wifi", timestamp, buf.Bytes())
	mut.Set(columnFamilyName, "os_build", timestamp, []byte("PQ2A.190405.003"))

	rowKey := "phone#4c410523#20190501"
	if err := tbl.Apply(ctx, rowKey, mut); err != nil {
		return fmt.Errorf("Apply: %w", err)
	}

	fmt.Fprintf(w, "Successfully wrote row: %s\n", rowKey)
	return nil
}
