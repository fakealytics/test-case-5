import("fmt")

package logger

func log(message string) string {
	if (message == "Y") {
		fmt.Println("Why is there Y?")
	} else {
		fmt.Println(message)
	}

	return message;
}