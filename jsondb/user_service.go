package jsondb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kdelwat/recipaliser"
)

var _ recipaliser.UserService = &UserService{}

type UserService struct {
	database *Database
}

func (us *UserService) User(name recipaliser.UserName) (recipaliser.User, error) {
	panic("implement me")
}

func (us *UserService) CreateUser(user *recipaliser.User) error {
	userFilename := filepath.Join(us.database.Path, user.Name+".json")

	// Ensure the user doesn't already exist.
	// Need to check the error twice to make sure it's not an error with stat
	// From https://stackoverflow.com/a/12518877
	if _, err := os.Stat(userFilename); err == nil {
		return recipaliser.UserAlreadyExists
	} else if !os.IsNotExist(err) {
		return err
	}

	user.NutrientReferenceValues = loadReferenceValuesFromTemplate(user)

	marshalledUser, err := json.Marshal(*user)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(userFilename, marshalledUser, 0644)

	if err != nil {
		return err
	}

	return nil
}

func loadReferenceValuesFromTemplate(user *recipaliser.User) map[string]recipaliser.ReferenceValues {
	availableTemplates, err := ioutil.ReadDir("./data")

	if err != nil {
		log.Fatal(err)
	}

	relevantTemplate, err := findRelevantTemplate(availableTemplates, user)

	if err != nil {
		log.Fatal(err)
	}

	// TODO: better path joining
	templateFile, err := ioutil.ReadFile("./data/" + relevantTemplate)

	if err != nil {
		log.Fatal(err)
	}

	var templateData map[string]interface{}

	if err := json.Unmarshal(templateFile, &templateData); err != nil {
		log.Fatal(err)
	}

	nutrientReferenceValues := map[string]recipaliser.ReferenceValues{}

	// TODO: fix this horrific duck typing
	templateReferenceValues := templateData["referenceValues"].(map[string]interface{})

	for nutrient := range templateReferenceValues {
		fmt.Println(nutrient)

		referenceValues := recipaliser.ReferenceValues{}

		templateReferenceTypes := templateReferenceValues[nutrient].(map[string]interface{})

		for referenceType := range templateReferenceTypes {
			fmt.Println(referenceType)
			referenceTypeValue := templateReferenceTypes[referenceType].(map[string]interface{})

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(referenceTypeValue["value"])
			referenceValues[referenceType] = recipaliser.ReferenceValue{
				Source: referenceTypeValue["source"].(string),
				Value:  referenceTypeValue["value"].(float64),
			}
		}

		nutrientReferenceValues[nutrient] = referenceValues
	}

	return nutrientReferenceValues
}

func findRelevantTemplate(availableTemplateFiles []os.FileInfo, user *recipaliser.User) (string, error) {
	for _, availableTemplate := range availableTemplateFiles {
		filename := strings.TrimSuffix(availableTemplate.Name(), ".json")

		if user.IsLactating && !strings.Contains(filename, "lactation") {
			continue
		}

		if user.IsPregnant && !strings.Contains(filename, "pregnancy") {
			continue
		}

		filenameParts := strings.Split(filename, "-")

		if user.Sex != filenameParts[0] {
			continue
		}

		ageLower, err := strconv.Atoi(filenameParts[1])

		if err != nil {
			log.Fatal(err)
		}

		ageUpper, err := strconv.Atoi(filenameParts[2])

		if err != nil {
			log.Fatal(err)
		}

		if user.Age < ageLower || user.Age > ageUpper {
			continue
		}

		return availableTemplate.Name(), nil
	}

	// TODO: remove this service implementation-specific error
	return "", recipaliser.NoRelevantUserTemplate
}
