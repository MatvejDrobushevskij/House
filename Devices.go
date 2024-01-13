package SmartDevice

import "fmt"

type SmartDevice struct {
	DeviceType string
	DeviceName string
	Brand      string
	Height     float32
	Weight     float32
	Color      string
	Quantity   int
}

func (sd SmartDevice) GetDeviceInfo() string {
	infoString := ""
	infoString += "\t\t\tТип умного устройства: " + sd.DeviceType + "\n" +
		"\t\t\tНаименование: " + fmt.Sprint(sd.DeviceName) + "\n" +
		"\t\t\tБренд: " + fmt.Sprint(sd.Brand) + "\n" +
		"\t\t\tВысота: " + fmt.Sprint(sd.Height) + " см\n" +
		"\t\t\tШирина: " + fmt.Sprint(sd.Weight) + " см\n" +
		"\t\t\tЦвет: " + fmt.Sprint(sd.Color) + "\n"
	return infoString
}

type HomeAppliance struct {
	ApplianceType string
	ApplianceName string
	Manufacturer  string
	Size          float32
	Weight        float32
	Color         string
	Quantity      int
}

func (ha HomeAppliance) GetApplianceInfo() string {
	infoString := ""
	infoString += "\t\t\tТип бытовой техники: " + ha.ApplianceType + "\n" +
		"\t\t\tНаименование: " + fmt.Sprint(ha.ApplianceName) + "\n" +
		"\t\t\tПроизводитель: " + fmt.Sprint(ha.Manufacturer) + "\n" +
		"\t\t\tРазмер: " + fmt.Sprint(ha.Size) + " см\n" +
		"\t\t\tВес: " + fmt.Sprint(ha.Weight) + " кг\n" +
		"\t\t\tЦвет: " + fmt.Sprint(ha.Color) + "\n"
	return infoString
}
