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
	Name                                      string `sql:"type:varchar(32);NOT NULL"`
	Email                                     string `sql:"type:varchar(64);unique;NOT NULL"`
	Mobile                                    string `sql:"type:char(10);unique;NOT NULL"`
	Role, DesiredCity                         string `sql:"type:varchar(16)"`
	Education, Company, ComputerCertification string `sql:"type:varchar(64)"`
	HasSystem                                 string `sql:"type:varchar(8);DEFAULT:NULL"`
	LastActive                                time.Time
	LocalityID                                uint `fk:"localities(id)"`
	Age, CurrentSalary, MaxHEScore            uint
	MinHEScode, Experience, TypingSpeed       uint
	ITI, MsExcel, Passport, Adharcard         bool
}

// FKContactID handels foreign key contact id in corresponding tables
type FKContactID struct {
	ContactID uint `fk:"contacts(id)"`
}

// ContactLanguage stores relation b/w contact and job_type
type ContactLanguage struct {
	FKContactID `sql:"NOT NULL"`
	LanguageID  uint `unique:"unq_contact_language(contact_id,language_id)" fk:"languages(id)" sql:"NOT NULL"`
}

// ContactProof stores the proof relation with contact
type ContactProof struct {
	FKContactID
	ProofID uint `unique:"unq_contact_proof(contact_id,proof_id)" fk:"proofs(id)" sql:"NOT NULL"`
}

// ContactJobType stores relation b/w contact and job_type
type ContactJobType struct {
	FKContactID
	JobTypeID uint `sql:"NOT NULL" unique:"unq_contact_job_type(contact_id,job_type_id)" fk:"job_types(id)"`
}

// Proof stores proof name
type Proof struct {
	PKey
	Name string `sql:"varchar(16);unique;NOT NULL"`
}

// Language module stores the languages
type Language struct {
	PKey
	Language string `sql:"type:varchar(8);unique;NOT NULL"`
}

// JobType model stores all job type available
type JobType struct {
	PKey
	Name string `sql:"type:varchar(16);unique;NOT NULL"`
}

// Locality model stores Locality details of the contact
type Locality struct {
	PKey
	City, Locality string `sql:"type:varchar(16);unique;NOT NULL"`
}
