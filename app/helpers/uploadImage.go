package helpers

import (
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/99designs/gqlgen/graphql"
)

func UploadImage(file *graphql.Upload, customfilePath *string) (*string, error) {
	var filePath string
	if customfilePath == nil {
		var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

		b := make([]rune, 20)
		for i := range b {
			b[i] = letters[rnd.Intn(len(letters))]
		}
		filePath = "./uploads/" + string(b) + file.Filename
	} else {
		filePath = "./uploads/" + *customfilePath
	}

	out, err := os.Create(filePath)
	if err != nil {
		log.Error().Err(err).Msg("error creating file")

		return nil, err
	}
	defer out.Close()

	_, err = io.Copy(out, file.File)
	if err != nil {
		log.Error().Err(err).Msg("error copying file")

		return nil, err
	}

	return &filePath, nil
}
