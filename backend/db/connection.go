package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	md "github.com/JuanMillan7818/FUT21/api_rest/models"
	"github.com/JuanMillan7818/FUT21/rest_fut21/models"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectedDB(nameDB string) *sql.DB {

	var dnsName string
	if nameDB == "" {
		dnsName = dns("")
	} else {
		dnsName = dns(nameDB)
	}
	db, err := sql.Open("mysql", dnsName)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	if nameDB == "" {
		inicializeDB(db)
	}
	return db
}

func inicializeDB(db *sql.DB) {
	query := getStatements()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}

	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer cancel()
	defer db.Close()
}

func SaveInDB(players []models.Player) {
	values := []interface{}{}
	var param string

	for _, item := range players {
		param += "(?,?,?,?,?), "
		values = append(
			values,
			item.FirstName,
			item.LastName,
			item.PositionFull,
			item.Natal.Name,
			item.Club.Name)
	}
	insertData(values, param[0:len(param)-2])
}

func insertData(players []interface{}, paramUnknown string) {
	db := ConnectedDB(os.Getenv("DBNAME"))
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sql := `INSERT INTO 
		Players (first_name, last_name, position, nationality, club)
		VALUES ` + paramUnknown

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(players...)
	if err != nil {
		panic(err)
	}

	tx.Commit()
	result.RowsAffected()

	defer db.Close()
}

func GetTeamPlayers(data *md.TeamRequest) md.TeamResponse {
	var list []md.Player
	db := ConnectedDB(os.Getenv("DBNAME"))

	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}

	sql := `SELECT * FROM 
		(SELECT * FROM Players FORCE INDEX(PRIMARY) WHERE LOWER(club)=?) 
		AS T1 GROUP BY first_name, last_name LIMIT ?,?`

	sql_aux := `SELECT COUNT(*) AS items FROM 
		(SELECT * FROM 
		(SELECT * FROM Players WHERE LOWER(club)=?) 
		AS T1 GROUP BY first_name, last_name) as T2;`

	stmt, err := db.Prepare(sql)
	stmt_aux, err_aux := db.Prepare(sql_aux)
	if err != nil || err_aux != nil {
		log.Panic(err)
	}

	data.Page--
	resul, err := stmt.Query(data.Name, (parseInt(os.Getenv("PAGINATION")) * data.Page), parseInt(os.Getenv("PAGINATION")))
	resul_aux, err_aux := stmt_aux.Query(data.Name)
	data.Page++
	if err != nil || err_aux != nil {
		log.Panic(err)
	}

	tx.Commit()

	for resul.Next() {
		var tmp md.Player
		if err := resul.Scan(
			&tmp.Id,
			&tmp.FirstName,
			&tmp.LastName,
			&tmp.PositionFull,
			&tmp.Natal,
			&tmp.Club); err != nil {
			log.Panic(err)
		}
		list = append(list, tmp)
	}

	var totalItemsAux string
	resul_aux.Next()
	if resul_aux.Scan(&totalItemsAux); err != nil {
		log.Panic(err)
	}
	totalItems, err := strconv.Atoi(totalItemsAux)
	if err != nil {
		log.Panic(err)
	}

	response := md.TeamResponse{
		Page:       data.Page,
		Items:      len(list),
		TotalItems: totalItems,
		TotalPages: int((math.Ceil(((float64)(totalItems) / float64(parseInt(os.Getenv("PAGINATION"))))))),
		Players:    list,
	}

	defer resul.Close()
	defer db.Close()
	return response
}

