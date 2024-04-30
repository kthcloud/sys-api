echo "Installing swaggo/swag"
go install github.com/swaggo/swag/cmd/swag@latest

# You might need to "source ./scripts/path-cmd.sh"
swag init --dir ../  -o ../docs/api