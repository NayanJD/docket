package models

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

func GenerateSecureToken(length int) string {
	rand.Seed(time.Now().UnixNano())
    
	b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}


type ClientStore struct {
	db                *gorm.DB
	tableName         string
	initTableDisabled bool
	maxLifetime       time.Duration
	maxOpenConns      int
	maxIdleConns      int
}

// ClientStoreItem data item
type ClientStoreItem struct {
	gorm.Model
	ID     string 
	Name   string	`gorm:"not null;unique"`
	Secret string	`gorm:"not null"`
	Domain string 
	Data   string
}

func (u ClientStoreItem) TableName() string{
	return "oauth2_clients"
}

func (s *ClientStoreItem) String() string {
	return fmt.Sprintf("ClientStoreItem ID: %v", s.ID)
}

func (c * ClientStoreItem) BeforeCreate(tx *gorm.DB)  (err error) {
	c.ID = GenerateSecureToken(8)
	c.Secret = GenerateSecureToken(16)

	return
}

// NewClientStore creates xorm mysql store instance
func NewClientStore(db *gorm.DB, options ...ClientStoreOption) (*ClientStore, error) {
	store := &ClientStore{
		// db:           getDB(),
		// tableName:    "oauth2_client",
		// maxLifetime:  time.Hour * 2,
		// maxOpenConns: 50,
		// maxIdleConns: 25,
	}

	for _, o := range options {
		o(store)
	}

	// var err error
	// if !store.initTableDisabled {
	// 	err = store.initTable()
	// }

	// if err != nil {
	// 	return store, err
	// }

	// store.db.SetMaxOpenConns(store.maxOpenConns)
	// store.db.SetMaxIdleConns(store.maxIdleConns)
	// store.db.SetConnMaxLifetime(store.maxLifetime)

	return store, nil
}

// func (s *ClientStore) initTable() error {
// 	err := s.db.AutoMigrate(&ClientStoreItem{})

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (s *ClientStore) toClientInfo(data string) (oauth2.ClientInfo, error) {
	var cm models.Client
	err := jsoniter.Unmarshal([]byte(data), &cm)
	return &cm, err
}

// GetByID retrieves and returns client information by id
func (s *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	if id == "" {
		return nil, nil
	}

	clientStoreItem := ClientStoreItem{}

	if err := GetDB().First(&clientStoreItem, "id = ?", id).Error; err != nil {
		return nil, nil;
	}

	log.Info().Msg(clientStoreItem.String())

	cm := models.Client{ID: clientStoreItem.ID, Secret: clientStoreItem.Secret, Domain: clientStoreItem.Domain, UserID: clientStoreItem.Data}

	return &cm, nil
}

// Create creates and stores the new client information
func (s *ClientStore) Create(info oauth2.ClientInfo) error {

	data, err := jsoniter.Marshal(info)
	if err != nil {
		return err
	}

	clientStoreItem := ClientStoreItem{ID: info.GetID(), Secret: info.GetSecret(), Domain: info.GetDomain(), Data: string(data)}

	GetDB().Create(&clientStoreItem)

	return nil
}

// ClientStoreOption is the configuration options type for client store
type ClientStoreOption func(s *ClientStore)