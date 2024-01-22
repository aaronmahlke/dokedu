package pdf

import (
	"context"
	"example/internal/db"
)

type Competence struct {
	Name         string
	Level        int
	Color        string
	IsCompetence bool
	Competences  []Competence
}

type CompetencesTemplateData struct {
	OrganisationName    string
	OrganisationAddress string
	OrganisationLogoURL string
	StudentName         string
	StudentBirthday     string
	SchoolYear          string
	Competences         []Competence
}

func (g *Generator) CompetencesReportData(report db.Report) (*CompetencesTemplateData, error) {
	ctx := context.Background()

	var data CompetencesTemplateData

	var student db.User
	err := g.cfg.DB.NewSelect().Model(&student).Where("id = ?", report.StudentUserID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var organisation db.Organisation
	err = g.cfg.DB.NewSelect().Model(&organisation).Where("id = ?", report.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	data.OrganisationName = organisation.Name
	data.OrganisationAddress = organisation.Address

	var competences []db.Competence
	err = g.cfg.DB.NewSelect().
		Model(&competences).
		Where("organisation_id = ?", report.OrganisationID).
		Order("sort_order").Scan(ctx)
	if err != nil {
		return nil, err
	}

	var userCompetences []db.UserCompetence
	err = g.cfg.DB.NewSelect().
		Model(&userCompetences).
		Where("user_id = ?", report.StudentUserID).
		Where("organisation_id = ?", report.OrganisationID).
		Where("created_at >= ?", report.From).
		Where("created_at <= (DATE ? + 1)", report.To).
		Order("created_at DESC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	// create a map with key competence.ParentID and value []db.Competence
	var competencesMap = make(map[string][]db.Competence)
	for _, c := range competences {
		competencesMap[c.CompetenceID.String] = append(competencesMap[c.CompetenceID.String], c)
	}

	// create a map with key userCompetence.CompetenceID and value []db.UserCompetence
	var userCompetencesMap = make(map[string][]db.UserCompetence)
	for _, uc := range userCompetences {
		userCompetencesMap[uc.CompetenceID] = append(userCompetencesMap[uc.CompetenceID], uc)
	}

	subjects := Subjects(competences)

	for i := range subjects {
		if data.Competences == nil {
			data.Competences = make([]Competence, len(subjects))
		}
		data.Competences[i].Name = subjects[i].Name

		data.Competences[i].Color = subjects[i].Color.String

		if len(data.Competences[i].Color) == 0 {
			data.Competences[i].Color = "stone"
		}

		competences, err := g.populateCompetences(report, subjects[i], competencesMap, userCompetencesMap, data.Competences[i].Color)
		if err != nil {
			return nil, err
		}

		if data.Competences[i].Competences == nil {
			data.Competences[i].Competences = make([]Competence, len(competences))
		}
		data.Competences[i].Competences = competences
	}

	return &data, nil
}

func Subjects(competences []db.Competence) []db.Competence {
	// find all subjects by selecting the ones with no parent
	var subjects []db.Competence
	for _, c := range competences {
		if c.CompetenceID.String == "" {
			subjects = append(subjects, c)
		}
	}

	// sort subjects by sort_order
	for i := range subjects {
		for j := range subjects {
			if subjects[i].SortOrder < subjects[j].SortOrder {
				subjects[i], subjects[j] = subjects[j], subjects[i]
			}
		}
	}

	return subjects
}

func GetCompetenceByParent(competences map[string][]db.Competence, parent db.Competence) []db.Competence {
	// find all competences with the given parent
	var result []db.Competence
	for _, c := range competences[parent.ID] {
		result = append(result, c)
	}

	// sort competences by sort_order
	for i := range result {
		for j := range result {
			if result[i].SortOrder < result[j].SortOrder {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}

func GetUserCompetencesByCompetenceId(userCompetences map[string][]db.UserCompetence, competence db.Competence) []db.UserCompetence {
	var result []db.UserCompetence
	for _, uc := range userCompetences[competence.ID] {
		result = append(result, uc)
	}

	return result
}

func (g *Generator) populateCompetences(report db.Report, parent db.Competence, competencesData map[string][]db.Competence, userCompetences map[string][]db.UserCompetence, color string) ([]Competence, error) {
	var data []Competence

	competences := GetCompetenceByParent(competencesData, parent)

	for i := range competences {
		if data == nil {
			data = make([]Competence, len(competences))
		}
		if data[i].Competences == nil {
			data[i].Competences = make([]Competence, len(competences))
		}

		data[i].Color = color

		data[i].Name = competences[i].Name
		if competences[i].CompetenceType == db.CompetenceTypeCompetence {
			data[i].IsCompetence = true
		}

		userLevels := GetUserCompetencesByCompetenceId(userCompetences, competences[i])
		if len(userLevels) == 0 {
			data[i].Level = 0
		} else {
			data[i].Level = userLevels[0].Level
		}

		for i := range competences {
			competences, err := g.populateCompetences(report, competences[i], competencesData, userCompetences, color)
			if err != nil {
				return nil, err
			}

			data[i].Competences = competences
		}
	}

	return data, nil
}
