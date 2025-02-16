package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i UserRepository -o ./mocks/ -s "_mimimock.go"
//go:generate minimock -i TransactionRepository -o ./mocks/ -s "_mimimock.go"
//go:generate minimock -i MerchRepository -o ./mocks/ -s "_mimimock.go"
//go:generate minimock -i InventoryRepository -o ./mocks/ -s "_mimimock.go"
