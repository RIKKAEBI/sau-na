package model

import (
	"database/sql/driver"
	"time"

	"braces.dev/errtrace"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type UUID uuid.UUID

func (UUID) GormDataType() string {
	return "binary(16)"
}
func (UUID) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "binary(16)"
}

func ParseUUID(id string) UUID {
	return UUID(uuid.Must(uuid.FromString(id)))
}

func (u *UUID) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	parseByte, err := uuid.FromBytes(bytes)
	*u = UUID(parseByte)
	return errtrace.Wrap(err)
}

func (u UUID) Value() (driver.Value, error) {
	byte, err := uuid.UUID(u).MarshalBinary()
	if err != nil {
		return nil, errtrace.Wrap(err)
	}
	return byte, nil
}

func (u UUID) String() string {
	return uuid.UUID(u).String()
}

// MarshalJSON -> convert to json string
func (u UUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(u)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

// UnmarshalJSON -> convert from json string
func (u *UUID) UnmarshalJSON(b []byte) error {
	s, err := uuid.FromBytes(b)
	*u = UUID(s)
	return errtrace.Wrap(err)
}

type ModelUUID struct {
	ID UUID `gorm:"type:binary(16)" json:"id"`
}

type ImmutableModel struct {
	ModelUUID
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime(6)"`
}

type Model struct {
	ModelUUID
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime(6)"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime(6)"`
}

type SoftModel struct {
	Model
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"type:datetime(6)"`
}

func (m *ModelUUID) BeforeCreate(db *gorm.DB) (err error) {
	uid := uuid.Must(uuid.NewV7())
	m.ID = UUID(uid)
	err = errtrace.Wrap(err)
	return
}
