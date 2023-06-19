package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"example/pkg/db"
	"example/pkg/graph/model"
	"example/pkg/middleware"
	"fmt"
	"strings"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

// Type is the resolver for the type field.
func (r *competenceResolver) Type(ctx context.Context, obj *db.Competence) (db.CompetenceType, error) {
	return obj.CompetenceType, nil
}

// Color is the resolver for the color field.
func (r *competenceResolver) Color(ctx context.Context, obj *db.Competence) (string, error) {
	if obj.Color.Valid {
		return obj.Color.String, nil
	}

	return "", nil
}

// Parents is the resolver for the parents field.
func (r *competenceResolver) Parents(ctx context.Context, obj *db.Competence) ([]*db.Competence, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	query := `
WITH RECURSIVE competence_hierarchy AS (
    SELECT *
    FROM competences
    WHERE id = ? AND organisation_id = ?

    UNION ALL

    SELECT c.*
    FROM competences c
    INNER JOIN competence_hierarchy ch ON c.id = ch.competence_id
	WHERE c.organisation_id = ?
)
SELECT *
FROM competence_hierarchy
WHERE id <> ?;
`

	// query without new lines
	q := strings.ReplaceAll(query, "\n", " ")

	var parents []*db.Competence
	err = r.DB.NewRaw(q, obj.ID, currentUser.OrganisationID, currentUser.OrganisationID, obj.ID).Scan(ctx, &parents)
	if err != nil {
		return nil, err
	}

	// reverse the order of the parents
	for i := len(parents)/2 - 1; i >= 0; i-- {
		opp := len(parents) - 1 - i
		parents[i], parents[opp] = parents[opp], parents[i]
	}

	return parents, nil
}

// Competences is the resolver for the competences field.
func (r *competenceResolver) Competences(ctx context.Context, obj *db.Competence, search *string) ([]*db.Competence, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var competences []*db.Competence
	query := r.DB.NewSelect().Model(&competences).Where("competence_id = ?", obj.ID).Where("organisation_id = ?", currentUser.OrganisationID)

	if search != nil && *search != "" {
		query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *search))
	}

	err = query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return competences, nil
}

