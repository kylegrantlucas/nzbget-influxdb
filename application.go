package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/urfave/cli"
)

type NZBFileGroups struct {
	Result []struct {
		ActiveDownloads  int           `json:"ActiveDownloads"`
		Category         string        `json:"Category"`
		CriticalHealth   int           `json:"CriticalHealth"`
		DeleteStatus     string        `json:"DeleteStatus"`
		Deleted          bool          `json:"Deleted"`
		DestDir          string        `json:"DestDir"`
		DownloadTimeSec  int           `json:"DownloadTimeSec"`
		DownloadedSizeHi int           `json:"DownloadedSizeHi"`
		DownloadedSizeLo int           `json:"DownloadedSizeLo"`
		DownloadedSizeMB int           `json:"DownloadedSizeMB"`
		DupeKey          string        `json:"DupeKey"`
		DupeMode         string        `json:"DupeMode"`
		DupeScore        int           `json:"DupeScore"`
		ExParStatus      string        `json:"ExParStatus"`
		ExtraParBlocks   int           `json:"ExtraParBlocks"`
		FailedArticles   int           `json:"FailedArticles"`
		FileCount        int           `json:"FileCount"`
		FileSizeHi       int           `json:"FileSizeHi"`
		FileSizeLo       int           `json:"FileSizeLo"`
		FileSizeMB       int           `json:"FileSizeMB"`
		FinalDir         string        `json:"FinalDir"`
		FirstID          int           `json:"FirstID"`
		Health           int           `json:"Health"`
		Kind             string        `json:"Kind"`
		LastID           int           `json:"LastID"`
		Log              []interface{} `json:"Log"`
		MarkStatus       string        `json:"MarkStatus"`
		MaxPostTime      int           `json:"MaxPostTime"`
		MaxPriority      int           `json:"MaxPriority"`
		MessageCount     int           `json:"MessageCount"`
		MinPostTime      int           `json:"MinPostTime"`
		MinPriority      int           `json:"MinPriority"`
		MoveStatus       string        `json:"MoveStatus"`
		NZBFilename      string        `json:"NZBFilename"`
		NZBID            int           `json:"NZBID"`
		NZBName          string        `json:"NZBName"`
		NZBNicename      string        `json:"NZBNicename"`
		ParStatus        string        `json:"ParStatus"`
		ParTimeSec       int           `json:"ParTimeSec"`
		Parameters       []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Parameters"`
		PausedSizeHi       int           `json:"PausedSizeHi"`
		PausedSizeLo       int           `json:"PausedSizeLo"`
		PausedSizeMB       int           `json:"PausedSizeMB"`
		PostInfoText       string        `json:"PostInfoText"`
		PostStageProgress  int           `json:"PostStageProgress"`
		PostStageTimeSec   int           `json:"PostStageTimeSec"`
		PostTotalTimeSec   int           `json:"PostTotalTimeSec"`
		RemainingFileCount int           `json:"RemainingFileCount"`
		RemainingParCount  int           `json:"RemainingParCount"`
		RemainingSizeHi    int           `json:"RemainingSizeHi"`
		RemainingSizeLo    int           `json:"RemainingSizeLo"`
		RemainingSizeMB    int           `json:"RemainingSizeMB"`
		RepairTimeSec      int           `json:"RepairTimeSec"`
		ScriptStatus       string        `json:"ScriptStatus"`
		ScriptStatuses     []interface{} `json:"ScriptStatuses"`
		ServerStats        []struct {
			FailedArticles  int `json:"FailedArticles"`
			ServerID        int `json:"ServerID"`
			SuccessArticles int `json:"SuccessArticles"`
		} `json:"ServerStats"`
		Status          string `json:"Status"`
		SuccessArticles int    `json:"SuccessArticles"`
		TotalArticles   int    `json:"TotalArticles"`
		URL             string `json:"URL"`
		UnpackStatus    string `json:"UnpackStatus"`
		UnpackTimeSec   int    `json:"UnpackTimeSec"`
		URLStatus       string `json:"UrlStatus"`
	} `json:"result"`
	Version string `json:"version"`
}

