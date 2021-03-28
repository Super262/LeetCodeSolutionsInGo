package unionfind

import (
	"sort"
	"strings"
)

func accountsMerge(accounts [][]string) [][]string {
	if accounts == nil || len(accounts) == 0 {
		return [][]string{}
	}
	emailFather := make(map[string]string)
	emailUser := make(map[string]string)
	for _, ac := range accounts {
		for i := 1; i < len(ac); i++ {
			emailFather[ac[i]] = ac[i]
			emailUser[ac[i]] = ac[0]
		}
	}
	for _, ac := range accounts {
		for i := 2; i < len(ac); i++ {
			union0721(&emailFather, &ac[1], &ac[i])
		}
	}
	emailCluster := make(map[string]map[string]int)
	currentFather := ""
	ok := false
	for _, ac := range accounts {
		for i := 1; i < len(ac); i++ {
			currentFather = findAndCompress0721(&emailFather, ac[i])
			_, ok = emailCluster[currentFather]
			if !ok {
				emailCluster[currentFather] = make(map[string]int)
			}
			emailCluster[currentFather][ac[i]] = 1
		}
	}
	result := make([][]string, len(emailCluster), len(emailCluster))
	i := 0
	j := 0
	for mainEmail, childSet := range emailCluster {
		j = 0
		result[i] = make([]string, len(childSet), len(childSet))
		for email, _ := range childSet {
			result[i][j] = email
			j++
		}
		sort.Strings(result[i])
		result[i] = append([]string{emailUser[mainEmail]}, result[i]...)
		i++
	}
	return result
}

func union0721(father *map[string]string, emailA *string, emailB *string) {
	rootA := findAndCompress0721(father, *emailA)
	rootB := findAndCompress0721(father, *emailB)
	(*father)[rootB] = rootA
}

func findAndCompress0721(father *map[string]string, head string) string {
	path := make([]string, 0)
	for strings.Compare((*father)[head], head) != 0 {
		path = append(path, head)
		head = (*father)[head]
	}
	for _, v := range path {
		(*father)[v] = head
	}
	return head
}
