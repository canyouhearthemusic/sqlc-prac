// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Reservation struct {
	ID        int32            `db:"id" json:"id"`
	RoomID    string           `db:"room_id" json:"room_id"`
	StartTime pgtype.Timestamp `db:"start_time" json:"start_time"`
	EndTime   pgtype.Timestamp `db:"end_time" json:"end_time"`
}
