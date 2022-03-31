package model

import "time"

type Vehicle struct {
	StationLocatorURL string `json:"station_locator_url"`
	TotalResults      int    `json:"total_results"`
	StationCounts     struct {
		Total int `json:"total"`
		Fuels struct {
			BD struct {
				Total int `json:"total"`
			} `json:"BD"`
			E85 struct {
				Total int `json:"total"`
			} `json:"E85"`
			ELEC struct {
				Total    int `json:"total"`
				Stations struct {
					Total int `json:"total"`
				} `json:"stations"`
			} `json:"ELEC"`
			HY struct {
				Total int `json:"total"`
			} `json:"HY"`
			LNG struct {
				Total int `json:"total"`
			} `json:"LNG"`
			CNG struct {
				Total int `json:"total"`
			} `json:"CNG"`
			LPG struct {
				Total int `json:"total"`
			} `json:"LPG"`
		} `json:"fuels"`
	} `json:"station_counts"`
	FuelStations []struct {
		AccessCode               string      `json:"access_code"`
		AccessDaysTime           string      `json:"access_days_time"`
		AccessDetailCode         interface{} `json:"access_detail_code"`
		CardsAccepted            interface{} `json:"cards_accepted"`
		DateLastConfirmed        string      `json:"date_last_confirmed"`
		ExpectedDate             interface{} `json:"expected_date"`
		FuelTypeCode             string      `json:"fuel_type_code"`
		GroupsWithAccessCode     string      `json:"groups_with_access_code"`
		ID                       int         `json:"id"`
		OpenDate                 string      `json:"open_date"`
		OwnerTypeCode            string      `json:"owner_type_code"`
		StatusCode               string      `json:"status_code"`
		StationName              string      `json:"station_name"`
		StationPhone             interface{} `json:"station_phone"`
		UpdatedAt                time.Time   `json:"updated_at"`
		FacilityType             string      `json:"facility_type"`
		GeocodeStatus            string      `json:"geocode_status"`
		Latitude                 float64     `json:"latitude"`
		Longitude                float64     `json:"longitude"`
		City                     string      `json:"city"`
		IntersectionDirections   interface{} `json:"intersection_directions"`
		Plus4                    interface{} `json:"plus4"`
		State                    string      `json:"state"`
		StreetAddress            string      `json:"street_address"`
		Zip                      string      `json:"zip"`
		Country                  string      `json:"country"`
		BdBlends                 interface{} `json:"bd_blends"`
		CngDispenserNum          interface{} `json:"cng_dispenser_num"`
		CngFillTypeCode          interface{} `json:"cng_fill_type_code"`
		CngPsi                   interface{} `json:"cng_psi"`
		CngRenewableSource       interface{} `json:"cng_renewable_source"`
		CngTotalCompression      interface{} `json:"cng_total_compression"`
		CngTotalStorage          interface{} `json:"cng_total_storage"`
		CngVehicleClass          interface{} `json:"cng_vehicle_class"`
		E85BlenderPump           interface{} `json:"e85_blender_pump"`
		E85OtherEthanolBlends    interface{} `json:"e85_other_ethanol_blends"`
		EvConnectorTypes         []string    `json:"ev_connector_types"`
		EvDcFastNum              int         `json:"ev_dc_fast_num"`
		EvLevel1EvseNum          interface{} `json:"ev_level1_evse_num"`
		EvLevel2EvseNum          int         `json:"ev_level2_evse_num"`
		EvNetwork                string      `json:"ev_network"`
		EvNetworkWeb             interface{} `json:"ev_network_web"`
		EvOtherEvse              interface{} `json:"ev_other_evse"`
		EvPricing                interface{} `json:"ev_pricing"`
		EvRenewableSource        interface{} `json:"ev_renewable_source"`
		HyIsRetail               interface{} `json:"hy_is_retail"`
		HyPressures              interface{} `json:"hy_pressures"`
		HyStandards              interface{} `json:"hy_standards"`
		HyStatusLink             interface{} `json:"hy_status_link"`
		LngRenewableSource       interface{} `json:"lng_renewable_source"`
		LngVehicleClass          interface{} `json:"lng_vehicle_class"`
		LpgPrimary               interface{} `json:"lpg_primary"`
		LpgNozzleTypes           interface{} `json:"lpg_nozzle_types"`
		NgFillTypeCode           interface{} `json:"ng_fill_type_code"`
		NgPsi                    interface{} `json:"ng_psi"`
		NgVehicleClass           interface{} `json:"ng_vehicle_class"`
		AccessDaysTimeFr         interface{} `json:"access_days_time_fr"`
		IntersectionDirectionsFr interface{} `json:"intersection_directions_fr"`
		BdBlendsFr               interface{} `json:"bd_blends_fr"`
		GroupsWithAccessCodeFr   string      `json:"groups_with_access_code_fr"`
		EvPricingFr              interface{} `json:"ev_pricing_fr"`
	} `json:"fuel_stations"`
}
