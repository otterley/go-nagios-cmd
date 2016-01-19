package nagios_result

import (
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestHostCheckResult(t *testing.T) {
	var r = HostCheckResult{
		Timestamp: time.Unix(123456789, 0),
		Hostname:  "host1",
		Status:    NAGIOS_OK,
		Output:    "OK"}
	var output = r.String()
	var desiredOutput = "[123456789] PROCESS_HOST_CHECK_RESULT;host1;0;OK\n"
	if strings.Compare(output, desiredOutput) != 0 {
		t.Errorf("output = %s\ndesiredOutput = %s", output, desiredOutput)
	}
}

func TestServiceCheckResult(t *testing.T) {
	var r = ServiceCheckResult{
		Timestamp: time.Unix(123456789, 0),
		Hostname:  "host1",
		Service:   "service",
		Status:    NAGIOS_OK,
		Output:    "OK"}
	var output = r.String()
	var desiredOutput = "[123456789] PROCESS_SERVICE_CHECK_RESULT;host1;service;0;OK\n"
	if strings.Compare(output, desiredOutput) != 0 {
		t.Errorf("output = %s\ndesiredOutput = %s", output, desiredOutput)
	}
}

func TestHostCheckResultWithImplicitTimestamp(t *testing.T) {
	var r = HostCheckResult{
		Hostname: "host1",
		Status:   NAGIOS_OK,
		Output:   "OK"}
	var output = r.String()
	re, _ := regexp.Compile(`^\[\d+\] PROCESS_HOST_CHECK_RESULT;host1;0;OK`)
	if !re.MatchString(output) {
		t.Errorf("output = %s; did not match regexp %s", output, re)
	}
}

func TestServiceCheckResultWithImplicitTimestamp(t *testing.T) {
	var r = ServiceCheckResult{
		Hostname: "host1",
		Service:  "service",
		Status:   NAGIOS_OK,
		Output:   "OK"}
	var output = r.String()
	re, _ := regexp.Compile(`^\[\d+\] PROCESS_SERVICE_CHECK_RESULT;host1;service;0;OK`)
	if !re.MatchString(output) {
		t.Errorf("output = %s; did not match regexp %s", output, re)
	}
}
