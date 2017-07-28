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

// Write a point using the HTTP client
func WriteMetrics(data models.ClientMetrics) {

	fmt.Println(data.ClientID, data.Timestamp, data.Secret, data.Metrics)

	// Create a point and add to batch
	tags := map[string]string{"client": "server01", "clientid": "1010101"}

	fields := map[string]interface{}{}
	for _, elem := range data.Metrics {
		//fmt.Println(elem.Name, elem.Group, elem.Value)

		intValue, _ := strconv.Atoi(elem.Value)
		fields[elem.Group+"_"+elem.Name] = intValue
	}

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

	pt, err := client.NewPoint("server01", tags, fields, time.Now())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	err = c.Write(bp)
	if err != nil {
		fmt.Println("Error: Could not write!", err.Error())
	}
}

