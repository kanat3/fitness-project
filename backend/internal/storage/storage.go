package storage

import (
	"database/sql"
	"fmt"
	"time"

	"fitness-project/backend/internal/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	db         *sql.DB
	dbSettings config.DatabaseConfig
}

type User struct {
	ID         int       `json:"id_users"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	ProfileImg string    `json:"profile_img"`
	Created    time.Time `json:"created"`
}

type Contacts struct {
	OptionLink string `json:"option_link"`
	TgLink     string `json:"tg_link"`
	InstLink   string `json:"inst_link"`
	VkLink     string `json:"vk_link"`
	RefUsers   int    `json:"refusers"`
	ID         int    `json:"id_contacts"`
}

type WorkoutPlan struct {
	ID           int       `json:"id_workout_plan"`
	DayPlan      string    `json:"day_plan"`
	DPlanCreated time.Time `json:"dplan_created"`
	WeekPlan     string    `json:"week_plan"`
	WPlanCreated time.Time `json:"wplan_created"`
}

type WorkoutPlanList struct {
	RefUsers       int `json:"refusers"`
	RefWorkoutPlan int `json:"refworkout_plan"`
}

type DietPlan struct {
	ID           int       `json:"id_diet_plan"`
	DayPlan      string    `json:"day_plan"`
	DPlanCreated time.Time `json:"dplan_created"`
	WeekPlan     string    `json:"week_plan"`
	WPlanCreated time.Time `json:"wplan_created"`
}

type DietPlanList struct {
	RefUsers    int `json:"refusers"`
	RefDietPlan int `json:"refdiet_plan"`
}

type Coach struct {
	ID         int       `json:"id_coach"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	ProfileImg string    `json:"profile_img"`
	Created    time.Time `json:"created"`
}

type FitClasses struct {
	ID   int    `json:"id_fit_classes"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type CoachClassesList struct {
	RefCoach    int `json:"refcoach"`
	RefFitClass int `json:"reffit_classes"`
}

type Bid struct {
	ID              int    `json:"id_bid"`
	OptionalGoal    string `json:"optional_goal"`
	OptionalMessage string `json:"optional_message"`
}

type BidByUser struct {
	RefBid   int `json:"refid_bid"`
	RefUsers int `json:"refusers"`
}

type FitClassesBid struct {
	RefBid      int `json:"refid_bid"`
	RefFitClass int `json:"reffit_classes"`
}

type Goal struct {
	ID     int    `json:"id_goal"`
	RefBid int    `json:"refbid"`
	Type   string `json:"type"`
}

type Sex struct {
	ID     int    `json:"id_sex"`
	RefBid int    `json:"refbid"`
	Type   string `json:"type"`
}

type Centers struct {
	ID      int    `json:"id_centers"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type CentersList struct {
	RefBid     int `json:"refbid"`
	RefCenters int `json:"refcenters"`
}

type DayTime struct {
	ID     int    `json:"id_day_time"`
	RefBid int    `json:"refbid"`
	Type   string `json:"type"`
}

type Status struct {
	ID     int    `json:"id_status"`
	RefBid int    `json:"refbid"`
	Type   string `json:"type"`
}

func (s *Storage) ConnectToDB() (err error) {
	const op = "storage.New"

	s.SetStorageSettings()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.dbSettings.Host, s.dbSettings.Port, s.dbSettings.User, s.dbSettings.Password,
		s.dbSettings.DBname,
	)

	db, err := sql.Open("postgres", connStr)
	s.db = db

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(error.Error(err))
		return err
	}

	return nil
}

func (s *Storage) Close() {
	s.Close()
}

func (s *Storage) PingDB() error {

	s.SetStorageSettings()

	ping := "host=" + s.dbSettings.Host + " port=" + s.dbSettings.Port + " user=" + s.dbSettings.User +
		" password=" + s.dbSettings.Password + " sslmode=disable"
	db, error := sql.Open("postgres", ping)

	if error != nil {
		fmt.Println(error.Error())
		return error
	}

	error = db.Ping()

	return nil
}

func (s *Storage) SetStorageSettings() {
	cfg := config.GetConfig()
	s.dbSettings = cfg.DatabaseConfig
}

func (s *Storage) HealCheck() error {
	_, error := s.db.Query("select * from users;")

	if error != nil {
		fmt.Println(error.Error())
		return error
	}

	return nil
}

