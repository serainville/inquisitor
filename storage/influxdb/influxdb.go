package influxdb

import (
	"fmt"
	"os"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/serainville/gologger"
	"github.com/serainville/inquisitor/models"
	"strconv"
)

var consoleLog = gologger.GetLogger(gologger.BASIC, gologger.ColoredLog)

func Client(data *models.ClientMetrics) {
	// NOTE: this assumes you've setup a user and have setup shell env variables,
	// namely INFLUX_USER/INFLUX_PWD. If not just omit Username/Password below.
	_, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://192.168.1.214:8086",
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	})
	if err != nil {
		//fmt.Println("Error creating InfluxDB Client: ", err.Error())
		consoleLog.Error("Error creating InfluxDB Client: " + err.Error())
	} else {
		consoleLog.Info("Connected to InfluxDB")
	}

}

// WriteMetrics Write a point using the HTTP client
func WriteMetrics(data models.ClientMetrics) error {

	// Create a point and add to batch
	tags := map[string]string{"accountid": data.AccountID, "hostname": data.Hostname}

	for _, mGroup := range data.Groups {
		fields := map[string]interface{}{}
		for _, mMetric := range mGroup.Metrics {
			intValue, _ := strconv.Atoi(mMetric.Value)
			fields[mMetric.Name] = intValue
		}

		err := storeMetrics(mGroup.Name, tags, fields)
		if err != nil {
			return err
		}
	}
	return nil
}

func storeMetrics(pointName string, tags map[string]string, fields map[string]interface{}) error {
	// Make client
	c, dberr := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://192.168.1.214:8086",
	})
	if dberr != nil {
		fmt.Println("Error creating InfluxDB Client: ", dberr.Error())
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "inquisitor",
		Precision: "s",
	})
	if err != nil {
		fmt.Println("Could not create point batch")
	}

	pt, err := client.NewPoint(pointName, tags, fields, time.Now())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	err = c.Write(bp)
	if err != nil {
		fmt.Println("Error: Could not write!", err.Error())
	}
	return err
}
