package shop

import (
	"errors"

	"github.com/oklog/ulid/v2"
)

type Shop struct {
	ID          string
	Name        string
	Description string
	Latitude    float64
	Longitude   float64
}

func NewShop(
	name string,
	description string,
	latitude float64,
	longitude float64,
) (*Shop, error) {
	return newShop(ulid.Make().String(), name, description, latitude, longitude)
}

func Reconstruct(
	id string,
	name string,
	description string,
	latitude float64,
	longitude float64,
) (*Shop, error) {
	return newShop(id, name, description, latitude, longitude)
}

func newShop(
	id string,
	name string,
	description string,
	latitude float64,
	longitude float64,
) (*Shop, error) {

	// IDのバリデーション
	if id == "" {
		return nil, errors.New("ID is required")
	}

	// nameのバリデーション
	if name == "" {
		return nil, errors.New("name is required")
	}

	// latitudeのバリデーション
	if latitude == 0 {
		return nil, errors.New("latitude is required")
	}

	// longitudeのバリデーション
	if longitude == 0 {
		return nil, errors.New("longitude is required")
	}

	return &Shop{
		ID:          id,
		Name:        name,
		Description: description,
		Latitude:    latitude,
		Longitude:   longitude,
	}, nil
}

func (s *Shop) GetID() string {
	return s.ID
}

func (s *Shop) GetName() string {
	return s.Name
}

func (s *Shop) GetDescription() string {
	return s.Description
}

func (s *Shop) GetLatitude() float64 {
	return s.Latitude
}

func (s *Shop) GetLongitude() float64 {
	return s.Longitude
}
