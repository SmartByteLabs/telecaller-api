package controller

import (
	"io"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/princeparmar/telecaller-app/database"
	"github.com/princeparmar/telecaller-app/model"
	"github.com/princeparmar/telecaller-app/schema"
	"github.com/princeparmar/telecaller-app/utils"
	"github.com/rightjoin/fuel"
)

// UploaderService handles the contact upload part
type UploaderService struct {
	fuel.Service
	upload fuel.POST
}

// Upload read the contact avalable in recieved csv and insert in database.
func (UploaderService) Upload(ad fuel.Aide) (*schema.UploaderResponse, error) {
	dbo := database.DatabaseManager.GetORM()

	// parse file
	reader, err := utils.NewCSVReaderFromRequest(ad.Request)
	if err != nil {
		return nil, err
	}

	var lineNo uint
	out := new(schema.UploaderResponse)
	for {
		line, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		lineNo++

		trn := dbo.Begin()
		if err := storeContactE2E(trn, line); err != nil {
			trn.Rollback()
			out.AddError(lineNo, line.String("mobile"), err)
			continue
		}
		trn.Commit()
		out.Successfull(line.String("mobile"))
	}

	return out, nil
}

func storeContactE2E(trn *gorm.DB, line utils.Line) error {
	// store locality
	localityID, err := storeLocality(trn, line)
	if err != nil {
		return err
	}

	// save contact info
	contactID, err := storeContact(trn, line, localityID)
	if err != nil {
		return err
	}

	return storeRel(trn, line, contactID, storeProof, storeJobType, storeComputerCertificate)
}

func storeRel(dbo *gorm.DB, l utils.Line, contactID uint, fns ...func(*gorm.DB, utils.Line, uint) error) error {
	for _, fn := range fns {
		err := fn(dbo, l, contactID)
		if err != nil {
			return err
		}
	}

	return nil
}

func storeLocality(dbo *gorm.DB, line utils.Line) (uint, error) {
	locality := model.Locality{
		City:     line.String("city"),
		Locality: line.String("locality"),
	}

	err := dbo.FirstOrCreate(&locality, "city = ? AND locality = ?", locality.City, locality.Locality).Error
	return locality.ID, err
}

func storeContact(dbo *gorm.DB, line utils.Line, localityID uint) (uint, error) {
	contact := model.Contact{
		Name:          line.String("name"),
		Email:         line.String("email"),
		Mobile:        line.String("mobile"),
		Role:          line.String("role"),
		DesiredCity:   line.String("desiredcity"),
		Age:           line.Uint("age"),
		LastActive:    line.Time("lastactivity"),
		CurrentSalary: line.Uint("currentsalarypermonth"),
		MaxHEScore:    line.Uint("helloenglishscoremaximum"),
		MinHEScode:    line.Uint("helloenglishscoreminimum"),
		Education:     line.String("education"),
		Experience:    line.Uint("roleexperienceinyears"),
		MsExcel:       line.Bool("skilledinmicrosoftexcel"),
		Passport:      line.Bool("passport"),
		Adharcard:     line.Bool("aadhaarcard"),
		TypingSpeed:   utils.TypingSpeed(line.String("typingspeed")),
		ITI:           line.Bool("itidiploma"),
		Company:       line.String("company"),
		HasSystem:     line.String("doyouhavelaptopdesktop"),
		CriminalCase:  line.Bool("criminalcasescomplaints"),
		LocalityID:    localityID,
	}
	err := dbo.Save(&contact).Error
	return contact.ID, err
}

func storeProof(dbo *gorm.DB, line utils.Line, contactID uint) error {
	proofs := append(line.StringArr("addressproof"), line.StringArr("photoidproof")...)
	for _, p := range proofs {
		if utils.StringCleaner(p) == "" {
			continue
		}
		proof := model.Proof{
			Name: p,
		}

		err := dbo.FirstOrCreate(&proof, "name = ?", proof.Name).Error
		if err != nil {
			return err
		}

		err = utils.SaveWithDuplicateIgnore(dbo, &model.ContactProof{
			FKContactID: model.FKContactID{ContactID: contactID},
			ProofID:     proof.ID,
		})
		if err != nil {
			if mysqlErr, _ := err.(*mysql.MySQLError); mysqlErr != nil && mysqlErr.Number == 1064 {
				continue
			}
			return err
		}

	}

	return nil
}

func storeComputerCertificate(dbo *gorm.DB, line utils.Line, contactID uint) error {
	for _, p := range line.StringArr("computercertification") {
		if utils.StringCleaner(p) == "" {
			continue
		}
		certificate := model.ComputerCertification{
			Name: p,
		}

		err := dbo.FirstOrCreate(&certificate, "name = ?", certificate.Name).Error
		if err != nil {
			return err
		}

		err = utils.SaveWithDuplicateIgnore(dbo, &model.ContactComputerCertificate{
			FKContactID:     model.FKContactID{ContactID: contactID},
			CertificationID: certificate.ID,
		})
		if err != nil {
			if mysqlErr, _ := err.(*mysql.MySQLError); mysqlErr != nil && mysqlErr.Number == 1064 {
				continue
			}
			return err
		}

	}
	return nil
}

func storeJobType(dbo *gorm.DB, line utils.Line, contactID uint) error {
	for _, p := range line.StringArr("jobtype") {
		if utils.StringCleaner(p) == "" {
			continue
		}
		jobType := model.JobType{
			Name: p,
		}

		err := dbo.FirstOrCreate(&jobType, "name = ?", jobType.Name).Error
		if err != nil {
			return err
		}

		err = utils.SaveWithDuplicateIgnore(dbo, &model.ContactJobType{
			FKContactID: model.FKContactID{ContactID: contactID},
			JobTypeID:   jobType.ID,
		})
		if err != nil {
			return err
		}

	}
	return nil
}

func storeLanguage(dbo *gorm.DB, line utils.Line, contactID uint) error {
	for _, p := range line.StringArr("languagesknown") {
		if utils.StringCleaner(p) == "" {
			continue
		}
		language := model.Language{
			Language: p,
		}

		err := dbo.FirstOrCreate(&language, "language = ?", language.Language).Error
		if err != nil {
			return err
		}

		err = utils.SaveWithDuplicateIgnore(dbo, &model.ContactLanguage{
			FKContactID: model.FKContactID{ContactID: contactID},
			LanguageID:  language.ID,
		})
		if err != nil {
			if mysqlErr, _ := err.(*mysql.MySQLError); mysqlErr != nil && mysqlErr.Number == 1064 {
				continue
			}
			return err
		}

	}
	return nil
}
