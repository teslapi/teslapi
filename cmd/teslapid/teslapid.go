package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	badger "github.com/dgraph-io/badger"
	"github.com/dgrijalva/jwt-go"
	"github.com/teslapi/teslapi/internal/recordings"
)

type config struct {
	StorageDir string
	DBDir      string
	DB         *badger.DB
}

type clipsPageData struct {
	Title string
	Clips []recordings.Clip
}

func main() {
	config := config{
		StorageDir: "./storage/TeslaUSB",
		DBDir:      "./storage/db",
	}

	db, err := badger.Open(badger.DefaultOptions(config.DBDir))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	config.DB = db

	// scan the directory for clips
	clips, err := scan(config)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		tmpl.Execute(w, clipsPageData{
			Title: "Clips",
			Clips: clips,
		})
	})

	http.HandleFunc("api/login", func(w http.ResponseWriter, r *http.Request) {
		type loginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		type loginResponse struct {
			Token string `json:"token"`
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		r.ParseForm()

		request := loginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		if request.Username != os.Getenv("TESLAPI_USERNAME") || request.Password != os.Getenv("TESLAPI_PASSWORD") {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("invalid request")
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * 30).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
			Issuer:    "teslapi",
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(os.Getenv("TESLAPI_KEY")))
		if err != nil {
			log.Fatal(err)
			return
		}

		response := loginResponse{Token: tokenString}
		body, err := json.Marshal(&response)
		if err != nil {
			log.Fatal(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	http.HandleFunc("/api/recordings", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		clips := []recordings.Clip{}
		db.View(func(txn *badger.Txn) error {
			opts := badger.DefaultIteratorOptions
			opts.PrefetchSize = 10
			it := txn.NewIterator(opts)
			defer it.Close()
			for it.Rewind(); it.Valid(); it.Next() {
				item := it.Item()
				k := item.Key()
				err := item.Value(func(v []byte) error {
					c := recordings.Clip{
						Name: string(k),
					}
					err := json.Unmarshal(v, &c)
					if err != nil {
						return err
					}

					// filter the params
					if params.Get("type") != "" && params.Get("camera") != "" {
						if params.Get("type") == c.Type && params.Get("camera") == c.Camera {
							clips = append(clips, c)
						}
					}

					if params.Get("type") != "" && params.Get("camera") == "" {
						if params.Get("type") == c.Type {
							clips = append(clips, c)
						}
					}

					if params.Get("type") == "" && params.Get("camera") != "" {
						if params.Get("camera") == c.Camera {
							clips = append(clips, c)
						}
					}

					if params.Get("type") == "" && params.Get("camera") != "" {
						if params.Get("camera") == c.Camera {
							clips = append(clips, c)
						}
					}

					return nil
				})
				if err != nil {
					return err
				}
			}
			return nil
		})
		w.Header().Set("Content-Type", "application/json")
		body, err := json.Marshal(&clips)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("%v", err)
	}
}

func scan(config config) ([]recordings.Clip, error) {
	clips := []recordings.Clip{}
	err := filepath.Walk(config.StorageDir, func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == false && info.Name() != ".DS_Store" {
			// determine the type
			clipType := "saved"
			if regexp.MustCompile(`RecentClips`).MatchString(root) {
				clipType = "recent"
			}

			// get the camera that recorded the clip
			camera := "front"
			if regexp.MustCompile(`right_repeater`).MatchString(info.Name()) {
				camera = "right"
			}
			if regexp.MustCompile(`left_repeater`).MatchString(info.Name()) {
				camera = "left"
			}
			c := recordings.Clip{
				Name:          info.Name(),
				Type:          clipType,
				Camera:        camera,
				FileLocation:  root,
				FileTimestamp: info.ModTime(),
				Uploaded:      false,
			}
			err := config.DB.Update(func(txn *badger.Txn) error {
				// TODO check for the object by id
				encoded, err := json.Marshal(c)
				if err != nil {
					return err
				}

				txn.Set([]byte(c.Name), encoded)
				return nil
			})

			if err != nil {
				return err
			}

			clips = append(clips, c)
		}

		return nil
	})

	if err != nil {
		return clips, err
	}

	return clips, err
}
