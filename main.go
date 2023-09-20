package main

import (
	"fmt"
	"github.com/srun-soft/api-sdk/apisdk"
	"net/url"
)

func main() {
	sdk := &apisdk.ApiConfig{
		//AppID:     "srunsoft",
		//AppSecret: "ba62f23e6212790052f387ee7af2943e4cfcece0",
		Scheme:  "https",
		Host:    "192.168.0.191",
		Port:    8001,
		Version: 2,
	}
	sdk.Version = 1
	v := url.Values{}
	v.Add("user_name", "yuantong")
	sdk.Request.Params = v
	res := sdk.UserView()
	fmt.Println(res.Code)
	fmt.Println(res.Message)
	if res.Code == 0 {
		m := res.Data.(map[string]interface{})
		fmt.Printf("%v", m["user_name"])
	}

}