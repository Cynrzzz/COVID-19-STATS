package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AutoGenerated struct {
	Fips          string      `json:"fips"`
	Country       string      `json:"country"`
	State         interface{} `json:"state"`
	County        interface{} `json:"county"`
	Hsa           interface{} `json:"hsa"`
	HsaName       interface{} `json:"hsaName"`
	Level         string      `json:"level"`
	Lat           interface{} `json:"lat"`
	LocationID    string      `json:"locationId"`
	Long          interface{} `json:"long"`
	Population    int         `json:"population"`
	HsaPopulation interface{} `json:"hsaPopulation"`
	Metrics       struct {
		TestPositivityRatio        float64 `json:"testPositivityRatio"`
		TestPositivityRatioDetails struct {
			Source string `json:"source"`
		} `json:"testPositivityRatioDetails"`
		CaseDensity                     float64     `json:"caseDensity"`
		WeeklyNewCasesPer100K           float64     `json:"weeklyNewCasesPer100k"`
		ContactTracerCapacityRatio      interface{} `json:"contactTracerCapacityRatio"`
		InfectionRate                   float64     `json:"infectionRate"`
		InfectionRateCI90               float64     `json:"infectionRateCI90"`
		IcuCapacityRatio                float64     `json:"icuCapacityRatio"`
		BedsWithCovidPatientsRatio      float64     `json:"bedsWithCovidPatientsRatio"`
		WeeklyCovidAdmissionsPer100K    float64     `json:"weeklyCovidAdmissionsPer100k"`
		VaccinationsInitiatedRatio      float64     `json:"vaccinationsInitiatedRatio"`
		VaccinationsCompletedRatio      float64     `json:"vaccinationsCompletedRatio"`
		VaccinationsAdditionalDoseRatio float64     `json:"vaccinationsAdditionalDoseRatio"`
	} `json:"metrics"`
	RiskLevels struct {
		Overall                    int `json:"overall"`
		TestPositivityRatio        int `json:"testPositivityRatio"`
		CaseDensity                int `json:"caseDensity"`
		ContactTracerCapacityRatio int `json:"contactTracerCapacityRatio"`
		InfectionRate              int `json:"infectionRate"`
		IcuCapacityRatio           int `json:"icuCapacityRatio"`
	} `json:"riskLevels"`
	CdcTransmissionLevel int `json:"cdcTransmissionLevel"`
	CommunityLevels      struct {
		CdcCommunityLevel interface{} `json:"cdcCommunityLevel"`
		CanCommunityLevel int         `json:"canCommunityLevel"`
	} `json:"communityLevels"`
	Actuals struct {
		Cases          int `json:"cases"`
		Deaths         int `json:"deaths"`
		PositiveTests  int `json:"positiveTests"`
		NegativeTests  int `json:"negativeTests"`
		ContactTracers int `json:"contactTracers"`
		HospitalBeds   struct {
			Capacity              int `json:"capacity"`
			CurrentUsageTotal     int `json:"currentUsageTotal"`
			CurrentUsageCovid     int `json:"currentUsageCovid"`
			WeeklyCovidAdmissions int `json:"weeklyCovidAdmissions"`
		} `json:"hospitalBeds"`
		HsaHospitalBeds struct {
			Capacity              interface{} `json:"capacity"`
			CurrentUsageTotal     interface{} `json:"currentUsageTotal"`
			CurrentUsageCovid     interface{} `json:"currentUsageCovid"`
			WeeklyCovidAdmissions interface{} `json:"weeklyCovidAdmissions"`
		} `json:"hsaHospitalBeds"`
		IcuBeds struct {
			Capacity          int `json:"capacity"`
			CurrentUsageTotal int `json:"currentUsageTotal"`
			CurrentUsageCovid int `json:"currentUsageCovid"`
		} `json:"icuBeds"`
		HsaIcuBeds struct {
			Capacity          interface{} `json:"capacity"`
			CurrentUsageTotal interface{} `json:"currentUsageTotal"`
			CurrentUsageCovid interface{} `json:"currentUsageCovid"`
		} `json:"hsaIcuBeds"`
		NewCases                          int         `json:"newCases"`
		NewDeaths                         int         `json:"newDeaths"`
		VaccinesDistributed               int         `json:"vaccinesDistributed"`
		VaccinationsInitiated             int         `json:"vaccinationsInitiated"`
		VaccinationsCompleted             int         `json:"vaccinationsCompleted"`
		VaccinationsAdditionalDose        int         `json:"vaccinationsAdditionalDose"`
		VaccinesAdministered              int         `json:"vaccinesAdministered"`
		VaccinesAdministeredDemographics  interface{} `json:"vaccinesAdministeredDemographics"`
		VaccinationsInitiatedDemographics interface{} `json:"vaccinationsInitiatedDemographics"`
	} `json:"actuals"`
	Annotations struct {
		Cases                           interface{} `json:"cases"`
		Deaths                          interface{} `json:"deaths"`
		PositiveTests                   interface{} `json:"positiveTests"`
		NegativeTests                   interface{} `json:"negativeTests"`
		ContactTracers                  interface{} `json:"contactTracers"`
		HospitalBeds                    interface{} `json:"hospitalBeds"`
		HsaHospitalBeds                 interface{} `json:"hsaHospitalBeds"`
		IcuBeds                         interface{} `json:"icuBeds"`
		HsaIcuBeds                      interface{} `json:"hsaIcuBeds"`
		NewCases                        interface{} `json:"newCases"`
		NewDeaths                       interface{} `json:"newDeaths"`
		VaccinesDistributed             interface{} `json:"vaccinesDistributed"`
		VaccinationsInitiated           interface{} `json:"vaccinationsInitiated"`
		VaccinationsCompleted           interface{} `json:"vaccinationsCompleted"`
		VaccinationsAdditionalDose      interface{} `json:"vaccinationsAdditionalDose"`
		VaccinesAdministered            interface{} `json:"vaccinesAdministered"`
		TestPositivityRatio             interface{} `json:"testPositivityRatio"`
		CaseDensity                     interface{} `json:"caseDensity"`
		WeeklyNewCasesPer100K           interface{} `json:"weeklyNewCasesPer100k"`
		ContactTracerCapacityRatio      interface{} `json:"contactTracerCapacityRatio"`
		InfectionRate                   interface{} `json:"infectionRate"`
		InfectionRateCI90               interface{} `json:"infectionRateCI90"`
		IcuCapacityRatio                interface{} `json:"icuCapacityRatio"`
		BedsWithCovidPatientsRatio      interface{} `json:"bedsWithCovidPatientsRatio"`
		WeeklyCovidAdmissionsPer100K    interface{} `json:"weeklyCovidAdmissionsPer100k"`
		VaccinationsInitiatedRatio      interface{} `json:"vaccinationsInitiatedRatio"`
		VaccinationsCompletedRatio      interface{} `json:"vaccinationsCompletedRatio"`
		VaccinationsAdditionalDoseRatio interface{} `json:"vaccinationsAdditionalDoseRatio"`
	} `json:"annotations"`
	LastUpdatedDate string      `json:"lastUpdatedDate"`
	URL             interface{} `json:"url"`
}

func main() {
	resp, err := http.Get("https://api.covidactnow.org/v2/country/US.json?apiKey=b1c2a342c50d43b886d0b69c05a311d7")
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	var a AutoGenerated
	json.Unmarshal([]byte(body), &a)
	log.Println("Total Deaths:", a.Actuals.Deaths)
}
