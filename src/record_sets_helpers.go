/*
Copyright 2018 Comcast Cable Communications Management, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import "strconv"

func typeSwitch(t string) string {
	switch t {
	case "A", "a":
		return "A"
	case "AAAA", "aaaa":
		return "AAAA"
	case "CNAME", "cname":
		return "CNAME"
	}
	return ""
}

func getRecordValue(records []string, recordValue interface{}, recordPrepend string) []string {
	var strVal string
	switch recordValue.(type) {
	case int:
		strVal = strconv.Itoa(recordValue.(int))
	case string:
		strVal = recordValue.(string)
	default:
		strVal = ""
	}
	if strVal != "" && strVal != "0" {
		records = append(records, recordPrepend+": "+strVal)
	}

	return records
}