type NZBStatus struct {
	Result struct {
		ArticleCacheHi      int  `json:"ArticleCacheHi"`
		ArticleCacheLo      int  `json:"ArticleCacheLo"`
		ArticleCacheMB      int  `json:"ArticleCacheMB"`
		AverageDownloadRate int  `json:"AverageDownloadRate"`
		DaySizeHi           int  `json:"DaySizeHi"`
		DaySizeLo           int  `json:"DaySizeLo"`
		DaySizeMB           int  `json:"DaySizeMB"`
		Download2Paused     bool `json:"Download2Paused"`
		DownloadLimit       int  `json:"DownloadLimit"`
		DownloadPaused      bool `json:"DownloadPaused"`
		DownloadRate        int  `json:"DownloadRate"`
		DownloadTimeSec     int  `json:"DownloadTimeSec"`
		DownloadedSizeHi    int  `json:"DownloadedSizeHi"`
		DownloadedSizeLo    int  `json:"DownloadedSizeLo"`
		DownloadedSizeMB    int  `json:"DownloadedSizeMB"`
		FeedActive          bool `json:"FeedActive"`
		ForcedSizeHi        int  `json:"ForcedSizeHi"`
		ForcedSizeLo        int  `json:"ForcedSizeLo"`
		ForcedSizeMB        int  `json:"ForcedSizeMB"`
		FreeDiskSpaceHi     int  `json:"FreeDiskSpaceHi"`
		FreeDiskSpaceLo     int  `json:"FreeDiskSpaceLo"`
		FreeDiskSpaceMB     int  `json:"FreeDiskSpaceMB"`
		MonthSizeHi         int  `json:"MonthSizeHi"`
		MonthSizeLo         int  `json:"MonthSizeLo"`
		MonthSizeMB         int  `json:"MonthSizeMB"`
		NewsServers         []struct {
			Active bool `json:"Active"`
			ID     int  `json:"ID"`
		} `json:"NewsServers"`
		ParJobCount      int  `json:"ParJobCount"`
		PostJobCount     int  `json:"PostJobCount"`
		PostPaused       bool `json:"PostPaused"`
		QueueScriptCount int  `json:"QueueScriptCount"`
		QuotaReached     bool `json:"QuotaReached"`
		RemainingSizeHi  int  `json:"RemainingSizeHi"`
		RemainingSizeLo  int  `json:"RemainingSizeLo"`
		RemainingSizeMB  int  `json:"RemainingSizeMB"`
		ResumeTime       int  `json:"ResumeTime"`
		ScanPaused       bool `json:"ScanPaused"`
		ServerPaused     bool `json:"ServerPaused"`
		ServerStandBy    bool `json:"ServerStandBy"`
		ServerTime       int  `json:"ServerTime"`
		ThreadCount      int  `json:"ThreadCount"`
		UpTimeSec        int  `json:"UpTimeSec"`
		URLCount         int  `json:"UrlCount"`
	} `json:"result"`
	Version string `json:"version"`
}

