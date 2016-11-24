package helper

import (
	"bytes"
	"io"
	"mime/multipart"
	netURL "net/url"
	"os"
	"reflect"
	"strconv"

	apiKit "github.com/emacsist/golang-weibo-sdk/kit"

	apiURL "github.com/emacsist/golang-weibo-sdk/url"

	"net/http"

	"github.com/Sirupsen/logrus"

	"errors"
)

// BuildGetQuery ： 构建查询参数，已经进行URL encode了的.
func BuildGetQuery(param interface{}, URL string, accessToken string) string {
	var query bytes.Buffer
	query.WriteString(URL + "?")
	queryParam := BuildPostQuery(param, accessToken)
	query.WriteString(queryParam)
	return query.String()
}

// BuildPostQuery : 构建Post请求，返回的是 name=Ava&friend=Jess&friend=Sarah&friend=Zoe 类似这样子的格式
func BuildPostQuery(param interface{}, accessToken string) string {
	paramsMap := netURL.Values{}
	paramsMap.Add(apiKit.FirstLower(apiURL.ParamNameAccessToken), accessToken)

	if param == nil {
		return paramsMap.Encode()
	}
	v := reflect.ValueOf(param)
	fileds := v.Elem()
	for i := 0; i < fileds.NumField(); i++ {
		field := fileds.Field(i)

		// Ignore fields that don't have the same type as a string
		if field.Type().Name() == "string" {
			//如果是一个字符串对象
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(string)
			if len(fieldValue) > 0 {
				paramsMap.Add(apiKit.FirstLower(fieldName), fieldValue)
				//kit.WriteKV(&query, fieldName, fieldValue)
			}
		} else if field.Type().Name() == "int" {
			//如果是整型
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(int)

			if fieldValue >= 0 {
				paramsMap.Add(apiKit.FirstLower(fieldName), strconv.Itoa(fieldValue))
				//kit.WriteKV(&query, fieldName, strconv.Itoa(fieldValue))
			}
		} else if field.Type().Name() == "int64" {
			//如果是整型
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(int64)

			if fieldValue >= 0 {
				paramsMap.Add(apiKit.FirstLower(fieldName), strconv.FormatInt(fieldValue, 10))
				//kit.WriteKV(&query, fieldName, strconv.FormatInt(fieldValue, 10))
			}
		} else if field.Type().Name() == "int32" {
			//如果是整型
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(int32)

			if fieldValue >= 0 {
				paramsMap.Add(apiKit.FirstLower(fieldName), strconv.FormatInt(int64(fieldValue), 10))
				//kit.WriteKV(&query, fieldName, strconv.FormatInt(int64(fieldValue), 10))
			}
		} else if field.Type().Name() == "int8" {
			//如果是整型
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(int8)

			if fieldValue >= 0 {
				paramsMap.Add(apiKit.FirstLower(fieldName), strconv.FormatInt(int64(fieldValue), 10))
				//kit.WriteKV(&query, fieldName, strconv.FormatInt(int64(fieldValue), 10))
			}
		} else if field.Type().Name() == "int16" {
			//如果是整型
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(int16)

			if fieldValue >= 0 {
				paramsMap.Add(apiKit.FirstLower(fieldName), strconv.FormatInt(int64(fieldValue), 10))
				//kit.WriteKV(&query, fieldName, strconv.FormatInt(int64(fieldValue), 10))
			}
		}
	}
	return paramsMap.Encode()
}

// BuildPostMultiPartQuery : 构建上传文件的请求，返回的是 request 对象
func BuildPostMultiPartQuery(param interface{}, URL string, accessToken string, fieldName string, file *os.File) (*http.Request, error) {

	if file == nil {
		return nil, errors.New("file is nil")
	}
	defer file.Close()
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	v := reflect.ValueOf(param)
	fileds := v.Elem()
	bodyWriter.WriteField(apiKit.FirstLower(apiURL.ParamNameAccessToken), accessToken)
	for i := 0; i < fileds.NumField(); i++ {
		field := fileds.Field(i)
		// Ignore fields that don't have the same type as a string
		if field.Type().Name() == "string" {
			//如果是一个字符串对象
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(string)
			if len(fieldValue) > 0 {
				bodyWriter.WriteField(apiKit.FirstLower(fieldName), fieldValue)
			}
		} else if field.Type().Name() == "int" {
			//如果是整型
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(int)

			if fieldValue >= 0 {
				bodyWriter.WriteField(apiKit.FirstLower(fieldName), strconv.Itoa(fieldValue))
			}
		} else if field.Type().Name() == "int64" {
			//如果是整型
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(int64)

			if fieldValue >= 0 {
				bodyWriter.WriteField(apiKit.FirstLower(fieldName), strconv.FormatInt(fieldValue, 10))
			}
		} else if field.Type().Name() == "int32" {
			//如果是整型
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(int32)

			if fieldValue >= 0 {
				bodyWriter.WriteField(apiKit.FirstLower(fieldName), strconv.FormatInt(int64(fieldValue), 10))
			}
		} else if field.Type().Name() == "int8" {
			//如果是整型
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(int8)

			if fieldValue >= 0 {
				bodyWriter.WriteField(apiKit.FirstLower(fieldName), strconv.FormatInt(int64(fieldValue), 10))
			}
		} else if field.Type().Name() == "int16" {
			//如果是整型
			fieldName := fileds.Type().Field(i).Name
			fieldValue := field.Interface().(int16)

			if fieldValue >= 0 {
				bodyWriter.WriteField(apiKit.FirstLower(fieldName), strconv.FormatInt(int64(fieldValue), 10))
			}
		}
	}

	fileWriter, fileError := bodyWriter.CreateFormFile(fieldName, file.Name())
	if fileError != nil {
		logrus.Errorf("CreateFormFile error :%v\n", fileError.Error())
		return nil, fileError
	}

	//iocopy
	_, fileError = io.Copy(fileWriter, file)
	if fileError != nil {
		logrus.Errorf("CopyFile error :%v\n", fileError.Error())
		return nil, fileError
	}

	//defer bodyWriter.Close()

	reqObject, reqErr := http.NewRequest("POST", URL, bodyBuf)
	reqObject.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	return reqObject, reqErr
}
