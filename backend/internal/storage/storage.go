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
	Password   string    `json:"password"`
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
	ID              int       `json:"id_bid"`
	OptionalGoal    string    `json:"optional_goal"`
	OptionalMessage string    `json:"optional_message"`
	Created         time.Time `json:"created"`
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

// set global
var StoragePSQL *Storage = nil

func (s *Storage) ConnectToDB() (err error) {
	const op = "storage.New"

	s.SetStorageSettings()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.dbSettings.Host, s.dbSettings.Port, s.dbSettings.User, s.dbSettings.Password,
		s.dbSettings.DBname,
	)

	db, err := sql.Open("postgres", connStr)
	s.db = db

	// set global
	StoragePSQL = s

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

/* Setters */

func (s *Storage) SetStructUsers(user User) error {

	const op = "storage.SetStructUsers"

	query := "INSERT INTO users (first_name, second_name, last_name, phone, email, profile_img, created, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	_, err := s.db.Exec(query, user.FirstName, user.SecondName, user.LastName, user.Phone, user.Email, user.ProfileImg, user.Created, user.Password)
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
		"INSERT INTO bid (optional_goal, optional_message, created) VALUES ($1, $2, $3)"

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

/* Getters */

func (s *Storage) GetUserByID(userID int) (User, error) {

	const op = "storage.GetUserByID"

	var user User

	err := s.db.QueryRow("SELECT id_users, first_name, second_name, last_name, phone, email, profile_img, created, password FROM users WHERE id_users = $1", userID).Scan(
		&user.ID,
		&user.FirstName,
		&user.SecondName,
		&user.LastName,
		&user.Phone,
		&user.Email,
		&user.ProfileImg,
		&user.Created,
		&user.Password,
	)

	if err != nil {
		return User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *Storage) GetUserContactsById(userID int) (Contacts, error) {

	const op = "storage.GetUserContactsById"

	var contacts Contacts

	err := s.db.QueryRow(fmt.Sprintf("SELECT option_link, tg_link, inst_link, vk_link, refusers FROM contacts WHERE refusers = %d", userID)).Scan(
		&contacts.OptionLink,
		&contacts.TgLink,
		&contacts.InstLink,
		&contacts.VkLink,
		&contacts.RefUsers,
	)

	if err != nil {
		return Contacts{}, fmt.Errorf("%s: %w", op, err)
	}

	return contacts, nil
}

func (s *Storage) GetLastWorkoutPlan(userID int) (WorkoutPlan, error) {

	const op = "storage.GetLastWorkoutPlan"

	query := (fmt.Sprintf("SELECT wp.id_workout_plan, wp.day_plan, wp.dplan_created, wp.week_plan, wp.wplan_created"+
		"FROM workout_plan wp"+
		"JOIN workout_plan_list wpl ON wp.id_workout_plan = wpl.refworkout_plan"+
		"WHERE wpl.refusers = %d"+
		"ORDER BY wp.wplan_created DESC"+
		"LIMIT 1", userID))

	var latestWorkoutPlan WorkoutPlan

	err := s.db.QueryRow(query, userID).Scan(
		&latestWorkoutPlan.ID,
		&latestWorkoutPlan.DayPlan,
		&latestWorkoutPlan.DPlanCreated,
		&latestWorkoutPlan.WeekPlan,
		&latestWorkoutPlan.WPlanCreated,
	)

	if err != nil {
		return WorkoutPlan{}, fmt.Errorf("%s: %w", op, err)
	}

	return latestWorkoutPlan, nil
}

/* TODO:
func (s *Storage) GetWorkoutPlanList(userID int) error {

	const op = "storage.GetWorkoutPlanList"

}
*/

func (s *Storage) GetLastDietPlan(userID int) (DietPlan, error) {

	const op = "storage.GetLastDietPlan"

	query := (fmt.Sprintf("SELECT dp.id_diet_plan, dp.day_plan, dp.dplan_created, dp.week_plan, wp.wplan_created"+
		"FROM diet_plan dp"+
		"JOIN diet_plan_list dpl ON dp.id_diet_plan = dpl.refdiet_plan"+
		"WHERE dpl.refusers = %d"+
		"ORDER BY dp.dplan_created DESC"+
		"LIMIT 1", userID))

	var latestDietPlan DietPlan

	err := s.db.QueryRow(query, userID).Scan(
		&latestDietPlan.ID,
		&latestDietPlan.DayPlan,
		&latestDietPlan.DPlanCreated,
		&latestDietPlan.WeekPlan,
		&latestDietPlan.WPlanCreated,
	)

	if err != nil {
		return DietPlan{}, fmt.Errorf("%s: %w", op, err)
	}

	return latestDietPlan, nil
}

/* TODO:
func (s *Storage) GetStructDietPlanList() error {

	const op = "storage.GetStructDietPlanList"

}
*/

func (s *Storage) GetCoachById(coachID int) (Coach, error) {

	const op = "storage.GetCoachById"

	var coach Coach

	err := s.db.QueryRow(fmt.Sprintf("SELECT id_coach, first_name, second_name, last_name, email, profile_img, created FROM coach WHERE id_coach = %d", coachID)).Scan(
		&coach.ID,
		&coach.FirstName,
		&coach.SecondName,
		&coach.LastName,
		&coach.Email,
		&coach.ProfileImg,
		&coach.Created,
	)

	if err != nil {
		return Coach{}, fmt.Errorf("%s: %w", op, err)
	}

	return coach, nil
}

/* TODO:
func (s *Storage) GetStructFitClasses() error {

	const op = "storage.GetStructFitClasses"

	return nil
}
*/

/* TODO:
func (s *Storage) GetStructCoachClassesList() error {

	const op = "storage.GetStructCoachClassesList"

}
*/

/*TODO
func (s *Storage) GetStructBid() error {

	const op = "storage.GetStructBid"

	return nil
}
*/

func (s *Storage) GetLastBidByUser(userID int) (Bid, error) {

	const op = "storage.GetBidByUser"

	var bid Bid

	err := s.db.QueryRow(fmt.Sprintf("SELECT id_bid, optional_goal, optional_message FROM bid"+
		"JOIN bid_by_user bbu ON bbu.refid_bid = bid.id_bid"+
		"WHERE bid.id_bid = %d"+
		"ORDER BY bid.created"+
		"LIMIT 1", userID)).Scan(
		&bid.ID,
		&bid.OptionalGoal,
		&bid.OptionalMessage,
		&bid.Created,
	)

	if err != nil {
		return Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	return bid, nil
}

/* TODO:
func (s *Storage) GetStructFitClassesBid() error {

	const op = "storage.GetStructFitClassesBid"

}
*/

/* TODO:
func (s *Storage) GetCenters() error {

	const op = "storage.GetStructCenters"
}
*/

/* TODO: add getters for frontend methods */

/* Functions */

func (s *Storage) IsUserByEmail(email string) (User, error) {

	const op = "storage.IsUserByEmail"

	query := "SELECT id_users, email, password FROM users WHERE email = '%s' LIMIT 1"

	var user User
	err := s.db.QueryRow(fmt.Sprintf(query, email)).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
