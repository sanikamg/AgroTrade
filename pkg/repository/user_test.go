package repository

import (
	"context"
	"errors"
	"golang_project_ecommerce/pkg/domain"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUpdateUserDetails(t *testing.T) {

	//create mock database connection
	mockDB, mock, err := sqlmock.New()
	if assert.NoError(t, err) {
		log.Println("Mock SQL created successfully")
	}
	defer mockDB.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB, DriverName: "postgres"}), &gorm.Config{})
	if assert.NoError(t, err) {
		log.Println("Mock SQL connected with GORM successfully")
	}
	//creating new user instance with mock database
	newUserRepositoryDB := NewUserRepository(db)

	//testcase1 with argument
	user := domain.Users{
		User_Id:     1,
		Username:    "vishnu",
		Name:        "vishnu",
		Phone:       "6282099388",
		Email:       "vishnup@gmail.com",
		BlockStatus: false,
		Password:    "vishnu",

		Verification: true,
	}

	updateDetails := domain.Users{
		Username: "vishnu",
		Name:     "vishnu",
		Phone:    "6282099388",
		Email:    "vishnup@gmail.com",
		Password: "vishnu",
	}

	//creating query for testcase
	expectedQuery := `update users set username=\$1,name=\$2,email=\$3,password=\$4 where phone=\$5`

	// Set up the mock expectation for the update query
	mock.ExpectQuery(expectedQuery).
		WithArgs(updateDetails.Username, updateDetails.Name, updateDetails.Email, updateDetails.Password, updateDetails.Phone).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "username", "name", "phone", "email", "block_status", "password", "verification"}).AddRow(1, "vishnu", "vishnu", "6282099388", "vishnup@gmail.com", false, "vishnu", true)) // Simulate a successful update

	applied, err := newUserRepositoryDB.UpdateUserDetails(context.Background(), updateDetails)

	assert.NoError(t, err)

	assert.Equal(t, user, applied)

	//err = mock.ExpectationsWereMet()

	if assert.NoError(t, err) {
		log.Println("Test1 Passed")
	}

	mock.ExpectQuery(expectedQuery).
		WithArgs(updateDetails.Username, updateDetails.Name, updateDetails.Email, updateDetails.Password, updateDetails.Phone).
		WillReturnError(errors.New("failed to complete user registration"))

	applied, err = newUserRepositoryDB.UpdateUserDetails(context.Background(), updateDetails)

	assert.Error(t, err)

	if assert.EqualError(t, err, "failed to complete user registration") {
		log.Println("Test 2 passed")
	}

}

//usermanagement

func TestFindUserById(t *testing.T) {
	//creating a mock database
	mockDB, mock, err := sqlmock.New()
	if assert.NoError(t, err) {
		log.Println("Mock SQL created successfully")
	}
	defer mockDB.Close()

	//connecting to form postgres
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB, DriverName: "postgres"}), &gorm.Config{})
	if assert.NoError(t, err) {
		log.Println("Mock SQL connected with GORM successfully")
	}

	//creating new user instance with mock database
	newUserRepositoryDB := NewUserRepository(db)

	//testcase1 with argument

	//expected output
	user := domain.Users{
		User_Id:     1,
		Username:    "vishnu",
		Name:        "vishnu",
		Phone:       "6282099388",
		Email:       "vishnup@gmail.com",
		BlockStatus: false,
		Password:    "vishnu",

		Verification: true,
	}

	//input
	userDetails := domain.Users{
		User_Id: 1,
	}

	//creating query for testcase
	expectedQuery := `select \* from users where user_id=\$1`

	// Set up the mock expectation for the update query
	mock.ExpectQuery(expectedQuery).
		WithArgs(userDetails.User_Id).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "username", "name", "phone", "email", "block_status", "password", "verification"}).AddRow(1, "vishnu", "vishnu", "6282099388", "vishnup@gmail.com", false, "vishnu", true)) // Simulate a successful update

	applied, err := newUserRepositoryDB.FindUserById(context.Background(), int(userDetails.User_Id))

	assert.NoError(t, err)

	assert.Equal(t, user, applied)

	if assert.NoError(t, err) {
		log.Println("Test1 Passed")
	}

}
