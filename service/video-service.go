package service

import (
	"MSquash/entity"
	"fmt"
)

type VideoService interface {
	Save(entity.Video) (entity.Video, error)
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func (service *videoService) Save(video entity.Video) (entity.Video, error) {
	for _, v := range service.videos {
		if v.URL == video.URL {
			return entity.Video{}, fmt.Errorf("video with URL %s already exists", video.URL)
		}
	}
	service.videos = append(service.videos, video)
	return video, nil
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

func New() VideoService {
	return &videoService{
		videos: []entity.Video{},
	}
}
