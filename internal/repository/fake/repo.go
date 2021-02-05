package fake

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/2785/demo/internal/models"
)

type FakeRepo struct {
	data []models.Beer
}

func New(f string) (*FakeRepo, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var j []models.Beer

	err = json.Unmarshal(b, &j)
	if err != nil {
		return nil, err
	}

	return &FakeRepo{data: j}, nil
}

func (fr *FakeRepo) GetById(ctx context.Context, id string) (*models.Beer, error) {
	for _, v := range fr.data {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}
