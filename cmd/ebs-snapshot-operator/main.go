import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ec2InstanceMetaData struct {
	privateIp			string `json:"privateIp"`
	instanceId			string `json:"instanceId"`
	billingProducts		string `json:"billingProducts"`
	instanceType		string `json:"instanceType"`
	accountId			int    `json:"accountId"`
	pendingTime			string `json:"pendingTime"`
	imageId				string `json:"imageId"`
	kernelId			string `json:"kernelId"`
	ramdiskId			string `json:"privateIp"`
	architecture		string `json:"architecture"`
	region				string `json:"region"`
	version				string `json:"version"`
	availabilityZone	string `json:"availabilityZone"`
	devpayProductCodes	string `json:"devpayProductCodes"`
}

// Get Region of Current Instance
url := "http://169.254.169.254/latest/dynamic/instance-identity/document"

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
var metadata ec2InstanceMetaData

// Use json.Decode for reading streams of JSON data
if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
	log.Println(err)
}

fmt.Println("Current Region = ", metadata.region)

sess, err := session.NewSession(&aws.Config{Region: aws.String(metadata.region)})