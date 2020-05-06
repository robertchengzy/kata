package utils

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// 文件名过滤器
func FileNameFiter(name string) string {
	reg, _ := regexp.Compile("[/\\\\:*?<>|]")
	return reg.ReplaceAllString(name, "")
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

type MyStruct struct {
	Name string
	Age  int64
}

func (s *MyStruct) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

const MIN = 0.000001

// MIN 为用户自定义的比较精度
func IsEqual(f1, f2 float64) bool {
	return math.Dim(f1, f2) < MIN
}

//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func main() {
	myData := make(map[string]interface{})
	myData["Name"] = "Tony"
	myData["Age"] = int64(23)

	result := &MyStruct{}
	err := result.FillStruct(myData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

// 过滤 emoji 表情
func FilterEmoji(content string) string {
	newContent := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			newContent += string(value)
		}
	}
	return newContent
}

// 版本比较 1.0.0  0.9.9
func ComapreVersion(version1, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	len1 := len(v1)
	len2 := len(v2)

	length := len1
	if len1 < len2 {
		length = len(v2)
	}

	if len1 < length {
		v1 = append(v1, "0")
	}

	if len2 < length {
		v2 = append(v2, "0")
	}

	for i := 0; i < length; i++ {
		num1, _ := strconv.ParseInt(v1[i], 10, 64)
		num2, _ := strconv.ParseInt(v2[i], 10, 64)
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}

	return 0
}

func GetPointTime(t *time.Time) (*time.Time, error) {
	now := time.Now()
	nowBytes, _ := now.MarshalText()

	t = new(time.Time)
	err := t.UnmarshalText(nowBytes)
	return t, err
}

func randInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

var letter = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// 订单号
func createVoiceId() string {
	return time.Now().Format("20060102150405") + letter[randInt64(1, 25)] + letter[randInt64(1, 25)]
}

const (
	regular = `^1([1-9])\d{9}$`
)

func ValidateMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func GetLastCycle(startTimeStr, endTimeStr string) (string, string, error) {
	startTime, err := time.Parse("2006-01-02", startTimeStr)
	if err != nil {
		return "", "", nil
	}
	endTime, err := time.Parse("2006-01-02", endTimeStr)
	if err != nil {
		return "", "", nil
	}

	banlance := endTime.Sub(startTime) / time.Hour / 24

	lastStartTime := startTime.AddDate(0, 0, -int(banlance)-1)
	lastEndTime := startTime.AddDate(0, 0, -1)

	return lastStartTime.Format("2006-01-02"), lastEndTime.Format("2006-01-02"), nil
}
