package highscore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"github.com/pkg/errors"
)

var userURL = "https://platfomer.firebaseio.com/dungeon/users/%s.json"
var highscoreURL = "https://platfomer.firebaseio.com/dungeon/highscore/%s.json"

type User struct {
	ID        string    `json:"id"`
	Connected time.Time `json:"connected"`
	Type      string    `json:"type"`
	Version   string    `json:"version"`
}

func NewUser(ID, clientType, version string) error {
	user := User{
		ID:        ID,
		Connected: time.Now(),
		Type:      clientType,
		Version:   version,
	}

	d, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf(userURL, ID), bytes.NewBuffer(d))
	// request.Header.Set("Content-Type", "application/json; charset=utf-8")
	// timeout := time.Duration(300 * time.Second)
	client := http.Client{
		// Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		return errors.Wrap(err, "save user")
	}
	defer response.Body.Close()
	d, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "read response failed with")
	}
	fmt.Println(string(d))
	return nil
}

func SaveScore(recordID string, userID string, t time.Duration, version string) error {
	record := Record{
		ID:      recordID,
		UserID:  userID,
		Time:    t,
		Version: version,
	}

	d, err := json.Marshal(record)
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf(highscoreURL, recordID), bytes.NewBuffer(d))
	// request.Header.Set("Content-Type", "application/json; charset=utf-8")
	// timeout := time.Duration(300 * time.Second)
	client := http.Client{
		// Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		return errors.Wrap(err, "save record")
	}
	defer response.Body.Close()
	d, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "read response failed with")
	}
	fmt.Println(string(d))
	return nil
}

type Record struct {
	ID      string        `json:"id"`
	UserID  string        `json:"name"`
	Time    time.Duration `json:"score"`
	Version string        `json:"version"`
}
type Records []Record

func (r Records) Limit(max int) Records {
	var records []Record
	for i, v := range r {
		if i >= max {
			break
		}
		records = append(records, v)
	}
	return records
}

func (r Records) ByVersion(ver string) Records {
	var records []Record
	for _, v := range r {
		if v.Version == ver {
			records = append(records, v)
		}
	}
	return records
}

func GetScore() (Records, error) {
	url := fmt.Sprintf(highscoreURL, "")
	// timeout := time.Duration(1 * time.Second)
	client := http.Client{
		// Timeout: timeout,
	}
	response, err := client.Get(url)
	if err != nil {
		return Records{}, err
	}
	defer response.Body.Close()

	highscore := map[string]Record{}
	if err := json.NewDecoder(response.Body).Decode(&highscore); err != nil {
		return Records{}, err
	}

	records := Records{}
	for _, v := range highscore {
		records = append(records, v)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].Time < records[j].Time
	})

	return records, nil
}
