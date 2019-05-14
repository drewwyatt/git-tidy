package reporting

import "fmt"

var startReportFormat = "⌛ %s...%s"
var continueReportFormat = "\r⌛ %s...%s"
var completeReportFormat = "\r✅  %s\n"
var failReportFormat = "\r❌ %s\n"
var successfullListing = "👍  %s\n"
var failedListing = "👎  %s\n%s\n"

func getFormatString(processes []process) string {
	if len(processes) > 0 {
		return continueReportFormat
	}

	return startReportFormat
}

type process struct {
	name      string
	succeeded bool
	errorMsg  string
}

func newProcess(name string) process {
	return process{name: name}
}

type Reporter struct {
	processes     []process
	longestString int
}

func NewReporter() Reporter {
	return Reporter{}
}

func (r *Reporter) Report(process string) {
	padding := r.updateProcessLengthAndGetPadding(process)
	fmt.Printf(getFormatString(r.processes), process, padding)
	r.processes = append(r.processes, newProcess(process))
}

func (r *Reporter) updateProcessLengthAndGetPadding(process string) string {
	var padding string
	if len(process) > r.longestString {
		r.longestString = len(process)
	} else {
		for i := len(process); i <= r.longestString; i++ {
			padding = padding + " "
		}
	}

	return padding
}

func (r *Reporter) Complete(message string) {
	str := fmt.Sprintf(completeReportFormat, message)
	fmt.Print(str)
	// fmt.Printf(completeReportFormat, message)
	r.listProcesses()
	r.processes = nil
	r.longestString = 0
}

func (r *Reporter) Fail(message string) {
	fmt.Printf(failReportFormat, message)
	r.listProcesses()
	r.processes = nil
}

func (r *Reporter) PassProcess(idx int) {
	r.processes[idx].succeeded = true
}

func (r *Reporter) FailProcess(idx int, msg string) {
	r.processes[idx].succeeded = false
	r.processes[idx].errorMsg = msg
}

func (r *Reporter) listProcesses() {
	for _, process := range r.processes {
		if process.succeeded {
			fmt.Printf(successfullListing, process.name)
		} else {
			fmt.Printf(failedListing, process.name, process.errorMsg)
		}
	}
}
