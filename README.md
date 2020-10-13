# api_backend_app_golang
<h3>Golang App which creates REST APIs which are JWT Authenticated.</h3>

Currently, it supports 3 REST API endpoints:
- ```/register```  - User Registration.
- ```/login```     - Logs in User.
- ```/protected``` - registered User with correct JWT Token can only access this endpoint.

<h3>How it works?</h3>

- ```/login``` Endpoint gives JWT (Json Web Token) token response to registered users.
- ```/protected``` Endpoint only needs JWT token in authentication header.

