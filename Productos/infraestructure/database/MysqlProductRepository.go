package database

import (
	"APICRUD/Productos/domain/entities"
	"APICRUD/Productos/domain/repositories"
	"database/sql"
	"fmt"
)

type MySQLProductoRepository struct {
	db *sql.DB
}

func NewMysqlProductoRepository(db *sql.DB) repositories.ProductoRepositories {
	return &MySQLProductoRepository{db: db}
}

func (r *MySQLProductoRepository) MostrarProductos() ([]*entities.Producto, error) {
	rows, err := r.db.Query("SELECT id_producto, nombre, descripcion, precio, stock, imagen_url, id_categoria, fecha_creacion FROM productos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []*entities.Producto
	for rows.Next() {
		prod := &entities.Producto{}
		if err := rows.Scan(
			&prod.Id_producto,
			&prod.Nombre,
			&prod.Descripcion,
			&prod.Precio,
			&prod.Stock,
			&prod.Imagen,
			&prod.Id_categoria,
			&prod.Fecha_creacion,
		); err != nil {
			return nil, err
		}
		productos = append(productos, prod)
	}
	return productos, nil
}

func (r *MySQLProductoRepository) AgregarProducto(producto *entities.Producto) error {
	_, err := r.db.Exec("INSERT INTO productos(nombre, descripcion, precio, stock, imagen_url, id_categoria, fecha_creacion) VALUES (?, ?, ?, ?, ?, ?, ?)",
		producto.Nombre, producto.Descripcion, producto.Precio, producto.Stock, producto.Imagen, producto.Id_categoria, producto.Fecha_creacion)
	return err
}

func (r *MySQLProductoRepository) ActualizarProducto(producto *entities.Producto) error {
	_, err := r.db.Exec("UPDATE productos SET nombre= ?, descripcion=?, precio=?, stock=?, imagen_url=?, id_categoria=?, fecha_creacion=? WHERE id_producto=?",
		producto.Nombre, producto.Descripcion, producto.Precio, producto.Stock, producto.Imagen, producto.Id_categoria, producto.Fecha_creacion, producto.Id_producto)
	return err
}

func (r *MySQLProductoRepository) EliminarProducto(id int) error {
	_, err := r.db.Exec("DELETE FROM productos WHERE id_producto = ?", id)
	return err
}

func (r *MySQLProductoRepository) BuscarPorID(id int) ([]*entities.Producto, error) {
	rows, err := r.db.Query("SELECT * FROM productos WHERE id_producto = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []*entities.Producto
	for rows.Next() {
		prod := &entities.Producto{}
		if err := rows.Scan(
			&prod.Id_producto,
			&prod.Nombre,
			&prod.Descripcion,
			&prod.Precio,
			&prod.Stock,
			&prod.Imagen,
			&prod.Id_categoria,
			&prod.Fecha_creacion,
		); err != nil {
			return nil, err
		}
		productos = append(productos, prod)
	}

	if len(productos) == 0 {
		return nil, fmt.Errorf("no se encontró el producto con ID %d", id)
	}

	return productos, nil
}