func GetNamesPlayers(search, order string, page int) md.TeamResponse {
	var list []md.Player
	var sql string
	var totalItemsAux string

	db := ConnectedDB(os.Getenv("DBNAME"))

	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}

	if strings.ToLower(order) == "desc" {
		sql = `SELECT * FROM Players FORCE INDEX(PRIMARY) 
		WHERE first_name LIKE ? OR last_name LIKE ? GROUP BY first_name desc, last_name desc
		LIMIT ?,?`
	} else {
		sql = `SELECT * FROM Players FORCE INDEX(PRIMARY) 
			WHERE first_name LIKE ? OR last_name LIKE ? GROUP BY first_name asc, last_name asc
			LIMIT ?,?`
	}

	sql_aux := `SELECT COUNT(*) AS items FROM 
		(SELECT * FROM Players WHERE first_name LIKE ? OR last_name LIKE ? GROUP BY first_name, last_name) as T2;`

	stmt, err := db.Prepare(sql)
	stmt_aux, err_aux := db.Prepare(sql_aux)
	if err != nil || err_aux != nil {
		log.Panic(err)
	}

	page--
	result, err := stmt.Query("%"+search+"%", "%"+search+"%", (parseInt(os.Getenv("PAGINATION")) * page), parseInt(os.Getenv("PAGINATION")))
	result_aux, err_aux := stmt_aux.Query("%"+search+"%", "%"+search+"%")
	if err != nil || err_aux != nil {
		log.Panic(err)
	}
	page++
	tx.Commit()

	for result.Next() {
		var tmp md.Player
		if err := result.Scan(
			&tmp.Id,
			&tmp.FirstName,
			&tmp.LastName,
			&tmp.PositionFull,
			&tmp.Natal,
			&tmp.Club); err != nil {
			log.Panic(err)
		}
		list = append(list, tmp)
	}

	result_aux.Next()
	if result_aux.Scan(&totalItemsAux); err != nil {
		log.Panic(err)
	}
	totalItems, err := strconv.Atoi(totalItemsAux)
	if err != nil {
		log.Panic(err)
	}

	if list == nil {
		list = []md.Player{}
	}
	response := md.TeamResponse{
		Page:       page,
		Items:      len(list),
		TotalItems: totalItems,
		TotalPages: int((math.Ceil(((float64)(totalItems) / float64(parseInt(os.Getenv("PAGINATION"))))))),
		Players:    list,
	}

	defer result.Close()
	defer db.Close()
	return response
}

func GetPlayersAll(page int) md.TeamResponse {
	var totalItemsAux string
	var list []md.Player

	db := ConnectedDB(os.Getenv("DBNAME"))

	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}

	sql := `SELECT * FROM Players FORCE INDEX(PRIMARY) GROUP BY first_name asc, last_name asc LIMIT ?,?`
	sql_aux := `SELECT COUNT(*) AS COUNT FROM Players GROUP BY first_name asc, last_name asc`

	stmt, err := db.Prepare(sql)
	stmt_aux, err_aux := db.Prepare(sql_aux)
	if err != nil || err_aux != nil {
		log.Panic(err)
	}

	page--
	result, err := stmt.Query((parseInt(os.Getenv("PAGINATION")) * page), parseInt(os.Getenv("PAGINATION")))
	page++
	result_aux, err_aux := stmt_aux.Query()
	if err != nil || err_aux != nil {
		log.Panic(err)
	}

	tx.Commit()

	for result.Next() {
		var tmp md.Player
		if err := result.Scan(
			&tmp.Id,
			&tmp.FirstName,
			&tmp.LastName,
			&tmp.PositionFull,
			&tmp.Natal,
			&tmp.Club); err != nil {
			log.Panic(err)
		}
		list = append(list, tmp)
	}

	result_aux.Next()
	if result_aux.Scan(&totalItemsAux); err != nil {
		log.Panic(err)
	}
	totalItems, err := strconv.Atoi(totalItemsAux)
	if err != nil {
		log.Panic(err)
	}

	if list == nil {
		list = []md.Player{}
	}
	response := md.TeamResponse{
		Page:       page,
		Items:      len(list),
		TotalItems: totalItems,
		TotalPages: int((math.Ceil(((float64)(totalItems) / float64(parseInt(os.Getenv("PAGINATION"))))))),
		Players:    list,
	}

	defer result.Close()
	defer db.Close()
	return response
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return n

}
func dns(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOSTNAME"), dbName)
}

func getStatements() string {
	return `CREATE DATABASE IF NOT EXISTS FUT21;
			USE FUT21;	
			CREATE TABLE IF NOT EXISTS Players (
				player_id INT NOT NULL AUTO_INCREMENT,
				first_name VARCHAR(100) NOT NULL,
				last_name VARCHAR(100) NOT NULL,
				position VARCHAR(100) NOT NULL,
				nationality VARCHAR(100) NOT NULL,
				club VARCHAR(100) NOT NULL,
				PRIMARY KEY (player_id)    
			);
			ALTER TABLE Players AUTO_INCREMENT = 1;`
}
