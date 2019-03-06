package parser

import (
	"go_pratice_code/learn_ccmouse_code/concurrent_crawler/engine"
	"go_pratice_code/learn_ccmouse_code/concurrent_crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">(\d+)cm</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">月收入:([^<]+)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">(\d+)kg</div>`)

// 没有性别字段
var genderRe = regexp.MustCompile(``)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var hukouRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798="">([^<]+)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798="">([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	// 名字
	profile.Name = name
	// 年龄
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	//  身高
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}
	// 体重
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}
	// 收入
	profile.Income = extractString(contents, incomeRe)
	// 性别
	profile.Gender = extractString(contents, genderRe)
	// 是否有车
	profile.Car = extractString(contents, carRe)
	// 教育状况
	profile.Education = extractString(contents, educationRe)
	// 户口
	profile.Hukou = extractString(contents, hukouRe)
	// 是否有房
	profile.House = extractString(contents, houseRe)
	// 婚姻状况
	profile.Marriage = extractString(contents, marriageRe)
	// 职业
	profile.Occupation = extractString(contents, occupationRe)
	// 星座
	profile.Xinzuo = extractString(contents, xinzuoRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
