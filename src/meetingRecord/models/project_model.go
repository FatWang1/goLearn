package models

import (
	"strings"
)

func HandleProjectsListData(projects []string) map[string]int {
	var projectsMap = make(map[string]int)

	for _, project := range projects {
		projectList := strings.Split(project, "&")
		for _, value := range projectList {
			projectsMap[value]++
		}
	}
	return projectsMap
}