// UserCompetences is the resolver for the userCompetences field.
func (r *competenceResolver) UserCompetences(ctx context.Context, obj *db.Competence, userID *string) ([]*db.UserCompetence, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	if userID == nil {
		return []*db.UserCompetence{}, nil
	}

	var userCompetences []*db.UserCompetence
	err = r.DB.NewSelect().
		Model(&userCompetences).
		Where("competence_id = ?", obj.ID).
		Where("user_id = ?", *userID).
		Where("organisation_id = ?", currentUser.OrganisationID).
		Order("created_at DESC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return userCompetences, nil
}

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.SignInPayload, error) {
	var user db.User
	err := r.DB.NewSelect().Model(&user).Where("email = ?", strings.ToLower(input.Email)).Scan(ctx)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	var organisation db.Organisation
	err = r.DB.NewSelect().Model(&organisation).Where("id = ?", user.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// user.Password is a sql.NullString, so we need to check if it is valid
	if !user.Password.Valid {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate a new token
	token, err := nanoid.New(32)
	if err != nil {
		return nil, errors.New("unable to generate a token")
	}

	// Save the token in the database
	session := db.Session{
		UserID: user.ID,
		Token:  token,
	}

	_, err = r.DB.NewInsert().Model(&session).Exec(ctx)
	if err != nil {
		return nil, errors.New("unable to generate a token")
	}

	return &model.SignInPayload{
		Token:       token,
		EnabledApps: organisation.EnabledApps,
	}, nil
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, input model.ResetPasswordInput) (*model.ResetPasswordPayload, error) {
	var user db.User
	err := r.DB.NewSelect().Model(&user).Where("recovery_token = ?", input.Token).Scan(ctx)
	if err != nil {
		return &model.ResetPasswordPayload{
			Success: false,
		}, nil
	}

	if user.RecoverySentAt.Add(time.Hour * 24).After(time.Now()) {
		return &model.ResetPasswordPayload{
			Success: false,
		}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return &model.ResetPasswordPayload{
			Success: false,
		}, nil
	}

	_, err = r.DB.NewUpdate().Model(&user).Set("password = ?", string(hashedPassword)).Set("recovery_token = NULL").Set("recovery_sent_at = NULL").Where("id = ?", user.ID).Exec(ctx)
	if err != nil {
		return &model.ResetPasswordPayload{
			Success: false,
		}, nil
	}

	return &model.ResetPasswordPayload{
		Success: true,
	}, nil
}

// ForgotPassword is the resolver for the forgotPassword field.
func (r *mutationResolver) ForgotPassword(ctx context.Context, input model.ForgotPasswordInput) (*model.ForgotPasswordPayload, error) {
	var user db.User
	err := r.DB.NewSelect().Model(&user).Where("email = ?", strings.ToLower(input.Email)).Scan(ctx)
	if err != nil {
		return &model.ForgotPasswordPayload{
			Success: false,
		}, nil
	}

	token := nanoid.Must(32)

	_, err = r.DB.NewUpdate().Model(&user).Set("recovery_token = ?", token).Set("recovery_sent_at = now()").Where("id = ?", user.ID).Exec(ctx)
	if err != nil {
		return &model.ForgotPasswordPayload{
			Success: false,
		}, nil
	}

	err = r.Mailer.SendPasswordReset(input.Email, user.FirstName, token)
	if err != nil {
		return &model.ForgotPasswordPayload{
			Success: false,
		}, nil
	}

	return &model.ForgotPasswordPayload{
		Success: true,
	}, nil
}

// SignOut is the resolver for the signOut field.
func (r *mutationResolver) SignOut(ctx context.Context) (bool, error) {
	currentUser := middleware.ForContext(ctx)
	if currentUser == nil {
		return false, errors.New("no user found in the context")
	}

	// TODO: suggestion: use a hard delete instead of a soft delete
	// TODO: suggestion: perhaps delete all sessions for the user?
	var session db.Session
	_, err := r.DB.NewUpdate().
		Model(&session).
		Set("deleted_at = now()").
		Where("user_id = ?", currentUser.ID).
		Where("token = ?", currentUser.Token).
		Exec(ctx)

	if err != nil {
		return false, err
	}

	return true, nil
}

// AcceptInvite is the resolver for the acceptInvite field.
func (r *mutationResolver) AcceptInvite(ctx context.Context, token string, input model.SignUpInput) (*model.SignInPayload, error) {
	panic(fmt.Errorf("not implemented: AcceptInvite - acceptInvite"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*db.User, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var organisation db.Organisation
	err = r.DB.NewSelect().Model(&organisation).Where("id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	// check if the email is in the allowed domains
	if isStringInArray(input.Email, organisation.AllowedDomains) {
		return nil, errors.New("email is not in the allowed domains (allowed domains: " + strings.Join(organisation.AllowedDomains, ", ") + ")")
	}

	// check if the email is already in the database
	//count, err := r.DB.middleware.GetUserByEmail(ctx, db.middleware.GetUserByEmailParams{
	//	OrganisationID: currentUser.OrganisationID,
	//	Email:          input.Email,
	//})
	var count int
	count, err = r.DB.NewSelect().Model(&db.User{}).Where("organisation_id = ?", currentUser.OrganisationID).Where("email = ?", input.Email).Count(ctx)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("email is already in the database")
	}

	// create a new user
	user := db.User{
		OrganisationID: currentUser.OrganisationID,
		Role:           input.Role,
		Email:          input.Email,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
	}

	// insert the user into the database
	err = r.DB.NewInsert().Model(&user).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*db.User, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	// update the user
	user := db.User{
		ID:             input.ID,
		OrganisationID: currentUser.OrganisationID,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
	}

	err = r.DB.NewUpdate().Model(&user).Where("id = ?", input.ID).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// InviteUser is the resolver for the inviteUser field.
func (r *mutationResolver) InviteUser(ctx context.Context, input model.CreateUserInput) (*db.User, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var organisation db.Organisation
	err = r.DB.NewSelect().Model(&organisation).Where("id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}
	if input.Email == "" {
		return nil, errors.New("email is required")
	}

	// check if the email is in the allowed domains
	if isStringInArray(input.Email, organisation.AllowedDomains) {
		return nil, errors.New("email is not in the allowed domains (allowed domains: " + strings.Join(organisation.AllowedDomains, ", ") + ")")
	}

	// check if the email is already in the database
	//count, err := r.DB.middleware.GetUserByEmail(ctx, db.middleware.GetUserByEmailParams{
	//	OrganisationID: currentUser.OrganisationID,
	//	Email:          input.Email,
	//})
	count, err := r.DB.NewSelect().Model(&db.User{}).Where("organisation_id = ?", currentUser.OrganisationID).Where("email = ?", input.Email).Count(ctx)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("email is already in the database")
	}

	// create a new user
	user := db.User{
		OrganisationID: currentUser.OrganisationID,
		Email:          input.Email,
		Role:           input.Role,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
	}

	// insert the user into the database
	err = r.DB.NewInsert().Model(&user).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ArchiveUser is the resolver for the archiveUser field.
func (r *mutationResolver) ArchiveUser(ctx context.Context, id string) (*db.User, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	// check whether the user is already archived
	count, err := r.DB.NewSelect().Model(&db.User{}).Where("id = ?", id).Where("organisation_id = ?", currentUser.OrganisationID).Where("deleted_at IS NOT NULL").Count(ctx)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("user is already archived")
	}

	user := &db.User{
		ID:             id,
		OrganisationID: currentUser.OrganisationID,
		DeletedAt: bun.NullTime{
			Time: time.Now(),
		},
	}

	// archive the user by setting the deleted_at field to the current time
	res, err := r.DB.NewUpdate().Model(user).Column("deleted_at").Where("id = ?", id).Where("organisation_id = ?", currentUser.OrganisationID).Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// CreateUserCompetence is the resolver for the createUserCompetence field.
func (r *mutationResolver) CreateUserCompetence(ctx context.Context, input model.CreateUserCompetenceInput) (*db.UserCompetence, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	if input.UserID == "" {
		return nil, errors.New("user id is required")
	}
	if input.CompetenceID == "" {
		return nil, errors.New("competence id is required")
	}

	userCompetence := db.UserCompetence{
		UserID:         input.UserID,
		Level:          input.Level,
		CreatedBy:      sql.NullString{String: currentUser.ID, Valid: true},
		CompetenceID:   input.CompetenceID,
		OrganisationID: currentUser.OrganisationID,
	}

	err = r.DB.NewInsert().Model(&userCompetence).Returning("*").Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &userCompetence, nil
}

// ArchiveUserCompetence is the resolver for the archiveUserCompetence field.
func (r *mutationResolver) ArchiveUserCompetence(ctx context.Context, id string) (*db.UserCompetence, error) {
	panic(fmt.Errorf("not implemented: ArchiveUserCompetence - archiveUserCompetence"))
}

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input model.CreateTagInput) (*db.Tag, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	// Check if tag with the same name already exists
	count, err := r.DB.NewSelect().Model(&db.Tag{}).Where("organisation_id = ?", currentUser.OrganisationID).Where("name = ?", input.Name).WhereAllWithDeleted().Count(ctx)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("Tag with the same name already exists")
	}

	// check if color is set
	color := input.Color
	if color == "" {
		color = "blue"
	}

	newTag := db.Tag{
		OrganisationID: currentUser.OrganisationID,
		Name:           input.Name,
		Color:          sql.NullString{String: color, Valid: true},
	}

	err = r.DB.NewInsert().Model(&newTag).Returning("*").Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &newTag, nil
}

// ArchiveTag is the resolver for the archiveTag field.
func (r *mutationResolver) ArchiveTag(ctx context.Context, id string) (*db.Tag, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	// set deleted_at field to the current time
	tag := db.Tag{
		ID:             id,
		OrganisationID: currentUser.OrganisationID,
		DeletedAt: bun.NullTime{
			Time: time.Now(),
		},
	}
	_, err = r.DB.NewUpdate().Model(&tag).Column("deleted_at").Where("id = ?", id).Where("organisation_id = ?", currentUser.OrganisationID).WhereAllWithDeleted().Returning("*").Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &tag, nil
}

// UpdateTag is the resolver for the updateTag field.
func (r *mutationResolver) UpdateTag(ctx context.Context, id string, input model.CreateTagInput) (*db.Tag, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var tag db.Tag
	tag.ID = id
	err = r.DB.NewSelect().
		Model(&tag).
		Where("organisation_id = ?", currentUser.OrganisationID).
		WherePK().
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	// update the tag
	tag.Name = input.Name

	color := input.Color
	if color == "" {
		color = "blue"
	}

	tag.Color = sql.NullString{String: color, Valid: true}
	_, err = r.DB.NewUpdate().Model(&tag).WherePK().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &tag, nil
}

// CreateReport is the resolver for the createReport field.
func (r *mutationResolver) CreateReport(ctx context.Context, input model.CreateReportInput) (*db.Report, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	report := db.Report{
		OrganisationID: currentUser.OrganisationID,
		UserID:         currentUser.ID,
		StudentUserID:  input.StudentUser,
		From:           input.From,
		To:             input.To,
		Format:         input.Format,
		FilterTags:     input.FilterTags,
		Kind:           input.Kind,
		Status:         "pending",
	}

	err = r.DB.NewInsert().Model(&report).Returning("*").Scan(ctx)
	if err != nil {
		return nil, err
	}

	// Call the report generation service
	err = r.ReportService.AddToQueue(report.ID)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

// UpdatePassword is the resolver for the updatePassword field.
func (r *mutationResolver) UpdatePassword(ctx context.Context, oldPassword string, newPassword string) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdatePassword - updatePassword"))
}

// UpdateCompetence is the resolver for the updateCompetence field.
func (r *mutationResolver) UpdateCompetence(ctx context.Context, input model.UpdateCompetenceInput) (*db.Competence, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var competence db.Competence
	competence.ID = input.ID
	err = r.DB.NewSelect().
		Model(&competence).
		Where("organisation_id = ?", currentUser.OrganisationID).
		WherePK().
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	if input.Color != nil {
		competence.Color = sql.NullString{String: *input.Color, Valid: true}
	}

	err = r.DB.NewUpdate().Model(&competence).WherePK().Returning("*").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &competence, nil
}

// Owner is the resolver for the owner field.
func (r *organisationResolver) Owner(ctx context.Context, obj *db.Organisation) (*db.User, error) {
	_, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var user db.User
	err = r.DB.NewSelect().Model(&user).Where("id = ?", obj.OwnerID).Where("organisation_id = ?", obj.ID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Organisation is the resolver for the organisation field.
func (r *queryResolver) Organisation(ctx context.Context) (*db.Organisation, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var organisation db.Organisation
	err = r.DB.NewSelect().Model(&organisation).Where("id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &organisation, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int, filter *model.UserFilterInput, search *string) (*model.UserConnection, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	// query the users
	var users []*db.User
	query := r.DB.NewSelect().Model(&users).Where("organisation_id = ?", currentUser.OrganisationID)

	if filter != nil {
		if filter.Role != nil && len(filter.Role) > 0 {
			query.Where("role IN (?)", bun.In(filter.Role))
		}
	}

	if search != nil && *search != "" {
		withoutSpace := strings.Replace(*search, " ", "", -1)
		// TODO: refactor this
		query.Where("first_name ILIKE ? OR last_name ILIKE ? OR first_name || last_name ILIKE ? OR last_name || first_name ILIKE ?", "%"+withoutSpace+"%", "%"+withoutSpace+"%", "%"+withoutSpace+"%", "%"+withoutSpace+"%")
	}

	count, err := query.ScanAndCount(ctx)
	if err != nil {
		return nil, err
	}

	return &model.UserConnection{
		Edges:      users,
		PageInfo:   nil,
		TotalCount: count,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*db.User, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var user db.User
	err = r.DB.NewSelect().Model(&user).Where("id = ?", id).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Competence is the resolver for the competence field.
func (r *queryResolver) Competence(ctx context.Context, id string) (*db.Competence, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var competence db.Competence
	err = r.DB.NewSelect().Model(&competence).Where("id = ?", id).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &competence, nil
}

// Competences is the resolver for the competences field.
func (r *queryResolver) Competences(ctx context.Context, limit *int, offset *int, filter *model.CompetenceFilterInput, search *string) (*model.CompetenceConnection, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var pageLimit = 10
	if limit != nil {
		pageLimit = *limit
	}

	var pageOffset = 0
	if offset != nil {
		pageOffset = *offset
	}

	var competences []*db.Competence
	query := r.DB.NewSelect().
		Model(&competences).
		Where("organisation_id = ?", currentUser.OrganisationID).
		Limit(pageLimit).
		Offset(pageOffset).
		Order("name ASC")

	if search != nil {
		query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *search))
	}

	if filter != nil {
		if filter.Type != nil {
			if len(filter.Type) == 1 {
				query.Where("competence_type = ?", filter.Type[0])
			} else {
				query.Where("competence_type IN (?)", bun.In(filter.Type))
			}
		}

		if filter.Parents != nil {
			if len(filter.Parents) == 1 {
				query.Where("competence_id = ?", filter.Parents[0])
			} else {
				query.Where("competence_id IN (?)", bun.In(filter.Parents))
			}
		}
	}

	count, err := query.ScanAndCount(ctx)

	if err != nil {
		return nil, err
	}

	return &model.CompetenceConnection{
		Edges:      competences,
		PageInfo:   nil,
		TotalCount: count,
	}, nil
}

// Report is the resolver for the report field.
func (r *queryResolver) Report(ctx context.Context, id string) (*db.Report, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	// query the report
	var report db.Report
	err = r.DB.NewSelect().Model(&report).Where("id = ?", id).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

// Reports is the resolver for the reports field.
func (r *queryResolver) Reports(ctx context.Context, limit *int, offset *int) (*model.ReportConnection, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	pageLimit := 10
	if limit != nil {
		pageLimit = *limit
	}

	pageOffset := 0
	if offset != nil {
		pageOffset = *offset
	}

	var reports []*db.Report
	count, err := r.DB.NewSelect().Model(&reports).Where("organisation_id = ?", currentUser.OrganisationID).Order("created_at DESC").Limit(pageLimit).Offset(pageOffset).ScanAndCount(ctx)
	if err != nil {
		return nil, err
	}

	return &model.ReportConnection{
		Edges:      reports,
		PageInfo:   nil,
		TotalCount: count,
	}, nil
}

// Tag is the resolver for the tag field.
func (r *queryResolver) Tag(ctx context.Context, id string) (*db.Tag, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var tag db.Tag
	err = r.DB.NewSelect().Model(&tag).Where("id = ?", id).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, limit *int, offset *int) ([]*db.Tag, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	pageLimit := 10
	if limit != nil {
		pageLimit = *limit
	}

	pageOffset := 0
	if offset != nil {
		pageOffset = *offset
	}

	var tags []*db.Tag
	err = r.DB.NewSelect().
		Model(&tags).
		Where("organisation_id = ?", currentUser.OrganisationID).
		Limit(pageLimit).
		Offset(pageOffset).
		Order("name").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// UserStudents is the resolver for the userStudents field.
func (r *queryResolver) UserStudents(ctx context.Context, limit *int, offset *int) (*model.UserStudentConnection, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	pageLimit := 10
	if limit != nil {
		pageLimit = *limit
	}

	pageOffset := 0
	if offset != nil {
		pageOffset = *offset
	}

	var userStudents []*db.UserStudent
	count, err := r.DB.NewSelect().Model(&userStudents).Where("organisation_id = ?", currentUser.OrganisationID).Limit(pageLimit).Offset(pageOffset).ScanAndCount(ctx)
	if err != nil {
		return nil, err
	}

	return &model.UserStudentConnection{
		Edges:      userStudents,
		PageInfo:   nil,
		TotalCount: count,
	}, nil
}

// UserStudent is the resolver for the userStudent field.
func (r *queryResolver) UserStudent(ctx context.Context, id string) (*db.UserStudent, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var userStudent db.UserStudent
	err = r.DB.NewSelect().Model(&userStudent).Where("id = ?", id).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &userStudent, nil
}

// Meta is the resolver for the meta field.
func (r *reportResolver) Meta(ctx context.Context, obj *db.Report) (string, error) {
	/// meta is a jsonb field, so we need to unmarshal it
	var meta map[string]interface{}
	err := json.Unmarshal(obj.Meta.RawMessage, &meta)
	if err != nil {
		return "", err
	}

	// return meta as a string
	return fmt.Sprintf("%v", meta), nil
}

// User is the resolver for the user field.
func (r *reportResolver) User(ctx context.Context, obj *db.Report) (*db.User, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var user db.User
	err = r.DB.NewSelect().Model(&user).Where("id = ?", obj.UserID).Where("organisation_id = ?", currentUser.OrganisationID).WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// StudentUser is the resolver for the studentUser field.
func (r *reportResolver) StudentUser(ctx context.Context, obj *db.Report) (*db.User, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var user db.User
	err = r.DB.NewSelect().
		Model(&user).
		Where("id = ?", obj.StudentUserID).
		Where("organisation_id = ?", currentUser.OrganisationID).
		WhereAllWithDeleted().
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// File is the resolver for the file field.
func (r *reportResolver) File(ctx context.Context, obj *db.Report) (*db.File, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var file db.File
	err = r.DB.NewSelect().Model(&file).Where("id = ?", obj.FileID).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &file, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *reportResolver) DeletedAt(ctx context.Context, obj *db.Report) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// Color is the resolver for the color field.
func (r *tagResolver) Color(ctx context.Context, obj *db.Tag) (string, error) {
	if obj.Color.Valid {
		return obj.Color.String, nil
	}

	return "", nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *tagResolver) DeletedAt(ctx context.Context, obj *db.Tag) (*time.Time, error) {
	if obj.DeletedAt.IsZero() {
		return &obj.DeletedAt.Time, nil
	}

	return nil, nil
}

// Student is the resolver for the student field.
func (r *userResolver) Student(ctx context.Context, obj *db.User) (*db.UserStudent, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var userStudent db.UserStudent
	err = r.DB.NewSelect().Model(&userStudent).Where("user_id = ?", obj.ID).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &userStudent, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *userResolver) DeletedAt(ctx context.Context, obj *db.User) (*time.Time, error) {
	if obj.DeletedAt.IsZero() {
		return &obj.DeletedAt.Time, nil
	}

	return nil, nil
}

// Competence is the resolver for the competence field.
func (r *userCompetenceResolver) Competence(ctx context.Context, obj *db.UserCompetence) (*db.Competence, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var competence db.Competence
	err = r.DB.NewSelect().Model(&competence).Where("id = ?", obj.CompetenceID).Where("organisation_id = ?", currentUser.OrganisationID).WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &competence, nil
}

// Entry is the resolver for the entry field.
func (r *userCompetenceResolver) Entry(ctx context.Context, obj *db.UserCompetence) (*db.Entry, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	if sql.NullString(obj.EntryID).String == "" {
		return nil, nil
	}

	var entry db.Entry
	err = r.DB.NewSelect().Model(&entry).Where("id = ?", obj.EntryID).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entry, nil
}

// User is the resolver for the user field.
func (r *userCompetenceResolver) User(ctx context.Context, obj *db.UserCompetence) (*db.User, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var user db.User
	err = r.DB.NewSelect().Model(&user).Where("id = ?", obj.UserID).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *userCompetenceResolver) CreatedBy(ctx context.Context, obj *db.UserCompetence) (*db.User, error) {
	currentUser, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, nil
	}

	var user db.User
	err = r.DB.NewSelect().Model(&user).Where("id = ?", obj.CreatedBy).Where("organisation_id = ?", currentUser.OrganisationID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// LeftAt is the resolver for the leftAt field.
func (r *userStudentResolver) LeftAt(ctx context.Context, obj *db.UserStudent) (*time.Time, error) {
	if !obj.LeftAt.IsZero() {
		return &obj.LeftAt.Time, nil
	}

	return nil, nil
}

// Birthday is the resolver for the birthday field.
func (r *userStudentResolver) Birthday(ctx context.Context, obj *db.UserStudent) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: Birthday - birthday"))
}

// Nationality is the resolver for the nationality field.
func (r *userStudentResolver) Nationality(ctx context.Context, obj *db.UserStudent) (*string, error) {
	panic(fmt.Errorf("not implemented: Nationality - nationality"))
}

// Comments is the resolver for the comments field.
func (r *userStudentResolver) Comments(ctx context.Context, obj *db.UserStudent) (*string, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// JoinedAt is the resolver for the joinedAt field.
func (r *userStudentResolver) JoinedAt(ctx context.Context, obj *db.UserStudent) (*time.Time, error) {
	if !obj.JoinedAt.IsZero() {
		return &obj.JoinedAt.Time, nil
	}

	return nil, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *userStudentResolver) DeletedAt(ctx context.Context, obj *db.UserStudent) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// EntriesCount is the resolver for the entriesCount field.
func (r *userStudentResolver) EntriesCount(ctx context.Context, obj *db.UserStudent) (int, error) {
	currentUser := middleware.ForContext(ctx)
	if currentUser == nil {
		return 0, errors.New("no user found in the context")
	}

	count, err := r.DB.NewSelect().Model(&db.EntryUser{}).Where("user_id = ?", obj.UserID).Where("organisation_id = ?", currentUser.OrganisationID).Count(ctx)

	if err != nil {
		return 0, err
	}

	return count, nil
}

// CompetencesCount is the resolver for the competencesCount field.
func (r *userStudentResolver) CompetencesCount(ctx context.Context, obj *db.UserStudent) (int, error) {
	currentUser := middleware.ForContext(ctx)
	if currentUser == nil {
		return 0, errors.New("no user found in the context")
	}

	count, err := r.DB.NewSelect().Model(&db.UserCompetence{}).Where("user_id = ?", obj.UserID).Where("organisation_id = ?", currentUser.OrganisationID).Count(ctx)

	if err != nil {
		return 0, err
	}

	return count, nil
}

// EventsCount is the resolver for the eventsCount field.
func (r *userStudentResolver) EventsCount(ctx context.Context, obj *db.UserStudent) (int, error) {
	currentUser := middleware.ForContext(ctx)
	if currentUser == nil {
		return 0, errors.New("no user found in the context")
	}

	count, err := r.DB.NewSelect().
		Model(&db.EntryEvent{}).
		Join("JOIN entry_users ON entry_users.entry_id = entry_event.entry_id").
		Where("entry_users.user_id = ?", obj.UserID).
		Where("entry_event.organisation_id = ?", currentUser.OrganisationID).
		Count(ctx)

	if err != nil {
		return 0, err
	}

	return count, nil
}

// Competence returns CompetenceResolver implementation.
func (r *Resolver) Competence() CompetenceResolver { return &competenceResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Organisation returns OrganisationResolver implementation.
func (r *Resolver) Organisation() OrganisationResolver { return &organisationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Report returns ReportResolver implementation.
func (r *Resolver) Report() ReportResolver { return &reportResolver{r} }

// Tag returns TagResolver implementation.
func (r *Resolver) Tag() TagResolver { return &tagResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// UserCompetence returns UserCompetenceResolver implementation.
func (r *Resolver) UserCompetence() UserCompetenceResolver { return &userCompetenceResolver{r} }

// UserStudent returns UserStudentResolver implementation.
func (r *Resolver) UserStudent() UserStudentResolver { return &userStudentResolver{r} }

type competenceResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type organisationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reportResolver struct{ *Resolver }
type tagResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userCompetenceResolver struct{ *Resolver }
type userStudentResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func isStringInArray(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}
