package authhelper

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i AuthHelper -o ./mocks/ -s "_mimimock.go"
