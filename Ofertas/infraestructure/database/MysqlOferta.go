package database

import (
	"APICRUD/Ofertas/domain/entities"
	"database/sql"
	"errors"
	"fmt"
	//"time"
)

type MysqlOfertasRepository struct {
	db *sql.DB
}

func NewMysqlOfertasRepository(db *sql.DB) *MysqlOfertasRepository {
	return &MysqlOfertasRepository{db: db}
}

func (repo *MysqlOfertasRepository) MostrarOfertas() ([]*entities.Ofertas, error) {
	rows, err := repo.db.Query("SELECT id_oferta, nombre, descripcion, fecha_inicio, fecha_fin FROM ofertas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ofertas := []*entities.Ofertas{}
	for rows.Next() {
		oferta := &entities.Ofertas{}
		err := rows.Scan(&oferta.Id_oferta, &oferta.Nombre, &oferta.Descripcion, &oferta.Fecha_inicio, &oferta.Fecha_fin)
		if err != nil {
			return nil, err
		}
		ofertas = append(ofertas, oferta)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ofertas, nil
}

func (repo *MysqlOfertasRepository) CrearOfertas(oferta *entities.Ofertas) error {
	query := "INSERT INTO ofertas (nombre, descripcion, fecha_inicio, fecha_fin) VALUES (?, ?, ?, ?)"
	_, err := repo.db.Exec(query, oferta.Nombre, oferta.Descripcion, oferta.Fecha_inicio, oferta.Fecha_fin)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MysqlOfertasRepository) Actualizar(oferta *entities.Ofertas) error {
	query := "UPDATE ofertas SET nombre = ?, descripcion = ?, fecha_inicio = ?, fecha_fin = ? WHERE id_oferta = ?"
	result, err := repo.db.Exec(query, oferta.Nombre, oferta.Descripcion, oferta.Fecha_inicio, oferta.Fecha_fin, oferta.Id_oferta)
	if err != nil {
		return err
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if numRows == 0 {
		return errors.New("no se encontró la oferta para actualizar")
	}
	return nil
}

func (repo *MysqlOfertasRepository) Eliminar(id int) error {
	query := "DELETE FROM ofertas WHERE id_oferta = ?"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if numRows == 0 {
		return errors.New("no se encontró la oferta para eliminar")
	}
	return nil
}

func (repo *MysqlOfertasRepository) MostrarPorID(id int) ([]*entities.Ofertas, error) {
	rows, err := repo.db.Query("SELECT * FROM ofertas WHERE id_oferta = ?", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ofertas []*entities.Ofertas

	for rows.Next() {

		oferts := &entities.Ofertas{}
		if err := rows.Scan(
			&oferts.Id_oferta,
			&oferts.Nombre,
			&oferts.Descripcion,
			&oferts.Fecha_inicio,
			&oferts.Fecha_fin,
		); err != nil {
			return nil, err
		}

		ofertas = append(ofertas, oferts)
	}

	if len(ofertas) == 0 {
		return nil, fmt.Errorf("no se encontro la oferta")
	}
	return ofertas, nil
}


// Verifica si hay nuevas ofertas activas (actuales)
func (repo *MysqlOfertasRepository) VerificarNuevasOfertas() ([]*entities.Ofertas, error) {
	query := `
		SELECT o.id_oferta, o.nombre, o.descripcion, o.fecha_inicio, o.fecha_fin
		FROM ofertas o
		JOIN productos_ofertas po ON o.id_oferta = po.id_oferta
		WHERE NOW() BETWEEN o.fecha_inicio AND o.fecha_fin
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ofertas []*entities.Ofertas
	for rows.Next() {
		oferta := &entities.Ofertas{}
		if err := rows.Scan(&oferta.Id_oferta, &oferta.Nombre, &oferta.Descripcion, &oferta.Fecha_inicio, &oferta.Fecha_fin); err != nil {
			return nil, err
		}
		ofertas = append(ofertas, oferta)
	}

	if len(ofertas) == 0 {
		return nil, errors.New("no hay nuevas ofertas activas")
	}

	return ofertas, nil
}

// Verifica si alguna oferta ha caducado
func (repo *MysqlOfertasRepository) VerificarOfertasCaducadas() ([]*entities.Ofertas, error) {
	query := `
		SELECT id_oferta, nombre, descripcion, fecha_inicio, fecha_fin
		FROM ofertas
		WHERE fecha_fin < NOW()
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ofertas []*entities.Ofertas
	for rows.Next() {
		oferta := &entities.Ofertas{}
		if err := rows.Scan(&oferta.Id_oferta, &oferta.Nombre, &oferta.Descripcion, &oferta.Fecha_inicio, &oferta.Fecha_fin); err != nil {
			return nil, err
		}
		ofertas = append(ofertas, oferta)
	}

	if len(ofertas) == 0 {
		return nil, errors.New("no hay ofertas caducadas")
	}

	return ofertas, nil
}