package weatherstack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

const (
	APIURL         string = "https://api.weatherstack.com"
	DateFormat     string = "2006-01-02"
	MaxDaysPerCall int    = 60
)

// type
//
type Service struct {
	accessKey             string
	maxRetries            uint
	secondsBetweenRetries uint32
}

type ServiceConfig struct {
	AccessKey             string
	MaxRetries            *uint
	SecondsBetweenRetries *uint32
}

func NewService(config ServiceConfig) (*Service, *errortools.Error) {
	service := new(Service)

	if config.AccessKey == "" {
		return nil, errortools.ErrorMessage("AccessKey not provided")
	}
	service.accessKey = config.AccessKey

	if config.MaxRetries != nil {
		service.maxRetries = *config.MaxRetries
	} else {
		service.maxRetries = 0
	}

	if config.SecondsBetweenRetries != nil {
		service.secondsBetweenRetries = *config.SecondsBetweenRetries
	} else {
		service.secondsBetweenRetries = 3
	}

	return service, nil
}

// generic Get method
//
func (service *Service) Get(endpoint string, values url.Values, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, endpoint, values, nil, responseModel)
}

func (service *Service) httpRequest(httpMethod string, endpoint string, values url.Values, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	client := new(http.Client)

	//add access_key
	values.Add("access_key", service.accessKey)

	url := fmt.Sprintf("%s/%s?%s", APIURL, endpoint, values.Encode())
	fmt.Println(url)

	buffer := new(bytes.Buffer)
	buffer = nil

	if bodyModel != nil {
		b, err := json.Marshal(bodyModel)
		if err != nil {
			return nil, nil, errortools.ErrorMessage(err)
		}
		//fmt.Println(string(b)) //temp
		buffer = bytes.NewBuffer(b)
	}

	ee := new(errortools.Error)

	request, err := func() (*http.Request, error) {
		// function necessary because a Buffer nil pointer differs from a nil value
		if buffer == nil {
			return http.NewRequest(httpMethod, url, nil)
		}
		return http.NewRequest(httpMethod, url, buffer)
	}()

	ee.SetRequest(request)

	if err != nil {
		ee.SetMessage(err)
		return request, nil, ee
	}

	// Add authorization token to header
	request.Header.Set("Accept", "application/json")

	if bodyModel != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	// Send out the HTTP request
	response, e := utilities.DoWithRetry(client, request, service.maxRetries, service.secondsBetweenRetries)
	ee.SetResponse(response)

	if response != nil {
		// Check HTTP StatusCode
		if response.StatusCode < 200 || response.StatusCode > 299 {
			fmt.Println(fmt.Sprintf("ERROR in %s", httpMethod))
			fmt.Println("url", url)
			fmt.Println("StatusCode", response.StatusCode)

			ee.SetMessage(fmt.Sprintf("Server returned statuscode %v", response.StatusCode))
		}

		if response.Body != nil {

			defer response.Body.Close()

			b, err := ioutil.ReadAll(response.Body)
			if err != nil {
				ee.SetMessage(err)
				return request, response, ee
			}

			if e != nil {
				// try to unmarshal to ErrorResponse
				errorResponse := ErrorResponse{}
				errError := json.Unmarshal(b, &errorResponse)

				if errError == nil {
					if errorResponse.Error.Info != "" {
						ee.SetMessage(errorResponse.Error.Info)
					}
				} else {
					// try to unmarshal to string
					errorString := ""
					errError = json.Unmarshal(b, &errorString)

					if errorString != "" {
						ee.SetMessage(errorString)
					}
				}

				errortools.CaptureInfo(errError)

				return request, response, ee
			}

			if responseModel != nil {
				err = json.Unmarshal(b, &responseModel)
				if err != nil {
					ee.SetMessage(err)
					return request, response, ee
				}
			}
		}
	}

	if e != nil {
		return request, response, e
	}

	return request, response, nil
}
