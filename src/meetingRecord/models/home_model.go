package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"meetingRecord/utils"
	"strconv"
	"strings"
)

type HomeBlock struct {
	Project      []ProjectLink
	Achievements string
	Goal         string
	Author       string
	CreateTime   string
	Link         string
	UpdateLink   string
	DeleteLink   string
	IsLogin      bool
}

type ProjectLink struct {
	ProjectName string
	ProjectUrl  string
}

type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

func MakeHomeBlocks(texts []Text, isLogin bool) template.HTML {
	htmlHome := ""

	for _, txt := range texts {
		block := HomeBlock{}
		block.IsLogin = isLogin
		block.Achievements = txt.Achievements
		block.Author = txt.User
		block.Goal = txt.Goal
		block.Project = createProjectLink(txt.Project)

		block.CreateTime = utils.FormatTimeStamp(txt.CreatedAt)
		block.Link = "/text/" + strconv.Itoa(int(txt.ID))
		block.UpdateLink = "/text/update?id=" + strconv.Itoa(int(txt.ID))
		block.DeleteLink = "/text/delete?id=" + strconv.Itoa(int(txt.ID))

		t, _ := template.ParseFiles("view/block/block.html")
		buffer := bytes.Buffer{}

		t.Execute(&buffer, block)
		htmlHome += buffer.String()
	}

	return template.HTML(htmlHome)
}

func createProjectLink(projects string) []ProjectLink {
	var projectLink []ProjectLink

	projectPara := strings.Split(projects, "&")
	for _, project := range projectPara {
		projectLink = append(projectLink, ProjectLink{project, "/?project=" + project})
	}
	return projectLink
}

func ConfigHomeFooterPageCode(page int) (pageCode HomeFooterPageCode) {

	sum := QueryTextSum()
	pageSize, _ := beego.AppConfig.Int("textsperpage")
	pageSum := (sum-1)/pageSize + 1
	pageCode.ShowPage = fmt.Sprintln(page, "/", pageSum)

	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}
	if page >= pageSum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}

	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)

	return pageCode
}
