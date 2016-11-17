package hateoas

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo"

	"gopkg.in/yaml.v2"
)

var swagger Swagger

// RootResponse offers HATEOAS links from the root endpoint of the API
func RootResponse(context echo.Context) error {
	hateoasObject := GetInfo()

	links, err := AddLinks(context.Path(), []string{})
	if err != nil {
		return context.JSON(http.StatusBadRequest, "An error occurred: "+err.Error())
	}

	hateoasObject.Links = links // Append the links to the full HATEOAS object that includes the general Swagger document information

	return context.JSON(http.StatusOK, hateoasObject)
}

// MergeSort takes two string arrays and shuffles them together (there has to be a better way to do this)
func MergeSort(first []string, second []string) string {
	var final []string

	for i := range first { // second should always be shorter than first because there's an empty string at the end of first
		if i < len(first) {
			final = append(final, first[i])
		}

		if i < len(second) {
			final = append(final, second[i])
		}
	}

	return strings.Join(final[:], "")
}

// EchoToSwagger converts paths from Echo syntax to Swagger syntax
func EchoToSwagger(path string) string {
	echoRegex := regexp.MustCompile(`\:(\w+)`)

	antiParameters := echoRegex.Split(path, -1)
	parameters := echoRegex.FindAllString(path, -1)

	for i := range parameters {
		parameters[i] = strings.Replace(parameters[i], ":", "{", 1)
		parameters[i] = parameters[i] + "}"
	}

	return MergeSort(antiParameters, parameters)
}

// Load loads a swagger.json file from an external URL
func Load(fileLocation string) error {
	response, err := http.Get(fileLocation)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("Received HTTP code " + strconv.Itoa(response.StatusCode) + " when attempting to retrieve Swagger document")
	}

	err = yaml.Unmarshal(contents, &swagger)
	if err != nil {
		return err
	}

	return nil
}

// GetInfo returns general information about the API (mainly for the root path)
func GetInfo() Root {
	returnStruct := Root{
		Title:       swagger.Info.Title,
		Description: swagger.Info.Description,
		Version:     swagger.Info.Version,
	}

	return returnStruct
}

// AddLinks searches through given paths
func AddLinks(path string, parameters []string) ([]Link, error) {
	allLinks := []Link{}
	contextPath := EchoToSwagger(path)
	contextRegex := "" // We populate this a few lines down from here

	if contextPath != "/" { // If we're dealing with a path that's not the root URL...
		// Make the path regex friendly by escaping / { and } characters
		contextPath = strings.Replace(contextPath, "/", `\/`, -1)
		contextPath = strings.Replace(contextPath, "{", `\{`, -1)
		contextPath = strings.Replace(contextPath, "}", `\}`, -1)

		contextRegex = `^` + contextPath + `\/[a-zA-Z{}]*$` // Append the usual regular exptression for finding dynamic arguments
	} else {
		contextRegex = `^\/[a-zA-Z{}]*$`
	}

	hateoasRegex := regexp.MustCompile(contextRegex)
	parameterRegex := regexp.MustCompile(`\{(.*?)\}`)

	log.Printf("%+v", swagger)

	for swaggerPaths := range swagger.Paths {
		match := hateoasRegex.MatchString(swaggerPaths)

		if match {
			antiParameters := parameterRegex.Split(swaggerPaths, -1)

			if swagger.Paths[swaggerPaths].Get != nil { // Make sure the matching path has a GET path
				link := Link{
					Rel:  swagger.Paths[swaggerPaths].Get.Summary,
					HREF: MergeSort(antiParameters, parameters),
				}

				allLinks = append(allLinks, link)

				log.Printf("%+V", allLinks)
			}
		}
	}

	return allLinks, nil
}
