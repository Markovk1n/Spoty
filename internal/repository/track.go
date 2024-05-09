package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/markovk1n/spoty/models"
)

type TrackRepository struct {
	db *sqlx.DB
}

func NewTrackRepository(db *sqlx.DB) *TrackRepository {
	return &TrackRepository{db: db}
}

func (t *TrackRepository) CreateCommentForTrack(userID int, input models.TrackComment) error {
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
	_, err = tx.Exec("insert into track_comments(track_id,comment_id)values($1,$2)", input.TrackID, CID)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (t *TrackRepository) GetTrackComm(trackId string) ([]models.TrackCommentResp, error) {
	query := `
	SELECT u.username, c.comment
	FROM track_comments ac
	JOIN comments c ON ac.comment_id = c.id
	JOIN users u ON c.user_id = u.id
	WHERE ac.track_id = $1
`

	rows, err := t.db.Query(query, trackId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var trackComments []models.TrackCommentResp

	for rows.Next() {
		var username, comment string
		err := rows.Scan(&username, &comment)
		if err != nil {
			log.Fatal(err)
		}
		trackComment := models.TrackCommentResp{
			Username: username,
			Comment:  comment,
		}
		trackComments = append(trackComments, trackComment)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return trackComments, nil
}
