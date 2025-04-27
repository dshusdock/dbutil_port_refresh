package constants

// /////////////TOP LEVEL SIDE MENU///////////////
type side_nav_btn_lbl struct {
	SYSTEM          string
	USER            string
	RECORDING       string
	BUTTON          string
	RESOURCE_AOR    string
	OPEN_CONNECTION string
	LINE            string
	ZONE            string
}

func SIDE_NAV_BTN_LBL() *side_nav_btn_lbl {
	return &side_nav_btn_lbl{
		SYSTEM:          "System",
		USER:            "User",
		RECORDING:       "Recording",
		BUTTON:          "Button",
		RESOURCE_AOR:    "Resource AOR",
		OPEN_CONNECTION: "Open Connection",
		LINE:            "Line",
		ZONE:            "Zone",
	}
}

// /////////////SYSTEM///////////////
type SysSubBtnLbl struct {
	ENTERPISE_INFO         string
	ZONE_INFO              string
	CCM_INFO               string
	MEDIA_MGR_INFO         string
	MEDIA_GWY_INFO         string
	DEVICE_ZONE_INFO       string
	IQMAX_TURRET_INVENTORY string
	TURRET_INFO            string
	JOB_DETAILS_INFO       string
	CDI_COUNTS             string
	LICENSE_INFO           string
}

func SYS_SUB_BTN_LBL() *SysSubBtnLbl {
	return &SysSubBtnLbl{
		ENTERPISE_INFO:         "Enterprise Info",
		ZONE_INFO:              "Zone Info",
		CCM_INFO:               "CCM Info",
		MEDIA_MGR_INFO:         "Media Manager Info",
		MEDIA_GWY_INFO:         "Media Gateway Info",
		DEVICE_ZONE_INFO:       "Device (BCP/Deploy/Current) Zone Info",
		IQMAX_TURRET_INVENTORY: "IQMAX Turret Inventory",
		TURRET_INFO:            "Turret Info",
		JOB_DETAILS_INFO:       "Job Details Info",
		CDI_COUNTS:             "CDI Counts",
		LICENSE_INFO:           "License Info",
	}
}

// /////////////USER///////////////
type user_sub_btn_lbl struct {
	USER_INFO                string
	COMMUNICATION_HISTORY    string
	JOB_EXECUTION_EVENT      string
	JOB_SUMMARY              string
	PERSONAL_EXTENSION       string
	PERSONALDIRNAMES_INFO    string
	USERCDIWITHNOUSERID_INFO string
	CALLS_PER_DAY            string
	CALLS_PER_DAY_PER_USER   string
}

func USER_SUB_BTN_LBL() *user_sub_btn_lbl {
	return &user_sub_btn_lbl{
		USER_INFO:                "User Info",
		COMMUNICATION_HISTORY:    "Communication History",
		JOB_EXECUTION_EVENT:      "Job Execution Event",
		JOB_SUMMARY:              "Job Summary",
		PERSONAL_EXTENSION:       "Personal Extension",
		PERSONALDIRNAMES_INFO:    "PersonalDirNamesInfo",
		USERCDIWITHNOUSERID_INFO: "UserCDIwithnoUserId Info",
		CALLS_PER_DAY:            "Calls Per Day",
		CALLS_PER_DAY_PER_USER:   "Calls Per Day Per User",
	}
}

// /////////////RECORDING///////////////
type rcrd_sub_btn_lbl struct {
	RECORDING_MIX              string
	RECORDING_MIX_WITH_MASK    string
	RECORDING_MIX_LOGONSESSION string
}

func RCRD_SUB_BTN_LBL() *rcrd_sub_btn_lbl {
	return &rcrd_sub_btn_lbl{
		RECORDING_MIX:              "Recording Mix",
		RECORDING_MIX_WITH_MASK:    "Recording Mix with Mask",
		RECORDING_MIX_LOGONSESSION: "Recording Mix LogOnSession",
	}
}

///////////////SYSTEM///////////////

type button_sub_btn_lbl struct {
	BUTTON_RESOURCE_APPEARANCE    string
	BUTTON_INFO                   string
	ARD_BUTTON_INFO               string
	FIND_NON_600BUTTONSUSERS_INFO string
	DUPLICATE_BUTTON_INFO         string
	UNLINKED_BUTTONS              string
	MISSING_DUPE_BUTTONS          string
}

