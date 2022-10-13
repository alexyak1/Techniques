package migrations

import (
	"fmt"
	"log"
	"main/database"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

type TechniqueUpdate struct {
	Id       *string `json:"id"`
	Name     string  `json:"name"`
	Belt     string  `json:"belt"`
	ImageURL string  `json:"image_url"`
	Type     string  `json:"type"`
}

func connectToDB() {
	config := database.Config{
		ServerName: "sql.freedb.tech",
		User:       "freedb_alexyak1",
		Hash:       "NEED_TO_SET_PASSWORD",
		DB:         "freedb_techniques",
	}
	connectionString := database.GetConnectionString(config)

	err := database.Connect(connectionString)
	if err != nil {
		fmt.Println("Failed while connectiong to database", err)
	}
}

func RunMigration() error {

	techniques, err := getAllTechniques()
	if err != nil {
		log.Println("TECHNIQUES IS NOT HERE")
		return err
	}

	imageIds := getImageIds(techniques)
	err = insertImageIdToTechniques(imageIds)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func getAllTechniques() ([]TechniqueUpdate, error) {
	fmt.Println("getAllTechniques")
	connectToDB()

	rows, err := database.Connector.DB().Query("SELECT id,name,belt,image_url,type FROM `techniques`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var techniques []TechniqueUpdate

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var techn TechniqueUpdate
		if err := rows.Scan(&techn.Id, &techn.Name, &techn.Belt,
			&techn.ImageURL, &techn.Type); err != nil {
			return techniques, err
		}
		techniques = append(techniques, techn)
	}
	if err = rows.Err(); err != nil {
		return techniques, err
	}

	return techniques, nil
}

func insertImageIdToTechniques(imageIds map[*string]string) error {
	connectToDB()

	for techniqueId, imageId := range imageIds {
		fmt.Println(imageId)

		_, err := database.Connector.DB().Exec("update techniques set image_id = ? where id = ?", imageId, techniqueId)
		if err != nil {
			return err
		}
	}
	return nil
}

func getImageIds(techniques []TechniqueUpdate) map[*string]string {
	var imageIds = make(map[*string]string)
	for _, technique := range techniques {
		re := regexp.MustCompile("d/(.*)/v")
		imageId := re.FindString(technique.ImageURL)
		imageId = imageId[2 : len(imageId)-2]
		fmt.Println(imageId)

		imageIds[technique.Id] = imageId
	}
	return imageIds
}
