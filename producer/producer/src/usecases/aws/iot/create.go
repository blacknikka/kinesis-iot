package iot

type IoTUsecase struct {
	IoT AWSIoT
}

func (aws *IoTUsecase) InitIoT() error {
	err := aws.IoT.Init()
	if err != nil {
		return err
	}

	return nil
}
