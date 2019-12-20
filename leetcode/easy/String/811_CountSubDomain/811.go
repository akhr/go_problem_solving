package countsubdomain811

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := map[string]int{}
	res := []string{}
	for _, cpdomain := range cpdomains {
		count, _ := strconv.Atoi(strings.Fields(cpdomain)[0])
		subDomains := getSubDomains(strings.Fields(cpdomain)[1])
		for _, subDomain := range subDomains {
			if _, ok := m[subDomain]; ok {
				m[subDomain] = m[subDomain] + count
			} else {
				m[subDomain] = count
			}
		}
	}
	for subDomain, count := range m {
		res = append(res, fmt.Sprintf("%d %s", count, subDomain))
	}
	return res
}

func getSubDomains(domain string) []string {
	sD := strings.Split(domain, ".")
	for i := len(sD) - 2; i >= 0; i-- {
		sD[i] = sD[i] + "." + sD[i+1]
	}
	return sD
}