func BTN_SUB_BTN_LBL() *button_sub_btn_lbl {
	return &button_sub_btn_lbl{
		BUTTON_RESOURCE_APPEARANCE:    "Button Resource Appearance",
		BUTTON_INFO:                   "Button Info",
		ARD_BUTTON_INFO:               "ARD Button Info",
		FIND_NON_600BUTTONSUSERS_INFO: "Find Non-600ButtonsUsers Info",
		DUPLICATE_BUTTON_INFO:         "DuplicateButton Info",
		UNLINKED_BUTTONS:              "Unlinked Buttons",
		MISSING_DUPE_BUTTONS:          "Missing/Dupe Buttons",
	}
}

///////////////SYSTEM///////////////

type rsrc_aor_sub_btn_lbl struct {
	RSRC_AOR_ON_SPKR_WITH_LINETYPES     string
	RSRC_AOR_ON_BTNS_WITH_LINETYPES     string
	USER_SPKR_DEVICE_RSRCAOR_INFO       string
	BTN_COUNT_PERZONE_FOR_EACH_RSRCAOR  string
	SPKR_COUNT_PERZONE_FOR_EACH_RSRCAOR string
}

func RSRC_AOR_SUB_BTN_LBL() *rsrc_aor_sub_btn_lbl {
	return &rsrc_aor_sub_btn_lbl{
		RSRC_AOR_ON_SPKR_WITH_LINETYPES:     "ResAOR On Spkr with LineTypes",
		RSRC_AOR_ON_BTNS_WITH_LINETYPES:     "ResAOR On Buttons with LineTypes",
		USER_SPKR_DEVICE_RSRCAOR_INFO:       "UserSpkrsDeviceResouceAOR_Info",
		BTN_COUNT_PERZONE_FOR_EACH_RSRCAOR:  "ButtonCount perZone for each ResAOR",
		SPKR_COUNT_PERZONE_FOR_EACH_RSRCAOR: "SpeakerCount perZone for each ResAOR",
	}
}

///////////////SYSTEM///////////////

type opencnx_sub_btn_lbl struct {
	OCC_ON_SPKR                      string
	OCC_ON_BTNS                      string
	ACTIVE_OCC_ON_BTNS               string
	OCCSPKR_WITH_NO_LISTEN_PERM_INFO string
	OCCSPKR_WITH_NO_SPKR_PERM_INFO   string
	OCCBTNS_WITH_NO_LISTEN_PERM_INFO string
	OCCBTNS_WITH_NO_SPKR_PERM_INFO   string
}

func OPENCNX_SUB_BTN_LBL() *opencnx_sub_btn_lbl {
	return &opencnx_sub_btn_lbl{
		OCC_ON_SPKR: "OCC On Spkr",
		OCC_ON_BTNS: "OCC On Buttons",
		ACTIVE_OCC_ON_BTNS: "Acive OCC on Spkr",
		OCCSPKR_WITH_NO_LISTEN_PERM_INFO: "OCC SpkrWithNoListnPerm Info",
		OCCSPKR_WITH_NO_SPKR_PERM_INFO: "OCC SpkrWithNoSpkrPerm Info",
		OCCBTNS_WITH_NO_LISTEN_PERM_INFO: "OCC ButtonsWithNoListnPerm Info",
		OCCBTNS_WITH_NO_SPKR_PERM_INFO: "OCC ButtonsWithNoListnPerm Info",
	}
}

///////////////SYSTEM///////////////
/*
const (
	ENTERPISE_INFO         = "Enterprise Info"
	ZONE_INFO              = "Zone Info"
	CCM_INFO               = "CCM Info"
	MEDIA_MGR_INFO         = "Media Manager Info"
	MEDIA_GWY_INFO         = "Media Gateway Info"
	DEVICE_ZONE_INFO       = "Device (BCP/Deploy/Current) Zone Info"
	IQMAX_TURRET_INVENTORY = "IQMAX Turret Inventory"
	TURRET_INFO            = "Turret Info"
	JOB_DETAILS_INFO       = "Job Details Info"
	CDI_COUNTS             = "CDI Counts"
	LICENSE_INFO           = "License Info"
)
*/
