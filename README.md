***Weather Information***

This service levareges the public api (https://openweathermap.org) to get current weather information for the lat & lon being passed to the service

**How to run this service locally?**
Copy or clone the repository to any place in your disk and run `go run main.go` under main directory. The project uses go modules so its not requiured to place it under GOPATH folder.

**Any config required?**
set environment variable `API_KEY=xxx` inorder to call the public api (https://openweathermap.org)