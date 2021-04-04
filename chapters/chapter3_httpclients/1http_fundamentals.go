package fundy


import (
	"io/ioutil"
	"net/http"
	"fmt"
	"log"
)

func fundamentals(){
	// basic http functionality
	response1, error := http.Get("http:www.google.com/robots.txt")
	//read response body. not shown
	defer rresponse1.Body.Close()
	response2, err := http.Head("http:www.google.com/robots.txt")
	// read reponse body. not shown
	defer response2.Body.Close()
	form := url.Values{}
	form.Add("foo", "bar")
	response, err := http.Post(
		"http:www.google.com/robots.txt",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	//read response body, not show
	defer response3.Body.Close
	//postform can be used instead of post
	response3, err = http.PostForm("http:www.google.com/robots.txt", form)
	//then read response body and close
}

func deleteRequest(){
	//delete by sending a call witout a http body
	request, error := http.NewRequest("DELETE", "http:www.google.com/robots.txt", nil)
	var client http.Client
	response, err := client.Do(request)
	//read response body and close
}

func putRequest() {
	//Put reuest, a patch looks similar
	request, error := http.NewRequest {
		"PUT",
		"http:www.google.com/robots.txt",
		strings.NewReader(form.Encode()),
	}
	var client http.Client
	response, err := client.Do(request)
	//read resonse body and close
}

func getWithMoreFunctionality() {
	response, error := http.Get("http:www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	//print http status
	fmt.Println(response.Status)
	//read and display response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	response.Body.Close

	//you can access the Sttatus parementer or StatusCode paramenter that accesses
	//only the integer portion of the status string
	//Response containes body param which is type io.ReadCloser
	//an io.Readcloser 
}