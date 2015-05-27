package rest

import (
  "net/http"
  "io/ioutil"
)
/**
 * Make HTTP GET request to specified REST API Endpoint
 * @param {string} url) ([]byte, error URL for API Endpoint
 * @return {[]byte} Byte array of response
 */
func GetContent(url string) ([]byte, error) {
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }

  //send request
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  //defer closing
  defer resp.Body.Close()

  //read content
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  return body, nil
}