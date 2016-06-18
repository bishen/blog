package fun

import (
	"fmt"
	"html/template"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Substr returns the substr from start to length.
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

// 去除html,然后在截取字符
func Subtext(src string, start, length int) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return Substr(strings.TrimSpace(src), start, length)
}

// HTML2str returns escaping text convert from html.
func HTML2str(html string) string {
	src := string(html)

	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//remove STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//remove SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

// DateFormat takes a time and a layout string and returns a string with the formatted date. Used by the template parser as "dateformat"
func DateFormat(t time.Time, layout string) (datestring string) {
	datestring = t.Format(layout)
	return
}

// DateFormat pattern rules.
var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", //A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// DateParse Parse Date use PHP time format.
func DateParse(dateString, format string) (time.Time, error) {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return time.ParseInLocation(format, dateString, time.Local)
}

// Date takes a PHP like date func to Go's time format.
func Date(t time.Time, format string) string {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return t.Format(format)
}

func Img(src string) template.HTML {
	if src == "" {
		return ""
	} else {
		return template.HTML("<img src=\"" + src + "\">")
	}
}

// Compare is a quick and dirty comparison function. It will convert whatever you give it to strings and see if the two values are equal.
// Whitespace is trimmed. Used by the template parser as "eq".
func Compare(a, b interface{}) (equal bool) {
	equal = false
	if strings.TrimSpace(fmt.Sprintf("%v", a)) == strings.TrimSpace(fmt.Sprintf("%v", b)) {
		equal = true
	}
	return
}

// CompareNot !Compare
func CompareNot(a, b interface{}) (equal bool) {
	return !Compare(a, b)
}

// NotNil the same as CompareNot
func NotNil(a interface{}) (isNil bool) {
	return CompareNot(a, nil)
}

// Str2html Convert string to template.HTML type.
func Str2html(raw string) template.HTML {
	return template.HTML(raw)
}

// Htmlquote returns quoted html string.
func Htmlquote(src string) string {
	//HTML编码为实体符号
	/*
	   Encodes `text` for raw use in HTML.
	       >>> htmlquote("<'&\\">")
	       '&lt;&#39;&amp;&quot;&gt;'
	*/

	text := string(src)

	text = strings.Replace(text, "&", "&amp;", -1) // Must be done first!
	text = strings.Replace(text, "<", "&lt;", -1)
	text = strings.Replace(text, ">", "&gt;", -1)
	text = strings.Replace(text, "'", "&#39;", -1)
	text = strings.Replace(text, "\"", "&quot;", -1)
	text = strings.Replace(text, "“", "&ldquo;", -1)
	text = strings.Replace(text, "”", "&rdquo;", -1)
	text = strings.Replace(text, " ", "&nbsp;", -1)

	return strings.TrimSpace(text)
}

// Htmlunquote returns unquoted html string.
func Htmlunquote(src string) string {
	//实体符号解释为HTML
	/*
	   Decodes `text` that's HTML quoted.
	       >>> htmlunquote('&lt;&#39;&amp;&quot;&gt;')
	       '<\\'&">'
	*/

	// strings.Replace(s, old, new, n)
	// 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换

	text := string(src)
	text = strings.Replace(text, "&nbsp;", " ", -1)
	text = strings.Replace(text, "&rdquo;", "”", -1)
	text = strings.Replace(text, "&ldquo;", "“", -1)
	text = strings.Replace(text, "&quot;", "\"", -1)
	text = strings.Replace(text, "&#39;", "'", -1)
	text = strings.Replace(text, "&gt;", ">", -1)
	text = strings.Replace(text, "&lt;", "<", -1)
	text = strings.Replace(text, "&amp;", "&", -1) // Must be done last!

	return strings.TrimSpace(text)
}

// AssetsJs returns script tag with src string.
func AssetsJs(src string) template.HTML {
	text := string(src)

	text = "<script src=\"" + src + "\"></script>"

	return template.HTML(text)
}

// AssetsCSS returns stylesheet link tag with src string.
func AssetsCSS(src string) template.HTML {
	text := string(src)

	text = "<link href=\"" + src + "\" rel=\"stylesheet\" />"

	return template.HTML(text)
}

// MapGet getting value from map by keys
// usage:
// Data["m"] = map[string]interface{} {
//     "a": 1,
//     "1": map[string]float64{
//         "c": 4,
//     },
// }
//
// {{ map_get m "a" }} // return 1
// {{ map_get m 1 "c" }} // return 4
func MapGet(arg1 interface{}, arg2 ...interface{}) (interface{}, error) {
	arg1Type := reflect.TypeOf(arg1)
	arg1Val := reflect.ValueOf(arg1)

	if arg1Type.Kind() == reflect.Map && len(arg2) > 0 {
		// check whether arg2[0] type equals to arg1 key type
		// if they are different, make conversion
		arg2Val := reflect.ValueOf(arg2[0])
		arg2Type := reflect.TypeOf(arg2[0])
		if arg2Type.Kind() != arg1Type.Key().Kind() {
			// convert arg2Value to string
			var arg2ConvertedVal interface{}
			arg2String := fmt.Sprintf("%v", arg2[0])

			// convert string representation to any other type
			switch arg1Type.Key().Kind() {
			case reflect.Bool:
				arg2ConvertedVal, _ = strconv.ParseBool(arg2String)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				arg2ConvertedVal, _ = strconv.ParseInt(arg2String, 0, 64)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				arg2ConvertedVal, _ = strconv.ParseUint(arg2String, 0, 64)
			case reflect.Float32, reflect.Float64:
				arg2ConvertedVal, _ = strconv.ParseFloat(arg2String, 64)
			case reflect.String:
				arg2ConvertedVal = arg2String
			default:
				arg2ConvertedVal = arg2Val.Interface()
			}
			arg2Val = reflect.ValueOf(arg2ConvertedVal)
		}

		storedVal := arg1Val.MapIndex(arg2Val)

		if storedVal.IsValid() {
			var result interface{}

			switch arg1Type.Elem().Kind() {
			case reflect.Bool:
				result = storedVal.Bool()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				result = storedVal.Int()
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				result = storedVal.Uint()
			case reflect.Float32, reflect.Float64:
				result = storedVal.Float()
			case reflect.String:
				result = storedVal.String()
			default:
				result = storedVal.Interface()
			}

			// if there is more keys, handle this recursively
			if len(arg2) > 1 {
				return MapGet(result, arg2[1:]...)
			}
			return result, nil
		}
		return nil, nil

	}
	return nil, nil
}
