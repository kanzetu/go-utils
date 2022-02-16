package x

import (
	"strconv"
	"strings"
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	
	"github.com/kanztu/go-utils/format"
)

func Stoi(s string) int{
	i, err := strconv.Atoi(s)
	Must(err)
	return i
}

func Itos(s int) string{
	return strconv.Itoa(s)
}

func Ftos(s float64) string{
	return strconv.FormatFloat(s, 'f', 0, 64)
}

func AddLeadingZeros(i int, z int) string{
	return fmt.Sprintf("%0" + Itos(z) + "d", i)
}

func Must(err error) {
	if err != nil {
		format.Error("Err", err.Error())
	}
}

func CreateDirIfNonExist(s string){
	if _, err := os.Stat(s); os.IsNotExist(err) {
		format.Verbose("DLer", "Folder " + s + " not exist, create it...")
		os.MkdirAll(s, 0777)
	}
}

func UrlDecode(s string) string{
	d, err := url.QueryUnescape(s)
	Must(err)
	return d
}

func SavePathName(s string) string{
	s = strings.ReplaceAll(s, ":", "-")
	s = strings.ReplaceAll(s, "/", "／")
	return s
}

func ResToByte(r *http.Response) []byte{
    b, err := ioutil.ReadAll(r.Body)
	Must(err)
	return b
}

func UniDecoder(u string) string{
    str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(u)), `\\u`, `\u`, -1))
	Must(err)
    return str
}

func ParseJson(s []byte) map[string]interface {}{
	var j interface{}
	json.Unmarshal(s, &j)
	return j.(map[string]interface{})
}

func ExtractFileName(url string) (string, string){
	ext_ok := []string{"gif", "png", "jpg", "jpeg", "mp4", "mkv", "webm"}
	r := strings.Split(url, "/")
	fname := strings.Split(r[len(r)-1], "?")[0]
	fname = strings.Split(fname, ";")[0]
	if len(r) == 1{
		return "", ""
	}else if len( fname ) == 0 {
		return "", ""	
	}
	for _, e := range ext_ok {
		ext := strings.Split(fname, ".")
		if e == ext[len(ext)-1] {
			return fname, "." + ext[len(ext)-1]	
		}
	}
	return "", ""
}
