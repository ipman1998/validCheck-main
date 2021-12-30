package checkStatus

import (
	"fmt"
	"net/http"
	"time"
	"validCheck/util"
)

var timeOut = time.Duration(5) * time.Second

func checkUri_All(uri string, ref string, userAgent string) (int, error) {
	client := http.Client{
		Timeout: timeOut,
	}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println(err.Error())
		return 404, err
	}
	req.Header = http.Header{
		"User-Agent": []string{userAgent},
		"Referer":    []string{ref},
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error(), "statusCode: ", resp.StatusCode)
		return 404, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func CheckStatus(uri string, jndiString string) {
	util.Status200Logger.Printf(`Input get from file: %s`,jndiString)

	finalUri := uri + "/" + jndiString
	ref := uri + "/" + jndiString
	query := "?url="
	userAgent := jndiString
	if status, err := checkUri_All(finalUri, "", ""); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf(
			`--------input---------
 case: only uri
 uri : %s
 referer: %s
 userAgent: %s
--------output---------
 statusCode: %d
 
`, finalUri, "", "", status)
		if status != 403 {
			util.Status200Logger.Printf(
				`--------input---------
 case: only uri
 uri : %s
 referer: %s
 userAgent: %s
--------output---------
 statusCode: %d

`, finalUri, "", "", status)
		}
	}

	finalUri = uri + "/" + query + jndiString
	if status, err := checkUri_All(finalUri, "", ""); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf(
			`--------input---------
 case: param
 uri : %s
 referer: %s
 userAgent: %s
--------output---------
 statusCode: %d

`, finalUri, "", "", status)
		if status != 403 {
			util.Status200Logger.Printf(
				`--------input---------
 case: param
 uri : %s
 referer: %s
 userAgent: %s
--------output---------
 statusCode: %d

`, finalUri, "", "", status)
		}
	}

	finalUri = uri
	if status, err := checkUri_All(finalUri, ref, ""); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf(
			`--------input--------
 case: ref
 uri : %s
 referer: %s
 userAgent: %s
--------output---------
 statusCode: %d

`, finalUri, ref, "", status)
		if status != 403 {
			util.Status200Logger.Printf(
				`--------input---------
 case: ref
 uri : %s
 referer: %s
 userAgent: %s
--------output---------
 statusCode: %d

`, finalUri, ref, "", status)
		}
	}

	finalUri = uri
	if status, err := checkUri_All(finalUri, "", userAgent); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf(
			`--------input---------
 case: userAgent
 uri : %s
 referer: %s
 userAgent: %s
--------output---------
 statusCode: %d

`, finalUri, "", userAgent, status)

		if status != 403 {
			util.Status200Logger.Printf(
				`--------input---------
 case: userAgent
 uri : %s
 referer: %s
 userAgent: %s
 --------output---------
 statusCode: %d

		`, finalUri, "", userAgent, status)
		}
	}

}
