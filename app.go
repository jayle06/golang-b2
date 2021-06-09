package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"ocg.com/hw-json/model"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type KeyValue struct {
	Key string
	Value int
}

type StringFloat struct {
	Str string
	Value float64
}

func GetAgeOfEachPeople(dateString string) (result int){
	var dateSlice = strings.Split(dateString, "-")
	year, _ := strconv.Atoi(dateSlice[0])
	monthInt, _ := strconv.Atoi(dateSlice[1])
	month := time.Month(monthInt)
	day, _ := strconv.Atoi(dateSlice[2])

	age := time.Now().Year() - year

	if month == time.Now().Month() && day > time.Now().Day() {
		age -= 1
	}
	if month > time.Now().Month() {
		age -= 1
	}
	return age
}

func AverageOfAge(users []model.User) (result float64){

	sum := 0
	count := len(users)
	for _, user := range users {
		sum += GetAgeOfEachPeople(user.BirthDate)
	}
	result = float64(sum) / float64(count)

	if math.IsNaN(result) {
		result = 0
	}
	return result
}

func AverageOfSalary(users []model.User) (result float64){
	sum := 0
	count := len(users)
	for _, user := range users {
		sum += user.Salary
	}
	result = float64(sum) / float64(count)

	if math.IsNaN(result) {
		result = 0
	}
	return result
}

func GetPeopleIsADeveloper(users []model.User) (result []model.User){
	for _, user := range users {
		if user.Job == "developer" {
			result = append(result, user)
		}
	}
	return result
}

func ConvertMapToSliceFloatAndSort(myMap map[string]float64)(result []StringFloat){
	for key, val := range myMap {
		result = append(result, StringFloat{key, val})
	}
	sort.Slice(result, func(i, j int) bool{
		return result[i].Value > result[j].Value
	})
	return result
}

func ConvertMapToSliceAndSort(myMap map[string]int)(result []KeyValue){
	for key, val := range myMap {
		result = append(result, KeyValue{key, val})
	}
	sort.Slice(result, func(i, j int) bool{
		return result[i].Value > result[j].Value
	})
	return result
}

func Contains(Str []string, e string) bool {
	for _, s := range Str {
		if s == e {
			return true
		}
	}
	return false
}

func GetNameOfCities(users []model.User) (result []string){
	for _, user := range users {
		if !Contains(result, user.City) {
			result = append(result, user.City)
		}
	}
	return result
}

func GetNameOfJobs(users []model.User) (result []string){
	for _, user := range users {
		if !Contains(result, user.Job) {
			result = append(result, user.Job)
		}
	}
	return result
}

func GetHottestJob(users []model.User) (result []KeyValue){
	var userJobMap = GroupPeopleByJob(users)
	var sliceConv = ConvertMapToSliceAndSort(userJobMap)
	var max = sliceConv[0].Value
	for _, e := range sliceConv {
		if e.Value == max {
			result = append(result, e)
		}
	}
	return result
}

func GroupPeopleByCity(users []model.User) (result map[string][]model.User){
	result = make(map[string][]model.User)
	for _, user := range users {
		result[user.City] = append(result[user.City], user)
	}
	return result
}

func GetGroupOfPeoPleByJob(users []model.User) (result map[string][]model.User){
	result = make(map[string][]model.User)
	for _, user := range users {
		result[user.Job] = append(result[user.Job], user)
	}
	return result
}

func GroupPeopleByJob(users []model.User) (result map[string]int) {
	result = make(map[string]int)
	for _, user := range users {
		result[user.Job]++
	}
	return result
}

func Top5JobsByNumber(users []model.User) (result []KeyValue) {
	var userJobMap = GroupPeopleByJob(users)
	var sliceConv = ConvertMapToSliceAndSort(userJobMap)

	for i:= 0; i<5; i++ {
		result = append(result, sliceConv[i])
	}
	return result
}

func Top5CitiesByNumber(users []model.User) (result []KeyValue){
	cityCountMap := make(map[string]int)
	for _, user := range users {
		cityCountMap[user.City]++
	}

	var sliceConv = ConvertMapToSliceAndSort(cityCountMap)

	for i:= 0; i<5; i++ {
		result = append(result, sliceConv[i])
	}
	return result
}

