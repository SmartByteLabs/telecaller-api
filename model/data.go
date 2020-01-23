package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// PKey primary key value
type PKey struct {
	ID uint `sql:"id;AUTO_INCREMENTED;PRIMARY_KEY"`
}

// Contact model to store the contact data.
type Contact struct {
	gorm.Model
	Name                  string `sql:"type:varchar(32)"`
	Email                 string `sql:"type:varchar(64);unique"`
	Mobile                string `sql:"type:char(10);unique"`
	Role                  string `sql:"type:varchar(16)"`
	DesiredCity           string `sql:"type:varchar(16)"`
	Education             string `sql:"type:varchar(64)"`
	Company               string `sql:"type:varchar(64)"`
	HasSystem             string `sql:"enum('laptop','desktop');DEFAULT:null"`
	ComputerCertification string `sql:"type:varchar(64)"`
	LastActive            time.Time
	LocalityID            uint `sql:""`
	Age                   uint
	CurrentSalary         uint
	MaxHEScore            uint
	MinHEScode            uint
	Experience            uint
	TypingSpeed           uint
	ITI                   bool
	MsExcel               bool
	Passport              bool
	Adharcard             bool
}

// ContactLanguage stores relation b/w contact and job_type
type ContactLanguage struct {
	ContactID  uint `sql:"unique('contact_id','language_id');foreignkey:contact('id')"`
	LanguageID uint
}

// ContactProof stores the proof relation with contact
type ContactProof struct {
	ContactID uint `sql:"uniq_proof('conatct_id')"`
	ProofID   uint
}

// Proof stores proof name
type Proof struct {
	PKey
	Name string `sql:"varchar(16)"`
}

// Language module stores the languages
type Language struct {
	PKey
	Language string `gorm:"type:varchar(8)"`
}

// ContactJobType stores relation b/w contact and job_type
type ContactJobType struct {
	ContactID uint `sql:"unique('contact_id','job_type_id')"`
	JobTypeID uint
}

// JobType model stores all job type available
type JobType struct {
	PKey
	Name string `sql:"type:varchar(16)"`
}

// Locality model stores Locality details of the contact
type Locality struct {
	PKey
	City     string `sql:"type:varchar(16)"`
	Locality string `sql:"type:varchar(16)"`
}
