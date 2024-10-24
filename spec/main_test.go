package spec_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/CDCgov/phinvads-go/internal/app"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// func apiCall(query string) {
// 	response, err := http.Get("http://localhost:4000/" + query)

// 	if err != nil {
// 		fmt.Print(err.Error())
// 		os.Exit(1)
// 	}

// 	responseData, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Status Code:", response.StatusCode)
// 	fmt.Println(string(responseData))
// }

var ENDPOINTS = []string{
	"/api",
	"/api/code-systems/",
	"/api/code-system-concepts/",
	"/api/value-sets/",
	"/api/views/",
	"/load-hot-topics/",
}

func queryForCode(url string) int {
	req, err := http.NewRequest("GET", url, nil)
	Expect(err).NotTo(HaveOccurred(), "could not create request")
	resp, err := http.DefaultClient.Do(req)
	Expect(err).NotTo(HaveOccurred(), "could not create response")
	// bodyBytes, err := io.ReadAll(resp.Body)
	// bodyString := string(bodyBytes)
	defer func() {
		err := resp.Body.Close()
		Expect(err).NotTo(HaveOccurred(), "could not close response body")
	}()
	return resp.StatusCode
}

var _ = Describe("Spec", Ordered, func() {
	var (
		app    *app.Application
		server *httptest.Server
	)

	// Set up the app before running tests
	BeforeAll(func() {
		// Read the dump file
		dumpData, err := os.ReadFile("./spec.dump")
		if err != nil {
			Expect(err).NotTo(HaveOccurred(), "could not read dump file")
		}

		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			app.GetMux().ServeHTTP(w, r)
			// Simulate the response from the server
			w.Write(dumpData)
		}))

		http.DefaultClient = server.Client()

		Expect(err).NotTo(HaveOccurred(), "could not write dump file")
		// Load the dump data into the app
		fmt.Println("Server started on", server.URL)
	})

	AfterAll(func() {
		server.Close()
		fmt.Println("Server closed on", server.URL)
	})

	Describe("While GETTING endpoints", func() {
		Context("inside phinvads", func() {
			for _, endpoint := range ENDPOINTS {
				endpoint := endpoint
				It(fmt.Sprintf("the %s endpoint PASSES,", endpoint), func() {
					resp := queryForCode(server.URL + endpoint)
					Expect(resp).To(Equal(http.StatusOK), "expected status 200 OK")
					fmt.Println("GET on", endpoint, "returned with status Code:", resp)
				})
			}
		})
	})
})
