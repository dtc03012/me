package handler

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"io/ioutil"
	"mime/multipart"
)

func (f *FileServer) UploadFileToDrive(ctx context.Context, driveSrv *drive.Service, file multipart.File, filename string) error {

	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	driveFile, err := driveSrv.Files.Create(&drive.File{Name: filename}).Media(file).Do()
	if err != nil {
		return err
	}

	err = f.db.InsertFileInfo(ctx, tx, filename, driveFile.Id)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (f *FileServer) GetFileFromDrive(ctx context.Context, driveSrv *drive.Service, fileName string) ([]byte, error) {

	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil
	}
	defer tx.Rollback()

	fileId, err := f.db.GetFileId(ctx, tx, fileName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Printf("file Id: %s\n", fileId)

	res, err := driveSrv.Files.Get(fileId).Download()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("error reading download %v", err)
	}

	return body, nil
}
