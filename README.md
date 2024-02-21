# Example Go API
This repo contains an example go REST API built using the chi module. The API was built following the tutorials in this YouTube video (https://www.youtube.com/watch?v=8uiZC0l4Ajw) and then my own customizations applied as part of the learning process.

# Customizations

On top of the main code built following the tutorial I have also added the below functionality as additional training:

1) /user method which provides outputs on user information (E.g. Username, Fullname etc...)
2) Added a currency field to the main /coins API method
3) Additional logging and logging of API timings to the terminal

# How to use

The API is hosted on http://localhost:8000

The methods:

## Coins:
http://localhost:8000/account/coins/?username=alex

Output: Balance, Currency

## User

Example URI: http://localhost:8000/account/user/?username=alex

Output: Username, FullName, Email, Country