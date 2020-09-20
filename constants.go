package main

// TechLevels - lists the tech levels supported by the app
var TechLevels = []string{"F", "G", "H", "J", "K"}

// Tons - lists the ship tonnage amounts supported by the app
var Tons = []string{"100", "200", "300", "400", "500", "600", "700", "800", "900", "1000", "1200", "1500", "2000", "2500",
	"3000", "3500", "4000", "5000", "6000", "7000", "7500", "8000", "9000", "10000", "12000", "12500", "13000", "14000",
	"15000", "16000", "180000", "20000", "22000", "25000", "30000", "35000", "40000", "50000", "60000", "70000", "80000",
	"90000", "100000", "120000", "125000", "150000", "200000", "225000", "250000", "275000", "300000", "350000", "375000",
	"400000", "500000", "600000", "700000", "650000", "800000", "900000", "1000000"}

// EngineLevel lists the engine levels we support, from 1 to 12 (i.e. J-1 to J-12, M-1 to M-12 * P-1 to P-12)
var EngineLevel = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
var Computer = [16]int{
	1, 2, 3, 4, 5, 7, 9, 11, 13, 15, 17, 20, 23, 26, 30, 35,
}