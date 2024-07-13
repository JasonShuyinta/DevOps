package services

import (
	"database/sql"
	"fmt"
	"log"
	"my-go-api/pkg/models"
	"strconv"
	"sync"
	"time"
)

func CreateTableIfNotExists(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS public.album (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		duration VARCHAR(255),  -- Duration is a string here
		artist VARCHAR(255),
		label VARCHAR(255),
		creation_timestamp TIMESTAMPTZ,
		update_timestamp TIMESTAMPTZ
	)`
	_, err := db.Exec(query)
	return err
}

func GetAllAlbums(client *sql.DB) ([]models.Albums, error) {
	rows, err := client.Query("SELECT * FROM album")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []models.Albums
	for rows.Next() {
		var album models.Albums
		if err := rows.Scan(&album.ID, &album.Title, &album.Description, &album.Duration,
			&album.Artist, &album.Label, &album.CreationTimestamp, &album.UpdateTimestamp); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return albums, nil
}

func SaveAlbum(client *sql.DB, album *models.Albums) error {
	query := `
		INSERT INTO album(title, description, duration, artist, label, creation_timestamp, update_timestamp)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	err := client.QueryRow(
		query,
		album.Title,
		album.Description,
		album.Duration,
		album.Artist,
		album.Label,
		album.CreationTimestamp,
		album.UpdateTimestamp,
	).Scan(&album.ID)

	return err
}

func GetAlbumByID(client *sql.DB, id int) (*models.Albums, error) {
	var album models.Albums
	query := `SELECT * FROM album WHERE id = $1`
	row := client.QueryRow(query, id)
	err := row.Scan(&album.ID, &album.Title, &album.Description, &album.Duration, &album.Artist,
		&album.Label, &album.CreationTimestamp, &album.UpdateTimestamp)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("album with id %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &album, nil
}

func DeleteAlbumByID(client *sql.DB, id int) error {
	query := `DELETE FROM album WHERE id = $1`

	result, err := client.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no album found with id %d", id)
	}

	return nil
}

func CreateMultipleAlbums(client *sql.DB, n int) error {
	for i := 1; i < n; i++ {
		album := &models.Albums{
			Title:             "Album" + strconv.Itoa(i),
			Description:       "Description of Album " + strconv.Itoa(i),
			Duration:          "00:45:00", // Example duration
			Artist:            "Artist" + strconv.Itoa(i),
			Label:             "Label" + strconv.Itoa(i),
			CreationTimestamp: time.Now(),
			UpdateTimestamp:   time.Now(),
		}

		err := SaveAlbum(client, album)
		if err != nil {
			log.Printf("Error saving album %d: %v\n", i, err)
			return err
		}
	}
	return nil
}

func SaveAlbumsEfficiently(client *sql.DB, albums []models.Albums, wg *sync.WaitGroup, errChan chan error) {
	// Not so efficient
	defer wg.Done()

	if len(albums) == 0 {
		return
	}

	query := `
		INSERT INTO album(title, description, duration, artist, label, creation_timestamp, update_timestamp)
		VALUES `
	values := ""
	args := []interface{}{}
	for i, album := range albums {
		idx := i * 7
		values += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d),", idx+1, idx+2, idx+3, idx+4, idx+5, idx+6, idx+7)
		args = append(args, album.Title, album.Description, album.Duration, album.Artist, album.Label, album.CreationTimestamp, album.UpdateTimestamp)
	}
	query += values[:len(values)-1]

	_, err := client.Exec(query, args...)
	if err != nil {
		errChan <- err
	}
}

// CreateMultipleAlbumsEfficient creates n albums with incrementing titles using batch insert.
func CreateMultipleAlbumsEfficient(client *sql.DB, n int) error {
	albums := make([]models.Albums, n)
	batchSize := n / 5
	for i := 0; i < n; i++ {
		albums[i] = models.Albums{
			Title:             "Album" + strconv.Itoa(i+1),
			Description:       "Description of Album " + strconv.Itoa(i+1),
			Duration:          "00:45:00",
			Artist:            "Artist" + strconv.Itoa(i+1),
			Label:             "Label" + strconv.Itoa(i+1),
			CreationTimestamp: time.Now(),
			UpdateTimestamp:   time.Now(),
		}
	}

	var wg sync.WaitGroup
	errChan := make(chan error, n/batchSize+1)

	for i := 0; i < n; i += batchSize {
		end := i + batchSize
		if end > n {
			end = n
		}
		wg.Add(1)
		go SaveAlbumsEfficiently(client, albums[i:end], &wg, errChan)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			log.Printf("Error saving albums: %v\n", err)
			return err
		}
	}
	return nil
}

func UpdateAlbum(client *sql.DB, album *models.Albums) error {
	query := `UPDATE album
	SET title=$1, description = $2, duration = $3, artist  = $4, label = $5
	WHERE id = $6`

	_, err := client.Exec(query, album.Title, album.Description, album.Duration, album.Artist, album.Label, album.ID)

	return err
}
