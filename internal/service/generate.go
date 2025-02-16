package service

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i AuthService -o ./mocks/ -s "_mimimock.go"
//go:generate minimock -i TransactionService -o ./mocks/ -s "_mimimock.go"
//go:generate minimock -i InfoService -o ./mocks/ -s "_mimimock.go"
//go:generate minimock -i BuyingService -o ./mocks/ -s "_mimimock.go"
