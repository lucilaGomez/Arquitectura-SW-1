package services

import (
	"proyecto/dtos"
	"proyecto/initializers"
	"testing"

	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateHotel(t *testing.T) {
	// Crear base de datos mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Configurar GORM con la base de datos mock
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open connection: %v", err)
	}

	initializers.DB = gormDB

	// Crear un objeto HotelDto
	hotelDto := dtos.HotelDto{
		Name:        "Sample Hotel",
		Description: "Sample description",
		Address:     "Sample Address",
		City:        "Sample City",
		Country:     "Sample Country",
		Amenities:   []string{"Sample amenity"},
		Photos:      []string{"Sample photo"},
	}

	// Definir expectativas
	// mock.ExpectBegin()
	// mock.ExpectExec("SELECT count(*) FROM `hotels` WHERE address = ? AND `hotels`.`deleted_at` IS NULL").WillReturnResult(sqlmock.NewResult(0, 1))
	// mock.ExpectExec("INSERT INTO hotels").WillReturnResult(sqlmock.NewResult(1, 1))
	//	mock.ExpectCommit()

	mock.ExpectQuery("SELECT count(*) FROM `hotels` WHERE address = ? AND `hotels`.`deleted_at` IS NULL").
		WithArgs(hotelDto.Address).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	mock.ExpectQuery("SELECT count(*) FROM `hotels` WHERE name != ? AND address = ? AND city = ? AND country = ? AND `hotels`.`deleted_at` IS NULL").
		WithArgs(hotelDto.Name, hotelDto.Address, hotelDto.City, hotelDto.Country).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	columns := []string{"id", "name"}

	mock.ExpectQuery("SELECT * FROM `amenities` WHERE name = ? AND `amenities`.`deleted_at` IS NULL ORDER BY `amenities`.`id` LIMIT ?").
		WithArgs(hotelDto.Amenities[0], 1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 1))

	// mock.ExpectBegin()
	// mock.ExpectExec( "INSERT INTO `hotels`(`created_at`,`updated_at`,`deleted_at`,`name`,`description`,`address`,`city`,`country`)").
	// 	WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), hotelDto.Name, hotelDto.Description, hotelDto.Address, hotelDto.City, hotelDto.Country).
	// 	WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectBegin()

	columns = []string{"created_at", "updated_at", "deleted_at", "name", "description", "address", "city", "country"}
	created_at := time.Now()
	updated_at := time.Now()
	deleted_at := time.Now()
	name := ""
	description := ""
	address := ""
	city := ""
	country := ""

	mock.ExpectQuery("INSERT INTO `hotels`(`created_at`,`updated_at`,`deleted_at`,`name`,`description`,`address`,`city`,`country`)"+
		" VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), // These are for created_at, updated_at, deleted_at
			hotelDto.Name, hotelDto.Description, hotelDto.Address, hotelDto.City, hotelDto.Country).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(created_at, updated_at, deleted_at, name, description, address, city, country))

	mock.ExpectCommit()

	// Llamar a la función CreateHotel
	result, err := CreateHotel(hotelDto)

	// Verificar resultados
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Sample Hotel", result.Name)
	assert.Equal(t, "Sample description", result.Description)
	assert.Equal(t, "Sample Address", result.Address)
	assert.Equal(t, "Sample City", result.City)
	assert.Equal(t, "Sample Country", result.Country)

	// Verificar que se cumplan todas las expectativas
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}

}
