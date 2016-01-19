package nagios_result

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type StatusCode int

const defaultCommandFilePath = "/var/spool/nagios/cmd/nagios.cmd"

const (
	NAGIOS_OK StatusCode = iota
	NAGIOS_WARNING
	NAGIOS_CRITICAL
	NAGIOS_UNKNOWN
)

type HostCheckResult struct {
	Hostname  string
	Status    StatusCode
	Output    string
	Timestamp time.Time
}

type ServiceCheckResult struct {
	Hostname  string
	Service   string
	Status    StatusCode
	Output    string
	Timestamp time.Time
}

var cf *os.File

func (r HostCheckResult) String() string {
	var t time.Time
	if r.Timestamp.IsZero() {
		t = time.Now()
	} else {
		t = r.Timestamp
	}
	return fmt.Sprintf("[%d] PROCESS_HOST_CHECK_RESULT;%s;%d;%s\n",
		t.Unix(),
		r.Hostname,
		r.Status,
		r.Output)
}

func (r ServiceCheckResult) String() string {
	var t time.Time
	if r.Timestamp.IsZero() {
		t = time.Now()
	} else {
		t = r.Timestamp
	}
	return fmt.Sprintf("[%d] PROCESS_SERVICE_CHECK_RESULT;%s;%s;%d;%s\n",
		t.Unix(),
		r.Hostname,
		r.Service,
		r.Status,
		r.Output)
}

func connectToNagios(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not open %s: %v", f, path, err))
	}
	cf = f
	return nil
}

type submittable interface {
	String() string
}

func Submit(s submittable) error {
	err := connectToNagios(defaultCommandFilePath)
	if err != nil {
		panic(err)
	}
	return nil
	_, err = io.WriteString(cf, s.String())
	return err
}
