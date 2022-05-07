package daemon

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func daemonResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, err)
}

const ShellToUse = "bash"

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func UpdateDaemon(w http.ResponseWriter, r *http.Request) {
	// the FormFile function takes in the POST input id file
	file, _, err := r.FormFile("file")
	if err != nil {
		daemonResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()
	daemonPath := r.Form.Get("path")

	out, err := os.OpenFile(daemonPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		daemonResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		daemonResponse(w, http.StatusInternalServerError, err)
	}

	script := r.Form.Get("script")
	err, ret1, ret2 := Shellout(script)
	if err != nil {
		err = fmt.Errorf("%v: %s,%s\n", err, ret1, ret2)
		daemonResponse(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Fprintf(w, "task has been executed successfully: %s, %s", ret1, ret2)
	w.WriteHeader(http.StatusOK)
}
