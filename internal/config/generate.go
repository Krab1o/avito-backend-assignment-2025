package config

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i JWTConfig -o ./mocks/ -s "_mimimock.go"
//go:generate minimock -i PGConfig -o ./mocks/ -s "_mimimock.go"
//go:generate minimock -i HTTPConfig -o ./mocks/ -s "_mimimock.go"
