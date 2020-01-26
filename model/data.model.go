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
	Name                                            string `sql:"type:varchar(32);NOT NULL"`
	Email                                           string `sql:"type:varchar(64);unique;NOT NULL"`
	Mobile                                          string `sql:"type:char(10);unique;NOT NULL"`
	Role, HasSystem                                 string `sql:"type:varchar(32);DEFAULT:NULL"`
	DesiredCity, Education, Company, Specialization string `sql:"type:varchar(128);DEFAULT:NULL"`
	TypingSpeed                                     string `sql:"type:varchar(16);DEFAULT:NULL"`
	LastActive                                      time.Time
	LocalityID                                      uint `fk:"localities(id)"`
	Age, CurrentSalary, MaxHEScore                  uint `sql:"DEFAULT:NULL"`
	MinHEScode, Experience                          uint `sql:"DEFAULT:NULL"`
	ITI, MsExcel, Passport, CriminalCase, Adharcard bool `sql:"DEFAULT:NULL"`
}

// FKContactID handles foreign key contact id in corresponding tables
type FKContactID struct {
	ContactID uint `fk:"contacts(id)" sql:"NOT NULL"`
}

// ContactLanguage stores relation b/w contact and job_type
type ContactLanguage struct {
	FKContactID
	LanguageID uint `unique:"unq_contact_language(contact_id,language_id)" fk:"languages(id)" sql:"NOT NULL"`
}

// ContactComputerCertificate stores relation b/w contact and certificate
type ContactComputerCertificate struct {
	FKContactID
	CertificationID uint `unique:"unq_contact_certificate(contact_id,certification_id)" fk:"computer_certifications(id)" sql:"NOT NULL"`
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

// ComputerCertification stores the computer certifaction
type ComputerCertification struct {
	PKey
	Name string
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
	City     string `unique:"uniq_locality(city,locality)"`
	Locality string `sql:"type:varchar(32);NOT NULL"`
}
