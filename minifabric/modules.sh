cd chaincode
cd academicRecords/go
go mod init academicRecords.go
go mod tidy
go mod vendor

cd ../..

cd decree/go
go mod init decree.go
go mod tidy
go mod vendor

cd ../..

cd registerBook/go
go mod init registerBook.go
go mod tidy
go mod vendor

cd ../..

cd XMLog/go
go mod init XMLog.go
go mod tidy
go mod vendor

cd ../..

mkdir academicRecords/go/vendor/errorMessages
mkdir decree/go/vendor/errorMessages
mkdir registerBook/go/vendor/errorMessages
mkdir XMLog/go/vendor/errorMessages

cp -R assetlib academicRecords/go/vendor
cp -R assetlib decree/go/vendor
cp -R assetlib registerBook/go/vendor
cp -R assetlib XMLog/go/vendor