func (s *Storage) SetStructUsers(user User) error {

	const op = "storage.SetStructUsers"

	query := "INSERT INTO users (first_name, second_name, last_name, phone, email, profile_img, created) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := s.db.Exec(query, user.FirstName, user.SecondName, user.LastName, user.Phone, user.Email, user.ProfileImg, user.Created)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructContacts(contact Contacts) error {

	const op = "storage.SetStructContacts"

	query :=
		"INSERT INTO contacts (option_link, tg_link, inst_link, vk_link, refusers) VALUES ($1, $2, $3, $4, $5)"

	_, err := s.db.Exec(query, contact.OptionLink, contact.TgLink, contact.InstLink, contact.InstLink, contact.RefUsers)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructWorkoutPlan(wp WorkoutPlan) error {

	const op = "storage.SetStructWorkoutPlan"

	query :=
		"INSERT INTO workout_plan (day_plan, dplan_created, week_plan, wplan_created) VALUES ($1, $2, $3, $4)"

	_, err := s.db.Exec(query, wp.DayPlan, wp.DPlanCreated, wp.WeekPlan, wp.WPlanCreated)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructWorkoutPlanList(wpl WorkoutPlanList) error {

	const op = "storage.SetStructWorkoutPlanList"

	query :=
		"INSERT INTO workout_plan_list (refusers, refworkout_plan) VALUES ($1, $2)"

	_, err := s.db.Exec(query, wpl.RefUsers, wpl.RefWorkoutPlan)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructDietPlan(dp DietPlan) error {

	const op = "storage.SetStructDietPlan"

	query :=
		"INSERT INTO diet_plan (day_plan, dplan_created, week_plan, wplan_created) VALUES ($1, $2, $3, $4)"

	_, err := s.db.Exec(query, dp.DayPlan, dp.DPlanCreated, dp.WeekPlan, dp.WPlanCreated)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructDietPlanList(dpl DietPlanList) error {

	const op = "storage.SetStructDietPlanList"

	query :=
		"INSERT INTO diet_plan_list (refusers, refdiet_plan) VALUES ($1, $2)"

	_, err := s.db.Exec(query, dpl.RefUsers, dpl.RefDietPlan)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructCoach(coach Coach) error {

	const op = "storage.SetStructCoach"

	query :=
		"INSERT INTO coach (first_name, second_name, last_name, email, profile_img, created) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := s.db.Exec(query, coach.FirstName, coach.SecondName, coach.LastName, coach.Email, coach.ProfileImg, coach.Created)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructFitClasses(fc FitClasses) error {

	const op = "storage.SetStructFitClasses"

	query :=
		"INSERT INTO fit_classes (name, type) VALUES ($1, $2)"

	_, err := s.db.Exec(query, fc.Name, fc.Type)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructCoachClassesList(ccl CoachClassesList) error {

	const op = "storage.SetStructCoachClassesList"

	query :=
		"INSERT INTO coach_classes_list (refcoach, reffit_classes) VALUES ($1, $2)"

	_, err := s.db.Exec(query, ccl.RefCoach, ccl.RefFitClass)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructBid(bid Bid) error {

	const op = "storage.SetStructBid"

	query :=
		"INSERT INTO bid (optional_goal, optional_message) VALUES ($1, $2)"

	_, err := s.db.Exec(query, bid.OptionalGoal, bid.OptionalMessage)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructBidByUser(bbu BidByUser) error {

	const op = "storage.SetStructBidByUser"

	query :=
		"INSERT INTO bid_by_user (refid_bid, refusers) VALUES ($1, $2)"

	_, err := s.db.Exec(query, bbu.RefBid, bbu.RefUsers)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructFitClassesBid(fcb FitClassesBid) error {

	const op = "storage.SetStructFitClassesBid"

	query :=
		"INSERT INTO fit_classes_bid (refid_bid, reffit_classes) VALUES ($1, $2)"

	_, err := s.db.Exec(query, fcb.RefBid, fcb.RefFitClass)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructGoal(goal Goal) error {

	const op = "storage.SetStructGoal"

	query :=
		"INSERT INTO goal (refbid, type) VALUES ($1, $2)"

	_, err := s.db.Exec(query, goal.RefBid, goal.Type)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructSex(sex Sex) error {

	const op = "storage.SetStructSex"

	query :=
		"INSERT INTO sex (refbid, type) VALUES ($1, $2)"

	_, err := s.db.Exec(query, sex.RefBid, sex.Type)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructCenters(center Centers) error {

	const op = "storage.SetStructCenters"

	query :=
		"INSERT INTO centers (name, address, phone) VALUES ($1, $2, $3)"

	_, err := s.db.Exec(query, center.Name, center.Address, center.Phone)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructCentersList(cl CentersList) error {

	const op = "storage.SetStructCentersList"

	query :=
		"INSERT INTO centers_list (refbid, refcenters) VALUES ($1, $2)"

	_, err := s.db.Exec(query, cl.RefBid, cl.RefCenters)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructDayTime(daytime DayTime) error {

	const op = "storage.SetStructDayTime"

	query :=
		"INSERT INTO day_time (refbid, type) VALUES ($1, $2)"

	_, err := s.db.Exec(query, daytime.RefBid, daytime.Type)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SetStructStatus(status Status) error {

	const op = "storage.SetStructStatus"

	query :=
		"INSERT INTO status (refbid, type) VALUES ($1, $2)"

	_, err := s.db.Exec(query, status.RefBid, status.Type)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
