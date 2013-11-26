package actions

import (
	"fmt"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/profitbricks"
	"github.com/dynport/gologger"
	"os"
)

const (
	CLI_DATACENTER_ID                 = "-d"
	CLI_NAME                          = "-n"
	CLI_SIZE                          = "-s"
	CLI_IMAGE_ID                      = "--image-id"
	CLI_RAM                           = "--ram"
	CLI_CORES                         = "--cores"
	CLI_OS_TYPE                       = "--os-type"
	CLI_INTERNET_ACCESS               = "--public"
	CLI_LAN_ID                        = "--lan-id"
	CLI_ROLLBACK_SNAPSHOT_STORAGE_ID  = "--storage-id"
	CLI_ROLLBACK_SNAPSHOT_SNAPSHOT_ID = "--snapshot-id"
	USAGE_USE_IDS                     = "ID [ID...]"
)

var (
	logger             = gologger.NewFromEnv()
	CreateStorage      *gocli.Action
	StartServer        *gocli.Action
	StopServer         *gocli.Action
	DeleteServer       *gocli.Action
	DeleteStorage      *gocli.Action
	ListAllImages      *gocli.Action
	ListAllDataCenters *gocli.Action
	CreateServer       *gocli.Action

	ListAllSnapshots *gocli.Action
	RollbackSnapshot *gocli.Action
)

var defaultDataCenterId = os.Getenv("PROFITBRICKS_DEFAULT_DC_ID")

func init() {
	StartServer = &gocli.Action{Handler: StartServerHandler, Description: "Start Server"}
	StopServer = &gocli.Action{Handler: StopServerHandler, Description: "Stop Server"}
	DeleteServer = &gocli.Action{Handler: DeleteServerHandler, Description: "Delete Server"}
	DeleteStorage = &gocli.Action{Handler: DeleteStorageHandler, Description: "Delete Storage"}
	ListAllImages = &gocli.Action{Handler: ListAllImagesHandler, Description: "List Images"}
	ListAllDataCenters = &gocli.Action{Handler: ListAllDataCentersHandler, Description: "List Snapshots"}

}

func dataCenterFlag() *gocli.Flag {
	required := true
	defaultValue := ""
	if defaultDataCenterId != "" {
		required = false
		defaultValue = defaultDataCenterId
	}
	return &gocli.Flag{
		Type:         gocli.STRING,
		Key:          "data_center_id",
		CliFlag:      CLI_DATACENTER_ID,
		Required:     required,
		DefaultValue: defaultValue,
		Description:  "Data Center Id",
	}
}

func init() {
	args := gocli.NewArgs(nil)
	args.RegisterFlag(dataCenterFlag())
	args.RegisterString(CLI_NAME, "name", true, "", "Storage Name")
	args.RegisterInt(CLI_SIZE, "size", true, 0, "Storage Size")
	args.RegisterString(CLI_IMAGE_ID, "image_id", false, "", "Mount Image Id")
	CreateStorage = &gocli.Action{Handler: CreateStorageHandler, Args: args, Description: "Create Storage"}
}

func init() {
	args := gocli.NewArgs(nil)
	args.RegisterFlag(dataCenterFlag())
	args.RegisterString(CLI_NAME, "name", true, "", "Name")
	args.RegisterInt(CLI_RAM, "ram", false, 1024, "RAM")
	args.RegisterInt(CLI_CORES, "cores", false, 1, "Cores")
	args.RegisterString(CLI_OS_TYPE, "os_type", false, "Linux", "OsType")
	args.RegisterBool(CLI_INTERNET_ACCESS, "internet_access", false, false, "Internet Access")
	args.RegisterInt(CLI_LAN_ID, "lan_id", false, 1, "LanId")
	args.RegisterString(CLI_IMAGE_ID, "image_id", false, "", "Image Id")
	CreateServer = &gocli.Action{Handler: CreateServerHandler, Args: args, Description: "Create Server"}
}

func CreateServerHandler(args *gocli.Args) error {
	req := &profitbricks.CreateServerRequest{
		DataCenterId:    args.MustGetString(CLI_DATACENTER_ID),
		ServerName:      args.MustGetString(CLI_NAME),
		Ram:             args.MustGetInt(CLI_RAM),
		Cores:           args.MustGetInt(CLI_CORES),
		OsType:          args.MustGetString(CLI_OS_TYPE),
		InternetAccess:  args.GetBool(CLI_INTERNET_ACCESS),
		LanId:           args.MustGetInt(CLI_LAN_ID),
		BootFromImageId: args.MustGetString(CLI_IMAGE_ID),
	}
	return profitbricks.NewFromEnv().CreateServer(req)
}

func ListAllDataCentersHandler(args *gocli.Args) error {
	dcs, e := profitbricks.NewFromEnv().GetAllDataCenters()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Id", "Name", "Version")
	for _, dc := range dcs {
		table.Add(dc.DataCenterId, dc.DataCenterName, dc.DataCenterVersion)
	}
	fmt.Println(table)
	return nil
}

func ListAllImagesHandler(args *gocli.Args) error {
	client := profitbricks.NewFromEnv()
	images, e := client.GetAllImages()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Id", "Type", "Region", "Name", "Size")
	for _, img := range images {
		table.Add(img.ImageId, img.ImageType, img.Region, img.ImageName, img.ImageSize)
	}
	fmt.Println(table)
	return nil
}

func doSomething(args *gocli.Args, f func(id string) error) error {
	if len(args.Args) == 0 {
		return fmt.Errorf(USAGE_USE_IDS)
	}
	for _, id := range args.Args {
		e := f(id)
		if e != nil {
			return e
		}
	}
	return nil
}

func DeleteStorageHandler(args *gocli.Args) error {
	return doSomething(args, profitbricks.NewFromEnv().DeleteStorage)
}

func DeleteServerHandler(args *gocli.Args) error {
	return doSomething(args, profitbricks.NewFromEnv().DeleteServer)
}

func StopServerHandler(args *gocli.Args) error {
	return doSomething(args, profitbricks.NewFromEnv().StopServer)
}

func StartServerHandler(args *gocli.Args) error {
	return doSomething(args, profitbricks.NewFromEnv().StartServer)
}

func CreateStorageHandler(args *gocli.Args) error {
	req := &profitbricks.CreateStorageRequest{
		DataCenterId: args.MustGetString(CLI_DATACENTER_ID),
		StorageName:  args.MustGetString(CLI_NAME),
		Size:         args.MustGetInt(CLI_SIZE),
		MountImageId: args.MustGetString(CLI_IMAGE_ID),
	}
	logger.Infof("creating storage with %#v", req)
	return profitbricks.NewFromEnv().CreateStorage(req)
}
