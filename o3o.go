package o3o

import (
	"math/rand"
	"strings"
	"time"
)

func O3O(key string) interface{} {
	emoticons := mapEmoticons()
	if key == "" {
		return getTags(emoticons)
	}
	if key == "random" || key == "*" {
		return randomEmoticons(emoticons)
	}
	yans, ok := emoticons[key]
	if !ok {
		return nil
	}
	return fetchRandom(yans)
}

func mapEmoticons() map[string][]string {
	s := map[string][]string{}
	for _, line := range yans {
		tags := strings.Split(line.tag, " ")
		for _, tag := range tags {
			if _, ok := s[tag]; ok {
				s[tag] = append(s[tag], line.yan...)
			} else {
				s[tag] = line.yan
			}
		}
	}
	return s
}

func randomEmoticons(emoticons map[string][]string) string {
	if emoticons == nil {
		emoticons = mapEmoticons()
	}
	yans := getYans(emoticons)
	return fetchRandom(yans)
}

func fetchRandom(s []string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	rand.Seed(time.Now().UTC().UnixNano())
	n := rand.Intn(length)
	return s[n]
}

func getTags(e map[string][]string) (tags []string) {
	for k, _ := range e {
		tags = append(tags, k)
	}
	return
}

func getYans(e map[string][]string) (yans []string) {
	for _, v := range e {
		yans = append(yans, v...)
	}
	return
}
