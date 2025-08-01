def run_go_service(name, path):
    local_resource(
        name=name,
        cmd="go run " + path + "/cmd/main.go",
        deps=[path + "/**/*.go"],
        allow_parallel=True,
    )

# Run services
run_go_service("auth_service", "services/auth_service")
run_go_service("users_service", "services/users_service")
run_go_service("api_gateway", "api-gateway")
