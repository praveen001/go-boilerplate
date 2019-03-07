package models

// ResolutionInfo .
type ResolutionInfo struct {
	Resolution           string
	FrameRate            float64
	Interlaced           bool
	VideoMode            string
	MaxGfxArea           int
	AdditionalResolution []string
}

//
var (
	ResolutionMap = map[string]*ResolutionInfo{
		"HD 1080i 29.97 Hz": &ResolutionInfo{
			Resolution: "1920x1080",
			FrameRate:  29.97002997002997,
			Interlaced: true,
			VideoMode:  "1080i60",
			MaxGfxArea: 20,
		},
		"HD 1080p 29.97 Hz": &ResolutionInfo{
			Resolution: "1920x1080",
			FrameRate:  29.97002997002997,
			Interlaced: false,
			VideoMode:  "1080p2997",
			MaxGfxArea: 20,
		},
		"HD 1080i 50 Hz": &ResolutionInfo{
			Resolution: "1920x1080",
			FrameRate:  25,
			Interlaced: true,
			VideoMode:  "1080i50",
			MaxGfxArea: 20,
		},
		"HD 1080i 60 Hz": &ResolutionInfo{
			Resolution: "1920x1080",
			FrameRate:  29.97002997002997,
			Interlaced: true,
			VideoMode:  "1080i60",
			MaxGfxArea: 20,
		},
		"SD NTSC (29.97 Hz)": &ResolutionInfo{
			Resolution: "720x480",
			FrameRate:  29.97002997002997,
			Interlaced: true,
			VideoMode:  "NTSC",
			MaxGfxArea: 100,
		},
		"SD PAL (25 Hz)": &ResolutionInfo{
			Resolution:           "720x576",
			FrameRate:            25,
			Interlaced:           true,
			VideoMode:            "PAL",
			MaxGfxArea:           100,
			AdditionalResolution: []string{"720x608"},
		},
		"720P 59.94Hz": &ResolutionInfo{
			Resolution: "1280x720",
			FrameRate:  59.94005994005994,
			Interlaced: false,
			VideoMode:  "720p5994",
			MaxGfxArea: 20,
		},
		"720P 29.97Hz": &ResolutionInfo{
			Resolution: "1280x720",
			FrameRate:  29.97002997002997,
			Interlaced: false,
			VideoMode:  "720p2997",
			MaxGfxArea: 20,
		},
		"720P 50Hz": &ResolutionInfo{
			Resolution: "1280x720",
			FrameRate:  50,
			Interlaced: false,
			VideoMode:  "720p50",
			MaxGfxArea: 20,
		},
		"720P 25Hz": &ResolutionInfo{
			Resolution: "1280x720",
			FrameRate:  25,
			Interlaced: false,
			VideoMode:  "720p25",
			MaxGfxArea: 20,
		},
		"4K 2160p 25 Hz": &ResolutionInfo{
			Resolution: "3840x2160",
			FrameRate:  25,
			Interlaced: false,
			VideoMode:  "2160p25",
			MaxGfxArea: 20,
		},
		"4K 2160p 50 Hz": &ResolutionInfo{
			Resolution: "3840x2160",
			FrameRate:  50,
			Interlaced: false,
			VideoMode:  "2160p50",
			MaxGfxArea: 20,
		},
		"4K 2160P 59.94Hz": &ResolutionInfo{
			Resolution: "3840x2160",
			FrameRate:  59.94005994005994,
			Interlaced: false,
			VideoMode:  "2160p5994",
			MaxGfxArea: 20,
		},
		"4K 2160P 29.97Hz": &ResolutionInfo{
			Resolution: "3840x2160",
			FrameRate:  29.97002997002997,
			Interlaced: false,
			VideoMode:  "2160p2997",
			MaxGfxArea: 20,
		},
	}
)
