package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/markovk1n/spoty/models"
)

type AlbumRepository struct {
	db *sqlx.DB
}

func NewAlbumRepository(db *sqlx.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (t *AlbumRepository) CreateCommentForAlbum(userID int, input models.AlbumComment) error {
	tx, err := t.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var CID int
	err = tx.Get(&CID, "insert into comments (comment,user_id)values($1,$2) returning id", input.Comment, userID)
	if err != nil {
		return err
	}
	_, err = tx.Exec("insert into album_comments(album_id,comment_id)values($1,$2)", input.AlbumID, CID)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (t *AlbumRepository) GetAlbumsComm(albumId string) ([]models.AlbumCommentResp, error) {

	query := `
		SELECT u.username, c.comment
		FROM album_comments ac
		JOIN comments c ON ac.comment_id = c.id
		JOIN users u ON c.user_id = u.id
		WHERE ac.album_id = $1
	`

	rows, err := t.db.Query(query, albumId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var albumComments []models.AlbumCommentResp

	for rows.Next() {
		var username, comment string
		err := rows.Scan(&username, &comment)
		if err != nil {
			log.Fatal(err)
		}
		albumComment := models.AlbumCommentResp{
			Username: username,
			Comment:  comment,
		}
		albumComments = append(albumComments, albumComment)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return albumComments, nil
}
