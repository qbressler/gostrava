package utils

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

var code string

func GrantAccess() string {
	cmd := "xdg-open"
	args := "https://www.strava.com/oauth/authorize?client_id=18876&response_type=code&redirect_uri=http://localhost:8080/exchange_token&approval_prompt=force&scope=read,activity:read"
	exec.Command(cmd, args).Start()
	server := &http.Server{Addr: ":8080", Handler: nil}
	go func() {
		if err := server.ListenAndServe(); err != nil {
		}

	}()

	http.HandleFunc("/", getCode)
	return code
}

func getCode(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	splt := strings.Split(r.URL.RawQuery, "&")
	for _, j := range splt {
		t := j[:4]
		if t == "code" {
			s := strings.Split(j, "=")
			code = s[1]

		}
	}
	fmt.Fprintf(w, "You can close your browser now. The code is %s Thanks... ", code)
}
