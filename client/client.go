package client

import (
	"ZCache/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	resp, _ := http.Get("http://10.67.2.252:8080/?a=123456&b=aaa&b=bbb")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var user User
	user.Name = "aaa"
	user.Age = 99
	if bs, err := json.Marshal(user); err == nil {
		//        fmt.Println(string(bs))
		req := bytes.NewBuffer([]byte(bs))
		tmp := `{"name":"junneyang", "age": 88}`
		req = bytes.NewBuffer([]byte(tmp))

		body_type := "application/json;charset=utf-8"
		resp, _ = http.Post("http://10.67.2.252:8080/test/", body_type, req)
		body, _ = ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println(err)
	}

	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://10.67.2.252:8080/?a=123456&b=aaa&b=bbb", nil)
	request.Header.Set("Connection", "keep-alive")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}

	req := `{"name":"junneyang", "age": 88}`
	req_new := bytes.NewBuffer([]byte(req))
	request, _ = http.NewRequest("POST", "http://10.67.2.252:8080/test/", req_new)
	request.Header.Set("Content-type", "application/json")
	response, _ = client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}
func Get(ipAddrPort string, key string) (response *http.Response, err error) {
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	url := fmt.Sprintf("http://%s/%s", ipAddrPort, key)
	request, _ := http.NewRequest(types.HttpGet, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func GetAll(ipAddrPort string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/keys", ipAddrPort)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpGet, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)

	return
}

func Export(ipAddrPort string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/export", ipAddrPort)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpGet, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func Import(ipAddrPort string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/import", ipAddrPort)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func Expension(ipAddrPort string, size string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/expension/%s", ipAddrPort, size)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func GetKeysNum(ipAddrPort string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/keys_num", ipAddrPort)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpGet, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func ImportRedis(ipAddrPort string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/import_Redis", ipAddrPort)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func ExportRedis(ipAddrPort string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/export_Redis", ipAddrPort)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpGet, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func GetDeleteAll(ipAddrPort string, ackChan chan int64) (err error) {
	response, err := DeleteAll(ipAddrPort)
	if err != nil {
		return
	}
	var resp types.ResponseAckData
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		newErr := json.Unmarshal(body, &resp)
		if newErr != nil {
			fmt.Println(newErr)
		}
		if "Success" == resp.Status {
			ackChan <- resp.CommitID
			return
		}
	}
	ackChan <- -1
	return
}
func DeleteAll(ipAddrPort string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/internal/keys", ipAddrPort)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpDelete, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)

	return
}

func GetDeleteAck(ipAddrPort string, key string, ackChan chan int64) (err error) {
	response, err := Delete(ipAddrPort, key)
	if err != nil {
		return
	}
	var resp types.ResponseAckData
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		newErr := json.Unmarshal(body, &resp)
		if newErr != nil {
			fmt.Println(newErr)
		}
		if "Success" == resp.Status {
			ackChan <- resp.CommitID
			return
		}
	}
	ackChan <- -1
	return
}

func Delete(ipAddrPort string, key string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/internal/%s", ipAddrPort, key)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpDelete, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func GetSetAck(ipAddrPort string, key string, value string, ackChan chan int64) (err error) {
	response, err := Set(ipAddrPort, key, value)
	if err != nil {
		return
	}
	var resp types.ResponseAckData
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		newErr := json.Unmarshal(body, &resp)
		if newErr != nil {
			fmt.Println(newErr)
		}
		if "Success" == resp.Status {
			ackChan <- resp.CommitID
			return
		}
	}
	ackChan <- -1
	return
}

func Set(ipAddrPort string, key string, value string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/internal/%s/%s", ipAddrPort, key, value)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPOST, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func GetUpdateAck(ipAddrPort string, key string, value string, ackChan chan int64) (err error) {
	response, err := Update(ipAddrPort, key, value)
	if err != nil {
		return
	}
	var resp types.ResponseAckData
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		newErr := json.Unmarshal(body, &resp)
		if newErr != nil {
			fmt.Println(newErr)
		}
		if "Success" == resp.Status {
			ackChan <- resp.CommitID
			return
		}
	}
	ackChan <- -1
	return
}

func Update(ipAddrPort string, key string, value string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/internal/%s/%s", ipAddrPort, key, value)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)

	return
}

func GetIncrAck(ipAddrPort string, key string, ackChan chan int64) (err error) {
	response, err := Incr(ipAddrPort, key)
	if err != nil {
		return
	}
	var resp types.ResponseAckData
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		newErr := json.Unmarshal(body, &resp)
		if newErr != nil {
			fmt.Println(newErr)
		}
		if "Success" == resp.Status {
			ackChan <- resp.CommitID
			return
		}
	}
	ackChan <- -1
	return
}

func Incr(ipAddrPort string, key string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/internal/incr/:%s", ipAddrPort, key)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func GetIncrByAck(ipAddrPort string, key string, value string, ackChan chan int64) (err error) {
	response, err := IncrBy(ipAddrPort, key, value)
	if err != nil {
		return
	}
	var resp types.ResponseAckData
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		newErr := json.Unmarshal(body, &resp)
		if newErr != nil {
			fmt.Println(newErr)
		}
		if "Success" == resp.Status {
			ackChan <- resp.CommitID
			return
		}
	}
	ackChan <- -1
	return
}

func IncrBy(ipAddrPort string, key string, value string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/internal/incrBy/%s/%s", ipAddrPort, key, value)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func GetDecrAck(ipAddrPort string, key string, ackChan chan int64) (err error) {
	response, err := Decr(ipAddrPort, key)
	if err != nil {
		return
	}
	var resp types.ResponseAckData
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		newErr := json.Unmarshal(body, &resp)
		if newErr != nil {
			fmt.Println(newErr)
		}
		if "Success" == resp.Status {
			ackChan <- resp.CommitID
			return
		}
	}
	ackChan <- -1
	return
}

func Decr(ipAddrPort string, key string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/internal/decr/%s", ipAddrPort, key)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func GetDecrByAck(ipAddrPort string, key string, value string, ackChan chan int64) (err error) {
	response, err := DecrBy(ipAddrPort, key, value)
	if err != nil {
		return
	}
	var resp types.ResponseAckData
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		newErr := json.Unmarshal(body, &resp)
		if newErr != nil {
			fmt.Println(newErr)
		}
		if "Success" == resp.Status {
			ackChan <- resp.CommitID
			return
		}
	}
	ackChan <- -1
	return
}

func DecrBy(ipAddrPort string, key string, value string) (response *http.Response, err error) {
	url := fmt.Sprintf("http://%s/internal/decrBy/%s/%s", ipAddrPort, key, value)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func CommitJob(ipAddrPort string, commitID int64) (result string, err error) {
	url := fmt.Sprintf("http://%s/internal/commit/%s", ipAddrPort, commitID)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	var resp types.ResponseAckData
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		newErr := json.Unmarshal(body, &resp)
		if newErr != nil {
			fmt.Println(newErr)
		}
		if "Success" == resp.Status {
			return "Success", nil
		}
	}
	return "Fail", nil
}

func DropJob(ipAddrPort string, commitID int64) (result string, err error) {
	url := fmt.Sprintf("http://%s/internal/drop/%s", ipAddrPort, commitID)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	var resp types.ResponseAckData
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		newErr := json.Unmarshal(body, &resp)
		if newErr != nil {
			fmt.Println(newErr)
		}
		if "Success" == resp.Status {
			return "Success", nil
		}
	}
	return "Fail", nil
}
