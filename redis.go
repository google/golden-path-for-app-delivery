// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RedisInstance struct {
	Status Status `json:"status"`
}
type Status struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

func getRedisInstance(redisName string) RedisInstance {
	// If we are not in dev, try to get the redis info from KCC
	token, err := getToken()
	namespace := getNamespace()
	if err != nil {
		panic(fmt.Errorf("unable to get token when constructing redis url: %v", err))
	}
	redisResourceURL := fmt.Sprintf("https://kubernetes/apis/redis.cnrm.cloud.google.com/v1beta1/namespaces/%s/redisinstances/%s", namespace, redisName)
	request, err := http.NewRequest("GET", redisResourceURL, nil)
	if err != nil {
		panic(fmt.Errorf("unable to create request when constructing redis url: %v", err))
	}
	authHeader := fmt.Sprintf("Bearer %s", token)
	request.Header.Set("Authorization", authHeader)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}
	resp, err := client.Do(request)
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("unable to read response querying redis information: %v", err))
	}
	redisInstance := RedisInstance{}
	err = json.Unmarshal(responseData, &redisInstance)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshall response from redis instance: %v", err))
	}
	return redisInstance
}

func getRedisURL() string {
	// Useful for when ConfigConnector is being used
	/*switch version{
	case "dev":
		return "redis-dev:6379"
	case "staging":
		redisInstance := getRedisInstance("redis-staging")
		redisUrl := fmt.Sprintf("%v:%v", redisInstance.Status.Host, redisInstance.Status.Port)
		return redisUrl
	case "canary", "prod":
		redisInstance := getRedisInstance("redis-prod")
		redisUrl := fmt.Sprintf("%v:%v", redisInstance.Status.Host, redisInstance.Status.Port)
		return redisUrl
	default:
		return "redis:6379"
	}*/
	return "redis:6379"
}

func getToken() (string, error) {
	// Fall back to the namespace associated with the service account token, if available
	if data, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token"); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			return ns, nil
		}
	} else {
		return "", err
	}
	return "", fmt.Errorf("unable to get token")
}
