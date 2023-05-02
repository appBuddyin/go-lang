package main

import (
	"github.com/joho/godotenv"
	htp "md-core/http"
	"md-core/utils"
	"reflect"
	"testing"
	"time"
)

var testLog = utils.CreateLog("TestRestTemplateInstance")

func test_simple_func(t *testing.T){
	type res struct{
		OtpId          string      `json:"otpId,omitempty"`
		Otp            string       `json:"otp,omitempty"`
		Expiry         time.Time    `json:"expiry,omitempty"`
		Active         bool         `json:"active,omitempty"`
		Timestamp      time.Time    `json:"timestamp,omitempty"`
	}{
		OtpId: ""
	}
}

func Test_restTemplate_Request(t *testing.T) {

	type args struct {
		serviceName string
		method  string
		url     string
		headers map[string][]string
		body    string
		object  interface{}
	}
	type Otp struct {
		OtpId          string      `json:"otpId,omitempty"`
		Otp            string       `json:"otp,omitempty"`
		Expiry         time.Time    `json:"expiry,omitempty"`
		Active         bool         `json:"active,omitempty"`
		Timestamp      time.Time    `json:"timestamp,omitempty"`
	}
	type  TokenResponse struct {
		MessageId  *      string `json:"messageId,omitempty"`
		OtpDetails Otp  `json:"otpDetails,omitempty"`
	}

	expectedResponseObj:=&htp.ResponseEntity{
		Header: htp.ResponseHeader{Message: htp.GetMsg(htp.OK), StatusCode: htp.OK}}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{"Get OTP",
		 args{
				serviceName : "md-auth",
				method:  "POST",
				url:     "/api/auth/v1/otp",
				headers: map[string][]string{
					"Content-Type": {"application/json","plain/text"},
				},
				body:    `{"phoneNumber":"+918800269004","reference":"1634818004086","skip":true}`,
				object:   &htp.ResponseEntity{},
			},
			expectedResponseObj,
		},
		{"Get Profile",
			args{
				serviceName : "md-auth",
				method:  "GET",
				url:     "/api/auth/v1/profile?vendorId=c287e8ae-efef-4d21-a8b7-831771184730&vendorType=DOCTOR",
				headers: map[string][]string{
					"Authorization": {"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImQyZDMwNWIxNTRhYTBlZDZlYjllMzljMDQ5ZmI2M2FmZjM4ODRkNDciLCJ1dWlkIjoiZGQ5OGY4MDctNGUzMS00MDZkLWJmMGItN2YxMjliODQ0YWVhIiwidHlwZSI6IkRPQ1RPUiIsImFyZ3MiOm51bGwsImV4cCI6MTYzNTI0MzkzNH0.yTa3Gvgwmc-gHZKC8patdqufJE007wdazbgtzGVFMaQ"},
					"Content-Type": {"application/json","plain/text"},
				},
				//body:    `{"phoneNumber":"+918800269004","reference":"1634818004086","skip":true}`,
				object:   &htp.ResponseEntity{},
			},
			expectedResponseObj,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := godotenv.Load("/Users/naren/projects/medi/md-auth/.env")
			if err != nil {
				t.Error("Error loading .env file",err)
			}
			s := &restTemplate{}
			got,error :=s.Request(tt.args.serviceName,tt.args.method, tt.args.url, tt.args.headers, tt.args.body, tt.args.object)
			if error!=nil {
				t.Errorf("Internal Error %v", error)
			}
			response:=*(got).(*htp.ResponseEntity)

			if !(response.Header.StatusCode==htp.OK) {
				t.Errorf("StatusCode is %v", response.Header.StatusCode)
			}
		})
	}
}
