package client


import (
"fmt"
"io/ioutil"
"net/http"
"encoding/json"
"bytes"
"ZCache/types"
	"ZCache/global"
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
func Get(ipAddr string, key string)(response * http.Response,err error){
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s/%s",ipAddr,global.Config.Port,key)
	request, _ := http.NewRequest(types.HttpGet, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func Delete(ipAddr string, key string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/%s",ipAddr,global.Config.Port,key)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpDelete, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func Set(ipAddr string, key string, value string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/%s/%s",ipAddr,global.Config.Port,key,value)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPOST, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}
func Update(ipAddr string, key string, value string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/%s/%s",ipAddr,global.Config.Port,key,value)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)

	return
}
func GetAll(ipAddr string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/keys",ipAddr,global.Config.Port)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpGet, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)

	return
}

func DeleteAll(ipAddr string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/keys",ipAddr,global.Config.Port)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpDelete, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)

	return
}

func Export(ipAddr string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/export",ipAddr,global.Config.Port)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpGet, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func Import(ipAddr string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/import",ipAddr,global.Config.Port)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func Expension(ipAddr string,size string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/expension/%s",ipAddr,global.Config.Port,size)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}


func GetKeysNum(ipAddr string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/keys_num",ipAddr,global.Config.Port)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpGet, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}


func Incr(ipAddr string,key string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/incr/:%s",ipAddr,global.Config.Port,key)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func IncrBy(ipAddr string,key string,value string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/incrBy/%s/%s",ipAddr,global.Config.Port,key,value)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func Decr(ipAddr string,key string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/decr/%s",ipAddr,global.Config.Port,key)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func DecrBy(ipAddr string,key string,value string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/decrBy/%s/%s",ipAddr,global.Config.Port,key,value)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func ImportRedis(ipAddr string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/import_Redis",ipAddr,global.Config.Port)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpPut, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}

func ExportRedis(ipAddr string)(response * http.Response,err error){
	url := fmt.Sprintf("http://%s:%s/export_Redis",ipAddr,global.Config.Port)
	req := `{}`
	req_byte := bytes.NewBuffer([]byte(req))
	client := &http.Client{}
	request, _ := http.NewRequest(types.HttpGet, url, req_byte)
	request.Header.Set("Content-type", "application/json")
	response, err = client.Do(request)
	return
}