func TopJobByNumberInEachCity(users []model.User) (result map[string][]KeyValue) {
	result = make(map[string][]KeyValue)
	var userSameCityMap = GroupPeopleByCity(users)
	var nameOfCitySlice = GetNameOfCities(users)
	var hottestJob []KeyValue
	for i := 0; i < len(nameOfCitySlice) ; i++ {
		hottestJob = GetHottestJob(userSameCityMap[nameOfCitySlice[i]])
		for _, hot := range hottestJob {
			result[nameOfCitySlice[i]] = append(result[nameOfCitySlice[i]], hot)
		}
	}
	return result
}

func AverageSalaryByJob(users []model.User) (result map[string]float64){
	result = make(map[string]float64)
	var usersInSameJob = GetGroupOfPeoPleByJob(users)
	var jobNames = GetNameOfJobs(users)
	for _, job := range jobNames {
		result[job] = AverageOfSalary(usersInSameJob[job])
	}
	return result
}

func FiveCitiesHasTopAverageSalary(users []model.User) (result []StringFloat) {
	var mapCityAverageSalary = make(map[string]float64)
	var userInSameCity = GroupPeopleByCity(users)
	var nameOfCity = GetNameOfCities(users)

	for _, city := range nameOfCity {
		mapCityAverageSalary[city] = AverageOfSalary(userInSameCity[city])
	}

	var sliceConv = ConvertMapToSliceFloatAndSort(mapCityAverageSalary)

	for i:= 0; i<5; i++ {
		result = append(result, sliceConv[i])
	}

	return result
}

func FiveCitiesHasTopSalaryForDeveloper(users []model.User) (result []StringFloat) {
	var mapCityAverageSalary = make(map[string]float64)
	var userInSameCity = GroupPeopleByCity(users)
	var nameOfCity = GetNameOfCities(users)

	for _, city := range nameOfCity {
		var usersInCity = userInSameCity[city]
		var devInCity = GetPeopleIsADeveloper(usersInCity)
		mapCityAverageSalary[city] = AverageOfSalary(devInCity)
	}

	var sliceConv = ConvertMapToSliceFloatAndSort(mapCityAverageSalary)

	for i:= 0; i<5; i++ {
		result = append(result, sliceConv[i])
	}

	return result
}

func AverageAgePerJob(users []model.User) (result map[string]float64){
	result = make(map[string]float64)
	var usersInSameJob = GetGroupOfPeoPleByJob(users)
	var jobs = GetNameOfJobs(users)

	for _, job := range jobs {
		result[job] = AverageOfAge(usersInSameJob[job])
	}
	return result
}

func AverageAgePerCity(users []model.User) (result map[string]float64){
	result = make(map[string]float64)
	var usersInSameCity = GroupPeopleByCity(users)
	var cities = GetNameOfCities(users)

	for _, city := range cities {
		result[city] = AverageOfAge(usersInSameCity[city])
	}
	return result
}


func main() {
	jsonFile, err := os.Open("person.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []model.User

	json.Unmarshal(byteValue, &users)

	//for i := 0; i < 5; i++ {
	//	fmt.Println(users[i])
	//}

	//2.1. get people in same city
	fmt.Println(GroupPeopleByCity(users))

	//2.2. count people have same job
	fmt.Println(GroupPeopleByJob(users))

	//2.3. 5 hot jobs
	fmt.Println(Top5JobsByNumber(users))

	//2.4 5 big cities
	fmt.Println(Top5CitiesByNumber(users))

	//2.5 Hottest in the city
	fmt.Println(TopJobByNumberInEachCity(users))

	//2.6 Average of Salary
	fmt.Println(AverageSalaryByJob(users))

	//2.7 5 cities has highest salary
	fmt.Println(FiveCitiesHasTopAverageSalary(users))

	//2.8 5 cities has highest salary for developer
	fmt.Println(FiveCitiesHasTopSalaryForDeveloper(users))

	//2.9 Average of age by each job
	fmt.Println(AverageAgePerJob(users))

	//2.10 Average of age in each city
	fmt.Println(AverageAgePerCity(users))
}
