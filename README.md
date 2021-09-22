# Phone Numbers

Phone Numbers is a single page application that lists and categorizes customer data by country phone numbers.

Phone numbers are categorized by country, validity state, country code and number.

This repository includes the backend side of the application.

You can find the frontend side of the application in the following [repository](https://github.com/imarco20/phone-numbers-frontend).

## How to run the application

### 1. Using Terminal:

- Navigate to the directory of the application on your computer, and run the following command
  ```go run ./cmd/api -port=port_of_your_choice```
- Then open your favorite browser and navigate to "http://localhost:8080". By default, the application runs on Port
  8080, but you can change this to run on the Port of your choice, by specifying it in the above command.

### 2. Using Makefile Commands:

- First, you should install Make on your machine
    ```
    # for linux users
    $ sudo apt install make
    
    # for macOS users
    $ brew install make
    
    # for windows users (using Chocolately package)
    > choco install make
    ```
- After installing make, you'll have access to all commands specified in the project's Makefile
- So, to run the applications, just enter:
  ```make run```