func main() {
	log.SetOutput(os.Stdout)

	// setting up cli settings
	app := cli.NewApp()
	app.Name = "nzbget-influxdb"
	app.Usage = "NZBGet -> InfluxDB ingestion daemon"
	app.Author = "Kyle Lucas"
	// app.Version = Version

	// setup cli flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "nzbURL, nurl",
			Usage: "Use a specific server",
		},
		cli.StringFlag{
			Name:  "nzbUsername, nu",
			Usage: "Use a specific server",
		},
		cli.StringFlag{
			Name:  "nzbScheme, ns",
			Value: "http",
			Usage: "Specify the scheme of the nzb server (http, https)",
		},
		cli.StringFlag{
			Name:  "nzbPassword, np",
			Usage: "Use a specific server",
		},
		cli.StringFlag{
			Name:  "influxUsername, iu",
			Usage: "The username for the influxDB instance",
		},
		cli.StringFlag{
			Name:  "influxPasword, ip",
			Usage: "The password for the influxDB instance",
		},
		cli.StringFlag{
			Name:  "influxDB, db",
			Usage: "The name for the influxDB database",
		},
		cli.StringFlag{
			Name:  "influxURL, iurl",
			Value: "http://localhost:8086",
			Usage: "The name for the influxDB database",
		},
		cli.IntFlag{
			Name:  "interval, i",
			Value: 20,
			Usage: "The amount of time in minutes to wait between speedtest runs",
		},
	}

	app.Action = func(c *cli.Context) {
		baseURL := c.String("nzbScheme") +
			"://" + c.String("nzbUsername") +
			":" +
			c.String("nzbPassword") +
			"@" +
			c.String("nzbURL") +
			"/jsonrpc"

		client := &http.Client{}
		db, err := influxDBClient(c.String("influxURL"), c.String("influxUsername"), c.String("influxPassword"))
		if err != nil {
			log.Printf("error connecting to influxdb: %v", err)
		}

		for {
			status, err := getStatus(client, baseURL)
			if err != nil {
				log.Printf("error getting nzbget status: %v", err)
			}

			log.Printf("result: %v", status)

			fileGroups, err := getFileGroups(client, baseURL)
			if err != nil {
				log.Printf("error getting nzbget status: %v", err)
			}

			log.Printf("result: %v", fileGroups)

			err = writeStatusMetrics(db, c.String("influxDB"), status)
			if err != nil {
				log.Printf("error writing status to influxdb: %v", err)
			}
			err = writeFileGroupsMetrics(db, c.String("influxDB"), fileGroups)
			if err != nil {
				log.Printf("error writing filegroups to influxdb: %v", err)
			}

			<-time.After(time.Duration(c.Int("interval")) * time.Second)
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("error running application: %v", err)
	}
}

func getStatus(client *http.Client, baseURL string) (NZBStatus, error) {
	var status NZBStatus
	err := getRequest(client, baseURL, "status", &status)
	if err != nil {
		log.Printf("error getting nzbget status: %v", err)
	}

	return status, nil
}

func getFileGroups(client *http.Client, baseURL string) (NZBFileGroups, error) {
	var fileGroups NZBFileGroups
	err := getRequest(client, baseURL, "listgroups", &fileGroups)
	if err != nil {
		log.Printf("error getting nzbget file groups: %v", err)
	}

	return fileGroups, nil
}

func getRequest(
	client *http.Client,
	baseURL string,
	endpoint string,
	responseStruct interface{},
) error {
	req, err := http.NewRequest("GET", baseURL+"/"+endpoint, nil)
	if err != nil {
		log.Printf("error building nzbget request: %v", err)
		return err
	}

	result, err := client.Do(req)
	if err != nil {
		log.Printf("error getting nzbget endpoint: %v", err)
		return err
	}
	defer result.Body.Close()

	err = json.NewDecoder(result.Body).Decode(responseStruct)
	if err != nil {
		log.Printf("error unmarshaling nzbget status: %v", err)
		return err
	}

	return nil
}

func influxDBClient(url string, username string, password string) (client.Client, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     url,
		Username: username,
		Password: password,
	})

	return c, err
}

func writeStatusMetrics(c client.Client, database string, status NZBStatus) error {
	tags := map[string]string{}
	fields := map[string]interface{}{}

	fields["download_rate"] = status.Result.DownloadRate
	fields["average_download_rate"] = status.Result.AverageDownloadRate
	fields["remaining_size_mb"] = status.Result.RemainingSizeMB
	fields["monthly_download_size_mb"] = status.Result.MonthSizeMB
	return writeMetrics(c, database, "status", tags, fields)
}

func writeFileGroupsMetrics(c client.Client, database string, fileGroups NZBFileGroups) error {
	tags := map[string]string{}
	fields := map[string]interface{}{}

	fields["count"] = len(fileGroups.Result)

	return writeMetrics(c, database, "filegroups", tags, fields)
}

func writeMetrics(c client.Client, database string, measurement string, tags map[string]string, fields map[string]interface{}) error {
	bp, err := client.NewBatchPoints(
		client.BatchPointsConfig{
			Database:  database,
			Precision: "s",
		},
	)
	if err != nil {
		return err
	}

	point, err := client.NewPoint(
		measurement,
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		return err
	}

	bp.AddPoint(point)

	err = c.Write(bp)

	return err
}
