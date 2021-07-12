package format

import (
		"time"
		"fmt"
		"os"
		"bufio"
)

var Version = "1.0"
var ENABLE_VERBOSE=1

var(
	BLACK="\033[0;30m";		DGRAY="\033[1;30m"
	RED="\033[0;31m";			LRED="\033[1;31m"
	GREEN="\033[0;32m";		LGREEN="\033[1;32m"
	ORANGE="\033[0;33m";	YELLOW="\033[1;33m"
	BLUE="\033[0;34m";			LBLUE="\033[1;34m"
	PURPLE="\033[0;35m";		LPURPLE="\033[1;35m"
	CYAN="\033[0;36m";		LCYAN="\033[1;36m"
	LGRAY="\033[0;37m";		WHITE="\033[1;37m"
	NC="\033[0m"
)


func GET_CURRENT_TIME() string{
	return DGRAY + time.Now().Format("2006-01-02 15:04:05") + NC
}

func Info(name string, str string){
	output := fmt.Sprintf("%s %s (%s) [+] %s %s",GET_CURRENT_TIME(),LGREEN,name,str,NC)
	fmt.Println(output)
}
func Error(name string, str string){
	output := fmt.Sprintf("%s %s (%s) [-] %s %s",GET_CURRENT_TIME(),LRED,name,str,NC)
	fmt.Println(output)
}

func Warning(name string, str string){
	output := fmt.Sprintf("%s %s (%s) [?] %s %s",GET_CURRENT_TIME(),YELLOW,name,str,NC)
	fmt.Println(output)
}

func Verbose(name string, str string){
	if ENABLE_VERBOSE != 0{
		output := fmt.Sprintf("%s %s (%s) [.] %s %s",GET_CURRENT_TIME(),LGRAY,name,str,NC)
		fmt.Println(output)
	}
}

func Special(name string, str string){
	output := fmt.Sprintf("%s %s (%s) [*] %s %s",GET_CURRENT_TIME(),LCYAN,name,str,NC)
	fmt.Println(output)
}

func Input() string{
	output := fmt.Sprintf("%s %s (%s)[>] ",GET_CURRENT_TIME(),LPURPLE,"Input")
	fmt.Print(output)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Print(NC)
	return text
}

func Testing(){
	ENABLE_VERBOSE=1
	Info("Test", "This is Info")
	Error("Test", "This is Error")
	Warning("Test", "This is Warning")
	Verbose("Test", "This is Verbose")
	Special("Test", "This is Special")
	a := Input()
	Info("Test", "Your input is " + a)
}