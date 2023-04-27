package service

import "MSquash/entity"

type TestCaseMangementService interface {
	Save(entity.TestCase) entity.TestCase
	FindAll() []entity.TestCase
}

type testCaseMangementService struct {
	testCases []entity.TestCase
}

func New() *testCaseMangementService {
	return &testCaseMangementService{}
}

func (service *testCaseMangementService) Save(testCase entity.TestCase) entity.TestCase {
	service.testCases = append(service.testCases, testCase)
	return testCase
}

func (service *testCaseMangementService) FindAll() []entity.TestCase {
	return service.testCases
}
