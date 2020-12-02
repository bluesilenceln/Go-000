package service

import (
	"Go-000/Week02/dao"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type Service struct {
	dao *dao.Dao
}

func NewService(addr string, d *dao.Dao) *Service {
	s := new(Service)
	s.dao = d
	Go(func() {
		http.HandleFunc("/user", s.user)
		_ = http.ListenAndServe(addr, nil)
	})

	return s
}

func (s *Service) user(w http.ResponseWriter, r *http.Request) {
	var res string
	id := r.URL.Query().Get("id" )
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name, err := s.dao.GetUserById(id)
	if err != nil {
		if errors.Is(err, dao.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Printf("user: id: %s %+v", id, err)
		return
	}

	res = name
	_, _ = w.Write([]byte(res))
}

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		f()
	}()
}