package space

import (
	"database/sql"
	"errors"

	"github.com/d4vi13/SeuCantinho/internal/database"
	models "github.com/d4vi13/SeuCantinho/internal/models/space"
)

type SpaceRepository struct{}

func (repository *SpaceRepository) Init() {}

func (repository *SpaceRepository) GetSpace(location string, substation string) (*models.Space, error) {
	// Statement para consultar um espaço no banco
	query := `SELECT id, location, substation, price, capacity, image FROM spaces WHERE location = $1 AND substation = $2`
	space := &models.Space{}

	row := database.QueryRow(query, location, substation)

	err := row.Scan(&space.Id, &space.Location, &space.Substation, &space.Price, &space.Capacity, &space.Img)
	if err != nil {
		// Espaço não existe no banco de dados
		if err == sql.ErrNoRows {
			return nil, errors.New("space not found")
		}
		return nil, err
	}

	// Retorna o objeto do espaço
	return space, nil
}

func (repository *SpaceRepository) Insert(space *models.Space) (int, error) {
	// Statement para inserir um novo espaço
	query := `INSERT INTO spaces (location, substation, price, capacity, image) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	// Preenche os dados e obtém o ID
	var id int
	err := database.QueryRow(query, space.Location, space.Substation, space.Price, space.Capacity, space.Img).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
