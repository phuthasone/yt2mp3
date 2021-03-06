package yt2mp3

import "github.com/otiai10/yt2mp3/factory"

type Converter struct {
	Client Client
}
type MyError struct {
	message string
}

func (e MyError) Error() string {
	return e.message
}
func NewError(message string) error {
	return MyError{
		message,
	}
}

func Init(args ...interface{}) (converter *Converter, err error) {
	if len(args) > 0 {
		if dummyClient, ok := args[0].(*DownloadClient); ok {
			converter = &Converter{
				dummyClient,
			}
		} else {
			err = NewError(
				"Invalid argument for Init: type `DownloadClient` required",
			)
		}
		return
	}
	err = CheckEnv()
	if err == nil {
		converter = &Converter{
			NewDownloadClient(),
		}
	}
	return
}

func CheckEnv() (err error) {
	return err
}
func (c Converter) SetOpt(key string, value interface{}) (ok bool, err error) {
	ok, err = c.Client.SetOpt(key, value)
	return
}
func (c Converter) Vid2mp3(vid string) (fpath string, err error) {
	// TODO: Invalid Vid Format Error (for example)
	fpath, err = c.Client.Execute(vid)
	return
}

func (c Converter) Url2mp3(url string) (fpath string, err error) {
	// TODO: error handling
	vid, _ := factory.Url2vid(url)
	fpath, err = c.Client.Execute(vid)
	return
}
