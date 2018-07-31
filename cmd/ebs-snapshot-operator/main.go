package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ec2InstanceMetaData struct {
	PrivateIP          string `json:"privateIp"`
	InstanceID         string `json:"instanceId"`
	BillingProducts    string `json:"billingProducts"`
	InstanceType       string `json:"instanceType"`
	AccountID          int    `json:"accountId"`
	PendingTime        string `json:"pendingTime"`
	ImageID            string `json:"imageId"`
	KernelID           string `json:"kernelId"`
	RamdiskID          string `json:"ramdiskId"`
	Architecture       string `json:"architecture"`
	Region             string `json:"region"`
	Version            string `json:"version"`
	AvailabilityZone   string `json:"availabilityZone"`
	DevpayProductCodes string `json:"devpayProductCodes"`
}

func main() {
	// Get Region of Current Instance
	url := "http://169.254.169.254/latest/dynamic/instance-identity/document"
	// url := "http://localhost:8080/metadata"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	var metadata ec2InstanceMetaData

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		log.Println(err)
	}

	fmt.Println("Current Region = ", metadata.Region)

	//sess, err := session.NewSession(&aws.Config{Region: aws.String(metadata.region)})
}
